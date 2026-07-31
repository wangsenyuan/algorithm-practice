# G. Criterion in Burlandia

[Problem link](https://codeforces.com/problemset/problem/2236/G)

**Contest:** [Codeforces Round (contest 2236)](https://codeforces.com/contest/2236)

time limit per test: 3 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

## Problem Statement

Regions in Burlandia form a tree with `n` vertices. Each region has a
friendliness value `a_i`.

There are `q` queries. In each query you are given two different vertices
`x` and `y`. Consider the unique path `v_1, v_2, ..., v_k` from `x` to `y`
(`v_1 = x`, `v_k = y`). Count the number of non-empty subsegments
`[l, r]` (`1 <= l <= r <= k`) of this path that are **hospitable**:

```text
a[v_l] ⊕ a[v_{l+1}] ⊕ ... ⊕ a[v_r]
  >= a[v_l] + a[v_{l+1}] + ... + a[v_r]
```

(Equivalently: the XOR of the friendliness values on the subsegment is at
least their sum.)

## Input

Each test consists of multiple test cases. The first line contains a single
integer `t` (`1 <= t <= 10^4`) — the number of test cases.

For each test case:

- the first line contains integers `n` and `q`
  (`2 <= n <= 10^5`, `1 <= q <= 10^5`);
- the second line contains `n` non-negative integers `a_1, ..., a_n`
  (`0 <= a_i < 2^20`);
- the next `n - 1` lines each contain integers `u, v` (`1 <= u, v <= n`) —
  tree edges;
- then `q` lines follow, each with integers `x, y`
  (`1 <= x, y <= n`, `x != y`) — a query.

It is guaranteed that the sum of `n` and the sum of `q` over all test cases
do not exceed `10^5`, and that the edges form a tree.

## Output

For each query, output the answer on a separate line.

## Sample Input 1

```text
3
4 3
0 0 4 1
1 2
1 3
1 4
1 4
2 3
2 4
4 3
0 4 1 2
1 3
1 4
2 4
1 2
2 3
2 4
4 3
3 2 4 4
1 2
2 4
3 4
1 2
1 3
2 3
```

## Sample Output 1

```text
3
6
6
6
10
3
2
5
4
```

## Note

In the third query of the third test case, the path from `2` to `3` is
`{2, 4, 3}`. Subsegments `[2; 3]` and `[1; 3]` fail the condition; the other
four subsegments succeed, so the answer is `4`.

## Solution Summary

### 1. Turn the inequality into a no-repeated-bit condition

For non-negative integers, XOR never exceeds their sum. Every bit carry makes
the sum strictly larger than the XOR. Therefore,

```text
XOR >= sum
```

is possible only when

```text
XOR = sum.
```

This equality holds exactly when no bit is set in two different values of the
subsegment. Thus a hospitable subsegment is a segment in which every one of the
20 bit positions occurs at most once.

For example, consider a path with values:

```text
1, 0, 1
```

Each side of the middle zero is valid by itself, but the whole path is invalid:
the lowest bit occurs in both `1`s.

An important consequence is that a valid segment contains at most 20 nonzero
vertices. It may contain arbitrarily many zero vertices, so the implementation
must skip or aggregate zero runs instead of visiting them one by one per query.

### 2. Rooted-path boundary `fp`

Root the tree at vertex `1`. During DFS, maintain
`lowDepthOfBit[b]`, the depth of the latest ancestor containing bit `b`.

For each vertex `u`, define `fp[u]` as the greatest forbidden starting depth
for a valid suffix of the root-to-`u` path:

```text
fp[u] = max(
    fp[parent[u]],
    lowDepthOfBit[b] for every bit b set in a[u]
)
```

Both unseen bit depths and the root's initial boundary are `-1`.

The propagation from `fp[parent[u]]` is necessary: a collision between two
earlier vertices still constrains segments ending at `u`, even if `a[u] = 0`.
After computing `fp[u]`, update the latest depth of every bit in `a[u]`.

A suffix ending at `u` and beginning at depth `s` is valid exactly when:

```text
s > fp[u].
```

Hence the number of valid root-path suffixes ending at `u` is:

```text
depth[u] - fp[u].
```

Let `pr[u]` be the prefix sum of these counts along the root-to-`u` path.

The latest-bit array is path-local DFS state. Each recursive call saves its own
20 entries and restores them before returning; otherwise, bit occurrences from
one child subtree could leak into a sibling subtree.

The recurrence also makes `fp` non-decreasing along every downward tree path.
This monotonicity allows query boundaries to be located with binary lifting.

### 3. Split a query at its LCA

For a query `(x, y)`, let:

```text
w = LCA(x, y).
```

Every path subsegment belongs to exactly one of three disjoint groups:

1. it lies strictly inside the `x`-arm below `w`;
2. it lies strictly inside the `y`-arm below `w`;
3. it contains `w`.

The first two groups use `fp` and `pr`. The third group needs bit masks from
both arms, because equal bits can collide across the LCA.

### 4. Count subsegments inside one arm

Consider the arm from `w` to an endpoint `x`, excluding `w`.

Find the deepest vertex `g` on this arm satisfying:

```text
fp[g] <= depth[w].
```

Because `fp` is monotone, binary lifting finds the first bad vertex and hence
its deepest good parent.

For every endpoint between the child of `w` and `g`, all starts strictly below
`w` are valid. If

```text
d = depth[g] - depth[w],
```

their total contribution is:

```text
1 + 2 + ... + d = d * (d + 1) / 2.
```

Every deeper endpoint `z` has `fp[z] > depth[w]`, so it contributes:

```text
depth[z] - fp[z].
```

Those contributions are obtained in `O(1)` from:

```text
pr[x] - pr[g].
```

Therefore:

```text
armCount(x, w)
    = d * (d + 1) / 2
    + pr[x] - pr[g].
```

The same calculation is performed for the `y`-arm.

### 5. Count subsegments containing the LCA

For an extension that includes `w`, the boundary is strict. On each arm, find
the deepest vertex `g` satisfying:

```text
fp[g] < depth[w].
```

Equality is not enough here: `fp[g] = depth[w]` means that the extension
collides with a bit in `a[w]`.

The segment from `w` through this strict-good `g` is internally valid, so it
contains at most 20 nonzero vertices. Precompute for every vertex its nearest
nonzero ancestor, then jump only through those nonzero vertices.

For one arm, build states:

```text
(mask, ways)
```

where:

- `mask` is the OR of values strictly below `w` up to an endpoint;
- `ways` is the number of endpoints producing that mask.

The mask excludes `a[w]`; the strict `fp` condition has already guaranteed that
the arm is compatible with `w`.

Zero runs do not change the mask. Suppose consecutive nonzero vertices occur at
depths:

```text
d1 < d2 < ... < dk
```

and the farthest strict-good endpoint has depth `D`. The state multiplicities
are obtained directly from depth gaps:

```text
mask 0:                 d1 - depth[w]
mask after d1:          d2 - d1
...
mask after dk:          D - dk + 1
```

If there is no nonzero vertex below `w`, the only state is:

```text
(0, D - depth[w] + 1).
```

Now combine the states of the two arms. A left endpoint and a right endpoint
form a valid subsegment containing `w` exactly when:

```text
left.mask & right.mask == 0.
```

Thus:

```text
crossCount =
    sum(left.ways * right.ways)
    over all state pairs with disjoint masks.
```

Returning to the example `1, 0, 1`, each arm has the choices `{mask 0, mask 1}`.
The pair `(1, 1)` is rejected because the masks intersect, while the other
three pairs are accepted.

### 6. Correctness argument

The algorithm is correct by the following facts:

1. A subsegment is hospitable exactly when no bit occurs twice.
2. `fp[u]` stores the greatest left boundary forced by any collision on the
   root-to-`u` path, so `s > fp[u]` exactly characterizes valid suffixes.
3. The arm formula counts all subsegments strictly below `w`: the triangular
   part covers endpoints whose boundary is at or above `w`, and the `pr`
   difference covers all deeper endpoints.
4. Every remaining subsegment contains `w` and is uniquely determined by one
   endpoint state from each arm.
5. Each arm state is internally valid and compatible with `w`; the combined
   segment is valid exactly when the two extra masks are disjoint.
6. The two arm-only groups and the `w`-containing group are disjoint and cover
   every path subsegment.

Therefore the returned count is exactly the number of hospitable subsegments.

### 7. Complexity

Binary-lifting and DFS preprocessing take:

```text
O(n log n + 20n)
```

time and `O(n log n)` memory.

Each query performs binary-lifting boundary searches, constructs at most 21
states per arm, and checks at most `21 * 21` state pairs:

```text
O(log n + 20^2)
```

time per query. Long zero paths contribute only through depth differences and
do not increase the number of states.
