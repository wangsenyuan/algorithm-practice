# F. Simple Cycles Edges

[Problem link](https://codeforces.com/problemset/problem/962/F)

**Contest:** [Codeforces Round #477 (Div. 2)](https://codeforces.com/contest/962)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given an undirected graph consisting of `n` vertices and `m` edges. The graph is not necessarily
connected. The graph has no multiple edges and no loops.

A **simple cycle** is a closed path that uses only distinct vertices.

More precisely, it is a sequence of vertices `v1, v2, ..., vk` with `k >= 3` such that:

- every consecutive pair `(vi, vi+1)` is an edge, and `(vk, v1)` is also an edge;
- all `k` vertices are different — no vertex appears twice.

So you may not “walk around” a cycle by revisiting a vertex. For example, in a triangle `1—2—3—1`, the walk
`1 → 2 → 3 → 1` is a simple cycle, but `1 → 2 → 1 → 2 → ...` is not.

An edge **belongs to** a simple cycle if that edge is one of the `k` edges of some simple cycle.

Find all edges that belong to **exactly one** simple cycle. If an edge lies on two or more different simple
cycles, do not output it (see the third sample).

## Input

The first line contains two integers `n` and `m`
(`1 <= n <= 100000`, `0 <= m <= min(n * (n - 1) / 2, 100000)`).

Each of the next `m` lines contains two integers `u` and `v` (`1 <= u, v <= n`, `u != v`) — an edge.

Edges are numbered from `1` in input order.

## Output

In the first line, print the number of edges that belong to exactly one simple cycle.

In the second line, print their indices in increasing order. If there are no such edges, print only
the first line with `0`.

## Example

### Input

```text
3 3
1 2
2 3
3 1
```

### Output

```text
3
1 2 3
```

### Input

```text
6 7
2 3
3 4
4 2
1 2
1 5
5 6
6 1
```

### Output

```text
6
1 2 3 5 6 7
```

### Input

```text
5 6
1 2
2 3
2 4
4 3
2 5
5 3
```

### Output

```text
0
```

## Solution

Use Tarjan's algorithm for edge-biconnected components.

Inside one edge-biconnected component, every two edges lie on a common cycle-like structure. Therefore an edge can
belong to exactly one simple cycle only when its whole component is exactly one cycle. That is easy to test:

- the component has the same number of edges and distinct vertices;
- every vertex has degree `2` inside this component.

The solver keeps a stack of edge indices during DFS. For a tree edge `u -> v`, after visiting `v`, if
`low[v] >= dfn[u]`, all edges popped from the stack up to this tree edge form one edge-biconnected component.
For each popped component, count its internal degrees and distinct vertices. If it is a simple cycle by the rule
above, add all its edge indices to the answer.

Back edges are pushed onto the stack only when they go to an ancestor, so every undirected edge is pushed exactly
once. Finally, sort the collected edge indices before printing them.

## Correctness

Tarjan's edge stack partitions the graph edges into edge-biconnected components. Any simple cycle is fully contained
inside one such component, so whether an edge belongs to one or multiple simple cycles can be decided from its
component alone.

If a component has `k` edges, `k` vertices, and every internal degree is `2`, then following the degree-2 links
around the connected component gives exactly one simple cycle. Every edge in this component is on that cycle and
there is no extra branch or chord that could create another simple cycle.

If a component fails this test, then it is not a single cycle. A bridge has too few edges or degree-1 endpoints, so
it belongs to no simple cycle. A larger biconnected structure with a branch, chord, or merged cycles gives at least
one vertex degree different from `2` or more edges than vertices, which means some edges lie on multiple simple
cycles or no edge set of the component is exactly one isolated cycle. Such edges must not be printed.

Thus the algorithm adds exactly the edges whose biconnected component is one simple cycle, which is exactly the
required set of edges that belong to exactly one simple cycle.

## Complexity

Each edge is pushed to and popped from the Tarjan stack once, and each component check scans only its own edges.
The total time is `O(n + m)`, plus sorting the answer. The memory usage is `O(n + m)`.

## Alternative Solution: Spanning Forest and Fundamental Cycles

Another way is to build an arbitrary spanning forest with DSU. Every edge rejected by DSU connects two vertices
that are already connected in the forest. Such a non-tree edge plus the unique forest path between its endpoints
forms one fundamental cycle.

For a fundamental cycle to be a valid answer cycle, every tree edge on its path must be used by no other
fundamental cycle. If two non-tree edges use the same forest edge, then their fundamental cycles can be combined
to create additional simple cycles, so the shared part and the related cycle edges cannot belong to exactly one
simple cycle.

We can count this with path difference on the forest:

1. Build the forest and collect all non-tree edges.
2. Precompute LCA on the forest.
3. For each non-tree edge `id = (u, v)`, add a unique marker `id + m` to `u` and `v`, and subtract it twice at
   `lca(u, v)`.
4. Accumulate markers bottom-up. A tree edge belongs to exactly one fundamental cycle only when the accumulated
   value on its child side is exactly one marker.
5. For each non-tree edge, compare how many tree edges on its fundamental path were marked only by it with the
   full path length. If all of them are private, output the whole fundamental cycle.

This version takes `O((n + m) log n)` time because of LCA preprocessing and queries, and `O((n + m) log n)` memory
for binary lifting.
