# F. Financiers Game

[Problem link](https://codeforces.com/problemset/problem/729/F)

**Contest:** [Technocup 2017 - Elimination Round 2](https://codeforces.com/contest/729)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

## Problem

There are `n` papers in a row. Paper `i` has income `a_i`, which may be positive, zero, or negative.

Igor and Zhenya take turns. Igor takes papers from the left and moves first. Zhenya takes papers from the right. Igor starts by taking `1` or `2` papers. After a player takes exactly `k` papers, the opponent must take `k` or `k + 1` papers from that opponent's side. A player cannot skip a turn.

The game ends when no papers remain, or when the player to move cannot take enough papers.

Both play optimally. Igor maximizes the difference between the sum of incomes he took and the sum Zhenya took. Zhenya minimizes that difference. Print the resulting difference.

## Constraints

- `1 ≤ n ≤ 4000`
- `-10^5 ≤ a_i ≤ 10^5`

## Input

```text
n
a_1 a_2 ... a_n
```

## Output

Print one integer, the optimal difference.

## Sample Input 1

```text
3
1 3 1
```

## Sample Output 1

```text
4
```

Igor takes the first two papers, sum `4`. One paper remains, and Zhenya would need to take `2` or `3`, so Zhenya cannot move.

## Sample Input 2

```text
5
-1 -2 -1 -2 -1
```

## Sample Output 2

```text
0
```

## Sample Input 3

```text
4
-4 -2 4 5
```

## Sample Output 3

```text
-13
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
