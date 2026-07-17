# G - Maximize Distance

[Problem link](https://atcoder.jp/contests/abc397/tasks/abc397_g)

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 625 points

## Problem Statement

You are given a directed graph with `N` vertices and `M` edges. The vertices
are numbered `1, 2, ..., N`, and edge `j` goes from vertex `u_j` to vertex
`v_j`. It is guaranteed that vertex `N` is reachable from vertex `1`.

Initially, every edge has weight `0`. Choose exactly `K` of the `M` edges and
change their weights to `1`.

Find the maximum possible shortest distance from vertex `1` to vertex `N`
after choosing the edges.

## Constraints

- `2 <= N <= 30`
- `1 <= K <= M <= 100`
- `1 <= u_j, v_j <= N`
- `u_j != v_j`
- Vertex `N` is reachable from vertex `1`
- All input values are integers

The graph may contain multiple edges between the same pair of vertices.

## Input

```text
N M K
u_1 v_1
u_2 v_2
...
u_M v_M
```

## Output

Print the answer.

## Samples

### Sample 1

Input:

```text
3 3 2
1 2
2 3
1 3
```

Output:

```text
1
```

Choosing edges `1` and `3` makes the shortest distance from vertex `1` to
vertex `3` equal to `1`. It is impossible to make it at least `2`.

### Sample 2

Input:

```text
4 4 3
1 2
1 3
3 2
2 4
```

Output:

```text
2
```

Choosing edges `1`, `2`, and `4` makes the shortest distance from vertex `1`
to vertex `4` equal to `2`. It is impossible to make it at least `3`.

### Sample 3

Input:

```text
2 2 1
1 2
1 2
```

Output:

```text
0
```

This sample demonstrates that the graph may contain multiple edges.

## Solution

### Binary Search on the Answer

For a fixed integer `d`, consider the decision problem:

> What is the minimum number of edges that must be marked so that the shortest
> distance from vertex `1` to vertex `N` is at least `d`?

Let this minimum be `minMarks(d)`. Distance `d` is achievable if and only if

```text
minMarks(d) <= K.
```

The problem requires marking exactly `K` edges, but this inequality is enough.
If only `c < K` edges are required, we can mark any additional `K-c` edges.
Increasing more edge weights from `0` to `1` cannot decrease a shortest
distance.

Feasibility is monotone: if distance `d` is achievable, every smaller distance
is also achievable. The answer is therefore found by binary search. It is at
most `N-1`, because some simple path from vertex `1` to vertex `N` contains at
most `N-1` edges.

The remaining task is to calculate `minMarks(d)`.

### Assigning Levels to Vertices

Assign every vertex `v` an integer level:

```text
0 <= level[v] <= d,
level[1] = 0,
level[N] = d.
```

For every original directed edge `u -> v`, require

```text
level[v] <= level[u] + 1.
```

The meaning of an edge is then:

- If `level[v] <= level[u]`, it can remain unmarked with weight `0`.
- If `level[v] = level[u] + 1`, it must be marked and have weight `1`.
- If `level[v] > level[u] + 1`, the level assignment is invalid because one
  edge cannot have weight greater than `1`.

Thus the cost of a valid level assignment is

```text
the number of edges u -> v with level[v] = level[u] + 1.
```

Why does such an assignment guarantee distance at least `d`? Every path from
vertex `1` to vertex `N` starts at level `0` and ends at level `d`. One edge can
increase the level by at most one, and every level-increasing edge is marked.
Consequently, every path contains at least `d` marked edges.

Conversely, suppose some set of marked edges gives shortest distance at least
`d`. Set

```text
level[v] = min(d, shortest distance from vertex 1 to v).
```

For an edge of weight `w`, shortest distances satisfy
`level[v] <= level[u] + w`. Hence these levels form a valid assignment, and
every edge whose level increases by one is among the marked edges. Minimizing
the level-assignment cost therefore gives exactly `minMarks(d)`.

### Why One Intermediate Vertex Is Insufficient

A transition through one vertex `w`, such as combining a cut from `u` to `w`
with a cut from `w` to `v`, controls only paths that pass through `w`. Other
paths may bypass it. The levels instead create boundaries across the entire
graph, so every possible path must cross enough level boundaries.

### Encoding a Level with Boolean Thresholds

Enumerating all `(d+1)^N` level assignments is impossible. Replace each level
with `d` Boolean threshold variables:

```text
node(v, j) is true iff level[v] >= j+1, for 0 <= j < d.
```

For example, when `d = 3`:

| `level[v]` | Thresholds |
| ---: | :--- |
| `0` | `000` |
| `1` | `100` |
| `2` | `110` |
| `3` | `111` |

Create one minimum-cut node for each threshold. Being on the source side of
the cut means the threshold is true; being on the sink side means it is false.

An implication `A => B` is represented by an edge `A -> B` of capacity
`INF`. The forbidden assignment `A = true, B = false` would cut this edge.
The implementation uses `INF = M+1`, which is larger than the cost of marking
all original edges.

### Edges in the Minimum-Cut Graph

Let `S` and `T` be the source and sink of the flow graph.

#### 1. Fix the endpoint levels

To force `level[1] = 0`, every threshold of vertex `1` must be false:

```text
node(1, j) -> T, capacity INF.
```

To force `level[N] = d`, every threshold of vertex `N` must be true:

```text
S -> node(N, j), capacity INF.
```

#### 2. Keep each threshold sequence monotone

The implication

```text
level[v] >= j+2  =>  level[v] >= j+1
```

is represented by

```text
node(v, j+1) -> node(v, j), capacity INF.
```

This allows only patterns of some `1`s followed by some `0`s, so every cut
chooses exactly one level for each vertex.

#### 3. Charge for marking an original edge

For every original edge `u -> v` and every `j`, add

```text
node(v, j) -> node(u, j), capacity 1.
```

This edge crosses the cut when `level[v] >= j+1` but `level[u] < j+1`. Once
level increases of two or more are forbidden, exactly one such capacity-`1`
edge crosses if and only if

```text
level[v] = level[u] + 1.
```

The cut therefore pays exactly once when the original edge must be marked.

#### 4. Forbid increasing the level by two or more

For every original edge `u -> v` and `0 <= j < d-1`, add

```text
node(v, j+1) -> node(u, j), capacity INF.
```

This represents

```text
level[v] >= j+2  =>  level[u] >= j+1,
```

which is equivalent to `level[v] <= level[u] + 1`.

After adding these edges, every finite cut represents a valid level assignment,
and its capacity is exactly the number of original edges that must be marked.
Thus the maximum flow, which equals the minimum cut, is `minMarks(d)`.

### Correctness Proof

We prove that the algorithm returns the maximum achievable shortest distance.

For a fixed `d`, the infinite-capacity edges enforce the endpoint levels,
monotone threshold sequences, and the constraint
`level[v] <= level[u] + 1` for every original edge. Hence every finite cut
corresponds to a valid level assignment. The capacity-`1` edges contribute
exactly once for each original edge whose level increases by one, so the cut
capacity equals the cost of that assignment.

Marking all edges whose level increases by one makes every path from vertex `1`
to vertex `N` contain at least `d` marked edges. Conversely, every marking that
achieves distance at least `d` produces a valid level assignment by using
capped shortest distances, with cost no greater than the number of marked
edges. Therefore the minimum-cut value is exactly `minMarks(d)`.

The predicate `minMarks(d) <= K` is consequently true exactly for achievable
distances. It is monotone, and binary search checks the entire possible range
`0..N-1`. Therefore the algorithm returns exactly the maximum achievable
shortest distance.

### Complexity

For one checked distance `d`, the flow graph has

```text
V = N*d + 2
E = O((N+M)*d)
```

vertices and edges. Using the general `O(V^2 E)` bound for Dinic's algorithm,
one check costs `O((Nd)^2 (N+M)d)`. Binary search performs `O(log N)` checks.
The flow graph uses `O((N+M)d)` space.

Since `d <= N-1` and `N <= 30`, the largest constructed graph has at most
`872` vertices and fewer than `7000` forward edges.
