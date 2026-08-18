# D. Array Replacement

[Problem link](https://codeforces.com/problemset/problem/2252/D)

**Contest:** [Codeforces Round 1115 (Div. 2)](https://codeforces.com/contest/2252)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given an array `a` of length `n`.

You can perform the following operation any number of times (possibly zero):

- Choose an index `i` (`2 ≤ i ≤ n − 1`) such that `a_{i−1}` and `a_{i+1}` have the same parity.
- Replace `a_i` with `a_{i−1} − a_i + a_{i+1}`.

Find the lexicographically smallest array that can be obtained after any number of operations.

A sequence `x` is lexicographically smaller than a sequence `y` of the same length if and only if, in the first position where `x` and `y` differ, the element in `x` is strictly smaller than the corresponding element in `y`.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 ≤ t ≤ 10^4`). The description of the test cases follows.

The first line of each test case contains a single integer `n` (`3 ≤ n ≤ 2 · 10^5`) — the length of the array `a`.

The second line of each test case contains `n` integers `a_1, a_2, …, a_n` (`−10^9 ≤ a_i ≤ 10^9`).

It is guaranteed that the sum of `n` over all test cases does not exceed `2 · 10^5`.

## Output

For each test case, output `n` integers — the lexicographically smallest array that can be obtained.

## Example

### Input

```text
3
10
100 108 114 118 120 5 7 19 13 11
3
1 2 3
4
10 10 8 4
```

### Output

```text
100 102 106 112 120 5 -1 -3 -1 11
1 2 3
10 6 4 4
```

### Note

In the second test case, the initial array is `[1, 2, 3]`. The only valid index to choose is `i = 2`, because `a_1 = 1` and `a_3 = 3` have the same parity. Replacing `a_2` with `1 − 2 + 3 = 2` leaves the array unchanged. Thus, the minimal array is `[1, 2, 3]`.

In the third test case, the initial array is `[10, 10, 8, 4]`. We can perform the following sequence of operations:

- Choose `i = 2` (`a_1 = 10` and `a_3 = 8` are both even). Replace `a_2` with `10 − 10 + 8 = 8`. The array becomes `[10, 8, 8, 4]`.
- Choose `i = 3` (`a_2 = 8` and `a_4 = 4` are both even). Replace `a_3` with `8 − 8 + 4 = 4`. The array becomes `[10, 8, 4, 4]`.
- Choose `i = 2` (`a_1 = 10` and `a_3 = 4` are both even). Replace `a_2` with `10 − 8 + 4 = 6`. The array becomes `[10, 6, 4, 4]`.

It can be shown that `[10, 6, 4, 4]` is the lexicographically smallest array obtainable.

## ideas
1. a[i-1] 和 a[i+1] 奇偶性相同, a[i-1] + a[i+1] 肯定是偶数
2. 所以 x = a[i-1] + a[i+1] - a[i], x和a[i]的奇偶性相同
3. a[1] 和 a[n]不变
4. 假设最先变的是a[i], x = a[i-1] + a[i+1] - a[i], x < a[i] (否则就没变要变?)
5. 然后改变a[i-1]后, y = a[i-2] + x - a[i-1] < a[i-1] 
6. a[i-2] + x < 2 * a[i-1] => a[i-2] + a[i-1] + a[i+1] < 2 *a[i-1] 
7. => a[i-2] + a[i+1] < a[i-1]
8. 但是这个时候, 就可以把a[i]又变小? 岂不是无限循环了
9. a[1] + a[3] - x >= 2 * (a[2] - x)
10. a[1] + a[3] >= 2 * a[2] - x
11. 