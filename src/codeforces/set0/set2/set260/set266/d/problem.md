# Problem

**题意简述**：给定一个连通无向带权图，点表示路口，边表示道路，餐厅可以建在任意路口或道路上的任意点。定义一个位置的代价为它到所有路口的最短路距离中的最大值。要求找到一个位置，使这个最大值最小，并输出该最小值。


## Input

- The first line contains two integers `n` and `m` — the number of junctions and the number of roads.
- Then `m` lines follow, each describing one road by three integers `ai`, `bi`, `wi`.
- `ai` and `bi` are the endpoints of the road, and `wi` is its length.
- Constraints visible in the statement: `1 ≤ ai, bi ≤ n`, `ai ≠ bi`, `1 ≤ wi ≤ 10^5`.
- It is guaranteed that each road connects two distinct junctions, there is at most one road between any two junctions, and the graph is connected.

## Output

Print a single real number — the minimum possible distance from the optimal restaurant location to the farthest junction.

The answer is accepted if its absolute or relative error does not exceed `10^-9`.

## Examples

### Example 1

**Input**

```text
2 1
1 2 1
```

**Output**

```text
0.50
```

### Example 2

**Input**

```text
3 3
1 2 1
2 3 1
1 3 1
```

**Output**

```text
1.00
```

### Example 3

**Input**

```text
3 2
1 2 100
2 3 1
```

**Output**

```text
50.50
```

## Algorithm

### Key Insight

这题不是直接在边上找一个“最优点”去三分，而是：

- 先固定答案 `R`
- 判断是否存在某个点或某条边上的某个位置，使得到所有点的距离都不超过 `R`

对点来说很好判断；对边来说，可以把“不满足某个点要求”的位置表示成一段禁止区间。只要所有点对应的禁止区间并起来**没有覆盖整条边**，就说明这条边上存在合法位置。

因此整题的核心是：

- **最短路预处理**
- **二分答案**
- **边上做区间覆盖判定**

### 1. 先做所有点对最短路

设 `dist[u][v]` 表示点 `u` 到点 `v` 的最短路长度。

因为餐厅可以建在点上，也可以建在边上，所以后续所有判断都只需要基于这些点到点最短路来完成。

当前实现里是对每个起点跑一次 Dijkstra。

### 2. 把答案乘 2

边权是整数，而最优点可能在边中点，所以答案可能是 `x.5`。

为了避免浮点误差，可以把所有距离都乘 2：

- 点到点距离变成 `2 * dist[u][v]`
- 边长 `w` 变成 `2 * w`
- 最终答案一定是整数，最后再除以 2 输出

### 3. 二分答案

设当前二分检查的半径为 `limit`（这里已经是乘 2 之后的值）。

问题变成：

- 是否存在某个点，或者某条边上的某个位置，使得到所有点的距离都不超过 `limit`

如果某个 `limit` 可行，那么更大的答案也一定可行，所以可以二分最小可行值。

### 4. 检查一个点是否可行

如果餐厅建在点 `u`，那么它的最差距离就是：

- `max_v 2 * dist[u][v]`

只要这个值 `<= limit`，当前答案就可行。

### 5. 检查一条边上是否存在可行点

考虑边 `(u, v)`，长度为 `w`，并把它也乘 2，记为 `2w`。

设餐厅位于这条边上，距离端点 `u` 的位置为 `t`，其中：

- `0 <= t <= 2w`

那么它到某个点 `i` 的最短距离是：

- `min(2 * dist[u][i] + t, 2 * dist[v][i] + 2w - t)`

要让这个点 `i` 满足要求，必须：

- `min(2 * dist[u][i] + t, 2 * dist[v][i] + 2w - t) <= limit`

反过来说，不满足要求时，说明两边都大于 `limit`：

- `2 * dist[u][i] + t > limit`
- `2 * dist[v][i] + 2w - t > limit`

整理后可得一段禁止区间：

- `t > limit - 2 * dist[u][i]`
- `t < 2 * dist[v][i] + 2w - limit`

也就是说，对每个点 `i`，我们都能求出一个在边上“不允许放餐厅”的区间。

### 6. 合并禁止区间

对一条边 `(u, v)`：

1. 枚举所有点 `i`
2. 算出它对应的禁止区间
3. 把所有禁止区间排序并合并

如果这些禁止区间**不能覆盖整条边**，说明存在某个位置 `t`：

- 对所有点都满足距离 `<= limit`

那么当前 `limit` 可行。

### 7. 为什么不能三分

一个容易犯的错误是：

- 认为边上的目标函数
  `max_i min(dist[u][i] + t, dist[v][i] + w - t)`
  一定是单峰函数，然后对它三分

这是不对的。

这个函数一般只是分段线性的上包络，不保证单峰，因此三分搜索可能错过真正最优点。

所以这里使用的是：

- **二分答案**
- **边上做可行性判定**

这样才是正确的。

### 8. 复杂度

设点数为 `n`，边数为 `m`。

- `n` 次 Dijkstra：`O(n * m log n)`
- 每次二分检查：
  - 检查所有点：`O(n^2)`
  - 检查所有边并合并区间：`O(m * n log n)`

总复杂度约为：

- `O(n * m log n + log W * (n^2 + m * n log n))`

其中 `W` 是答案范围。
