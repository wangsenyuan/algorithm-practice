# Digital Table Expansion

## Problem Description

Mr. Bender has a digital table of size n × n, where each cell can be switched on or off. He wants the field to have at least c switched-on squares to be happy.

### Setup
- The table has rows numbered 1 to n (top to bottom) and columns numbered 1 to n (left to right)
- Initially, exactly one cell at coordinates (x, y) is switched on; all other cells are off
- Each second, cells that are adjacent to already-on cells switch on

### Expansion Rule
For a cell at coordinates (x, y), the side-adjacent cells are:
- (x - 1, y) — above
- (x + 1, y) — below
- (x, y - 1) — left
- (x, y + 1) — right

The switched-on region expands in a diamond/Manhattan distance pattern.

### Task
Determine in how many seconds Mr. Bender's condition (at least c cells switched on) will be satisfied.

## Input Format
A single line containing four space-separated integers:
- `n`: table size (1 ≤ n ≤ 10⁹)
- `x`: initial row of switched-on cell (1 ≤ x ≤ n)
- `y`: initial column of switched-on cell (1 ≤ y ≤ n)
- `c`: minimum number of switched-on cells needed (1 ≤ c ≤ 10⁹; c ≤ n²)

## Output Format
A single integer — the minimum number of seconds needed.

## Examples

### Example 1
**Input:**
```
6 4 3 1
```
**Output:**
```
0
```

### Example 2
**Input:**
```
9 3 8 10
```
**Output:**
```
2
```

### ideas
1. 如果在一个无限的平面上，经过w秒后，会有 x = 2 * w长的边， 4 * x - 4 个新的格子
2. 8 * w - 4
3. 但是现在是在n * n 的平面内
4. 感觉要用二分(在w秒的时候)，有多少格子跑到了外面去
5. 思路应该没有错。但是处理起来，似乎还有点难