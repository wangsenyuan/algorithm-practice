# D. Exceptional Segments

[Problem link](https://codeforces.com/problemset/problem/2225/D)

**Contest:** [Educational Codeforces Round 189 (Rated for Div. 2)](https://codeforces.com/contest/2225)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

## Problem

You are given two integers `n` and `x`.

Consider the sequence `[1, 2, 3, …, n]`. You need to find the number of its
subsegments that contain `x` and have XOR equal to `0`. In other words, you need
to count the number of pairs `(l, r)` such that `1 ≤ l ≤ x ≤ r ≤ n` and
`l ⊕ (l + 1) ⊕ ⋯ ⊕ r = 0`, where `⊕` denotes the bitwise exclusive OR.

For example, if `n = 7` and `x = 6`, then the following segments are suitable:

- `(4, 7)`, because `x` lies in this segment and `4 ⊕ 5 ⊕ 6 ⊕ 7 = 0`
- `(1, 7)`, because `x` lies in this segment and
  `1 ⊕ 2 ⊕ 3 ⊕ 4 ⊕ 5 ⊕ 6 ⊕ 7 = 0`

Since the answer can be very large, output it modulo `998244353`.

## Constraints

- `1 ≤ t ≤ 2 · 10^5`
- `1 ≤ x ≤ n ≤ 10^18`

## Input

Each test contains multiple test cases. The first line contains the number of
test cases `t` (`1 ≤ t ≤ 2 · 10^5`). The description of the test cases follows.

The only line of each test case contains two integers `n` and `x`
(`1 ≤ x ≤ n ≤ 10^18`).

## Output

For each test case, output one integer — the number of suitable segments modulo
`998244353`.

## Samples

### Sample 1

Input:

```
5
5 5
8 1
15 8
10 10
5989566119 1996588700
```

Output:

```
1
2
10
0
99996
```

## Solution

Let

```text
pref(k) = 1 ^ 2 ^ ... ^ k
```

and define `pref(0) = 0`. The xor of segment `[l, r]` is
`pref(r) ^ pref(l - 1)`, so it is zero iff:

```text
pref(r) = pref(l - 1)
```

For every valid segment containing `x`, the index `l - 1` lies in `[0, x - 1]`
and `r` lies in `[x, n]`. Therefore we only need to count equal prefix-xor
values across these two ranges.

The prefix xor pattern is:

```text
k mod 4 = 0 -> pref(k) = k
k mod 4 = 1 -> pref(k) = 1
k mod 4 = 2 -> pref(k) = k + 1
k mod 4 = 3 -> pref(k) = 0
```

Values produced by `k mod 4 = 0` and `k mod 4 = 2` identify the index `k`
uniquely. Since `[0, x - 1]` and `[x, n]` do not overlap, those values cannot
match between the two sides.

Only two values can contribute:

- `pref(k) = 0`, from `k = 0` or `k mod 4 = 3`
- `pref(k) = 1`, from `k mod 4 = 1`

So the answer is:

```text
count0(0, x - 1) * count0(x, n) +
count1(0, x - 1) * count1(x, n)
```

where counts are taken modulo `998244353`.

The implementation counts numbers in an interval with a fixed remainder modulo
`4`, plus the special index `0` for prefix value `0`.

Complexity: `O(1)` per test case.
