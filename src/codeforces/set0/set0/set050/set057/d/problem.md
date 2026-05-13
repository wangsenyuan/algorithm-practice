# D. Average lifespan (Codeforces 57D)

Stewie the Rabbit explores a new parallel universe shaped as a rectangular grid with `n` rows and `m` columns. The universe is small: each cell holds **at most one** particle.

Each particle is either **static** or **dynamic**.

- **Static** particles never move. No two static particles share a **row** or a **column**, and none occupy **diagonally adjacent** cells.
- A **dynamic** particle appears in a **uniformly random empty** cell, then chooses a **uniformly random empty** cell as its destination (the destination may equal the start). It moves along a **shortest path** through cells not occupied by static particles. When it reaches the destination, it disappears. Only one dynamic particle exists at a time.

The particle may move between **edge-adjacent** cells; each such step takes **one galactic second**.

Find the **expected** (average) lifespan of one such particle.

## Input

The first line contains two integers `n` and `m` (`2 <= n, m <= 1000`) — grid dimensions.

The next `n` lines each contain `m` characters:

- `'X'` — static particle in that cell  
- `'.'` — empty cell  

The grid is guaranteed to satisfy the static-particle rules above.

## Output

Print a single real number — the average lifespan, with **at least 6** decimal digits shown.

The answer is accepted if the absolute or relative error is at most `10^-6`.

## Examples

### Example 1

**Input**

```text
2 2
..
.X
```

**Output**

```text
0.888888888889
```

### Example 2

**Input**

```text
3 3
...
.X.
...
```

**Output**

```text
2.000000000000
```

Original: [Codeforces 57D](https://codeforces.com/problemset/problem/57/D)


## Solution

Let `F` be the number of empty cells.  The start cell and destination cell are chosen independently and uniformly from all empty cells, and they may be the same cell.

So if `S` is the sum of shortest-path distances over all **unordered pairs of different empty cells**, then the required expectation is:

```text
2 * S / F^2
```

The factor `2` is because every unordered pair `{a, b}` appears as two ordered choices: `(a, b)` and `(b, a)`.  Pairs `(a, a)` contribute `0`.

### Base Manhattan Contribution

Without static cells, the shortest distance between `(r1, c1)` and `(r2, c2)` is:

```text
abs(r1 - r2) + abs(c1 - c2)
```

We first sum this Manhattan distance over all unordered pairs of empty cells.

For the row part, let:

```text
row[i] = number of empty cells in row i
```

When processing rows from top to bottom, suppose previous rows contain:

```text
cnt = total empty cells in previous rows
sum = sum(row_index * empty_count) over previous rows
```

Then row `i` contributes:

```text
i * cnt * row[i] - sum * row[i]
```

This is exactly:

```text
sum((i - previous_row) for every pair of cells in different rows)
```

The column part is computed symmetrically using:

```text
col[j] = number of empty cells in column j
```

This gives the total Manhattan contribution in `O(nm)`.

### Why Obstacles Only Add `+2`

Because no two static cells share a row or column, and no two are diagonally adjacent, a shortest path is blocked only by a thin monotone barrier.  If a Manhattan path exists, the distance stays Manhattan.  If all Manhattan paths are blocked, going around the barrier needs exactly two extra steps, so the distance becomes:

```text
Manhattan + 2
```

Therefore the remaining task is to count how many unordered pairs of empty cells have all Manhattan shortest paths blocked.

There are three kinds of blockers.

### 1. One Static Cell in the Same Row

For a static cell `X` at `(r, c)`, pairs in row `r` with one endpoint left of `c` and the other endpoint right of `c` cannot pass straight through that row.

Let:

```text
left  = number of empty cells in row r before c
right = number of empty cells in row r after c
```

Then this static cell contributes:

```text
2 * left * right
```

to `S`, because each such unordered pair needs an extra `+2` distance.

The implementation scans each row and maintains `left`.  Since there is at most one `X` in a row, this is enough for all same-row direct blockers.

### 2. One Static Cell in the Same Column

The column case is identical.

For a static cell `X` at `(r, c)`, pairs in column `c` with one endpoint above `r` and the other endpoint below `r` need an extra `+2`.

Let:

```text
up   = number of empty cells in column c before r
down = number of empty cells in column c after r
```

Then the contribution is:

```text
2 * up * down
```

### 3. Monotone Zigzag Chains

The first two cases are not sufficient.  Several static cells can form a zigzag barrier even when no single `X` lies directly between the two endpoints.

Example:

```text
...X.
.X...
```

For the pair `(0, 4)` and `(1, 0)`, the Manhattan distance is `5`, but every Manhattan path is blocked by the two static cells.  The true distance is `7`.

This happens when static cells occupy consecutive rows and their columns move monotonically.

Consider a chain:

```text
(r,     c1)
(r + 1, c2)
...
(r + t, ct)
```

where the rows are consecutive and:

```text
c1 < c2 < ... < ct
```

This increasing chain blocks all Manhattan paths between:

- a cell in row `r` strictly left of `c1`;
- a cell in row `r + t` strictly right of `ct`.

There are:

```text
c1 * (m - 1 - ct)
```

such unordered pairs, and each needs `+2`, so the contribution is:

```text
2 * c1 * (m - 1 - ct)
```

For a decreasing chain:

```text
c1 > c2 > ... > ct
```

the blocked endpoints are reversed:

- a cell in row `r` strictly right of `c1`;
- a cell in row `r + t` strictly left of `ct`.

The contribution is:

```text
2 * (m - 1 - c1) * ct
```

The implementation sorts static cells by row, splits them into maximal groups of consecutive rows, and checks every monotone subchain inside each group.  Since there is at most one static cell per row, this is direct.

The same idea is applied after swapping rows and columns.  That counts monotone chains in consecutive columns.

### Why There Is No Double Counting

Each blocked unordered pair has a unique tight rectangle whose opposite corners are the two cells.  If no Manhattan path exists inside that rectangle, the static cells that block it form one boundary chain of that rectangle:

- either a single static cell in the same row;
- or a single static cell in the same column;
- or a monotone chain touching consecutive rows;
- or, after swapping axes, a monotone chain touching consecutive columns.

The static-cell constraints prevent two independent blockers of the same tight rectangle.  In particular, static cells cannot share a row/column, and cannot be diagonally adjacent, so a blocked Manhattan rectangle has one such boundary chain.  Thus every pair that needs `+2` is counted once.

### Complexity

Let `s` be the number of static cells.  Because no two static cells share a row or a column:

```text
s <= min(n, m) <= 1000
```

The algorithm does:

```text
Base Manhattan sum:       O(nm)
Direct row/column blocks: O(nm)
Monotone chain counting:  O(s^2)
Memory:                  O(n + m + s)
```

This fits easily for `n, m <= 1000`.
