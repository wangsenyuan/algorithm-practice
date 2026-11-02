# O - Minimum Value

[Problem link](https://atcoder.jp/contests/aalc001/tasks/aalc001_o)

**Contest:** [AAL Contest 001: Let's use segtree!](https://atcoder.jp/contests/aalc001)

time limit: 3 sec

memory limit: 1024 MiB

score: 4 points

## Problem

You are given a sequence `A` of `N` nonnegative integers.

Process `Q` queries. Each query gives `l` and `r`. Print

```text
sum over l ≤ s ≤ t ≤ r of min(A_s, A_{s+1}, ..., A_t)
```

## Constraints

- `1 ≤ N ≤ 2 · 10^5`
- `1 ≤ Q ≤ 2 · 10^5`
- `0 ≤ A_i ≤ 10^6`
- `1 ≤ l ≤ r ≤ N`
- All input values are integers

## Input

```text
N Q
A_1 A_2 ... A_N
l_1 r_1
l_2 r_2
...
l_Q r_Q
```

## Output

Print `Q` lines. The `i`-th line is the answer to the `i`-th query.

## Sample Input 1

```text
4 4
3 1 4 2
1 1
1 2
2 4
1 4
```

## Sample Output 1

```text
3
5
11
17
```

The first query is the single element `3`. The second query adds the minima of `(1, 1)`, `(1, 2)`, and `(2, 2)`, which are `3`, `1`, and `1`.

## Sample Input 2

```text
5 3
5 2 7 1 4
2 4
1 3
3 5
```

## Sample Output 2

```text
14
20
15
```

## Sample Input 3

```text
6 4
0 1000000 2 999999 3 4
1 6
2 5
2 2
4 6
```

## Sample Output 3

```text
2000031
2000017
1000000
1000015
```

## Solution

每个子数组的最小值恰好出现在一个最左位置。把答案拆成每个下标 `i` 的贡献，再按查询端点把贡献分成四块。四块都是查询平面上的矩形加，用一次从左到右的扫描和九棵树状数组读出来。

下标在下面都是 0-based。题目的查询 `[l, r]` 对应 `[ql, qr] = [l-1, r-1]`。

### 谁是最左最小值

`left[i]` 是 `i` 左边最近的 `j`，满足 `A_j ≤ A_i`；没有就是 `-1`。`right[i]` 是 `i` 右边最近的 `j`，满足 `A_j < A_i`；没有就是 `N`。单调栈各扫一遍。

左边用 `≤`，右边用 `<`。相等的值归左边那个位置，每个子数组只被统计一次。

记 `p = left[i] + 1`，`q = right[i] - 1`。`i` 是 `[s, t]` 的最左最小值，当且仅当

```text
p ≤ s ≤ i ≤ t ≤ q
```

`A_i = 0` 时贡献是 0，直接跳过。

### 一次查询切出四块

查询 `[ql, qr]` 里，`i` 必须满足 `ql ≤ i ≤ qr`。起点只能从 `max(ql, p)` 取到 `i`，终点只能从 `i` 取到 `min(qr, q)`。贡献是

```text
A_i · (i - max(ql, p) + 1) · (min(qr, q) - i + 1)
```

`fullL = i - p + 1`，`fullR = q - i + 1`。按查询是否盖住 `p` 和 `q` 分成四种，互不重叠：

- 两边都盖住，`ql ≤ p` 且 `qr ≥ q`。贡献是常数 `A_i · fullL · fullR`。这样的查询是 `ql ∈ [0, p]`，`qr ∈ [q, N-1]`。
- 只切右端，`ql ≤ p` 且 `i ≤ qr < q`。贡献是 `C · (qr - i + 1)`，其中 `C = A_i · fullL`。展开成 `(qr + 1) · C - C · i`。查询矩形是 `ql ∈ [0, p]`，`qr ∈ [i, q-1]`。`q = i` 时右边没有可切的位置，这块是空的。
- 只切左端，`p < ql ≤ i` 且 `qr ≥ q`。贡献是 `D · (i - ql + 1)`，其中 `D = A_i · fullR`。展开成 `D · (i + 1) - ql · D`。查询矩形是 `ql ∈ [p+1, i]`，`qr ∈ [q, N-1]`。`p = i` 时左边没有可切的位置。
- 两边都切，`p < ql ≤ i` 且 `i ≤ qr < q`。贡献是

```text
A_i · (i - ql + 1) · (qr - i + 1)
  = (qr + 1) · A_i · (i + 1)
    - A_i · i · (i + 1)
    - ql · (qr + 1) · A_i
    + ql · A_i · i
```

查询矩形是 `ql ∈ [p+1, i]`，`qr ∈ [i, q-1]`。

展开之后，式子里有两类数。`A_i`、`i`、`fullL`、`fullR` 只跟位置 `i` 有关，对这一块里的每个查询都一样。`ql` 和 `qr` 才跟这次查询有关。所以先把只跟 `i` 有关的数记到「哪些查询会用到 `i` 的这一块」上，查询到来时再把这对 `(ql, qr)` 代进去。

以 `A = (3, 1, 4, 2)` 的 `i = 1` 为例。`A_1 = 1`，`p = 0`，`q = 3`，`fullL = 2`，`fullR = 3`。只切左端的那一块，贡献是

```text
D · (i + 1) - ql · D = 6 - 3 · ql
```

这里的 `6` 和 `3` 是 `D · (i + 1)` 和 `D`，算完就不再变。会用到这块的查询是 `ql = 1` 且 `qr ∈ [3, 3]`，也就是只有 `(ql, qr) = (1, 3)`。代入 `ql = 1` 得到 `6 - 3 = 3`。

一次查询的答案，是每个 `i` 在自己那一块里代入这对 `(ql, qr)` 之后的和。

### 矩形加、点查

按 `ql` 从左往右扫描。一个矩形在左端点 `l1` 加上，在 `l2 + 1` 减去。同一时刻，树状数组只保留「左端点已经覆盖当前 `ql`」的矩形。

右端点用差分：在 `r1` 加 `v`，在 `r2 + 1` 减 `v`。前缀 `qr` 就是所有盖住这个右端点的值之和。九棵树分别存：

- 常数块：`A_i · fullL · fullR`
- 切右端：`C` 和 `C · i`
- 切左端：`D · (i + 1)` 和 `D`
- 两边都切：`A_i`、`A_i · i`、`A_i · (i + 1)`、`A_i · i · (i + 1)`

读出前缀后按上面的展开式组合。`ql · (qr + 1)` 不超过 `N^2`，放得进 `int64`。

### 样例 1 的第三问

`A = (3, 1, 4, 2)`。四个位置的 `(p, q)` 是 `(0, 0)`、`(0, 3)`、`(2, 2)`、`(2, 3)`。

第三问是 1-based 的 `[2, 4]`，也就是 `[ql, qr] = [1, 3]`。

- `i = 1`：左端被切、右端盖住。`D = 3`，贡献 `3 · (1 - 1 + 1) = 3`。对应子数组 `[1, 1]`、`[1, 2]`、`[1, 3]`，最小值都是 `1`。
- `i = 2`：两边都盖住。贡献 `4`，只有 `[2, 2]`。
- `i = 3`：两边都盖住。`fullL = 2`，`fullR = 1`，贡献 `2 · 2 · 1 = 4`。对应 `[2, 3]` 和 `[3, 3]`。

合计 `3 + 4 + 4 = 11`。

### 128 位

单个系数最大大约是 `A_i · N^2 ≤ 4 · 10^16`，放得进 `int64`。许多位置加在同一棵树上时，前缀和可以到 `10^21` 以上，展开式里的正项和负项各自都会超过 `2^63`，相减之后才回到答案。树状数组因此用有符号 128 位。最终答案不超过 `10^6 · N · (N + 1) / 2 ≤ 2 · 10^16`，低 64 位就是输出。

### 复杂度

每个位置最多写常数个矩形，每个矩形是两次差分。扫描时每个查询做常数次前缀和。`N, Q ≤ 2 · 10^5`，总时间 `O((N + Q) log N)`。

