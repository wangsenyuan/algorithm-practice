# C. Soldier and Cards

[Problem link](https://codeforces.com/problemset/problem/546/C)

**Contest:** [Codeforces Round 304 (Div. 2)](https://codeforces.com/contest/546)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Two soldiers play a war-like card game. There are `n` cards with distinct values `1` through `n`, split between them. The stacks may have different lengths.

Each fight, both players take the top card of their stack. The larger card wins. The winner puts the opponent's card at the bottom of his stack, then puts his own card under that. If a player's stack becomes empty, he loses.

Print the number of fights and the winner (`1` or `2`). If the game never ends, print `-1`.

## Constraints

- `2 ≤ n ≤ 10`
- `1 ≤ k1 ≤ n-1` and `k1 + k2 = n`
- The `n` card values are a permutation of `1..n`

## Input

```text
n
k1 a1 a2 ... ak1
k2 b1 b2 ... bk2
```

Cards are listed from the top of the stack to the bottom.

## Output

Print `fights winner`, or `-1`.

## Sample Input 1

```text
4
2 1 3
2 4 2
```

## Sample Output 1

```text
6 2
```

## Sample Input 2

```text
3
1 2
2 1 3
```

## Sample Output 2

```text
-1
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
