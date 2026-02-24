# Problem D

Harry Potter has a difficult homework. Given a rectangular table consisting of $n \times m$ cells. Each cell contains an integer. Harry knows two spells: the first spell changes the sign of the integers in a selected row; the second changes the sign in a selected column. Harry's task is to make the sum of the numbers in each row and each column non-negative using these spells.

Alone, the boy cannot cope. Help the young magician!

---

## Input

- First line: two integers $n$ and $m$ ($1 \le n, m \le 100$) — the number of rows and columns.
- Next $n$ lines: each contains $m$ integers. The $j$-th integer in the $i$-th line is $a_{i,j}$ ($|a_{i,j}| \le 100$), the number in row $i$ and column $j$.
- Rows are numbered from $1$ to $n$, columns from $1$ to $m$.

## Output

- **First line:** Print $a$ — the number of applications of the first spell — then $a$ distinct space-separated row numbers to apply the spell to.
- **Second line:** Print $b$ — the number of applications of the second spell — then $b$ distinct space-separated column numbers to apply the spell to.

If there are several solutions, output any of them.
