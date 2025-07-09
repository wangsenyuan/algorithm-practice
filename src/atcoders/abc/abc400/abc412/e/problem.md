# Problem E

**Time Limit:** 4 sec / **Memory Limit:** 1024 MiB

**Score:** 500 points

## Problem Statement

For a positive integer $n$, define $A_n$ as the least common multiple of $1, 2, \ldots, n$.

You are given positive integers $L$ and $R$. How many distinct integers are contained in the sequence $(A_L, A_{L+1}, \ldots, A_R)$?

## Constraints

- $1 \leq L \leq R \leq 10^{14}$
- $R - L \leq 10^7$
- $L$ and $R$ are integers.

## Input

The input is given from Standard Input in the following format:

```
L R
```

## Output

Print the number of distinct integers in the sequence $(A_L, A_{L+1}, \ldots, A_R)$.

## Sample Input 1

```
2 4
```

## Sample Output 1

```
3
```

## Sample Input 2

```
1 1
```

## Sample Output 2

```
1
```

## Sample Input 3

```
10 20
```

## Sample Output 3

```
6
```


### ideas
1. lcm(x, y) = x和y中，质因数的最高次的乘积
2. lcm(1.....n) 假设知道把所有数，都按照质因数分解后，那么F(n) = (p2, p3, .... pk)
3. p2表示1...n中2的最高次，p3表示3的最高次。。。。pk=k的最高次
4. F(n + 1) >= F(n), n+1，考虑它的质因数分解，如果p2 > F(n)的p2，那么它肯定产生了一个新的数
5. ...只要有一个pi，那么它就产生了一个新的数
6. L...R需要1e7次迭代。所以这里的问题是要快速质因数分解
7. 好像，不需要 >= 1e7 的质数，因为这样的质数，最多出现一次（第二次出现肯定在外部了）
8. 