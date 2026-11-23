# C2. Marenol (hard version)

[Problem link](https://codeforces.com/problemset/problem/2254/C2)

**Contest:** [Codeforces Round 1114 (Div. 3)](https://codeforces.com/contest/2254)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

This is the hard version of the problem. Here the task is the minimum number of operations that turn `a` into `b`.

## Problem

Yousef gives two binary strings `a` and `b` of the same length `n`. On `a`, either of these replacements may be used any number of times:

- replace a substring `001` with `100`, or `100` with `001`
- replace a substring `011` with `110`, or `110` with `011`

Print the minimum number of operations that make `a` equal to `b`. If no sequence of operations does it, print `-1`.

## Constraints

- `1 ≤ t ≤ 10^4`
- `1 ≤ n ≤ 2 · 10^5`
- `a` and `b` have length `n` and contain only `0` and `1`
- the sum of `n` over all test cases does not exceed `2 · 10^5`

## Input

The first line contains `t`.

Each test case is three lines:

```text
n
a
b
```

## Output

For each test case print one integer, the minimum number of operations, or `-1`.

## Sample Input 1

```text
4
0100
0001
```

## Sample Output 1

```text
1
```

`a[2..4] = 100` becomes `001`, so `0100` turns into `0001` in one operation.

## Sample Input 2

```text
4
0100
0010
```

## Sample Output 2

```text
-1
```

No sequence of operations turns `a` into `b`.

## Sample Input 3

```text
6
110000
000011
```

## Sample Output 3

```text
4
```

One minimum sequence is `110000 → 100100 → 100001 → 001001 → 000011`.

## Sample Input 4

```text
8
10101010
10101010
```

## Sample Output 4

```text
0
```

The strings are already equal.

## Sample Input 5

```text
5
01001
10010
```

## Sample Output 5

```text
3
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
