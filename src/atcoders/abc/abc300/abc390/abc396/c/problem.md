# C - Buy Balls

[Problem link](https://atcoder.jp/contests/abc396/tasks/abc396_c)

**Contest:** [AtCoder Beginner Contest 396](https://atcoder.jp/contests/abc396)

time limit: 2 sec

memory limit: 1024 MiB

score: 300 points

## Problem

There are `N` black balls and `M` white balls. The value of the `i`-th black
ball is `B_i`, and the value of the `j`-th white ball is `W_j`.

Choose zero or more balls so that the number of black balls chosen is at least
the number of white balls chosen. Among all such choices, find the maximum
possible sum of the values of the chosen balls.

## Constraints

- `1 <= N, M <= 2 * 10^5`
- `-10^9 <= B_i, W_j <= 10^9`
- All input values are integers

## Input

```text
N M
B_1 B_2 ... B_N
W_1 W_2 ... W_M
```

## Output

Print the answer.

## Samples

### Sample 1

```text
Input
4 3
8 5 -1 3
3 -2 -4

Output
19
```

Choose black balls `1, 2, 4` and white ball `1`: `8 + 5 + 3 + 3 = 19`.

### Sample 2

```text
Input
4 3
5 -10 -2 -5
8 1 4

Output
15
```

Choose black balls `1, 3` and white balls `1, 3`: `5 + (-2) + 8 + 4 = 15`.

### Sample 3

```text
Input
3 5
-36 -33 -31
12 12 28 24 27

Output
0
```

Choosing no balls is allowed.

## Solution

Sort both arrays in non-increasing order. If we take exactly `i + 1` black
balls, the optimal choice is the prefix of the largest `i + 1` blacks. Among
whites, we may take at most `i + 1` of them, so we only ever consider prefixes
of the largest whites.

Scan black prefixes from left to right:

- `sum1` = sum of the first `i + 1` blacks
- while the white pointer `j` satisfies `j <= i`, extend the white prefix and
  keep `best = max` of all white prefix sums seen so far (including empty =
  `0`)
- update `ans = max(ans, sum1 + best)`

`best` is the best white contribution allowed under the constraint
`#white <= #black`. Negative blacks still get considered as we grow the black
prefix (they may unlock a better white prefix), but `ans` starts at `0`, so
choosing nothing remains available.

Time `O(N log N + M log M)`, space `O(1)` extra beyond the arrays.
