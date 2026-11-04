# A. Adding Digits

[Problem link](https://codeforces.com/problemset/problem/260/A)

**Contest:** [Codeforces Round 158 (Div. 2)](https://codeforces.com/contest/260)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Vasya has integers `a` and `b`. He lengthens `a` exactly `n` times. One lengthening appends a single decimal digit on the right, and the number after the append must be divisible by `b`.

Print any integer he can obtain, without leading zeros. If no such integer exists, print `-1`.

## Constraints

- `1 ≤ a, b, n ≤ 10^5`

## Input

```text
a b n
```

## Output

Print the integer, or `-1`.

## Sample Input 1

```text
5 4 5
```

## Sample Output 1

```text
524848
```

## Sample Input 2

```text
12 11 1
```

## Sample Output 2

```text
121
```

## Sample Input 3

```text
260 150 10
```

## Sample Output 3

```text
-1
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
