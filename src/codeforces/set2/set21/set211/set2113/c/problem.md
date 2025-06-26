# Problem Description

The boy Smilo is playing Minecraft! To prepare for the battle with the dragon, he needs a lot of golden apples, and for that, he requires a lot of gold. Therefore, Smilo goes to the mine.

The mine is a rectangular grid of size $n \times m$, where each cell can be either gold ore, stone, or an empty cell. Smilo can blow up dynamite in any empty cell. When dynamite explodes in an empty cell with coordinates $(x, y)$, all cells within a square of side $2k + 1$ centered at cell $(x, y)$ become empty. If gold ore was located strictly inside this square (not on the boundary), it disappears. However, if the gold ore was on the boundary of this square, Smilo collects that gold.

Dynamite can only be detonated inside the mine, but the explosion square can extend beyond the mine's boundaries.

Determine the maximum amount of gold that Smilo can collect.

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ $(1 \leq t \leq 10^4)$. The description of the test cases follows.

The first line of each test case contains three integers $n$, $m$, and $k$ $(1 \leq n, m, k \leq 500)$ — the number of rows, columns, and the explosion parameter $k$, respectively.

Each of the following $n$ lines contains $m$ characters, each of which is equal to:
- `'.'` — empty cell
- `'#'` — stone  
- `'g'` — gold

It is guaranteed that at least one of the cells is empty.

It is guaranteed that the sum $n \cdot m$ across all test cases does not exceed $2.5 \cdot 10^5$.

## Output

For each test case, output a single integer — the maximum amount of gold that can be obtained.

## ideas
1. 第一个起爆点之后，可以逐步扩大， 且之后都不再损失
2. 所以答案 = all - 角落内的最少的gold
3. 当然有个问题是，这个角落里面必须要是empty