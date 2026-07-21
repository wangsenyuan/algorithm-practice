# C. Firewood

[Problem link](https://codeforces.com/problemset/problem/1958/C)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

## Problem

It's pretty cold in Berland. Monocarp has a big log of wood weighing `2^n`
grams. He needs to burn `k` grams today and keep the remaining `2^n - k` grams
for tomorrow.

In one minute, Monocarp can split one log in half. If a log weighs `x`, each
resulting log weighs `x/2`. He cannot split a log of weight `1`.

Initially he has one log of weight `2^n`. He must cut so that some of the
resulting logs weigh exactly `k` grams in total. Find the minimum number of
minutes needed.

## Constraints

- `1 <= t <= 10^4`
- `1 <= n <= 60`
- `1 <= k <= 2^n - 1`

## Input

The first line contains one integer `t` — the number of test cases.

Each test case consists of one line containing two integers `n` and `k`.

## Output

For each test case, print one integer — the minimum number of minutes.

## Example

```text
Input
4
2 2
2 1
10 3
50 36679020707840

Output
1
2
10
16
```

In the first test case, one cut yields two logs of weight `2`.

In the second test case, cut the log of weight `4` once, then cut one of the
resulting logs, obtaining weights `2, 1, 1`.
