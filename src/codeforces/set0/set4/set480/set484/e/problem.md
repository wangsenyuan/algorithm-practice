# Sign on Fence

## 题意

长度为 n 的数组 h，每个元素表示一个单位宽度条块的高度。

每次询问 (l, r, w)：在区间 [l, r] 内选一个长度为 w 的连续子区间，矩形高度为该子区间内所有高度的**最小值**，求所有选法中矩形高度的**最大值**。

## Input

- 第一行：n (1 ≤ n ≤ 10^5)
- 第二行：n 个整数 hi (1 ≤ hi ≤ 10^9)
- 第三行：m (1 ≤ m ≤ 10^5)
- 接下来 m 行：每行 l, r, w (1 ≤ l ≤ r ≤ n, 1 ≤ w ≤ r - l + 1)

## Output

对每个询问输出一行：最大可能的矩形高度。

## Examples

### Example 1

**Input:**

```
5
1 2 2 3 3
3
2 5 3
2 5 2
1 5 5
```

**Output:**

```
2
3
1
```

## Ideas

1. 如果将整个区间分成 m 段，每段计算出这个区间的最小值
2. 但是可能正好有那种在两个区间中间的询问
3. 如果在每个 l 处记录 dp[l][w] = h 表示在 l 处长度为 w 时的最优高度
4. 那么构造 w 个区间树，这个只有当 w 比较小的时候才可以
5. 如果 w 比较大？
6. 如果按照高度/宽度 (h, w) 组成的区域进行分割，貌似只会有 O(n) 个这样的区域
7. 这样的区间可以表示成 (l, r, h)
8. 那么剩下就比较简单了

## Editorial

二分答案：固定高度 h，判断能否在 [l, r] 内放下宽 w、高 h 的矩形。

数据结构：可持久化线段树，维护区间内**连续 1 的最大长度** (maxOnes)。叶节点只有 0 和 1。

需要维护的值：

- `len` — 区间长度
- `prefOnes` — 前缀连续 1 的长度
- `sufOnes` — 后缀连续 1 的长度
- `maxOnes` — 区间内连续 1 的最大长度

合并方式：

- `maxOnes = max(maxOnes(L), maxOnes(R), sufOnes(L) + prefOnes(R))`
- `prefOnes`：若 `len(L) == prefOnes(L)` 则为 `prefOnes(L) + prefOnes(R)`，否则为 `prefOnes(L)`
- `sufOnes`：若 `len(R) == sufOnes(R)` 则为 `sufOnes(L) + sufOnes(R)`，否则为 `sufOnes(R)`

建树：按高度从高到低，依次在对应位置置 1。每个高度对应一个版本。

查询：在版本 h 的树上查询 [l, r] 的 maxOnes，若 ≥ w 则可行。

复杂度：建树 O(n log n)，单次查询 O(log n)。
