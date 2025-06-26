# Problem C: Maximum Beauty of Nice Arrays

## Problem Description

The distance of a value $x$ in an array $c$, denoted as $d_x(c)$, is defined as the largest gap between any two occurrences of $x$ in $c$.

Formally, $d_x(c) = \max(j - i)$ over all pairs $i < j$ where $c_i = c_j = x$. If $x$ appears only once or not at all in $c$, then $d_x(c) = 0$.

The beauty of an array is the sum of the distances of each distinct value in the array. Formally, the beauty of an array $c$ is equal to $\sum_{1 \leq x \leq n} d_x(c)$.

Given an array $a$ of length $n$, an array $b$ is **nice** if it also has length $n$ and its elements satisfy $1 \leq b_i \leq a_i$ for all $1 \leq i \leq n$. Your task is to find the maximum possible beauty of any nice array.

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \leq t \leq 10^4$). The description of the test cases follows.

The first line of each test case contains a single integer $n$ ($1 \leq n \leq 2 \cdot 10^5$) — the length of array $a$.

The second line of each test case contains $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq n$) — the elements of array $a$.

It is guaranteed that the sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

## Output

For each test case, output a single integer representing the maximum possible beauty among all nice arrays.


## ideas
1. 搞反了，是要找到最大值，而不是最小值