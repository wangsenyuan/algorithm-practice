# G - Big Banned Grid (ABC413)

**Contest:** [ABC413](https://atcoder.jp/contests/abc413) — Denso Create Programming Contest 2025 (AtCoder Beginner Contest 413)  
**Task:** [https://atcoder.jp/contests/abc413/tasks/abc413_g](https://atcoder.jp/contests/abc413/tasks/abc413_g)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 575 points

## Problem Statement

There is an `H x W` grid. Let `(i,j)` denote the cell at the `i`-th row
(`1 <= i <= H`) from the top and `j`-th column (`1 <= j <= W`) from the left.

Each cell in the grid either has an obstacle placed on it or has nothing placed
on it. There are `K` cells with obstacles: cells `(r_1, c_1), (r_2, c_2), ...,
(r_K, c_K)`.

Takahashi is initially at cell `(1,1)` and wants to reach cell `(H,W)` by
repeatedly moving to an adjacent cell (up, down, left, right) that has nothing
placed on it.

More precisely, he can repeat the following operation as many times as he likes.
Let `(i,j)` be his current cell. Choose one of the following four operations and
perform the chosen operation:

- If `1 < i` and cell `(i-1,j)` has nothing placed on it, move to cell `(i-1,j)`.
  Otherwise, do not move.
- If `1 < j` and cell `(i,j-1)` has nothing placed on it, move to cell `(i,j-1)`.
  Otherwise, do not move.
- If `i < H` and cell `(i+1,j)` has nothing placed on it, move to cell `(i+1,j)`.
  Otherwise, do not move.
- If `j < W` and cell `(i,j+1)` has nothing placed on it, move to cell `(i,j+1)`.
  Otherwise, do not move.

Determine whether he can move from cell `(1,1)` to cell `(H,W)`.

## Constraints

- `1 <= H <= 2 * 10^5`
- `1 <= W <= 2 * 10^5`
- `0 <= K <= 2 * 10^5`
- `1 <= r_i <= H` for `1 <= i <= K`
- `1 <= c_i <= W` for `1 <= i <= K`
- `(r_i, c_i) != (1, 1)` for `1 <= i <= K`
- `(r_i, c_i) != (H, W)` for `1 <= i <= K`
- `(r_i, c_i) != (r_j, c_j)` for `1 <= i < j <= K`
- All input values are integers.


## Solution

Instead of running reachability on the full `H x W` grid, only the `K` obstacles
matter. The answer is determined by how obstacles connect the four grid borders
under **8-direction adjacency** (including diagonals).

### DSU on obstacles + borders

Create one DSU node for each obstacle, plus four virtual border nodes:

- `top`: row `0`
- `right`: column `W + 1`
- `bottom`: row `H + 1`
- `left`: column `0`

For each obstacle at `(r, c)`, look at all eight neighbors in the expanded
coordinate system. Skip the four corner positions `(0,0)`, `(0,W+1)`, `(H+1,0)`,
and `(H+1,W+1)`.

- If a neighbor is outside the grid on the top / bottom / left / right side,
  union the obstacle with the corresponding border node.
- If a neighbor is another obstacle, union the two obstacle nodes.

After processing all obstacles, check four border pairs:

| Pair           | Meaning                                |
| -------------- | -------------------------------------- |
| top ↔ bottom   | a horizontal wall blocks every path    |
| left ↔ right   | a vertical wall blocks every path      |
| top ↔ left     | the start corner `(1,1)` is sealed off |
| bottom ↔ right | the end corner `(H,W)` is sealed off   |

If any pair is in the same DSU component, print `No`; otherwise print `Yes`.

### Why this works

A 4-direction path from `(1,1)` to `(H,W)` is blocked exactly when obstacles,
expanded by 8-connectivity, link two borders that must stay separated for such a
path to exist. So the problem reduces to four connectivity checks on `K + 4`
nodes, not a full-grid BFS.

### Complexity

- Time: `O(K alpha(K))`
- Space: `O(K)`

### Why use 8-connectivity between obstacles

Movement is 4-directional, but obstacles should be merged diagonally as well.

If two obstacles touch only at a diagonal, they still pinch off a narrow passage.
For example, obstacles at `(2,2)` and `(3,3)` leave no room for a 4-direction path
to slip between them diagonally. So two obstacles belong to the same "wall" when
they are 8-adjacent, not only 4-adjacent.

### Why add virtual border nodes

An obstacle on row `1` has an 8-neighbor at row `0`. That means the obstacle touches
the top border of the grid. The same idea applies to the other three sides.

Instead of drawing the full `H x W` grid, we represent each side as one DSU node and
union an obstacle with the border node whenever one of its eight neighbors falls
outside the grid on that side.

This turns "does a wall reach the top?" into "is this obstacle in the same DSU
component as the top node?"

### Why skip the four corner neighbors

When scanning the eight neighbors of an obstacle, positions like `(0,0)` sit on two
borders at once. If we union both borders there, we would connect `top` and `left`
(or `bottom` and `right`) during construction itself.

We do not want that. Those pairs are checked separately at the end. So neighbors at

- `(0,0)`, `(0,W+1)`
- `(H+1,0)`, `(H+1,W+1)`

are ignored.

All other out-of-grid neighbors connect the obstacle to exactly one border node.

### What the four final checks mean

After all unions, each DSU component describes one 8-connected obstacle cluster,
possibly touching one or more borders.

1. **`top` connected to `bottom`**

   The cluster forms a horizontal barrier from the top side to the bottom side. Every
   path from the upper part of the grid to the lower part is blocked. Since `(1,1)` is
   above `(H,W)`, no route exists.

2. **`left` connected to `right`**

   The cluster forms a vertical barrier from the left side to the right side. Every
   path from the left part to the right part is blocked. Again, start and end lie on
   opposite sides.

3. **`top` connected to `left`**

   The cluster touches both borders that meet at `(1,1)`. The start cell only has two
   empty directions: right and down. If the top and left borders are linked through
   obstacles near that corner, `(1,1)` is sealed off.

4. **`bottom` connected to `right`**

   Symmetric to the previous case at `(H,W)`, whose only exits are up and left.

If none of these four bad connections happen, a valid path exists. Otherwise the
answer is `No`.

### How the code implements this

1. Map each obstacle coordinate to an index `0 .. K-1`.
2. Create a DSU of size `K + 4`:
   - `K` = top
   - `K + 1` = right
   - `K + 2` = bottom
   - `K + 3` = left
3. For every obstacle, scan all eight offsets, apply the corner skip rule, then:
   - union with a border node if the neighbor is outside on one side
   - union with another obstacle if the neighbor is an obstacle cell
4. Test the four pairs `{top, bottom}`, `{right, left}`, `{top, left}`,
   `{bottom, right}`; return `false` if any pair shares a root.

The DSU only stores `K + 4` elements, so this stays fast even when `H` and `W` are
as large as `2 * 10^5`.

### Example: Sample Input 1

Grid size is `4 x 5` with obstacles at `(1,4)`, `(2,3)`, `(3,2)`, `(3,4)`,
`(4,2)`.

These obstacles 8-connect into a cluster that links the top border and the bottom
border. The `{top, bottom}` check fails, so the answer is `No`. This matches the
official sample without running BFS on all `20` cells.