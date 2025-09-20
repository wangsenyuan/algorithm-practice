# Problem Description

Two positive integers a and b have a sum of s and a bitwise XOR of x. How many possible values are there for the ordered pair (a, b)?

## Input

The first line of the input contains two integers s and x (2 ≤ s ≤ 10^12, 0 ≤ x ≤ 10^12), the sum and bitwise xor of the pair of positive integers, respectively.

## Output

Print a single integer, the number of solutions to the given conditions. If no solutions exist, print 0.

## Examples

### Example 1
```
Input:
9 5

Output:
4
```

### Example 2
```
Input:
3 3

Output:
2
```

### Example 3
```
Input:
5 2

Output:
0
```

## Note

In the first sample, we have the following solutions: (2, 7), (3, 6), (6, 3), (7, 2).

In the second sample, the only solutions are (1, 2) and (2, 1).



### ideas
1. a + b = s, a ^ b = x
2. 从低位开始考虑, x[0] = 1, a[0] = 1 or b[1] = 1, 不会有进位 
3. x[0] = 0, a[0] = b[0] = (0/1)
4. dp[i][0/1] 表示到i位为止，a[:i] ^ b[:i] = x[:i], 且 a[:i] + b[i] = 也一致，且有无/进位