# E. Prime Destruction

[Problem link](https://codeforces.com/problemset/problem/2266/E)

**Contest:** [Codeforces Round 1122 (Div. 3)](https://codeforces.com/contest/2266)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

## Problem

You are given a multiset `a` consisting of `n` positive integers. You may perform
the following operation any number of times, possibly zero:

Choose an integer `x > 1` from the multiset and a prime divisor `p` of `x`.
Remove one occurrence of `x` and add `p` copies of `x / p`.

You are also given an integer `k` (`1 <= k <= n`). Let `f(k)` be the minimum
number of operations required, starting from the original multiset, until every
integer in the multiset is at most `k`. Find `f(k)`.

## Input

The first line contains the number of test cases `t` (`1 <= t <= 10^4`).

Each test case starts with two integers `n` and `k`
(`1 <= k <= n <= 2 * 10^5`). The next line contains `n` integers
`a[i]` (`1 <= a[i] <= n`).

The sum of `n` over all test cases is at most `2 * 10^5`.

## Output

For each test case, print `f(k)` as a single integer.

## Sample

```text
Input
6
1 1
1
6 2
6 6 4 3 2 1
8 1
8 6 4 3 2 1 8 6
12 3
12 10 9 8 7 6 5 4 3 2 1 12
10 9
10 9 8 7 6 5 4 3 2 1
5 5
5 4 3 2 1

Output
0
4
25
15
1
0
```

## Status

`solve` is intentionally TODO.
