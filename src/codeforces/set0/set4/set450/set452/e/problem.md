# Problem

You are given three strings $s_1, s_2, s_3$. For each integer $l$ such that $1 \le l \le \min(|s_1|, |s_2|, |s_3|)$, you need to find how many triples $(i_1, i_2, i_3)$ exist such that the three substrings

- $s_1[i_1 \dots i_1 + l - 1]$
- $s_2[i_2 \dots i_2 + l - 1]$
- $s_3[i_3 \dots i_3 + l - 1]$

are pairwise equal (i.e., all three substrings are exactly the same string). Print each answer modulo $10^9 + 7$.

## Input

- The first three lines contain three non-empty strings $s_1, s_2, s_3$.
- The sum of their lengths does not exceed $3 \cdot 10^5$.
- All strings consist only of lowercase English letters.

## Output

Output $\min(|s_1|, |s_2|, |s_3|)$ integers, separated by spaces, where the $l$-th number is the required count for this $l$, taken modulo $10^9 + 7$.

## Examples

### Example 1

**Input:**

```text
abc
bc
cbc
```

**Output:**

```text
3 1
```

### Example 2

**Input:**

```text
abacaba
abac
abcd
```

**Output:**

```text
11 2 0 0
```

## Notes

Consider a string $t = t_1 t_2 \dots t_{|t|}$, where $t_i$ denotes the $i$-th character and $|t|$ is the length of $t$.

The notation $t[i \dots j]$ (for $1 \le i \le j \le |t|$) denotes the substring $t_i t_{i+1} \dots t_j$.

## Ideas

1. 对于长度l, 计算有多少个位置(i1, i2, i3)的组合，满足 s1[i1:i1+l] = s2[i2:i2+l] = s3[i3:i3+l]
2. 假设有这样一个树形结构，路径表示一个字符串（深度表示这个字符串的长度）, 在每个叶子节点上，存在三组坐标set1, set2, set3；set1表示这个字符串结束时的下标的集合
3. 那么这个节点的贡献 = len(set1) * len(set2) * len(set3)
4. 感觉是后缀树（后缀树包含了所有子串），记录s1的状态（到对应节点计数）

## Editorial

This problem requires one to use one of the datastructures, such as suffix array, suffix tree or suffix automata. The easiest solution uses a compressed suffix tree. Build one suffix tree on all three strings. For simplicity add some non-alphabetic character at the end of each string.

For every node in the tree store how many times the corresponding suffix occurs in each string. Then traverse the tree once. If the tree had no shortcuts, for every node that is $a$ characters away from the root you would have increased the answer for $a$ by the product of numbers of occurrences of the suffix in each of the strings.

Since you do have shortcuts, you need to update the answer for all the lengths from $a$ to $b$, where $a$ and $b$ are the distances of two ends of the shortcut from the root. One way to do it with constant time updates and linear time to print all the answers is the following:

If the array of answers is $v$, then instead of computing $v$ we can compute the array of differences $p$, such that $p_i = v_i - v_{i-1}$. This way when you traverse the shortcut, rather than adding some value at all the positions from $a$ to $b$, you only need to add that value at position $a$, and subtract it at position $b$. When $p$ is computed, it is easy to restore $v$ in one pass.

## Solution Summary

### Algorithm: Suffix Automaton (SAM) + Difference Array

**Step 1: Build SAM on concatenation**

Concatenate all three strings with unique separators: `s1 + '#' + s2 + '$' + s3 + '%'`. The separators ensure no cross-string substring is created. Build a SAM on this concatenated string.

**Step 2: Count occurrences per string**

Walk each original string `s1`, `s2`, `s3` through the SAM from the root. At each character position, we land in some SAM state — increment that state's count for the corresponding string. This records how many times each state is an endpoint of a substring from each string.

**Step 3: Propagate counts up suffix links**

Sort SAM states by length (descending). For each state, add its counts to its suffix-link parent. After propagation, `cnt[d][v]` = total number of occurrences of all substrings represented by state `v` in string `d`.

**Step 4: Difference array for answer**

Each SAM state `v` represents substrings of lengths `(parent.length + 1)` to `v.length`. The contribution to each such length is `cnt[0][v] * cnt[1][v] * cnt[2][v]`. Use a difference array: add the product at `lo = parent.length + 1`, subtract it at `hi + 1 = v.length + 1`. Finally, prefix-sum the array to get all answers.

### Example 1 Walkthrough

```
s1 = "abc", s2 = "bc", s3 = "cbc"
```

Sorted by length: `arr[0] = "bc"`, `arr[1] = "abc"`, `arr[2] = "cbc"`. Answer length = min(3,2,3) = 2.

**For l = 1 (single characters):**

| char | positions in "bc" | positions in "abc" | positions in "cbc" | triples |
|------|-------------------|--------------------|--------------------|---------|
| a    | 0                 | 1                  | 0                  | 0*1*0=0 |
| b    | 1                 | 1                  | 1                  | 1*1*1=1 |
| c    | 1                 | 1                  | 2                  | 1*1*2=2 |

Total for l=1: 0 + 1 + 2 = **3**

**For l = 2 (length-2 substrings):**

| substr | positions in "bc" | positions in "abc" | positions in "cbc" | triples |
|--------|-------------------|--------------------|--------------------|---------|
| ab     | 0                 | 1                  | 0                  | 0       |
| bc     | 1                 | 1                  | 1                  | 1*1*1=1 |
| cb     | 0                 | 0                  | 1                  | 0       |

Total for l=2: **1**

**Answer: `3 1`**

### Example 2 Walkthrough

```
s1 = "abacaba", s2 = "abac", s3 = "abcd"
```

Sorted by length: `arr[0] = "abac"`, `arr[1] = "abcd"`, `arr[2] = "abacaba"`. Answer length = 4.

**For l = 1:**

| char | in "abac" | in "abcd" | in "abacaba" | triples |
|------|-----------|-----------|--------------|---------|
| a    | 2         | 1         | 4            | 2*1*4=8 |
| b    | 1         | 1         | 2            | 1*1*2=2 |
| c    | 1         | 1         | 1            | 1*1*1=1 |
| d    | 0         | 1         | 0            | 0       |

Total for l=1: 8 + 2 + 1 = **11**

**For l = 2:**

| substr | in "abac" | in "abcd" | in "abacaba" | triples |
|--------|-----------|-----------|--------------|---------|
| ab     | 1         | 1         | 2            | 1*1*2=2 |
| ba     | 1         | 0         | 2            | 0       |
| ac     | 1         | 0         | 1            | 0       |
| bc     | 0         | 1         | 0            | 0       |
| cd     | 0         | 1         | 0            | 0       |

Total for l=2: **2**

**For l = 3:**

| substr | in "abac" | in "abcd" | in "abacaba" | triples |
|--------|-----------|-----------|--------------|---------|
| aba    | 1         | 0         | 2            | 0       |
| bac    | 1         | 0         | 1            | 0       |
| abc    | 0         | 1         | 0            | 0       |
| bcd    | 0         | 1         | 0            | 0       |

Total for l=3: **0**

**For l = 4:** No length-4 substring appears in all three strings. Total: **0**

**Answer: `11 2 0 0`**

### How the SAM computes this efficiently

Instead of enumerating every possible substring, the SAM groups substrings into equivalence classes (states). Each state `v` represents a set of substrings with lengths in `(link.length, v.length]` that all share the same set of ending positions. After propagating counts up suffix links, we know exactly how many times each equivalence class appears in each string. The product of three counts gives the number of triples, and the difference array distributes this across all lengths the state covers — all in $O(|s_1| + |s_2| + |s_3|)$ time.
