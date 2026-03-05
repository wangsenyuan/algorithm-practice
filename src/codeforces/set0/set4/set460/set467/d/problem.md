# Problem

**题意：** 给出一篇由 $m$ 个单词组成的文章，以及 $n$ 条同义词规则（每条形如「单词 $x_i$ 可替换为 $y_i$」，单向）。可将文章中任意单词按规则替换为其同义词，次数任意。字母大小写不区分。求在保证替换后语义不变的前提下：先使文章中字母 R（不区分大小写）的个数最少，若有多种则再使文章总长度（所有单词长度之和）最短。输出这两个值：最少 R 的个数、以及该情况下的最短总长度。

## Input

- The first line contains a single integer $m$ ($1 \le m \le 10^5$) — the number of words in the initial essay.
- The second line contains the words of the essay, separated by a single space. The total length of all words is at most $10^5$ characters.
- The next line contains a single integer $n$ ($0 \le n \le 10^5$) — the number of synonym pairs in the dictionary.
- Each of the next $n$ lines contains two space-separated non-empty words $x_i$ and $y_i$, meaning that word $x_i$ can be replaced with word $y_i$ (but not vice versa). The total length of all words in the dictionary is at most $5 \cdot 10^5$ characters.

All words consist only of uppercase and lowercase English letters.

## Output

Print two integers: the minimum number of letters R in an optimal essay, and the minimum total length of such an essay.

## Examples

### Example 1

**Input:**

```
3
AbRb r Zz
4
xR abRb
aA xr
zz Z
xr y
```

**Output:**

```
2 6
```

### Example 2

**Input:**

```
2
RuruRu fedya
1
ruruRU fedor
```

**Output:**

```
1 10
```
