### Problem

Calculate the value of the sum: $n \bmod 1 + n \bmod 2 + n \bmod 3 + \ldots + n \bmod m$. 

As the result can be very large, you should print the value modulo $10^9 + 7$ (the remainder when divided by $10^9 + 7$).

The modulo operator $a \bmod b$ stands for the remainder after dividing $a$ by $b$. For example $10 \bmod 3 = 1$.

### Input

The only line contains two integers $n$, $m$ ($1 \le n, m \le 10^{13}$) — the parameters of the sum.

### Output

Print integer $s$ — the value of the required sum modulo $10^9 + 7$.

### Examples

#### Example 1
**Input**
```
3 4
```
**Output**
```
4
```

#### Example 2
**Input**
```
4 4
```
**Output**
```
1
```

#### Example 3
**Input**
```
1 1
```
**Output**
```
0
```


### ideas
1. when n < i, 那么 n % i = n
2. 那么这样的数有 m - n 个
3. 现在考虑 i <= n, 1, 2, 3.... n
4. n = i * ? + r(i)
5.  比如 n = 15, [1，3, 5, 15]
6.  0, 1, 0, 3, 0, 3, 1, 7, 6, 5, 4, 3, 2, 1, 0
7.  