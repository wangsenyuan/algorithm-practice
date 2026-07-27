# D1. Constrained Tree

[Problem link](https://codeforces.com/problemset/problem/513/D1)

**Contest:** [Rockethon 2015](https://codeforces.com/contest/513)

time limit: 2 seconds

memory limit: 256 megabytes

## Problem

Construct a binary tree with `n` nodes. The nodes are labeled `1..n` by a
pre-order traversal:

```text
root, left subtree, right subtree
```

There are `c` constraints. Each constraint contains two labels `a`, `b` and
one direction:

- `a b LEFT`: node `b` must belong to the subtree rooted at the left child of
  node `a`;
- `a b RIGHT`: node `b` must belong to the subtree rooted at the right child
  of node `a`.

Find any binary tree satisfying all constraints.

## Constraints

For subproblem D1:

- `1 <= n <= 100`
- `1 <= c <= 50`
- `1 <= a, b <= n`
- Every direction is either `LEFT` or `RIGHT`

## Input

```text
n c
a_1 b_1 direction_1
...
a_c b_c direction_c
```

## Output

If no valid binary tree exists, print:

```text
IMPOSSIBLE
```

Otherwise, print `n` space-separated node labels representing the in-order
traversal of any valid tree:

```text
left subtree, root, right subtree
```

The labels in the output are still the labels assigned by the tree's pre-order
traversal.

## Sample 1

```text
Input
3 2
1 2 LEFT
1 3 RIGHT

Output
2 1 3
```

The only valid tree has node `2` in the left subtree of node `1` and node `3`
in its right subtree.

## Sample 2

```text
Input
3 2
1 2 RIGHT
1 3 LEFT

Output
IMPOSSIBLE
```

Because labels come from a pre-order traversal, node `2` cannot be in the
right subtree of node `1` while the later node `3` is in its left subtree.

## Solution

### Contiguous pre-order intervals

The important property is that every subtree occupies a contiguous interval
of pre-order labels.

Suppose a subtree contains exactly the labels `[l, r]`. Its root must be `l`,
because the root is visited first in pre-order. If the last label in its left
subtree is `split`, then its two children contain:

```text
left subtree:  [l+1, split]
right subtree: [split+1, r]
```

The left subtree is empty when `split = l`, and the right subtree is empty
when `split = r`.

Therefore, constructing the tree is equivalent to choosing one valid `split`
for every interval `[l, r]`.

### Convert constraints into split bounds

The implementation uses zero-based labels internally. For each root `a`, it
precomputes:

- `maxLeft[a]`: the largest label that is required to be in `a`'s left
  subtree;
- `minRight[a]`: the smallest label that is required to be in `a`'s right
  subtree;
- `maxRequired[a]`: the largest label mentioned by any constraint whose root
  is `a`;
- `invalid[a]`: whether a constraint asks for `b <= a`.

A descendant of `a` must have a larger pre-order label, so any constraint with
`b <= a` is immediately impossible.

For a candidate subtree `[l, r]`, all constrained descendants of `l` must be
inside the interval:

```text
maxRequired[l] <= r.
```

Every `LEFT` constraint requires its target to be at most `split`, while every
`RIGHT` constraint requires its target to be greater than `split`. Hence:

```text
split >= maxLeft[l]
split <  minRight[l].
```

Together with `l <= split <= r`, the complete candidate range is:

```text
low  = max(l, maxLeft[l])
high = min(r, minRight[l]-1)
```

Contradictory constraints naturally make `low > high`, leaving no valid split.

### Interval DP

Let `f(l, r)` mean that a valid subtree can be constructed from the pre-order
interval `[l, r]`.

For every `split` in `[low, high]`:

1. Recursively construct `[l+1, split]` as the left subtree, unless it is
   empty.
2. Recursively construct `[split+1, r]` as the right subtree, unless it is
   empty.
3. If both sides are possible, store `split` in `dp[l][r]` and mark the
   interval as possible.

The DP table uses:

```text
-2: not calculated
-1: impossible
>= l: the chosen split
```

Memoization ensures that each interval is solved only once. If `f(0, n-1)` is
impossible, print `IMPOSSIBLE`.

### Reconstruct the in-order traversal

For a possible interval `[l, r]`, retrieve its stored `split`. In-order
traversal visits:

```text
in-order(left subtree)
root l
in-order(right subtree)
```

Thus reconstruction is:

```text
output(l+1, split)
append(l)
output(split+1, r)
```

The implementation converts the zero-based label back to `l+1` when appending
it to the answer.

### Correctness

We prove by induction on the length of `[l, r]` that `f(l, r)` succeeds
exactly when a valid subtree with those pre-order labels exists.

For a one-node interval, the only possible split is the root itself. It is
accepted exactly when that node has no impossible descendant requirement.

For a larger interval, any valid binary tree has some last label `split` in
its left subtree. Pre-order contiguity forces its left and right subtrees to be
exactly `[l+1, split]` and `[split+1, r]`. All constraints belonging to root
`l` force this split into `[low, high]`, and the induction hypothesis says
that both recursive intervals are accepted. Therefore the DP considers and
accepts the tree's split.

Conversely, when the DP accepts a split, its bounds place every constrained
target of root `l` on the required side. By the induction hypothesis, both
recursive subtrees satisfy all constraints belonging to their nodes. Joining
them under root `l` therefore produces a valid subtree. This proves the DP is
both complete and sound.

The reconstruction prints the accepted left subtree, root, and accepted right
subtree, so it produces the required in-order traversal.

### Complexity

There are `O(n^2)` intervals, and each interval tries at most `O(n)` split
points:

```text
Time:  O(n^3 + c)
Space: O(n^2)
```

This is easily sufficient for D1, where `n <= 100`.
