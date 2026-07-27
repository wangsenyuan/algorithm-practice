# B. Arpa’s obvious problem and Mehrdad’s terrible solution

[Problem link](https://codeforces.com/problemset/problem/742/B)

**Contest:** [Codeforces Round #383 (Div. 2)](https://codeforces.com/contest/742)

time limit: 1 second

memory limit: 256 megabytes

## Problem

Given an array `a` of `n` integers and an integer `x`, count the number of pairs
of indices `i, j` (`1 <= i < j <= n`) such that `a[i] XOR a[j] = x`.

## Constraints

- `1 <= n <= 10^5`
- `0 <= x <= 10^5`
- `1 <= a[i] <= 10^5`

## Input

```text
n x
a1 a2 ... an
```

## Output

Print a single integer — the number of pairs.

## Sample 1

```text
Input
2 3
1 2

Output
1
```

Only pair `(1, 2)`: `1 XOR 2 = 3`.

## Sample 2

```text
Input
6 1
5 1 2 3 4 1

Output
2
```

Pairs `(3, 4)` (`2 XOR 3 = 1`) and `(1, 5)` (`5 XOR 4 = 1`).
