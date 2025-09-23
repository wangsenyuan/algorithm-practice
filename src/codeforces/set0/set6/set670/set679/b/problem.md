# Limak and the Tower of Blocks

## Problem Description

Limak is a little polar bear. He plays by building towers from blocks. Every block is a cube with positive integer length of side. Limak has infinitely many blocks of each side length.

A block with side $a$ has volume $a^3$. A tower consisting of blocks with sides $a_1, a_2, \ldots, a_k$ has the total volume $a_1^3 + a_2^3 + \ldots + a_k^3$.

Limak is going to build a tower. First, he asks you to tell him a positive integer $X$ — the required total volume of the tower. Then, Limak adds new blocks greedily, one by one. Each time he adds the biggest block such that the total volume doesn't exceed $X$.

Limak asks you to choose $X$ not greater than $m$. Also, he wants to maximize the number of blocks in the tower at the end (however, he still behaves greedily). Secondarily, he wants to maximize $X$.

Can you help Limak? Find the maximum number of blocks his tower can have and the maximum $X \leq m$ that results in this number of blocks.

## Input

The only line of the input contains one integer $m$ ($1 \leq m \leq 10^{15}$), meaning that Limak wants you to choose $X$ between 1 and $m$, inclusive.

## Output

Print two integers — the maximum number of blocks in the tower and the maximum required total volume $X$, resulting in the maximum number of blocks.

## Examples

### Input
```
48
```

### Output
```
9 42
```

### Input
```
6
```

### Output
```
6 6
```

## Note

In the first sample test, there will be 9 blocks if you choose $X = 23$ or $X = 42$. Limak wants to maximize $X$ secondarily so you should choose 42.

In more detail, after choosing $X = 42$ the process of building a tower is:

1. Limak takes a block with side 3 because it's the biggest block with volume not greater than 42. The remaining volume is $42 - 27 = 15$.
2. The second added block has side 2, so the remaining volume is $15 - 8 = 7$.
3. Finally, Limak adds 7 blocks with side 1, one by one.

So, there are 9 blocks in the tower. The total volume is $3^3 + 2^3 + 7 \cdot 1^3 = 27 + 8 + 7 = 42$.



### ideas
1. 先2分数量？
2. 如果当前的x < 8, 那么贡献 = x
3. 然后如果在x = 8~27之间, x = 8, f(x) = 1, x = 9 = 1 + f(x - 8) = x - 8 + 1, ... f(16) = 2
4.  f(17) = 3, f(18) = 4... f(24) = 3, f(25) = 4, f(26) = 6, f(27) = 1
5.  也就是说不是x越大，这个f(x)越大, 不具备2分的情况
6.  f(x) = f(x - z * z * z) + 1
7.  对于m，找到最大的z, z * z * z <= m
8.  要么x < z3, 要么 x >= z3 如果 x >= z3 => f(x) = f(x - z3) + 1
9.  x - z3 是一个非常快的下降的数
10. 对于给定的x，计算它的f(x)是非常快的
11. 1e5, 
12. 放过来，假设我们想达到f(x) = y, 可以得到的最大的x是多少？
13. 这个y似乎是比较小的？
14. 假设是 $1^3, 2^3, 3^3, .....1e5^3$
15. 假设 x = 26 * 27
16. 125 = $5^3$
17. 124 / 27 = 4
18. 所以个数 7 + (26 / 8) + (63 / 27) + (124 / 64)
19. 7 + 3 + 2 + 1 + (215 / 125)  ...
20. 也就是手，对于5后面的数，它最多在区间(i, i+1)中间，增加1
21. 所以y ~~ 1e5左右
22. dp[y] 表示在y的情况下，能到达的最大的x
23. dp[y+1] = ?,假设x出在区间 (i, i+1), 那么 x 必须落在区间 [i+1, i+2]中间，
24. 