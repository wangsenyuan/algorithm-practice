# F - Strongly Connected 2 (ABC450)

**Contest:** [ABC450](https://atcoder.jp/contests/abc450) — AtCoder Beginner Contest 450  
**Task:** [https://atcoder.jp/contests/abc450/tasks/abc450_f](https://atcoder.jp/contests/abc450/tasks/abc450_f)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 525 points

## Problem Statement

There is a directed graph with `N` vertices and `N - 1 + M` edges, with edges and
vertices numbered.

- For `1 <= i <= M`, edge `i` is a directed edge from vertex `X_i` to vertex `Y_i`.
- For `1 <= i <= N - 1`, edge `M + i` is a directed edge from vertex `i + 1` to
  vertex `i`.

You may choose some subset (possibly empty) of edges `1, 2, ..., M` and delete them.
There are `2^M` such choices.

Find the number of choices for which the graph after deletion is **strongly
connected**, modulo `998244353`.

## Constraints

- `2 <= N <= 2 * 10^5`
- `1 <= M <= 2 * 10^5`
- `1 <= X_i < Y_i <= N`
- All input values are integers.

## Input

The input is given from standard input in the following format:

```text
N M
X_1 Y_1
X_2 Y_2
:
X_M Y_M
```

## Output

Print the answer.

## Sample Input 1

```text
4 3
1 4
1 3
2 4
```

## Sample Output 1

```text
5
```

The graph is strongly connected when the set of deleted edge indices is one of
`{}`, `{1}`, `{2}`, `{3}`, or `{2, 3}` — five ways in total.

The fixed edges are `4 -> 3`, `3 -> 2`, and `2 -> 1` (edges `M+1`, `M+2`, `M+3`).

## Sample Input 2

```text
10 11
1 4
1 4
3 9
2 5
3 4
9 10
6 9
4 10
1 3
8 10
4 7
```

## Sample Output 2

```text
1297
```

The given graph may contain multi-edges (duplicate pairs `(X_i, Y_i)` are allowed).

## Solution

The fixed edges `i+1 -> i` form a backbone from `N` down to `1`. Every optional edge
goes forward (`X < Y`). Strong connectivity is equivalent to: after deletions, one can
still travel from `1` to `N` using optional edges and backward steps, and return from
`N` to `1`.

### DP meaning

Sort optional edges by `(l, r)` ascending.

Let `dp[v]` be the number of deletion choices among edges processed so far that keep
the “forward reachability” state represented at vertex `v`. Intuitively, `dp[v]`
counts partial schemes where the current “frontier” of forward connectivity ends at
`v`.

- Start with `dp[1] = 1` (only the backbone, nothing processed yet).
- Process each optional edge `(l, r)` in sorted order.

### Transition for one edge `(l, r)`

1. **Multiply `dp[r..N]` by `2`**

   For states with frontier at least `r`, this edge may be deleted or kept without
   forcing a new connection through `r` yet. That gives a factor of `2` for this
   binary choice.

2. **Add `sum(dp[l..r-1])` into `dp[r]`**

   To eventually reach `N`, some schemes must jump from a vertex in `[l, r-1]` to `r`
   using this forward edge (while still being able to use backward edges inside
   `[l, r]`). All weight in `dp[l..r-1]` is transferred to `dp[r]`.

Backward edges let you move from `r` down to `l`, so keeping `(l, r)` connects the
segment even if earlier forward edges were deleted.

### Why sorting works

All edges are processed by increasing right endpoint. When handling `(l, r)`, only
edges with right endpoint `< r` are already fixed, so `dp[l..r-1]` is final for this
step. Multiplying `dp[r..N]` before the addition correctly accounts for later choices
on the current edge.

Multiple edges with the same `(l, r)` or sharing endpoints are handled naturally:
each edge applies its own multiply-then-add update.

### Segment tree

`dp` is stored on `[1..N]` with a lazy segment tree because each step needs:

- range multiply by `2` on `[r, N]`
- range sum on `[l, r-1]`
- point add at `r`

This is `O(M log N)` time and `O(N)` space.

### Answer

After all optional edges are processed, `dp[N]` is the number of deletion sets that
leave the graph strongly connected. Return `dp[N] mod 998244353`.

### Example sanity check (sample 1)

Edges `(1,4), (1,3), (2,4)` after sorting. The DP yields `dp[4] = 5`, matching the
five valid deletion sets `{}`, `{1}`, `{2}`, `{3}`, `{2,3}`.