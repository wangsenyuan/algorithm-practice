# D - Minimize Range (ABC450)

**Contest:** [ABC450](https://atcoder.jp/contests/abc450) — AtCoder Beginner Contest 450  
**Task:** [https://atcoder.jp/contests/abc450/tasks/abc450_d](https://atcoder.jp/contests/abc450/tasks/abc450_d)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 400 points

## Problem Statement

You are given a sequence `A` of `N` positive integers and a positive integer `K`. You
may perform the following operation on `A` any number of times:

- Choose an integer `i` with `1 <= i <= N`, and add `K` to `A_i`.

Find the minimum possible value of `max(A) - min(A)`.

## Constraints

- `1 <= N <= 2 * 10^5`
- `1 <= K <= 10^9`
- `1 <= A_i <= 10^9`
- All input values are integers.

## Input

The input is given from standard input in the following format:

```text
N K
A_1 A_2 ... A_N
```

## Output

Print the answer on a single line.

## Sample Input 1

```text
3 10
3 21 9
```

## Sample Output 1

```text
4
```

One optimal sequence of operations:

1. Choose `i = 1`: `A = (13, 21, 9)`
2. Choose `i = 3`: `A = (13, 21, 19)`
3. Choose `i = 1`: `A = (23, 21, 19)`

Then `max(A) - min(A) = 23 - 19 = 4`. It is impossible to make
`max(A) - min(A) <= 3`, so the answer is `4`.

## Sample Input 2

```text
5 6
4 100 5 10 450
```

## Sample Output 2

```text
2
```
