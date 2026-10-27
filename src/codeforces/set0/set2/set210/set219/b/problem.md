# B. Special Offer! Super Price 999 Bourles!

[Problem link](https://codeforces.com/problemset/problem/219/B)

**Contest:** [Codeforces Round 135 (Div. 2)](https://codeforces.com/contest/219)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

The starting price is `p` bourles. The price may be lowered by at most `d` bourles. Among prices `x` with `p - d ≤ x ≤ p`, choose the one with the largest number of trailing nines. If several prices tie, choose the largest one.

Only nines at the end of the decimal representation count. The printed price has no leading zeros.

## Constraints

- `1 ≤ p ≤ 10^18`
- `0 ≤ d < p`

## Input

```text
p d
```

## Output

Print the chosen price.

## Sample Input 1

```text
1029 102
```

## Sample Output 1

```text
999
```

## Sample Input 2

```text
27191 17
```

## Sample Output 2

```text
27189
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
