# D - Takahashi the Wall Breaker (ABC400)

**Contest:** [ABC400](https://atcoder.jp/contests/abc400) — AtCoder Beginner Contest 400  
**Task:** [https://atcoder.jp/contests/abc400/tasks/abc400_d](https://atcoder.jp/contests/abc400/tasks/abc400_d)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 400 points

## Problem Statement

Takahashi is about to go buy eel at a fish shop.

The town where he lives is divided into a grid of H rows and W columns. Each cell is either a road or a wall. Let us denote the cell at the i-th row from the top (1 ≤ i ≤ H) and the j-th column from the left (1 ≤ j ≤ W) as cell (i, j). Information about each cell is given by H strings S_1, S_2, …, S_H, each of length W. Specifically, if the j-th character of S_i (1 ≤ i ≤ H, 1 ≤ j ≤ W) is `.`, cell (i, j) is a road; if it is `#`, cell (i, j) is a wall.

He can repeatedly perform the following two types of actions in any order:

- Move to an adjacent cell (up, down, left, or right) that is within the town and is a road.
- Choose one of the four directions (up, down, left, or right) and perform a front kick in that direction. When he performs a front kick, for each of the cells at most 2 steps away in that direction from the cell he is currently in, if that cell is a wall, it becomes a road. If some of the cells at most 2 steps away are outside the town, a front kick can still be performed, but anything outside the town does not change.

He starts in cell (A, B), and he wants to move to the fish shop in cell (C, D). It is guaranteed that both the cell where he starts and the cell with the fish shop are roads. Find the minimum number of front kicks he needs in order to reach the fish shop.

## Constraints

- 1 ≤ H ≤ 1000
- 1 ≤ W ≤ 1000
- Each S_i is a string of length W consisting of `.` and `#`.
- 1 ≤ A, C ≤ H
- 1 ≤ B, D ≤ W
- (A, B) ≠ (C, D)
- H, W, A, B, C, and D are integers.
- The cell where Takahashi starts and the cell with the fish shop are roads.

## Input

The input is given from Standard Input in the following format:

```text
H W
S_1
S_2
⋮
S_H
A B C D
```

## Output

Print the minimum number of front kicks needed for Takahashi to reach the fish shop.

## Sample Input 1

```text
10 10
..........
#########.
#.......#.
#..####.#.
##....#.#.
#####.#.#.
.##.#.#.#.
###.#.#.#.
###.#.#.#.
#.....#...
1 1 7 1
```

## Sample Output 1

```text
1
```

Takahashi starts in cell (1, 1). By repeatedly moving to adjacent road cells, he can reach cell (7, 4). If he performs a front kick to the left from cell (7, 4), cells (7, 3) and (7, 2) turn from walls to roads. Then, by continuing to move through road cells (including those that have become roads), he can reach the fish shop in cell (7, 1).

In this case, the number of front kicks performed is 1, and it is impossible to reach the fish shop without performing any front kicks, so print 1.

## Sample Input 2

```text
2 2
.#
#.
1 1 2 2
```

## Sample Output 2

```text
1
```

Takahashi starts in cell (1, 1). When he performs a front kick to the right, cell (1, 2) turns from a wall to a road. The cell two steps to the right of (1, 1) is outside the town, so it does not change. Then, he can move to cell (1, 2) and then to the fish shop in cell (2, 2).

In this case, the number of front kicks performed is 1, and it is impossible to reach the fish shop without performing any front kicks, so print 1.

## Sample Input 3

```text
1 3
.#.
1 1 1 3
```

## Sample Output 3

```text
1
```

When performing a front kick, it is fine if the fish shop's cell is within the cells that could be turned into a road. Specifically, the fish shop's cell is a road from the beginning, so it remains unchanged; particularly, the shop is not destroyed by the front kick.

## Sample Input 4

```text
20 20
####################
##...##....###...###
#.....#.....#.....##
#..#..#..#..#..#..##
#..#..#....##..#####
#.....#.....#..#####
#.....#..#..#..#..##
#..#..#.....#.....##
#..#..#....###...###
####################
####################
##..#..##...###...##
##..#..#.....#.....#
##..#..#..#..#..#..#
##..#..#..#..#..#..#
##.....#..#..#..#..#
###....#..#..#..#..#
#####..#.....#.....#
#####..##...###...##
####################
3 3 18 18
```

## Sample Output 4

```text
3
```

## Solution

Find the minimum number of front kicks to go from `(A, B)` to `(C, D)` on an
`H × W` grid of roads (`.`) and walls (`#`). Moves are free; each kick costs 1.

### Short recap

- **Move:** step to an adjacent in-bounds cell that is a road (`.`) — cost 0.
- **Front kick:** pick a direction; every wall among the cells at distance 1 and
  2 in that direction becomes a road — cost 1. Cells outside the grid are ignored.
- **Goal:** minimize kicks; movement count is irrelevant.

### Key observations

1. **State = position, not broken walls.** We only need `dist[r][c]` = minimum
   kicks to *reach* cell `(r, c)`. Which walls were broken is implied: if
   `dist[r][c] = k`, Takahashi can stand on `(r, c)` after exactly `k` kicks along
   some path. We never mutate the input grid.

2. **Two transition types → 0-1 BFS.** Edge weights are only 0 or 1, so Dijkstra
   becomes **0-1 BFS** with a deque:
   - **0-cost (push front):** from `(r, c)`, move to a neighbor `(nr, nc)` with
     `grid[nr][nc] == '.'` and strictly better `dist`.
   - **1-cost (push tail):** from `(r, c)`, kick in one of four directions; for
     the cells at offsets 1 and 2 along that ray (if in bounds), relax
     `dist = dist[r][c] + 1`.

3. **Movement checks the destination, not the source.** A move requires the
   *target* cell to be `.` in the original grid. Once a kick assigns
   `dist[r][c] = k` to a wall, we can leave that cell toward an adjacent road
   without re-checking whether the source was originally `#`. That models “the
   kick broke this wall, then I walked out.”

4. **Kick relaxes both distance-1 and distance-2 cells.** A kick affects up to
   two cells in a direction. The inner loop in `solve` visits offset 1, then
   offset 2, and relaxes each in-bounds cell — wall or road — with cost `+1`.
   Roads keep their old distance (already reachable with fewer kicks).

5. **Stale entries.** Nodes store `(r, c, dist)`; on pop, skip if
   `dist[r][c] != cur.dist` (standard Dijkstra / 0-1 BFS lazy deletion).

### Algorithm (matching `solve`)

**Step 0 — parse and 0-index.** Convert `(A, B)` and `(C, D)` to 0-based
`(a, b)` and `(c, d)`.

**Step 1 — initialize.**

```go
dist[r][c] = inf   for all cells
dist[a][b] = 0
dq.PushTail(node{a, b, 0})
```

**Step 2 — 0-1 BFS.** While the deque is non-empty:

1. Pop front → `(r, c, d)`. Skip if `dist[r][c] != d`.
2. **Moves (cost 0):** for each of 4 neighbors `(nr, nc)` in bounds with
   `grid[nr][nc] == '.'` and `dist[nr][nc] > d`, set `dist[nr][nc] = d` and
   `PushFront`.
3. **Kicks (cost 1):** for each direction, let `(nr, nc)` be offset 1, then
   offset 2. If in bounds and `dist[nr][nc] > d + 1`, set `dist[nr][nc] = d + 1`
   and `PushTail`.

**Step 3 — answer.** Return `dist[c][d]`.

Direction deltas use the closed loop `dd = {-1, 0, 1, 0, -1}` so direction `i`
uses `(dd[i], dd[i+1])`.

### Why it works

Every path from start to goal is a sequence of moves (free) and kicks (paid).
Split it at each kick: between consecutive kicks, Takahashi walks only on roads
that are either original `.` cells or walls broken by an earlier kick in that
segment.

`dist[r][c]` is the minimum kicks among all such paths ending at `(r, c)`:

- **Move relaxation** adds a 0-weight edge from `(r, c)` to an adjacent original
  road whenever we already know a `k`-kick route to `(r, c)`.
- **Kick relaxation** adds a 1-weight edge from `(r, c)` to each cell within
  kick range: standing at `(r, c)` after `k` kicks, one more kick makes those
  cells reachable with `k + 1` kicks.

0-1 BFS processes non-decreasing kick counts in order (deque front = same kicks,
back = one more kick), so the first time we settle `(r, c)` we have the minimum.
No separate “broken wall” bitmap is needed because any cell first entered with
`k` kicks is reachable, and outgoing moves only require the neighbor to be an
original road (or a cell already opened by a kick on the way in).

### Walkthrough: sample 2 (`2 × 2`, start `(1,1)`, goal `(2,2)`)

Grid (1-indexed):

```text
. #
# .
```

0-indexed: start `(0, 0)`, goal `(1, 1)`. Only `(0, 1)` is a wall blocking a
direct path.

| step | action | `dist` highlights |
| --- | --- | --- |
| init | `dist[0][0] = 0`, push `(0,0,0)` | |
| pop `(0,0,0)` | moves: none (only neighbor `(0,1)` is `#`) | |
| | kick right: `dist[0][1] ← 1`, push tail `(0,1,1)` | `(0,1)` now reachable with 1 kick |
| pop `(0,1,1)` | move down: `grid[1][1] == '.'`, `dist[1][1] ← 1`, push front | goal reached with 1 kick |
| pop `(1,1,1)` | already optimal | |

Output: **1**. One kick breaks `(1,2)`, then two free moves `(1,2) → (2,2)`.

Sample 3 (`1 × 3`, `. # .`, `(1,1)` → `(1,3)`): a single kick to the right
relaxes `dist[0][1]` and `dist[0][2]`; `dist[0][2] = 1` is the answer — kicking
through the shop cell is allowed because it was already a road.

### Complexity

- **Time:** `O(H · W)` — each cell is pushed at most a constant number of times
  (4 moves + 4×2 kick targets per pop); deque operations are `O(1)`.
- **Space:** `O(H · W)` for `dist` and the deque.

Both satisfy `H, W ≤ 1000`.
