# C. News Distribution

[Problem link](https://codeforces.com/problemset/problem/1167/C)

**Contest:** [Educational Codeforces Round 65 (Rated for Div. 2)](https://codeforces.com/contest/1167)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

A social network has `n` users and `m` groups of friends. Two users are friends when they belong to at least one common group.

News spreads from a user to that user's friends, then to their friends, and so on. Spreading stops when every friend of someone who already knows the news also knows it.

For each user `x`, print how many users know the news if only `x` starts with it.

## Constraints

- `1 ≤ n, m ≤ 5 · 10^5`
- the `i`-th group has `k_i` users, `0 ≤ k_i ≤ n`
- the `k_i` user ids in a group are distinct and lie in `1..n`
- the sum of `k_i` does not exceed `5 · 10^5`

## Input

```text
n m
k_1 x_1 x_2 ... x_{k_1}
...
k_m x_1 x_2 ... x_{k_m}
```

## Output

Print `n` integers. The `i`-th integer is the number of users who know the news when user `i` starts distributing it.

## Sample Input 1

```text
7 5
3 2 5 4
0
2 1 2
1 1
2 6 7
```

## Sample Output 1

```text
4 4 1 4 4 2 2
```

Users `2`, `5`, and `4` share a group, and user `1` shares a group with user `2`, so `{1, 2, 4, 5}` has size `4`. User `3` is alone. Users `6` and `7` form a component of size `2`.

## Status

I/O and the official sample are in place. `solve` is left as a TODO.
