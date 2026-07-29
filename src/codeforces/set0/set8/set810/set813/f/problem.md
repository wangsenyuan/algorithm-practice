# F. Bipartite Checking

[Problem link](https://codeforces.com/problemset/problem/813/F)

**Contest:** [Educational Codeforces Round 22](https://codeforces.com/contest/813)

time limit per test: 6 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

You are given an undirected graph with `n` vertices and initially no edges.
There are `q` queries; each query toggles an undirected edge between `x` and
`y` (add it if absent, remove it if present).

After each query, report whether the graph is bipartite.

## Constraints

- `2 <= n, q <= 10^5`
- `1 <= x_i < y_i <= n`

## Input

The first line contains two integers `n` and `q`.

Each of the next `q` lines contains two integers `x_i` and `y_i` — the edge to
toggle.

## Output

Print `q` lines. The `i`-th line is `YES` if the graph is bipartite after the
`i`-th query, and `NO` otherwise.

## Sample 1

```text
Input
3 5
2 3
1 3
1 2
1 2
1 2

Output
YES
YES
NO
YES
NO
```

## ideas
1. 在一棵树上, 添加了一条边, 且这条边连接的节点, dep[u] & 1 == dep[v] & 1, 那么就造成了一个奇数长度的环, 就不是二部图
2. 添加边的时候, 还是能处理的. 这个时候, 添加这条边合并了两个分组, 那么就去更新小的分组的dep[?]; 这样整体, 摊销是n*lgn
3. 但是删除边呢?
4. 删除边的时候, 有可能要重新计算dep[?], 否则后续的判断就会出错. 但是更新这个似乎是个很麻烦的事情.
5. 如果不更新这个dep[?]. 先只处理添加边的query, 生成一个tree
6. 然后添加一条边的时候, 连接 dep[u] & 1 = dep[v] & 1 就没法作为判断依据了

## editorial

The graph changes online, but all queries are known before we start processing.
Use that fact to turn deletions into intervals of time, then solve the intervals
offline with a rollback DSU.

### 1. Turn toggles into active intervals

For every edge, remember the query index where it was most recently added.
When the same edge appears again at index `i`, it is removed, so it existed on
the half-open interval `[start, i)`. If it remains active after the final
query, its interval is `[start, q)`.

For example, if an edge is toggled on at time `2` and off at time `7`, it is
active after queries `2, 3, 4, 5, 6`, namely on `[2, 7)`.

### 2. Segment tree over query time

Store each active interval in a segment tree over `[0, q)`. An edge is put
only in nodes whose whole segment is contained in its active interval.

For an edge active on `[2, 7)` in a tree over times `0..7`, one possible
decomposition is:

```text
[2, 3)  [3, 4)  [4, 6)  [6, 7)
```

When DFS enters one of those nodes, that edge is guaranteed to be active for
every query represented by the node. Therefore it is safe to add the edge to
DSU before descending and keep it there for the entire subtree.

At a leaf for time `t`, the union of the edge lists on the root-to-leaf path
is exactly the graph after query `t`:

- an active edge has an interval containing `t`, so one of its stored
  segment-tree nodes lies on this path;
- an inactive edge has no stored node containing `t`, so it is absent.

`addInterval` performs this interval decomposition, and `segments[node]`
stores the edges for one time-tree node.

### 3. DSU with color parity

A bipartite graph requires every edge `(u, v)` to satisfy

\[
color(u) \oplus color(v) = 1.
\]

The DSU stores `parity[x]` as

\[
color(x) \oplus color(parent(x)).
\]

Thus `find(x)` returns both the root and `px`, where

\[
px = color(x) \oplus color(root_x).
\]

Suppose we add `(u, v)` and obtain `(ru, pu)` and `(rv, pv)` from `find`.

#### Different components

After attaching `ru` below `rv`, `parity[ru]` must be

\[
color(ru) \oplus color(rv).
\]

Because the new edge requires opposite endpoint colors,

\[
(color(ru) \oplus pu) \oplus (color(rv) \oplus pv) = 1.
\]

Rearranging gives

\[
color(ru) \oplus color(rv) = pu \oplus pv \oplus 1.
\]

This is exactly the assignment in the code:

```go
dsu.parity[ru] = pu ^ pv ^ 1
```

The smaller component is attached to the larger one. There is no path
compression: union by size alone keeps the DSU height `O(log n)`, and every
future mutation remains easy to undo.

#### Same component

If `ru == rv`, the DSU already knows the relative colors of `u` and `v`.

- `pu != pv`: they already have opposite colors, so the edge is compatible.
- `pu == pv`: they have the same color, so the new edge creates an odd cycle.

The variable `bad` counts currently active contradictory edges. The current
graph is bipartite exactly when `bad == 0`.

### 4. Rollback

Before processing a time-tree node, save `snapshot := len(dsu.history)`.
Every successful merge records the old parent, old parent parity, and old
parent-component size. Every contradiction records the old `bad` value.

After the node's subtree is complete, `rollback(snapshot)` pops those records
and restores the DSU to precisely the state it had before entering the node.
This removes all edges that were only valid for that time segment, while
preserving the edges inherited from ancestor segments.

At a leaf, `solve` writes `YES` if `bad == 0`; otherwise it writes `NO`.

### Correctness sketch

Every edge is present in the DSU at exactly the leaves corresponding to its
active interval. At each such leaf, the parity DSU enforces the opposite-color
condition for every active edge and increments `bad` exactly when an odd cycle
is forced. Hence `bad == 0` exactly when the graph at that query is bipartite.

### Complexity

Each active interval is stored in `O(log q)` time-tree nodes. Each insertion,
find, and rollback operation costs `O(log n)` because of union by size. The
total time is `O(q log q log n)` and the auxiliary space is `O(q log q + n)`.
