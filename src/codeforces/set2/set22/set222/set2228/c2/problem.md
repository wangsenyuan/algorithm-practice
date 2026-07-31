# C2. Cirno and Number (Hard Version)

[Problem link](https://codeforces.com/problemset/problem/2228/C2)

**Contest:** [Codeforces contest 2228](https://codeforces.com/contest/2228)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem Statement

This is the hard version of the problem. The difference between the versions is
that in this version, `1 <= n <= 10`. You can hack only if you solved all
versions of this problem.

You are given a non-negative integer `a` and a non-empty, strictly increasing
sequence of digits `d` of length `n`, where `0 <= d_i <= 9`.

Find the minimum value of `|a - b|` over all non-negative integers `b` whose
decimal representation contains only digits from `d`.

## Input

Each test contains multiple test cases. The first line contains the number of
test cases `t` (`1 <= t <= 10^4`).

Each test case:

- the first line contains two integers `a` and `n`
  (`0 <= a <= 10^17`, `1 <= n <= 10`);
- the second line contains `n` integers `d_1, d_2, ..., d_n`. It is guaranteed
  that `0 <= d_1 < d_2 < ... < d_n <= 9`.

## Output

For each test case, output the minimum value of `|a - b|`.

## Sample Input 1

```text
4
0 1
0
11 2
1 2
222 3
3 4 5
3333 4
6 7 8 9
```

## Sample Output 1

```text
0
0
111
2334
```

## Note

- Test 1: `a = 0`, `b = 0`, `|a - b| = 0`.
- Test 2: `a = 11`, `b = 11`, `|a - b| = 0`.
- Test 3: `a = 222`, `b = 333`, `|a - b| = 111`.
- Test 4: `a = 3333`, `b = 999`, `|a - b| = 2334`.
