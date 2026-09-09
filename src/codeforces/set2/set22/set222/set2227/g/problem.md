# G. Drowning

[Problem link](https://codeforces.com/problemset/problem/2227/G)

**Contest:** [Codeforces Round 1096 (Div. 3)](https://codeforces.com/contest/2227)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Yousef has an array `a` consisting of `n` positive integers.

He defines a reduction operation on any array `c` of length `|c| >= 3`:

- Choose an index `i` (`1 < i < |c|`) such that `c_{i-1} + c_{i+1} > c_i`.
- Replace the triplet `{c_{i-1}, c_i, c_{i+1}}` with a single integer `x = c_{i-1} - c_i + c_{i+1}`.

The new integer `x` occupies the position previously held by the triplet, and the length of the array decreases by `2`.

An array is considered good if it can be reduced to a single element by performing the operation above zero or more times. Note that an array of length `1` is always good.

Yousef wants you to count the number of pairs `(l, r)` (`1 <= l <= r <= n`) such that the subarray `a[l, r]` is good.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases. The description of each test case follows.

The first line of each test case contains an integer `n` (`1 <= n <= 2 * 10^5`) — the size of the array.

The second line contains `n` integers `a_1, a_2, …, a_n` (`1 <= a_i <= 10^9`) — the elements of the array.

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output a single integer — the number of good subarrays.

## Example

### Input

```text
4
3
10 20 10
5
1 1 1 1 1
4
5 1 5 1
1
100
```

### Output

```text
3
9
5
1
```

### Note

In the first example, `a = [10, 20, 10]`. Subarrays `[10]`, `[20]`, and `[10]` are all good (`3` total). The subarray `[10, 20, 10]` is not good. To reduce it, we must pick `i = 2`. The condition `a_1 + a_3 > a_2` becomes `10 + 10 > 20`, which is `20 > 20` (false).

In the second example, `a = [1, 1, 1, 1, 1]`:

- All `5` subarrays of length `1` are good.
- All `4` subarrays of length `2` are not good.
- All `3` subarrays of length `3` (which are `[1, 1, 1]`) are good because `1 + 1 > 1`.
- All `2` subarrays of length `4` are not good.
- The subarray of length `5` is good: `[1, 1, 1, 1, 1]` with `i = 2` becomes `[1, 1, 1]`, then with `i = 2` becomes `[1]`.
- Total good subarrays `= 5 + 3 + 1 = 9`.

## Status

I/O and official samples are in place. `solve` is left as a TODO.
