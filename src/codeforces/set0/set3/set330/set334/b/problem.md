# B. Eight Point Sets

[Problem link](https://codeforces.com/problemset/problem/334/B)

**Contest:** [Codeforces Round 194 (Div. 2)](https://codeforces.com/contest/334)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

An eight-point set is respectable when it is the intersection of three distinct vertical lines and three distinct horizontal lines, with the center point removed.

There are integers `x1 < x2 < x3` and `y1 < y2 < y3` such that the set is every point `(xi, yj)` except `(x2, y2)`.

You are given eight points. Print `respectable` if they form such a set, and `ugly` otherwise.

## Constraints

- Each coordinate is an integer from `0` to `10^6`
- The eight points are otherwise unrestricted

## Input

Eight lines, each containing two integers `xi yi`.

## Output

Print `respectable` or `ugly`.

## Sample Input 1

```text
0 0
0 1
0 2
1 0
1 2
2 0
2 1
2 2
```

## Sample Output 1

```text
respectable
```

## Sample Input 2

```text
0 0
1 0
2 0
3 0
4 0
5 0
6 0
7 0
```

## Sample Output 2

```text
ugly
```

## Sample Input 3

```text
1 1
1 2
1 3
2 1
2 2
2 3
3 1
3 2
```

## Sample Output 3

```text
ugly
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
