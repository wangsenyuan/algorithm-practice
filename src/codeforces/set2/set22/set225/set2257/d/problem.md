# D. Bermuda Rectangle

[Problem link](https://codeforces.com/problemset/problem/2257/D)

**Contest:** [Codeforces Round 1117 (Div. 2)](https://codeforces.com/contest/2257)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

## Problem

The Bermuda Rectangle has integer side lengths, bottom-left corner `(0, 0)`,
and area `S`. For every query rectangle with side lengths `x` and `y`, also
anchored at `(0, 0)`, count its cells that can lie inside at least one possible
Bermuda Rectangle of area `S`.

## Input

The first line contains the number of test cases `t` (`1 <= t <= 10^4`).

For each test case:

- the first line contains `S` and `q` (`1 <= S <= 10^14`,
  `1 <= q <= 3 * 10^5`);
- each of the next `q` lines contains `x` and `y` (`1 <= x, y <= S`).

The sum of `q` is at most `3 * 10^5`; the sum of `sqrt(S)` is at most `10^7`.

## Output

For every query, output the requested cell count on its own line.

## Sample

```text
Input
3
6 4
2 3
4 5
6 6
1 1
5 2
2 2
3 4
8 2
3 1
5 6

Output
6
11
14
1
3
6
3
15
```

## Status

Implemented with divisor enumeration, staircase-area prefix sums, and binary
search. The sample tests cover each official test case independently.

## Ideas

List every divisor of `S` in increasing order:

\[
w_1 < w_2 < \dots < w_n, \qquad h_i = S / w_i.
\]

The possible rectangles are all anchored at `(0, 0)`, so their union forms a
descending staircase. With `w_0 = 0`, staircase segment `i` has width
`w_i - w_{i-1}` and height `h_i`.

For a query `(x, y)`, find:

- `i`: the first index with `w_i >= x`;
- `j`: the largest index with `h_j >= y`. Equivalently, it is the largest
  index with `w_j <= floor(S / y)`.

If `j >= i`, every cell of the query rectangle is covered, so the answer is
`x * y`. Otherwise, the covered area consists of the left portion clipped at
height `y`, the complete middle staircase segments, and the final partial
segment:

\[
f(x, y) = w_j y
 + \sum_{k=j+1}^{i-1}(w_k-w_{k-1})h_k
 + (x-w_{i-1})h_i.
\]

Precompute the staircase-area prefix sums

\[
pre_i = \sum_{k=1}^{i}(w_k-w_{k-1})h_k,
\]

so the middle term is `pre_{i-1} - pre_j`. Each query then takes two binary
searches and `O(1)` arithmetic after enumerating the divisors.

## Summary

The union of all valid rectangles is a monotone staircase indexed by the
divisors of `S`. For a query, binary-search the segment containing `x` and the
last segment taller than `y`; prefix areas then give the clipped staircase area
in `O(log d)` time, where `d` is the number of divisors of `S`. Enumerating the
divisors costs `O(sqrt(S))` per test case.
