# Problem Statement

Dreamoon has a string $s$ and a pattern string $p$. He first removes exactly $x$ characters from $s$ obtaining string $s'$ as a result. Then he calculates $f(s')$ that is defined as the maximal number of non-overlapping substrings equal to $p$ that can be found in $s'$. He wants to make this number as big as possible.

More formally, let's define $g(x)$ as maximum value of $f(s')$ over all $s'$ that can be obtained by removing exactly $x$ characters from $s$. Dreamoon wants to know $g(x)$ for all $x$ from $0$ to $|s|$ where $|s|$ denotes the length of string $s$.

## Input

The first line of the input contains the string $s$ ($1 \leq |s| \leq 2000$).

The second line of the input contains the string $p$ ($1 \leq |p| \leq 500$).

Both strings will only consist of lower case English letters.

## Output

Print $|s| + 1$ space-separated integers in a single line representing the $g(x)$ for all $x$ from $0$ to $|s|$.

## Examples

### Example 1
**Input:**
```
aaaaa
aa
```

**Output:**
```
2 2 1 1 0 0
```

### Example 2
**Input:**
```
axbaxxb
ab
```

**Output:**
```
0 1 1 2 1 1 0 0
```


### ideas
1. 对于给定的x，dp[i][x]表示目前s[:i]前缀删除x个字符后能够匹配的p[:j]的最大前缀
2. dp[i+1][x] = 根据lps计算的新的位置，（不删除s[i+1]), 或者 dp[i][x-1] (删除s[i+1])
3. 这里没法知道个数，但是如果将dp[i][x]表示成 匹配p的重复的位置，似乎也就知道这个位置了
4. got