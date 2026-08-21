# F. Whiplash

[Problem link](https://codeforces.com/problemset/problem/2254/F)

**Contest:** [Codeforces Round 1114 (Div. 3)](https://codeforces.com/contest/2254)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

## Problem

You are given an even integer `n` and two arrays `a` and `b`, each containing
`n` non-negative integers. You may repeat the following operation on `a` any
number of times:

1. choose an index `i`;
2. replace every `a_j` with `a_j xor a_i` for `j != i`;
3. leave `a_i` unchanged.

Determine whether a finite sequence of operations can transform `a` into `b`.

## Input

The first line contains the number of test cases `t` (`1 <= t <= 10^4`).

For each test case:

- the first line contains an even integer `n` (`2 <= n <= 2 * 10^5`);
- the next line contains `n` integers `a_i` (`0 <= a_i < 2^30`);
- the next line contains `n` integers `b_i` (`0 <= b_i < 2^30`).

The sum of `n` over all test cases is at most `2 * 10^5`.

## Output

For each test case, print `YES` if the transformation is possible; otherwise
print `NO`. Letter case does not matter.

## Sample

```text
Input
6
2
1 2
1 0
4
1 2 4 7
6 7 5 3
4
1 2 4 8
8 4 2 1
4
1 2 3 4
1 2 4 5
4
1 2 0 3
3 3 0 3
6
3 5 6 9 10 12
6 5 3 12 15 9

Output
NO
YES
YES
NO
NO
YES
```

## Notes

