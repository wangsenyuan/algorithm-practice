# D - Clouds

[Problem link](https://atcoder.jp/contests/abc434/tasks/abc434_d)

**Contest:** [Sky Inc, Programming Contest 2025 (AtCoder Beginner Contest 434)](https://atcoder.jp/contests/abc434)

time limit: 2 sec

memory limit: 1024 MiB

score: 425 points

The sky is represented by a `2000 x 2000` grid. When looking up at the sky, the cell at the `r`-th row
from the top and `c`-th column from the left is called `(r, c)`.

Currently, there are clouds `1, 2, ..., N` floating in this sky. The cell `(r, c)` is covered by
cloud `i` if and only if it satisfies `U_i <= r <= D_i` and `L_i <= c <= R_i`.

For `k = 1, 2, ..., N`, answer the following question:

- Remove only cloud `k` from the `N` clouds. At this point, there are `N - 1` clouds floating in the
  sky. How many cells are not covered by any cloud?

## Constraints

- `1 <= N <= 2 * 10^5`
- `1 <= U_i <= D_i <= 2000`
- `1 <= L_i <= R_i <= 2000`
- All input values are integers.

## Input

```text
N
U_1 D_1 L_1 R_1
U_2 D_2 L_2 R_2
...
U_N D_N L_N R_N
```

## Output

Print `N` lines. The `i`-th line should contain the answer when `k = i`.

## Sample Input 1

```text
5
2 4 1 4
3 3 3 5
1 3 4 6
4 5 3 5
5 5 4 6
```

## Sample Output 1

```text
3999983
3999976
3999982
3999978
3999977
```

### Note

The total number of cells is `2000 * 2000 = 4000000`.

- When cloud 1 is removed, 17 cells remain covered by some cloud, so the answer is `3999983`.
- When cloud 2 is removed, 24 cells remain covered, so the answer is `3999976`.
- When cloud 3 is removed, 18 cells remain covered, so the answer is `3999982`.
- When cloud 4 is removed, 22 cells remain covered, so the answer is `3999978`.
- When cloud 5 is removed, 23 cells remain covered, so the answer is `3999977`.

## Solution

For each cloud `k`, after removing it, the newly uncovered cells are exactly the cells that were
covered by **only cloud `k`** before removal.

So if:

- `empty` is the number of cells covered by no cloud initially;
- `only[k]` is the number of cells covered by exactly cloud `k`;

then the answer for removing cloud `k` is:

```text
empty + only[k]
```

The grid size is fixed at `2000 x 2000`, so we can process all cells directly after doing rectangle
updates with a 2D difference array.

### Encoding the Covering Cloud

Normally, a difference array would add `1` for each rectangle and later tell us only the number of
clouds covering each cell. Here we also need to know **which** cloud is the unique covering cloud
when the coverage count is exactly `1`.

The implementation assigns cloud `k` the value:

```text
N + k
```

where `k` is zero-based. For every cell covered by cloud `k`, this value is added through a 2D imos
rectangle update.

After taking prefix sums:

- value `0` means no cloud covers the cell;
- a value in `[N, 2N)` means exactly one cloud covers the cell;
- any value `>= 2N` means at least two clouds cover the cell.

Why does this work?

Each single cloud contributes one value between `N` and `2N - 1`. The smallest possible sum of two
different clouds is:

```text
N + (N + 1) = 2N + 1
```

So no cell covered by two or more clouds can have a value less than `2N`. Therefore, if the final
value is in `[N, 2N)`, it must be covered by exactly one cloud, and its id is:

```text
id = value - N
```

### 2D Difference Array

For a rectangle:

```text
U <= row <= D
L <= col <= R
```

with added value `v`, use the usual 2D imos updates:

```text
diff[U][L]       += v
diff[U][R + 1]   -= v
diff[D + 1][L]   -= v
diff[D + 1][R+1] += v
```

The code stores one extra border, so `R + 1` and `D + 1` are safe.

Then scan the grid and compute prefix sums:

```text
diff[i][j] += diff[i][j - 1] + diff[i - 1][j] - diff[i - 1][j - 1]
```

For each real cell:

- if the value is `0`, increment `empty`;
- if the value is in `[N, 2N)`, increment `only[value - N]`.

Finally print:

```text
empty + only[k]
```

for every cloud `k`.

### Correctness

A cell is uncovered after removing cloud `k` in exactly two cases:

1. it was already uncovered before removing anything;
2. it was covered by cloud `k` and by no other cloud.

The first case contributes `empty` to every answer. The second case contributes `only[k]` only to
the answer for cloud `k`.

The 2D difference array applies each cloud's encoded value to exactly the cells in its rectangle, so
after prefix sums each cell contains the sum of encoded values of all clouds covering it. The encoding
distinguishes uncovered cells, uniquely covered cells, and multiply covered cells as described above,
so `empty` and every `only[k]` are counted correctly.

Thus `empty + only[k]` is exactly the number of cells not covered after removing cloud `k`.

### Complexity

There are `2000 * 2000` grid cells. Each cloud adds one rectangle update, and then the grid is scanned
once.

Total complexity is `O(N + 2000^2)`, and memory usage is `O(2000^2 + N)`.
