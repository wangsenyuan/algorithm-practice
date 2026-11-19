# B. Universal Solution

[Problem link](https://codeforces.com/problemset/problem/1380/B)

**Contest:** [Educational Codeforces Round 91 (Rated for Div. 2)](https://codeforces.com/contest/1380)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

A bot plays rock-paper-scissors with strategy `s`, a string of length `n`. `R` is rock, `S` is scissors, and `P` is paper. Rock beats scissors, scissors beat paper, and paper beats rock.

Choose your own string `t` of length `n`. The bot is then started at every position `pos` from `1` to `n`. In the game that starts at `pos`, your `i`-th move `t_i` plays against `s` shifted by `pos`. `win(pos)` is the number of moves you win in that game.

Maximize the average

```text
(win(1) + win(2) + ... + win(n)) / n
```

Print any string `t` that achieves the maximum.

## Constraints

- `1 ≤ t ≤ 1000`
- `1 ≤ n ≤ 2 · 10^5`
- `s` has length `n` and contains only `R`, `S`, and `P`
- The sum of `n` over all test cases does not exceed `2 · 10^5`

## Input

The first line contains `t`, the number of test cases.

Each test case is one line, the string `s`.

## Output

For each test case, print one string of length `n`. If several answers are optimal, print any of them.

## Sample Input 1

```text
RRRR
```

## Sample Output 1

```text
PPPP
```

Paper is the only move that beats rock, and every character of `s` is rock.

## Sample Input 2

```text
RSP
```

## Sample Output 2

```text
RSP
```

One optimal string is `RSP`. The three alignments win `0`, `3`, and `0` games, so the average is `1`. `PPP`, `RRR`, and `SSS` are optimal as well.

## Sample Input 3

```text
S
```

## Sample Output 3

```text
R
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
