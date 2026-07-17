# D - Cubes

[Problem link](https://atcoder.jp/contests/abc397/tasks/abc397_d)

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 425 points

## Problem Statement

You are given a positive integer `N`. Determine whether there exists a pair of
positive integers `(x, y)` such that `x^3 - y^3 = N`. If such a pair exists,
print one such pair `(x, y)`.

## Constraints

- `1 <= N <= 10^18`
- All input values are integers

## Input

```
N
```

## Output

If there is no pair of positive integers `(x, y)` satisfying
`x^3 - y^3 = N`, print `-1`. If there is such a pair, print `x` and `y`
separated by a space. If there are multiple solutions, any one is accepted.

## Samples

### Sample 1

Input:

```
397
```

Output:

```
12 11
```

`12^3 - 11^3 = 397`.

### Sample 2

Input:

```
1
```

Output:

```
-1
```

### Sample 3

Input:

```
39977273855577088
```

Output:

```
342756 66212
```
