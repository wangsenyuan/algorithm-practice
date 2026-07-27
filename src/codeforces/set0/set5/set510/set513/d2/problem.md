# Codeforces 513D2 — Constrained Tree

Problem: <https://codeforces.com/problemset/problem/513/D2>

## Statement

Construct a binary tree with vertices `1..n` whose preorder traversal is exactly

```text
1, 2, ..., n
```

There are `c` constraints:

- `a b LEFT`: vertex `b` must belong to the left subtree of `a`;
- `a b RIGHT`: vertex `b` must belong to the right subtree of `a`.

Print the inorder traversal of any valid tree, or `IMPOSSIBLE` if no such tree
exists.

For D2:

```text
1 <= n <= 1,000,000
1 <= c <= 100,000
```

## Key observation: every subtree is an interval

Because the preorder is fixed as `1, 2, ..., n`, the vertices of every subtree
form one consecutive interval.

Suppose the subtree rooted at `a` ends at `r`, and its left subtree ends at
`split[a]`. Then:

```text
root:          a
left subtree:  [a + 1, split[a]]
right subtree: [split[a] + 1, r]
```

`split[a] = a` means that the left subtree is empty.

Therefore, all constraints with the same root can be compressed into three
extreme values:

- `maxLeft[a]`: the largest vertex required in the left subtree;
- `minRight[a]`: the smallest vertex required in the right subtree;
- `maxRequired[a]`: the largest vertex required in either subtree.

These extremes are sufficient because both subtrees are intervals. If the left
subtree reaches `maxLeft[a]`, it also contains every smaller required-left
vertex. If it ends before `minRight[a]`, it excludes every required-right
vertex. Finally, the whole subtree must reach `maxRequired[a]`.

Any constraint with `b <= a` is immediately impossible: descendants of `a`
always appear after `a` in preorder.

## Greedy construction

Define:

```text
build(a, need)
```

as constructing a subtree rooted at `a` that contains all labels through
`need`, and returning the **smallest possible ending label** of that subtree.

There are two cases.

### 1. No vertex is required to be in the left subtree

Choose an empty left subtree:

```text
split[a] = a
```

If `a` already reaches `need`, the subtree ends at `a`. Otherwise, construct
the right subtree with:

```text
build(a + 1, need)
```

Putting unnecessary vertices on the left cannot shorten the complete subtree,
so the empty left subtree is optimal.

### 2. Some vertex must be in the left subtree

The left child is necessarily `a + 1`. Construct the shortest possible left
subtree that reaches every required-left vertex:

```text
leftEnd = build(a + 1, maxLeft[a])
```

Now check the first vertex required on the right:

```text
leftEnd < minRight[a]
```

If this is false, even the shortest valid left subtree already contains a
vertex required on the right. Making the left subtree larger cannot repair the
conflict, so the answer is impossible.

Otherwise set:

```text
split[a] = leftEnd
```

If the left subtree already reaches `need`, the subtree is complete. If not,
construct the right subtree:

```text
build(leftEnd + 1, need)
```

The top-level call is:

```text
build(1, n)
```

so all `n` vertices are included.

## Why choosing the shortest left subtree is safe

Consider any valid subtree rooted at `a`.

- Its left subtree must reach `maxLeft[a]`.
- `build(a + 1, maxLeft[a])` returns the smallest possible end among all such
  left subtrees.
- Replacing the original left subtree by this shorter one cannot swallow a
  required-right vertex. It only moves the beginning of the right subtree
  leftward.
- Any vertices removed from the left interval can instead be placed at the
  beginning of the right subtree, whose preorder interval is still valid.

Thus, if some solution exists, one also exists with the greedy shortest left
subtree. Conversely, if that shortest left subtree reaches `minRight[a]`, every
possible left subtree does, so no solution exists.

By induction on the constructed interval, `build(a, need)` either returns the
smallest feasible end or correctly reports impossibility.

## Avoiding recursion

A valid tree may be a chain of one million vertices, so recursive construction
or recursive inorder traversal can overflow the Go call stack or add avoidable
overhead.

The implementation uses:

- an explicit stack of frames for `build`;
- an explicit stack of intervals for inorder traversal.

After construction, `split[a]` uniquely describes the two child intervals, so
the inorder traversal can be generated iteratively.

## Complexity

Each constraint is aggregated once, and each vertex enters and leaves each
explicit stack only once.

```text
Time:  O(n + c)
Space: O(n + c)
```

This fits D2, unlike the interval DP used for D1.
