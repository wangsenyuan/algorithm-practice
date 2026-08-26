# D. Silhouette

[Problem link](https://codeforces.com/problemset/problem/2254/D)

**Contest:** [Codeforces Round 1114 (Div. 3)](https://codeforces.com/contest/2254)

time limit per test: 2.5 seconds

memory limit per test: 256 megabytes

## Problem

There is a secret array `a` of `n` strictly positive integers. The shadow of
each element is the sum of every array element that is strictly smaller:

```text
b[i] = sum(a[j]) for all j with a[j] < a[i]
```

Given the shadow array `b`, reconstruct the lexicographically smallest valid
positive-integer array `a`. If no such array exists, print `-1`.

## Input

The first line contains `t` (`1 <= t <= 10^4`), the number of test cases.

For each test case:

- the first line contains `n` (`1 <= n <= 2 * 10^5`);
- the second line contains `n` integers `b[i]` (`0 <= b[i] <= 2 * 10^14`).

The sum of `n` across test cases is at most `2 * 10^5`.

## Output

For each test case, print the lexicographically smallest valid array `a`,
whose elements are between `1` and `10^18`, or `-1` if reconstruction is
impossible.

## Sample

```text
Input
8
1
0
5
0 4 0 4 14
3
4 0 0
3
0 0 0
3
0 1 1
4
1 1 1 1
7
0 4 4 4 4 4 9
5
0 0 0 3 3

Output
1
2 5 2 5 6
3 2 2
1 1 1
1 2 2
-1
-1
1 1 1 2 2
```

## ideas
1. 这个题目还挺难的
2. 假设b是排序好的, 