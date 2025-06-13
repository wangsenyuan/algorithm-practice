# Problem Statement

We've got a rectangular $n \times m$-cell maze. Each cell is either passable, or is a wall (impassable). A little boy found the maze and cyclically tiled a plane with it so that the plane became an infinite maze. Now on this plane, cell $(x, y)$ is a wall if and only if cell $(x \bmod n, y \bmod m)$ is a wall.

In this problem, $a \bmod b$ is the remainder of dividing number $a$ by number $b$.

The little boy stood at some cell on the plane and he wondered whether he can walk infinitely far away from his starting position. From cell $(x, y)$ he can go to one of the following cells: $(x, y-1)$, $(x, y+1)$, $(x-1, y)$ and $(x+1, y)$, provided that the cell he goes to is not a wall.

---

## Input

The first line contains two space-separated integers $n$ and $m$ ($1 \leq n, m \leq 1500$) — the height and the width of the maze that the boy used to cyclically tile the plane.

Each of the next $n$ lines contains $m$ characters — the description of the labyrinth. Each character is either a `#`, that marks a wall, a `.`, that marks a passable cell, or an `S`, that marks the little boy's starting point.

The starting point is a passable cell. It is guaranteed that character `S` occurs exactly once in the input.

---

## Output

Print `Yes` (without the quotes), if the little boy can walk infinitely far from the starting point. Otherwise, print `No` (without the quotes).

## Ideas
1. 先得到边界。
2. 然后下边界要和上边界先连，左和右要相连
3. 如果能够从上到下，或者从左到右，那么就可以