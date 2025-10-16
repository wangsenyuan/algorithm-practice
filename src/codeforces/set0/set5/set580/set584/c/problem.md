## Problem Statement

Marina loves strings of the same length and Vasya loves when there is a third string, different from them in exactly t characters. Help Vasya find at least one such string.

More formally, you are given two strings s1, s2 of length n and number t. Let's denote as f(a, b) the number of characters in which strings a and b are different. Then your task will be to find any string s3 of length n, such that f(s1, s3) = f(s2, s3) = t. If there is no such string, print -1.

## Input

The first line contains two integers n and t (1 ≤ n ≤ 10^5, 0 ≤ t ≤ n).

The second line contains string s1 of length n, consisting of lowercase English letters.

The third line contain string s2 of length n, consisting of lowercase English letters.

## Output

Print a string of length n, differing from string s1 and from s2 in exactly t characters. Your string should consist only from lowercase English letters. If such string doesn't exist, print -1.

## Examples

### Example 1

**Input**
```
3 2
abc
xyc
```

**Output**
```
ayd
```

### Example 2

**Input**
```
1 0
c
b
```

**Output**
```
-1
```


### ideas
1. 假设有 m个位置 s1[i] != s2[i], 那么就有w = n - m 个位置 s1[i] = s2[i]
2. 如果那w个位置中，有x个，s3和它们不同, 那么还剩余 y = t - x个位置
3. 那么 y <= m
4. y = min(m, t), x = t - y, x <= w
5. 这样子就有答案了 