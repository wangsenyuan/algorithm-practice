## XOR and Favorite Number

Bob has a favorite number $k$ and $a_i$ of length $n$. Now he asks you to answer $m$ queries. Each query is given by a pair $l_i$ and $r_i$ and asks you to count the number of pairs of integers $i$ and $j$, such that $l \leq i \leq j \leq r$ and the xor of the numbers $a_i, a_{i + 1}, ..., a_j$ is equal to $k$.

### Input

The first line of the input contains integers $n$, $m$ and $k$ ($1 \leq n, m \leq 100000$, $0 \leq k \leq 1000000$) — the length of the array, the number of queries and Bob's favorite number respectively.

The second line contains $n$ integers $a_i$ ($0 \leq a_i \leq 1000000$) — Bob's array.

Then $m$ lines follow. The $i$-th line contains integers $l_i$ and $r_i$ ($1 \leq l_i \leq r_i \leq n$) — the parameters of the $i$-th query.

### Output

Print $m$ lines, answer the queries in the order they appear in the input.

### Examples

#### Example 1

**Input:**

```
6 2 3
1 2 1 1 0 3
1 6
3 5
```

**Output:**

```
7
0
```

#### Example 2

**Input:**

```
5 3 1
1 1 1 1 1
1 5
2 4
1 3
```

**Output:**

```
9
4
4
```

### Note

In the first sample the suitable pairs of $i$ and $j$ for the first query are: $(1, 2)$, $(1, 4)$, $(1, 5)$, $(2, 3)$, $(3, 6)$, $(5, 6)$, $(6, 6)$. Not a single of these pairs is suitable for the second query.

In the second sample xor equals $1$ for all subarrays of an odd length.


### ideas
1. 先不考虑l...r，考虑整个序列中，如何找到 xor = k 的计数？
2. xor[i...j] = xor[j] ^ xor[i-1] 把区间变成前缀
3. 对于一个给定的k，从左往右扫过去的时候, xor[j...?] = k = xor[?] = xor[j] ^ k 的freq
4. 使用MO算法，将所有的查询排序，这样子就有sqrt(N)的复杂性，然后对于每个查询，再使用log(n) 去找到l的位置
5. Q * sqrt(N) * log(N) ~ 100 000 * 1000 * 20;
6. 似乎不大行呐～
7. 而且还有map的消耗
8. 好像不用 MO，把所有的查询放置在R端，然后在R端的时候，根据每个查询去search就好了
9. 错了～
10. 假设处理到i, 那么对于xor[j] = xor[i] ^ k 的位置j, 都可以+1， 那么这样查询的时候，只要找l后面的sum就可以了
11. 但是问题出在对于i来说，这样的j貌似有很多哎～
12. 好像，记录最近的那个，并把前面的j的值，更新到新的位置就可以了～
13. 