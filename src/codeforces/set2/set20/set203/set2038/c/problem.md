# C. DIY

[Problem link](https://codeforces.com/problemset/problem/2038/C)

**Contest:** [2024-2025 ICPC, NERC, Southern and Volga Russian Regional Contest (Unrated, Online Mirror, ICPC Rules, Preferably Teams)](https://codeforces.com/contest/2038)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

You are given a list of `n` integers `a_1, a_2, …, a_n`. You need to pick 8 elements from the list and use them as coordinates of four points. These four points should be corners of a rectangle which has its sides parallel to the coordinate axes. Your task is to pick coordinates in such a way that the resulting rectangle has the maximum possible area. The rectangle can be degenerate, i.e. its area can be 0. Each integer can be used as many times as it occurs in the list (or less).

## Input

The first line contains one integer `t` (`1 <= t <= 25000`) — the number of test cases.

The first line of each test case contains one integer `n` (`8 <= n <= 2 * 10^5`).

The second line of each test case contains `n` integers `a_1, a_2, …, a_n` (`-10^9 <= a_i <= 10^9`).

The sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, print the answer as follows:

- if it is impossible to construct a rectangle which meets the constraints from the statement, print a single line containing the word `NO` (case-insensitive);
- otherwise, in the first line, print `YES` (case-insensitive). In the second line, print 8 integers `x_1, y_1, x_2, y_2, x_3, y_3, x_4, y_4` — the coordinates of the corners of the rectangle. You can print the corners in any order.

## Example

### Input

```text
3
16
-5 1 1 2 2 3 3 4 4 5 5 6 6 7 7 10
8
0 0 -1 2 2 1 1 3
8
0 0 0 0 0 5 0 5
```

### Output

```text
YES
1 2 1 7 6 2 6 7
NO
YES
0 0 0 5 0 0 0 5
```

## Solution

An axis-aligned rectangle is four corners `(x1, y1), (x2, y1), (x2, y3), (x1, y3)`.
That uses each of `x1`, `x2`, `y1`, `y3` twice, so a value can appear as a
coordinate only as many times as `⌊count/2⌋`. Sort `a` and emit one token
per such pair.

Four tokens are required (two x-coordinates and two y-coordinates). Fewer
than four means `NO`. Otherwise the maximum area uses the two smallest
tokens as the low x and low y, and the two largest as the high x and high y.
Assigning the two highs to the two axes is a 2-way product comparison
`(u-x1)(v-y1)` versus `(v-x1)(u-y1)`; swapping the two lows is the same
choice with the axes flipped. Middle tokens cannot enlarge both spans, so
they are discarded. A zero product is allowed (degenerate rectangle).

### Correctness sketch

Every feasible rectangle corresponds to four pair-tokens. For a sorted
token list, both side lengths are differences of two chosen values, so
replacing an interior token by a more extreme unused token never decreases
either length. Hence the optimum lives on `{smallest, second-smallest,
second-largest, largest}`. The product comparison enumerates the two
distinct partitions of those four values into an x-pair and a y-pair.

### Complexity

Sorting dominates: `O(n log n)` time and `O(n)` extra memory per test.
The sum of `n` over tests is at most `2 · 10^5`.
