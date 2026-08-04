# D1. Construct an Array (Easy Version)

[Problem link](https://codeforces.com/problemset/problem/2245/D1)

**Contest:** [Spectral::Cup 2026 Round 3 (Codeforces Round 1110, Div. 1 + Div. 2)](https://codeforces.com/contest/2245)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

This is the easy version of the problem. The difference between the versions is that in
this version, `n ≤ 1000` and `m = n(n + 1) / 2`.

You are given two integers `n` and `m`. You need to construct an integer array `a` of
length `n` that satisfies `m` restrictions. Each restriction can be represented by a
tuple `(o, i, j)` such that `o ∈ {1, 2}` and `1 ≤ i ≤ j ≤ n`:

- If `o = 1`, the sum `a_i + a_j` must be non-negative.
- If `o = 2`, the sum `a_i + a_j` must be negative.

## Constraints

- `1 ≤ t ≤ 10^4`
- `1 ≤ n ≤ 1000`
- `m = n(n + 1) / 2`
- `o ∈ {1, 2}`, `1 ≤ i ≤ j ≤ n`
- Each pair `(i, j)` with `1 ≤ i ≤ j ≤ n` occurs in at most one restriction
- Sum of `m` over all test cases does not exceed `10^6`
- If a valid array exists, there exists one with `|a_i| ≤ 10^9`

## Input

The first line contains the number of test cases `t`.

The first line of each test case contains two integers `n` and `m`.

Each of the next `m` lines contains three integers `o`, `i`, and `j`.

## Output

For each test case, if no such array `a` exists, output `NO`.

Otherwise, first output `YES` on a single line. Then output `n` integers
`a_1, a_2, …, a_n` (`|a_i| ≤ 10^9`).

If there exist multiple arrays satisfying the requirement, you may output any of them.

You can output the answer in any case (upper or lower).

## Example

### Input

```text
6
1 1
1 1 1
1 1
2 1 1
2 3
1 1 1
1 1 2
1 2 2
2 3
1 1 1
1 2 2
2 1 2
3 6
1 1 1
1 1 2
1 1 3
2 2 2
2 2 3
2 3 3
3 6
2 1 1
1 1 2
2 2 3
1 3 3
1 2 2
2 1 3
```

### Output

```text
YES
0
YES
-1
YES
0 0
NO
YES
1 -1 -1
NO
```

## Note

In the first test case, the only restriction is that `a_1 + a_1` is non-negative,
implying that `a_1` is non-negative. Thus, `a_1` can be any non-negative integer.

In the second test case, the only restriction is that `a_1 + a_1` is negative, implying
that `a_1` is negative. Thus, `a_1` can be any negative integer.

In the fourth test case, the first and second restrictions imply that both `a_1` and
`a_2` are non-negative. However, the third restriction forces `a_1 + a_2` to be
negative, which leads to a contradiction.

## Solution

Signs come from the diagonal; mixed pairs become an abs-value order; topo-sort
assigns those abs values.

### 1. Diagonal fixes the sign

For each `i`, the self-constraint `(o, i, i)` decides the sign of `a_i`:

- `o = 1` → `2a_i ≥ 0` → `a_i ≥ 0`
- `o = 2` → `2a_i < 0` → `a_i < 0`

Call these **P** (non-negative) and **N** (negative).

### 2. Same-sign pairs are rigid

- Two values in **P**: sum is always `≥ 0` → that pair **must** be type `1`.
  Type `2` is impossible.
- Two values in **N**: sum is always `< 0` → that pair **must** be type `2`.
  Type `1` is impossible.

So any “two positives but type 2” or “two negatives but type 1” → `NO`.

### 3. Mixed pairs only constrain magnitudes

Take `i ∈ P`, `j ∈ N`, and write `a_i = +x`, `a_j = -y` with `x, y > 0`:

| constraint | meaning |
|---|---|
| type `1` (`a_i + a_j ≥ 0`) | `x ≥ y` → `\|a_i\| ≥ \|a_j\|` |
| type `2` (`a_i + a_j < 0`) | `x < y` → `\|a_i\| < \|a_j\|` |

Mixed edges are just ordering constraints on absolute values.

### 4. Build a graph on abs-order

For every positive `i` and negative `j`:

- type `1` (`\|i\| ≥ \|j\|`): edge `j → i` (give `j` smaller abs first)
- type `2` (`\|i\| < \|j\|`): edge `i → j` (give `i` smaller abs first)

An edge `u → v` means: assign `|a_u|` before `|a_v|` in increasing abs order.

### 5. Topological order → answer

Topo-sort that graph; assign absolute values `1, 2, …, n` in that order:

- node in **P** → `+cnt`
- node in **N** → `-cnt`

- If there is a **cycle**, the abs constraints contradict → `NO`
- If topo finishes all `n` nodes, still **re-check every restriction** (this
  catches bad same-sign pairs that were never put into the graph)

### One-liner

Fix signs from diagonals; turn mixed constraints into a DAG on absolute values;
topo-assign `±1…±n`; verify.
