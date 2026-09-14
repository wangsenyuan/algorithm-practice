# D. MEX Multiset

[Problem link](https://codeforces.com/problemset/problem/2259/D)

**Contest:** [Codeforces Round 1119 (Div. 3)](https://codeforces.com/contest/2259)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

The MEX (minimum excluded value) of a collection of integers is the smallest non-negative integer that does not appear in it. For example, `mex([2, 2, 1]) = 0` and `mex([0, 3, 1, 2]) = 4`. The MEX of an empty collection is `0`.

You are given an array `a` of `n` non-negative integers. Assign each element to one of three labeled multisets `A`, `B`, and `C` (a set may be empty) so that

- `mex(A) = mex(B)`, and
- `mex(C) = 0`.

If several assignments are valid, print any of them.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 <= t <= 10^4`). The description of the test cases follows.

The first line of each test case contains a single integer `n` (`1 <= n <= 2 * 10^5`) — the length of the array.

The second line contains `n` integers `a_1, a_2, …, a_n` (`0 <= a_i <= 10^9`).

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, print `NO` if no assignment exists.

Otherwise print `YES` on the first line, and on the second line a string of length `n` consisting of characters `A`, `B`, and `C`, where the `i`-th character is the set that receives `a_i`.

You may print `YES` and `NO` in any case.

## Example

### Input

```text
5
6
1 0 0 1 2 1
4
0 0 0 0
3
0 2 2
4
6 7 6 7
5
0 0 0 1 2
```

### Output

```text
YES
CABCCC
YES
ABBB
NO
YES
AAAA
YES
ABBCC
```

One possible output is shown. Any valid assignment is accepted.

## Status

I/O and official sample cases are in place. `solve` is left as a TODO.
