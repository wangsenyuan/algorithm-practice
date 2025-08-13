# Problem E: Road Reconstruction

## Problem Description

The great Shaass is the new king of the Drakht empire. The empire has $n$ cities which are connected by $n - 1$ bidirectional roads. Each road has a specific length and connects a pair of cities. There's a unique simple path connecting each pair of cities.

His majesty the great Shaass has decided to tear down one of the roads and build another road with the same length between some pair of cities. He should build such road that it's still possible to travel from each city to any other city. He might build the same road again.

You as his advisor should help him to find a way to make the described action. You should find the way that minimizes the total sum of pairwise distances between cities after the action. So calculate the minimum sum.

## Input

The first line of the input contains an integer $n$ denoting the number of cities in the empire, $(2 \leq n \leq 5000)$.

The next $n - 1$ lines each contains three integers $a_i$, $b_i$ and $w_i$ showing that two cities $a_i$ and $b_i$ are connected using a road of length $w_i$, $(1 \leq a_i, b_i \leq n, a_i \neq b_i, 1 \leq w_i \leq 10^6)$.

## Output

On the only line of the output print the minimum pairwise sum of distances between the cities.

## Note

Please do not use the `%lld` specificator to read or write 64-bit integers in C++. It is preferred to use the `cin`, `cout` streams or the `%I64d` specificator.

## Examples

### Example 1

**Input:**
```
3
1 2 2
1 3 4
```

**Output:**
```
12
```

### Example 2

**Input:**
```
6
1 2 1
2 3 1
3 4 1
4 5 1
5 6 1
```

**Output:**
```
29
```

### Example 3

**Input:**
```
6
1 3 1
2 3 1
3 4 100
4 5 2
4 6 1
```

**Output:**
```
825
```


### ideas
1. 假设将路径 (u, v) 断开，然后在其他地方进行了连接
2. 在两边内部的节点间的总距离是不变的
3. 然后假设把这个路径添加了另外两个节点(a, b)上
4. 那么变化, 假设有两个点(l, r)分别在两边
5. 原来的 dist(l, r) = dist(l, u) + (u, v) + dist(v, r)
6. 现在 dist(l, r) = dist(l, a) + (u, v) + dist(b, r)
7. 变化 = -dist(l...u) + dist(l, a) - dist(v, r) + dist(b, r)
8. a和b是独立的（相互可以单独选）
9. 也就是说，如果选中了一条边(u, v) （n-1）
10. 然后在两边选中两个新的点(a, b) (可以分别使用n) 去代替(u, v)
11. 第二步必须保证在(o(n))时间内完成，这个可以用re-root的方式计算 