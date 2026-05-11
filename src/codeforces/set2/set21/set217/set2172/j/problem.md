# Codeforces 2172J - Sliding Tiles

https://codeforces.com/problemset/problem/2172/J

## Statement

You have a special sliding puzzle on an `n x n` grid.

Between each pair of adjacent columns, there is a vertical bar of height `hi`
for `1 <= i < n`. The bar between columns `i` and `i + 1` is positioned at the
bottom of the grid. A height `hi` means the bar extends upward through the
bottom `hi` rows, blocking tile movement between those two columns in those
rows.

The grid contains several tiles. Each tile occupies exactly one cell. Tiles can
slide freely unless blocked by:

- the grid boundary;
- a vertical bar;
- another tile.

There are two tilt operations:

- Tilt right: all tiles slide to the right as far as possible.
- Tilt down: all tiles slide downward as far as possible.

During a tilt, all tiles move simultaneously and stop only when blocked by the
edge of the grid, a bar, or another tile.

A group operation is defined as:

1. tilt the grid to the right;
2. then tilt the grid downward.

Initially, the `i`-th column contains `ai` tiles stacked from the bottom of the
column.

Perform the group operation exactly once and determine how many tiles are in
each column afterward.

## Input

The first line contains an integer `n` — the size of the board.

The second line contains `n` integers `a1, a2, ..., an`, where `ai` is the number
of tiles initially in column `i`.

The third line contains `n - 1` integers `h1, h2, ..., h(n-1)`, where `hi` is the
height of the bar between columns `i` and `i + 1`.

Constraints:

- `2 <= n <= 5 * 10^5`
- `0 <= ai <= n`
- `0 <= hi <= n - 1`

## Output

Print `n` integers: the number of tiles in each column after performing the
group operation exactly once.

## Samples

```text
5
5 5 2 3 0
3 0 4 1
```

```text
3 3 4 2 3
```

```text
6
4 3 0 3 0 2
1 0 0 1 3
```

```text
1 0 3 3 2 3
```

## Solution

Think about one horizontal row at a time.

For row `y` counted from the bottom:

- column `i` initially contains a tile in this row iff `a[i] >= y`;
- the bar between columns `i` and `i + 1` blocks this row iff `h[i] >= y`.

So, during the right tilt, row `y` is split into independent intervals by all
blocked bars. Inside each interval, all occupied cells slide to the right, so if
the interval is `[L, R]` and it has `cnt` tiles, then after the right tilt the
occupied cells in this row are exactly:

```text
R - cnt + 1, R - cnt + 2, ..., R
```

The following down tilt does not move tiles between columns. It only stacks the
tiles inside each column, so the final answer for a column is just the number of
rows where that column is occupied after the right tilt.

## Sweep rows from bottom to top

When moving from row `y - 1` to row `y`:

- columns with `a[i] = y - 1` stop being occupied;
- bars with `h[i] = y - 1` stop blocking, so the intervals on the two sides
  merge.

This is useful because intervals only merge as `y` increases. We can maintain
the current row intervals with DSU.

For each DSU component, store:

- `l`, `r`: its column interval;
- `cnt`: how many currently active columns in this row interval contain a tile;
- `last`: the first row height since this component's current `(r, cnt)` state
  became valid.

If a component remains unchanged for rows:

```text
last, last + 1, ..., y - 1
```

then it contributes `y - last` rows to the suffix:

```text
[r - cnt + 1, r]
```

So before changing a component, flush its pending contribution into a difference
array over columns:

```text
diff[r - cnt + 1] += y - last
diff[r + 1]       -= y - last
```

Then update `last = y`.

## Events

Initialize row `1`:

- every column with `a[i] > 0` is active;
- every bar with `h[i] = 0` is already passable, so union its two columns.

Then process `y = 2..n`:

1. For each column with `a[i] = y - 1`, flush its component and decrement
   `cnt`.
2. For each bar with `h[i] = y - 1`, flush both neighboring components and
   union them.

At the end, flush every remaining component through row `n`.

Finally prefix-sum the difference array to get the number of occupied rows in
each column after the right tilt, which is also the final column height after
the down tilt.

## Complexity

Each column is removed once, each bar is merged once, and each DSU operation is
amortized inverse-Ackermann time.

The total complexity is:

```text
O(n alpha(n))
```

with `O(n)` memory.

## ideas
