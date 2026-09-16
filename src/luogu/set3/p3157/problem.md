# P3157 [CQOI2011] 动态逆序对

[Problem link](https://www.luogu.com.cn/problem/P3157)

time limit per test: 1.50s

memory limit per test: 500.00MB

The inversion count of a sequence `a` is the size of the set

`{(i, j) | i < j and a_i > a_j}`.

You are given a permutation of `1..n`. Then `m` values are deleted one by one. Before each deletion, report the inversion count of the current sequence.

## Input

The first line contains two integers `n` and `m`, the length of the permutation and the number of deletions.

The next `n` lines contain the initial permutation, one integer per line, each in `1..n`.

The next `m` lines contain the deleted values, one integer per line, in deletion order.

## Output

Print `m` lines. The `i`-th line is the inversion count just before the `i`-th deletion.

## Sample

### Input

```text
5 4
1
5
3
4
2
5
1
4
2
```

### Output

```text
5
2
2
1
```

### Note

Sequences just before each deletion:

- `1, 5, 3, 4, 2`
- `1, 3, 4, 2`
- `3, 4, 2`
- `3, 2`

## Constraints

For all tests, `1 <= n <= 10^5` and `1 <= m <= 5 * 10^4`.

## Status

I/O and the official sample are in place. `solve` is left as a TODO.
