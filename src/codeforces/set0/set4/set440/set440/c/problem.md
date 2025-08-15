# Problem C: Prof. Vasechkin's Sum Representation

## Problem Description

Prof. Vasechkin wants to represent positive integer n as a sum of addends, where each addend is an integer number containing only 1s. For example, he can represent 121 as 121 = 111 + 11 + (–1). Help him to find the least number of digits 1 in such sum.

## Input

The first line of the input contains integer **n** (1 ≤ n < 10^15).

## Output

Print the expected minimal number of digits 1.

## Examples

### Example 1
**Input:**
```
121
```

**Output:**
```
6
```

**Explanation:** 121 can be represented as 111 + 11 + (–1), which contains 3 + 2 + 1 = 6 digits 1 in total.


### ideas
1. n = x - y
2. x = 1111 （比n大的11的序列） 
3. y = 确定的
4. 或者 n = x + y(x = 11111)