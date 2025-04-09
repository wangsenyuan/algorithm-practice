You've got a string s = s1s2... s|s| of length |s|, consisting of lowercase English letters. There also are q queries, each query is described by two integers li, ri (1 ≤ li ≤ ri ≤ |s|). The answer to the query is the number of substrings of string s[li... ri], which are palindromes.

String s[l... r] = slsl + 1... sr (1 ≤ l ≤ r ≤ |s|) is a substring of string s = s1s2... s|s|.

String t is called a palindrome, if it reads the same from left to right and from right to left. Formally, if t = t1t2... t|t| = t|t|t|t| - 1... t1.

### ideas
1. 就是 s[l...r]中间的回文子串（连续的）
2. 如果知道f[r] = s[1...r]中间的回文子串的前缀和
3. res = f[r] - f[l] - 那些在l(中点 = l)前面的，但是延展到l后面的子串
4. 后面那部分怎么查呢？
5. 怎么预先计算出来就可以了
6. dp[c][i] 表示以c为中点，但是以i为有边界的，在i后面的回文的数量
7. 咋感觉又复杂了呢？