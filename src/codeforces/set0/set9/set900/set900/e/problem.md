# Problem Description

Vasya wrote down two strings s of length n and t of length m consisting of small English letters 'a' and 'b'. What is more, he knows that string t has a form "abab...", namely there are letters 'a' on odd positions and letters 'b' on even positions.

Suddenly in the morning, Vasya found that somebody spoiled his string. Some letters of the string s were replaced by character '?'.

Let's call a sequence of positions i, i + 1, ..., i + m - 1 as occurrence of string t in s, if 1 ≤ i ≤ n - m + 1 and t₁ = sᵢ, t₂ = sᵢ₊₁, ..., tₘ = sᵢ₊ₘ₋₁.

The boy defines the beauty of the string s as maximum number of disjoint occurrences of string t in s. Vasya can replace some letters '?' with 'a' or 'b' (letters on different positions can be replaced with different letter). Vasya wants to make some replacements in such a way that beauty of string s is maximum possible. From all such options, he wants to choose one with the minimum number of replacements. Find the number of replacements he should make.

## Input

The first line contains a single integer n (1 ≤ n ≤ 10⁵) — the length of s.

The second line contains the string s of length n. It contains small English letters 'a', 'b' and characters '?' only.

The third line contains a single integer m (1 ≤ m ≤ 10⁵) — the length of t. The string t contains letters 'a' on odd positions and 'b' on even positions.

## Output

Print the only integer — the minimum number of replacements Vasya has to perform to make the beauty of string s the maximum possible.

## ideas
1. t = ababa...
2. 所以，对于s[i:i+m]，如果能够得到t, 那么在对应位置上的修改是确定的（奇数位置为a，偶数位置为b)
3. dp[i] = j 表示从i开始能够匹配的最长的abab的长度
4. dp[i] = dp[i+2] + 2 如果 s[i:i+1] = ab(或者修改后能够得到)
5. 如果是?a, dp[i] = 1, 如果是b?, dp[i] = 0
6. 搞错了。 是要得到最小的修改， 而不是最大的beauty
7. 