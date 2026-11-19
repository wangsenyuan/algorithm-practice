# C. Constanze's Machine

[Problem link](https://codeforces.com/problemset/problem/1245/C)

**Contest:** [Codeforces Round 597 (Div. 2)](https://codeforces.com/contest/1245)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

The machine writes each spoken letter as itself, except for two letters:

- pronouncing `w` writes `uu`
- pronouncing `m` writes `nn`

The received message is `s`. Count the strings the machine could have turned into `s`. Print that count modulo `10^9+7`. If no such string exists, print `0`.

## Constraints

- `1 ≤ |s| ≤ 10^5`
- `s` contains only lowercase Latin letters

## Input

```text
s
```

## Output

Print one integer, the number of original strings modulo `10^9+7`.

## Sample Input 1

```text
ouuokarinn
```

## Sample Output 1

```text
4
```

The four strings are `ouuokarinn`, `ouuokarim`, `owokarim`, and `owokarinn`.

## Sample Input 2

```text
banana
```

## Sample Output 2

```text
1
```

## Sample Input 3

```text
nnn
```

## Sample Output 3

```text
3
```

The three strings are `nm`, `mn`, and `nnn`.

## Sample Input 4

```text
amanda
```

## Sample Output 4

```text
0
```

The machine never writes `m`.

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
