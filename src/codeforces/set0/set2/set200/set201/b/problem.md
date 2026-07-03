# B. Guess That Car!

[Problem link](https://codeforces.com/problemset/problem/201/B)

**Contest:** [Codeforces Beta Round #12 (Div. 2)](https://codeforces.com/contest/201)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

The parking lot is an `n × m` grid of squares. Square `(i, j)` has `1 <= i <= n` (north to south) and
`1 <= j <= m` (west to east). Each square contains one car with rarity coefficient `c_{i,j}`.

Yura stands on a junction of grid lines at coordinates `(l_i, l_j)` where `0 <= l_i <= n` and
`0 <= l_j <= m`. The distance from `(l_i, l_j)` to the center of square `(i, j)` is the Euclidean
distance. Guessing the car in that square costs `c_{i,j} · d^2`.

Choose a starting junction to minimize the total guessing time for all cars. If multiple junctions
tie, pick the smallest `l_i`; if still tied, pick the smallest `l_j`.

## Input

The first line contains `n` and `m` (`1 <= n, m <= 1000`).

Each of the next `n` lines contains `m` integers `c_{i,j}` (`0 <= c_{i,j} <= 100000`).

## Output

Print two lines:

1. The minimum total time
2. Optimal `l_i` and `l_j`

## Example

### Input

```text
2 3
3 4 5
3 9 1
```

### Output

```text
392
1 1
```

### Input

```text
3 4
1 0 0 0
0 0 3 0
0 0 5 5
```

### Output

```text
240
2 3
```

### Note

In the first example, with junction `(1, 1)` the total is
`3·8 + 3·8 + 4·8 + 9·8 + 5·40 + 1·40 = 392`.

## Solution

Let the junction be `(r, c)` with `0 <= r <= n` and `0 <= c <= m`. The total cost is

```text
sum_{i,j} c[i,j] * ((i - r)^2 + (j - c)^2)
```

Expand and regroup:

```text
sum_i (sum_j c[i,j]) * (i - r)^2  +  sum_j (sum_i c[i,j]) * (j - c)^2
```

Define row weights `row[i] = sum_j c[i,j]` and column weights `col[j] = sum_i c[i,j]`. The cost
splits into two independent 1D problems:

```text
f(r) = sum_i row[i] * (i - r)^2
g(c) = sum_j col[j] * (j - c)^2
answer = min f(r) + min g(c)
```

So the optimal row index and optimal column index can be found separately.

### Integer coordinates

To avoid floating point, place junction `(r, c)` at `(4r, 4c)` and the center of cell `(i, j)` at
`(4i + 2, 4j + 2)`. Then

```text
(i - r)^2 + (j - c)^2  =  ((4r - 4i - 2)^2 + (4c - 4j - 2)^2) / 16
```

The denominator is the same for every cell, so minimizing the total cost is equivalent to
minimizing the integer sums using `(4r - 4i - 2)^2` and `(4c - 4j - 2)^2`.

### Algorithm

1. Compute `row[i]` and `col[j]`.
2. For each `r = 0..n`, compute `f(r) = sum_i row[i] * (4r - 4i - 2)^2`; keep the smallest value
   and the smallest `r` on ties.
3. For each `c = 0..m`, compute `g(c)` similarly.
4. Output `f(r0) + g(c0)` and `(r0, c0)`.

### Correctness

The objective is separable because cross terms never appear after expanding
`c[i,j] * ((i-r)^2 + (j-c)^2)`. Therefore the global minimum is obtained by independently minimizing
the row part and the column part. The integer scaling is a common factor on every term, so it does
not change the argmin.

Tie-breaking is handled by scanning `r` and `c` in increasing order and updating only on a strictly
smaller cost.

### Complexity

Preprocessing the weights takes `O(n m)`. Each of the two 1D scans is `O(n^2 + m^2)` in the worst
case, which is fine for `n, m <= 1000`.