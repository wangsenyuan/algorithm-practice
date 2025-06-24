# G - Merchant Takahashi

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 550 points

## Problem Statement
The Kingdom of AtCoder has $N$ towns: towns $1, 2, \ldots, N$. To move from town $i$ to town $j$, you must pay a toll of $C \times |i - j|$ yen.

Takahashi, a merchant, is considering participating in zero or more of $M$ upcoming markets.

The $i$-th market ($1 \leq i \leq M$) is described by the pair of integers $(T_i, P_i)$, where the market is held in town $T_i$ and he will earn $P_i$ yen if he participates.

For all $1 \leq i < M$, the $i$-th market ends before the $(i+1)$-th market begins. The time it takes for him to move is negligible.

He starts with $10^{10^{100}}$ yen and is initially in town $1$. Determine the maximum profit he can make by optimally choosing which markets to participate in and how to move.

Formally, let $10^{10^{100}} + X$ be his final amount of money if he maximizes the amount of money he has after the $M$ markets. Find $X$.

## Constraints
- $1 \leq N \leq 2 \times 10^5$
- $1 \leq C \leq 10^9$
- $1 \leq M \leq 2 \times 10^5$
- $1 \leq T_i \leq N$ ($1 \leq i \leq M$)
- $1 \leq P_i \leq 10^{13}$ ($1 \leq i \leq M$)
- All input values are integers.


## ideas
1. 因为移动的时间可以忽略，所以只需要关心，Taka在前i个market活动中，目前处在什么位置？
2. dp[i] 表示处在位置i的时候，他的最大收益
3. dp[j] = P[j] - $C \times |t[j] - t[i]| $ + dp[i]
4. 要把绝对值去掉，就要区分是在T[j]的前面，还是后面
5. 假设T[j] = u, T[i] = v, 且 u >= v
6. dp[j] = P[j] - C * (u - v) + dp[i] = P[j] - C * u + (C * v + dp[i])
7. 假设 u < v
8. dp[j] = P[j] - C * (v - u) + dp[i] = P[j] + C * u - C * v + dp[i] = P[j] + C * u + (dp[i] - C * v)
9. 所以需要两棵树