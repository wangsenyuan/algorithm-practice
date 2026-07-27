# C - Upgrade Required

[Problem link](https://atcoder.jp/contests/abc426/tasks/abc426_c)

**Contest:** [AtCoder Beginner Contest 426](https://atcoder.jp/contests/abc426)

time limit: 2 sec

memory limit: 1024 MiB

score: 300 points

## Problem

There are `N` OS versions numbered `1, 2, ..., N` (older to newer), and `N`
PCs. Initially, the `i`-th PC has version `i`.

Process `Q` operations in order. The `i`-th operation:

- Upgrade all PCs whose current version is `X_i` or earlier to version
  `Y_i` (`Y_i > X_i`). Then print how many PCs were upgraded in this
  operation.

Later operations see the results of earlier upgrades.

## Constraints

- All input values are integers
- `2 <= N <= 10^6`
- `1 <= Q <= 2 * 10^5`
- `1 <= X_i < Y_i <= N`

## Input

```text
N Q
X_1 Y_1
...
X_Q Y_Q
```

## Output

Print `Q` lines. The `i`-th line is the number of PCs upgraded in the `i`-th
operation.

## Sample 1

```text
Input
8 5
2 6
3 5
1 7
5 7
7 8

Output
2
1
0
3
7
```
