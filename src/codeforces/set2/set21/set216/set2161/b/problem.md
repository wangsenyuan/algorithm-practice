# B. Make Connected (Codeforces 2161B)

**Limits:** 1.5 seconds per test, 512 MB  
**I/O:** standard input / standard output

Source: [https://codeforces.com/problemset/problem/2161/B](https://codeforces.com/problemset/problem/2161/B)

## Problem Statement

You are given an `n x n` grid where some cells are colored black, and the rest are white.

You may paint some white cells black to achieve the following conditions:

1. There is at least one black cell.
2. All black cells must be orthogonally connected. That is, it should be possible to go from any black cell to any other by crossing several vertical or horizontal cell borders while visiting only black cells. You cannot go directly through the corner of a cell.
3. There are no three consecutive black cells aligned vertically or horizontally.

You cannot paint black cells white.

Determine whether it is possible to paint some white cells black in order to satisfy all the conditions.

## Input

Each test contains multiple test cases.

The first line contains the number of test cases `t` (`1 <= t <= 1000`).

For each test case:

- The first line contains an integer `n` (`1 <= n <= 100`) — the size of the grid.
- The following `n` lines contain `n` characters each:
  - `.` — a white cell;
  - `#` — a black cell.

It is guaranteed that the sum of `n` over all test cases does not exceed `2000`.

## Output

For each test case, print `YES` if it is possible to paint some white cells black to satisfy all the conditions, and `NO` otherwise.

You may print each letter in any case. For example, `yEs`, `yes`, `Yes`, and `YES` are all accepted.

## Example

### Input

```text
11
1
.
1
#
3
.##
.##
...
3
#..
.#.
..#
3
###
...
...
3
#.#
...
.#.
4
####
#..#
#..#
####
3
..#
...
.#.
3
..#
#..
...
5
#.#.#
.#.#.
#.#.#
.#.#.
#.#.#
5
...#.
...#.
.....
##...
.....
```

### Output

```text
YES
YES
YES
YES
NO
NO
NO
YES
YES
NO
YES
```

## Note

In the first test case, there are no black cells, so we must paint one cell black.

In the second and third test cases, the grid satisfies all the conditions from the beginning.

In the fourth test case, one possible solution is:

```text
##.
.##
..#
```

In the fifth test case, the grid already violates the condition that no three consecutive black cells should be aligned vertically or horizontally, so there is no solution.

In the sixth test case, it can be shown that it is impossible to achieve a connected grid without violating the condition about three consecutive black cells.


## Solution

The final black cells must be connected, but we are forbidden to create three consecutive black cells in one row or one column.

This severely restricts the shape of any connected component.  A long straight segment is impossible, so when the component keeps moving, it must alternate between horizontal and vertical steps.

The typical shape is a staircase:

```text
##
.##
..##
...##
```

or one of its mirrored versions.  Consecutive cells may share a row for one step, then must turn; after that they may share a column for one step, then must turn again.

So, except for the small `2 x 2` square case, all black cells must be embeddable into one alternating path.

### Special Case: `2 x 2` Square

The shape:

```text
##
##
```

is valid:

- it is connected;
- it has no three consecutive black cells in any row or column.

But it is not a simple alternating path in the same sense as the staircase walk used by the implementation.  Therefore the code handles this case directly when there are exactly four black cells.

### Checking Staircases

Let the existing black cells be fixed.  We may paint extra cells black, so it is enough to check whether there exists a valid staircase path that visits all existing black cells.

The implementation tries candidate starting cells near the first existing black cell.  This is sufficient because if a staircase contains that black cell, then the start of the relevant tested pattern is either that cell or one of its neighboring cells.

For each candidate start, it tests four possible staircase directions:

1. right, then down, then right, then down, ...
2. down, then right, then down, then right, ...
3. left, then down, then left, then down, ...
4. down, then left, then down, then left, ...

While walking such a path, it counts how many original black cells are encountered.  If the walk leaves the grid after having seen all original black cells, then all required black cells can be connected by painting the missing cells along this staircase.

The walk itself is valid because it alternates direction every step, so it never creates three consecutive black cells horizontally or vertically.

### Empty Grid

If there are no black cells initially, the answer is always `YES`: paint any single cell black.

### Algorithm

1. Collect all cells that are initially black.
2. If there are none, return `YES`.
3. If there are exactly four cells and they form a `2 x 2` square, return `YES`.
4. For every offset `dx, dy` in `[-1, 0, 1]`, use the cell near the first black cell as a candidate staircase start.
5. From that start, try the four alternating direction patterns.
6. If any pattern visits all original black cells, return `YES`; otherwise return `NO`.

### Complexity

There are at most `9` candidate starts and `4` direction patterns.  Each walk has length at most `O(n)` before leaving the grid.

```text
Time:   O(n^2) to collect black cells, plus O(n) checks
Memory: O(number of black cells)
```

With `n <= 100`, this is easily enough.
