# C. Unusual Product

[Problem link](https://codeforces.com/problemset/problem/405/C)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: stdin

output: stdout

Little Chris is a huge fan of linear algebra. This time he has been given homework
about the unusual square of a square matrix.

The dot product of two integer vectors `x` and `y` of size `n` is the sum of the
products of the corresponding components. The **unusual square** of an `n x n`
matrix `A` is the sum of `n` dot products in `GF(2)` (all operations modulo 2).
The `i`-th term is the dot product of the `i`-th row vector and the `i`-th column
vector of `A`.

All elements of `A` are binary (`0` or `1`). For example, for

```text
1 1 1
0 1 1
1 1 0
```

the unusual square is

```text
(1·1 + 1·0 + 1·1) + (0·1 + 1·1 + 1·0) + (1·1 + 0·1 + 0·0) = 0 + 1 + 1 = 0
```

(modulo 2).

Chris must process `q` queries. Each query is one of:

1. `1 i` — flip every bit in row `i` (`1` becomes `0`, `0` becomes `1`);
2. `2 i` — flip every bit in column `i`;
3. `3` — find the unusual square of the current matrix.

Given the initial matrix, output the answers for all queries of type `3`.

## Input

The first line contains an integer `n` (`1 <= n <= 1000`) — the number of rows and
columns in `A`.

Each of the next `n` lines contains `n` space-separated bits `a_ij`
(`0 <= a_ij <= 1`) — the `i`-th row of `A`.

The next line contains an integer `q` (`1 <= q <= 10^6`) — the number of queries.

Each of the next `q` lines is one query:

- `1 i` — flip row `i` (`1 <= i <= n`);
- `2 i` — flip column `i` (`1 <= i <= n`);
- `3` — output the unusual square.

Since the input and output can be very large, use fast I/O. In C++, avoid `cin` /
`cout`.

## Output

Let `m` be the number of type-3 queries. Print one string `s` of length `m`, where
the `i`-th character is the unusual square for the `i`-th type-3 query in input
order.

## Example

### Input

```text
3
1 1 1
0 1 1
1 0 0
12
3
3
2 2
2 2
1 3
3
3
1 2
2 1
1 1
3
```

### Output

```text
01001
```
