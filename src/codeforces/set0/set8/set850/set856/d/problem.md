# D. Masha and Cactus

## Key idea

Adding an extra edge `(u, v)` creates exactly one cycle on the tree path from
`u` to `v`. The final graph is a cactus iff every tree vertex is covered by at
most one chosen extra-edge path. So the problem is maximum weight selection of
vertex-disjoint tree paths.

Root the tree at `1`. Let `dp[x]` be the best value using only paths fully
inside the subtree of `x`.

If `x` is not covered by a chosen path:

```text
dp[x] = sum dp[child]
```

If `x` is covered, the chosen path must have `lca(u, v) = x`. Start from the
same child-subtree sum at `x`, add the path weight, and replace the already
counted `dp` values on the covered branches by the contribution of subtrees
that hang off the path.

For a processed vertex `y`, store:

```text
base[y] = sum dp[child of y] - dp[y]
```

When a candidate path with `lca = x` is considered, all descendants on the path
have already been processed. Querying the sum of `base[y]` over the tree path
`u..v` gives exactly the adjustment needed after starting with
`sum dp[child of x]`.

Thus:

```text
candidate = sum dp[child of x] + weight(u, v) + path_sum(base, u, v)
dp[x] = max(candidate over edges with lca x, sum dp[child of x])
```

Heavy-light decomposition supports each path sum in `O(log^2 n)`, and a segment
tree supports point updates of `base[x]` after `dp[x]` is fixed.

## Complexity

`O((n + m) log^2 n)` time and `O(n + m)` memory.

## Sources

- Codeforces 856D: Masha and Cactus
