# D. Sanae, Cross and Color

[Problem link](https://codeforces.com/problemset/problem/2228/D)

**Contest:** [Codeforces Round 1098 (Div. 2)](https://codeforces.com/contest/2228)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

## Problem

Given `n` distinct integer-coordinate points, choose integers `k1` and `k2`.
The lines `x = k1 + 0.5` and `y = k2 + 0.5` split the plane into four regions.
The choice is valid when every region contains at least one point; its regions
color points red (top-left), green (top-right), blue (bottom-left), and yellow
(bottom-right). Count distinct resulting colorings, rather than line choices.

## Input

The first line contains the number of test cases `t` (`1 <= t <= 10^4`).

For each test case:

- the first line contains `n` (`4 <= n <= 2 * 10^6`);
- the next `n` lines contain distinct points `(x_i, y_i)` where
  `1 <= x_i, y_i <= n`.

The sum of `n` over all test cases is at most `2 * 10^6`.

## Output

For every test case, print the number of distinct valid colorings.

## Sample

```text
Input
5
4
1 1
2 2
3 3
4 4
4
1 4
4 1
1 1
4 4
8
7 2
5 7
2 7
1 3
6 7
3 6
7 5
1 6
8
6 1
3 6
1 4
1 1
4 2
5 5
3 4
4 1
6
5 5
5 4
3 5
1 5
5 3
2 2

Output
0
1
12
8
4
```

## Status

Implemented with coordinate buckets and prefix/suffix extrema. The tests cover
each official test case independently.

## Ideas

A coloring is determined by the two point partitions induced by the vertical
and horizontal lines, not by the exact line coordinates. Since every coordinate
is in `1..n`, aggregate each `x` directly instead of sorting it: store the
minimum and maximum `y` in that column, then skip empty columns while sweeping
the vertical cut from left to right.

Maintain `loPre` and `hiPre`, the minimum and maximum `y` among points left of
the cut. Suffix extrema give `loSuf` and `hiSuf` for the right side. A valid
horizontal cut must leave both sides with a point above and below it, so its
endpoints must lie between:

\[
y_1 = \max(loPre, loSuf), \qquad y_2 = \min(hiPre, hiSuf).
\]

There is no valid cut when `y1 >= y2`.

Within a valid range, crossing an absent `y` coordinate does not change any
point's color. Mark every globally present `y` and prefix-sum those marks. If
`cnt` distinct `y` coordinates occur in `[y1, y2]`, this vertical partition
contributes `cnt - 1` distinct colorings.

Every bucket, suffix value, and vertical partition is processed once, yielding
`O(n)` time and `O(n)` space.

## Summary

Bucket points by `x`, sweep the nonempty columns, and intersect the left/right
`y` spans using prefix and suffix extrema. A global prefix count of present
`y` values converts each feasible span into its number of distinct horizontal
color partitions, so no sorting or dynamic data structure is needed.
