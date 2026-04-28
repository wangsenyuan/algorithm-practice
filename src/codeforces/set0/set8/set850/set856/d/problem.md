# D. Masha and Cactus

## Key idea

Adding an extra edge `(u, v)` creates exactly one cycle on the tree path from
`u` to `v`. The final graph is a cactus iff every tree vertex is covered by at
most one chosen extra-edge path. So the problem is maximum weight selection of
vertex-disjoint tree paths.

The "vertex-disjoint paths" formulation is the important reduction:

- choosing `(u, v)` uses every vertex on the original tree path `u..v`;
- if two chosen paths share a vertex, that vertex belongs to two cycles, which
  is not allowed in a cactus;
- if all chosen paths are vertex-disjoint, then every added edge creates one
  independent cycle, so the graph is a valid cactus.

Root the tree at `1`. Let `dp[x]` be the best value using only paths fully
inside the subtree of `x`.

If `x` is not covered by a chosen path:

```text
dp[x] = sum dp[child]
```

If `x` is covered by a chosen path inside its subtree, that path must have
`lca(u, v) = x`. Any path whose LCA is below `x` is already handled inside one
child subtree. Any path whose LCA is above `x` is not fully inside `x`'s
subtree, so it is not part of `dp[x]`.

Start with:

```text
sum = sum dp[child of x]
```

This assumes every child subtree independently takes its best answer. If we now
choose a path `u..v` with `lca(u, v) = x`, the vertices on that path become
unavailable. For every vertex on the path, we may still keep answers from the
child subtrees that branch away from the path, but we must discard the `dp` of
the child that continues along the path.

One way to write the transition is to publish a correction for each processed
vertex `y`:

```text
base[y] = sum dp[child of y] - dp[y]
```

Why this value works:

- `sum dp[child of y]` is the contribution of all subtrees hanging below `y`
  before deciding whether `y` is covered;
- `dp[y]` is what the parent side would have counted for the whole subtree of
  `y`;
- the difference is exactly the correction needed when `y` itself lies on the
  newly chosen path and the whole `dp[y]` is no longer available as a closed
  subtree answer.

When a candidate path with `lca = x` is considered, all descendants on the path
have already been processed and published into the segment tree. Querying the
tree path `u..v` therefore gives the sum of `base[y]` for the already-processed
proper descendants of `x` that lie on the path. That is exactly the adjustment
needed after starting with `sum dp[child of x]`.

Thus:

```text
candidate = sum dp[child of x] + weight(u, v) + published_path_sum(u, v)
dp[x] = max(candidate over edges with lca x, sum dp[child of x])
```

At the moment we evaluate candidate paths for `x`, `dp[x]` has not been
published to its parent yet. In the implementation, the segment tree update for
`x` happens only after `dp[x]` is finalized, so candidate queries for `x` see
only already-processed descendants. The explicit `sum dp[child of x]` supplies
the LCA contribution.

Heavy-light decomposition supports each path sum in `O(log^2 n)`, and a segment
tree supports point updates after the relevant DP value is fixed.

## Editorial formula and implementation details

The official editorial uses:

```text
f[v] = best answer inside subtree v
s[v] = sum f[child of v]
g[u] = f[p] - f[u], where p is parent of u
```

The subtle point is that the `f[p]` in the `g[u]` line should be understood as
the value contributed by `p` after the path deletes the parent-child branch
toward `u`. In code this is the no-cycle-through-`p` base value:

```text
g[u] = s[parent[u]] - f[u]
```

It is not the finalized `f[parent[u]]`. The finalized `f[parent[u]]` may have
chosen an extra edge whose cycle already passes through `parent[u]`. If an
ancestor later chooses another path through the same vertex, keeping that
finalized choice would put the vertex in two cycles and overcount.

In the code:

```go
// after fp[u] is known
for _, v := range adj[u] {
    tr.Update(dfn[v], s[u]-fp[v])
}
```

This stores `g[v] = s[u] - fp[v]`, where `u` is the parent of `v`.

Now consider an extra edge `(x, y, c)` with `v = lca(x, y)`. If this edge is
chosen, the tree path `x..y` is removed from the independent subtree choices.
Let:

```text
x' = child of v on the path to x
y' = child of v on the path to y
```

If one endpoint is exactly `v`, that side is skipped.

The remaining isolated subtree roots after deleting the path are:

- child subtrees of `v` except `x'` and `y'`;
- child subtrees of every internal path vertex that are not on the path;
- child subtrees of `x` and `y`.

The code computes this as:

```text
c + s[v]
  + side(x)
  + side(y)
```

For one side, for example from endpoint `x` up to the child `x'`:

```text
side(x) = s[x] + sum g[node on path x..x'] - f[x']
```

This matches the code:

```go
tmp += s[x]
x1 := nextAncestor(x, v)
tmp += query(x, x1) - fp[x1]
```

`query(x, x1)` sums `g` on the downward chain from `x1` to `x`. The final
`-fp[x1]` removes the child subtree of `v` that lies on the selected path,
because the initial `s[v]` included it.

For example, if `v` has child `x'`, `s[v]` starts by including `f[x']`. But the
path from `v` to `x` uses vertices inside `x'`, so the whole independent answer
`f[x']` is no longer available. Instead, only the off-path branches along that
side remain, and those are represented by `s[x] + sum g`.

This is exactly why using:

```go
tr.Update(dfn[v], fp[u]-fp[v])
```

is wrong. It stores finalized `f[u] - f[v]`. If `f[u]` already selected a cycle
through `u`, a higher path can count that cycle together with another cycle
through `u`, which violates the cactus condition. The correct update is:

```go
tr.Update(dfn[v], s[u]-fp[v])
```

The implementation order is bottom-up by vertex number. This is valid because
the input guarantees `parent[i] < i`, so every child has a larger index than its
parent. During DFS DP, after finishing a vertex `u`, we update the segment tree
for every child `v` with:

```text
s[u] - fp[v]
```

Later, when an ancestor evaluates a path passing through that child branch, the
HLD path sum can include this correction in `O(log^2 n)`.

## Complexity

`O((n + m) log^2 n)` time and `O(n + m)` memory.

## Sources

- Codeforces 856D: Masha and Cactus


### editorial

Let us use dynamic programming for a rooted tree and some data structures. Denote as fv the maximal total beauty of edges that have both ends in a subtree of v, such that if we add them all to the subtree it would be a cactus.

To calculate fv let us consider two cases: v belongs to some cycle, or it doesn't. If it doesn't belong to any cycle, fv is equal to the sum of fu for all children u of v.

If v belongs to a cycle, let us iterate over all possible cycles it can belong to. Such cycle is generated by an added edge (x, y) such that LCA(x, y) = v. Try all possible such edges and then temporarily delete a path from x to y from a tree, calculate the sum of fu for all u — roots of the isolated subtrees after the deletion of the path, and add it to the beauty of (x, y).

Now we have an O(nm) solution.

To speed up this solution let us use some data structures. First, we need to calculate LCA for all endpoints of the given edges, any fast enough standard algorithm is fine. The second thing to do is to be able to calculate the sum of fu for all subtrees after removing the path. To do it, use the following additional values: gu = fp - fu, where p is the parent of u, and sv = sum(fu), where u are the children of v.

Now the sum of fu for all subtrees after x - y path removal is the sum of the following values: sx, sy, sv - fx' - fy', the sum of gi for all i at [x, x'), the sum of gi for all i at [y, y'), where x' is the child of v that has x in its subtree, and y' is the child of v that has y in its subtree. We need some data structure for a tree that supports value change in a vertex and the sum for a path, range tree or Fenwick are fine. The complexity is O((n + m)log(n)).
