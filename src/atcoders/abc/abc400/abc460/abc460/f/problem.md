# F - Farthest Pair Query (ABC460)

**Contest:** [ABC460](https://atcoder.jp/contests/abc460) — AtCoder Beginner Contest 460  
**Task:** [https://atcoder.jp/contests/abc460/tasks/abc460_f](https://atcoder.jp/contests/abc460/tasks/abc460_f)

**Time limit:** 4 sec / **Memory limit:** 1024 MiB  
**Score:** 550 points

## Problem Statement

There is a tree with `N` vertices numbered `1..N`. The `i`-th edge connects
`U_i` and `V_i`. Initially every vertex is black.

Process `Q` queries in order. Each query gives an integer `x`:

- If vertex `x` is white, repaint it black; if black, repaint it white.
- Output the maximum distance between two black vertices (tree distance = number
  of edges on the simple path).

It is guaranteed that at every step at least two vertices remain black.

## Constraints

- `3 <= N <= 10^5`
- `1 <= U_i, V_i <= N`
- The graph is a tree
- `1 <= Q <= 10^5`
- `1 <= x <= N` for each query
- At all times at least two black vertices exist
- All input values are integers

## Input

```text
N
U_1 V_1
U_2 V_2
⋮
U_{N-1} V_{N-1}
Q
query_1
query_2
⋮
query_Q
```

Each query is a single integer `x`.

## Output

Print `Q` lines. The `i`-th line is the answer after the `i`-th query.

## Sample Input 1

```text
7
1 2
2 3
2 4
6 7
3 5
7 3
9
1
4
2
6
3
1
1
4
6
```

## Sample Output 1

```text
4
3
3
2
2
3
2
3
4
```

After each query, the black set changes by toggling one vertex; the answer is the
maximum pairwise distance among black vertices. The statement lists one optimal
pair per step.


## Solution

The answer for the current black vertices is the diameter of that set: the
largest tree distance between any two black vertices.

We maintain this diameter with a segment tree. Each segment-tree node stores the
two endpoints of the diameter among the black vertices in its interval, together
with the diameter length. A leaf is either:

- black: both endpoints are that vertex and the distance is `0`;
- white: it contains no endpoints and its distance is `-inf`.

When two child intervals are merged, the diameter of their union is one of:

1. the left child's diameter;
2. the right child's diameter;
3. a path joining one endpoint of the left diameter to one endpoint of the
   right diameter.

There are only four cross pairs, so a merge needs a constant number of distance
queries. A toggle changes one leaf and rebuilds the `O(log N)` nodes on its path
to the root.

The segment-tree leaves use DFS preorder. `pos[u]` is the leaf position of
vertex `u`, while `at[i]` is the vertex stored at position `i`. Therefore a
query for vertex `x` must update leaf `pos[x-1]`, not leaf `x-1`.

## Tree distances with LCA

For two vertices `u` and `v`, their distance is:

```text
dist(u, v) = depth[u] + depth[v] - 2 * depth[lca(u, v)]
```

Because every segment-tree merge performs several distance queries, LCA should
be answered in `O(1)`. We use an Euler tour followed by a sparse-table RMQ.

### Building the Euler tour

When DFS enters vertex `u`, append `u` to `euler` and record this index as
`first[u]`. After DFS finishes each child and returns to `u`, append `u` again:

```text
visit u:
    append u
    for every child v:
        visit v
        append u
```

The repeated append is essential. The Euler sequence represents the current
vertex while DFS walks along tree edges. Entering a child moves one level down;
returning from that child moves back to `u`, so `u` must be recorded again.

This gives the key LCA property: between the first occurrences of any two
vertices `u` and `v`, the vertex with minimum depth is their LCA.

Appending each vertex only once, as in an ordinary preorder, does not have this
property. For example, consider these edges:

```text
1 - 2 - 4
|
3
```

A preorder can be:

```text
1, 2, 4, 3
```

The interval from `4` to `3` does not contain their LCA `1`, so taking the
minimum depth in this interval is incorrect. The Euler tour is:

```text
1, 2, 4, 2, 1, 3, 1
```

The interval from the first `4` to the first `3` contains the return path
`4, 2, 1, 3`; its minimum-depth vertex is correctly `1`.

Appending `u` after every child also matters when `u` has multiple children.
It places `u` between the Euler-tour portions of consecutive child subtrees.
Thus a query whose endpoints lie in two different child subtrees necessarily
includes `u`, their LCA, in the queried interval.

Starting from a vertex `w`, its ancestors appear later as DFS returns upward:
first its parent, eventually its grandparent, and so on. They are not
necessarily consecutive because DFS may visit remaining sibling subtrees
between two returns.

Therefore the Euler-tour interval between two vertices is not only an upward
path. It can contain:

1. movement upward from the first vertex toward the LCA;
2. complete detours into sibling subtrees;
3. movement downward from the LCA toward the second vertex.

For the same example, the interval from the first `4` to the first `3` is:

```text
4, 2, 1, 3
```

If vertex `2` also had another child `5` visited after `4`, the interval would
instead contain a detour:

```text
4, 2, 5, 2, 1, 3
```

This interval is not the simple path from `4` to `3`, but that does not harm
RMQ. Every detour enters a descendant and therefore only increases depth. It
cannot introduce a vertex shallower than the LCA. Consequently, even with both
upward and downward movement and arbitrary sibling-subtree detours, the
minimum-depth vertex in the interval is still exactly the LCA.

### Sparse-table RMQ

Let:

```text
dp[i][j] = the minimum-depth vertex in
           euler[i .. i + 2^j - 1]
```

For `j = 0`, the interval contains one Euler-tour entry. For larger `j`, split
the interval into two equal halves and keep the shallower of their answers:

```text
dp[i][j] = shallower(
    dp[i][j-1],
    dp[i + 2^(j-1)][j-1],
)
```

To find `lca(u, v)`, set `l = first[u]` and `r = first[v]`, swapping them when
necessary so `l <= r`. Let:

```text
length = r - l + 1
w = floor(log2(length))
```

In Go, `w` is computed as `bits.Len(uint(length)) - 1`. Query two blocks of
length `2^w`:

```text
[l, l + 2^w - 1]
[r - 2^w + 1, r]
```

Together they cover the whole interval and may overlap. Overlap is harmless
because taking the minimum depth is idempotent. The shallower of the two stored
vertices is the LCA.

It is important to use `floor(log2(length))`. Using `bits.Len(length)` directly
would select a block larger than the query interval and can also index beyond
the sparse-table levels.

## Correctness outline

For every segment-tree node, assume both children store valid diameter endpoints.
Any longest path in their union lies completely in one child or has one endpoint
in each child. In a tree, a farthest cross pair can be found among the two
diameter endpoints of each child, so checking the two child diameters and four
cross pairs produces the union's diameter. The property holds at leaves and is
preserved by every merge; therefore the root stores the diameter of all black
vertices after each toggle.

The Euler-tour RMQ returns the correct LCA, so every distance used by these
merges is correct.

## Complexity

The Euler tour has `2N-1` entries. Building its sparse table takes
`O(N log N)` time and memory. Each LCA and distance query takes `O(1)`.

Each toggle updates `O(log N)` segment-tree nodes, and each merge performs only
a constant number of distance queries. Therefore all queries take
`O(Q log N)` time. The total complexity is:

```text
Time:   O(N log N + Q log N)
Memory: O(N log N)
```
