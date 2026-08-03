# D. An Alternative Way

[Problem link](https://codeforces.com/problemset/problem/2241/D)

**Contest:** [Codeforces Round 1107 (Div. 3)](https://codeforces.com/contest/2241)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

You are given two arrays `a` and `b`, each of length `n`. You are allowed to
perform the following operation on array `a` any number of times (including zero):

1. Choose two indices `l` and `r` such that `1 ≤ l ≤ r ≤ n`.
2. For each index `i` from `l` to `r` (both inclusive):
   - Set `a_i := a_i - 1` if `i - l` is odd.
   - Set `a_i := a_i + 1` if `i - l` is even.

Determine whether you can make the array `a` equal to the array `b` by performing
the operation any number of times.

## Constraints

- `1 ≤ t ≤ 10^4`
- `1 ≤ n ≤ 2 · 10^5`
- `1 ≤ a_i, b_i ≤ 10^9`
- Sum of `n` over all test cases ≤ `2 · 10^5`

## Input

The first line contains a single integer `t` — the number of test cases.

Each test case consists of:

- a single integer `n` — the length of the arrays `a` and `b`
- `n` integers `a_1, a_2, ..., a_n`
- `n` integers `b_1, b_2, ..., b_n`

## Output

For each test case, print `YES` if you can make array `a` equal to array `b` and
`NO` otherwise.

You can output `YES` and `NO` in any case.

## Samples

### Sample 1

Input:

```
7
3
1 2 3
1 2 3
4
1 4 5 2
1 5 4 3
1
9
8
6
6 7 6 7 6 7
7 6 7 6 7 6
9
9 8 7 6 5 4 3 2 1
9 9 8 2 4 4 3 5 3
3
1 1 2
2 1 1
2
1 2
1 1
```

Output:

```
YES
YES
NO
YES
NO
YES
NO
```

### Notes

- First test case: arrays `a` and `b` are already equal.
- Second test case: choose `l = 2`, `r = 4`. Then
  `a_2 := 4 + 1 = 5`, `a_3 := 5 - 1 = 4`, `a_4 := 2 + 1 = 3`, so
  `a = [1, 5, 4, 3] = b`.
- Third test case: it is impossible to make `a` equal to `b`.
