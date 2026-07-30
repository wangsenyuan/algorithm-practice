# D - Good Tuple Problem

[Problem link](https://atcoder.jp/contests/abc327/tasks/abc327_d)

**Contest:** [AtCoder Beginner Contest 327](https://atcoder.jp/contests/abc327)

time limit: 2 sec

memory limit: 1024 MiB

score: 400 points

## Problem Statement

A pair of sequences of length `M` consisting of positive integers at most `N`,
`(S, T) = ((S_1, ..., S_M), (T_1, ..., T_M))`, is said to **be a good pair of
sequences** when there exists a sequence `X = (X_1, ..., X_N)` of length `N`
consisting of `0` and `1` such that:

- `X_{S_i} != X_{T_i}` for each `i = 1, 2, ..., M`.

You are given a pair of sequences `(A, B)`. If `(A, B)` is a good pair of
sequences, print `Yes`; otherwise, print `No`.

## Constraints

- `1 <= N, M <= 2 * 10^5`
- `1 <= A_i, B_i <= N`
- All input values are integers.

## Input

```text
N M
A_1 A_2 ... A_M
B_1 B_2 ... B_M
```

## Output

If `(A, B)` is a good pair of sequences, print `Yes`; otherwise, print `No`.

## Sample Input 1

```text
3 2
1 2
2 3
```

## Sample Output 1

```text
Yes
```

`X = (0, 1, 0)` works.

## Sample Input 2

```text
3 3
1 2 3
2 3 1
```

## Sample Output 2

```text
No
```

## Sample Input 3

```text
10 1
1
1
```

## Sample Output 3

```text
No
```

## Sample Input 4

```text
7 8
1 6 2 7 5 4 2 2
3 2 7 2 1 2 3 3
```

## Sample Output 4

```text
Yes
```

## ideas
1. 按照图的逻辑处理