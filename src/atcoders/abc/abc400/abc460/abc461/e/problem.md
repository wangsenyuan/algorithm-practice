# E - E-liter (ABC461)

**Contest:** [ABC461](https://atcoder.jp/contests/abc461) — AtCoder Beginner Contest 461  
**Task:** [https://atcoder.jp/contests/abc461/tasks/abc461_e](https://atcoder.jp/contests/abc461/tasks/abc461_e)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 475 points

## Problem Statement

There is an `N × N` grid. Initially, all cells are painted white.

Process `Q` queries in order. Each query is one of:

- **Type 1:** given `R`, paint the `R`-th row from the top black.
- **Type 2:** given `C`, paint the `C`-th column from the left white.

After each query, output the number of black cells in the grid.

## Constraints

- `1 <= N, Q <= 3 × 10^5`
- For type 1 queries, `1 <= R <= N`.
- For type 2 queries, `1 <= C <= N`.
- All input values are integers.

## Input

```text
N Q
query_1
query_2
⋮
query_Q
```

Each `query_i` is one of:

```text
1 R
```

```text
2 C
```

## Output

Print `Q` lines. The `i`-th line is the number of black cells after the `i`-th query.

## Sample Input 1

```text
3 4
1 1
1 3
2 2
1 1
```

## Sample Output 1

```text
3
6
4
5
```

Grid changes (`.` = white, `#` = black):

```text
...    ###    ###    #.#    ###
... -> ... -> ... -> ... -> ...
...    ...    ###    #.#    #.#
```

## Sample Input 2

```text
300000 1
2 300000
```

## Sample Output 2

```text
0
```

Beware of overflow in larger cases.

## Solution

### Key idea

Cell `(r, c)` is black iff row `r` was painted black **after** column `c` was last
painted white. Later operations overwrite earlier ones on the affected row or
column.

Track:

- `lastRow[r]` — query index when row `r` was last painted black (`-1` if never)
- `lastCol[c]` — query index when column `c` was last painted white (`-1` if never)

Then `(r, c)` is black iff `lastRow[r] > lastCol[c]`.

### Incremental updates

Maintain the total black count `tot` and update it per query instead of rescanning
the grid.

**Type 1 — paint row `R` at time `i`:**

- First time row `R` is painted: `tot += N` (all `N` cells in the row turn black).
- Repaint row `R` (previous paint at `lastRow[R]`): columns whitened in
  `(lastRow[R], i]` were blocking that row; repainting restores those cells, so
  `tot +=` number of such columns.
- Set `lastRow[R] = i` and record the row paint in a Fenwick tree `s1`.

**Type 2 — paint column `C` white at time `i`:**

- Rows painted in `(lastCol[C], i]` had a black cell in column `C`; whitening
  removes them, so `tot -=` number of such rows.
- Set `lastCol[C] = i` and record the column paint in a Fenwick tree `s2`.

`s1` stores row-paint events; `s2` stores column-white events. Range counts on
these BITs answer “how many rows/columns were updated in an interval?” in
`O(log Q)`.

### Complexity

- Time: `O(Q log Q)`
- Space: `O(N + Q)`
- Use 64-bit integers for `tot` when `N` is large (`N²` can exceed 32-bit range).
