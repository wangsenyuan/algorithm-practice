# Problem Description

You are running through a rectangular field. This field can be represented as a matrix with 3 rows and m columns. (i, j) denotes a cell belonging to i-th row and j-th column.

You start in (2, 1) and have to end your path in (2, m). From the cell (i, j) you may advance to:

- (i - 1, j + 1) — only if i > 1
- (i, j + 1), or
- (i + 1, j + 1) — only if i < 3

However, there are n obstacles blocking your path. k-th obstacle is denoted by three integers ak, lk and rk, and it forbids entering any cell (ak, j) such that lk ≤ j ≤ rk.

You have to calculate the number of different paths from (2, 1) to (2, m), and print it modulo 10^9 + 7.

## Input

The first line contains two integers n and m (1 ≤ n ≤ 10^4, 3 ≤ m ≤ 10^18) — the number of obstacles and the number of columns in the matrix, respectively.

Then n lines follow, each containing three integers ak, lk and rk (1 ≤ ak ≤ 3, 2 ≤ lk ≤ rk ≤ m - 1) denoting an obstacle blocking every cell (ak, j) such that lk ≤ j ≤ rk. Some cells may be blocked by multiple obstacles.

## Output

Print the number of different paths from (2, 1) to (2, m), taken modulo 10^9 + 7. If it is impossible to get from (2, 1) to (2, m), then the number of paths is 0.

## Example

**Input:**
```
2 5
1 3 4
2 2 3
```

**Output:**
```
2
```

### ideas
1. 如果没有障碍物，貌似是一个matrix的pow问题
2. 一个3 * 3 的矩阵, (i, j) 表示从i行到第j行的转移, (0, 0) = 1, (0, 1) = 1, (0, 2) = 0, (1, 0) = 1, (1, 1) = 1, (1, 2) = 1, ...
3. 对于障碍物，那么先计算到它前一列，然后看最后一列～
4. 每个障碍物都是一条的，如果有两个障碍物上下重叠了，那么就阻断了这两行的移动
5. 所以，还是的分情况，阻断第一行，阻断第二行，阻断第三行，。。。
6. 不过确实可以搞了