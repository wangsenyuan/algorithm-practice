# Problem D: Army Strength

## Problem Statement

Winter is here at the North and the White Walkers are close. John Snow has an army consisting of n soldiers. While the rest of the world is fighting for the Iron Throne, he is going to get ready for the attack of the White Walkers.

He has created a method to know how strong his army is. Let the i-th soldier's strength be $a_i$. For some k he calls $i_1, i_2, ..., i_k$ a clan if $i_1 < i_2 < i_3 < ... < i_k$ and $\gcd(a_{i_1}, a_{i_2}, ..., a_{i_k}) > 1$. He calls the strength of that clan $k \cdot \gcd(a_{i_1}, a_{i_2}, ..., a_{i_k})$. Then he defines the strength of his army by the sum of strengths of all possible clans.

Your task is to find the strength of his army. As the number may be very large, you have to print it modulo $10^9 + 7$.

**Greatest common divisor (gcd)** of a sequence of integers is the maximum possible integer so that each element of the sequence is divisible by it.

## Input

The first line contains integer $n$ $(1 \leq n \leq 200000)$ — the size of the army.

The second line contains $n$ integers $a_1, a_2, ..., a_n$ $(1 \leq a_i \leq 1000000)$ — denoting the strengths of his soldiers.

## Output

Print one integer — the strength of John Snow's army modulo $10^9 + 7$.

## Examples

### Example 1
**Input:**
```
3
3 3 1
```

**Output:**
```
12
```

### Example 2
**Input:**
```
4
2 3 4 6
```

**Output:**
```
39
```

## Note

In the first sample the clans are $\{1\}, \{2\}, \{1, 2\}$ so the answer will be $1 \cdot 3 + 1 \cdot 3 + 2 \cdot 3 = 12$



### ideas
1. 考虑有f[2], f[4] 分别表示能整除2的数量和能整除4的数量
2. f[4] = m, 4的贡献  = 4 * (1 * C(m, 1) + 2 * C(m, 2) + ... + i * C(m, i), ...)
3. f[2] = n, 2的贡献 = 2 * (1 * C(n, 1) + 2 * C(n, 2) + ... + i * C(n, i) + ... - 4的那部分奇数) - 
4. 且 m <= n 肯定成立的
5. dp[4] = 4 * ....
6. dp[2] = 2 * (C....) - 2 * dp[4] / 4?
7. 似乎是成立的～
8. 对于给定的n, C(n, 1) + 2 * C(n, 2) + ... + i * C(n, i) + ... n * C(n, n)
9. 这个的有没有公式可以用的？
10. 这个求和式是：$\sum_{i=1}^{n} i \cdot C(n, i) = n \cdot 2^{n-1}$
