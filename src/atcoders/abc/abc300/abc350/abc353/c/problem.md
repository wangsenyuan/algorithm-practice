# C - Sigma Problem

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 300 points

## Problem Statement
For positive integers $x$ and $y$, define $f(x, y)$ as the remainder of $(x + y)$ divided by $10^8$.

You are given a sequence of positive integers $A = (A_1, \ldots, A_N)$ of length $N$. Find the value of the following expression:

$$
\sum_{i=1}^{N-1} \sum_{j=i+1}^{N} f(A_i, A_j).
$$

## Constraints
- $2 \leq N \leq 3 \times 10^5$
- $1 \leq A_i < 10^8$
- All input values are integers.

## ideas
1. 还不是简单的 n * sum(A) % 1e8
2. 将A排序，对结果是没有影响的
3. 然后对于a[i]来说，如果有w个a[j] + a[i] < 1e8, 那么a[i]的贡献 = w (*a[i]) 
4. 如果有w个数a[j], a[i] + a[j] >= 1e8 但是小于 2e8
5. w * a[i] + sum(a[j]) - w * 1e8
6. 似乎也可以搞
7. 但是如果a[i], a[i+1]远远超过w, 那么就有很多次空转
8. 不会，因为a[i] < 1e8， 所以只有两种情况 a[i] + a[j] < 1e8, or > 1e8