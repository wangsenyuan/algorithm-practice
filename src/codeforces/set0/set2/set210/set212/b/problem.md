# Problem B

For a string $s = s_1 s_2 \ldots s_{|s|}$ (where $|s|$ is the length of $s$), a **substring** $s[a,b]$ is $s_a s_{a+1} \ldots s_b$ with $1 \le a \le b \le |s|$.

The **trace** of a non-empty string $t$ is the set of distinct characters in $t$. For example, the trace of `"aab"` is $\{\texttt{a}, \texttt{b}\}$.

Consider a string $s$ and the set of its substrings whose trace equals some set $C$. Denote by $r(C, s)$ the number of substrings in this set that are **maximal by inclusion**: a substring $s[a,b]$ of length $n = b - a + 1$ in this set is maximal if there is no other substring $s[x,y]$ in the set with length $> n$ such that $1 \le x \le a \le b \le y \le |s|$. Two substrings are different if they have different positions in $s$, even if their content is the same.

Polycarpus must solve: given string $s$ and non-empty character sets $C_1, C_2, \ldots, C_m$, find $r(C_i, s)$ for each $C_i$. Help him so he doesn’t get expelled.

---

## Input

- First line: a non-empty string $s$ ($1 \le |s| \le 10^6$).
- Second line: integer $m$ ($1 \le m \le 10^4$).
- Next $m$ lines: the $i$-th line is a string $c_i$ whose **trace** is $C_i$ (all characters in $c_i$ are distinct). The sets $C_i$ may repeat. All strings use lowercase English letters.

## Output

Print $m$ integers: the $i$-th integer is $r(C_i, s)$.

---

## Examples

### Example 1

**Input:**

```
aaaaa
2
a
a
```

**Output:**

```
1
1
```

### Example 2

**Input:**

```
abacaba
3
ac
ba
a
```

**Output:**

```
1
2
4
```

### ideas
1. 阅读理解～
2. 对于给定的C，找出最长的子串w, 由C组成，且如果存在两端外的字符，它们不能属于C
3. 计数
4. 假设s[l...r]的trace = C, 如果 s[l-1] 属于C, 那么可以扩张r(C)
5. 或者这里trace[l...r] = C, 那么 trace[l-1...r] != C
6. 这样子，似乎只需要更新26次（因为26次后，肯定包含了完整的字符集）
7. 也就是说每次l...r = C要找到左边最远的某个l1, s[l1....r] 比C多最少的字符（最好是1个）
8. abC
9. brute force = s[l-1] 肯定不在C里面，然后加入l-1, 然后一直找最左边第一个不在C中的字符， 它是新的边界
10. 这个时候，只需要迭代C，假设w不在其中，找出它的最近的位置（就是最新的边界）
11. 搞定