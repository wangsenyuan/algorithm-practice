# F. Tea Blend

[Problem link](https://codeforces.com/problemset/problem/2275/F)

**Contest:** [Codeforces Round 1125 (Div. 3)](https://codeforces.com/contest/2275)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

## Problem

Recently, K1o0n bought `n` packs of tea with strengths `a_1, ..., a_n`.

The packs are arranged in the order they were bought. For any `i` and `j`,
he brews a blend `(i, j)` by pouring the first `j` packs into the teapot,
then adding one more spoon from pack `i`. Its strength is

```text
F(i, j) = a_i * (a_1 * a_2 * ... * a_j)
```

A blend is **capricious** if the number of natural divisors of `F(i, j)` is odd.
Count the number of pairs `(i, j)` with `1 <= i <= n` and `1 <= j <= n`
that are capricious.

## Constraints

- `1 <= t <= 10^4`
- `1 <= n <= 2 * 10^5`
- `1 <= a_i <= 10^6`
- The sum of `n` over all test cases does not exceed `2 * 10^5`

## Input

The first line contains `t`, the number of test cases.

Each test case contains:
- The integer `n`
- `n` integers `a_1, ..., a_n`

## Output

For each test case, print the number of capricious blends.
The answer may exceed 32 bits.

## Sample

```text
Input
4
4
2 3 6 4
4
974153 980051 988027 1
8
12 3 5 45 1 7 63 2
1
13

Output
4
4
11
1
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
