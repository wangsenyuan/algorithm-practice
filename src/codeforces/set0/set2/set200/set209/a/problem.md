# Zebroid Subsequence Count

## Problem Statement
Polycarpus has arranged n marbles in a row from left to right. The marbles are colored either red or blue, and they form what is called a "zebroid" pattern.

## Definition
A zebroid is a non-empty sequence of marbles where the colors alternate between red and blue. For example:
- Valid zebroids: (red, blue, red) or (blue)
- Invalid zebroid: (red, red)

## Task
Find the number of ways to select a zebroid subsequence from the given sequence. The answer should be computed modulo $10^9 + 7$.

## Input Format
- A single integer n ($1 \leq n \leq 10^6$)
  - Represents the number of marbles in the sequence

## Output Format
- A single integer
  - The number of possible zebroid subsequences
  - Modulo $10^9 + 7$

## ideas
- 对于给定的n, 序列就是(red, blue, red, ...)
- 那么其中的任何连续的一段都是ok的？ n * (n + 1) / 2
- 但是 n = 4不成立， 4 * 5 / 2 = 10, 但是答案是11
- subsequence是可以跳着选的吧
- rbrb, 单个的有4个，两个rb/br的有3个， rbr/brb 2个， rbrb是1个
- 但是第一个和地四个组成rb 也是1个。 所以是11个
- dp[i]表示到以i结尾的数, dp[i] = dp[i-1] + dp[i-3] + dp[i-5]
- 
