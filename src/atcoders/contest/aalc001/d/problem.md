# D - Inversion Count

[Problem link](https://atcoder.jp/contests/aalc001/tasks/aalc001_d)

**Contest:** [AAL Contest 001: Let's use segtree!](https://atcoder.jp/contests/aalc001)

time limit: 2 sec

memory limit: 1024 MiB

score: 3 points

## Problem

You are given a sequence `A` of length `N`. Every element is `0`, `1`, or `2`.

Process `Q` queries. Each query is one of:

- `1 i x`: set `A_i` to `x`.
- `2 l r`: print the inversion count of `(A_l, A_{l+1}, ..., A_r)`.

The inversion count of a sequence `B` of length `k` is the number of pairs
`(i, j)` with `1 ≤ i < j ≤ k` and `B_i > B_j`.

## Constraints

- `1 ≤ N ≤ 2 · 10^5`
- `1 ≤ Q ≤ 2 · 10^5`
- `0 ≤ A_i ≤ 2`
- For a type `1` query, `1 ≤ i ≤ N` and `0 ≤ x ≤ 2`
- For a type `2` query, `1 ≤ l ≤ r ≤ N`
- All input values are integers

## Input

```text
N Q
A_1 A_2 ... A_N
query_1
query_2
...
query_Q
```

A query is either `1 i x` or `2 l r`.

## Output

Let `q` be the number of type `2` queries. Print `q` lines. The `j`-th line is
the answer to the `j`-th type `2` query.

## Sample Input 1

```text
5 5
2 0 1 2 0
2 1 5
2 2 4
1 2 2
2 1 3
2 2 5
```

## Sample Output 1

```text
5
0
2
4
```

Initially `A = (2, 0, 1, 2, 0)`. After the third query, `A = (2, 2, 1, 2, 0)`.

## Sample Input 2

```text
6 6
0 2 1 0 2 1
2 1 6
1 1 2
2 1 4
1 4 2
2 3 6
2 4 4
```

## Sample Output 2

```text
5
5
2
0
```

## Solution

值只有 `0, 1, 2`，所以一段区间只要记下三种值各有多少个，以及这段内部的逆序对数。线段树的每个节点存 `cnt[0..2]` 和 `val`。

左右两段 `a`、`b` 拼起来时，逆序对分成三段内部的和跨过中点的：

```text
val = a.val + b.val
    + a.cnt[2] * (b.cnt[0] + b.cnt[1])
    + a.cnt[1] * b.cnt[0]
```

跨过中点的只有两种：左边的 `2` 配右边的 `0` 或 `1`，左边的 `1` 配右边的 `0`。个数直接相加。

单点修改把叶子改成「只有一个 `x`，逆序对是 `0`」，再从叶子往上按同一公式重算。区间查询把 `[l, r)` 拆成若干段，按从左到右的顺序用同一公式合并。

### 样例 1

`A = (2, 0, 1, 2, 0)`。整段里左边的 `2` 能配到后面三个更小的数，中间的 `1` 能配到最后的 `0`，一共 `5`。`[2, 4] = (0, 1, 2)` 已经有序，答案是 `0`。

把 `A_2` 改成 `2` 之后，`A = (2, 2, 1, 2, 0)`。`[1, 3] = (2, 2, 1)` 有两个逆序对，`[2, 5] = (2, 1, 2, 0)` 有四个。

### 复杂度

建树、每次修改和每次查询都是 `O(log N)`。`N, Q ≤ 2 · 10^5`。
