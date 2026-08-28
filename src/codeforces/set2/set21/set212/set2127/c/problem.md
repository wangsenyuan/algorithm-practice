# C. Trip Shopping

[Problem link](https://codeforces.com/problemset/problem/2127/C)

**Contest:** [Atto Round 1 (Codeforces Round 1041, Div. 1 + Div. 2)](https://codeforces.com/contest/2127)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Ali and Bahamin decided to spend their summer vacation on the beautiful southern coasts of Iran. They also agreed to do some shopping during the trip — but instead of setting a fixed budget, they decided to determine how much they would spend by playing a game.

The game is played on two arrays `a` and `b`, each containing `n` integers.

The game will last for `k` rounds. In one round:

- First, Ali selects two indices `i` and `j` (`1 <= i < j <= n`);
- Then, Bahamin rearranges the four integers `a_i`, `a_j`, `b_i`, and `b_j` arbitrarily. Note that Bahamin can swap numbers between two arrays. He can also keep the two arrays unchanged.

After all the `k` rounds, the value of the game is defined as `v = sum |a_i - b_i|` over `1 <= i <= n`. Ali and Bahamin will spend exactly `v` coins during their trip.

However, their goals are quite different:

- Ali wants to spend as little as possible, that is, to minimize `v`;
- Bahamin wants to spend as much as possible, that is, to maximize `v`.

You have to find the final amount of coins they will spend if both Ali and Bahamin play optimally.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases.

Each test case contains:

- one line with two integers `n` and `k` (`2 <= n <= 2 * 10^5`, `1 <= k <= n`) — the length of `a` and `b`, and the number of rounds;
- one line with `n` integers `a_1, a_2, ..., a_n` (`1 <= a_i <= 10^9`) — the elements of `a`;
- one line with `n` integers `b_1, b_2, ..., b_n` (`1 <= b_i <= 10^9`) — the elements of `b`.

The sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, print one integer — the final amount of coins they will spend if both Ali and Bahamin play optimally.

## Example

### Input

```text
5
2 1
1 7
3 5
3 2
1 5 3
6 2 4
5 4
1 16 10 10 16
3 2 2 15 15
4 1
23 1 18 4
19 2 10 3
10 10
4 3 2 100 4 1 2 4 5 5
1 200 4 5 6 1 10 2 3 4
```

### Output

```text
8
9
30
16
312
```

### Note

In the first test case, Ali can only choose `(i, j) = (1, 2)`, and Bahamin can rearrange all four numbers. Thus, he can assign `a = [5, 1]` and `b = [3, 7]`. And the value of the game will be `v = |5 - 3| + |1 - 7| = 8`. It can be shown that this is the maximum possible value reachable for Bahamin — other arrangements like `a = [5, 7]`, `b = [1, 3]` are also possible, but they don't have larger values.

In the second test case, the best strategy for Bahamin is to keep the two arrays unchanged, regardless of what indices Ali selects. And the value of the game will be `v = |1 - 6| + |5 - 2| + |3 - 4| = 9`.

## Status

I/O and official samples are in place. `solve` is left as a TODO.
