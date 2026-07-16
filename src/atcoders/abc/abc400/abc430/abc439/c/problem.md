# C - 2026

[Problem link](https://atcoder.jp/contests/abc439/tasks/abc439_c)

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 300 points

## Problem Statement

A positive integer `n` is called a **good integer** when it satisfies the
following condition:

- There exists exactly one pair of integers `(x, y)` that satisfies
  `0 < x < y` and `x^2 + y^2 = n`.

For example, when `n = 2026`, it can be verified that `(x, y) = (1, 45)` is the
only such pair. Thus `2026` is a good integer.

You are given a positive integer `N`. Enumerate all good integers not
exceeding `N`.

## Constraints

- `1 <= N <= 10^7`
- `N` is an integer

## Input

```
N
```

## Output

Let there be `k` good integers not exceeding `N`, and let
`(a_1, a_2, ..., a_k)` be them in ascending order. Output:

```
k
a_1 a_2 ... a_k
```

If `k = 0`, output the second line as an empty line.

## Samples

### Sample 1

Input:

```
10
```

Output:

```
2
5 10
```

### Sample 2

Input:

```
1
```

Output:

```
0

```

### Sample 3

Input:

```
50
```

Output:

```
14
5 10 13 17 20 25 26 29 34 37 40 41 45 50
```
