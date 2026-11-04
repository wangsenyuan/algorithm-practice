# B. Routine Problem

[Problem link](https://codeforces.com/problemset/problem/337/B)

**Contest:** [Codeforces Round 196 (Div. 2)](https://codeforces.com/contest/337)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

A monitor has width-to-height ratio `a:b`. A movie has width-to-height ratio `c:d`. The movie keeps its ratio, fits entirely on the screen, and occupies as much of the screen as possible.

Print the ratio of the empty area to the whole screen as an irreducible fraction `p/q`.

## Constraints

- `1 ≤ a, b, c, d ≤ 1000`

## Input

```text
a b c d
```

## Output

Print `p/q`, where `p` and `q` are coprime and `q > 0`.

## Sample Input 1

```text
1 1 3 2
```

## Sample Output 1

```text
1/3
```

The screen is square and the movie is `3:2`. Matching the widths leaves the movie covering `2/3` of the height, so `1/3` of the screen is empty.

## Sample Input 2

```text
4 3 2 2
```

## Sample Output 2

```text
1/4
```

The screen is `4:3` and the movie is square. Matching the heights leaves the movie covering `3/4` of the width.

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
