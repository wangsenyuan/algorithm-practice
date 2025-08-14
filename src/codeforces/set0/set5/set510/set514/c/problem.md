# Problem C: Watto and Mechanism

## Problem Description

Watto, the owner of a spare parts store, has recently got an order for the mechanism that can process strings in a certain way. Initially the memory of the mechanism is filled with $n$ strings. Then the mechanism should be able to process queries of the following type: "Given string $s$, determine if the memory of the mechanism contains string $t$ that consists of the same number of characters as $s$ and differs from $s$ in exactly one position".

Watto has already compiled the mechanism, all that's left is to write a program for it and check it on the data consisting of $n$ initial lines and $m$ queries. He decided to entrust this job to you.

## Input

The first line contains two non-negative numbers $n$ and $m$ ($0 \leq n \leq 3 \cdot 10^5$, $0 \leq m \leq 3 \cdot 10^5$) — the number of the initial strings and the number of queries, respectively.

Next follow $n$ non-empty strings that are uploaded to the memory of the mechanism.

Next follow $m$ non-empty strings that are the queries to the mechanism.

The total length of lines in the input doesn't exceed $6 \cdot 10^5$. Each line consists only of letters 'a', 'b', 'c'.

## Output

For each query print on a single line "YES" (without the quotes), if the memory of the mechanism contains the required string, otherwise print "NO" (without the quotes).

## Examples

### Example 1

**Input:**
```
2 3
aaaaa
acacaca
aabaa
ccacacc
caaac
```

**Output:**
```
YES
NO
NO
```

**Explanation:** 
- Query "aabaa" returns "YES" because it differs from "aaaaa" in exactly one position
- Query "ccacacc" returns "NO" because no string in memory differs from it in exactly one position
- Query "caaac" returns "NO" because no string in memory differs from it in exactly one position


### ideas
1. 每个字符串只有a,b,c组成
2. 然后 len(s) = len(t) 
3. cnt[s][a,b,c] = cnt[t][a,b,c]
4. 且只有一个位置不同，假设这个位置为s[i] = a, t[i] = b
5. 那么cnt[s][a] - 1 = cnt[t][a]
6. cnt[s][b] + 1 = cnt[t][b]
7. 然后就是更精确的匹配要怎么做？
8.  dp[i] 表示到i为止还匹配的的集合
9.  fp[i] 表示后缀到i为止还匹配的集合
10. 那么 dp[i-1] & fp[i+1] 有交集，且其中的s[i] != t[i]
11. 还有点难么～
12. 突破点是不是在a,b,c上？
13. 假设对所有的s，保存这样一个信息 (i, s[i], pref, suf)
14. 那么对于t的每个i来说，就是查找i, (a, b, c) != t[i], pref, suf 的检查
15. pref，suf可以用suf tree的状态节点来处理
16. 可以不用suffix tree， 用 trie就可以了
17. 