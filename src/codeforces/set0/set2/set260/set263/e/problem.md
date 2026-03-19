# Problem

You are given a table of size `n x m`. At the intersection of row `i` (`1 <= i <= n`) and column `j` (`1 <= j <= m`) there is a non-negative integer `a[i][j]`. You are also given a non-negative integer `k`.

Your task is to find a pair of integers `(a, b)` such that:

- `k <= a <= n - k + 1`
- `k <= b <= m - k + 1`
- Let `mval` be the maximum value of function `f(x, y)` over all integers `x, y` satisfying:
  - `k <= x <= n - k + 1`
  - `k <= y <= m - k + 1`
  Then the required pair must satisfy `f(a, b) = mval`.

Note: in the original source text used for this file, the explicit formula for `f(x, y)` is omitted.

## Input

The first line contains three space-separated integers `n`, `m`, and `k` (`1 <= n, m <= 1000`).

The next `n` lines each contain `m` integers. The `j`-th number on the `i`-th of these lines is `a[i][j]` (`0 <= a[i][j] <= 10^6`).

Numbers in each line are separated by spaces.

## Output

Print the required pair of integers `a` and `b`, separated by a space.

If there are multiple correct answers, you may print any of them.

## Examples

### Example 1

Input

```text
4 4 2
1 2 3 4
1 1 1 1
2 2 2 2
4 3 2 1
```

Output

```text
3 2
```

### Example 2

Input

```text
5 7 3
8 2 3 4 2 3 3
3 4 6 2 3 4 6
8 7 6 8 4 5 7
1 2 3 2 1 3 2
4 5 3 2 1 2 1
```

Output

```text
3 3
```

## Thoughts

The missing formula is:

`f(x, y) = sum a[i][j] * max(0, k - abs(i - x) - abs(j - y))`

So each cell contributes inside a diamond of Manhattan radius `k - 1`.

For a fixed center `(x, y)`, split the diamond into 4 closed quadrants:

- top-left
- top-right
- bottom-left
- bottom-right

Inside one quadrant, for example `i <= x` and `j <= y`, the contribution becomes:

`a[i][j] * (k - (x - i) - (y - j))`

which can be rewritten as:

`a[i][j] * ((k - x - y) + (i + j))`

This is useful because over one quadrant we only need two aggregated values:

- the sum of `a[i][j]`
- the sum of `(i + j) * a[i][j]`

Both sums are taken over the triangular region:

`i <= x, j <= y, x - i + y - j < k`

That region is a rectangle cut by one diagonal, so it can be handled in `O(1)` per cell after preprocessing:

- column prefix sums
- diagonal prefix sums

Using one scan, we can compute the full top-left closed-quadrant contribution for every `(x, y)`. Then rotate the matrix 3 times to get the other three quadrants with the same code.

One detail: the 4 closed quadrants overlap on the horizontal axis, the vertical axis, and the center cell itself. So after summing the 4 quadrant tables, we subtract:

- the horizontal-axis contribution
- the vertical-axis contribution
- one extra copy of the center contribution

The horizontal and vertical axis parts reduce to a 1D weighted segment sum:

`sum a[t] * (k - abs(t - pos))`

which can be computed for every position using prefix sums of:

- `a[t]`
- `t * a[t]`

Overall complexity is `O(nm)`, which is fast enough for `n, m <= 1000`.