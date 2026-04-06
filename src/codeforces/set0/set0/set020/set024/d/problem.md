# D. Broken robot

You received as a gift a very clever robot walking on a rectangular board. Unfortunately, you understood that it is broken and behaves rather strangely (randomly). The board consists of `N` rows and `M` columns of cells. The robot is initially at some cell on the *i*-th row and the *j*-th column. Then at every step the robot could go to some other cell. The aim is to go to the bottommost (`N`-th) row.

The robot can stay at its current cell, move to the left, move to the right, or move to the cell below the current one. If the robot is in the leftmost column it cannot move to the left, and if it is in the rightmost column it cannot move to the right. At every step, all possible moves are equally probable.

Return the **expected number of steps** to reach the bottommost row.

## Input

On the first line you will be given two space-separated integers `N` and `M` (`1 ≤ N, M ≤ 1000`).

On the second line you will be given another two space-separated integers `i` and `j` (`1 ≤ i ≤ N`, `1 ≤ j ≤ M`) — the initial row and column. Note that `(1, 1)` is the upper-left corner of the board and `(N, M)` is the bottom-right corner.

## Output

Output the expected number of steps on a line by itself, with at least 4 digits after the decimal point.

## Examples

### Example 1

**Input**

```
10 10
10 4
```

**Output**

```
0.0000000000
```

### Example 2

**Input**

```
10 14
5 14
```

**Output**

```
18.0038068653
```


### ideas
1. 完全不知道咋做～

## detailed explanation

Let `E[r][c]` be the expected number of steps needed to reach the bottom row if the robot starts at row `r`, column `c`.

Our target is `E[i][j]`.

### 1. Base case

If `r = N`, we are already on the bottom row, so

- `E[N][c] = 0`

for every column `c`.

### 2. Write the expectation equation

For a cell that is not on the bottom row, use the standard expectation formula:

- `E = 1 + average(of all reachable next states)`

because we spend one step immediately, then continue from the next position.

The set of possible moves depends on the column:

- stay in place
- move left, if it exists
- move right, if it exists
- move down

All possible moves are equally likely.

### 3. Equations on one fixed row

Assume row `r + 1` is already known, and we now want to compute all values on row `r`.

For convenience let

- `X[c] = E[r][c]`
- `D[c] = E[r+1][c]`

Then we solve for `X[1..M]`.

#### Left border (`c = 1`)

Possible moves are:

- stay at `(r,1)`
- move right to `(r,2)`
- move down to `(r+1,1)`

So

- `X[1] = 1 + (X[1] + X[2] + D[1]) / 3`

which becomes

- `2 * X[1] - X[2] = 3 + D[1]`

#### Middle cell (`1 < c < M`)

Possible moves are:

- stay
- left
- right
- down

So

- `X[c] = 1 + (X[c-1] + X[c] + X[c+1] + D[c]) / 4`

which becomes

- `3 * X[c] - X[c-1] - X[c+1] = 4 + D[c]`

#### Right border (`c = M`)

Similarly,

- `X[M] = 1 + (X[M-1] + X[M] + D[M]) / 3`

which becomes

- `2 * X[M] - X[M-1] = 3 + D[M]`

### 4. This is a tridiagonal linear system

For one row, all unknowns are `X[1], X[2], ..., X[M]`, and each equation only involves:

- the current variable
- its left neighbor
- its right neighbor

So the system is tridiagonal and can be solved in linear time.

### 5. Linear-time elimination

We express each variable in the form

- `X[c] = a[c] * X[c+1] + b[c]`

and eliminate from left to right.

#### First column

From

- `2 * X[1] - X[2] = 3 + D[1]`

we get

- `X[1] = (1/2) * X[2] + (3 + D[1]) / 2`

so

- `a[1] = 1/2`
- `b[1] = (3 + D[1]) / 2`

#### General column

Suppose

- `X[c-1] = a[c-1] * X[c] + b[c-1]`

Substitute it into

- `3 * X[c] - X[c-1] - X[c+1] = 4 + D[c]`

Then

- `(3 - a[c-1]) * X[c] - X[c+1] = 4 + D[c] + b[c-1]`

so

- `X[c] = a[c] * X[c+1] + b[c]`

with

- `a[c] = 1 / (3 - a[c-1])`
- `b[c] = (4 + D[c] + b[c-1]) / (3 - a[c-1])`

### 6. Recover the last column and back-substitute

For the last column,

- `2 * X[M] - X[M-1] = 3 + D[M]`

and since

- `X[M-1] = a[M-1] * X[M] + b[M-1]`

we get

- `(2 - a[M-1]) * X[M] = 3 + D[M] + b[M-1]`

thus

- `X[M] = (3 + D[M] + b[M-1]) / (2 - a[M-1])`

After that, recover all previous values by going right-to-left:

- `X[c] = a[c] * X[c+1] + b[c]`

### 7. Process rows from bottom to top

The bottom row is already known:

- all zeros

Then compute row `N-1`, then `N-2`, and so on, until row `i`.

Each row costs `O(M)`, so the full complexity is:

- `O(N * M)` time
- `O(M)` memory

which easily fits the limits `N, M <= 1000`.

### 8. Special case `M = 1`

There is only one column, so the robot has only two possible moves:

- stay
- go down

each with probability `1/2`.

If the expected number of steps from one row above is `X`, then

- `X = 1 + (X + D) / 2`

which gives

- `X = D + 2`

So every remaining row contributes exactly `2` expected steps.
