# C. Annoying Game

[Problem link](https://codeforces.com/problemset/problem/2158/C)

**Contest:** [Codeforces Round 1067 (Div. 2)](https://codeforces.com/contest/2158)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given two integer arrays `a` and `b`, both of length `n`, and a total number of turns `k`.

Alice and Bob play a game by taking turns modifying array `a`. Alice goes first. The game lasts for exactly `k` turns.

On their turn, a player must choose an index `i` (`1 <= i <= n`) and perform one of the following operations:

- Add: increase `a_i` by `b_i`, i.e. set `a_i := a_i + b_i`.
- Subtract: decrease `a_i` by `b_i`, i.e. set `a_i := a_i - b_i`.

After the `k`-th turn is complete, the final score is the maximum non-empty subarray sum of the modified array `a`. Alice's goal is to maximize the final score, while Bob's goal is to minimize it.

Assuming both players play optimally, determine the final score.

The maximum non-empty subarray sum of `a` is `max_{1 <= i <= j <= n} S(i, j)`, where `S(i, j) = a_i + … + a_j`. Empty subarrays are not considered.

## Input

The first line contains the number of test cases `t` (`1 <= t <= 10^4`).

The first line of each test case contains two integers `n` and `k` (`1 <= n <= 2 * 10^5`, `1 <= k <= 2 * 10^5`) — the length of the arrays and the total number of turns.

The second line contains `n` integers `a_1, a_2, …, a_n` (`-10^9 <= a_i <= 10^9`).

The third line contains `n` integers `b_1, b_2, …, b_n` (`0 <= b_i <= 10^9`).

The sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, print a single integer — the final score after `k` turns, assuming both players play optimally.

## Example

### Input

```text
5
5 200000
3 -1 9 -5 4
0 0 0 0 0
4 5
10 10 10 10
1 1 1 1
3 1
2 -7 3
1 11 3
3 2
2 -7 3
1 11 3
1 1
-3
2
```

### Output

```text
11
41
9
3
-1
```

### Note

In the first test case every `b_i` is `0`, so moves cannot change `a`. The maximum non-empty subarray sum is `11`.

## Status

I/O and official samples are in place. `solve` is left as a TODO.
