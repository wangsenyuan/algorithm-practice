# E - Min of Restricted Sum

[Problem link](https://atcoder.jp/contests/abc396/tasks/abc396_e)

**Contest:** [AtCoder Beginner Contest 396](https://atcoder.jp/contests/abc396)

time limit: 3 sec

memory limit: 1024 MiB

score: 450 points

## Problem

You are given integers `N`, `M` and three sequences of length `M`:
`X`, `Y`, and `Z` (each `X_i`, `Y_i` is in `1..N`).

A length-`N` sequence of non-negative integers `A` is **good** iff for every
`i` (`1 <= i <= M`):

```text
A[X_i] XOR A[Y_i] = Z_i
```

Determine whether a good sequence exists. If it does, find one that minimizes
`sum(A)`. Any such minimum-sum sequence is accepted.

## Constraints

- `1 <= N <= 2 * 10^5`
- `0 <= M <= 10^5`
- `1 <= X_i, Y_i <= N`
- `0 <= Z_i <= 10^9`
- All input values are integers

## Input

```text
N M
X_1 Y_1 Z_1
...
X_M Y_M Z_M
```

## Output

If no good sequence exists, print `-1`.

Otherwise print `N` integers `A_1 ... A_N` separated by spaces.

## Samples

### Sample 1

```text
Input
3 2
1 3 4
1 2 3

Output
0 3 4
```

### Sample 2

```text
Input
3 3
1 3 4
1 2 3
2 3 5

Output
-1
```

### Sample 3

```text
Input
5 8
4 2 4
2 3 11
3 4 15
4 5 6
3 2 11
3 3 0
3 1 9
3 4 15

Output
0 2 9 6 0
```

## Solution Summary

Interpret every variable `A_i` as a vertex of an undirected graph. A
constraint

```text
A_x XOR A_y = z
```

becomes an edge between `x` and `y` with label `z`.

### Assign relative XOR values

Process every connected component separately. Choose an arbitrary root and
temporarily assign it the relative value `0`. During DFS, if vertex `u` has
relative value `dp[u]`, an edge `(u, v, z)` forces

```text
dp[v] = dp[u] XOR z.
```

If an already visited vertex is reached with a different forced value, the
constraints are inconsistent, so no good sequence exists.

Otherwise, all relative values in the component are determined. They describe
every solution in that component up to one common XOR constant `C`:

```text
A[u] = dp[u] XOR C.
```

The root's final value is `C`; changing `C` changes every value in the
component while preserving every edge constraint.

### Minimize the sum bit by bit

Different bits of XOR are independent, so choose each bit of `C` separately.
For a fixed bit `d`, count how many component vertices have that bit equal to
`0` or `1` in `dp`:

```text
cnt0 = number of zero bits
cnt1 = number of one bits
```

- If bit `d` of `C` is `0`, the final values contain `cnt1` one bits.
- If bit `d` of `C` is `1`, every bit is flipped, so the final values contain
  `cnt0` one bits.

Choose the option with fewer one bits. The implementation directly constructs
the resulting `A[u]`: it keeps the `dp` bit when `cnt0 >= cnt1`, and flips it
otherwise.

Only bits `0..29` need to be considered because `Z_i <= 10^9 < 2^30`.
All relative values therefore have no higher set bits, and setting any higher
bit in `C` would only increase the sum.

### Correctness Proof

We prove that the algorithm either correctly reports impossibility or returns
a good sequence with minimum sum.

For each DFS tree edge `(u, v, z)`, the algorithm assigns
`dp[v] = dp[u] XOR z`, so `dp[u] XOR dp[v] = z`. If DFS reaches an already
assigned vertex, equality with the newly forced value is necessary and
sufficient for that edge to be compatible with all previous assignments.
Consequently, a mismatch proves that no assignment can satisfy the component.
If there is no mismatch, every edge constraint is satisfied by the relative
values.

For a consistent component, let `r` be its root. In any good sequence, set
`C = A[r]`. Following the constraints along any path from `r` to `u` gives
`A[u] = dp[u] XOR C`. Conversely, applying any common `C` to all relative
values preserves every constraint because

```text
(dp[x] XOR C) XOR (dp[y] XOR C)
    = dp[x] XOR dp[y]
    = Z.
```

Thus the algorithm considers exactly all possible assignments for the
component by choosing `C`.

For each bit, choosing `C` equal to `0` contributes `cnt1 * 2^d` to the sum,
while choosing it equal to `1` contributes `cnt0 * 2^d`. The algorithm chooses
the smaller contribution. Bits are independent, so minimizing every bit
individually minimizes the component's total sum. Connected components are
also independent, so minimizing each component minimizes the sum of the whole
sequence.

Therefore, the algorithm reports `-1` exactly when no good sequence exists;
otherwise, it returns a valid minimum-sum sequence.

### Complexity

Every vertex and edge is processed by DFS, and constructing the answer checks
30 bits per vertex.

- Time: `O(N + M + 30N)`, equivalently `O(N + M)`.
- Space: `O(N + M)`.
