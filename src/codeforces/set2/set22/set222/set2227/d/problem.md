# D. Palindromex

[Problem link](https://codeforces.com/problemset/problem/2227/D)

**Contest:** [Codeforces contest 2227](https://codeforces.com/contest/2227)

## Problem

You are given an array `a` of `2n` integers. Every integer `x` in `[0, n-1]`
appears exactly twice.

Find a contiguous subarray that is a palindrome whose MEX is maximized, and
output that maximum MEX.

A subarray is a palindrome if it reads the same forwards and backwards.

The MEX of a multiset is the smallest non-negative integer that does not appear
in it.

## Constraints

- `1 <= t <= 10^4`
- `1 <= n <= 10^5`
- `0 <= a_i <= n-1`
- Every value in `[0, n-1]` appears exactly twice
- Sum of `2n` over all test cases `<= 2 * 10^5`

## Input

```text
t
case_1
...
case_t
```

Each test case:

```text
n
a1 a2 ... a_{2n}
```

## Output

For each test case, print one integer — the maximum MEX of any palindromic
subarray.

## Sample 1

```text
Input
4
1 2 0 3 3 0 2 1

Output
4
```

## Sample 2

```text
Input
2
0 1 0 1

Output
2
```

## Sample 3

```text
Input
2
1 1 0 0

Output
1
```

## Sample 4

```text
Input
3
2 0 2 1 1 0

Output
1
```

## Sample 5

```text
Input
4
0 1 3 0 3 1 2 2

Output
2
```

## Sample 6

```text
Input
3
0 1 2 1 0 2

Output
3
```
