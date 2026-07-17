# F. Imbalance Value of a Tree

[Problem link](https://codeforces.com/problemset/problem/915/F)

time limit per test: 4 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

You are given a tree `T` with `n` vertices. Vertex `i` has value `a_i`.

Define `I(x, y)` as the difference between the maximum and minimum of `a_i` on
the simple path connecting `x` and `y`.

Compute

```text
∑_{1 ≤ i ≤ j ≤ n} I(i, j)
```

(when `i = j`, `I(i, i) = 0`).

## Constraints

- `1 ≤ n ≤ 10^6`
- `1 ≤ a_i ≤ 10^6`
- The given edges form a tree

## Input

```
n
a_1 a_2 ... a_n
x_1 y_1
...
x_{n-1} y_{n-1}
```

## Output

Print one integer — the required sum.

## Sample

### Sample 1

Input:

```
4
2 2 3 1
1 2
1 3
1 4
```

Output:

```
6
```

## Ideas

For a pair of vertices `(x, y)`, let `P(x, y)` be their simple path. Its
contribution can be separated into two parts:

```text
I(x, y) = max(a_v, v in P(x, y)) - min(a_v, v in P(x, y))
```

Therefore, we can calculate the sum of all path maximums and the sum of all
path minimums independently:

```text
answer = sum(f[u]) - sum(g[u])
```

Here, `f[u]` is the contribution assigned to `u` when processing path
maximums, and `g[u]` is the analogous contribution when processing path
minimums.

### Calculate path maximums

Sort the vertices by `a[u]` in nondecreasing order and activate them one by
one. The activated vertices form a forest, whose connected components can be
maintained with a DSU.

When activating `u`, suppose its already-active neighboring components have
sizes:

```text
s1, s2, ..., sk
```

Adding `u` joins these components. Every pair of vertices that becomes
connected at this moment has a path passing through `u`. All vertices that
were activated earlier have values at most `a[u]`, so the maximum value on
each such path is `a[u]`.

Start with a component containing only `u`, so `cur = 1`. When merging an
adjacent component of size `si`, the number of newly connected pairs is:

```text
cur * si
```

One endpoint can be chosen from the current component and the other from the
component being merged. Thus the update is:

```text
f[u] += a[u] * cur * si
cur += si
```

Apply this operation to every active neighboring component of `u`.

We only count pairs with different endpoints. A pair `(u, u)` contributes
`a[u]` to both the maximum sum and the minimum sum, so the two contributions
cancel.

### Calculate path minimums

The calculation of `g[u]` is symmetric. Reset the DSU, sort the vertices by
`a[u]` in nonincreasing order, and activate them one by one.

Now every previously activated vertex has value at least `a[u]`. Therefore,
when two vertices first become connected while activating `u`, the minimum
value on their path is `a[u]`. The same merging formula gives:

```text
g[u] += a[u] * cur * si
cur += si
```

### Equal values

Equal values do not need special handling. Vertices with the same value may
be processed in any fixed order.

Every pair is counted exactly once: at the moment its endpoints first become
connected. If several vertices on that path share the maximum or minimum
value, the processing order merely decides which tied vertex owns the pair's
contribution. Since all tied vertices have the same value, the numerical
contribution is unchanged.

The important point is to activate and merge equal-valued vertices
sequentially. If an entire equal-value group were queried first and merged
only afterward, connections created inside the group would require separate
accounting.

The total complexity is `O(n log n)` for sorting and almost `O(n)` for all
DSU operations. The memory complexity is `O(n)`.
