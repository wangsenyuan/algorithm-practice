# Problem F - Choose a Square

## Problem Description

Petya recently found a game "Choose a Square". In this game, there are $n$ points numbered from $1$ to $n$ on an infinite field. The $i$-th point has coordinates $(x_i, y_i)$ and cost $c_i$.

You have to choose a square such that:
- Its sides are parallel to coordinate axes
- The lower left and upper right corners belong to the line $y = x$
- All corners have integer coordinates

The score you get is the sum of costs of the points covered by the selected square minus the length of the side of the square. Note that the length of the side can be zero.

Petya asks you to calculate the maximum possible score in the game that can be achieved by placing exactly one square.

## Input

The first line of the input contains one integer $n$ $(1 \leq n \leq 5 \cdot 10^5)$ — the number of points on the field.

Each of the following $n$ lines contains three integers $x_i, y_i, c_i$ $(0 \leq x_i, y_i \leq 10^9, -10^6 \leq c_i \leq 10^6)$ — coordinates of the $i$-th point and its cost, respectively.

## Output

In the first line print the maximum score Petya can achieve.

In the second line print four integers $x_1, y_1, x_2, y_2$ $(0 \leq x_1, y_1, x_2, y_2 \leq 2 \cdot 10^9, x_1 = y_1, x_2 = y_2, x_1 \leq x_2)$ separated by spaces — the coordinates of the lower left and upper right corners of the square which Petya has to select in order to achieve the maximum score.

## Examples

### Example 1
**Input:**
```
6
0 0 2
1 0 -5
1 1 3
2 3 4
1 4 -4
3 1 -1
```

**Output:**
```
4
1 1 3 3
```

### Example 2
**Input:**
```
5
3 3 0
3 3 -3
0 2 -1
3 1 3
0 0 -2
```

**Output:**
```
0
1 1 1 1
```


### ideas
1. F(r, r) - F(l, r) - F(r, l) + F(l, l) - r + l
  

### solution
设正方形左下角为 (a,a)，右上角为 (b,b)。
如果正方形包含点 (x,y)，那么有
a <= x <= b
a <= y <= b
即
a <= min(x,y) <= max(x,y) <= b

定义 x' = min(x,y)，y' = max(x,y)。正方形要包含 (x',y')。
直观的理解方式是，如果 x > y，那么交换 x 和 y。相当于把 y=x 下方的点翻转到 y=x 上方，这不影响答案。
于是判断条件简化成 x >= a 且 y <= b。

分数为 sum(c) - (b - a) = (a + sum(c)) - b。
枚举正方形包含的纵坐标最大的点 y，那么 b 越小越好，所以 b = y。
现在问题变成最大化 a + sum(c)，其中 sum(c) 中的点满足 x >= a。

核心思路：枚举右，维护左。
枚举 b=y，维护每个 a 对应的 a + sum(c)，然后计算 [0,y] 中的这些 a + sum(c) 的最大值。

用线段树维护，线段树中的叶子 a 保存满足 x >= a 且 y <= b 的点的 a + sum(c)。初始值为 a。
当我们遍历到 (x,y) 时，对于包含 (x,y) 的正方形，满足 a <= x，所以把线段树 [0,x] 中的数都增加 c。这需要 Lazy 线段树。
对于所有 y 相同的点，当前正方形都要包含（当前在枚举 b=y 的正方形），所以要先对这些点执行一遍线段树的区间加（[0,x] 中的数都增加 c），然后再求线段树 [0,y] 中的最大值。
为什么不能一个一个处理？注意本题 c 可能是负数，如果有两个负数 c 需要做区间更新，执行第一个区间更新就立刻算区间最大值，会导致答案偏大。

为方便使用线段树，需要把 x 和 y 离散化。
由于需要输出具体方案，线段树还需要维护区间最大值的下标。
