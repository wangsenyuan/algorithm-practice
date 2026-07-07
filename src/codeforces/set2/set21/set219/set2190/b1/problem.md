# B1 - Sub-RBS (Easy Version)

[Problem link](https://codeforces.com/problemset/problem/2190/B1)

**Contest:** [Codeforces Round 1073 (Div. 1)](https://codeforces.com/contest/2190)

time limit: 2 seconds per test

memory limit: 256 MB

This is the easy version of Sub-RBS. You are given a **regular bracket sequence** `s` of even
length `n`.

Sequence `a` is **better** than sequence `b` if either:

- `b` is a proper prefix of `a`, or
- at the first index where they differ, `a` has `(` and `b` has `)`.

Among all non-empty regular bracket subsequences `t` of `s`, find the maximum possible length of
`t` such that `t` is better than `s`. If no such `t` exists, print `-1`.

## Constraints

- 1 <= t <= 10^4
- 2 <= n <= 2 * 10^5, `n` is even
- `s` is a regular bracket sequence
- sum of `n` over all test cases <= 2 * 10^5
- `s` consists of `(` and `)`

## Input

The first line contains `t` — the number of test cases.

Each test case:

```text
n
s
```

## Output

For each test case, print one integer — the maximum length described above, or `-1`.

## Sample Input

```text
3
2
()
8
(()(()))
6
(())()
```

## Sample Output

```text
-1
6
-1
```

## Note

In the first test case, the only non-empty regular bracket subsequence is `t = s = ()`. Since `t`
is not better than `s`, the answer is `-1`.

In the second test case, choose `t = ((()))`. At the first differing index `i = 3`, `t[i] = '('`
and `s[i] = ')'`, so `t` is better than `s`. No longer regular bracket subsequence is better than
`s`, so the answer is `6`.
