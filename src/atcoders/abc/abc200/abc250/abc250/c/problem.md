# C - Adjacent Swaps

[Problem link](https://atcoder.jp/contests/abc250/tasks/abc250_c)

**Contest:** [AtCoder Beginner Contest 250](https://atcoder.jp/contests/abc250)

time limit: 2 seconds

memory limit: 1024 MiB

## Problem

Initially, `N` balls are arranged as `1, 2, ..., N`. For each operation with
value `x`, swap ball `x` with the ball immediately to its right. If `x` is at
the rightmost position, swap it with the ball immediately to its left instead.
Print the final arrangement.

## Input

```text
N Q
x_1
...
x_Q
```

- `2 <= N <= 2 * 10^5`
- `1 <= Q <= 2 * 10^5`
- `1 <= x_i <= N`

## Output

Print the final labels from left to right, separated by spaces.

## Sample

```text
Input
5 5
1
2
3
4
5

Output
1 2 3 5 4
```

## Status

`solve` is intentionally left as a TODO. The tests preserve the three official
cases and are skipped until an implementation is added.
