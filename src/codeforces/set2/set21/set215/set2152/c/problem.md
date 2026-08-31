# C. Triple Removal

[Problem link](https://codeforces.com/problemset/problem/2152/C)

**Contest:** [Squarepoint Challenge (Codeforces Round 1055, Div. 1 + Div. 2)](https://codeforces.com/contest/2152)

time limit per test: 2 seconds

memory limit per test: 1024 megabytes

## Problem

For a binary array, one operation chooses three equal elements at positions
`i < j < k`, removes them, and costs `min(k - j, j - i)` using the current
array indices. The cost of an array is the minimum total cost needed to remove
all its elements; it is `-1` when that cannot be done.

Given a binary array `a`, answer independent range queries: for every `[l, r]`,
find the cost of subarray `a[l..r]`.

## Input

The first line contains the number of test cases `t` (`1 <= t <= 10^4`).

For each test case:

- the first line contains `n` and `q` (`1 <= n, q <= 250000`);
- the next line contains `n` binary values;
- each of the next `q` lines contains a range `l, r` (`1 <= l <= r <= n`).

The sums of `n` and of `q` over all test cases are each at most `250000`.

## Output

For every query, print the corresponding minimum cost, or `-1` if the subarray
cannot be made empty.

## Sample

```text
Input
2
12 4
0 0 1 1 0 1 0 1 0 1 1 0
1 12
2 7
5 10
6 11
6 3
0 0 0 1 1 1
1 3
4 6
1 6

Output
4
2
3
-1
1
1
2
```

## Status

`solve` is intentionally left as a TODO. The tests preserve the official cases
and are skipped until an implementation is added.
