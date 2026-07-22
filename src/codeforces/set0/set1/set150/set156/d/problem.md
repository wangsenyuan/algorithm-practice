# D. Clues

[Problem link](https://codeforces.com/problemset/problem/156/D)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Sherlock has `n` clues and `m` mutual direct links already found between them
(undirected, at most one link between any pair).

Clues are considered linked if they are in the same connected component via
direct links (possibly through other clues). To solve the crime, all clues must
be linked together.

Let `T` be the minimum number of additional direct links needed to connect all
clues. Count the number of ways to choose exactly `T` additional direct links
so that the graph becomes connected. Two ways differ if some pair of clues is
linked in one way but not the other.

Print the answer modulo `k`.

## Constraints

- `1 <= n <= 10^5`
- `0 <= m <= 10^5`
- `1 <= k <= 10^9`
- Edges: `1 <= a, b <= n`, `a != b`, no duplicate edges

## Input

```text
n m k
a b
...
```

## Output

Print a single integer — the number of ways modulo `k`.

## Examples

### Sample 1

```text
Input
2 0 1000000000

Output
1
```

### Sample 2

```text
Input
3 0 100

Output
3
```

### Sample 3

```text
Input
4 1 1000000000
1 4

Output
8
```


### ideas
1. 假设目前有W个components, 那么还需要W-1个link
2. 假设W = 2, 答案 = s1 * s2 (这两个compoent中任意取1个, 然后连接它们)
3. 如果W = 3, 先合并了两个, 现在合并第三个 (和正确结果不一致)
4. 而且, 这样子, 和操作顺序也有关系, 比如先处理第二个和第三个, 再处理第一个, 结果似乎不一样
5. 那么现在其实是吧W个点, 组合成一棵树的方案.
6. 这个好像有个公式的.
7.

## Solution

Use DSU to merge every pair of clues that is already directly linked. Suppose
the resulting graph has `W` connected components, whose sizes are

```text
s[1], s[2], ..., s[W]
```

and therefore

```text
s[1] + s[2] + ... + s[W] = n.
```

If `W = 1`, all clues are already connected. We add no links, and there is
exactly one valid choice: the empty set of links. Its count before taking the
modulo is `1`.

Otherwise, at least `W - 1` new links are necessary: one link can reduce the
number of connected components by at most one. It is also possible to connect
`W` components with exactly `W - 1` links. After contracting every original
component to one super-vertex, these added links must consequently form a tree
on the `W` super-vertices.

The remaining task is to count these trees while remembering that an edge
between components `i` and `j` can be realized by any of

```text
s[i] * s[j]
```

pairs of original vertices.

### Cayley's formula for labeled trees

For `x` labeled vertices, the number of different trees is

```text
x^(x - 2).
```

This is Cayley's formula. A convenient proof uses a Prüfer sequence.

Repeatedly remove the smallest labeled leaf and write down its neighbor. For a
tree on `x` vertices, this produces a sequence of exactly `x - 2` labels. The
process is reversible, so trees and Prüfer sequences are in one-to-one
correspondence.

Here, a leaf means a vertex whose current degree is `1`. "Write down its
neighbor" means append the neighbor's label to the Prüfer sequence, then delete
the chosen leaf and its edge. We stop when only two vertices remain.

For example, consider the tree with edges

```text
1-4, 2-4, 4-5, 5-3
```

which looks like

```text
1 - 4 - 5 - 3
    |
    2
```

Encode it step by step:

1. The leaves are `1, 2, 3`. Remove the smallest leaf `1`; its neighbor is
   `4`, so append `4`. The sequence is now `[4]`.
2. The leaves are now `2, 3`. Remove `2`; its neighbor is `4`, so append `4`.
   The sequence is now `[4, 4]`.
3. The remaining tree is `4 - 5 - 3`, whose leaves are `3, 4`. Remove the
   smallest leaf `3`; its neighbor is `5`, so append `5`. The sequence is now
   `[4, 4, 5]`.

Only vertices `4` and `5` remain, so the process stops. There were `5`
vertices, and the resulting sequence has `5 - 2 = 3` entries as expected.

Always choosing the smallest leaf makes the encoding deterministic: the same
tree always produces the same sequence. The reverse decoding procedure also
reconstructs exactly one tree from each sequence.

Every one of the `x - 2` positions may contain any of the `x` labels. Hence the
number of sequences, and therefore the number of labeled trees, is

```text
x * x * ... * x = x^(x - 2).
```

We cannot directly use `W^(W - 2)` here, because the super-vertices represent
components of different sizes. Different component-tree edges can have
different numbers of choices for their actual endpoints.

### Weighted Cayley formula

First consider only three components, with sizes `a`, `b`, and `c`. There are
three possible component trees, according to which component is in the middle.

If `A` is the middle component, the tree is

```text
B - A - C
```

The edge `A-B` has `a * b` possible pairs of original endpoints, and `A-C`
has `a * c` possible pairs. This tree therefore represents

```text
(a * b) * (a * c) = a^2 * b * c
```

original added-link sets. Similarly, choosing `B` or `C` as the middle
component gives `a * b^2 * c` or `a * b * c^2`. Summing all three trees gives

```text
a^2*b*c + a*b^2*c + a*b*c^2
= a*b*c*(a + b + c)
= a*b*c*n.
```

This is exactly the weighted formula for `W = 3`:

```text
product(s[i]) * n^(W - 2)
= a*b*c*n.
```

Now generalize this calculation.

#### Weight of one fixed component tree

Fix a tree `T` on all `W` components, and let `deg[i]` be the degree of
component `i` in this tree. Component `i` supplies an endpoint for every one of
its `deg[i]` incident edges. Each endpoint can independently be any of its
`s[i]` original vertices, so component `i` contributes

```text
s[i]^deg[i]
```

choices. Multiplying over all components, the number of original added-link
sets represented by this one component tree is

```text
product(s[i]^deg[i]) over all i.
```

For example, if the tree contains edges `(1, 2)` and `(1, 3)`, component `1`
supplies endpoints for two edges while components `2` and `3` supply one each:

```text
(s[1] * s[2]) * (s[1] * s[3])
= s[1]^2 * s[2] * s[3].
```

#### Convert degrees into Prüfer-sequence occurrences

The important Prüfer-sequence property is

```text
occ[i] = deg[i] - 1,
```

where `occ[i]` is the number of times label `i` occurs in the sequence. One way
to understand this property is that label `i` is recorded whenever one of its
neighboring leaves is removed. This happens for all but the last of its
incident edges; after that, `i` itself is a leaf or one of the final two
vertices. Therefore it is recorded exactly `deg[i] - 1` times.

Consequently,

```text
deg[i] = 1 + occ[i].
```

Substitute this into the weight of one tree:

```text
product(s[i]^deg[i])
= product(s[i]^(1 + occ[i]))
= product(s[i]) * product(s[i]^occ[i]).
```

The first `product(s[i])` is a common base factor: every component is a vertex
of the tree, so it has at least one incident edge and contributes at least one
factor `s[i]`.

The second product is encoded directly by the Prüfer sequence. Every occurrence
of label `i` contributes one additional factor `s[i]`. For example, if a Prüfer
sequence is

```text
[2, 2, 4]
```

then its component tree has weight

```text
product(s[i]) * s[2] * s[2] * s[4].
```

#### Sum the weights of all component trees

Every Prüfer sequence has `W - 2` positions. At any one position, we may choose
any component label `i`, and choosing it contributes weight `s[i]`. Therefore,
the sum of the possible weights for one position is

```text
s[1] + s[2] + ... + s[W] = n.
```

The positions are independent. For two positions, for example, the weighted
sum is

```text
(s[1] + ... + s[W]) * (s[1] + ... + s[W]) = n^2.
```

Therefore, all `W - 2` positions together contribute

```text
n^(W - 2).
```

Finally, multiply by the common base factor `product(s[i])`:

```text
answer = product(s[i]) * n^(W - 2) mod k.
```

Another way to remember the formula is

```text
product(s[i]) : one endpoint-choice factor for every component
n^(W - 2)     : the extra endpoint-choice factors encoded by the Prüfer sequence
```

This also explains the small cases:

- `W = 2`: the answer is `s[1] * s[2]`; there is only one component-tree
  edge, and we choose its two endpoints.
- `W = 3`: the answer is
  `s[1] * s[2] * s[3] * (s[1] + s[2] + s[3])`. The last factor selects which
  component is the middle vertex of the three-vertex tree, with the correct
  endpoint weight.
- `W = 1`: handle it separately and return `1 mod k`, because the general
  expression would have a negative exponent even though no new link is needed.

For sample 3, the component sizes are `2, 1, 1`, so `W = 3` and

```text
answer = (2 * 1 * 1) * 4^(3 - 2) = 8.
```

### Implementation

The DSU stores each component size in `cnt[root]`. The solution multiplies all
component sizes modulo `k`, then uses binary exponentiation to multiply by
`n^(W - 2) mod k`. Only multiplication is needed, so the method also works when
`k` is composite; no modular inverse is required. The code first handles
`k = 1`, for which every answer modulo `k` is `0`, and then handles `W = 1`.

DSU processing takes `O((n + m) alpha(n))` time, where `alpha` is the inverse
Ackermann function. Binary exponentiation takes `O(log W)` time. The total
space usage is `O(n)`.
