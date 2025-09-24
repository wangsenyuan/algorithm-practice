# Problem E: Subsequence Counting

## Problem Statement

After getting kicked out of her reporting job for not knowing the alphabet, Bessie has decided to attend school at the Fillet and Eggs Eater Academy. She has been making good progress with her studies and now knows the first k English letters.

Each morning, Bessie travels to school along a sidewalk consisting of m + n tiles. In order to help Bessie review, Mr. Moozing has labeled each of the first m sidewalk tiles with one of the first k lowercase English letters, spelling out a string t. Mr. Moozing, impressed by Bessie's extensive knowledge of farm animals, plans to let her finish labeling the last n tiles of the sidewalk by herself.

Consider the resulting string s (|s| = m + n) consisting of letters labeled on tiles in order from home to school. For any sequence of indices p₁ < p₂ < ... < pₑ we can define subsequence of the string s as string sₚ₁sₚ₂...sₚₑ. Two subsequences are considered to be distinct if they differ as strings. Bessie wants to label the remaining part of the sidewalk such that the number of distinct subsequences of tiles is maximum possible. However, since Bessie hasn't even finished learning the alphabet, she needs your help!

**Note:** Empty subsequence also counts.

## Input

The first line of the input contains two integers n and k (0 ≤ n ≤ 1,000,000, 1 ≤ k ≤ 26).

The second line contains a string t (|t| = m, 1 ≤ m ≤ 1,000,000) consisting of only first k lowercase English letters.

## Output

Determine the maximum number of distinct subsequences Bessie can form after labeling the last n sidewalk tiles each with one of the first k lowercase English letters. Since this number can be rather large, you should print it modulo 10⁹ + 7.

**Important:** You are not asked to maximize the remainder modulo 10⁹ + 7! The goal is to maximize the initial value and then print the remainder.

## Examples

### Example 1
**Input:**
```
1 3
ac
```

**Output:**
```
8
```

### Example 2
**Input:**
```
0 2
aaba
```

**Output:**
```
10
```

## Explanation

**Example 1:** The optimal labeling gives 8 different subsequences: "" (the empty string), "a", "c", "b", "ac", "ab", "cb", and "acb".

**Example 2:** The entire sidewalk is already labeled. There are 10 possible different subsequences: "" (the empty string), "a", "b", "aa", "ab", "ba", "aaa", "aab", "aba", and "aaba". Note that some strings, including "aa", can be obtained with multiple sequences of tiles, but are only counted once.


### ideas
1. 考虑在n = 0的时候，如何计算不同的子序列
2. aa "", a, aa (3个)
3. ab "", a, b, ab (4个)
4. 如果已经存在一个a，那么当前的a就不能重复算
5. 推广，就是前面已经存在子序列的情况下，abc, 那么ab + c就不能重复算
6. 但是记录所有已经存在的子序列，肯定是不行的（预期是指数级的）
7. 但是奇怪的是，这里要求的是最多数量的求模
8. 要么就是这个不用求模的情况下，就能算出多少个字符(能枚举出来)
9. 要么就是这个最多的子序列的，符合某种模式
10. 用到的字符越多越好。
11. 考虑 m = 0 的情况，考虑 123123123
12. 剩下的感觉就是 1...k1...k这样的循环模式（第一个不一定是1，要和前一个字符有差别）
13. 第二个也可能有些不同。但是前k个应该是按照贪心确定的，然后按照这个模式滚动下去
14. dp[i] = fp[?] + 当前使用x时的计数
15. 如果x不是y(x != y)的后面， 可以直接加上去。但是如果x前面是x的时候，
16. dp[i] = sum(fp[?]) + 1 - fp[x]