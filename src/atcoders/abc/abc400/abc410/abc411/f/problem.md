# F - Contraction

https://atcoder.jp/contests/abc411/tasks/abc411_f

**Time Limit:** 4 sec / **Memory Limit:** 1024 MiB

**Score:** 525 points

## Problem Statement

You are given an undirected graph `G_0` with `N` vertices and `M` edges.
Vertices are numbered `1, 2, ..., N`, and edges are numbered
`1, 2, ..., M`. Edge `i` connects vertices `U_i` and `V_i`.

Takahashi has a graph `G` and `N` pieces numbered `1, 2, ..., N`.
Initially, `G = G_0`, and piece `i` is placed on vertex `i`.

He performs `Q` operations. In the `i`-th operation, an integer `X_i` is given.
Let the original edge `X_i` connect vertices `U_{X_i}` and `V_{X_i}`.

If the pieces `U_{X_i}` and `V_{X_i}` are placed on different vertices of the
current graph `G`, and those two vertices are connected by an edge in `G`, then
contract that edge:

- merge the two endpoints into one vertex,
- remove self-loops,
- replace multiple edges with a single simple edge,
- move all pieces on the two merged vertices to the new vertex.

All other pieces remain on their corresponding vertices.

If the two pieces are already on the same vertex, or if their current vertices
are not adjacent in `G`, do nothing.

For each operation, output the number of edges in `G` after the operation.

## Edge Contraction

Contracting an edge between vertices `u` and `v` means creating a new vertex
`w` and connecting `w` to each other vertex `x` exactly when at least one of
`u-x` or `v-x` existed before the contraction. Then vertices `u` and `v`, and
all edges incident to them, are removed.

## Constraints

- `2 <= N <= 3 * 10^5`
- `1 <= M <= 3 * 10^5`
- `1 <= U_i < V_i <= N`
- `(U_i, V_i) != (U_j, V_j)` if `i != j`.
- `1 <= Q <= 3 * 10^5`
- `1 <= X_i <= M`
- All input values are integers.

## Solution

Maintain each current contracted vertex as a DSU component. For every component
root `u`, store a set:

```text
adj[u] = current neighboring component roots of u
```

Because the graph is simple after each contraction, each neighbor appears at
most once in this set.

For a query edge `id`, let its original endpoints be `(x, y)`. The pieces
`x` and `y` currently belong to:

```text
u = find(x)
v = find(y)
```

If `u == v`, the pieces are already in the same contracted vertex, so nothing
happens.

Otherwise, the operation contracts an edge only when `u` and `v` are adjacent.
In the current implementation, queried pairs that are not adjacent do not change
the graph.

When contracting adjacent components `u` and `v`, merge the smaller adjacency
set into the larger one. This is the standard small-to-large trick.

The edge count `tot` is maintained directly:

1. Remove all edges incident to `u` and all edges incident to `v`:

   ```text
   tot -= len(adj[u])
   tot -= len(adj[v])
   ```

2. The edge between `u` and `v` was subtracted twice, but it existed only once,
   so add one back:

   ```text
   tot++
   ```

3. Delete the internal edge `u-v`.

4. For every old neighbor `w` of `v`, connect `w` to `u`. If `w` was already
   adjacent to `u`, this would become a multi-edge after contraction, so keep
   only one edge.

5. Make `v`'s DSU parent `u`.

6. Add back the number of remaining edges incident to the merged component:

   ```text
   tot += len(adj[u])
   ```

After each query, output `tot`.

The small-to-large merge makes the total number of moved adjacency entries
`O((N + M) log N)` over all contractions, so the method handles the constraints.
