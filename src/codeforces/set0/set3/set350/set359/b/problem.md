# B. Permutation

[Problem link](https://codeforces.com/problemset/problem/359/B)

**Contest:** [Codeforces Round 209 (Div. 2)](https://codeforces.com/contest/359)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Find a permutation `a` of length `2n`, using each integer from `1` to `2n` once, such that

```text
|a_1 - a_2| + |a_3 - a_4| + ... + |a_{2n-1} - a_{2n}|
    - | (a_1 - a_2) + (a_3 - a_4) + ... + (a_{2n-1} - a_{2n}) |
    = 2k
```

A solution exists. Any valid permutation is accepted.

## Constraints

- `1 ≤ n ≤ 50000`
- `0 ≤ 2k ≤ n`

## Input

```text
n k
```

## Output

Print `2n` integers, the permutation.

## Sample Input 1

```text
1 0
```

## Sample Output 1

```text
1 2
```

`|1 - 2| - |1 - 2| = 0`.

## Sample Input 2

```text
2 1
```

## Sample Output 2

```text
3 2 1 4
```

`|3 - 2| + |1 - 4| - |3 - 2 + 1 - 4| = 2`.

## Sample Input 3

```text
4 0
```

## Sample Output 3

```text
2 7 4 6 1 3 5 8
```

The paired absolute values sum to `12`, and the absolute value of the signed sum is also `12`.

## Status

I/O and the official samples are in place. `solve` is left as a TODO. The tests accept any permutation that satisfies the equation.
