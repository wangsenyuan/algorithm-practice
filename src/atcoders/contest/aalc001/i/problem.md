# I - Subarray

[Problem link](https://atcoder.jp/contests/aalc001/tasks/aalc001_i)

**Contest:** [AAL Contest 001: Let's use segtree!](https://atcoder.jp/contests/aalc001)

time limit: 2 sec

memory limit: 1024 MiB

score: 4 points

## Problem

You are given a sequence `A` of `N` integers.

Process `Q` queries. Each query is one of:

- `1 i x`: set `A_i` to `x`.
- `2 l r`: from `(A_l, ..., A_r)`, choose between 0 and 2 nonempty contiguous subarrays that do not share positions. Print the maximum possible sum of the chosen elements.

Choosing nothing gives sum `0`.

## Constraints

- `1 ≤ N ≤ 10^5`
- `1 ≤ Q ≤ 10^5`
- `-10^9 ≤ A_i ≤ 10^9`
- For a type `1` query, `1 ≤ i ≤ N` and `-10^9 ≤ x ≤ 10^9`
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

Let `q` be the number of type `2` queries. Print `q` lines. The `j`-th line is the answer to the `j`-th type `2` query.

## Sample Input 1

```text
7 7
4 -5 3 -1 2 -6 5
2 1 7
2 2 6
1 4 10
2 1 7
1 6 1
2 1 7
2 2 2
```

## Sample Output 1

```text
9
5
20
25
0
```

The first query chooses `(3, -1, 2)` and `(5)`, summing to `9`. After the update `A_4 = 10`, the array is `(4, -5, 3, 10, 2, -6, 5)`, and `(3, 10, 2)` together with `(5)` sums to `20`. The last query is a single negative element, so choosing nothing gives `0`.

## Sample Input 2

```text
4 5
-8 -3 -5 -2
2 1 4
1 2 7
2 1 4
1 4 6
2 1 4
```

## Sample Output 2

```text
0
7
13
```

## Solution

线段树每个节点表示一个区间，维护至多两段不重叠子数组的最大和。一段可以是空的，空段的和是 `0`，所以负数不会被选进去。

节点里除了区间和 `sum`、最大前缀和 `pref`、最大后缀和 `suf`，还有这些值：

- `best1`：区间内一段子数组的最大和。
- `best`：至多两段的最大和。
- `bestPref`：从左端点开始的一段，再加上它右边至多一段。
- `bestSuf`：到右端点结束的一段，再加上它左边至多一段。
- `best2`：从左端点开始的一段，加上到右端点结束的一段。

只靠 `best` 和 `best1` 不够。两段分别在左右孩子里时，答案是 `left.best1 + right.best1`。还有一段跨过中点的情况：它由左边的一段后缀和右边的一段前缀拼成，另一段必须落在这段的左侧或右侧。所以

```text
best = max(left.best, right.best, left.best1 + right.best1,
           left.bestSuf + right.pref, left.suf + right.bestPref)
```

`best2` 不能写成 `left.pref + right.suf`。最大前缀会把后面更好的一段吃进去。例如前缀停在一个负数之前，剩下的 `(9, 9)` 比把最大前缀用完再往后找更优。跨过中点的那一段必须盖住整个右孩子，才能继续碰到父区间的右端点：

```text
best2 = max(left.pref + right.suf, left.best2 + right.sum, left.sum + right.best2, pref, suf)
```

`bestPref` 不要求第二段碰到右端点，所以可以把 `best2` 的右端那段再接到右孩子的一个前缀上。`bestSuf` 对称。

单点修改和区间查询都是普通线段树。`N, Q ≤ 10^5`，每次 `O(log N)`。
