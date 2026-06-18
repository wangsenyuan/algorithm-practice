# E - x + y ≡ x + y (ABC460)

**Contest:** [ABC460](https://atcoder.jp/contests/abc460) — AtCoder Beginner Contest 460  
**Task:** [https://atcoder.jp/contests/abc460/tasks/abc460_e](https://atcoder.jp/contests/abc460/tasks/abc460_e)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 450 points

## Problem Statement

For positive integers `a` and `b`, define `concat(a, b)` as the integer formed by
writing `a` and `b` in decimal and concatenating them.

Formally, let `A` and `B` be the decimal strings of `a` and `b`. The integer
represented by the string `A + B` is `concat(a, b)`.

Example: `concat(123, 45) = 12345`.

You are given positive integers `N` and `M`. Count pairs of positive integers
`(x, y)` with `x <= N` and `y <= N` such that:

```text
concat(x, y) ≡ x + y   (mod M)
```

Output the count modulo `998244353`.

You are given `T` test cases; solve each one.

## Constraints

- `1 <= T <= 10^4`
- `1 <= N <= 10^18`
- `2 <= M <= 10^9`
- All input values are integers

## Input

The input is given from Standard Input in the following format:

```text
T
case_1
case_2
⋮
case_T
```

Each test case:

```text
N M
```

## Output

Print `T` lines. The `i`-th line is the answer for the `i`-th test case, modulo
`998244353`.

## Sample Input 1

```text
4
3 2
123 456
20260530 460
123456789123456789 998244353
```

## Sample Output 1

```text
3
0
922576091
422081792
```

For the first test case, the valid pairs are `(2, 1)`, `(2, 2)`, and `(2, 3)` — three
in total.

