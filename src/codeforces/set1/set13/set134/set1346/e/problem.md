# E. Magic Tricks

[Problem link](https://codeforces.com/problemset/problem/1346/E)

**Contest:** [Kotlin Heroes: Episode 4](https://codeforces.com/contest/1346)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

There are `n` balls in a row. The special ball starts at position `k`. Positions are numbered from `1` to `n`, left to right.

Masha then performs `m` swaps. The `i`-th swap exchanges the balls on positions `x_i` and `y_i`. She may fake any subset of these swaps. A faked swap looks real to the audience, but the balls do not move.

For each final position, print the minimum number of swaps she must fake so the special ball ends there. If that position is unreachable, print `-1`.

## Constraints

- `2 ≤ n ≤ 2 · 10^5`
- `1 ≤ m ≤ 2 · 10^5`
- `1 ≤ k ≤ n`
- `1 ≤ x_i, y_i ≤ n`
- `x_i ≠ y_i`

## Input

```text
n m k
x_1 y_1
x_2 y_2
...
x_m y_m
```

## Output

Print `n` integers. The `i`-th integer is the minimum number of faked swaps that leave the special ball on position `i`, or `-1` when that is impossible.

## Sample Input 1

```text
4 5 1
3 4
2 1
4 1
3 1
3 1
```

## Sample Output 1

```text
2 0 3 1
```

## Sample Input 2

```text
5 7 4
3 2
3 2
4 2
3 4
4 1
3 2
5 2
```

## Sample Output 2

```text
2 2 0 3 1
```

## Sample Input 3

```text
7 15 5
5 3
4 2
6 1
2 4
1 6
3 7
5 6
4 2
6 4
2 6
6 3
6 3
7 6
2 6
7 2
```

## Sample Output 3

```text
-1 1 1 1 2 1 0
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
