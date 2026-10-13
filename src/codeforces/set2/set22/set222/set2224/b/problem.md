# B. Zhily and Mex and Max

[Problem link](https://codeforces.com/problemset/problem/2224/B)

**Contest:** [Codeforces Round 1097 (Div. 2, Based on Zhili Cup 2026)](https://codeforces.com/contest/2224)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

You are given an array `a` of `n` non-negative integers. You may rearrange it.

The MEX of a collection is the smallest non-negative integer that does not occur in it.

Maximize the sum, over every prefix, of that prefix's MEX plus its maximum:

```text
sum_{i=1..n} (mex(a_1, ..., a_i) + max(a_1, ..., a_i))
```

## Input

The first line contains `t` (`1 <= t <= 10^4`), the number of test cases.

Each test case is two lines:

- `n` (`1 <= n <= 2 * 10^5`)
- `n` integers `a_1, a_2, …, a_n` (`0 <= a_i <= 10^9`)

The sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, print the maximum sum.

## Example

### Input

```text
5
5
0 0 0 0 0
2
0 1
5
1 1 1 1 0
6
1 1 4 5 1 4
1
1000000000
```

### Output

```text
5
4
13
30
1000000000
```

In the first test case every prefix has MEX `1` and maximum `0`, so the sum is `5`.

In the third test case, rearranging into `[0, 1, 1, 1, 1]` gives `(1+0) + (2+1) + (2+1) + (2+1) + (2+1) = 13`.

## Status

I/O and official sample cases are in place. `solve` is left as a TODO.
