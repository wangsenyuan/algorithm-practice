# Problem D

You are given a sequence of $n$ integers $a_1, a_2, \ldots, a_n$.

You have to construct two sequences of integers $b$ and $c$ with length $n$ that satisfy:

- For every $i$ $(1 \leq i \leq n)$: $b_i + c_i = a_i$
- $b$ is non-decreasing, which means that for every $1 < i \leq n$, $b_i \geq b_{i-1}$ must hold
- $c$ is non-increasing, which means that for every $1 < i \leq n$, $c_i \leq c_{i-1}$ must hold

You have to minimize $\max(b_i, c_i)$. In other words, you have to minimize the maximum number in sequences $b$ and $c$.

Also there will be $q$ changes, the $i$-th change is described by three integers $l, r, x$. You should add $x$ to $a_l, a_{l+1}, \ldots, a_r$.

You have to find the minimum possible value of $\max(b_i, c_i)$ for the initial sequence and for sequence after each change.

## Input

The first line contains an integer $n$ $(1 \leq n \leq 10^5)$.

The second line contains $n$ integers $a_1, a_2, \ldots, a_n$ $(1 \leq i \leq n, -10^9 \leq a_i \leq 10^9)$.

The third line contains an integer $q$ $(1 \leq q \leq 10^5)$.

Each of the next $q$ lines contains three integers $l, r, x$ $(1 \leq l \leq r \leq n, -10^9 \leq x \leq 10^9)$, describing the next change.

## Output

Print $q + 1$ lines.

On the $i$-th $(1 \leq i \leq q + 1)$ line, print the answer to the problem for the sequence after $i - 1$ changes.

## ideas
1. 一头雾水
2. 在没有变化前， b和c是否是固定的。如果n=1，都不是固定的，b1 + c1 = a1, 为了让max(b1, c1) 最小，b1和c1尽量相等，如果a1是奇数，那么就有两种可能性
3. 虽然b1, c1不确定，但是max(b1, c1)是确定的
4. b是升序的，c是降序的，那么先考虑如何保证这个性质，
5. 如果c[1]非常大，那么b[1]就足够小，所以应该可以证明肯定存在这样的b和c满足条件
6. b + c = a
7. b = a - c => max(a - c, c) 最小
8. a[1] - c[1] <= a[2] - c[2] => a[1] - a[2] <= c[1] - c[2]
9. 也就是说，在给定c[1]的时候， c[2] <= c[1] - (a[1] - a[2]) （c[2]只有上限，且要求c非递增，所以同样证明了肯定有这样的序列）
10. c[2] = c[1] - max(a[1] - a[2], 0)
11. 也就是在确定c[1]的情况下， c[i]可以利用这个公式推导出来
12. 当a[1] <= a[2]的时候， c[2] = c[1], 当a[1] > a[2]的时候， c[2] = c[1] - (a[1] - a[2])
13. 所以，这样子c[i] = c[1] - sum(max(a[i-1] - a[i], 0))
14. b[i] = a[i] - c[i]
15. 如何计算其中的 最小化的 max(b[i], c[i]) 呢，
16. 如果a[i] = a[i+1], c[i] = c[i+1], 那么 b[i] = b[i+1]
17. 如果a[i] < a[i+1] => c[i] = c[i+1] => b[i] < b[i+1]
18. 如果a[i] > a[i+1] => c[i] > c[i+1]， c[i+1] = c[i] - (a[i] - a[i+1]) => b[i+1] = a[i+1] - c[i] + a[i] - a[i+1] = a[i] - c[i] = b[i]
19. 也就是说，当a[i] <= a[i+1]时， c[i] = c[i+1], 当a[i] >= a[i+1]时， b[i] = b[i+1]
20. 所以，可以预期的是前半段c[i] >= b[i], 后半段 c[i] <= b[i]
21. 所以，最大值要么发生在头，要么发生在尾部
22. 假设c[1]取某个值 => c[n] 计算出来 => b[n] 
23. 也就是要计算c[1]和b[n]的最小值
24. c[n] = c[1] - delta
25. b[n] = a[n] - c[n] = a[n] - c[1] + delta
26. c[1] + b[n] = a[n] + delta => got
27. 现在的问题，就是如何计算delta
28. delta = sum of (max(a[i] - a[i+1], 0))
29. 所以要计算 a[i] > a[i+1]的差值
30. range update + lazy