# B3. Recover Polygon (hard)

[Problem link](https://codeforces.com/problemset/problem/690/B3)

time limit per test: 5 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

The lair is a strictly convex polygon on the integer lattice. Every polygon
vertex is a lattice point.

For each unit cell, the contamination level is the number of the cell's four
corners that are inside or on the polygon border. The checker is damaged: for
each cell, it only reports whether the level is one of `{1, 2, 3}`.

In other words, the input gives exactly the cells whose level is not `0` and not
`4`. Recover the polygon.

## Input

The input contains multiple test cases.

Each test case starts with two integers `N` and `M`:

- `N` is the grid size (`5 <= N <= 100000`);
- `M` is the number of reported cells (`8 <= M <= 200000`).

The next line contains `M` pairs:

```text
x_1 y_1 x_2 y_2 ... x_M y_M
```

These are the coordinates of cells whose contamination level is `1`, `2`, or
`3`. Cells are identified by the coordinates of their upper-right corner. Thus
cell `(x, y)` has corners:

```text
(x-1, y-1), (x, y-1), (x-1, y), (x, y)
```

The last line contains `0 0` and is not a test case.

The sum of all `M` over the file is at most `200000`.

## Output

For each test case:

- print `V`, the number of polygon vertices;
- print the next `V` lines, each containing one vertex.

Vertices must be printed in clockwise order, starting from the lexicographically
smallest vertex.

## Example

### Input

```text
8 19
2 3 2 4 2 5 3 3 3 5 4 3 4 5 4 6 5 2 5 3 5 6 6 2 6 3 6 4 6 5 6 6 6 7 7 6 7 7
5 8
2 2 2 3 2 4 3 2 3 4 4 2 4 3 4 4
0 0
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
```

## Solution Summary

Reported cells are exactly the boundary cells: they have at least one corner
inside/on the polygon and at least one corner outside the polygon.

For a lattice point `p` to be inside or on the polygon, every valid cell touching
`p` must be non-zero. If one adjacent cell has level `0`, then none of its
corners is inside/on the polygon, so `p` is outside.

For a corner of a reported cell, the four cells touching that corner are:

```text
(x, y), (x+1, y), (x, y+1), (x+1, y+1)
```

where `(x, y)` is the corner coordinate. The input directly tells whether a cell
has level `1..3`. The missing part is deciding whether an unreported neighboring
cell has level `4` or level `0`.

Use the convex hull of reported cell centers for that:

- a reported cell center is represented as `(2*x-1, 2*y-1)` to avoid fractions;
- an unreported cell is treated as level `4` exactly when its center lies inside
  or on the convex hull of reported cell centers;
- otherwise it is level `0`.

Then scan every corner of every reported cell. If all four touching cells are
non-zero, add that corner as a candidate polygon point.

The candidate set contains all real vertices, and may also contain extra lattice
points on edges or inside the polygon. Taking the convex hull of all candidates
removes those extra points and leaves exactly the polygon vertices.

The monotonic-chain hull returns vertices counterclockwise from the
lexicographically smallest point. The implementation keeps that first point and
reverses the remaining vertices to produce clockwise order.

## Complexity

Let `M` be the number of reported boundary cells.

- Building the hull of reported centers costs `O(M log M)`.
- There are at most `4M` candidate corners, and each checks four neighboring
  cells. Each missing-cell classification uses `O(log M)` point-in-convex-hull
  time.
- Building the final candidate hull costs `O(M log M)`.

Overall complexity is `O(M log M)` per test case, with `O(M)` memory.
