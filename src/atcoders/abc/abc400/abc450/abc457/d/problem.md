# D - Raise Minimum (ABC457)

**Contest:** [ABC457](https://atcoder.jp/contests/abc457) — AtCoder Beginner Contest 457  
**Task:** [https://atcoder.jp/contests/abc457/tasks/abc457_d](https://atcoder.jp/contests/abc457/tasks/abc457_d)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 400 points

## Problem Statement

You are given an integer sequence `A = (A_1, A_2, ..., A_N)` and an integer `K`.

You may perform the following operation at most `K` times:

- Choose an index `i` with `1 <= i <= N`, and add `i` to `A_i`.

After all operations, consider the minimum value among all elements of the
sequence. Find the maximum possible value of this minimum.

## Constraints

- `1 <= N <= 2 * 10^5`
- `1 <= A_i <= 10^18`
- `1 <= K <= 10^18`
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```text
N K
A_1 A_2 ... A_N
```

## Output

Print the maximum possible value of `min(A_1, A_2, ..., A_N)` after performing
at most `K` operations.

## Sample Input 1

```text
3 3
1 2 3
```

## Sample Output 1

```text
3
```

For example, choose `i = 1` twice and `i = 2` once. The sequence becomes
`(3, 4, 3)`, whose minimum is `3`. It is impossible to make every element at
least `4`.

## Sample Input 2

```text
4 5
10 1 10 1
```

## Sample Output 2

```text
7
```

## Sample Input 3

```text
20 457
8 9 10 9 8 8 4 6 8 1 5 10 2 8 2 6 8 1 6 6
```

## Sample Output 3

```text
132
```
