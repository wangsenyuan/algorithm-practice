# F - Rectangle GCD

[Problem link](https://atcoder.jp/contests/abc254/tasks/abc254_f)

**Contest:** [AtCoder Beginner Contest 254](https://atcoder.jp/contests/abc254)

time limit: 2 sec

memory limit: 1024 MiB

score: 500 points

You are given a positive integer `N` and two length-`N` positive integer sequences
`A = (A_1, ..., A_N)` and `B = (B_1, ..., B_N)`.

Consider an `N × N` grid where cell `(i, j)` contains `C_{i,j} = A_i + B_j`.

Process `Q` queries. Each query gives `h_1, h_2, w_1, w_2`
(`1 <= h_1 <= h_2 <= N`, `1 <= w_1 <= w_2 <= N`). Find the GCD of all values in the rectangle
whose top-left corner is `(h_1, w_1)` and bottom-right corner is `(h_2, w_2)`.

## Constraints

- `1 <= N, Q <= 2 * 10^5`
- `1 <= A_i, B_i <= 10^9`
- `1 <= h_1 <= h_2 <= N`
- `1 <= w_1 <= w_2 <= N`
- All input values are integers

## Input

```text
N Q
A_1 A_2 ... A_N
B_1 B_2 ... B_N
query_1
query_2
...
query_Q
```

Each query:

```text
h_1 h_2 w_1 w_2
```

## Output

Print `Q` lines. The `i`-th line is the answer to the `i`-th query.

## Sample Input 1

```text
3 5
3 5 2
8 1 3
1 2 2 3
1 3 1 3
1 1 1 1
2 2 2 2
3 3 1 1
```

## Sample Output 1

```text
2
1
11
6
10
```

For the first query, the rectangle contains `C_{1,2}=4`, `C_{1,3}=6`, `C_{2,2}=6`, `C_{2,3}=8`, so
the answer is `gcd(4,6,6,8)=2`.

## Sample Input 2

```text
1 1
9
100
1 1 1 1
```

## Sample Output 2

```text
109
```

## Solution

For a query rectangle `h1..h2` and `w1..w2`, every cell is:

```text
C[i][j] = A[i] + B[j]
```

Start from the top-left value:

```text
base = A[h1] + B[w1]
```

All other values in the rectangle can be compared with this base.

For the same row and different columns:

```text
(A[h1] + B[j]) - (A[h1] + B[j-1]) = B[j] - B[j-1]
```

For the same column and different rows:

```text
(A[i] + B[w1]) - (A[i-1] + B[w1]) = A[i] - A[i-1]
```

Using the identity `gcd(x, y) = gcd(x, y - x)`, the GCD of the whole rectangle is:

```text
gcd(
    A[h1] + B[w1],
    gcd of |A[i] - A[i-1]| for h1 < i <= h2,
    gcd of |B[j] - B[j-1]| for w1 < j <= w2
)
```

So the 2D rectangle query becomes two 1D range-GCD queries over adjacent differences.

### Sparse tables

Build two difference arrays:

```text
DA[i] = |A[i] - A[i-1]|  for i >= 2
DB[i] = |B[i] - B[i-1]|  for i >= 2
```

The implementation stores them in sparse tables:

- `dp[k][i]` = GCD of `DA[i .. i + 2^k - 1]`
- `fp[k][i]` = GCD of `DB[i .. i + 2^k - 1]`

The transition is:

```text
dp[k][i] = gcd(dp[k-1][i], dp[k-1][i + 2^(k-1)])
fp[k][i] = gcd(fp[k-1][i], fp[k-1][i + 2^(k-1)])
```

For a query interval of length `len`, choose `k = floor(log2(len))` and combine the two overlapping
blocks:

```text
gcd(table[k][l], table[k][r - 2^k + 1])
```

This works for GCD because duplicated elements do not change the result.

### Answering a query

For query `(h1, h2, w1, w2)`:

1. Start with `base = A[h1] + B[w1]`.
2. If `h1 < h2`, query the sparse table for `DA[h1+1 .. h2]`.
3. If `w1 < w2`, query the sparse table for `DB[w1+1 .. w2]`.
4. Return the GCD of these values.

If the row range or column range has length one, the corresponding difference range is empty, and
its GCD contribution is `0`.

### Correctness

All cells in the rectangle can be reached from the top-left cell by moving along rows and columns.
Each horizontal move changes the value by a neighboring difference from `B`, and each vertical move
changes the value by a neighboring difference from `A`. Therefore every cell value is congruent to
`base` modulo the GCD of the selected row and column differences, so the rectangle GCD must divide
the formula above.

Conversely, the rectangle contains `base`, and it also contains adjacent pairs whose differences are
exactly the selected `A` and `B` differences. Any common divisor of all rectangle values must divide
`base` and all those adjacent differences. Therefore no larger value than the formula can be the
GCD. The formula is exactly the rectangle GCD.

### Complexity

Building the two sparse tables takes `O(N log N)` time and memory. Each query uses `O(1)` sparse-table
range queries, so it is answered in `O(1)` time.
