# Problem Description

There is an n × m rectangular grid, each cell of the grid contains a single integer: zero or one. Let's call the cell on the i-th row and the j-th column as (i, j).

Let's define a "rectangle" as four integers a, b, c, d (1 ≤ a ≤ c ≤ n; 1 ≤ b ≤ d ≤ m). Rectangle denotes a set of cells of the grid {(x, y) : a ≤ x ≤ c, b ≤ y ≤ d}. Let's define a "good rectangle" as a rectangle that includes only the cells with zeros.

You should answer the following q queries: calculate the number of good rectangles all of which cells are in the given rectangle.

## Input

There are three integers in the first line: n, m and q (1 ≤ n, m ≤ 40, 1 ≤ q ≤ 3·10^5). Each of the next n lines contains m characters — the grid. Consider grid rows are numbered from top to bottom, and grid columns are numbered from left to right. Both columns and rows are numbered starting from 1.

Each of the next q lines contains a query — four integers that describe the current rectangle, a, b, c, d (1 ≤ a ≤ c ≤ n; 1 ≤ b ≤ d ≤ m).

## Output

For each query output an answer — a single integer in a separate line.

## Examples

### Example 1

**Input:**
```
5 5 5
00101
00000
00001
01000
00001
1 2 2 4
4 5 4 5
1 2 5 2
2 2 4 5
4 2 5 3
```

**Output:**
```
10
1
7
34
5
```

### Example 2

**Input:**
```
4 7 5
0000100
0000010
0011000
0000000
1 7 2 7
3 1 3 1
2 3 4 5
1 2 2 7
2 2 4 7
```

**Output:**
```
3
1
16
27
52
```


### ideas
1. n <= 40, m <= 40. 如果一开始就生成好, n * n * m * m = 40 * 40 * 40 * 40 = 1e4 * 256
2. 时间上似乎也ok？
3. dp[a][b][c][d] 表示从(a, b)开始，右下角在(x, y), x <= c && y <= d的good的区间
4. 查询的时候，放过来，固定(c, d)迭代(a, b)就可以了