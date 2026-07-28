# C. View Angle

[Problem link](https://codeforces.com/problemset/problem/257/C)

**Contest:** [Codeforces Round #159 (Div. 2)](https://codeforces.com/contest/257)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

There are `n` mannequins on the plane. Find the minimum angle with vertex at the
origin such that every mannequin lies inside or on the boundary of the angle.
Output the angle in degrees.

## Constraints

- `1 <= n <= 10^5`
- `|x_i|, |y_i| <= 1000`
- No mannequin at the origin
- All mannequins are at distinct points
- Absolute or relative error up to `10^{-6}` is accepted

## Input

The first line contains `n`.

Each of the next `n` lines contains two integers `x_i y_i` — coordinates of a
mannequin.

## Output

Print one real number — the minimum angle in degrees.

## Sample 1

```text
Input
2
2 0
0 2

Output
90.0000000000
```

## Sample 2

```text
Input
3
2 0
0 2
-2 2

Output
135.0000000000
```

## Sample 3

```text
Input
4
2 0
0 2
-2 0
0 -2

Output
270.0000000000
```

## Sample 4

```text
Input
2
2 1
1 2

Output
36.8698976458
```

## ideas

1. Convert each point to a polar angle with `atan2(y, x)` in `[0, 2π)`.
2. Sort the angles.
3. The uncovered complementary arc is the largest gap between consecutive
   angles (including the wrap-around gap from last to first).
4. Answer (degrees) = `360 - max_gap_degrees`.
5. If `n = 1`, the answer is `0`.

## summary

Minimum origin-centered viewing angle covering all points equals a full turn
minus the largest empty angular sector between consecutive points on the unit
circle of directions.
