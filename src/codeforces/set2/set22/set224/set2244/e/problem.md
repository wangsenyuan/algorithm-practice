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

## Solution

Goal: make `s[l..r]` strictly alternating with at most `k` subsegment
flips. A window is beautiful iff it matches one of the two globally
aligned targets `0101…` or `1010…` (`pat[i&1]` on the absolute index).
Trying both covers both relative patterns on any `[l, r]`.

### 1. Mismatch runs

Against one target, set `val[i] = 1` on a mismatch. One flip covers a
whole maximal mismatch run (and nothing else), so the min number of
flips equals the number of those runs.

### 2. Prefix of run starts

`pref[i]` is the number of run starts on `[0..i]`. Sweep left to right
and increment `cnt` when a mismatch begins (`val[i] = 1` and
`i == 0` or `val[i-1] == 0`).

### 3. Query on `[l, r]` (0-index)

| piece | meaning |
|---|---|
| `k1 = pref[r] − pref[l−1]` | run starts inside the window |
| `+1` if `val[l] = 1`, `l > 0`, and `val[l−1] = 1` | left cut splits a run; the leftover inside the window is an extra run |

`YES` iff `k1 ≤ k` for at least one of the two targets.

### 4. Offline sweep

Bucket queries by `r`, then one left-to-right pass per target answers
every query in `O(n + q)`.

### One-liner

Match against global `01` and `10`; min flips = mismatch-run count,
with a `+1` when the left boundary splits a run.
