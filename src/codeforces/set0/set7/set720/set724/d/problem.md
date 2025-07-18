# Problem Description

You are given a string `s`, consisting of lowercase English letters, and the integer `m`.

## Task

One should choose some symbols from the given string so that any contiguous subsegment of length `m` has at least one selected symbol. Note that here we choose positions of symbols, not the symbols themselves.

Then one uses the chosen symbols to form a new string. All symbols from the chosen position should be used, but we are allowed to rearrange them in any order.

## Formal Definition

We choose a subsequence of indices $1 \leq i_1 < i_2 < ... < i_t \leq |s|$. The selected sequence must meet the following condition: for every $j$ such that $1 \leq j \leq |s| - m + 1$, there must be at least one selected index that belongs to the segment $[j, j + m - 1]$, i.e. there should exist a $k$ from $1$ to $t$, such that $j \leq i_k \leq j + m - 1$.

Then we take any permutation $p$ of the selected indices and form a new string $s_{i_{p_1}}s_{i_{p_2}}...s_{i_{p_t}}$.

**Find the lexicographically smallest string that can be obtained using this procedure.**

## Input

- The first line of the input contains a single integer $m$ ($1 \leq m \leq 100,000$).
- The second line contains the string $s$ consisting of lowercase English letters. It is guaranteed that this string is non-empty and its length doesn't exceed $100,000$. It is also guaranteed that the number $m$ doesn't exceed the length of the string $s$.

## Output

Print the single line containing the lexicographically smallest string that can be obtained using the procedure described above.

## Examples

### Example 1
**Input:**
```
3
cbabc
```

**Output:**
```
a
```

**Explanation:** One can choose the subsequence $\{3\}$ and form a string "a".

### Example 2
**Input:**
```
2
abcab
```

**Output:**
```
aab
```

**Explanation:** One can choose the subsequence $\{1, 2, 4\}$ (symbols on these positions are 'a', 'b' and 'a') and rearrange the chosen symbols to form a string "aab".

### Example 3
**Input:**
```
3
bcabcbaccba
```

**Output:**
```
aaabb
```

### ideas
1. 假设前一个选择的位置i,那么在 i + 1...i+k中间必须选一个新的字符，那么就选择其中最小的那个字符
2. 如果有相同的，就选择最靠右边的那个