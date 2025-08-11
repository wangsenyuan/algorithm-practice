# Martian Lucky Numbers

## Problem Description

You know that the Martians use a number system with base $k$. Digit $b$ ($0 \leq b < k$) is considered lucky, as the first contact between the Martians and the Earthlings occurred in year $b$ (by Martian chronology).

A **digital root** $d(x)$ of number $x$ is a number that consists of a single digit, resulting after cascading summing of all digits of number $x$. The word "cascading" means that if the first summing gives us a number that consists of several digits, then we sum up all digits again, and again, until we get a one digit number.

### Example
For example, $d(35047) = d((3 + 5 + 0 + 4)_7) = d(15_7) = d((1 + 5)_7) = d(6_7) = 6_7$. In this sample the calculations are performed in the 7-base notation.

If a number's digital root equals $b$, the Martians also call this number **lucky**.

## Task

You have string $s$, which consists of $n$ digits in the $k$-base notation system. Your task is to find how many distinct substrings of the given string are lucky numbers. **Leading zeroes are permitted** in the numbers.

### Note on Substrings
Substring $s[i...j]$ of the string $s = a_1a_2...a_n$ ($1 \leq i \leq j \leq n$) is the string $a_ia_{i+1}...a_j$. Two substrings $s[i_1...j_1]$ and $s[i_2...j_2]$ of the string $s$ are different if either $i_1 \neq i_2$ or $j_1 \neq j_2$.

## Input

The first line contains three integers $k$, $b$ and $n$ ($2 \leq k \leq 10^9$, $0 \leq b < k$, $1 \leq n \leq 10^5$).

The second line contains string $s$ as a sequence of $n$ integers, representing digits in the $k$-base notation: the $i$-th integer equals $a_i$ ($0 \leq a_i < k$) — the $i$-th digit of string $s$. The numbers in the lines are space-separated.

## Output

Print a single integer — the number of substrings that are lucky numbers.

## Important Note

Please, do not use the `%lld` specifier to read or write 64-bit integers in C++. It is preferred to use the `cin`, `cout` streams or the `%I64d` specifier.

## Examples

### Example 1
**Input:**
```
10 5 6
3 2 0 5 6 1
```

**Output:**
```
5
```

**Explanation:** The following substrings have the sought digital root:
- $s[1...2] = "3 2"$
- $s[1...3] = "3 2 0"$
- $s[3...4] = "0 5"$
- $s[4...4] = "5"$
- $s[2...6] = "2 0 5 6 1"$

### Example 2
**Input:**
```
7 6 4
3 5 0 4
```

**Output:**
```
1
```

### Example 3
**Input:**
```
257 0 3
0 0 256
```

**Output:**
```
3
```

### ideas
1. 假设到i为止，可以获得的d(x)的集合为dp[i], 那么新加入一个数后，从d(x)计算出新的nd(x)是比较容易的
2. 但是问题出在这个集合的大小；
3. 感觉和前缀和类似吧。
4. 假设到目前为止的前缀d(x) = w, 那么有一段包括i的子区间的d(x) = b
5. 那么我们考虑前面某个区间的d(x) = u, 那么有 (u + b)(k) = w
6. 那么这个u的值要怎么计算出来？
7. 如果 u + b 有进位，那么 u + b = (1, y - 1)的组合
8. 如果没有进位，那么就是 u + b = y
9. got