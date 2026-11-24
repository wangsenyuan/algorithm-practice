# C. Exams

[Problem link](https://codeforces.com/problemset/problem/479/C)

**Contest:** [Codeforces Round 274 (Div. 2)](https://codeforces.com/contest/479)

time limit per test: 1 second

memory limit per test: 256 megabytes

## Problem

Student Valera is taking `n` exams. Exam `i` can be taken on the scheduled day `a_i`,
or Valera can arrange to take it early on day `b_i` (`b_i < a_i`).

The teacher records the scheduled date `a_i` in the record book. Valera wants the
recorded dates to be in non-decreasing order. Find the minimum possible day on which
he can take the last exam.

## Input

The first line contains an integer `n` (`1 <= n <= 5000`) — the number of exams.

Each of the next `n` lines contains two integers `a_i` and `b_i`
(`1 <= b_i < a_i <= 10^9`) — the scheduled and early dates for exam `i`.

## Output

Print the minimum possible day of the last exam.

## Examples

```text
Input
3
5 2
3 1
4 2

Output
2
```

```text
Input
3
6 1
5 2
4 3

Output
6
```
