# B. Polo the Penguin and Matrix

[Problem link](https://codeforces.com/problemset/problem/289/B)

**Contest:** [Codeforces Round 177 (Div. 2)](https://codeforces.com/contest/289)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Polo has an `n × m` matrix of integers. In one move he can add or subtract `d` from a single matrix element. Find the minimum number of moves to make all elements equal. If it is impossible, print `-1`.

## Constraints

- `1 ≤ n, m ≤ 100`
- `1 ≤ d ≤ 10^4`
- `1 ≤ a_ij ≤ 10^4`

## Input

```text
n m d
a11 a12 ... a1m
...
an1 an2 ... anm
```

## Output

Print the minimum number of moves, or `-1` if impossible.

## Sample Input 1

```text
2 2 2
2 4
6 8
```

## Sample Output 1

```text
4
```

One optimal target is `4`: `(2→4)` costs 1, `(6→4)` costs 1, `(8→4)` costs 2, total 4.

## Sample Input 2

```text
1 2 7
6 7
```

## Sample Output 2

```text
-1
```

`6` and `7` differ by 1, which is not a multiple of `d = 7`.

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
