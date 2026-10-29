# K - Multiset

[Problem link](https://atcoder.jp/contests/aalc001/tasks/aalc001_k)

**Contest:** [AAL Contest 001: Let's use segtree!](https://atcoder.jp/contests/aalc001)

time limit: 2 sec

memory limit: 1024 MiB

score: 4 points

## Problem

`S` is a multiset of integers. It starts with `N` elements `A_1, A_2, ..., A_N`.

Process `Q` queries. Each query is one of:

- `1 x`: insert one copy of `x` into `S`.
- `2 x`: delete one copy of `x` from `S`. `S` contains at least one `x` just before this query.
- `3`: sort the elements of `S` in nondecreasing order as `x_1, x_2, ..., x_k`, keeping duplicates. Print the sum of `i * x_i` modulo `998244353`.

An empty multiset has `k = 0`, and the sum is `0`.

## Constraints

- `1 ≤ N ≤ 2 · 10^5`
- `1 ≤ Q ≤ 2 · 10^5`
- `0 ≤ A_i ≤ 10^9`
- For a type `1` or `2` query, `0 ≤ x ≤ 10^9`
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

A query is `1 x`, `2 x`, or `3`.

## Output

Let `q` be the number of type `3` queries. Print `q` lines. The `j`-th line is the answer to the `j`-th type `3` query.

## Sample Input 1

```text
4 8
3 1 3 2
3
1 0
3
2 3
3
2 3
3
1 2
```

## Sample Output 1

```text
26
35
20
8
```

The initial multiset sorts to `(1, 2, 3, 3)`, so the first query prints `1*1 + 2*2 + 3*3 + 4*3 = 26`.

## Sample Input 2

```text
1 5
998244353
3
2 998244353
3
1 1000000000
3
```

## Sample Output 2

```text
0
0
1755647
```

The third query runs while `S` is empty, so it prints `0`.

## Solution

排序后的加权和是 `sum i * x_i`。不同的值最多 `N + Q` 个，先离散化。两棵树状数组分别记每个值的个数和元素和。个数是真实个数，元素和与答案都在模 `998244353` 下维护。

小于等于 `x` 的元素下标不变。设它们有 `cnt` 个，严格大于 `x` 的元素和是 `tail`。

- 插入时 `x` 还不在这 `cnt` 个里面，新下标是 `cnt + 1`，后面每个元素的下标加 1。增量是 `(cnt + 1) * x + tail`。
- 删除时这 `cnt` 个已经包含被删的那一份，它自己的下标就是 `cnt`。增量是 `-cnt * x - tail`。

`cnt` 是个数，不能拿去当下标。`tail` 要从 `x` 的下一个离散值查到末尾。相等的值算进 `cnt`，不算进 `tail`。

### 复杂度

每次查询是两次前缀和。`N, Q ≤ 2 · 10^5`，总时间 `O((N + Q) log (N + Q))`。
