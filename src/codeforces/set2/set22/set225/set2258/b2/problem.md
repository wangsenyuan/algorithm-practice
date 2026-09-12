# B2. Carrot Chopdown (Hard Version)

[Problem link](https://codeforces.com/problemset/problem/2258/B2)

**Contest:** [Codeforces Round 1118 (Div. 2)](https://codeforces.com/contest/2258)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

This is the hard version of the problem. The difference between the versions is that in this version, you need to solve the problem for different values of `k`. You can hack only if you solved all versions of this problem.

You are given `n` delicious carrots with sizes `a_1, a_2, …, a_n`. You are also given a cutting machine, which works as follows.

- For each operation, you choose a set of carrots (you can choose chopped carrots again) and a positive integer `x` (not necessarily the same for each operation).
- After that, consider every chosen carrot, let its length be `l`. If `l <= x`, this carrot is unaffected; otherwise, it is divided into two carrots of sizes `x` and `l - x`.

We'll sell some of the final carrots to an interesting guy who wants them all to be the same length.

We are asking you to determine the maximum number of carrots we can sell after using this machine exactly `k` times. Solve the problem for each `k = 1, 2, …, m` independently.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 <= t <= 10^4`). The description of the test cases follows.

The first line of each test case contains `n` and `m` (`1 <= n, m <= 2 * 10^5`), denoting the number of carrots and the maximum possible length of a carrot.

The second line of each test case contains `n` integers `a_1, a_2, …, a_n` (`1 <= a_i <= m`), denoting the initial carrot sizes.

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5` and the sum of `m` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output `m` integers — the answer for each `k = 1, 2, …, m`.

## Example

### Input

```text
6
5 4
1 2 3 4 4
5 8
1 1 8 8 8
1 8
6
7 9
1 7 5 1 7 5 3
4 1
1 1 1 1
3 5
3 1 5
```

### Output

```text
6 14 14 14
6 12 26 26 26 26 26 26
2 3 6 6 6 6 6 6
7 17 29 29 29 29 29 29 29
4
3 7 9 9 9
```

### Note

In the first test case, the given carrots are `[1, 2, 3, 4, 4]`.

For `k = 1`, it is best to choose `x = 2` with the set `[2, 3, 4, 4]`. After the operation, we'll get `[1, 2, 2, 1, 2, 2, 2, 2]`. We can sell `6` carrots of length `2`.

For `k = 2`, it is best to choose `x = 2` with the set `[2, 3, 4, 4]`. After the operation, we'll get `[1, 2, 2, 1, 2, 2, 2, 2]`. For the second operation, it is best to choose `x = 1` with the set of all of the carrots. Resulting in `14` carrots, all of the same length. We can sell `14` carrots of length `1`.

## Status

I/O and official samples are in place. `solve` is left as a TODO.
