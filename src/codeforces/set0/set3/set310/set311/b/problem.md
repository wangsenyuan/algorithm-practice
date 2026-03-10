# Problem

$n$ 个山丘排成一条线（编号 $1 \sim n$），山丘 $i$ 与 $i-1$ 距离 $d_i$。$p$ 个饲养员从山丘 $1$ 出发，以速度 $1$ 向山丘 $n$ 行进，途中可在任意山丘接走该处的猫。

$m$ 只猫，第 $i$ 只于时刻 $t_i$ 到达山丘 $h_i$ 并开始等待。饲养员到达山丘 $h_i$ 的时刻必须 $\ge t_i$ 才能接走该猫，否则猫的等待时间为饲养员到达时刻减去 $t_i$。

求：安排 $p$ 个饲养员的出发时间，使**所有猫的等待时间之和**最小。

**提示**：定义 $a_i = t_i - \sum_{j=2}^{h_i} d_j$，表示接走猫 $i$ 所需的最早出发时间。问题转化为将 $m$ 个 $a_i$ 分成 $p$ 段，每段代价为 $\sum (\max - a_i)$，求最小总代价。

## Input

- 第一行：$n$, $m$, $p$（$2 \le n \le 10^5$, $1 \le m \le 10^5$, $1 \le p \le 100$）
- 第二行：$d_2, d_3, \ldots, d_n$（$1 \le d_i < 10^4$）
- 接下来 $m$ 行：每行 $h_i$, $t_i$（$1 \le h_i \le n$, $0 \le t_i \le 10^9$）

## Output

最小等待时间之和。

## Examples

**Input:**

```
4 6 2
1 3 5
1 0
2 1
4 9
1 10
2 10
3 12
```

**Output:** `3`

## Divide And Conquer DP

### 1. 转化

先求出每只猫所在山丘到 `1` 号山丘的距离：

- `dist[h] = d2 + d3 + ... + dh`

如果一个饲养员在时刻 `s` 从 `1` 号山丘出发，那么他到达猫 `i` 所在山丘 `h_i` 的时刻是：

- `s + dist[h_i]`

要接到这只猫，必须满足：

- `s + dist[h_i] >= t_i`

等价于：

- `s >= t_i - dist[h_i]`

定义：

- `a_i = t_i - dist[h_i]`

那么 `a_i` 就表示“想让猫 `i` 不再继续等，最早要在什么时候出发”。

### 2. 为什么排序后一定是分段

把所有 `a_i` 排序后记为：

- `a_1 <= a_2 <= ... <= a_m`

如果某个饲养员负责一批猫，那么他的出发时间肯定取这批猫里最大的 `a`，也就是这批猫排序后最后一个值。

假设某个饲养员负责区间 `[l, r]`，那么最优出发时间就是：

- `a_r`

因为：

- 出发更早，接不到最后那只猫。
- 出发更晚，只会让这一段里的所有猫多等，没有任何好处。

于是这一段的总代价就是：

- `(a_r - a_l) + (a_r - a_{l+1}) + ... + (a_r - a_r)`
- `= (r - l + 1) * a_r - (prefix[r] - prefix[l-1])`

其中：

- `prefix[i] = a_1 + a_2 + ... + a_i`

所以整个问题就变成：

- 把排好序的 `m` 个数分成不超过 `p` 段
- 每段代价用上面的公式计算
- 求总代价最小

### 3. DP 定义

令：

- `dp[k][i]` = 把前 `i` 只猫分成恰好 `k` 段的最小代价

最后答案取：

- `min(dp[1][m], dp[2][m], ..., dp[p][m])`

因为题目给了 `p` 个饲养员，但允许有些人不接猫，所以本质上是“不超过 `p` 段”。

转移时，枚举最后一段从哪里开始。

如果最后一段是 `(j+1 ... i)`，那么前 `j` 只猫已经被前 `k-1` 段处理完了，所以：

- `dp[k][i] = min(dp[k-1][j] + cost(j+1, i))`
- 其中 `0 <= j < i`

代价函数：

- `cost(j+1, i) = (i-j) * a_i - (prefix[i] - prefix[j])`

因此：

- `dp[k][i] = min(dp[k-1][j] + (i-j) * a_i - (prefix[i] - prefix[j]))`

### 4. 为什么可以用 Divide And Conquer Optimization

把上式写成：

- `dp[k][i] = i * a_i - prefix[i] + min(dp[k-1][j] + prefix[j] - j * a_i)`

虽然这看起来也能往斜率优化上推，但本题更直接的做法是用分治优化。

关键点是：

- 排序后的这类分段代价满足四边形不等式 / Monge 性质
- 因此最优决策点 `opt[k][i]` 随着 `i` 增大是单调不减的

这就允许我们在计算一整层 `dp[k][*]` 时，用分治只在一个缩小后的决策区间里找最优 `j`。

分治函数的含义是：

- 计算 `dp[k][left...right]`
- 已知这些位置的最优决策点一定落在 `[optL, optR]` 之间

对中点 `mid`：

1. 枚举 `j in [optL, min(mid-1, optR)]`
2. 找到使 `dp[k-1][j] + cost(j+1, mid)` 最小的 `bestPos`
3. 递归左半边时，最优点范围变成 `[optL, bestPos]`
4. 递归右半边时，最优点范围变成 `[bestPos, optR]`

这样每层的复杂度是：

- `O(m log m)` 到 `O(m)` 之间的常数级分治开销

总复杂度通常记为：

- `O(p * m * log m)`

在本题里 `p <= 100`，`m <= 1e5`，可以通过。

### 5. 代码里的几个数组分别表示什么

在 [solution.go](solution.go) 里：

- `a[i]`
  - 排序后的 `t_i - dist[h_i]`

- `pref[i]`
  - `a` 的前缀和

- `prev[i]`
  - 上一层 DP，也就是 `dp[k-1][i]`

- `cur[i]`
  - 当前层 DP，也就是 `dp[k][i]`

主循环：

```go
for groups := 1; groups <= p; groups++ {
    compute(cur, prev, pref, a, groups, groups, m, groups-1, m-1)
    ans = min(ans, cur[m])
    prev, cur = cur, prev
}
```

含义是：

- 枚举使用 `1..p` 段
- 每次通过 `compute` 用分治算出这一层
- 用 `ans` 维护“至多 `p` 段”的最优值

### 6. `compute` 里做了什么

核心代码：

```go
mid := (left + right) / 2
bestPos := -1
bestVal := inf
upper := min(mid-1, optR)
for j := optL; j <= upper; j++ {
    cand := prev[j] + (mid-j)*a[mid] - (pref[mid] - pref[j])
    if cand < bestVal {
        bestVal = cand
        bestPos = j
    }
}
dp[mid] = bestVal
compute(... left half ..., optL, bestPos)
compute(... right half ..., bestPos, optR)
```

这正对应前面的分治优化模板：

- 中点暴力找最优决策
- 再利用决策单调性，把左右两边的搜索区间缩小

### 7. 为什么第一层可以单独写

当 `groups == 1` 时，不需要枚举断点：

- `dp[1][i] = cost(1, i)`
- `= i * a_i - prefix[i]`

所以代码里直接写成：

```go
if groups == 1 {
    bestVal = mid*a[mid] - pref[mid]
    bestPos = 0
}
```

这就是第一层的初始化。

### 8. 一个容易错的点

`a_i = t_i - dist[h_i]` 可以是负数。

这表示：

- 如果想让这只猫完全不等待，饲养员需要在时间 `0` 之前就出发

题目并没有限制出发时间必须非负，所以：

- 不能把出发时间强行写成 `max(a_i, 0)`

之前错误实现的问题就出在这里。  
本题允许负出发时刻，这一点必须保留。

## CHT Optimization

这一题也可以用 Convex Hull Trick。

### 1. 从转移式改写成直线查询

原始转移是：

- `dp[k][i] = min(dp[k-1][j] + (i-j) * a_i - (prefix[i] - prefix[j]))`

把和 `j` 无关的部分提出来：

- `dp[k][i] = i * a_i - prefix[i] + min(dp[k-1][j] + prefix[j] - j * a_i)`

对于固定的 `j`，定义一条直线：

- 斜率 `m = -j`
- 截距 `b = dp[k-1][j] + prefix[j]`

那么这一项就变成：

- `m * x + b`

其中查询点：

- `x = a_i`

所以每层 DP 做的事情就是：

1. 按 `j` 从小到大插入直线
2. 按 `i` 从小到大查询 `x = a_i`
3. 取最小值

### 2. 为什么可以用单调队列维护凸壳

这里有两个关键单调性：

- 插入顺序里，`j` 单调递增，所以斜率 `-j` 单调递减
- 排序后 `a_i` 单调递增，所以查询点 `x` 单调递增

这正是最经典的 CHT 使用条件之一：

- 单调斜率插入
- 单调查询

因此不需要 Li Chao Tree，也不需要二分交点，直接用一个 deque 就够了。

### 3. 为什么判断“中间直线没用”可以用叉乘式子

假设凸壳尾部有三条线 `A, B, C`，如果 `B` 永远不会成为最优线，就应该把它删掉。

在代码里用了等价判定：

- `(c.b-a.b) * (a.m-b.m) <= (b.b-a.b) * (a.m-c.m)`

这其实就是用交点顺序比较来判断：

- 若 `intersection(A, B) >= intersection(A, C)`，则 `B` 无用

因为本题里斜率是单调的，这样维护出的就是下凸壳。

### 4. 查询时为什么可以弹队头

查询点 `x = a_i` 单调递增。

如果当前队头第二条线在 `x` 处已经不比第一条差，那么随着 `x` 继续增大，第一条线以后也不会更优，所以可以永久丢掉。

代码里对应：

```go
for head+1 < len(hull) && value(hull[head+1], x) <= value(hull[head], x) {
    head++
}
```

这样每条线最多进队一次、出队一次，所以每层是线性的。

### 5. `solve_cht` 的整体复杂度

每一层 DP：

- 插入 `O(m)` 条线
- 查询 `O(m)` 次

所以每层是：

- `O(m)`

总复杂度：

- `O(p * m)`

再加上最开始对 `a_i` 的排序：

- `O(m log m)`

总计：

- `O(m log m + p * m)`

这比 divide and conquer 的 `O(p * m * log m)` 更快一些。

### 6. 和当前代码的对应关系

在 [solution.go](solution.go) 里，`solve_cht` 的局部定义对应如下：

- `line{m, b}`
  - 一条直线 `y = m*x + b`

- `value`
  - 计算某条线在 `x` 处的值

- `bad`
  - 判断中间那条线是否可以删掉

- `addLine`
  - 向凸壳尾部插入一条新线

- `advanceHead`
  - 在当前查询点 `x` 下，把队头移动到最优位置

然后每层主循环的逻辑是：

1. 先把合法的 `j` 转成直线插入
2. 再用这条壳去计算 `dp[k][j+1]`

由于插入和查询都是按顺序进行，所以整层只需线性时间。
