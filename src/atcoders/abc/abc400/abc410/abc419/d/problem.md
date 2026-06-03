# D - Substr Swap (ABC419)

**Contest:** [ABC419](https://atcoder.jp/contests/abc419) — AtCoder Beginner Contest 419  
**Task:** [https://atcoder.jp/contests/abc419/tasks/abc419_d](https://atcoder.jp/contests/abc419/tasks/abc419_d)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 400 points

## Problem Statement

You are given length-`N` lowercase English strings `S` and `T`, and `M` pairs of
integers `(L_1, R_1), (L_2, R_2), ..., (L_M, R_M)`.

For `i = 1, 2, ..., M` in order, perform the following operation:

Swap the `L_i`-th through `R_i`-th characters of `S` and the `L_i`-th through
`R_i`-th characters of `T`.

For example, if `S` is `abcdef`, `T` is `ghijkl`, and `(L_i, R_i) = (3, 5)`, then
after the operation `S` becomes `abijkf` and `T` becomes `ghcdel`.

Find the string `S` after performing all `M` operations.

## Constraints

- `1 <= N <= 5 * 10^5`
- `1 <= M <= 2 * 10^5`
- `S` and `T` are length-`N` lowercase English strings
- `1 <= L_i <= R_i <= N`
- `N`, `M`, `L_i`, and `R_i` are integers

## Input

The input is given from standard input in the following format:

```text
N M
S
T
L_1 R_1
L_2 R_2
:
L_M R_M
```

## Output

Output the string `S` after performing all `M` operations.

## Sample Input 1

```text
5 3
apple
lemon
2 4
1 5
5 5
```

## Sample Output 1

```text
lpple
```

Initially, `S` and `T` are `apple` and `lemon`, respectively.

- After the operation for `i = 1`, `S` and `T` are `aemoe` and `lppln`.
- After the operation for `i = 2`, `S` and `T` are `lppln` and `aemoe`.
- After the operation for `i = 3`, `S` and `T` are `lpple` and `aemon`.

Thus, the string `S` after the three operations is `lpple`.

## Sample Input 2

```text
10 5
lemwrbogje
omsjbfggme
5 8
4 8
1 3
6 6
1 4
```

## Sample Output 2

```text
lemwrfogje
```
