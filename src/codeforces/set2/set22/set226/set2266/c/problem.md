# C. AND, OR, Sort!

[Problem link](https://codeforces.com/problemset/problem/2266/C)

**Contest:** [Codeforces Round 1122 (Div. 3)](https://codeforces.com/contest/2266)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

## Problem

Given a binary string `s`, an operation chooses an index `i` and replaces
`s[i]` with either the bitwise AND or bitwise OR of the prefix `s[0..i]`.
Find the minimum number of operations needed to make `s` non-decreasing.

## Input

The first line contains `t` test cases. Each test case contains `n` followed
by a binary string `s` of length `n`. The total length across all test cases
is at most `2 * 10^5`.

## Output

For each test case, print the minimum number of operations.

## Sample

```text
Input
6
4
0011
4
1000
5
01000
8
01001101
7
0101010
7
0111101

Output
0
3
1
2
3
1
```

## Status

`solve` is intentionally TODO. The sample tests are retained and skipped until
the algorithm is implemented.
