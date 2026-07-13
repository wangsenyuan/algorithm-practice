# C - Belt Conveyor

Problem: [AtCoder Beginner Contest 265 C](https://atcoder.jp/contests/abc265/tasks/abc265_c)

## Problem Statement

There is a grid with `H` rows and `W` columns. Each cell contains one of the characters `U`, `D`, `L`, or `R`.

You begin at cell `(1, 1)`. At every step, follow the direction written in your current cell:

- `U`: move one row upward.
- `D`: move one row downward.
- `L`: move one column to the left.
- `R`: move one column to the right.

If the indicated move would leave the grid, stop and print your current row and column. If you continue moving forever, print `-1`.

## Constraints

- `1 <= H, W <= 500`
- Every cell contains `U`, `D`, `L`, or `R`.
- `H` and `W` are integers.

## Input

```text
H W
G_1
G_2
...
G_H
```

Here, `G_i` is a string of length `W` describing row `i`.

## Output

If movement stops at `(i, j)`, print:

```text
i j
```

If movement never stops, print:

```text
-1
```

## Sample 1

```text
Input
2 3
RDU
LRU

Output
1 3
```

## Sample 2

```text
Input
2 3
RRD
ULL

Output
-1
```

## Sample 3

```text
Input
9 44
RRDDDDRRRDDDRRRRRRDDDRDDDDRDDRDDDDDDRRDRRRRR
RRRDLRDRDLLLLRDRRLLLDDRDLLLRDDDLLLDRRLLLLLDD
DRDLRLDRDLRDRLDRLRDDLDDLRDRLDRLDDRLRRLRRRDRR
DDLRRDLDDLDDRLDDLDRDDRDDDDRLRRLRDDRRRLDRDRDD
RDLRRDLRDLLLLRRDLRDRRDRRRDLRDDLLLLDDDLLLLRDR
RDLLLLLRDLRDRLDDLDDRDRRDRLDRRRLDDDLDDDRDDLDR
RDLRRDLDDLRDRLRDLDDDLDDRLDRDRDLDRDLDDLRRDLRR
RDLDRRLDRLLLLDRDRLLLRDDLLLLLRDRLLLRRRRLLLDDR
RRRRDRDDRRRDDRDDDRRRDRDRDRDRRRRRRDDDRDDDDRRR

Output
9 5
```

## Solution

Simulate the movement from `(1, 1)` while recording every visited cell.

The grid uses 0-based coordinates in the code, so the initial position is `(x, y) = (0, 0)`. A two-dimensional boolean array `vis` records whether each cell has already been visited.

At every iteration:

1. If `(x, y)` is already marked in `vis`, return `-1`.
2. Mark `(x, y)` as visited.
3. Read the direction in `grid[x][y]`.
4. If moving in that direction would leave the grid, the movement stops at the current cell. Return `(x + 1, y + 1)` to convert the coordinates back to 1-based indexing.
5. Otherwise, update `(x, y)` to the next cell and continue.

### Why revisiting a cell means an infinite loop

The direction written in a cell never changes. Therefore, whenever we are at a particular cell, the next action is always the same.

If we visit a cell for the second time, the movement from that point onward is identical to what happened after the first visit. The same sequence of cells will repeat forever, so the answer is `-1`.

If no cell is revisited, the simulation can visit at most `H * W` cells. Thus, it must eventually reach a cell whose direction points outside the grid, and that cell is the answer.

### Correctness Proof

We prove that the algorithm returns the required answer.

During the simulation, `(x, y)` is exactly the current position because the algorithm applies the direction written in the current cell according to the movement rules.

- If the indicated move goes outside the grid, no further move is possible. The algorithm returns the current cell using 1-based coordinates, which is the required final position.
- If the algorithm reaches an already visited cell, that cell has the same fixed outgoing direction as on its previous visit. From then on, every following position also repeats, so movement never ends. The algorithm correctly returns `-1`.

These are the only two ways the simulation can end. Therefore, the algorithm always returns the correct answer.

### Complexity

Each cell is processed at most once before the algorithm stops or detects a repeated cell.

- Time: `O(H * W)`
- Space: `O(H * W)` for the visited array
