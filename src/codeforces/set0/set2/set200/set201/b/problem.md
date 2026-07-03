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

### ideas
1. 
