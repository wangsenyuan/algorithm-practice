Limak is a little polar bear. He doesn't have many toys and thus he often plays with polynomials.

He considers a polynomial valid if its degree is $n$ and its coefficients are integers not exceeding $k$ by the absolute value. More formally:

Let $a_0, a_1, ..., a_n$ denote the coefficients, so $P(x) = a_0 + a_1 \cdot x + a_2 \cdot x^2 + ... + a_n \cdot x^n$. 

Then, a polynomial $P(x)$ is valid if all the following conditions are satisfied:
- $a_i$ is integer for every $i$
- $|a_i| \leq k$ for every $i$
- $a_n \neq 0$

Limak has recently got a valid polynomial $P$ with coefficients $a_0, a_1, a_2, ..., a_n$. He noticed that $P(2) \neq 0$ and he wants to change it. He is going to change one coefficient to get a valid polynomial $Q$ of degree $n$ that $Q(2) = 0$. Count the number of ways to do so. You should count two ways as distinct if coefficients of target polynomials differ.

### Input

The first line contains two integers $n$ and $k$ ($1 \leq n \leq 200000$, $1 \leq k \leq 10^9$) — the degree of the polynomial and the limit for absolute values of coefficients.

The second line contains $n + 1$ integers $a_0, a_1, ..., a_n$ ($|a_i| \leq k$, $a_n \neq 0$) — describing a valid polynomial. It's guaranteed that $P(2) \neq 0$.

### Output

Print the number of ways to change one coefficient to get a valid polynomial $Q$ that $Q(2) = 0$.

### Examples

**Input**
```
3 1000000000
10 -9 -3 5
```

**Output**
```
3
```

**Input**
```
3 12
10 -9 -3 5
```

**Output**
```
2
```

**Input**
```
2 20
14 -7 19
```

**Output**
```
0
```

### Note

In the first sample, we are given a polynomial $P(x) = 10 - 9x - 3x^2 + 5x^3$.

Limak can change one coefficient in three ways:

1. He can set $a_0 = -10$. Then he would get $Q(x) = -10 - 9x - 3x^2 + 5x^3$ and indeed $Q(2) = -10 - 18 - 12 + 40 = 0$.
2. Or he can set $a_2 = -8$. Then $Q(x) = 10 - 9x - 8x^2 + 5x^3$ and indeed $Q(2) = 10 - 18 - 32 + 40 = 0$.
3. Or he can set $a_1 = -19$. Then $Q(x) = 10 - 19x - 3x^2 + 5x^3$ and indeed $Q(2) = 10 - 38 - 12 + 40 = 0$.

In the second sample, we are given the same polynomial. This time though, $k$ is equal to 12 instead of $10^9$. Two first of ways listed above are still valid but in the third way we would get $|a_1| > k$ which is not allowed. Thus, the answer is 2 this time.


### ideas
1. too hard～
2. 没法把 2带进去～
3. P(2) 等于多少，都算不出来～
4. P(2) = a0 + a1 * 2 + a2 * 4 + a3 * 8... + an * pow(2, n)
5. Q(2) = 大部分是一样的，只有1个系数和P不相同
6. P(2)如果看作是2的系数，那么可以把它变成一个2进制表示
7. an * pow(2, n) 相当于把an左移n位，。。。。
8. 这样子，可以算出来每一位的系数（变成0/1）
9. 居然有负值