# L - Swap between Arrays

[Problem link](https://atcoder.jp/contests/aalc001/tasks/aalc001_l)

**Contest:** [AAL Contest 001: Let's use segtree!](https://atcoder.jp/contests/aalc001)

time limit: 2 sec

memory limit: 1024 MiB

score: 4 points

## Problem

You are given three sequences `A`, `B`, and `C`, each of length `N`.

Process `Q` queries. Each query is one of:

- `1 l r`: for every `k` from `l` to `r`, swap `A_k` and `B_k`.
- `2 l r`: for every `k` from `l` to `r`, swap `B_k` and `C_k`.
- `3 l r`: for every `k` from `l` to `r`, swap `C_k` and `A_k`.
- `4 i`: print the smallest `j` with `i ≤ j ≤ N` such that `A_i + ... + A_j > 0`. If no such `j` exists, print `N + 1`.

## Constraints

- `1 ≤ N ≤ 10^5`
- `1 ≤ Q ≤ 10^5`
- `-10^9 ≤ A_i, B_i, C_i ≤ 10^9`
- For a type `1`, `2`, or `3` query, `1 ≤ l ≤ r ≤ N`
- For a type `4` query, `1 ≤ i ≤ N`
- All input values are integers

## Input

```text
N Q
A_1 A_2 ... A_N
B_1 B_2 ... B_N
C_1 C_2 ... C_N
query_1
query_2
...
query_Q
```

A query is `1 l r`, `2 l r`, `3 l r`, or `4 i`.

## Output

Let `q` be the number of type `4` queries. Print `q` lines. The `j`-th line is the answer to the `j`-th type `4` query.

## Sample Input 1

```text
5 7
3 -4 5 -2 1
-1 6 -3 4 -5
2 -7 1 8 -6
4 2
1 2 4
4 1
2 3 5
4 3
3 1 3
4 2
```

## Sample Output 1

```text
3
1
4
4
```

The first query has `A_2 = -4 ≤ 0` and `A_2 + A_3 = 1 > 0`, so it prints `3`. After the second query, `A = (3, 6, -3, 4, 1)`. The third query has `A_1 = 3 > 0`, so it prints `1`.

## Sample Input 2

```text
6 7
0 -3 3 -5 7 -2
-4 2 -1 0 5 -8
6 -2 0 4 -3 1
4 1
4 6
2 2 6
1 3 5
4 3
3 6 6
4 6
```

## Sample Output 2

```text
5
7
4
7
```

## Solution

三个数组按位置对齐。一次区间交换只是在每个位置上重排 `A`、`B`、`C` 三个槽，查询又只看交换之后的 `A`。所以线段树的每个结点同时维护三套信息，懒标记是这三个槽的一个置换。

下标在树里用半开区间 `[l, r)`。查询 `1 l r` 对应的区间是 `[l-1, r)`。

### 结点

槽 `0, 1, 2` 分别是当前的 `A, B, C`。每个槽有两个值：

- `sum`：这个结点整段的和。
- `best`：从结点左端点出发的最大前缀和。

`best` 可以是负数。它回答的是「有没有一条从前端开始的前缀超过某个门槛」，所以取的是最大的那条，不是任意子段和。单点的 `best` 就是这个元素自己。

左右孩子合并时，整段和直接相加。从前端出发的最大前缀只有两种走法：停在左孩子内部，或者走完整个左孩子再接右孩子的一条前缀。

```text
best = max(left.best, left.sum + right.best)
```

右孩子内部、但不接触左端点的前缀，不会成为父结点的 `best`。

元素绝对值到 `10^9`，一段最多 `10^5` 个，和用 `int64`。

### 置换标记

`perm` 的含义是 `new[i] = old[perm[i]]`。三种操作都是对换：

- `1 l r` 交换 `A` 和 `B`：`{1, 0, 2}`，槽变成 `(B, A, C)`。
- `2 l r` 交换 `B` 和 `C`：`{0, 2, 1}`，槽变成 `(A, C, B)`。
- `3 l r` 交换 `C` 和 `A`：`{2, 1, 0}`，槽变成 `(C, B, A)`。

对换是奇置换。把三个槽一起轮换 `(A, B, C) -> (C, A, B)` 只会生成偶置换，第三次轮换才回到原点，中间那两次都不是题目要的交换。所以标记必须能表示全部 6 个置换，不能只记一个轮换次数。

`apply` 先用 `perm` 重排本结点已经算好的 `sum` 和 `best`。孩子先不动。若结点上还没有标记，新标记就是 `perm`。若已经有标记 `lazy`，孩子目前仍是更旧的顺序，而本结点的视图是 `lazy` 作用之后的结果。再套一层 `perm` 时：

```text
当前视图[i] = 孩子[lazy[i]]
新视图[i]   = 当前视图[perm[i]] = 孩子[lazy[perm[i]]]
```

新标记就是 `lazy ∘ perm`，下标写成 `composed[i] = lazy[perm[i]]`。

例如先交换 `B, C`，标记是 `{0, 2, 1}`，视图为 `(A, C, B)`。再交换 `A, B`，也就是对换前两个槽：`(A, C, B)` 变成 `(C, A, B)`。合成标记是 `{2, 0, 1}`。

下传时把同一标记 `apply` 到两个孩子，然后清掉本结点的标记。叶子没有孩子，`apply` 已经改过叶子上的值，下传时发现孩子下标越界就只清标记。

区间更新是普通线段树：整段被覆盖就 `apply`，否则下传、递归两半、再 `merge`。

### 查询

`4 i` 要的是最小的 `j`，使 `A_i + ... + A_j > 0`。比较是严格大于。`A_i` 自己大于 0 时，答案就是 `i`。

从根往下走，丢掉完全位于 `pos` 左边的结点（`pos` 是 0-based 的 `i-1`）。

- `pos` 落在右孩子里时，左孩子整个在起点之前，不能用它的前缀。只递归右孩子。
- `pos` 落在左孩子里时，先在左孩子里找。找到就返回。没找到时，递归带出来的 `leftSum` 是 `A[pos..mid)` 的和，而且这段里每一条前缀都 `<= 0`。右孩子全部在 `pos` 之后，它的 `best` 正是从 `mid` 往右的最大前缀。`leftSum + right.best > 0` 表示右孩子里存在一条前缀，接上 `leftSum` 之后第一次变成正数。

存在这样的前缀时，用 `walk` 找出最左的那条。门槛是 `need = -leftSum`，也就是前缀和要 `> need`。左孩子的 `best > need` 就继续往左；否则左边任何前缀都不够，把门槛减去左孩子的整段和，再到右边。走到叶子时，父结点的判断已经保证这个叶子会跨过门槛。

都不存在时返回 `-1`，输出 `N + 1`。找到的 0-based 下标加 1 变成题目的下标。

`A[pos]` 必须参加比较。叶子若直接返回 0，正的 `A_i` 会被跳过，后面的负和也会少掉这一项。

### 样例 1 的前两次查询

`A = (3, -4, 5, -2, 1)`。从 `i = 2` 开始，`A_2 = -4`，`A_2 + A_3 = 1 > 0`，答案是 `3`。

接着在 `[2, 4]` 上交换 `A` 和 `B`。`B = (-1, 6, -3, 4, -5)`，所以 `A` 变成 `(3, 6, -3, 4, 1)`。下一次从 `i = 1` 开始，`A_1 = 3 > 0`，答案是 `1`。

### 复杂度

每次交换或查询沿一条或数条树链向下，每层 `O(1)`。`N, Q ≤ 10^5`，总时间 `O((N + Q) log N)`。
