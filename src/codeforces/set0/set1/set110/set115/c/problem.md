# C. Plumber

[Problem link](https://codeforces.com/problemset/problem/115/C)

time limit per test: 3 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Little John has drawn a grid of `n` rows and `m` columns. In each cell he will
draw a pipe segment of one of four types numbered `1`–`4`.

Each pipe segment has two ends. For example, segment `1` has ends at the top
and left sides.

The piping system is **leaking** if there is at least one pipe segment whose end
is not connected to another pipe's end or to the border of the grid.

You are given a partially filled grid: each cell is either one of `"1"`–`"4"`
or `"."` (empty). Count the number of ways to fill all empty cells with pipe
segments `1`–`4` so that the final system is non-leaking. Print the answer
modulo `1000003` (`10^6 + 3`).

Rotations and flips of the grid are **not** considered the same: two
configurations that match only after rotation or flipping are different.

## Constraints

- `1 <= n, m` and `n * m <= 5 * 10^5`
- Each cell is one of `"1"`, `"2"`, `"3"`, `"4"`, or `"."`

## Input

The first line contains two integers `n` and `m`.

Then `n` lines follow, each with exactly `m` characters describing the grid.

## Output

Print a single integer — the number of non-leaking completions modulo
`1000003`. If there are none, print `0`.

## Examples

### Sample 1

```text
Input
2 2
13
..

Output
2
```

### Sample 2

```text
Input
3 1
1
4
.

Output
0
```

### Sample 3

```text
Input
2 2
3.
.1

Output
1
```
