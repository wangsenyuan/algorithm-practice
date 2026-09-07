# C2. We Be Flipping (Hard Version)

[Problem link](https://codeforces.com/problemset/problem/2229/C2)

**Contest:** [Spectral::Cup 2026 Round 2 (Codeforces Round 1100, Div. 1 + Div. 2)](https://codeforces.com/contest/2229)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

This is the hard version of the problem. The difference between the versions is that in this version, you must maximise the sum. You can hack only if you solved all versions of this problem.

You have an array `a` of length `n` which consists of non-zero (but possibly negative) integers. You will perform the following operation at most `n` times (possibly none):

- select an index `i` (`1 <= i <= n`) such that `a_i > 0`
- then for each `j` where `1 <= j <= i` do `a_j := -a_j`

Output a valid sequence of operations of length at most `n` which maximises the sum of `a` at the end.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 <= t <= 10^4`). The description of the test cases follows.

The first line of each test case contains an integer `n` (`2 <= n <= 2 * 10^5`) — the length of the array `a`.

The second line of each test case contains `n` integers `a_1, a_2, …, a_n` (`-10^9 <= a_i <= 10^9`, `a_i != 0`).

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output a single integer `k` (`0 <= k <= n`) — the number of operations you will perform.

Then output `k` integers `b_1, …, b_k` where `b_i` is the index you perform the `i`-th operation on.

After performing the operations the sum of `a` should be maximal.

## Example

### Input

```text
5
5
-1 -2 -3 -5 -4
4
5 7 10 19
5
1 -3 2 -1 10
4
16 -13 -18 -16
11
2 -10 -11 3 -10 15 7 18 16 17 -9
```

### Output

```text
0
0
2
1 3
0
6
6 3 1 5 4 7
```

### Note

In the first test case, no operations are possible.

In the second test case, the sum is already maximal.

In the third test case, operations are made as follows:

- `[1, -3, 2, -1, 10]` with `i = 1` becomes `[-1, -3, 2, -1, 10]`
- `[-1, -3, 2, -1, 10]` with `i = 3` becomes `[1, 3, -2, -1, 10]`

This has sum `11`, which can be proven to be maximal.

## Status

I/O and official samples are in place. `solve` is left as a TODO.
