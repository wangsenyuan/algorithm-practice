# F - Run

[Problem link](https://atcoder.jp/contests/aalc001/tasks/aalc001_f)

**Contest:** [AAL Contest 001: Let's use segtree!](https://atcoder.jp/contests/aalc001)

time limit: 2 sec

memory limit: 1024 MiB

score: 3 points

## Problem

You are given a sequence `A` of `N` nonnegative integers.

Process `Q` queries. Each query is one of:

- `1 i x`: set `A_i` to `x`.
- `2 l r`: print the largest positive integer `k` such that some `p` satisfies
  `l ≤ p ≤ p+k-1 ≤ r` and `A_p = A_{p+1} = ... = A_{p+k-1}`.

That is the longest run of equal values inside `[l, r]`.

## Constraints

- `1 ≤ N ≤ 2 · 10^5`
- `1 ≤ Q ≤ 2 · 10^5`
- `0 ≤ A_i ≤ 10^9`
- For a type `1` query, `1 ≤ i ≤ N` and `0 ≤ x ≤ 10^9`
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
8 7
1 1 2 2 2 3 2 2
2 1 8
2 4 8
1 6 2
2 3 8
1 4 7
2 1 8
2 4 4
```

## Sample Output 1

```text
3
2
6
4
1
```

After the third query, `A = (1, 1, 2, 2, 2, 2, 2, 2)`, so the longest run inside
`[3, 8]` has length `6`.

## Sample Input 2

```text
6 7
0 1000000000 1000000000 0 0 0
2 1 6
2 2 5
1 3 0
2 2 6
1 2 0
2 1 6
1 6 1000000000
```

## Sample Output 2

```text
3
2
4
6
```

## Solution

一段连续相同的值，要么完全落在左半段，要么完全落在右半段，要么刚好跨过中点。跨过中点的那一段，只能是左段的相同值后缀接上右段的相同值前缀，而且两段相接处的值必须相等。线段树每个节点记下这些信息：

| 字段 | 含义 |
|---|---|
| `firstValue`, `lastValue` | 这段最左、最右的值 |
| `pref` | 从左端起、值都等于 `firstValue` 的长度 |
| `suf` | 从右端起、值都等于 `lastValue` 的长度 |
| `best` | 这段内部最长的连续相同值长度 |
| `len` | 区间长度 |

左右两段 `a`、`b` 拼接时：

```text
best = max(a.best, b.best)
if a.lastValue == b.firstValue {
    best = max(best, a.suf + b.pref)
}
pref = a.pref
suf  = b.suf
if a.pref == a.len && a.lastValue == b.firstValue {
    pref += b.pref
}
if b.suf == b.len && a.lastValue == b.firstValue {
    suf += a.suf
}
```

`firstValue` 取左边，`lastValue` 取右边，`len` 相加。前缀只有在左段整段都是同一个值、并且这个值等于右段开头时才会变长；后缀对称。叶子长度是 `1`，四个长度字段都是 `1`。

单点修改只改叶子的 `firstValue` 和 `lastValue`，再沿路径用同一公式重算。查询按半开区间 `[l, r)` 拆开，从左到右合并后取 `best`。

### 样例 1

`A = (1, 1, 2, 2, 2, 3, 2, 2)`。整段最长是中间三个 `2`，答案 `3`。`[4, 8] = (2, 2, 3, 2, 2)` 最长是 `2`。

把 `A_6` 改成 `2` 之后，`A = (1, 1, 2, 2, 2, 2, 2, 2)`，`[3, 8]` 全是 `2`，答案 `6`。再把 `A_4` 改成 `7`，整段变成 `(1, 1, 2, 7, 2, 2, 2, 2)`，最长是末尾四个 `2`。单点 `[4, 4]` 的答案是 `1`。

### 复杂度

建树、每次修改和每次查询都是 `O(log N)`。`N, Q ≤ 2 · 10^5`，总时间 `O((N + Q) log N)`。
