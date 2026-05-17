# F. The 67th Tree Problem (Codeforces 2218F)

**Limits:** 4 seconds per test, 256 MB  
**I/O:** standard input / standard output

Source: [https://codeforces.com/problemset/problem/2218/F](https://codeforces.com/problemset/problem/2218/F)

---

You are given two integers `x` and `y`.

Construct a **tree** with **`x + y`** nodes, **rooted at node `1`**, such that:

- exactly **`x`** vertices have **even** subtree size;
- exactly **`y`** vertices have **odd** subtree size.

If several trees work, print any. If none exists, print `NO`.

**Subtree definition:** the subtree of a vertex `u` is the set of all vertices that lie on a simple path from that vertex to the root **through `u`** (including `u` itself).

---

*The statement opens with a short fictional intro on the Codeforces page; the formal task is the construction above.*

## Solution Explanation

Root the tree at node `1`. For every vertex `u`, let `sz[u]` be the size of its
subtree.

Two very useful facts are enough:

1. Every leaf has subtree size `1`, so every tree has at least one odd vertex.
   Therefore `y = 0` is impossible.
2. If `sz[u]` is even, then `u` must have at least one child `v` with odd
   `sz[v]`.

The second fact follows from

```text
sz[u] = 1 + sum(sz[child])
```

If all children of `u` had even subtree size, then the sum of child subtree
sizes would be even, so `sz[u]` would be odd. That contradicts `sz[u]` being
even.

Now choose one odd child for every even vertex. The chosen odd children are all
different, because each vertex has only one parent. So every even vertex needs a
distinct odd child below it, which gives the necessary condition:

```text
x <= y
```

There is one more special case. If `x = 0`, then all vertices must have odd
subtree size. In particular the root has subtree size `x + y = y`, so `y` must
be odd. If `y` is even, the root itself would be an even vertex, which is not
allowed.

So the impossible cases are:

- `y == 0`;
- `x > y`;
- `x == 0` and `y` is even.

It remains to show that all other cases are constructible.

### Building Blocks

The construction uses only two simple shapes attached to the root.

#### A direct leaf

Add one edge:

```text
1 - a
```

If `a` has no children, then `sz[a] = 1`, so it contributes one odd vertex.

#### A length-2 branch

Add two edges:

```text
1 - a - b
```

Here `b` is a leaf, so `sz[b] = 1` is odd. Vertex `a` has exactly one child
`b`, so `sz[a] = 2` is even.

Thus one length-2 branch contributes:

- one even vertex: `a`;
- one odd vertex: `b`.

The root's parity is controlled by the total number of vertices `n = x + y`:

- if `n` is even, then the root is one even vertex;
- if `n` is odd, then the root is one odd vertex.

### Case 1: `x = 0`

We already know `y` must be odd. A star is enough:

```text
1 connected to every other vertex
```

The root has subtree size `y`, which is odd, and every other vertex is a leaf,
also odd. Therefore the tree has `0` even vertices and `y` odd vertices.

This is the branch:

```go
if x == 0 {
    connect all other nodes directly to 1
}
```

### Case 2: `x > 0` and `x + y` is even

Now the root already contributes one even vertex. We still need `x - 1` more
even vertices.

Create `x - 1` length-2 branches from the root:

```text
1 - a1 - b1
1 - a2 - b2
...
```

Each branch gives one even vertex and one odd vertex. Together with the root,
this gives exactly `x` even vertices.

How many vertices have been used?

```text
1 root + 2 * (x - 1) branch vertices = 2x - 1
```

The total number of vertices is `x + y`, so the number of remaining vertices is

```text
(x + y) - (2x - 1) = y - x + 1
```

Because `x <= y`, this number is at least `1`. Attach all remaining vertices
directly to the root as leaves. They are all odd vertices.

The final counts are:

- even: root plus `x - 1` branch-middle vertices, total `x`;
- odd: `x - 1` branch leaves plus `y - x + 1` direct leaves, total `y`.

This is why the code loops `x - 1` times in the even-`n` case.

### Case 3: `x > 0` and `x + y` is odd

Now the root is odd, so it cannot help with the `x` even vertices. We create
exactly `x` length-2 branches:

```text
1 - a1 - b1
1 - a2 - b2
...
1 - ax - bx
```

Each branch contributes one even vertex and one odd vertex. The root contributes
one more odd vertex.

The number of vertices used is:

```text
1 root + 2x branch vertices = 2x + 1
```

The number of remaining vertices is

```text
(x + y) - (2x + 1) = y - x - 1
```

In this case `x + y` is odd, so `x` and `y` have different parity. Since
`x <= y` and `x > 0`, the smallest possible odd-`n` valid gap is actually
`y - x >= 1`, and because the parities differ, `y - x - 1 >= 0`.

Attach all remaining vertices directly to the root as leaves.

The final counts are:

- even: the `x` branch-middle vertices;
- odd: root, the `x` branch leaves, and `y - x - 1` direct leaves, total `y`.

This is why the code loops `x` times in the odd-`n` case.

### Correctness

The impossible checks are necessary:

- `y == 0` is impossible because every tree has at least one leaf, and every
  leaf has odd subtree size.
- `x > y` is impossible because every even vertex has a distinct odd child.
- `x == 0` with even `y` is impossible because the root has subtree size `y`,
  so the root would be even.

For every remaining input, the construction builds a connected tree:

- all edges connect a new vertex either to the root or to the previous vertex in
  a length-2 branch;
- every vertex from `1` to `x + y` is used exactly once;
- the number of edges is exactly `x + y - 1`.

The subtree parities are exactly as counted above:

- direct leaves have subtree size `1`;
- the middle vertex of every length-2 branch has subtree size `2`;
- the root has subtree size `x + y`.

Therefore the constructed tree has exactly `x` even-subtree vertices and exactly
`y` odd-subtree vertices.

### Complexity

The algorithm prints `x + y - 1` edges and does only constant work per edge.

Time complexity: `O(x + y)` per test case.  
Memory complexity: `O(x + y)` for the edge list.
