# E - RGB

[Problem link](https://atcoder.jp/contests/aalc001/tasks/aalc001_e)

**Contest:** [AAL Contest 001: Let's use segtree!](https://atcoder.jp/contests/aalc001)

time limit: 2 sec

memory limit: 1024 MiB

score: 3 points

## Problem

You are given a sequence `A` of `N` nonnegative integers and a string `S` of length
`N` made of `R`, `G`, and `B`.

Process `Q` queries. Each query is one of:

- `1 i x`: set `A_i` to `x`.
- `2 l r`: print the sum of `A_i * A_j * A_k` over all triples `l ≤ i < j < k ≤ r`
  whose characters `S_i, S_j, S_k` are `R`, `G`, `B` in that order, modulo
  `998244353`.

`S` never changes.

## Constraints

- `1 ≤ N ≤ 2 · 10^5`
- `1 ≤ Q ≤ 2 · 10^5`
- `0 ≤ A_i ≤ 10^9`
- `S` has length `N` and consists of `R`, `G`, and `B`
- For a type `1` query, `1 ≤ i ≤ N` and `0 ≤ x ≤ 10^9`
- For a type `2` query, `1 ≤ l ≤ r ≤ N`
- All input values are integers

## Input

```text
N Q
A_1 A_2 ... A_N
S
query_1
query_2
...
query_Q
```

A query is either `1 i x` or `2 l r`.

## Output

Let `q` be the number of type `2` queries. Print `q` lines. The `j`-th line is
the answer to the `j`-th type `2` query, modulo `998244353`.

## Sample Input 1

```text
5 5
1 2 3 4 5
RGBGB
2 1 5
1 2 10
2 1 5
2 2 4
2 1 3
```

## Sample Output 1

```text
36
100
0
30
```

For the first query the triples are `(1, 2, 3)`, `(1, 2, 5)`, and `(1, 4, 5)`.
Their products sum to `1*2*3 + 1*2*5 + 1*4*5 = 36`.

## Sample Input 2

```text
3 5
1000000000 1000000000 1000000000
RGB
2 1 3
1 2 998244353
2 1 3
1 2 998244354
2 1 3
```

## Sample Output 2

```text
747068572
0
716070898
```

## Solution

`S` 不变，每个位置只属于 `R`、`G`、`B` 里的一种。线段树的一段只要记下这三种颜色的权值和，以及已经配好的前缀乘积：

| 字段 | 含义 |
|---|---|
| `r`, `g`, `b` | 该颜色位置上 `A` 的和 |
| `rg` | 所有 `i < j` 且颜色是 `R, G` 的 `A_i * A_j` 之和 |
| `gb` | 所有 `j < k` 且颜色是 `G, B` 的 `A_j * A_k` 之和 |
| `val` | 所有 `i < j < k` 且颜色是 `R, G, B` 的 `A_i * A_j * A_k` 之和 |

左右两段 `a`、`b` 拼接时，三元组要么整段在左边，要么整段在右边，要么跨过中点。跨过中点只有两种：左边已经配好的 `RG` 乘右边的一个 `B`，或者左边的一个 `R` 乘右边已经配好的 `GB`。

```text
val = a.val + b.val + a.rg * b.b + a.r * b.gb
rg  = a.rg  + b.rg  + a.r  * b.g
gb  = a.gb  + b.gb  + a.g  * b.b
```

`r`、`g`、`b` 直接相加。所有运算都在模 `998244353` 下进行。

叶子只在自己的颜色上放 `A_i`，另外五个字段是 `0`。单点修改改掉这个权值，再沿路径用同一公式重算。查询把 `[l, r)` 拆开，从左到右合并后取 `val`。

### 样例 1

`A = (1, 2, 3, 4, 5)`，`S = RGBGB`。满足 `R, G, B` 顺序的三元组是 `(1, 2, 3)`、`(1, 2, 5)`、`(1, 4, 5)`，乘积和 `6 + 10 + 20 = 36`。

把 `A_2` 改成 `10` 之后，前两个三元组的乘积变成 `30` 和 `50`，总和是 `100`。`[2, 4] = GBG` 里没有 `R`，答案是 `0`。`[1, 3]` 只剩 `(1, 2, 3)`，乘积是 `30`。

### 复杂度

每个位置只更新、查询一次路径，`N, Q ≤ 2 · 10^5`，总时间 `O((N + Q) log N)`。
