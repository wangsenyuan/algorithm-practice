# A. Sonya and Queries

[Problem link](https://codeforces.com/problemset/problem/713/A)

**Contest:** [Codeforces Round 371 (Div. 1)](https://codeforces.com/contest/713)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Sonya starts with an empty multiset of non-negative integers and processes `t` queries:

- `+ ai` inserts one copy of `ai`.
- `- ai` deletes one copy of `ai`. The multiset always contains at least one copy.
- `? s` counts how many integers currently match the pattern `s`.

The pattern consists of `0` and `1`. `0` means an even digit and `1` means an odd digit. Digits are compared from the right. If the pattern is shorter, it is padded with `0` on the left. If the integer is shorter, it is padded with `0` on the left.

For example, `010` matches `92`, `2212`, `50`, and `414`. It does not match `3`, `110`, `25`, or `1030`.

For each `?` query, print the count. Repeated values are counted once per copy.

## Constraints

- `1 ≤ t ≤ 10^5`
- There is at least one `?` query
- `0 ≤ ai < 10^18`, written without leading zeros unless the number is `0`
- Each pattern has length at most `18`

## Input

```text
t
c_1 arg_1
c_2 arg_2
...
c_t arg_t
```

`c_i` is `+`, `-`, or `?`.

## Output

Print one integer per `?` query.

## Sample Input 1

```text
12
+ 1
+ 241
? 1
+ 361
- 241
? 0101
+ 101
? 101
- 101
? 101
+ 4000
? 0
```

## Sample Output 1

```text
2
1
2
1
1
```

The five patterns match, in order: `1` and `241`; `361`; `101` and `361`; `361`; `4000`.

## Sample Input 2

```text
4
+ 200
+ 200
- 200
? 0
```

## Sample Output 2

```text
1
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
