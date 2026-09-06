# B. Swaps

[Problem link](https://codeforces.com/problemset/problem/1573/B)

**Contest:** [Codeforces Round 743 (Div. 2)](https://codeforces.com/contest/1573)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given two arrays `a` and `b` of length `n`. Array `a` contains each odd integer from `1` to `2n` in an arbitrary order, and array `b` contains each even integer from `1` to `2n` in an arbitrary order.

You can perform the following operation on those arrays:

- choose one of the two arrays
- pick an index `i` from `1` to `n-1`
- swap the `i`-th and the `(i+1)`-th elements of the chosen array

Compute the minimum number of operations needed to make array `a` lexicographically smaller than array `b`.

For two different arrays `x` and `y` of the same length `n`, we say that `x` is lexicographically smaller than `y` if in the first position where `x` and `y` differ, the array `x` has a smaller element than the corresponding element in `y`.

An answer always exists.

## Input

The first line contains the number of test cases `t` (`1 <= t <= 10^4`).

The first line of each test case contains a single integer `n` (`1 <= n <= 10^5`) — the length of the arrays.

The second line contains `n` integers `a_1, a_2, …, a_n` (`1 <= a_i <= 2n`, all `a_i` are odd and pairwise distinct).

The third line contains `n` integers `b_1, b_2, …, b_n` (`1 <= b_i <= 2n`, all `b_i` are even and pairwise distinct).

The sum of `n` over all test cases does not exceed `10^5`.

## Output

For each test case, print one integer: the minimum number of operations needed to make array `a` lexicographically smaller than array `b`.

## Example

### Input

```text
3
2
3 1
4 2
3
5 3 1
2 4 6
5
7 5 9 1 3
2 4 6 10 8
```

### Output

```text
0
2
3
```

### Note

In the first example, `a` is already lexicographically smaller than `b`, so no operations are required.

In the second example, one way is to swap `5` and `3` and then swap `2` and `4`, which results in `[3, 5, 1]` and `[4, 2, 6]`.

## ideas
1. a[0] != b[0]
2. 所以, 有一个简单的策略是调整a中的, 找到最近的a[i] < b[0] (或者 a[0] < b[j])
3. 有没有情况是, 找到a[i] < b[j], 将它们调整到(0, 0)
