# Problem G: Binary String Substring Analysis

## Problem Description

You are given a binary string $s_1s_2\ldots s_n$ of length $n$. A string $s$ is called binary if it consists only of zeros and ones.

For a string $p$, we define the function $f(p)$ as the maximum number of occurrences of any character in the string $p$. For example:
- $f(00110) = 3$ (character '0' appears 3 times)
- $f(01) = 1$ (both characters appear only once)

You need to find the sum $f(s_ls_{l+1}\ldots s_r)$ for all pairs $1 \leq l \leq r \leq n$.

## Input

Each test consists of multiple test cases. The first line contains a single integer $t$ ($1 \leq t \leq 10^4$) — the number of test cases. Then follows their descriptions.

The first line of each test case contains a single integer $n$ ($1 \leq n \leq 2 \cdot 10^5$) — the length of the binary string.

The second line of each test case contains a string of length $n$, consisting of 0s and 1s — the binary string $s$.

**Constraints:**
- It is guaranteed that the sum of $n$ across all test cases does not exceed $2 \cdot 10^5$.

## Output

For each test case, output the sum $f(s_ls_{l+1}\ldots s_r)$ for all pairs $1 \leq l \leq r \leq n$.

## Example

### Input
```
6
1
0
2
01
4
0110
6
110001
8
10011100
11
01011011100
```

### Output
```
1
3
14
40
78
190
```

## Explanation

**Test Case 1:** The string $s$ has one substring, and the value $f(0) = 1$.

**Test Case 2:** All substrings of the string $s$ are:
- $0$ → $f(0) = 1$
- $01$ → $f(01) = 1$ 
- $1$ → $f(1) = 1$

Answer: $1 + 1 + 1 = 3$

**Test Case 3:** All substrings of the string $s$ are:
- $0$ → $f(0) = 1$
- $01$ → $f(01) = 1$
- $011$ → $f(011) = 2$ (character '1' appears 2 times)
- $0110$ → $f(0110) = 2$ (character '1' appears 2 times)
- $1$ → $f(1) = 1$
- $11$ → $f(11) = 2$ (character '1' appears 2 times)
- $110$ → $f(110) = 2$ (character '1' appears 2 times)
- $1$ → $f(1) = 1$
- $10$ → $f(10) = 1$
- $0$ → $f(0) = 1$

Answer: $1 + 1 + 2 + 2 + 1 + 2 + 2 + 1 + 1 + 1 = 14$


## ideas
1. 考虑区间l...r，且假设这个区间里面1的个数多于0的个数，那么就是要在答案中增加1的个数
2. pref[r] - pref[l-1] >= r - pref[r] - (l - 1 - pref[l-1])
3. 2 * pref[r] - r >= 2 * pref[l-1] - (l - 1)
4. 就是要在 2 * pref[i] - i 处记录 pref[i]的值
5. cnt * (这样的i) - 这样的sum 
6. 考虑 0比1多的情况 
7. 2 * pref[r] - r < 2 * pref[l-1] - (l - 1)
8.  (r - pref[r]) 