# Problem Statement

There is a board with a grid consisting of $n$ rows and $m$ columns, the rows are numbered from 1 from top to bottom and the columns are numbered from 1 from left to right. In this grid we will denote the cell that lies on row number $i$ and column number $j$ as $(i, j)$.

A group of six numbers $(a, b, c, d, x_0, y_0)$, where $0 \leq a, b, c, d$, is a cross, and there is a set of cells that are assigned to it. Cell $(x, y)$ belongs to this set if at least one of two conditions are fulfilled:

- $|x_0 - x| \leq a$ and $|y_0 - y| \leq b$
- $|x_0 - x| \leq c$ and $|y_0 - y| \leq d$

The picture shows the cross $(0, 1, 1, 0, 2, 3)$ on the grid $3 \times 4$.

Your task is to find the number of different groups of six numbers, $(a, b, c, d, x_0, y_0)$ that determine the crosses of an area equal to $s$, which are placed entirely on the grid. The cross is placed entirely on the grid, if any of its cells is in the range of the grid (that is for each cell $(x, y)$ of the cross $1 \leq x \leq n; 1 \leq y \leq m$ holds). The area of the cross is the number of cells it has.

Note that two crosses are considered distinct if the ordered groups of six numbers that denote them are distinct, even if these crosses coincide as sets of points.

---

## Input

The input consists of a single line containing three integers $n$, $m$ and $s$ ($1 \leq n, m \leq 500$, $1 \leq s \leq n \cdot m$). The integers are separated by a space.

## Output

Print a single integer — the number of distinct groups of six integers that denote crosses with area $s$ and that are fully placed on the $n \times m$ grid.

> Please, do not use the %lld specifier to read or write 64-bit integers in C++. It is preferred to use the cin, cout streams or the %I64d specifier.

### ideas
1. 一共是6个维度。
2. 对于一个cross，可以认为有3部分组成，中间部分，上下，左右
3. 假设中间部分的长宽是(h, w)
4. 那么有 h * w +  w * (th - h) + h * (tw - w) = s
5. th上下的宽度， tw是左右的宽度
6. 且必须满足h,w是奇数（这样子才有中心)
7. 那么这样子就是4个维度，但是第4个维度，可以由其他三个维度确定