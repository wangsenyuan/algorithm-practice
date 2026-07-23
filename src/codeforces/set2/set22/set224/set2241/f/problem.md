# F. A Bit Odd

[Problem link](https://codeforces.com/problemset/problem/2241/F)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Alice and Bob have got a binary string `s` of length `n`. They have decided to play a
game on it, taking turns alternately, with Alice moving first.

In each move, the player must select a subsequence which has an odd number of
inversions and delete it. The player who cannot make a move loses.

Determine who wins the game, assuming both players play optimally.

- A binary string consists only of the characters `0` and `1`.
- A sequence `a` is a subsequence of a string `b` if `a` can be obtained from `b` by
  the deletion of several (possibly zero or all) characters.
- An inversion in a binary string `s` is a pair of indices `(i, j)` such that
  `i < j`, `s_i = 1`, and `s_j = 0`.

## Constraints

- `1 ≤ t ≤ 10^4`
- `1 ≤ n ≤ 2 · 10^5`
- `s` consists only of `0` and `1`
- Sum of `n` over all test cases ≤ `2 · 10^5`

## Input

```
t
n
s
...
```

The first line contains a single integer `t` — the number of test cases.

Each test case consists of:

- a single integer `n` — the length of the binary string `s`
- a binary string `s` of length `n`

## Output

For each test case, print `Alice` if Alice wins the game and `Bob` otherwise.

## Samples

### Sample 1

Input:

```
3
5
10101
4
0100
6
011001
```

Output:

```
Alice
Alice
Bob
```

### Notes

- First test case: Alice can choose the entire string (odd number of inversions),
  leaving Bob with an empty string.
- Second test case: Alice can choose indices `1, 2, 4` forming `010`, leaving a
  single `0` with zero inversions.
- Third test case: Bob can guarantee a win regardless of Alice's first move.


## ideas
1. 如果一开始就是odd number inversions =》 alice
2. 假设目前是偶数逆序对, 删掉一个 10, 如果0的前面有w个1, 1的后面有v个0
3. 那么删除它们后逆序对 - w - v. 那么w和v的parity要一致, 这样才能继续保证是偶数的(否则马上失败了)
4. 假设存在这样的对, 删除掉后, 111100 如果剩余的 w是偶数, 那么alice获胜, 否则bob获胜

## Summary

Only the substring from the first `1` to the last `0` matters. Leading `0`s and
trailing `1`s cannot be part of any inversion, so they do not affect whether a
legal move exists.

After trimming this inactive border, split the remaining string into maximal
runs of equal characters. The game has a simple characterization:

- if every run has even length, the current player loses;
- if at least one run has odd length, the current player wins.

The all-even case is losing because any legal move, which must delete a
subsequence with an odd number of `1 ... 0` pairs, necessarily breaks the
all-even run structure and leaves the opponent a position with an odd run.

Conversely, if there is an odd run in the active middle, the current player can
choose an odd-inversion subsequence whose deletion restores the all-even run
structure for the opponent. Therefore such a position is winning.

So the implementation trims the inactive border, scans the active runs, and
returns `Alice` as soon as it finds an odd-length run. If no such run exists, it
returns `Bob`.

The complexity is `O(n)` per test case with `O(1)` extra working memory.
