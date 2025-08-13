# Problem C: Coin Combinations

## Problem Description

In the Isle of Guernsey there are $n$ different types of coins. For each $i$ ($1 \leq i \leq n$), coin of type $i$ is worth $a_i$ cents. It is possible that $a_i = a_j$ for some $i$ and $j$ ($i \neq j$).

Bessie has some set of these coins totaling $t$ cents. She tells Jessie $q$ pairs of integers. For each $i$ ($1 \leq i \leq q$), the pair $b_i, c_i$ tells Jessie that Bessie has a strictly greater number of coins of type $b_i$ than coins of type $c_i$. It is known that all $b_i$ are distinct and all $c_i$ are distinct.

Help Jessie find the number of possible combinations of coins Bessie could have. Two combinations are considered different if there is some $i$ ($1 \leq i \leq n$), such that the number of coins Bessie has of type $i$ is different in the two combinations. Since the answer can be very large, output it modulo $1000000007$ ($10^9 + 7$).

If there are no possible combinations of coins totaling $t$ cents that satisfy Bessie's conditions, output $0$.

## Input

- The first line contains three space-separated integers: $n$, $q$ and $t$ ($1 \leq n \leq 300$; $0 \leq q \leq n$; $1 \leq t \leq 10^5$)
- The second line contains $n$ space-separated integers: $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq 10^5$)
- The next $q$ lines each contain two distinct space-separated integers: $b_i$ and $c_i$ ($1 \leq b_i, c_i \leq n$; $b_i \neq c_i$)

**Note:** It's guaranteed that all $b_i$ are distinct and all $c_i$ are distinct.

## Output

A single integer, the number of valid coin combinations that Bessie could have, modulo $1000000007$ ($10^9 + 7$).

## Examples

### Example 1

**Input:**
```
4 2 17
3 1 2 5
4 2
3 4
```

**Output:**
```
3
```

### Example 2

**Input:**
```
3 2 6
3 1 1
1 2
2 3
```

**Output:**
```
0
```

### Example 3

**Input:**
```
3 2 10
1 2 3
1 2
2 1
```

**Output:**
```
0
```

## Note

For the first sample, the following 3 combinations give a total of 17 cents and satisfy the given conditions:
- $\{0$ of type 1, $1$ of type 2, $3$ of type 3, $2$ of type 4$\}$
- $\{0, 0, 6, 1\}$
- $\{2, 0, 3, 1\}$

No other combinations exist. Note that even though 4 occurs in both $b_i$ and $c_i$, the problem conditions are still satisfied because all $b_i$ are distinct and all $c_i$ are distinct.

## ideas
1. 给定的关系，表明了一个顺序，也就是说， 加入要使用类似u的硬币，那么比它少的必须先处理掉，且i的数量，要更多
2. 所有的bi不相同，所有的c不相同，=> 所有的关系最多组成一条直线，不会是个树形（图形）的
3. 就是说可能存在 a < b < c  < d...
4. 那么可以每条线单独处理吗？
5. 比如处理一条线可以得到一个dp[x] (使用这条线得到x cents的数量)
6. 那么就是用 x1 + x2 + ... = t 时，这些结果的乘积
7. 但是这里还有两个难点。a. 如何在给定的直线上，计算fp[x] (这条线获得x cents)的计数？b. 如何快速合并dp?
8. 在一条线上时，后面的要比前面的多，所以需要知道当前使用了多少个
9. 但是使用的数量的上限 t/a[?] 这个还是太大了
10. 如果所有的a[i] = 1, 那么这个迭代就是 t * t * n了
11. 这个貌似可以用后缀和来优化，就是如果i使用x个, 那么整个后缀也必须使用x个
12. 那么可以整个后缀考虑, fp[suf + w] += fp[w]就可以了
13. 那么这样子，就把个数的迭代给去掉了。那么可以在 n * t时间内处理掉
14.  现在怎么合并 fp[j] * dp[i] => dp[i+j] 呢？t * t 肯定不行～
15. 假设要计算dp[x] 那么 dp[x] = dp[y] * fp[x - y] 
16. 假设计算出来了 dp[x-1], 能计算出dp[x] 吗？
17. 怎么感觉要用到FTT？