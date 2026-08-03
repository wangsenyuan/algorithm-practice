# D1. XOR Sorting (Easy Version)

[Problem link](https://codeforces.com/problemset/problem/2247/D1)

**Contest:** [Codeforces Round 1111 (Div. 2)](https://codeforces.com/contest/2247)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

This is the easy version of the problem. The only difference between the versions
is that in this version, `q = 0`.

Note that zero-based indexing is used in this problem.

For an array `b` consisting of `m` positive integers, define `f(b)` as follows.

For a non-negative integer `k`, we say that `b` can be `k`-sorted if it can be
sorted in non-decreasing order by performing the following operation any number
of times:

- Choose two indices `i` and `j` (`0 ≤ i < j ≤ m - 1`) such that `i ⊕ j ≤ k`\*.
  Note that the condition applies to the indices `i` and `j`, not to the
  elements `b_i` and `b_j`.
- Swap the elements `b_i` and `b_j`.

The value `f(b)` is defined as the smallest non-negative integer `k` such that
the array `b` can be `k`-sorted.

You are given an array `a` of length `n`, consisting of positive integers. You
will perform `q` updates on `a`. Each update has the following form:

- `i x`: assign `a_i = x`.

Note that the updates are persistent. In other words, each update affects all
subsequent states of the array.

For each of the `q + 1` states of `a` — the initial state and the state after
each of the `q` updates — find the value of `f(a)`.

\* `⊕` denotes the bitwise XOR operation.

## Constraints

- `1 ≤ t ≤ 10^4`
- `1 ≤ n ≤ 10^6`, `q = 0`
- `1 ≤ a_i ≤ 10^9`
- `0 ≤ i_j < n`, `1 ≤ x_j ≤ 10^9`
- Sum of `n` over all test cases ≤ `10^6`
- Sum of `q` over all test cases ≤ `10^6`

## Input

Each test contains multiple test cases. The first line contains the number of
test cases `t` (`1 ≤ t ≤ 10^4`). The description of the test cases follows.

The first line of each test case contains two integers `n` and `q`
(`1 ≤ n ≤ 10^6`, `q = 0`) — the length of the array `a` and the number of
updates.

The second line of each test case contains `n` integers
`a_0, a_1, ..., a_{n-1}` (`1 ≤ a_i ≤ 10^9`) — the array `a`.

The `j`-th of the following `q` lines contains two integers `i_j` and `x_j`
(`0 ≤ i_j < n`, `1 ≤ x_j ≤ 10^9`) — the description of the `j`-th update. This
update means that the assignment `a_{i_j} = x_j` is performed.

It is guaranteed that the sum of `n` over all test cases does not exceed `10^6`.

It is guaranteed that the sum of `q` over all test cases does not exceed `10^6`.

## Output

For each test case, output `q + 1` integers — the values of `f(a)` for the
initial state of the array and after each of the `q` updates, in order.

## Samples

### Sample 1

Input:

```
3
3 0
2 3 4
2 0
1000000000 999999999
6 0
2 5 3 4 1 6
```

Output:

```
0
1
4
```

## Note

In the first example, the array is `a = [2, 3, 4]`. It is already sorted, so
`f(a) = 0`.

In the second example, the array is `a = [10^9, 10^9 - 1]`. We can swap `a_0`
and `a_1`, transforming `a` as follows:
`[10^9, 10^9 - 1] → [10^9 - 1, 10^9]`. Therefore, `f(a) = 0 ⊕ 1 = 1`.

In the third example, the array is `a = [2, 5, 3, 4, 1, 6]`. We can perform
swaps using the index pairs `(0, 1)` and `(0, 4)`, transforming `a` as follows:
`[2, 5, 3, 4, 1, 6] → [5, 2, 3, 4, 1, 6]`,
`[5, 2, 3, 4, 1, 6] → [1, 2, 3, 4, 5, 6]`. It can be shown that no smaller
value of `k` is sufficient, so `f(a) = max(0 ⊕ 1, 0 ⊕ 4) = 4`.


### ideas
1. 0 ^ 1 = 1, 
2. 2 ^ 3 = 1
3. 1 ^ 2 = 3
4. 0 ^ 2 = 2
5. 0, 2, 1 => 1, 2, 0 => 2, 1, 0 => 0, 1, 2 (k <= 2)
6. 0, 1, 2, 4, 5, 3
7. 3 ^ 5 = 6 (110)
8. 4 ^ 5 = 1 (001)