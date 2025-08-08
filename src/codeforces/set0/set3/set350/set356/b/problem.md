# Xenia and Hamming Distance

## Problem Description

The Hamming distance between two strings $s = s_1s_2\ldots s_n$ and $t = t_1t_2\ldots t_n$ of equal length $n$ is $\sum_{i=1}^{n} [s_i \ne t_i]$, where $[P]$ equals 1 if predicate $P$ is true and 0 otherwise.

String $a$ is the concatenation of $n$ copies of string $x$. String $b$ is the concatenation of $m$ copies of string $y$. It is guaranteed that the resulting strings $a$ and $b$ have the same length.

Given $n, x, m, y$, compute the Hamming distance between $a$ and $b$.

## Input

- The first line contains two integers $n$ and $m$ $(1 \le n, m \le 10^{12})$.
- The second line contains a non-empty string $x$.
- The third line contains a non-empty string $y$.

Both strings consist of at most $10^6$ lowercase English letters.

## Output

Print a single integer — the required Hamming distance.

## Examples

### Example 1

Input:
```
100 10
a
aaaaaaaaaa
```

Output:
```
0
```

### Example 2

Input:
```
1 1
abacaba
abzczzz
```

Output:
```
4
```

### Example 3

Input:
```
2 3
rzr
az
```

Output:
```
5
```

## Note

- In the first test case string $a$ is the same as string $b$ and equals 100 letters 'a', so the distance is 0.
- In the second test case the strings differ at the 3rd, 5th, 6th and 7th characters, so the distance is 4.
- In the third test case string $a$ is $rzrrzr$ and string $b$ is $azazaz$. They differ in all characters except the second one, so the distance is 5.

### ideas
1. 考虑某个位置x, x = n * u + i, 同时 x = m * v + j
2. 也就是说位置S[x] = s[i], T[x] = t[j]
3. 那么这时候(x, i) 同余（关于n), (x, j)同余（关于m) 
4. 那么 => (i - j) 能够整除(n - m)
5. 也就是说, 如果 ds = s * (n - m)
6. 迭代j的时候，i的位置是确定的
7. 对于给定的j, j + (n - m), j + 2 * (n - m), j + ?*(n-m) < n 这些位置i可以用j在某个地方匹配
8. h = lcm(n, m)后，就可以认为开始循环了， 然后最后的再计算一遍
9. 