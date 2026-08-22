# E. Masha and the Garland

[Problem link](https://codeforces.com/problemset/problem/2244/E)

**Contest:** [Codeforces Round 1109 (Div. 3)](https://codeforces.com/contest/2244)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Masha has a New Year garland consisting of `n` bulbs. Each bulb can be either
off or on. The state of the garland is given by a binary string `s` of length
`n`, where `'0'` denotes an off bulb and `'1'` denotes an on bulb.

Masha considers the garland beautiful if the states of adjacent bulbs strictly
alternate. That is, there are no two adjacent bulbs that are both on or both
off. For example, the garlands `'01010'` and `'10101'` are beautiful, while
`'0110'` and `'000'` are not.

Yura can apply the following operation to the garland: choose a subsegment and
flip the state of all bulbs in it, that is, turn all on bulbs off and all off
bulbs on.

Masha invites him to play the following game `q` times: she chooses a segment
from the `l`-th to the `r`-th bulb inclusive, and Yura must make this segment
beautiful using at most `k` operations.

However, Yura is not sure whether this is always possible, so he asks you to
determine for each game whether he can make the chosen segment beautiful using
no more than `k` operations. Note that the games are independent, and the
garland itself is not actually modified.

## Constraints

- `1 ≤ t ≤ 10^4`
- `1 ≤ n, q ≤ 2 · 10^5`
- `1 ≤ l ≤ r ≤ n`
- `0 ≤ k ≤ n`
- `s` consists only of characters `'0'` and `'1'`
- Sum of `n` over all test cases does not exceed `2 · 10^5`
- Sum of `q` over all test cases does not exceed `2 · 10^5`

## Input

The first line contains a single integer `t` — the number of test cases.

The first line of each test case contains two integers `n` and `q` — the length
of the garland and the number of games.

The second line contains a binary string `s` of length `n`, consisting only of
characters `'0'` and `'1'`.

The next `q` lines each contain three integers `l`, `r`, and `k` — the
boundaries of the segment and the maximum allowed number of operations.

## Output

For each game, output `YES` if the segment can be made beautiful using at most
`k` operations, and `NO` otherwise.

## Example

### Input

```text
2
5 5
00110
1 5 1
1 5 2
2 4 1
1 2 0
3 4 0
4 2
1010
1 4 0
2 3 1
```

### Output

```text
YES
YES
YES
NO
NO
YES
YES
```
