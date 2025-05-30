# Problem Statement

A subsequence of length $|x|$ of string $s = s_1s_2... s_{|s|}$ (where $|s|$ is the length of string $s$) is a string $x = s_{k_1}s_{k_2}... s_{k_{|x|}}$ ($1 \leq k_1 < k_2 < ... < k_{|x|} \leq |s|$).

You've got two strings — $s$ and $t$. Let's consider all subsequences of string $s$, coinciding with string $t$. Is it true that each character of string $s$ occurs in at least one of these subsequences? In other words, is it true that for all $i$ ($1 \leq i \leq |s|$), there is such subsequence $x = s_{k_1}s_{k_2}... s_{k_{|x|}}$ of string $s$, that $x = t$ and for some $j$ ($1 \leq j \leq |x|$) $k_j = i$?

---

## Input

The first line contains string $s$, the second line contains string $t$. Each line consists only of lowercase English letters. The given strings are non-empty, the length of each string does not exceed $2 \cdot 10^5$.

---

## Output

Print `Yes` (without the quotes), if each character of the string $s$ occurs in at least one of the described subsequences, or `No` (without the quotes) otherwise.


### ideas
1. 对于i，如果存在t的一个前缀t[:j]能匹配s[:i-1], 一个后缀t[k:] 匹配s[i+1:]
2. 且在j...k中间存在s[i], 那么就是ok的
3. 