# Problem Statement

## Local Minima and Maxima Definition

An element $b_i$ ($1 \leq i \leq m$) in an array $b_1, b_2, \ldots, b_m$ is a **local minimum** if at least one of the following holds:

1. $2 \leq i \leq m-1$ and $b_i < b_{i-1}$ and $b_i < b_{i+1}$, or
2. $i = 1$ and $b_1 < b_2$, or
3. $i = m$ and $b_m < b_{m-1}$

Similarly, an element $b_i$ ($1 \leq i \leq m$) in an array $b_1, b_2, \ldots, b_m$ is a **local maximum** if at least one of the following holds:

1. $2 \leq i \leq m-1$ and $b_i > b_{i-1}$ and $b_i > b_{i+1}$, or
2. $i = 1$ and $b_1 > b_2$, or
3. $i = m$ and $b_m > b_{m-1}$

> Note: Local minima and maxima are not defined for arrays with only one element.

## Problem Description

There is a hidden permutation* $p$ of length $n$. The following two operations are applied to permutation $p$ alternately, starting from operation 1, until there is only one element left in $p$:

1. **Operation 1** — remove all elements of $p$ which are not local minima.
2. **Operation 2** — remove all elements of $p$ which are not local maxima.

More specifically, operation 1 is applied during every odd iteration, and operation 2 is applied during every even iteration, until there is only one element left in $p$.

For each index $i$ ($1 \leq i \leq n$), let $a_i$ be the iteration number that element $p_i$ is removed, or $-1$ if it was never removed.

It can be proven that there will be only one element left in $p$ after at most $\lceil \log_2 n \rceil$ iterations (in other words, $a_i \leq \lceil \log_2 n \rceil$).

You are given the array $a_1, a_2, \ldots, a_n$. Your task is to construct any permutation $p$ of $n$ elements that satisfies array $a$.

> *A permutation of length $n$ is an array consisting of $n$ distinct integers from $1$ to $n$ in arbitrary order. For example, $[2,3,1,5,4]$ is a permutation, but $[1,2,2]$ is not a permutation (2 appears twice in the array), and $[1,3,4]$ is also not a permutation ($n=3$ but there is 4 in the array).

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \leq t \leq 10^4$). The description of the test cases follows.

The first line of each test case contains a single integer $n$ ($2 \leq n \leq 2 \cdot 10^5$) — the number of elements in permutation $p$.

The second line of each test case contains $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq \lceil \log_2 n \rceil$ or $a_i = -1$) — the iteration number that element $p_i$ is removed.

It is guaranteed that:
- The sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.
- There exists at least one permutation $p$ that satisfies array $a$.

## Output

For each test case, output $n$ integers representing the elements of the permutation satisfying array $a$.

If there are multiple solutions, you may output any of them.

## ideas
1. not a clue~
2. 在第一轮中，所有不是最小值的，都被删除掉了,假设i被删除掉，但是i+1没有(a[i] < a[i+1])
3. 那么必然有p[i] > p[i+1] (i+1是最小值)
4. 应该看哪些被保留下来了，假设保留了i, 那么就有 p[i-1] > p[i] and p[i] < p[i+1]
5. 那么这样子可以组成一组关系（每个位置，有a[i]个边）
6. 形成一个dag后，那么从deg = 0 的地方开始处理，
7.  