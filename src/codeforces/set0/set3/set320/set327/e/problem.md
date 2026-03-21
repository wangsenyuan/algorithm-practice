# Problem

Iahub wants to meet his girlfriend Iahubina. They both live on the Ox axis (the horizontal axis). Iahub lives at point `0` and Iahubina at point `d`.

Iahub has `n` positive integers `a1, a2, ..., an`. The sum of those numbers is `d`. Suppose `p1, p2, ..., pn` is a permutation of `{1, 2, ..., n}`. Then, let `b1 = ap1`, `b2 = ap2` and so on. The array `b` is called a **route**. There are `n!` different routes, one for each permutation `p`.

Iahub's travel schedule is: he walks `b1` steps on the Ox axis, then he makes a break at point `b1`. Then, he walks `b2` more steps on the Ox axis and makes a break at point `b1 + b2`. Similarly, at the `j`-th time (`1 <= j <= n`) he walks `bj` more steps on the Ox axis and makes a break at point `b1 + b2 + ... + bj`.

Iahub is very superstitious and has `k` integers which give him bad luck. He calls a route **good** if he never makes a break at a point corresponding to one of those `k` numbers. For his own curiosity, answer how many good routes he can make, modulo `1000000007` (`10^9 + 7`).

## Input

- First line: integer `n` (`1 <= n <= 24`).
- Second line: `n` integers `a1, a2, ..., an` (`1 <= ai <= 10^9`).
- Third line: integer `k` (`0 <= k <= 2`).
- Fourth line: `k` positive integers — the numbers that give Iahub bad luck. Each does not exceed `10^9`.

## Output

Output a single integer — the answer to Iahub's dilemma modulo `1000000007` (`10^9 + 7`).

## Examples

### Example 1

Input

```text
3
2 3 5
2
5 7
```

Output

```text
1
```

### Example 2

Input

```text
3
2 2 2
2
1 3
```

Output

```text
6
```

## Note

In the first case consider six possible orderings:

- `[2, 3, 5]`. Iahub will stop at positions `2`, `5` and `10`. Among them, `5` is bad luck for him.
- `[2, 5, 3]`. Iahub will stop at positions `2`, `7` and `10`. Among them, `7` is bad luck for him.
- `[3, 2, 5]`. He will stop at the unlucky `5`.
- `[3, 5, 2]`. This is a valid ordering.
- `[5, 2, 3]`. He got unlucky twice (`5` and `7`).
- `[5, 3, 2]`. Iahub would reject, as it sends him to position `5`.

In the second case, note that it is possible that two different ways have the identical set of stopping points. In fact, all six possible ways have the same stops: `[2, 4, 6]`, so there's no bad luck for Iahub.


### ideas
1. 如果k = 0, ans = n!
2. k = 1, 那么只有一个位置, s, 假设存在一个序列sum (a[...]) = s(或者多个)
3. 那么这个序列，不能是前缀;  
4. dp[s]表示组成sum = s的方案数（set）
5. 把数组分成两半
6. dp[s1][i] 表示sum = s1, 且集合的size = i的计数
7. 然后就可以计算
8. 如果 k = 2呢？
9. 一个答案 = s1的那个，另外一个是, 排除掉s1, 但是出现s2的（否则s1就已经是bad了）

## Notes

这题的关键不是“按排列直接数”，而是先把“出现坏前缀和”转成“某个前缀集合的元素和等于指定值”。

### 1. 一个坏点怎么数

设坏点是 `s`。

如果一个排列在某个时刻第一次走到前缀和 `s`，那么一定存在一个前缀集合 `S`，满足：

- `sum(S) = s`

并且这个排列可以看成：

1. 先把 `S` 里的元素按任意顺序排在前面
2. 再把剩余元素按任意顺序排在后面

如果 `|S| = t`，那么这样的排列数就是：

- `t! * (n-t)!`

所以：

- 对每个满足 `sum(S)=s` 的子集 `S`，贡献是 `|S|! * (n-|S|)!`

于是只要我们知道：

- “和为 `s` 的子集有多少个，并按子集大小分类”

就能算出所有经过坏点 `s` 的排列数。

### 2. 为什么要按子集大小分类

因为不同大小的前缀集合，对应的排列数不同：

- 大小为 `t` 的前缀集合，贡献 `t! * (n-t)!`

所以不能只统计“有多少个子集和为 `s`”，还必须统计：

- `cnt[s][t] = 和为 s、大小为 t 的子集个数`

然后累加：

- `sum(cnt[s][t] * t! * (n-t)!)`

### 3. 两个坏点怎么数

设两个坏点为：

- `x < y`

一个排列同时经过 `x` 和 `y`，意味着它存在两个前缀：

- 前缀 `S1` 的和为 `x`
- 更长前缀 `S2` 的和为 `y`

并且显然：

- `S1 ⊂ S2`

于是等价地，我们可以拆成两个互不相交的集合：

- `A = S1`，满足 `sum(A) = x`
- `B = S2 \ S1`，满足 `sum(B) = y - x`

然后整个排列结构就是：

1. 先排 `A`
2. 再排 `B`
3. 最后排其余元素

如果：

- `|A| = p`
- `|B| = q`

那么这样的排列数就是：

- `p! * q! * (n-p-q)!`

所以“同时经过 `x` 和 `y`”的排列总数就是：

- 枚举所有不相交的 `(A, B)`
- 满足 `sum(A)=x, sum(B)=y-x`
- 每对贡献 `|A|! * |B|! * (n-|A|-|B|)!`

### 4. 为什么最后是 inclusion-exclusion

设：

- `Bad(x)` = 经过坏点 `x` 的排列集合
- `Bad(y)` = 经过坏点 `y` 的排列集合

那么好排列数就是：

- `All - Bad(x) - Bad(y) + Bad(x) ∩ Bad(y)`

因为题目中 `k <= 2`，所以只需要做到：

1. 会数单个坏点
2. 会数同时经过两个坏点

就结束了。

### 5. 为什么 meet-in-the-middle 合适

`n <= 24`，直接枚举全部子集是：

- `2^24`，勉强大但不太舒服

更自然的是折半：

- 左半最多 12 个数
- 右半最多 12 个数

这样：

- 每半的子集数最多 `2^12 = 4096`

就很轻松了。

对于一个坏点 `s`：

- 分别统计左右两半的 `sum -> 按大小计数`
- 再枚举左边和 `sumL`，去右边找 `s - sumL`

对于两个坏点：

- 每个元素有三种状态：放进 `A`、放进 `B`、或者都不放
- 所以每半最多 `3^12 = 531441` 种状态，也仍然可行

然后同样把左右两半拼起来。

### 6. 这题真正的思维转换

最重要的一步是把“坏前缀”看成“前缀元素集合”。

因为一旦这样看：

- 前缀和条件变成子集和条件
- 排列数只剩下一个简单的阶乘乘法

于是整题就从“排列上的路径限制”变成了：

- 子集和计数
- 按大小分类
- inclusion-exclusion

这就是实现里 `countOneBad` 和 `countTwoBad` 的本质。
