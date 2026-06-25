# D. Permutation Construction

[Problem link](https://codeforces.com/problemset/problem/2222/D)

**Contest:** [Spectral::Cup 2026 Round 1 (Codeforces Round 1094, Div. 1 + Div. 2)](https://codeforces.com/contest/2222)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given an array `a` consisting of `n` integers.

For an inversion `(i, j)` in a permutation `p`, its value is defined as

```text
sum_{k=i}^{j-1} a_k
```

The beauty of a permutation is the sum of the values over all its inversions.

You have to construct a permutation `p` of length `n` that maximizes its beauty.

An inversion in a permutation `p` of length `n` is a pair of indices `(i, j)` such that `1 <= i < j <= n`
and `p_i > p_j`.

If there are multiple valid answers, you may print any of them.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases.

Each test case contains:

- one line with a single integer `n` (`1 <= n <= 2 * 10^5`);
- one line with `n` integers `a_1, a_2, ..., a_n` (`-10^9 <= a_i <= 10^9`).

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, print one line containing `n` integers — a permutation `p` that maximizes its beauty.

## Example

### Input

```text
7
1
0
2
1000000000 3
1
2
4
-1 -2 -3 -4
5
-1 2 -3 2 -1
6
1 -1 3 -4 1 -3
7
-3 -2 -1 4 -1 -2 -3
```

### Output

```text
1
2 1
1
1 2 3 4
3 4 1 5 2
5 2 4 1 6 3
1 4 6 7 2 3 5
```

### Note

- In the first test case, the only permutation is `[1]`.
- In the second test case, for `p = [2, 1]`, the only inversion is `(1, 2)` and its value is `10^9`.
- In the third test case, the only permutation is `[1]`.
- In the fourth test case, one optimal permutation is `[1, 2, 3, 4]`.
- In the fifth test case, one optimal permutation is `[3, 4, 1, 5, 2]`.

### ideas
1. 这个题目咋这么奇怪的, 好绕口
2. 