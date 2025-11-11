# Problem E: Two Strings and Types

## Description

Analyzing the mistakes people make while typing search queries is a complex and an interesting work. As there is no guaranteed way to determine what the user originally meant by typing some query, we have to use different sorts of heuristics.

Polycarp needed to write a code that could, given two words, check whether they could have been obtained from the same word as a result of typos. Polycarpus suggested that the most common typo is skipping exactly one letter as you type a word.

Implement a program that can, given two distinct words S and T of the same length n determine how many words W of length n + 1 are there with such property that you can transform W into both S, and T by deleting exactly one character. Words S and T consist of lowercase English letters. Word W also should consist of lowercase English letters.

## Input

The first line contains integer n (1 ≤ n ≤ 100 000) — the length of words S and T.

The second line contains word S.

The third line contains word T.

Words S and T consist of lowercase English letters. It is guaranteed that S and T are distinct words.

## Output

Print a single integer — the number of distinct words W that can be transformed to S and T due to a typo.

## Examples

### Example 1

**Input:**
```
7
reading
trading
```

**Output:**
```
1
```

### Example 2

**Input:**
```
5
sweet
sheep
```

**Output:**
```
0
```

### Example 3

**Input:**
```
3
toy
try
```

**Output:**
```
2
```

## Note

In the first sample test the two given words could be obtained only from word "treading" (the deleted letters are marked in bold).

In the second sample test the two given words couldn't be obtained from the same word by removing one letter.

In the third sample test the two given words could be obtained from either word "tory" or word "troy".


### ideas
1. 找到这样的X，长度为n+1，删除1个字符后，可以变成S，删除另外一个字符后，可以变成T
2. let X = $x_1x_2...x_i...x_j....x_nx_{n+1}$
3. 删除i后， X = pref + suf
4. 假设前i个位置， S[:i] = T[:i], 现在可以考虑在位置i处添加一个字符
5. 有3种情况， s[i] != t[i], 那么可以添加一个t[i], 此时，要求t的后缀在某个地方添加了某个s[j]后， s[i:] = t[i+1] + ?
6. 或者添加s[i]
7. 或者 s[i] = t[i], 那么这里添加任意字符，但是后缀必须一致
8. 第三个情况很好处理
9. 主要第一种和第二种情况
10. dp[i][0] 表示后缀， s[i:] = t[i:]
11. dp[i][1] 表示后缀   s[i:] = t[i+1:] 且在某个地方增加了一个字符后，相等
12. dp[i][-1] 表示后缀 s[i:] = t[i-1:] 且在某个地方删除了一个字符后，相等
13. dp[i][0] 很好处理 = s[i] = t[i] && dp[i+1][0]
14. dp[i][1] = s[i] = t[i+1] && dp[i+1][1] or dp[i+1][0] (替换t[i]为s[i]后)
15. dp[i][-1] = s[i] = t[i-1] && dp[i+1][-1] or dp[i+1][0]
16. 
17. 