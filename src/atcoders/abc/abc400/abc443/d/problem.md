There is an $N \times N$ grid, and there is one piece placed in each column. The piece in column $i$ is placed in row $R_i$ from the top.

You can perform the following operation zero or more times:

- Choose a piece that is not in the topmost row and move that piece to the cell directly above it.

Find the minimum number of operations needed to satisfy the following condition for all integers $i$ satisfying $1 \le i \le N-1$:

Let the piece in column $i$ be in row $x$ from the top and the piece in column $i+1$ be in row $y$ from the top. Then, $|x-y| \le 1$.

You are given $T$ test cases; solve each of them.

## Constraints

All input values are integers.

- $1 \le T \le 50000$
- $2 \le N \le 3 \times 10^5$
- $1 \le R_i \le N$

For a single input, the sum of $N$ does not exceed $3 \times 10^5$.

## Input

The input is given from Standard Input in the following format:

```
T
case 1
case 2
⋮
case T
```

Each test case is given in the following format:

```
N
R_1 R_2 ... R_N
```

## Output

Output $T$ lines.

The $i$-th line should contain the answer for the $i$-th test case.

## Sample Input 1

```
5
5
5 2 1 3 4
2
1 1
3
1 3 1
9
9 9 8 2 4 4 3 5 3
20
7 4 6 2 15 5 17 15 1 8 18 1 5 1 12 11 2 7 8 14
```

## Sample Output 1

```
4
0
1
16
105
```

This input contains five test cases.

For the first test case, by performing operations as follows, you can satisfy the condition in the problem statement with four operations, which is the minimum.

Move the piece in the 5-th column from the left to the cell directly above it. The pieces are in rows $5, 2, 1, 3, 3$ from left to right.

Move the piece in the 1-st column from the left to the cell directly above it. The pieces are in rows $4, 2, 1, 3, 3$ from left to right.

Move the piece in the 1-st column from the left to the cell directly above it. The pieces are in rows $3, 2, 1, 3, 3$ from left to right.

Move the piece in the 4-th column from the left to the cell directly above it. The pieces are in rows $3, 2, 1, 2, 3$ from left to right.

For the second test case, the condition is satisfied without performing any operations.


### ideas
1. 只能往上移动
2. 已经在top的肯定不需要移动。按照高度处理
3. 现在处理到i, 如果左边比它高的位置是l, 右边比它高的位置是r（有可能不存在）
4. 那么R[i] 要满足 R[r] - x <= r - i => x >= R[r] - (r - i)
5. 且 x >= R[l] - (i - l) 所以找到x >= max(a, b)
6. 那么这个x - R[i] = 就是结果