# Moving Boxes

## Problem Description

There is an infinite line consisting of cells. There are $n$ boxes in some cells of this line. The $i$-th box stands in the cell $a_i$ and has weight $w_i$. All $a_i$ are distinct, moreover, $a_{i-1} < a_i$ holds for all valid $i$.

You would like to put together some boxes. Putting together boxes with indices in the segment $[l,r]$ means that you will move some of them in such a way that their positions will form some segment $[x,x+(r-l)]$.

In one step you can move any box to a neighboring cell if it isn't occupied by another box (i.e. you can choose $i$ and change $a_i$ by $1$, all positions should remain distinct). You spend $w_i$ units of energy moving the box $i$ by one cell. You can move any box any number of times, in arbitrary order.

## Queries

Sometimes weights of some boxes change, so you have queries of two types:

1. `id nw` — weight $w_{id}$ of the box $id$ becomes $nw$.
2. `l r` — you should compute the minimum total energy needed to put together boxes with indices in $[l,r]$. 

Since the answer can be rather big, print the remainder it gives when divided by $1000000007 = 10^9+7$. Note that the boxes are not moved during the query, you only should compute the answer.

**Important**: You should minimize the answer, not its remainder modulo $10^9+7$. So if you have two possible answers $2⋅10^9+13$ and $2⋅10^9+14$, you should choose the first one and print $10^9+6$, even though the remainder of the second answer is $0$.

## Input

- The first line contains two integers $n$ and $q$ $(1 \leq n,q \leq 2⋅10^5)$ — the number of boxes and the number of queries.
- The second line contains $n$ integers $a_1,a_2,…a_n$ $(1 \leq a_i \leq 10^9)$ — the positions of the boxes. All $a_i$ are distinct, $a_{i-1} < a_i$ holds for all valid $i$.
- The third line contains $n$ integers $w_1,w_2,…w_n$ $(1 \leq w_i \leq 10^9)$ — the initial weights of the boxes.
- Next $q$ lines describe queries, one query per line.

Each query is described in a single line, containing two integers $x$ and $y$:
- If $x < 0$, then this query is of the first type, where $id = -x$, $nw = y$ $(1 \leq id \leq n, 1 \leq nw \leq 10^9)$
- If $x > 0$, then the query is of the second type, where $l = x$ and $r = y$ $(1 \leq l_j \leq r_j \leq n)$
- $x$ cannot be equal to $0$

## Output

For each query of the second type print the answer on a separate line. Since answer can be large, print the remainder it gives when divided by $1000000007 = 10^9+7$.

## Example

### Input
```
5 8
1 2 6 7 10
1 1 1 1 2
1 1
1 5
1 3
3 5
-3 5
-1 10
1 4
2 5
```

### Output
```
0
10
3
4
18
7
```

## Note

Let's go through queries of the example:

1. `1 1` — there is only one box so we don't need to move anything.
2. `1 5` — we can move boxes to segment $[4,8]$: $1⋅|1-4|+1⋅|2-5|+1⋅|6-6|+1⋅|7-7|+2⋅|10-8|=10$.
3. `1 3` — we can move boxes to segment $[1,3]$.
4. `3 5` — we can move boxes to segment $[7,9]$.
5. `-3 5` — $w_3$ is changed from $1$ to $5$.
6. `-1 10` — $w_1$ is changed from $1$ to $10$. The weights are now equal to $w=[10,1,5,1,2]$.
7. `1 4` — we can move boxes to segment $[1,4]$.
8. `2 5` — we can move boxes to segment $[5,8]$.


### ideas
1. 先考虑整个区间上要怎么计算,可以确定的一点是，必然有一个位置是锚点，假设这个位置是a[i]
2. 那么有 (a[i] - (i-1) - a[1]) * w[1] + (a[i] - (i - 2) - a[2]) * w[2] + ... + (a[i] - 1 - a[i-1]) * w[i-1] + (a[i] - a[i]) * w[i]
3.  + (a[i+1] - (a[i] + 1)) * w[i+1] + ... + (a[i + j] - (a[i] + j)) * w[i+j] + ..
4.  s1 = a[i] * (w[1] + w[2] + ... * w[i]) - (w[i+1] + w[i+2] + ... + w[n]) * a[i]
5.  s2 = -(i-1) * w[1] - (i-2) * w[2] .. - 1 * w[i-1] - 0 * w[i] - 1 * w[i+1] - 2 * w[i+2] ... - (n - i) * w[i]
6.  s3 = -a[1] * w[1] - a[2] * w[2] - .... - a[i] * w[i] + a[i+1] * w[i+1] + a[i+2] * w[i+2] + ... w[n] * a[n]
7.  可以看到 s1基本只和a[i]有关系（需要知道前后区间w的sum)
8.  s2和i有关系，但乘法有点麻烦
9.  s3似乎也容易搞
10. 这个应该要用到凸优化方案
11. 也就是随着i变化的时候，如何快速的计算出新的最优解
12. 每个节点维护，这段内最优的线段集，然后合并线段的时候，利用凸优化快速的计算出新的集