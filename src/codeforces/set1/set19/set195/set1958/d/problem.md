# D. Staircase

[Problem link](https://codeforces.com/problemset/problem/1958/D)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

## Problem

There is a staircase consisting of `n` steps. Each step is either intact or
broken. For each broken step, an integer `a_i` is given denoting the difficulty
of repairing it.

Every day, you can either:

- repair an arbitrary broken step with effort `a_i`; or
- repair two adjacent broken steps `i` and `i+1` with effort
  `2 * (a_i + a_{i+1})`.

You want to repair all broken steps in the **minimum possible number of days**.
Among all such schedules, find the **minimum total effort**.

## Constraints

- `1 <= t <= 10^4`
- `1 <= n <= 3 * 10^5`
- `0 <= a_i <= 10^8`
- If `a_i = 0`, the `i`-th step does not need repair; otherwise it is broken
- Sum of `n` over all test cases does not exceed `3 * 10^5`

## Input

The first line contains one integer `t` — the number of test cases.

Each test case consists of two lines:

- the first line contains one integer `n`
- the second line contains `n` integers `a_1, a_2, ..., a_n`

## Output

For each test case, print one integer — the minimum total effort to repair all
broken steps in the minimum number of days.

## Example

```text
Input
6
5
0 0 0 0 0
4
0 13 15 8
4
13 15 0 8
8
1 2 3 4 5 6 7 8
5
99999999 100000000 99999999 99999999 99999999
5
2 3 4 3 2

Output
0
59
64
72
899999993
24
```

In the second test case: repair steps `3` and `4` on day 1, then step `2` on
day 2. Effort `2*(15+8)+13 = 59`.

In the third test case: repair step `4` on day 1, then steps `1` and `2` on
day 2. Effort `8+2*(13+15) = 64`.


### ideas
1. 当只有单个的时候, 那么贡献 = a[i]
2. else 连续的一段, 只能尽量选择操作2(减少天数)
3. 但是如果是奇数的, 那么可以选择奇数位置的最大值, 节省下来