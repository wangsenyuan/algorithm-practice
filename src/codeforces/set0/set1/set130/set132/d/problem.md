# Problem D: Shakespeare Numbers

## Problem Statement

Shakespeare is a widely known esoteric programming language in which programs look like plays by Shakespeare, and numbers are given by combinations of ornate epithets. In this problem we will have a closer look at the way the numbers are described in Shakespeare.

Each constant in Shakespeare is created from non-negative powers of 2 using arithmetic operations. For simplicity we'll allow only addition and subtraction and will look for a representation of the given number which requires a minimal number of operations.

You are given an integer n. You have to represent it as n = a₁ + a₂ + ... + aₘ, where each of aᵢ is a non-negative power of 2, possibly multiplied by -1. Find a representation which minimizes the value of m.

## Input

The only line of input contains a positive integer n, written as its binary notation. The length of the notation is at most 10^6. The first digit of the notation is guaranteed to be 1.

## Output

Output the required minimal m. After it output m lines. Each line has to be formatted as "+2^x" or "-2^x", where x is the power coefficient of the corresponding term. The order of the lines doesn't matter.

## Examples

### Example 1
**Input:**
```
1111
```

**Output:**
```
2
+2^4
-2^0
```

### Example 2
**Input:**
```
1010011
```

**Output:**
```
4
+2^0
+2^1
+2^4
+2^6
```


### ideas
1. 一段连续的1，可以转换成两个1，一个+，一个-
2. 11100111 这个可以用4个，但貌似也能用3个
3. dp[i][0] 表示到目前为止，已经把前缀处理好了的最优操作数
4. dp[i][1] 表示到目前为止，前缀都处理好了，但是当前位，多出来一个1
5. 如果当前位 = 1
6.   dp[i][0] = dp[i-1][0] + 1 (+1， 对后续没有影响)
7.   dp[i][1] = dp[i-1][1] 这里不处理（对后续还有影响）
8. 如果当前位 = 0
9.   dp[i][0] = dp[i-1][0] (不操作)
10.           = dp[i-1][1] + 1 (增加一个-2 ^ i的操作)
11.  dp[i][1] = dp[i-1][0] + 1 (增加 + 1)
12. 