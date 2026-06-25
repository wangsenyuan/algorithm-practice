# C. Median Partition

[Problem link](https://codeforces.com/problemset/problem/2222/C)

**Contest:** [Spectral::Cup 2026 Round 1 (Codeforces Round 1094, Div. 1 + Div. 2)](https://codeforces.com/contest/2222)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given an array `a` with an odd length `n` consisting of positive integers. Partition the array into
several subarrays with odd lengths that all have the same median. Find the maximum number of subarrays.

More formally, find a strictly increasing sequence `k` of length `(p + 1)` with `k_1 = 1` and `k_{p+1} = n + 1`,
such that for every `1 <= i <= p`, the medians of the segments `[k_i, k_{i+1} - 1]` are all equal. The parity of
`k_i` and `k_{i+1}` must be different. Maximize `p`.

A subarray is obtained by deleting some (possibly zero) elements from the beginning and end of `a`.

The median of an odd-length array is the element at position `ceil(x / 2)` after sorting in non-decreasing order.

## Input

The first line contains an integer `t` (`1 <= t <= 1000`) — the number of test cases.

Each test case contains:

- one line with a single integer `n` (`1 <= n < 5000`, `n` is odd);
- one line with `n` integers `a_1, a_2, ..., a_n` (`1 <= a_i <= 10^9`).

It is guaranteed that the sum of `n^2` over all test cases does not exceed `5000^2`.

## Output

For each test case, print one integer — the maximum number of subarrays.

## Example

### Input

```text
10
5
3 3 2 4 3
7
9 5 7 7 4 7 7
9
1 1 1 1 1 1 1 1 1
15
3 1 2 3 3 2 2 2 5 1 2 3 4 5 2
3
2 2 2
5
1 2 3 4 5
5
2 1 3 2 2
7
2 2 1 2 3 2 2
9
2 1 2 3 2 1 2 3 2
```

### Output

```text
3
3
9
1
1
1
3
1
5
5
```

### Note

- In the first test case, one optimal partition is `[3], [3], [2, 4, 3]`.
- In the second test case, one optimal partition is `[9, 5, 7], [7], [4, 7, 7]`.
- In the third test case, every element can form its own subarray.

### ideas
1. 每段的median就是整段的median
2. 这个应该可以反证. 
3. dp[i] = j表示到i为止, 且形成了奇数段的median = w的最多分段数