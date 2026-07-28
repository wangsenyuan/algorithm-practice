# C. Below the Diagonal

[Problem link](https://codeforces.com/problemset/problem/266/C)

**Contest:** [Codeforces Round #163 (Div. 2)](https://codeforces.com/contest/266)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

You are given a square matrix of size `n × n`. Exactly `n - 1` cells contain
ones; the rest contain zeros. You may apply these operations:

1. Swap the `i`-th and `j`-th rows.
2. Swap the `i`-th and `j`-th columns.

Transform the matrix so that every one lies **strictly below** the main
diagonal: a cell `(i, j)` is below the diagonal iff `i > j`.

You do not need to minimize the number of operations, but it must not exceed
`10^5`. Any valid sequence is accepted.

## Constraints

- `2 <= n <= 1000`
- Exactly `n - 1` ones; all positions are distinct
- `1 <= x_k, y_k <= n`
- Number of operations `m` must satisfy `0 <= m <= 10^5`

## Input

The first line contains `n`.

Then follow `n - 1` lines, each with two integers `x_k y_k` — the row and column
of a cell that contains a one.

## Output

Print `m` — the number of operations.

Then print `m` lines, each with three integers `t i j`:

- `t = 1` — swap rows `i` and `j`
- `t = 2` — swap columns `i` and `j`
- `i ≠ j`

## Sample 1

```text
Input
2
1 2

Output
2
2 1 2
1 1 2
```

## Sample 2

```text
Input
3
3 1
1 3

Output
3
2 2 3
1 1 3
1 1 2
```

## Sample 3

```text
Input
3
2 1
3 2

Output
0
```

## ideas
1. 看起来是肯定有答案的. 要找到一个策略, 维持某个不变性;
2. 感觉这里有个图, swap(r1, r2), swap(r2, r3) 这样子貌似是有意义的
3. 这个时候, r1行, 包含的是r2行的数据, r2行包含的是r3行的数据, r3包含的是r1行的数据
4. 任何两个cell, 如果一开始就不在同一行, 那么永远也不会在同一行. 
5. 