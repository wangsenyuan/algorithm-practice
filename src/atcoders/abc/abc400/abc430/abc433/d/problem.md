# D - 183183 (ABC433)

**Contest:** [ABC433](https://atcoder.jp/contests/abc433) — AtCoder Beginner Contest 433  
**Task:** [https://atcoder.jp/contests/abc433/tasks/abc433_d](https://atcoder.jp/contests/abc433/tasks/abc433_d)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 400 points

## Problem Statement

For positive integers `x` and `y`, define `f(x, y)` as follows:

- Write `x` and `y` in decimal notation without leading zeros, concatenate them
  in this order to form a string `S`, and interpret `S` as a decimal integer.

For example, `f(12, 3) = 123` and `f(100, 40) = 10040`.

You are given positive integers `N`, `M` and a sequence of `N` positive integers
`A = (A_1, A_2, ..., A_N)`.

Find the number of pairs of integers `(i, j)` that satisfy all of the following:

- `1 <= i, j <= N`
- `f(A_i, A_j)` is a multiple of `M`

## Constraints

- `1 <= N <= 2 * 10^5`
- `1 <= M <= 10^9`
- `1 <= A_i <= 10^9`
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```text
N M
A_1 A_2 ... A_N
```

## Output

Output the number of pairs `(i, j)` that satisfy all the conditions.

## Sample Input 1

```text
2 11
2 42
```

## Sample Output 1

```text
2
```

- `(i, j) = (1, 1)`: `f(A_1, A_1) = 22` is a multiple of `11`.
- `(i, j) = (1, 2)`: `f(A_1, A_2) = 242` is a multiple of `11`.
- `(i, j) = (2, 1)`: `f(A_2, A_1) = 422` is not a multiple of `11`.
- `(i, j) = (2, 2)`: `f(A_2, A_2) = 4242` is not a multiple of `11`.

The answer is `2`.

## Sample Input 2

```text
4 7
2 8 16 183
```

## Sample Output 2

```text
4
```

## Sample Input 3

```text
5 5
1000000000 1000000000 1000000000 1000000000 1000000000
```

## Sample Output 3

```text
25
```

## Sample Input 4

```text
12 13
80 68 862370 82217 8 56 5 168 672624 6 286057 11864
```

## Sample Output 4

```text
10
```
