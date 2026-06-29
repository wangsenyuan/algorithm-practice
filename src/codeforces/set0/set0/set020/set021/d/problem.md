# D. Traveling Graph

[Problem link](https://codeforces.com/problemset/problem/21/D)

**Contest:** [Codeforces Alpha Round 21 (Codeforces format)](https://codeforces.com/contest/21)

time limit per test: 0.5 seconds

memory limit per test: 64 megabytes

input: standard input

output: standard output

You are given an undirected weighted graph. Find the length of the shortest cycle which starts from
vertex `1` and passes through all the edges at least once. The graph may contain multiple edges
between a pair of vertices and loops (edges from a vertex to itself).

## Input

The first line contains two integers `n` and `m` (`1 <= n <= 15`, `0 <= m <= 2000`) — the number of
vertices and the number of edges.

Each of the next `m` lines contains three integers `x`, `y`, and `w` (`1 <= x, y <= n`,
`1 <= w <= 10000`) — the endpoints of an edge and its length.

## Output

Print the minimal cycle length, or `-1` if no such cycle exists.

## Example

### Input

```text
3 3
1 2 1
2 3 1
3 1 1
```

### Output

```text
3
```

### Input

```text
3 2
1 2 3
2 3 4
```

### Output

```text
14
```

### Note

- In the first example, the triangle `1 -> 2 -> 3 -> 1` uses every edge exactly once, so the
  answer is `3`.
- In the second example, every valid cycle must start at `1`, traverse `(1, 2)`, then `(2, 3)`, and
  return to `1` without an edge between `3` and `1`. One shortest route is
  `1 -> 2 -> 3 -> 2 -> 1`, with total length `14`.

## Solution

This is the undirected Chinese postman problem on a very small graph (`n <= 15`).

Every original edge must be used at least once, so the answer always contains:

```text
sum = total weight of all input edges
```

If after using every edge once all non-isolated vertices have even degree and are reachable from
vertex `1`, then those edges already form an Eulerian connected multigraph, and we can start from
`1`, traverse every edge once, and return to `1`.

The only problem is odd-degree vertices. In an Eulerian cycle every vertex must have even degree.
Whenever we walk an extra path between two odd vertices, all internal vertices gain degree `2`, and
the two endpoints each gain degree `1`; therefore those two odd vertices become even. So the task is:

```text
pair all odd-degree vertices with minimum total shortest-path distance
```

Then:

```text
answer = sum + minimum_pairing_cost
```

### Connectivity

Only vertices with positive degree matter. If some such vertex is not reachable from vertex `1`,
then no cycle starting at `1` can cover its edges, so the answer is `-1`.

The implementation checks this after Floyd-Warshall. If `deg[i] > 0` and `dist[1][i]` is still
infinite, the required cycle is impossible.

Loops are harmless for parity: a loop contributes `2` to the degree of its vertex. It still adds its
weight to `sum`, but it never creates an odd-degree vertex.

### Shortest paths

For duplicating paths, only the shortest distance between every pair of vertices matters.

Build `dist` from the input edges, keeping the minimum edge weight for multiple edges, and run
Floyd-Warshall:

```text
dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])
```

Because `n <= 15`, `O(n^3)` is tiny.

### Matching odd vertices

Let `odd` be the list of all odd-degree vertices. Its size is always even. We use bitmask DP:

```text
dp[mask] = minimum cost to pair all odd vertices whose bits are set in mask
```

Transition:

1. Pick the first set bit `i` in `mask`.
2. Pair it with any other set bit `j`.
3. Remove both bits and add their shortest-path distance.

```text
dp[mask] = min(dp[mask without i and j] + dist[odd[i]][odd[j]])
```

The final extra cost is `dp[(1 << len(odd)) - 1]`.

### Correctness

Any valid cycle starts at `1`, returns to `1`, and uses every required edge. If we mark every time an
edge is traversed, the used edges form a connected multigraph where every vertex has even degree.
The first traversal of each input edge contributes exactly `sum`; every additional traversal is some
extra walk in the graph.

Vertices that already have even degree must stay even. Odd-degree vertices must be changed to even,
and one duplicated path flips parity only at its two endpoints. Therefore the extra walks must pair
up all odd-degree vertices. For each pair, using anything other than a shortest path cannot be
better, so the cheapest possible extra cost is the minimum pairing cost over shortest-path
distances.

Conversely, if we choose such a minimum pairing and duplicate the corresponding shortest paths, all
vertices with edges become even, and they are connected to vertex `1`. The resulting multigraph has
an Eulerian cycle starting at `1`, which traverses all original edges at least once. Thus
`sum + dp[all_odd]` is both achievable and minimal.

### Complexity

Floyd-Warshall takes `O(n^3)`. If there are `p` odd vertices, the pairing DP takes `O(p^2 * 2^p)`,
with `p <= n <= 15`.

Memory usage is `O(n^2 + 2^p)`.
