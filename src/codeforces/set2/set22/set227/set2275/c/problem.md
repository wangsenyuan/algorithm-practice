# C. Unrequited Love

[Problem link](https://codeforces.com/problemset/problem/2275/C)

**Contest:** [Codeforces Round 1125 (Div. 3)](https://codeforces.com/contest/2275)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

A triad starting at index `x` uses the positions `x`, `x+2`, and `x+4`. Its audience-love value is

```text
a_x + a_{x+2} - a_{x+4}
```

Each index can be used by at most one triad. Count the unordered pairs of distinct triads that share no index and have the same love value.

## Constraints

- `1 ≤ t ≤ 10^4`
- `6 ≤ n ≤ 2 · 10^5`
- `-10^4 ≤ a_i ≤ 10^4`
- The sum of `n` over all test cases does not exceed `2 · 10^5`

## Input

The first line contains `t`, the number of test cases.

Each test case is two lines:

```text
n
a_1 a_2 ... a_n
```

## Output

For each test case, print the number of pairs. The answer may exceed 32 bits.

## Sample Input 1

```text
6
5 1 4 2 2 -4
```

## Sample Output 1

```text
1
```

`5 + 4 - 2 = 1 + 2 - (-4)`.

## Sample Input 2

```text
7
0 0 0 0 0 0 0
```

## Sample Output 2

```text
2
```

Every triad has love `0`.

## Sample Input 3

```text
12
-4 3 0 -4 -4 -1 -4 1 -3 -2 -2 -1
```

## Sample Output 3

```text
5
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
