# Problem Statement

You are given an array $a_1, a_2, \ldots, a_n$ of length $n$ and a positive integer $k$, but some parts of the array $a$ are missing. Your task is to fill the missing part so that the maximum subarray sum* of $a$ is exactly $k$, or report that no solution exists.

## Formal Description

You are given a binary string $s$ and a partially filled array $a$, where:

- If you remember the value of $a_i$, $s_i = 1$ to indicate that, and you are given the real value of $a_i$.
- If you don't remember the value of $a_i$, $s_i = 0$ to indicate that, and you are given $a_i = 0$.

All the values that you remember satisfy $|a_i| \leq 10^6$. However, you may use values up to $10^{18}$ to fill in the values that you do not remember. It can be proven that if a solution exists, a solution also exists satisfying $|a_i| \leq 10^{18}$.

\* The maximum subarray sum of an array $a$ of length $n$, i.e. $a_1, a_2, \ldots a_n$ is defined as $\max_{1 \leq i \leq j \leq n} S(i,j)$ where $S(i,j) = a_i + a_{i+1} + \ldots + a_j$.

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \leq t \leq 10^4$). The description of the test cases follows.

For each test case:
- The first line contains two numbers $n, k$ ($1 \leq n \leq 2 \cdot 10^5, 1 \leq k \leq 10^{12}$).
- The second line contains a binary (01) string $s$ of length $n$.
- The third line contains $n$ numbers $a_1, a_2, \ldots, a_n$ ($|a_i| \leq 10^6$). If $s_i = 0$, then it's guaranteed that $a_i = 0$.

It is guaranteed that the sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

## Output

For each test case:
1. First output `Yes` if a solution exists or `No` if no solution exists. You may print each character in either case, for example `YES` and `yEs` will also be accepted.
2. If there's at least one solution, print $n$ numbers $a_1, a_2, \ldots, a_n$ on the second line. $|a_i| \leq 10^{18}$ must hold.

## ideas
1. 因为可以使用的范围很大，所以只需要使用一个数就可以了