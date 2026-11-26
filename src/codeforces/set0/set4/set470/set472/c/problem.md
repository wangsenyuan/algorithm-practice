# C. Design Tutorial: Make It Nondeterministic

[Problem link](https://codeforces.com/problemset/problem/472/C)

**Contest:** [Codeforces Round 270 (Div. 2)](https://codeforces.com/contest/472)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

## Problem

There are `n` people. Each person `i` has a first name `f_i` and a last name `s_i`.
Each person chooses one of these names as a handle.

Given a permutation `p`, determine whether the handles can be selected so that the
handles are lexicographically increasing in the order `p[0], p[1], ..., p[n-1]`.

All `2n` given names are distinct.

## Input

The first line contains an integer `n` (`1 <= n <= 10^5`).

The next `n` lines contain two lowercase strings `f_i` and `s_i`
(`1 <= |f_i|, |s_i| <= 50`).

The following line contains `n` distinct integers `p_1, p_2, ..., p_n`
(`1 <= p_i <= n`) describing the required person order.

## Output

Print `YES` if a valid choice exists, otherwise print `NO`.

## Examples

```text
Input
3
gennady korotkevich
petr mitrichev
gaoyuan chen
1 2 3

Output
NO
```

```text
Input
3
gennady korotkevich
petr mitrichev
gaoyuan chen
3 1 2

Output
YES
```

```text
Input
2
galileo galilei
nicolaus copernicus
2 1

Output
YES
```
