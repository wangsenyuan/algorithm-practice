# Problem

**题意简述**：给定一个 `n x m` 的黑白网格，至少有一个黑格。若任意两个黑格之间，都存在一条仅经过上下左右相邻黑格的路径，且这条路径最多只拐弯一次，则称该网格是凸的。判断给定网格是否为凸网格。

## Input

- The first line contains two integers `n` and `m` (`1 ≤ n, m ≤ 50`) — the size of the grid.
- Each of the next `n` lines contains `m` characters `"B"` or `"W"`.
- Character `"B"` denotes a black cell and `"W"` denotes a white cell.
- It is guaranteed that the grid has at least one black cell.

## Output

Print `"YES"` if the grid is convex, otherwise print `"NO"`.

## Examples

### Example 1

**Input**

```text
3 4
WWBW
BWWW
WWWB
```

**Output**

```text
NO
```

### Example 2

**Input**

```text
3 1
B
B
W
```

**Output**

```text
YES
```
