### Problem

You are given `n` axis-aligned rectangles labeled from `1` to `n`.  
Each rectangle has integer coordinates and its edges are parallel to the axes.  
Rectangles may **touch** each other (share borders or corners), but they do **not overlap** in their interiors (no point belongs to the interior of more than one rectangle).

Your task is to determine whether there exists a **non-empty subset** of these rectangles that exactly forms a **square**.

Formally, we ask if there exists:

- a subset of the rectangles, and  
- a square,

such that:

1. Every point belonging to the **interior or border** of the square belongs to the interior or border of **at least one** rectangle in the subset; and  
2. Every point belonging to the **interior or border** of at least one rectangle in the subset belongs to the interior or border of that square.

In other words, the union of the chosen rectangles is **exactly** a square (including its boundary).

### Input

- First line: integer `n` (`1 ≤ n ≤ 10^5`) — the number of rectangles.
- Next `n` lines: each describes rectangle `i` with four integers:
  - `x1 y1 x2 y2` — the bottom-left corner is `(x1, y1)` and the top-right corner is `(x2, y2)`,  
    with `0 ≤ x1 < x2 ≤ 3000`, `0 ≤ y1 < y2 ≤ 3000`.

Additional guarantee:

- No two rectangles overlap in their interiors.

### Output

- If such a subset exists:
  - Print `YES` (without quotes) and a space, then `k` — the number of rectangles in the subset.
  - On the next line, print `k` integers — the labels of rectangles in the subset in any order.  
    If multiple valid subsets exist, you may print **any** of them.
- If no such subset exists, print `NO` (without quotes).

### Examples

**Input**
```text
9
0 0 1 9
1 0 9 1
1 8 9 9
8 1 9 8
2 2 3 6
3 2 7 3
2 6 7 7
5 3 7 6
3 3 5 6
```

**Output**
```text
YES 5
5 6 7 8 9
```

**Input**
```text
4
0 0 1 9
1 0 9 1
1 8 9 9
8 1 9 8
```

**Output**
```text
NO
```

### ideas
1. 需要判断是否存在一个子集能组成一个正方形
2. x <= 3000, 所以平面上一共有3000 * 3000个点
3. 假设dp[i][j] 表示被(0, 0, i, j)完全覆盖的点的数量
4. 如果 dp[i][j] - dp[i - w][j] - dp[i][j-w] + dp[i-w][j-w] = w * w 
5. 那么这个w区域满足条件
6. (dp[i][j] 应该可以在O(n) + O(p))的时间算出来
7. w <= 3000, 在每个点的右上角开始算