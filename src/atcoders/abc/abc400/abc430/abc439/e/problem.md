# E - Kite

[Problem link](https://atcoder.jp/contests/abc439/tasks/abc439_e)

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 450 points

## Problem Statement

`N` people numbered `1` to `N` are trying to fly kites by the riverbank.
The river flows in a straight line; use a 2D coordinate system where the
`x`-axis is along the river and the `y`-axis is height.

Person `i` stands at `(A_i, 0)` and flies a kite at `(B_i, 1)`.

Persons `i` and `j` (`i ≠ j`) cannot fly at the same time if the line segment
connecting `(A_i, 0)` and `(B_i, 1)` intersects the segment connecting
`(A_j, 0)` and `(B_j, 1)` (including touching at endpoints).

What is the maximum number of people who can fly kites at the same time?

## Constraints

- `1 <= N <= 2 * 10^5`
- `0 <= A_i <= 10^9`
- `0 <= B_i <= 10^9`
- All input values are integers

## Input

```
N
A_1 B_1
A_2 B_2
...
A_N B_N
```

## Output

Output the maximum number of people who can fly kites at the same time.

## Samples

### Sample 1

Input:

```
3
3 5
1 4
2 6
```

Output:

```
2
```

### Sample 2

Input:

```
5
1 2
1 3
1 4
1 5
1 6
```

Output:

```
1
```

### Sample 3

Input:

```
10
440423913 766294629
725560240 59187619
965580535 585990756
550925213 623321125
549392044 122410708
21524934 690874816
529970099 244587368
757265587 736247509
576136367 993115118
219853537 21553211
```

Output:

```
4
```


### ideas
1. sort by A inc
2. assume a[i] <= b[i], then if A[i] < A[j], and A[j] <= B[i], then i & j can't fly at the same time
3. dp[i]表示到位置i(不是下标i)时的最优解
4. dp[i] = dp[j] + 1 (位置j后面和i之间可以放一只风筝)
5. 理解不大对,
6. 如果 a[i] < a[j], 且 b[i] < b[j] (ok)
7.

## Solution

每个人的风筝线连接下方的点 `(A_i, 0)` 和上方的点 `(B_i, 1)`。要判断两条线是否相交，
关键不是区间 `[A_i, B_i]` 是否重叠，而是两个人在上下两条水平线上的相对顺序是否一致。

### 两条风筝线何时不相交

考虑两个人 `i` 和 `j`，不妨先看 `A_i < A_j` 的情况。

在高度 `y` 处，第 `i` 条线的横坐标为：

```text
x_i(y) = (1-y) * A_i + y * B_i,  0 <= y <= 1
```

如果同时有 `B_i < B_j`，那么在下端和上端，第 `i` 条线都严格位于第 `j` 条线左边。
两个横坐标之差是关于 `y` 的一次函数，并且在 `y=0`、`y=1` 时都严格小于 `0`，所以在
整个高度范围内都不会变成 `0`，两条线不相交。

反过来，如果 `B_i >= B_j`：

- `B_i = B_j` 时，两条线在上端点接触；
- `B_i > B_j` 时，两条线在下端的顺序是 `i,j`，在上端的顺序变成 `j,i`，中间必然相交。

题目把端点接触也算作相交，所以两条线能够同时存在的充要条件是：

```text
A_i < A_j 且 B_i < B_j，
```

或者两个不等式同时反向。也就是说，选择出来的人按 `A` 严格递增排列后，`B` 也必须严格
递增。问题因此变成求二维点 `(A_i, B_i)` 的最长严格递增链。

### 排序与动态规划

先把所有点按照 `A` 从小到大排序。定义：

```text
dp[b] = 当前找到的、以压缩后的 B 坐标 b 结尾的最长合法链长度
```

对于当前点 `(A_i, B_i)`，设 `B_i` 离散化后的下标为 `j`。它前面只能接一个 `B` 严格
小于 `B_i` 的点，因此转移为：

```text
dp[j] = max(dp[j], max(dp[0..j-1]) + 1)
```

需要支持前缀最大值查询和单点取最大值，代码用树状数组 `bit` 完成：

- `bit.get(j-1)` 返回所有严格小于当前 `B_i` 的结尾状态的最大值；
- `bit.update(j, dp[j])` 把当前状态加入数据结构。

`B` 只影响大小关系，所以先排序、去重并离散化，不会改变合法链。

### 为什么相同的 `A` 必须分组处理

相同 `A` 的两条线在下端点接触，不能同时选择。因此，同一个 `A` 组内的状态不能互相转移。

代码先为整组计算所有 `dp[j]`，此时树状数组中只包含 `A` 更小的组；等整组计算完以后，
才统一执行 `bit.update`。这个延迟更新保证了转移中的 `A` 也是严格递增的。

同样，查询只到 `j-1`，而不是 `j`，保证 `B` 严格递增；相同 `B` 的两条线在上端点接触，
也不会被同时选入答案。

### Correctness Proof

按照 `A` 排序后，任意两个被选择的点对应的风筝线不相交，当且仅当它们的 `A` 和 `B`
顺序都严格一致。因此，任意合法选择恰好对应一个 `A`、`B` 都严格递增的链，反之亦然。

处理某个相同 `A` 的组时，树状数组恰好保存所有 `A` 更小的点形成的最优链。对于组内的点
`(A_i, B_i)`，查询 `B < B_i` 的最大值并加 `1`，枚举了所有能够合法接在它前面的链，
所以得到以该点结尾的最优答案。整组延迟更新排除了相同 `A` 的转移，严格前缀查询排除了
相同 `B` 的转移。

因此，算法计算了所有严格二维递增链的最大长度，也就是能够同时放风筝的最大人数。

### Complexity

排序和离散化需要 `O(N log N)` 时间。每个点进行一次树状数组查询和更新，每次为
`O(log N)`。

- 时间复杂度：`O(N log N)`；
- 空间复杂度：`O(N)`。
