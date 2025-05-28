# Problem Statement

It is known that fleas in Berland can jump only vertically and horizontally, and the length of the jump is always equal to $s$ centimeters. A flea has found herself at the center of some cell of the checked board of the size $n \times m$ centimeters (each cell is $1 \times 1$ centimeters). She can jump as she wishes for an arbitrary number of times, she can even visit a cell more than once. The only restriction is that she cannot jump out of the board.

The flea can count the amount of cells that she can reach from the starting position $(x, y)$. Let's denote this amount by $d_{x, y}$. Your task is to find the number of such starting positions $(x, y)$, which have the maximum possible value of $d_{x, y}$.

## Input

The first line contains three integers $n$, $m$, $s$ ($1 \leq n, m, s \leq 10^6$) — length of the board, width of the board and length of the flea's jump.

## Output

Output the only integer — the number of the required starting positions of the flea.

---

### Ideas

1. 考虑水平方向，`c % s` 相同的列，可以被覆盖到。