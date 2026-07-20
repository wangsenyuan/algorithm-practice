# E. The Turtle Strikes Back

[Problem link](https://codeforces.com/problemset/problem/2194/E)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

There is an `n × m` pizza. The pleasure of the slice in row `i`, column `j`
is `a_{i,j}`. It is guaranteed that at least one `a_{i,j}` is non-negative.

Michelangelo chooses a path from `(1,1)` to `(n,m)` moving only right or down,
and eats all slices on the path. He wants to **maximize** the total pleasure
of the path.

Before that, Raphael must choose **exactly one** cell and flip the sign of its
pleasure (`a_{i,j} → -a_{i,j}`).

Raphael wants to minimize the maximum pleasure Michelangelo can then achieve.
Output that value (the minimum Michelangelo’s optimal path sum that Raphael
can force).

## Constraints

- `1 ≤ t ≤ 10^4`
- `1 ≤ n, m ≤ 10^6`
- `1 ≤ n · m ≤ 10^6`
- `-10^9 ≤ a_{i,j} ≤ 10^9`
- At least one `a_{i,j}` is non-negative
- Sum of `n · m` over all test cases ≤ `10^6`

## Input

```
t
n m
a_1,1 ... a_1,m
...
a_n,1 ... a_n,m
...
```

## Output

For each test case, print one integer — the answer Raphael can guarantee.

## Samples

### Sample 1

Input:

```
1
3 3
1 -2 3
4 -5 2
1 6 -1
```

Output:

```
3
```

Without sauce, an optimal path has pleasure `11`. If Raphael flips `(3,2)`
(`6 → -6`), Michelangelo’s best path has pleasure `3`.

### Sample 2

Input:

```
1
2 4
-1 -1 -1 1
-1 -1 -1 -1
```

Output:

```
-5
```

## Solution

Fix the cell `x = (i,j)` whose sign Raphael flips. Michelangelo's best path
after that operation is the better of two kinds of paths:

1. a path that passes through `x`;
2. a path that avoids `x`.

We can compute the best value of both kinds in `O(1)` for every possible
flipped cell after two dynamic-programming passes.

### 1. Prefix and suffix path DP

Let

- `dp[i][j]` be the maximum original sum of a path from `(0,0)` to `(i,j)`;
- `fp[i][j]` be the maximum original sum of a path from `(i,j)` to
  `(n-1,m-1)`.

Their transitions are

```text
dp[i][j] = a[i][j] + max(dp[i-1][j], dp[i][j-1])
fp[i][j] = a[i][j] + max(fp[i+1][j], fp[i][j+1])
```

with the unavailable directions omitted on the grid boundary.

Therefore, the maximum original sum of a path passing through `(i,j)` is

```text
throughOriginal(i,j) = dp[i][j] + fp[i][j] - a[i][j].
```

The subtraction is needed because both DP values contain `a[i][j]`.

### 2. Best path passing through the flipped cell

After flipping `(i,j)`, only that cell's contribution changes from
`a[i][j]` to `-a[i][j]`. The best prefix ending immediately before it and the
best suffix starting immediately after it are unchanged.

Let

```text
prefix = max(dp[i-1][j], dp[i][j-1])
suffix = max(fp[i+1][j], fp[i][j+1]).
```

Again, unavailable neighbors are omitted. Then the best path forced through
the flipped cell has value

```text
throughFlipped(i,j) = prefix + suffix - a[i][j].
```

This also handles the start and end cells by treating their missing prefix or
suffix as zero. In particular, for a `1 × 1` grid the value is simply
`-a[0][0]`.

### 3. Best path avoiding the flipped cell

Every right/down path visits exactly one cell on each anti-diagonal

```text
i + j = d.
```

Consequently, a path avoids `(i,j)` exactly when its cell on anti-diagonal
`d = i+j` is some other cell `(r,c)` on that diagonal.

For any such cell, the best original path passing through it has value

```text
dp[r][c] + fp[r][c] - a[r][c].
```

A path avoiding `(i,j)` does not use the flipped value, so its sum remains the
same as in the original grid. Thus

```text
avoid(i,j) = max(dp[r][c] + fp[r][c] - a[r][c])
             over r+c=i+j and (r,c)!=(i,j).
```

To answer this exclusion query in `O(1)`, for every anti-diagonal the code
stores its two largest path values together with their column indices. If the
largest entry belongs to column `j`, it uses the second largest; otherwise it
uses the largest. A column identifies a unique cell on one anti-diagonal.

The start and end anti-diagonals contain only one cell. Every valid path must
use that cell, so there is no avoiding-path candidate there.

### 4. Minimax result

For a fixed flipped cell `(i,j)`, Michelangelo chooses the better available
path:

```text
bestAfterFlip(i,j) = max(throughFlipped(i,j), avoid(i,j)).
```

Raphael tries every cell and chooses the minimum:

```text
answer = min(bestAfterFlip(i,j)).
```

## Correctness Proof

We prove that the algorithm returns the minimum possible value of
Michelangelo's optimal path after exactly one sign flip.

Consider any cell `x = (i,j)` selected by Raphael.

Every path from the top-left to the bottom-right either passes through `x` or
avoids `x`.

- Among paths passing through `x`, the best prefix before `x` is given by the
  appropriate predecessor `dp` value, and the best suffix after `x` is given
  by the appropriate successor `fp` value. The flipped cell contributes
  `-a[i][j]`. Therefore `throughFlipped(i,j)` is exactly the maximum sum among
  all paths passing through `x` after the flip.
- Every path visits exactly one cell on anti-diagonal `i+j`. Hence a path
  avoids `x` if and only if it passes through another cell on that diagonal.
  For each other cell `(r,c)`, `dp[r][c] + fp[r][c] - a[r][c]` is exactly the
  best original path through it. Such a path does not contain `x`, so flipping
  `x` does not change its sum. Taking the maximum over the other cells gives
  exactly `avoid(i,j)`.

The maximum of these two quantities is therefore exactly Michelangelo's
optimal result after flipping `x`. Finally, the algorithm takes this value's
minimum over every possible `x`, which is exactly Raphael's optimal choice.
Thus the returned answer is correct.

## Complexity Analysis

Let `N = n · m`.

- Time: `O(N)` for the two DP passes, diagonal preprocessing, and evaluation
  of every flipped cell.
- Space: `O(N + n + m)` for the grid, the two DP tables, and the two best
  entries of every anti-diagonal.
