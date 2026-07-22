# D. Building Bridge

[Problem link](https://codeforces.com/problemset/problem/250/D)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Two villages are separated by a river with banks `x = a` and `x = b`
(`0 < a < b`).

The west village is at `O = (0, 0)`. There are `n` straight pathways to the left
bank ending at `A_i = (a, y_i)`.

The east village is somewhere on the east bank. There are `m` paths to the
right bank ending at `B_j = (b, y'_j)`, and the path from the east village to
`B_j` has length `l_j`.

Choose exactly one left point `A_i` and one right point `B_j`, and connect them
by a straight bridge, minimizing:

```text
|O A_i| + |A_i B_j| + l_j
```

where `|XY|` is Euclidean distance.

## Constraints

- `1 <= n, m <= 10^5`
- `0 < a < b < 10^6`
- `|y_i|, |y'_j| <= 10^6`, both sequences strictly ascending
- `1 <= l_j <= 10^6`
- No two `A_i` coincide; no two `B_j` coincide

## Input

```text
n m a b
y_1 y_2 ... y_n
y'_1 y'_2 ... y'_m
l_1 l_2 ... l_m
```

## Output

Print two integers — 1-based indices of the chosen left and right bank points.
Any optimal answer within relative/absolute error `1e-6` of the jury's path
length is accepted.

## Example

```text
Input
3 2 3 5
-2 -1 4
-1 2
7 3

Output
2 2
```

### ideas
1. 在左右选择(i, j), 使得 |OAi| + |AiBj| + Lj 最小
2. 第二项 = sqrt(d * d + dy * dy)
3. 选中j来说, 连接O-Bj, 存在一个点w, 就是查看这个点附近的i

## Solution

枚举右岸端点 `B_j = (b, y'_j)`。对固定的 `j`，需要在所有左岸端点
`A_i = (a, y_i)` 中最小化

```text
f_j(y_i) = sqrt(a^2 + y_i^2)
         + sqrt((b-a)^2 + (y'_j-y_i)^2)
         + l_j
```

其中 `l_j` 对当前 `j` 是常数，所以先考虑前两段路程。

### 连续情况下的最优交点

如果允许在直线 `x = a` 上任选一点 `A`，由三角不等式：

```text
|OA| + |AB_j| >= |OB_j|
```

等号成立当且仅当 `O`、`A`、`B_j` 共线，并且 `A` 位于线段
`OB_j` 上。

直线 `OB_j` 的斜率为 `y'_j / b`，因此它与 `x = a` 的交点纵坐标为

```text
y* = a * y'_j / b
```

这正是连续情况下唯一的最优位置。函数 `f_j(y)` 是凸函数：在 `y*`
左侧随 `y` 增大而减小，在 `y*` 右侧随 `y` 增大而增大。

### 离散端点只需检查两个邻居

实际只能选择给定的纵坐标 `y_i`。将所有左岸端点按 `y_i` 排序后，设
`p = lower_bound(y*)`，则离散最优点只可能是：

- 第一个满足 `y_i >= y*` 的点 `p`；
- `p` 的前一个点。

因为更靠左的点处在函数的单调递减区间，不会优于最靠近 `y*` 的左侧
端点；更靠右的点同理。

代码中的 `arr` 保存 `(原下标, y_i)` 并按 `y_i` 排序。对每个 `B_j`：

1. 计算 `ya = a * y'_j / b`；
2. 用二分搜索找到 `lower_bound(ya)`；
3. 用 `play` 计算该位置及其前驱对应的完整路程；
4. 用 `best` 维护所有 `j` 中的全局最小值，并保存两个原始下标。

`play(y0, y1, l)` 直接计算

```text
sqrt(a^2 + y0^2) + sqrt((b-a)^2 + (y1-y0)^2) + l
```

因此不同 `B_j` 的额外路径长度 `l_j` 也被正确纳入全局比较。

### Correctness Proof

对任意固定的 `B_j`，连续函数 `f_j(y)` 在直线 `OB_j` 与 `x = a`
的交点 `y* = a*y'_j/b` 处取得最小值，并在 `y*` 两侧分别单调递减、
单调递增。因此，在排序后的离散集合 `{y_i}` 中，最优的 `A_i` 必定是
`lower_bound(y*)` 或它的前驱。算法检查了这两个可能的端点，所以找到了
固定 `j` 时的最短路线。

算法对每个右岸端点 `B_j` 都执行上述检查，并比较包含 `l_j` 的完整路线
长度。因此最终保存的 `(i, j)` 在所有合法端点对中具有最小总长度，算法
正确。

### Complexity Analysis

- 排序左岸端点：`O(n log n)`；
- 对每个右岸端点进行一次二分和至多两次距离计算：`O(m log n)`；
- 总时间复杂度：`O(n log n + m log n)`；
- 额外空间复杂度：`O(n)`。
