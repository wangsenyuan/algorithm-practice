# D. Cycle

[Problem link](https://codeforces.com/problemset/problem/135/D)

time limit per test: 3 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Little Petya very much likes rectangular tables that consist of characters `0` and `1`.
Recently he has received one such table as a gift from his mother. The table
contained `n` rows and `m` columns. The rows are numbered from top to bottom from
`1` to `n`, the columns are numbered from the left to the right from `1` to `m`.
Petya immediately decided to find the longest cool cycle whatever it takes.

A cycle is a sequence of pairwise distinct cells where each two consecutive
cells have a common side; besides, the first cell has a common side with the
last cell. A cycle is called cool if it fulfills all the following conditions
simultaneously:

- The cycle entirely consists of the cells that contain `1`.
- Each cell that belongs to the cycle has a common side with exactly two other
  cells that belong to the cycle.
- Each cell of the table that contains `1` either belongs to the cycle or is
  positioned outside of it (see definition below).

To define the notion of "outside" formally, let's draw a cycle on a plane. Let
each cell of the cycle `(i, j)` (`i` is the row number, `j` is the column number)
correspond to the point `(i, j)` on the coordinate plane. Let a straight line
segment join each pair of points that correspond to the cells belonging to the
cycle and sharing a side. Thus, we will get a closed polyline that has no
self-intersections and self-touches. The polyline divides the plane into two
connected parts: the part of an infinite area and the part of a finite area. It
is considered that cell `(r, c)` lies outside of the cycle if it does not belong
to the cycle and the corresponding point on the plane with coordinates `(r, c)`
lies in the part with the infinite area.

Help Petya to find the length of the longest cool cycle in the table. The cycle
length is defined as the number of cells that belong to the cycle.

## Solution

Think about a valid cool cycle as the boundary of some finite empty region.
Every cell inside the cycle must be `0`, because the statement says every `1`
cell must either belong to the cycle or be outside it. Therefore, if a cool cycle
has a non-empty interior containing grid points, then the interior is made from
one or more `0` cells, and the answer can be found by looking at bounded
components of `0`.

There is one special case: a `2 x 2` block of `1` cells is already a cycle of
length `4`. Its geometric inside is the small square between the four cell
centers, so it contains no cell center and no `0` component. The code checks this
case directly:

```text
11
11
```

For all longer cycles, the inside contains at least one `0` cell.

### Why `0` Cells Are Joined Diagonally

The cycle is drawn through cell centers. If two `0` cells touch even by a corner,
then they are in the same open region with respect to such a polyline boundary:
a cycle of side-adjacent `1` cells cannot pass between two diagonally touching
cell centers without self-touching at a grid corner.

So the implementation groups `0` cells by 8-neighbor connectivity:

```text
up, down, left, right, and the four diagonals
```

This is done by `getComponents`.

Any `0` component that touches the border is connected to the infinite outside
region, so it cannot be the inside of a cycle. The array `ok[root]` marks exactly
the `0` components that do not touch any border cell.

### Candidate From One Bounded `0` Component

For each bounded `0` component, `play(root)` tries to construct the cycle around
it.

First, it BFS-es through all `0` cells of this component using 8-neighbor moves.
While doing this, whenever it sees a neighboring `1` cell, it adds that `1` cell
to `todo`.

So after this scan:

```text
todo = all `1` cells adjacent to this bounded `0` component
```

If this component is truly the inside of a cool cycle, then `todo` must be
exactly the cells of that cycle.

The code then verifies two properties.

### Property 1: The Boundary Must Be Connected

Cycle cells are side-adjacent, so `todo` must be connected by 4-neighbor moves.

The helper `bfs(todo[0])` walks only through cells whose `vis` value marks them
as members of `todo`. If the number of reached cells is not `len(todo)`, then
the surrounding `1` cells do not form one connected side-adjacent boundary, so
this component cannot produce a valid cycle.

### Property 2: No Boundary Cell Can Have Degree More Than 2

In a simple cycle, every cycle cell is adjacent by side to exactly two other
cycle cells. The code checks, for every cell in `todo`, how many of its four
neighbors are also in `todo`.

If any cell has more than two such neighbors, the boundary branches or touches
itself, so it is not a valid simple cycle.

Together with 4-connectivity around a bounded `0` component, this rules out
invalid boundaries. If the checks pass, `len(todo)` is the length of the cycle.

### Why This Also Excludes `1` Cells Inside

Suppose a cycle surrounds a bounded region but there is an extra `1` cell inside
that is not on the cycle. Then that inside `1` cell is adjacent, possibly through
other inside cells, to the same bounded region. During the 8-neighbor scan of the
inside `0` component, it would be collected into `todo`.

Then one of the checks fails:

- either the collected `1` cells are not one 4-connected boundary;
- or some collected cell has too many side-neighbors among collected cells.

This matches the cool-cycle requirement that every `1` cell must be either on
the cycle or outside it.

### Algorithm

1. Build 8-connected components of all `0` cells with DSU.
2. Mark every `0` component touching the grid border as invalid, because it is
   part of the outside region.
3. Check every `2 x 2` all-`1` block and set the answer to at least `4`.
4. For every remaining bounded `0` component:
   - collect all neighboring `1` cells into `todo`;
   - check that `todo` is 4-connected;
   - check that no cell in `todo` has more than two 4-neighbors inside `todo`;
   - if valid, update the answer with `len(todo)`.

The DSU and all BFS scans are linear in the number of cells. Each `0` component
is processed once, so the total complexity is `O(nm)`.
