# Problem Statement

## Definitions
- An **allowed letter** is a lowercase letter that is one of the first k letters of the Latin alphabet.
- A string t is called **pleasant** if t is a subsequence of string s.
- A sequence t is a **subsequence** of sequence s if t can be obtained from s by deleting several (possibly zero or all) elements from arbitrary positions.

## Input
- First line: Two integers n and k (1 ≤ n ≤ 10^6; 1 ≤ k ≤ 26) — the length of string s and the number of allowed letters.
- Second line: String s of length n, consisting only of allowed letters.
- Third line: Integer q (1 ≤ q ≤ 2·10^5) — the number of queries.
- Next q lines: Each line contains a string ti, consisting only of allowed letters.

### Additional Constraint
- The total length of all ti does not exceed 10^6.

## Output
For each query, output one integer — the minimum number of allowed letters that need to be appended to the string on the right so that it stops being pleasant.

## ideas
1. 肯定会有个答案， 比如k = 1时， f(t) = len(s) + 1 - len(t) (将t增加到比s多一位)
2. 如果t不是s的子序列，那么f(t) = 0
3. 如果t是s的子序列，目前匹配到的位置是s[i]（s[:i]包含t的最小前缀）
4. 接下来append一个x，要怎么处理呢？
5. dp[i][x]表示如果s[i]处，匹配了x，需要增加的最少的字符数
6. dp[i][x] = dp[next[i][x]][?] + 1
7. dp[i] = 如果在s[i]处，使用某个字符不匹配s[i]后，需要增加的最少字符数
8. dp[i] = min(dp[j], j = next[i][x]) + 1 (在i处使用x)
9. 