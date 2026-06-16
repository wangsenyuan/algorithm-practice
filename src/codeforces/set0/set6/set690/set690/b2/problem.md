# B2. Recover Polygon (medium)

[Problem link](https://codeforces.com/problemset/problem/690/B2)

time limit per test: 4 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Heidi has verified her Zombie Contamination checker. Now the lair is a **strictly
convex polygon** on the integer lattice: every vertex is a lattice point, and the
polygon has no `180°` vertices.

For each unit cell, the contamination level is `0`–`4`: the number of the cell's
four corners inside or on the polygon border.

Given the full grid of levels, recover the exact shape of the lair.

## Input

The input contains multiple test cases.

The first line of each test case contains an integer `N` (`5 <= N <= 500`) — the
grid size.

Each of the next `N` lines contains `N` characters `0`–`4`, the contamination in
each cell.

Rows are given in **decreasing** `y`:

- the first row is cells `(1, N), (2, N), ..., (N, N)`
- the last row is cells `(1, 1), (2, 1), ..., (N, 1)`

Within a row, columns are in **increasing** `x`.

The last line of the file is `0`. Do not treat it as a test case. The sum of all
`N` in one file is at most `5000`.

## Output

For each test case:

- Print `V`, the number of polygon vertices.
- Print the next `V` lines: two integers per line — vertex coordinates in
  **clockwise** order, starting from the **lexicographically smallest** vertex
  `(x, y)` (`x` smaller first; if equal, smaller `y`).

## Examples

### Input

```text
8
00000000
00000110
00012210
01234200
02444200
01223200
00001100
00000000
5
00000
01210
02420
01210
00000
7
0000000
0122100
0134200
0013200
0002200
0001100
0000000
0
```

### Output

```text
4
2 3
2 4
6 6
5 2
4
2 2
2 3
3 3
3 2
3
2 5
4 5
4 2
```

## Note

The solution exists and is unique for every test case. In the correct answer, every
vertex coordinate is between `2` and `N - 2` (inclusive).

Vertex `(x1, y1)` is lexicographically smaller than `(x2, y2)` if `x1 < x2`, or if
`x1 = x2` and `y1 < y2`.

## Solution Summary

Each cell value is the sum of four unknown boolean values: whether each corner
lattice point is inside or on the polygon.

Let `inside[x][y]` be `1` if lattice point `(x, y)` is inside/on the polygon, and
`0` otherwise. Because the true polygon vertices are strictly inside the grid,
the outer boundary lattice points are outside. Therefore, after reading the cell
values bottom-up, we can recover every `inside` value by inversion.

In the problem's coordinate system, cell `(x, y)` has corners:

```text
(x-1, y-1), (x, y-1), (x-1, y), (x, y)
```

Thus:

```text
cell[x][y] =
    inside[x-1][y-1]
  + inside[x][y-1]
  + inside[x-1][y]
  + inside[x][y]
```

so:

```text
inside[x][y] =
    cell[x][y]
  - inside[x-1][y-1]
  - inside[x][y-1]
  - inside[x-1][y]
```

The implementation scans `y` from bottom to top and `x` from left to right, so
all three values on the right-hand side are already known.

After recovering all lattice points inside/on the polygon, the original polygon
is exactly the convex hull of those points. Boundary lattice points lying on an
edge are removed by the monotonic-chain hull condition `cross <= 0`, leaving only
the true vertices.

The standard monotonic-chain hull returns vertices in counterclockwise order
starting from the lexicographically smallest point. The output requires clockwise
order, so the implementation keeps the first vertex and reverses the remaining
vertices.

## Complexity

For each test case, inversion scans `O(N^2)` cells. The hull is built from at
most `(N+1)^2` recovered lattice points, so sorting dominates with
`O(P log P)`, where `P` is the number of inside/on lattice points. Memory usage
is `O(N^2)`.
