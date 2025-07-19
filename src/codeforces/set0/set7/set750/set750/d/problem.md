# Firework Problem

## Problem Description

One tradition of welcoming the New Year is launching fireworks into the sky. Usually a launched firework flies vertically upward for some period of time, then explodes, splitting into several parts flying in different directions. Sometimes those parts also explode after some period of time, splitting into even more parts, and so on.

Limak, who lives in an infinite grid, has a single firework. The behaviour of the firework is described with a recursion depth $n$ and a duration for each level of recursion $t_1, t_2, ..., t_n$. 

Once Limak launches the firework in some cell, the firework starts moving upward. After covering $t_1$ cells (including the starting cell), it explodes and splits into two parts, each moving in the direction changed by 45 degrees (see the pictures below for clarification). So, one part moves in the top-left direction, while the other one moves in the top-right direction. 

Each part explodes again after covering $t_2$ cells, splitting into two parts moving in directions again changed by 45 degrees. The process continues till the $n$-th level of recursion, when all $2^{n-1}$ existing parts explode and disappear without creating new parts. After a few levels of recursion, it's possible that some parts will be at the same place and at the same time — it is allowed and such parts do not crash.

Before launching the firework, Limak must make sure that nobody stands in cells which will be visited at least once by the firework. Can you count the number of those cells?

## Input

The first line of the input contains a single integer $n$ ($1 \leq n \leq 30$) — the total depth of the recursion.

The second line contains $n$ integers $t_1, t_2, ..., t_n$ ($1 \leq t_i \leq 5$). On the $i$-th level each of $2^{i-1}$ parts will cover $t_i$ cells before exploding.

## Output

Print one integer, denoting the number of cells which will be visited at least once by any part of the firework.

## Examples

### Example 1

**Input:**
```
4
4 2 2 3
```

**Output:**
```
39
```

### Example 2

**Input:**
```
6
1 1 1 1 1 3
```

**Output:**
```
85
```

### Example 3

**Input:**
```
1
3
```

**Output:**
```
3
```

## Note

For the first sample, the drawings below show the situation after each level of recursion. Limak launched the firework from the bottom-most red cell. It covered $t_1 = 4$ cells (marked red), exploded and divided into two parts (their further movement is marked green). All explosions are marked with an 'X' character. 

On the last drawing, there are 4 red, 4 green, 8 orange and 23 pink cells. So, the total number of visited cells is $4 + 4 + 8 + 23 = 39$.

For the second sample, the drawings below show the situation after levels 4, 5 and 6. The middle drawing shows directions of all parts that will move in the next level.


### ideas
1. 如果不考虑重叠的部分， 答案很简单 = t[i] * (1 << i)
2. 怎么计算重叠区域呢？
3. 但是这个区域的大小，似乎比较小 5 * n * 5 * n?
4. 高度 = sum(t[i]), 宽带 = 2 * (sum[1:]) + 1
5. 所以这个范围是比较小的。
6. 只需要计算，有哪些格子被覆盖到了
7. 每个格子有5个状态，是否发生了爆炸，是否向上、向左、向下、向右发生了爆炸
8. 还不够，还必须知道爆炸的距离(*5) 
9. h * w * 5 * 5 = 还ok
10. 