# C. Subarrays

[Problem link](https://codeforces.com/problemset/problem/2276/C)

**Contest:** [Educational Codeforces Round 195 (Rated for Div. 2)](https://codeforces.com/contest/2276)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

## Problem

Monocarp has an array `a` of `n` positive integers and a favorite integer `k` with `k < n`.

A subarray from index `l` to `r` is `k`-balanced when both of these hold:

- the sum of `a_l + ... + a_r` is divisible by `k`
- the length `r - l + 1` is strictly greater than `k`

Count the `k`-balanced subarrays.

## Constraints

- `1 ≤ t ≤ 10^4`
- `1 ≤ k < n ≤ 2 · 10^5`
- `1 ≤ a_i ≤ 10^9`
- The sum of `n` over all test cases does not exceed `2 · 10^5`

## Input

The first line contains `t`, the number of test cases.

Each test case is two lines:

```text
n k
a_1 a_2 ... a_n
```

## Output

For each test case, print the number of `k`-balanced subarrays.

## Sample Input 1

```text
5 2
2 1 3 1 4
```

## Sample Output 1

```text
2
```

The two subarrays are `[2, 1, 3]` (sum `6`) and `[3, 1, 4]` (sum `8`).

## Sample Input 2

```text
9 3
1 1 1 1 1 1 1 1 1
```

## Sample Output 2

```text
5
```

Every element is `1` and `k = 3`. The suitable lengths are `6` and `9`: four subarrays of length `6` and one of length `9`.

## Sample Input 3

```text
4 3
2 8 5 7
```

## Sample Output 3

```text
0
```

The only subarray longer than `3` is the whole array. Its sum is `22`, which is not divisible by `3`.

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
