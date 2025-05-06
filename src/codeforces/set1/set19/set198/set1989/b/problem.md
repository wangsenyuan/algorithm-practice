You are given two strings 𝑎
 and 𝑏
, both consisting of lowercase Latin letters.

A subsequence of a string is a string which can be obtained by removing several (possibly zero) characters from the original string. A substring of a string is a contiguous subsequence of that string.

For example, consider the string abac:

a, b, c, ab, aa, ac, ba, bc, aba, abc, aac, bac and abac are its subsequences;
a, b, c, ab, ba, ac, aba, bac and abac are its substrings.
Your task is to calculate the minimum possible length of the string that contains 𝑎
 as a substring and 𝑏
 as a subsequence.

 ### ideas
 1. 假设答案是s，那么a是s的substring, b是s的子序列
 2. 如果b是a的subsequence，那么答案 = len(a)
 3. 所以，在a中找到b的最多的subsequence，（并且是连续的），然后把其他的添加到两头