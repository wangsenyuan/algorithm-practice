# B. Grid Counting

[Problem link](https://codeforces.com/problemset/problem/2150/B)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Alice has an `n × n` grid. Initially all cells are white. She wants to color some
cells black so that the following hold.

Let the black cells be `(x_1, y_1), ..., (x_m, y_m)`, and let `a` be an array of
length `n`:

- For each row `k`, there are exactly `a_k` black cells in that row.
- For each `k` (`1 <= k <= n`), there is exactly one black cell with
  `max(x_i, y_i) = k`.
- For each `k` (`1 <= k <= n`), there is exactly one black cell with
  `max(x_i, n + 1 - y_i) = k`.

Count the number of valid grids modulo `998244353`. Two grids differ if some cell
is black in one and white in the other.

For example, when `a = [2, 2, 1, 0, 0]`, one valid set of black cells is
`{(1,1), (2,2), (3,3), (2,4), (1,5)}`.

## Constraints

- `1 <= t <= 10^4`
- `2 <= n <= 2 * 10^5`
- `0 <= a_i <= n`
- Sum of `n` over all test cases does not exceed `2 * 10^5`

## Input

The first line contains `t` — the number of test cases.

For each test case:

- The first line contains `n`
- The second line contains `a_1, a_2, ..., a_n`

## Output

For each test case, print the number of valid grids modulo `998244353`.

## Example

```text
Input
5
5
2 2 1 0 0
2
2 0
2
1 1
4
3 1 0 0
4
0 0 0 0

Output
1
1
0
2
0
```

### Note

- Test 1: the only valid grid is the one in the statement.
- Test 2: the only valid grid has black cells `{(1,1), (1,2)}`.
- Test 3: no valid grids.

## Solution

The two diagonal-square conditions force every column to contain exactly one black cell.

For a black cell at row `x` and column `y`, the two special values are:

```text
max(x, y)
max(x, n + 1 - y)
```

For column `y`, the cell can be placed only in rows:

```text
x <= min(y, n + 1 - y)
```

Otherwise one of the two values would be determined by `x` too late and the required
one-per-level structure would be broken. Conversely, choosing exactly one row for each
column under this bound gives a valid grid. Therefore the problem becomes:

> There are `n` columns. Column `y` may be assigned to any row
> `1 .. min(y, n+1-y)`. Count assignments where row `i` receives exactly `a_i`
> columns.

Process rows from bottom to top. When we are at row `row`, some columns become available:

- column `row`;
- column `n + 1 - row`;
- for odd `n`, the middle column is counted only once.

These are exactly the columns whose maximum allowed row is `row`.

Let `available` be the number of currently available columns that have not yet been
assigned to a lower row. At row `row`, choose `a_row` of them:

```text
C(available, a_row)
```

Then those chosen columns are consumed:

```text
available -= a_row
```

If at some point `a_row > available`, the answer is `0`. After all rows are processed,
`available` must also be `0`, because every column must be assigned exactly once.

The final answer is the product of these binomial choices over all rows, modulo
`998244353`.

Factorials and inverse factorials are precomputed, so each combination is `O(1)`.
The time complexity is `O(n log MOD)` per test case because of one modular inverse
precomputation, and the memory complexity is `O(n)`.
