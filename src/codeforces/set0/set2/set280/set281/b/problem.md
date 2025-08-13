# Problem B: Nearest Fraction

## Problem Description

You are given three positive integers $x$, $y$, $n$. Your task is to find the nearest fraction to $\frac{x}{y}$ whose denominator is no more than $n$.

Formally, you should find such pair of integers $a$, $b$ ($1 \leq b \leq n$; $0 \leq a$) that the value $|\frac{a}{b} - \frac{x}{y}|$ is as minimal as possible.

If there are multiple "nearest" fractions, choose the one with the minimum denominator. If there are multiple "nearest" fractions with the minimum denominator, choose the one with the minimum numerator.

## Input

A single line contains three integers $x$, $y$, $n$ ($1 \leq x, y, n \leq 10^5$).

## Output

Print the required fraction in the format "a/b" (without quotes).

## Examples

### Example 1
**Input:**
```
3 7 6
```

**Output:**
```
2/5
```

### Example 2
**Input:**
```
7 2 4
```

**Output:**
```
7/2
```


## ideas
1. 先将 x/y化简 如果 y <= n => a/b = x / y (gcd)
2. y > n, 先处理 a / b < x / y
3. 能二分吗？
4. 好像不能二分。这个时候可能0.1得不到，但是0.05的差距反而能够得到
5. 还有一个就是迭代b(1...n)
6. 但是这样，似乎有点慢？