# D. Replace with Occurrences

[Problem link](https://codeforces.com/problemset/problem/2137/D)

**Contest:** [Codeforces Round 1047 (Div. 3)](https://codeforces.com/contest/2137)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

For an array `a`, let `f(x)` be the number of times `x` occurs in `a`. Given `b` of length `n`, construct an array `a` of length `n` such that `f(a_i) = b_i` for every index `i`. Each `a_i` must lie between `1` and `n`. If no such array exists, report that.

## Constraints

- `1 ≤ t ≤ 10^4`
- `1 ≤ n ≤ 2 · 10^5`
- `1 ≤ b_i ≤ n`
- The sum of `n` over all test cases does not exceed `2 · 10^5`

## Input

The first line contains `t`, the number of test cases.

Each test case is two lines:

```text
n
b_1 b_2 ... b_n
```

## Output

For each test case, print `-1` if no array exists. Otherwise print `n` integers `a_1 a_2 ... a_n`. Any valid array is accepted.

## Sample Input 1

```text
4
1 2 3 4
```

## Sample Output 1

```text
-1
```

No array matches `b`.

## Sample Input 2

```text
6
1 2 2 3 3 3
```

## Sample Output 2

```text
4 5 5 6 6 6
```

`4`, `5`, and `6` occur `1`, `2`, and `3` times, so the frequencies are `1, 2, 2, 3, 3, 3`.

## Sample Input 3

```text
6
6 6 6 6 6 6
```

## Sample Output 3

```text
2 2 2 2 2 2
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO. A nil result prints `-1`. The tests accept any array whose frequencies match `b`.
