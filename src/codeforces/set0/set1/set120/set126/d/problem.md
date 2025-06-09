# Fibonacci Sum Decomposition

## Problem Description

Fibonacci numbers have the following form:

- F₁ = 1
- F₂ = 2
- Fᵢ = Fᵢ₋₁ + Fᵢ₋₂, i > 2

Let's consider some non-empty set S = {s₁, s₂, ..., sₖ}, consisting of different Fibonacci numbers. Let's find the sum of values of this set's elements.

Let's call the set S a number n's decomposition into Fibonacci sum.

It's easy to see that several numbers have several decompositions into Fibonacci sum. For example:
- For 13 we have:
  - 13
  - 5 + 8
  - 2 + 3 + 8
  (three decompositions)
- For 16 we have:
  - 3 + 13
  - 1 + 2 + 13
  - 3 + 5 + 8
  - 1 + 2 + 5 + 8
  (four decompositions)

By the given number n determine the number of its possible different decompositions into Fibonacci sum.

## Input

The first line contains an integer t — the number of tests (1 ≤ t ≤ 10⁵). Each of the following t lines contains one test.

Each test is an integer n (1 ≤ n ≤ 10¹⁸).

**Note:** Please do not use the %lld specificator to read or write 64-bit integers in C++. It is preferred to use the cin, cout streams or the %I64d specificator.

## Output

For each input data test print a single number on a single line — the answer to the problem.

## ideas
1. 考虑f[i] 是第i个fib 数， dp[i]表示得到它的ways
2. dp[i] = dp[i-1] * dp[i-2] (f[i] = f[i-1] + f[i-2]) (有重复使用的情况)
3. 1, 2, 3, 5, 8, 13， 21
4. 5 = (2, 3) (f[4] = 1)
5. 8 = (5, 3), (5, 2, 1) = 2
6. 13 = (8, 5), (8, 3, 2) = 2
7. 21 = （13，8) (13...) = 3
8. dp[n] = dp[n-2] + 1 ? 
9. 好像是的，因为 f(n-1) > f(n) / 2 (至少是f(n)的一半)， 如果把f(n-1)拆分了，那么必然造成set中出现重复的数
10. 然后把 num 拆分成最大的几个数字的组合， num = s1 + s2 + .. + sk (s1 < s2 < .. < sk)
11. 这个表示一定存在吗？
12. 假设 <= f[i]的集合都存在
13. 那么 f(i) -> f(i+1) 肯定也是存在的。原因还在于 f(i)至少是f(i+1)的一半
14. 结果是dp[s1] * dp[s2] * dp[...] 感觉不是的
15. 因为 s2中可能出现s1
16. 把num一直拆分成fib的数的乘积
17. 那么连续的两个数，可以选择合并在一起
18. 也可以选择不合并
19. dp[i] = dp[i-1] + dp[i-2] * (如果能合并)