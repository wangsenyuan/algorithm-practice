# F - No Passage (ABC413)

**Contest:** [ABC413](https://atcoder.jp/contests/abc413) — Denso Create Programming Contest 2025 (AtCoder Beginner Contest 413)  
**Task:** [https://atcoder.jp/contests/abc413/tasks/abc413_f](https://atcoder.jp/contests/abc413/tasks/abc413_f)

**Time limit:** 2.5 sec / **Memory limit:** 1024 MiB  
**Score:** 525 points

## Problem Statement

There is an `H x W` grid. Let `(i,j)` denote the cell at the `i`-th row from the
top and `j`-th column from the left. Among these, `K` cells are goals. The `i`-th
goal (`1 <= i <= K`) is cell `(R_i, C_i)`.

Takahashi and Aoki play a game using this grid and a single piece placed on the
grid. Takahashi and Aoki repeatedly perform the following series of operations until
the piece reaches a goal cell:

1. Aoki chooses an integer `a` between `1` and `4`, inclusive.
2. Then, Takahashi chooses an integer `b` between `1` and `4`, inclusive, where
   `a != b` must be satisfied. Let `(i,j)` be the cell where the piece is placed
   before the operation. Based on the chosen integer `b` and the piece's position,
   move the piece:
   - When `b = 1`: If `(i-1,j)` is within the grid, move the piece from cell
     `(i,j)` to cell `(i-1,j)`; if it is outside the grid, do nothing.
   - When `b = 2`: If `(i+1,j)` is within the grid, move the piece from cell
     `(i,j)` to cell `(i+1,j)`; if it is outside the grid, do nothing.
   - When `b = 3`: If `(i,j-1)` is within the grid, move the piece from cell
     `(i,j)` to cell `(i,j-1)`; if it is outside the grid, do nothing.
   - When `b = 4`: If `(i,j+1)` is within the grid, move the piece from cell
     `(i,j)` to cell `(i,j+1)`; if it is outside the grid, do nothing.

Takahashi's objective is to minimize the number of moves until the piece reaches a
goal. Aoki's objective is to prevent the piece from reaching the goal; if that is
impossible, his objective is to maximize the number of moves until the piece reaches
a goal.

For all pairs of integers `(i,j)` satisfying `1 <= i <= H` and `1 <= j <= W`, solve
the following problem and find the sum of all solutions:

Start the game with the piece at cell `(i,j)`. Assume both players act optimally
toward their respective objectives. If Takahashi can make the piece reach a goal,
the solution is the minimum number of moves; otherwise, the solution is `0`.

## Constraints

- `2 <= H <= 3000`
- `2 <= W <= 3000`
- `1 <= K <= min(HW, 3000)`
- `1 <= R_i <= H`
- `1 <= C_i <= W`
- `(R_i, C_i) != (R_j, C_j)` for `1 <= i < j <= K`
- All input values are integers.

## Input

The input is given from standard input in the following format:

```text
H W K
R_1 C_1
R_2 C_2
:
R_K C_K
```

## Output

Print the answer.

## Sample Input 1

```text
2 3 2
1 2
2 1
```

## Sample Output 1

```text
2
```

When `(i,j) = (1,2), (2,1)`, the starting cell is a goal, so the solution is `0`.

When `(i,j) = (1,1), (2,2)`, no matter which `a` Aoki chooses, Takahashi can make
the piece reach a goal in `1` move from the starting cell, so the solution is `1`.

When `(i,j) = (1,3), (2,3)`, Takahashi cannot reach a goal, so the solution is `0`.

The sum of these is `0 * 2 + 1 * 2 + 0 * 2 = 2`. Thus, print `2`.

## Sample Input 2

```text
9 3 9
1 3
6 1
4 1
1 2
2 1
7 1
9 3
8 1
9 2
```

## Sample Output 2

```text
43
```

## Sample Input 3

```text
10 10 36
3 8
5 10
3 10
6 10
2 10
2 8
7 10
1 10
1 8
7 6
7 8
2 5
1 6
8 8
7 5
2 4
9 8
7 4
4 3
10 10
10 8
8 10
10 6
6 2
4 2
10 5
8 3
1 2
2 1
4 1
10 4
10 3
8 1
6 1
10 2
9 1
```

## Sample Output 3

```text
153
```


### ideas
1. A 可以阻止T朝向一个方向运动
2. dp[i][j]表示从goals移动到(i, j)，在有阻挡的情况下的最优解
3. (当(i, j))被第二次touch到的时候，可以被计算