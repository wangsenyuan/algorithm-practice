# Problem: The Little Elephant and Array

The Little Elephant loves playing with arrays. He has array $a$, consisting of $n$ positive integers, indexed from $1$ to $n$. Let's denote the number with index $i$ as $a_i$.

Additionally, the Little Elephant has $m$ queries to the array, each query is characterised by a pair of integers $l_j$ and $r_j$ ($1 \leq l_j \leq r_j \leq n$). For each query $l_j, r_j$ the Little Elephant has to count, how many numbers $x$ exist, such that number $x$ occurs exactly $x$ times among numbers $a_{l_j}, a_{l_j+1}, ..., a_{r_j}$.

Help the Little Elephant to count the answers to all queries.

---

## Input

The first line contains two space-separated integers $n$ and $m$ ($1 \leq n, m \leq 10^5$) — the size of array $a$ and the number of queries to it.

The next line contains $n$ space-separated positive integers $a_1, a_2, ..., a_n$ ($1 \leq a_i \leq 10^9$).

Next $m$ lines contain descriptions of queries, one per line. The $j$-th of these lines contains the description of the $j$-th query as two space-separated integers $l_j$ and $r_j$ ($1 \leq l_j \leq r_j \leq n$).

---

## Output

In $m$ lines print $m$ integers — the answers to the queries. The $j$-th line should contain the answer to the $j$-th query.

## ideas
1. 显然x不能超过m
2. 对于区间l...r，可以找到这个区间内，可以计算每个数出现的次数（用莫队算法）
3. 就用莫队算法就可以了