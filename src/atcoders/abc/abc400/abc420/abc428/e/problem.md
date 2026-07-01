# E - Farthest Vertex

[Problem link](https://atcoder.jp/contests/abc428/tasks/abc428_e)

**Contest:** [AtCoder Beginner Contest 428](https://atcoder.jp/contests/abc428)

time limit: 2 sec

memory limit: 1024 MiB

score: 475 points

There is a tree with `N` vertices numbered `1` to `N`. The `i`-th edge connects vertices `A_i` and
`B_i`.

The distance between vertices `u` and `v` is the number of edges on the unique path between them.

For each `v = 1, 2, ..., N`, among all vertices, output the one farthest from `v`. If several
vertices tie for maximum distance, output the **largest** vertex number.

## Constraints

- `2 <= N <= 5 * 10^5`
- `1 <= A_i < B_i <= N`
- The input graph is a tree
- All input values are integers

## Input

```text
N
A_1 B_1
A_2 B_2
...
A_{N-1} B_{N-1}
```

## Output

Print `N` lines. The `i`-th line should contain the answer for `v = i`.

## Sample Input 1

```text
3
1 2
2 3
```

## Sample Output 1

```text
3
3
1
```

- From vertex `1`, the farthest vertex is `3`.
- From vertex `2`, vertices `1` and `3` tie; answer is `3`.
- From vertex `3`, the farthest vertex is `1`.

## Sample Input 2

```text
5
1 2
2 3
2 4
1 5
```

## Sample Output 2

```text
4
5
5
5
4
```

## Solution

For every vertex `u`, we need the vertex with the maximum distance from `u`. If several vertices
have the same distance, the largest vertex number must be chosen.

Use reroot DP. For each vertex, keep the best two candidates seen from that vertex. A candidate is:

- `leaf`: the farthest vertex id represented by this candidate.
- `to`: the next vertex on the path from the current vertex toward `leaf`.
- `dist`: the distance from the current vertex to `leaf`.

Candidates are compared by `(dist, leaf)`: larger distance is better, and if distances are equal,
larger vertex id is better.

### First DFS

Root the tree at vertex `1`.

For each vertex `u`, initialize its candidate list with `u` itself:

```text
(leaf = u, to = u, dist = 0)
```

After solving a child `v`, the best candidate from that child subtree becomes a candidate for `u`:

```text
(leaf = best[v].leaf, to = v, dist = best[v].dist + 1)
```

Insert it into `u`'s top two candidates. After this DFS, `dp[u]` contains the best candidates whose
farthest vertex lies inside `u`'s rooted subtree.

### Second DFS

The first DFS does not know about vertices outside a node's subtree. The second DFS passes one
extra candidate `from` from the parent side into each node.

When visiting `u`, insert `from` into `dp[u]`. Then `dp[u][0]` is the best candidate among all
vertices in the whole tree, so:

```text
answer[u] = dp[u][0].leaf
```

For each child `v`, we must pass the best candidate visible from `u` that does not immediately go
through `v`, because `v` already handles its own subtree from the first DFS.

- If `dp[u][0].to != v`, pass `dp[u][0]`.
- Otherwise, pass `dp[u][1]`.

The candidate is moved one edge farther when passed to the child:

```text
(leaf = chosen.leaf, to = u, dist = chosen.dist + 1)
```

The initial parent-side candidate for the root is made invalid with a very negative distance, so it
never beats a real vertex.

### Correctness

Consider any vertex `u`. Every other vertex belongs to exactly one direction from `u`: either `u`
itself, one of `u`'s child subtrees, or the parent side of `u`.

The first DFS records the best candidates from `u` itself and from each child subtree. Therefore it
correctly covers all downward directions in the rooted tree.

The second DFS supplies the missing parent-side direction. When moving from `u` to a child `v`, the
best candidate to pass is the best candidate at `u` that does not enter `v`'s own subtree. If the
current best candidate goes through `v`, the second best candidate is used instead; otherwise the
best candidate is safe to pass. Thus the child receives exactly the best candidate among all
vertices outside its subtree.

After inserting this parent-side candidate, `dp[u]` contains the best candidate from every possible
direction around `u`. Since candidates are always compared by larger distance and then larger vertex
id, `dp[u][0].leaf` is exactly the required answer for `u`.

### Complexity

Each edge is processed once in each DFS, and each insertion keeps only two candidates.

The time complexity is `O(N)`, and the memory complexity is `O(N)`.
