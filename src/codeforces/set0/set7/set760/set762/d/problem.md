# Problem D: Maximum Path Sum

## Problem Description

You are given a rectangular table of size 3 × n. Each cell contains an integer. You can move from one cell to another if they share a side.

Find a path from the upper left cell to the bottom right cell of the table that:
- Doesn't visit any cell twice
- Has the maximum possible sum of numbers written in the cells

## Input

- **First line**: An integer n (1 ≤ n ≤ 10^5) — the number of columns in the table
- **Next three lines**: Each line contains n integers — the description of the table
- The j-th number in the i-th line corresponds to the cell a[i][j] (-10^9 ≤ a[i][j] ≤ 10^9)

## Output

Output the maximum sum of numbers on a valid path from the upper left cell to the bottom right cell.

## Examples

### Example 1
**Input:**
```
3
1 1 1
1 -1 1
1 1 1
```

**Output:**
```
7
```

### Example 2
**Input:**
```
5
10 10 10 -1 -1
-1 10 10 10 10
-1 10 10 10 10
```

**Output:**
```
110
```


### ideas
1. 3行的情况下，似乎是存在*往回走的路*的
2. 但是对于一个(i, j) 如果它从右边过来，那么它的上边/下边，必须有能退出去的通道
3. 折返的只能是在中间行，且看起来是只需要完后折返一格
4. 考虑折返两格的情况，从第一行下来，然后折返两格，然后从第三行返回
5. 正好组成一个 *田*; 在左上角进入，但是在右下角离开
6. 那么通过调整路径，可以变成先访问第一列，再访问第二列，再访问第3列，达到同样的效果
7. 但是如果是个*日*型的结构，要从左上进入，从右下离开，就必须在中间行折返