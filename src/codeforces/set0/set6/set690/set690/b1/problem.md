# B1. Recover Polygon (easy)

[Problem link](https://codeforces.com/problemset/problem/690/B1)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Heidi needs to locate a zombie lair modeled as an axis-parallel rectangle on an
`N x N` integer lattice.

For each unit cell, the **Zombie Contamination level** is an integer from `0` to
`4`: the number of the cell's four corners that lie inside or on the border of the
rectangle.

Given the contamination levels on the whole grid, decide whether they could come
from **one** rectangle with **non-zero area** (sides parallel to the axes, vertices
on lattice points).

## Input

The first line contains an integer `N` (`5 <= N <= 50`) — the grid size.

Each of the next `N` lines contains `N` characters `0`–`4`, the contamination in
each cell.

Rows are given in **decreasing** `y`:

- the first row is cells `(1, N), (2, N), ..., (N, N)`
- the last row is cells `(1, 1), (2, 1), ..., (N, 1)`

Within a row, columns are in **increasing** `x`.

## Output

Print `Yes` if there exists a single non-zero-area axis-parallel rectangular lair
inside the grid that produces exactly the given levels. Otherwise print `No`.

## Example

### Input

```text
6
000000
000000
000001
210002
420001
210000
000000
```

### Output

```text
Yes
```

## Note

If the lair exists, it is an axis-parallel rectangle with corners at grid points
`(x1, y1)`, `(x1, y2)`, `(x2, y1)`, `(x2, y2)`, with positive area and contained in
the grid. The input levels must match the contamination produced by that rectangle.


## ideas
1. 肯定有4个角, 它们是1
2. 然后边通过2连接