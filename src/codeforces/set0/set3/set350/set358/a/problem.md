# A. Dima and Continuous Line

[Problem link](https://codeforces.com/problemset/problem/358/A)

**Contest:** [Codeforces Round 208 (Div. 2)](https://codeforces.com/contest/358)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

`n` distinct points lie on the x-axis. Connect them in the given order by semicircles above the axis: the first point to the second, the second to the third, and so on. The semicircle between two points has that segment as its diameter.

Decide whether this chain intersects itself.

## Constraints

- `1 ≤ n ≤ 10^3`
- `-10^6 ≤ x_i ≤ 10^6`
- The coordinates are distinct and are not necessarily sorted

## Input

```text
n
x_1 x_2 ... x_n
```

## Output

Print `yes` if the chain has a self-intersection. Otherwise print `no`.

## Sample Input 1

```text
4
0 10 5 15
```

## Sample Output 1

```text
yes
```

The semicircle from `0` to `10` crosses the semicircle from `5` to `15`.

## Sample Input 2

```text
4
0 15 5 10
```

## Sample Output 2

```text
no
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
