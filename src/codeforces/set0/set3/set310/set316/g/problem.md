# Problem G: Good Substrings

## Problem Description

Smart Beaver recently got interested in a new word game. The point is as follows: count the number of distinct good substrings of some string $s$. To determine if a string is good or not the game uses rules. Overall there are $n$ rules. Each rule is described by a group of three $(p, l, r)$, where $p$ is a string and $l$ and $r$ $(l \leq r)$ are integers.

## Rule Definition

We'll say that string $t$ complies with rule $(p, l, r)$, if the number of occurrences of string $t$ in string $p$ lies between $l$ and $r$, inclusive.

**Examples:**
- String "ab" complies with rules ("ab", 1, 2) and ("aab", 0, 1)
- String "ab" does not comply with rules ("cd", 1, 2) and ("abab", 0, 1)

## Substring Definition

A substring $s[l...r]$ $(1 \leq l \leq r \leq |s|)$ of string $s = s_1s_2...s_{|s|}$ $(|s|$ is the length of $s)$ is string $s_ls_{l+1}...s_r$.

Consider a number of occurrences of string $t$ in string $p$ as a number of pairs of integers $l, r$ $(1 \leq l \leq r \leq |p|)$ such that $p[l...r] = t$.

## Good String Definition

We'll say that string $t$ is **good** if it complies with all $n$ rules. Smart Beaver asks you to help him to write a program that can calculate the number of distinct good substrings of string $s$.

Two substrings $s[x...y]$ and $s[z...w]$ are considered to be distinct if and only if $s[x...y] \neq s[z...w]$.

## Input

The first line contains string $s$.

The second line contains integer $n$.

Next $n$ lines contain the rules, one per line. Each of these lines contains a string and two integers $p_i, l_i, r_i$, separated by single spaces $(0 \leq l_i \leq r_i \leq |p_i|)$.

It is guaranteed that all the given strings are non-empty and only contain lowercase English letters.

## Scoring System

### Subproblem G1 (30 points):
- $0 \leq n \leq 10$
- The length of string $s$ and the maximum length of string $p$ is $\leq 200$

### Subproblem G2 (70 points):
- $0 \leq n \leq 10$
- The length of string $s$ and the maximum length of string $p$ is $\leq 2000$

### Subproblem G3 (100 points):
- $0 \leq n \leq 10$
- The length of string $s$ and the maximum length of string $p$ is $\leq 50000$

## Output

Print a single integer — the number of good substrings of string $s$.

## Examples

### Example 1

**Input:**
```
aaab
2
aa 0 0
aab 1 1
```

**Output:**
```
3
```

### Example 2

**Input:**
```
ltntlnen
3
n 0 0
ttlneenl 1 4
lelllt 1 1
```

**Output:**
```
2
```

### Example 3

**Input:**
```
a
0
```

**Output:**
```
1
```

## Note

There are three good substrings in the first sample test: «aab», «ab» and «b».

In the second test only substrings «e» and «t» are good.


### ideas
1. len(s) <= 50000, len(p) <= 50000
2. 最多10条规则
3. 要在s中找出满足n条规则的t(substring), 计算它们的数量
4. 将n条规则变成一个suffix tree？（10的话，似乎不一定用得到？）
5. 对于规则p, l, r 是不是可以知道满足条件的的子字符串？ 要计算某一段t，在其中的数量，似乎也挺难的？
6. 如果知道t，那么把t作为pattern，生成kmp，然后去运行p，这样可以在O(p)的时间内，计算出t的出现次数
7. 但是没法使用所有的子串去运行这个逻辑
8. 将s变成index tree， 然后去匹配p？