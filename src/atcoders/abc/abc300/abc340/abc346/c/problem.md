# C - Σ (ABC346)

**Contest:** [ABC346](https://atcoder.jp/contests/abc346) — AtCoder Beginner Contest 346  
**Task:** [https://atcoder.jp/contests/abc346/tasks/abc346_c](https://atcoder.jp/contests/abc346/tasks/abc346_c)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 250 points

## Problem Statement

You are given a length-`N` sequence of positive integers `A = (A_1, …, A_N)` and a
positive integer `K`.

Find the sum of integers from `1` to `K` (inclusive) that do **not** appear in `A`.

## Constraints

- `1 <= N <= 2 × 10^5`
- `1 <= K <= 2 × 10^9`
- `1 <= A_i <= 2 × 10^9`
- All input values are integers

## Input

```text
N K
A_1 A_2 … A_N
```

## Output

Print the answer.

## Sample Input 1

```text
4 5
1 6 3 1
```

## Sample Output 1

```text
11
```

Among `1..5`, the values not in `A` are `2`, `4`, and `5`, so the answer is
`2 + 4 + 5 = 11`.

## Sample Input 2

```text
1 3
346
```

## Sample Output 2

```text
6
```

## Sample Input 3

```text
10 158260522
877914575 24979445 623690081 262703497 24979445 1822804784 1430302156 1161735902 923078537 1189330739
```

## Sample Output 3

```text
12523196466007058
```

## Solution

Answer = `(1 + 2 + … + K) − sum of distinct values in A that lie in [1, K]`.

Use `int64` for the arithmetic.
