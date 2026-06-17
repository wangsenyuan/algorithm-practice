# D - Repeatedly Repainting (ABC460)

**Contest:** [ABC460](https://atcoder.jp/contests/abc460) — AtCoder Beginner Contest 460  
**Task:** [https://atcoder.jp/contests/abc460/tasks/abc460_d](https://atcoder.jp/contests/abc460/tasks/abc460_d)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 425 points

## Problem Statement

There is an `H` by `W` grid. Cell `(i, j)` is the cell in row `i` (from the top)
and column `j` (from the left).

Each cell is white or black. The grid is given by `H` strings `S_1, …, S_H` of
length `W`:

- `.` — cell `(i, j)` is white
- `#` — cell `(i, j)` is black

Perform the following operation **10^100** times. In each step, update **all**
cells simultaneously:

1. A **white** cell becomes **black** if and only if it has at least one **black**
   **8-neighbor** (Chebyshev distance 1: `max(|x - x'|, |y - y'|) = 1`).
2. A **black** cell becomes **white**.

Output the final color of every cell.

## Constraints

- `1 <= H * W <= 10^6`
- `H`, `W` are positive integers
- Each `S_i` is a string of length `W` consisting of `.` and `#`

## Input

```text
H W
S_1
S_2
⋮
S_H
```

## Output

Print `H` lines. The `i`-th line is a string of length `W`:

- `.` if cell `(i, j)` is white after all operations
- `#` if cell `(i, j)` is black after all operations

## Sample Input 1

```text
3 4
#.#.
.#..
#...
```

## Sample Output 1

```text
#.#.
.#..
#..#
```

After one operation the grid changes; after `10^100` operations it reaches the
output above (see the statement figures for the intermediate grids).

## Sample Input 2

```text
3 3
###
###
###
```

## Sample Output 2

```text
...
...
...
```

## Sample Input 3

```text
5 7
.#.....
.......
..#....
.......
....#..
```

## Sample Output 3

```text
.#.##.#
....#..
#.#.###
#.....#
###.#.#
```

## ideas
1. 一开始是白色的格子, 需要计算出,它在什么时刻(偶数还是奇数时刻)变成了黑色
2. 在后续时刻, 它就是在那里震荡了