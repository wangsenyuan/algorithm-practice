Recently, the bear started studying data structures and faced the following problem.

You are given a sequence of integers $x_1, x_2, \ldots, x_n$ of length $n$ and $m$ queries, each of them is characterized by two integers $l_i, r_i$. Let's introduce $f(p)$ to represent the number of such indexes $k$, that $x_k$ is divisible by $p$. The answer to the query $l_i, r_i$ is the sum: $\sum_{p \in S(l_i, r_i)} f(p)$, where $S(l_i, r_i)$ is a set of prime numbers from segment $[l_i, r_i]$ (both borders are included in the segment).

Help the bear cope with the problem.

---

## Input

- The first line contains integer $n$ ($1 \leq n \leq 10^6$).
- The second line contains $n$ integers $x_1, x_2, \ldots, x_n$ ($2 \leq x_i \leq 10^7$). The numbers are not necessarily distinct.
- The third line contains integer $m$ ($1 \leq m \leq 50000$).
- Each of the following $m$ lines contains a pair of space-separated integers, $l_i$ and $r_i$ ($2 \leq l_i \leq r_i \leq 2 \cdot 10^9$) — the numbers that characterize the current query.

## Output

Print $m$ integers — the answers to the queries in the order the queries appear in the input.

## ideas
1. 每个数，num，它的质因数，不会超过log(num)个
2. 然后用它的质因数p，去更新查询的空间（某个查询l, r, l <= p and p <= r 加1）
3. 怎么快速的去处理呢？
4. 假设有一个num，有两个p1, p2; 且每个区间，包含了只包含了p1或者p2, 那么它+1，
5. 如果包含两个那么+2
6. 那么是不是就是差分数组就可以了？