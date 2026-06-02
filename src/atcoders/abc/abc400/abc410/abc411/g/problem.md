# G - Count Cycles

https://atcoder.jp/contests/abc411/tasks/abc411_g

**Time Limit:** 6 sec / **Memory Limit:** 1024 MiB

**Score:** 600 points

## Problem Statement

There is an undirected graph `G` with `N` vertices and `M` edges. The graph has
no self-loops, but it may have multi-edges.

Vertices are numbered from `1` to `N`, and edges are numbered from `1` to `M`.
Edge `i` connects vertices `U_i` and `V_i`.

Find the number of cycles contained in `G`, modulo `998244353`.

More formally, count the subsets:

```text
{e_1, e_2, ..., e_k} subset of {1, 2, ..., M}, with k >= 2
```

such that there exists a permutation `(e'_1, e'_2, ..., e'_k)` of those edges
and a vertex sequence `(v_1, v_2, ..., v_k)` satisfying:

- `v_1, v_2, ..., v_k` are pairwise distinct.
- For every `j`, edge `e'_j` connects
  `v_j` and `v_{(j mod k) + 1}`.

## Constraints

- `2 <= N <= 20`
- `2 <= M <= 2 * 10^5`
- `1 <= U_i < V_i <= N`
- All input values are integers.

## Solution

Since `N <= 20`, use bitmask DP over vertex sets.

First, compress multi-edges into an adjacency matrix:

```text
adj[u][v] = number of edges between u and v
```

If a simple cycle uses an unordered vertex edge `{u, v}`, then there are
`adj[u][v]` choices for the actual input edge. So the number of edge-subset
cycles for one vertex cycle is the product of edge multiplicities along it.

To avoid counting the same vertex set from many starting points, choose the
smallest-numbered vertex in the mask as the start:

```text
st = lowest set bit of mask
```

Define:

```text
dp[mask][v]
```

as the number of paths that:

- start at `st`,
- visit exactly the vertices in `mask`,
- end at vertex `v`,
- only add vertices whose index is greater than `st`.

Initialize:

```text
dp[1 << i][i] = 1
```

For each state, close the path back to `st`:

```text
answer += dp[mask][v] * adj[v][st]
```

Then extend the path by an unused vertex `w > st`:

```text
dp[mask | (1 << w)][w] += dp[mask][v] * adj[v][w]
```

The restriction `w > st` guarantees that `st` is the minimum vertex in the
cycle, so the same cycle is not counted once for every vertex as the start.

This DP also counts the length-2 case where the path is `u -> v` and then closes
with another edge `v -> u`. Such a cycle is valid only when two distinct
parallel edges are chosen, but the DP initially also includes the same edge used
twice. Across all two-vertex masks, the number of invalid same-edge choices is
exactly `M`, so subtract `M`.

Every remaining cycle is counted twice, once in each direction around the
cycle. Therefore:

```text
answer = (answer - M) / 2
```

The division is done modulo `998244353` using the modular inverse of `2`.

The complexity is `O(2^N * N^2)`, which is feasible for `N <= 20`.
