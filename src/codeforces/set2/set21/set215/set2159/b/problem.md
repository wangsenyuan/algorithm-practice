# B. Rectangles

[Problem link](https://codeforces.com/problemset/problem/2159/B)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

You are given a binary grid `G` of dimensions `n × m`.

A rectangle is a tuple `(u, d, l, r)` that satisfies:

- `1 <= u < d <= n`
- `1 <= l < r <= m`
- Cells `(u, l)`, `(u, r)`, `(d, l)`, and `(d, r)` all contain a `1`

The area of rectangle `(u, d, l, r)` is `(d - u + 1) * (r - l + 1)`.

For each cell `(i, j)`, find the minimum area of any rectangle `(u, d, l, r)` such that
`u <= i <= d` and `l <= j <= r`. If no such rectangle exists, the answer is `0`.

## Constraints

- `1 <= t <= 10^4`
- `2 <= n, m` and `n * m <= 250000`
- Each cell of `G` is `0` or `1`
- Sum of `n * m` over all test cases does not exceed `250000`

## Input

The first line contains `t` — the number of test cases.

For each test case:

- The first line contains `n` and `m`
- Each of the next `n` lines contains a binary string of length `m`

## Output

For each test case, print an `n × m` grid. Entry `(i, j)` is the minimum covering
rectangle area for that cell, or `0` if none exists.

## Example

```text
Input
3
3 5
10101
10100
00101
4 6
011101
010001
100010
101110
5 5
11100
10110
11111
01101
00111

Output
6 6 6 9 9
6 6 6 9 9
0 0 9 9 9
0 10 8 8 10 10
0 10 8 8 10 10
10 10 8 8 10 0
10 10 8 8 10 0
6 6 6 0 0
6 6 4 4 0
6 4 4 4 6
0 4 4 6 6
0 0 6 6 6
```

### Note

In the first test case, the grid is:

```text
1 0 1 0 1
1 0 1 0 0
0 0 1 0 1
```

There are two rectangles: `(1, 2, 1, 3)` with area `6`, and `(1, 3, 3, 5)` with area `9`.

In the third test case, six rectangles achieve the minimum area for at least one cell:

- `(1, 3, 1, 2)` area `6`
- `(1, 2, 1, 3)` area `6`
- `(3, 4, 2, 3)` area `4`
- `(2, 3, 3, 4)` area `4`
- `(3, 5, 4, 5)` area `6`
- `(4, 5, 3, 5)` area `6`


## Solution

### Goal

For every cell `(i, j)`, compute the minimum area of a valid corner-rectangle that
covers it. A valid rectangle `(u, d, l, r)` needs four `1`s at its corners; its area
is height `×` width = `(d - u + 1) * (r - l + 1)`.

Naively trying all rectangles is too slow. The key is that we do **not** need every
rectangle — only a special subset that still preserves the minimum area for every
cell.

### Exhaustive minimal set

Fix the top and bottom rows `u < d`. Among all columns where both `G[u][c]` and
`G[d][c]` are `1`, list them as `c1 < c2 < ... < ck`.

Any valid rectangle with these top/bottom rows must use two of these columns as
left/right borders. For a cell covered by some wide rectangle
`(u, d, c_a, c_b)`, it is also covered by a **narrower consecutive** rectangle
`(u, d, c_j, c_{j+1})` that still contains that cell, and the narrower one has
smaller (or equal) area.

So for fixed `(u, d)`, the only rectangles we need are consecutive pairs:

```text
(u, d, c_j, c_{j+1})   for j = 1 .. k-1
```

These form the **exhaustive minimal set** `S`: enough to achieve the true minimum
area for every covered cell, and `|S|` is `O(n^2 m)` in the worst case, but the
total width sum over `S` is also `O(n^2 m)`.

### Complexity trick: transpose

Enumerating all row pairs costs `O(n^2 m)`. Since `n * m <= 250000`, if we ensure
`n <= m` then `n <= sqrt(n*m) <= ~500`, and

```text
O(n^2 m) = O(n * (n*m)) = O(n * 250000) = O(nm * min(n, m))
```

which is acceptable. If `n > m`, transpose the grid, solve, then transpose the
answer back.

### Computing answers without storing all rectangles

For a fixed top row `top`, scan bottom rows `bot = top+1 .. n-1`.

For each `(top, bot)`:

1. Build `left[j]` = nearest column `<= j` where both rows are `1`
   (1-based index, or `0` if none).
2. Build `right[j]` = nearest column `> j` where both rows are `1`
   (1-based index, or `0` if none).

Then for each column `j`, the smallest rectangle of height
`h = bot - top + 1` that covers column `j` is:

- If column `j` itself is a `1`-column (both rows `1`):
  - pair with previous `1`-column: width `j - left[j-1] + 1`
  - pair with next `1`-column: width `right[j] - j`
- Otherwise, if both a left and a right `1`-column exist:
  - the consecutive pair spanning `j` has width `right[j] - left[j] + 1`

Store that area into `smin[bot][j]`.

After finishing all bottoms for this fixed `top`, propagate upward per column:

```text
smin[r][j] = min(smin[r][j], smin[r+1][j])   for r = n-2 .. top
ans[r][j]  = min(ans[r][j],  smin[r][j])
```

Why this works: any rectangle with top = `top` and bottom `>= r` covers row `r`.
Taking the min over larger bottoms (via the suffix min) gives the best area among
all such rectangles that cover cell `(r, j)`.

Repeating for every `top` and taking the global min over all `top` yields the
final answer for every cell. Unreachable cells stay `inf` and become `0`.

### Walkthrough on sample 1

```text
1 0 1 0 1
1 0 1 0 0
0 0 1 0 1
```

Row pair `(0, 1)`: `1`-columns are `{0, 2}` → one consecutive rectangle
`(0,1,0,2)` with area `2 * 3 = 6`. It covers columns `0..2` on rows `0..1`.

Row pair `(0, 2)`: `1`-columns are `{2, 4}` → rectangle `(0,2,2,4)` with area
`3 * 3 = 9`. It covers columns `2..4` on rows `0..2`.

Row pair `(1, 2)`: only one shared `1`-column `{2}` → no rectangle.

So answers become:

```text
6 6 6 9 9
6 6 6 9 9
0 0 9 9 9
```

### Complexity

- Time: `O(nm * min(n, m))`
- Memory: `O(nm)` for the answer / temporary arrays

Reference: [Codeforces Round 1058 Editorial](https://codeforces.com/blog/entry/147322)
