# Problem D: Fishing Pond

## Problem Description

While Grisha was celebrating New Year with Ded Moroz, Misha gifted Sasha a small rectangular pond of size $n \times m$, divided into cells of size $1 \times 1$, inhabited by tiny evil fishes (no more than one fish per cell, otherwise they'll strife!).

The gift bundle also includes a square scoop of size $r \times r$, designed for fishing. If the lower-left corner of the scoop-net is located at cell $(x, y)$, all fishes inside the square $(x, y)...(x + r - 1, y + r - 1)$ get caught. Note that the scoop-net should lie completely inside the pond when used.

Unfortunately, Sasha is not that skilled in fishing and hence throws the scoop randomly. In order to not frustrate Sasha, Misha decided to release $k$ fishes into the empty pond in such a way that the expected value of the number of caught fishes is as high as possible.

## Task

Help Misha! In other words, put $k$ fishes in the pond into distinct cells in such a way that when the scoop-net is placed into a random position among $(n - r + 1) \cdot (m - r + 1)$ possible positions, the average number of caught fishes is as high as possible.

## Input

The only line contains four integers $n$, $m$, $r$, $k$:
- $1 \leq n, m \leq 10^5$
- $1 \leq r \leq \min(n, m)$
- $1 \leq k \leq \min(n \cdot m, 10^5)$

## Output

Print a single number — the maximum possible expected number of caught fishes.

**Note:** Your answer is considered correct if its absolute or relative error does not exceed $10^{-9}$. Namely, let your answer be $a$, and the jury's answer be $b$. Your answer is considered correct if $|a - b| \leq 10^{-9}$.

## Examples

### Example 1
**Input:**
```
3 3 2 3
```

**Output:**
```
2.0000000000
```

### Example 2
**Input:**
```
12 17 9 40
```

**Output:**
```
32.8333333333
```

## Note

In the first example you can put the fishes in cells $(2, 1)$, $(2, 2)$, $(2, 3)$. In this case, for any of four possible positions of the scoop-net (highlighted with light green), the number of fishes inside is equal to two, and so is the expected value.

## Key Points

- The pond is $n \times m$ in size
- The scoop-net is $r \times r$ in size
- Place $k$ fishes to maximize expected catch
- The scoop-net must be completely inside the pond
- Calculate expected value over all possible scoop-net positions

### ideas
1. 完全没有想法～
2. 把k条鱼，撒到 n * m 的格子中，使的一个r * r 的网，随机撒下去后，能够获得的鱼的期望值最高
3. 这个网的期望值  = r * r * 每个格子的平均期望值
4. 如果 k >= n * m => 1.0 （每个格子的期望值）
5. 好像不对。网的左上点，（n - r + 1) * (m - r + 1)个点
6. 在这些点的概率是不一样的
7. 假设有k = n*m-1, 那么可以猜测，右下角的点，不应该放鱼，因为这个点可以被捕到的概率 = 1 / (n1 * m1)
8. 但是它旁边的点（比如左边）= 2 / (n1 * m1), 有两个点下网点可以覆盖到这个点
9. 所以这个问题，换个形式，是把k条鱼，放到n*m个格子中，使的随机一个 r * r块的期望值最大
10. 那么这里，每个格子被覆盖的数量 = 能够覆盖它的，落网的左上角的数量 , (x, y) 
11. h = min(x + r - 1, n) - max(x - r + 1, 1) + 1
12. w = min(y + r - 1, m) - max(y - r + 1, 1) + 1
13. h * w
14. 这个越大的，越应该分配到鱼
15. sum(...k) / (n1 * m1)
16. 也就是说要怎么在n * m 非常大的情况下，找到最大的k个（h * w)的数
17. 可以用pq?
18. 最中间的数，肯定是最大的