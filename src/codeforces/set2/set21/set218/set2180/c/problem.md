Ostad thinks that the usual way of factoring numbers is too mathematical, so he invented a new notion called XOR-factorization, which is more computer-science-like. For a given integer $n$, a sequence of integers $a_1, a_2, \ldots, a_k$ with $0 \le a_i \le n$ for all $i$ is called a **XOR-factorization** of $n$ if and only if

\[
a_1 \oplus a_2 \oplus \cdots \oplus a_k = n,
\]

where $\oplus$ denotes the bitwise XOR operation.

You are given integers $n$ and $k$. Find a XOR-factorization $a_1, a_2, \ldots, a_k$ of $n$ that maximizes the sum $a_1 + a_2 + \cdots + a_k$.

It can be proven that under the problem conditions, a XOR-factorization always exists.

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 10^4$). The description of the test cases follows.

Each of the next $t$ lines contains two integers $n$ and $k$ ($1 \le n \le 10^9$, $1 \le k \le 10^5$).

It is guaranteed that the sum of $k$ over all test cases does not exceed $10^5$.

## Output

For each test case, output $k$ integers $a_1, a_2, \ldots, a_k$ such that $0 \le a_i \le n$.

We can show that an answer always exists. If there are multiple valid answers, you may print any of them in any order.

## Example

### Input

```
4
5 4
4 3
8 2
1 1
```

### Output

```
1 4 5 5
4 4 4
0 8
1
```

## Note

In the first test case, we can factor $5$ as $1 \oplus 4 \oplus 5 \oplus 5$ with a sum of $15$, and it can be shown that no other XOR-factorization has a higher sum.

In the second test case, we can factor $4$ as $4 \oplus 4 \oplus 4$ with a sum of $12$, which is trivially the maximum possible.


### ideas
1. 如果 k is odd => [n, n, ... n] 
2. k is even， at least [0, n, n, ..., n] = (k - 1) * n
3. 有没有更好的结果？
4. 假设n[d] = 0, d[d+1] = 1, 且d是最高的
5. 假设到目前为止，分配了m个数， a[?] 和n的高位一致， 其中m是奇数
6. 现在处理d位，（需要偶数个a[?][d] = 1), k - m 是奇数（因为k是偶数）
7. 那么只能是 k - m - 1 个 1 << d, 但是，这样子，所有剩余的数 < n 成立了
8. 比如 n(25) = 11001 , k = 4, (k - 1) * n = 75
9. [10111, 10111, 10111]
10. 