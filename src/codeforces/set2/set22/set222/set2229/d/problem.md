# D. Me When Median Problem

[Problem link](https://codeforces.com/problemset/problem/2229/D)

**Contest:** [Spectral::Cup 2026 Round 2 (Codeforces Round 1100, Div. 1 + Div. 2)](https://codeforces.com/contest/2229)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem Statement

You are given two arrays of positive integers `a` and `b`, both of length `n`.
You will perform the following operation exactly `n - 1` times:

- let `m` be the current length of `a` and `b` (the lengths stay equal);
- select an integer `i` (`1 <= i < m`);
- let `S` be the multiset `{a_i, a_{i+1}, b_i, b_{i+1}}`;
- sort the elements of `S` so that `s_1 <= s_2 <= s_3 <= s_4`;
- replace `a_i, a_{i+1}` with `s_2`, and replace `b_i, b_{i+1}` with `s_3`.
  More formally, replace `a` with
  `[a_1, ..., a_{i-1}, s_2, a_{i+2}, ..., a_m]` and replace `b` with
  `[b_1, ..., b_{i-1}, s_3, b_{i+2}, ..., b_m]`.

After all operations, exactly one element remains in both `a` and `b`.
Determine the maximum attainable value of `min(a_1, b_1)`.

## Input

Each test contains multiple test cases. The first line contains the number of
test cases `t` (`1 <= t <= 10^4`).

Each test case:

- the first line contains an integer `n` (`1 <= n <= 10^5`) — the length of
  `a` and `b`;
- the second line contains `n` integers `a_1, ..., a_n` (`1 <= a_i <= 2 * n`);
- the third line contains `n` integers `b_1, ..., b_n` (`1 <= b_i <= 2 * n`).

The sum of `n` over all test cases does not exceed `10^5`.

## Output

For each test case, output the maximum attainable value of `min(a_1, b_1)`.

## Sample Input 1

```text
6
1
1
2
3
2 4 5
1 3 6
4
7 5 4 8
4 6 7 8
8
8 7 13 11 1 10 4 5
11 11 12 8 9 2 3 13
9
16 1 9 12 5 18 10 10 16
14 6 7 11 12 17 18 3 17
6
3 6 12 4 10 12
2 3 2 7 8 9
```

## Sample Output 1

```text
1
3
6
8
14
8
```

## Note

In the first test case, no operations are needed, so the answer is
`min(1, 2) = 1`.

In the second test case, one optimal sequence:

1. `i = 1`: `S = {2, 4, 1, 3}` → `a = [2, 5]`, `b = [3, 6]`
2. `i = 1`: `S = {2, 5, 3, 6}` → `a = [3]`, `b = [5]`

Then `min(3, 5) = 3`.


### ideas
1. 这个也很难呐, 即使知道, 要用二分来处理, 也不知道怎么check
2. 把a, b全部变成0 for v < w, else 1
3. 那么就是尽量的保留1, 
4. (0, 1) (0, 1) => (0, 1)
5. (0, 1) (1, 1) => (1, 1)
6. (0, 0), (1, 1) => (0, 1)
7. (0, 0), (0, 1) => (0, 0)
8. 