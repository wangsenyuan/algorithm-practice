# D1. Little String (Easy Version)

[Problem link](https://codeforces.com/problemset/problem/2189/D1)

**Contest:** [Codeforces contest 2189](https://codeforces.com/contest/2189)

## Problem

This is the easy version. Here the string `s` contains no `?`.

For a binary string `w = w_1...w_n`, define `f(w)` as the number of
permutations `p` of `[0, 1, ..., n-1]` such that for every `i` in `1..n`:

- if `w_i = 1`, then some subarray of `p` has MEX equal to `i`;
- if `w_i = 0`, then no subarray of `p` has MEX equal to `i`.

Given `s` (characters `0`/`1` only) and a positive integer `c`, consider all
strings `w` obtained by replacing `?` in `s` with `0`/`1` (in this version there
are no `?`, so `w = s`). Find the smallest `f(w)` that is **not** divisible by
`c`, or determine that none exists. Output the answer modulo `10^9+7`, or `-1`.

## Constraints

- `1 <= t <= 10^4`
- `3 <= n <= 2 * 10^5`
- `1 <= c <= 10^9`
- `s` has length `n` and consists of `0` and `1`
- Sum of `n` over all test cases `<= 2 * 10^5`

## Input

```text
t
case_1
...
case_t
```

Each test case:

```text
n c
s
```

## Output

For each test case, print one integer.

## Sample 1

```text
Input
3 3
001

Output
-1
```

## Sample 2

```text
Input
3 1
111

Output
-1
```

## Sample 3

```text
Input
4 100
1001

Output
4
```

## Sample 4

```text
Input
6 100
111001

Output
96
```

## Sample 5

```text
Input
6 100
111101

Output
64
```

## Sample 6

```text
Input
5 8
10001

Output
12
```

## Sample 7

```text
Input
4 100
1110

Output
-1
```

## Sample 8

```text
Input
21 123456789
111000111000111000111

Output
336892528
```

## Sample 9

```text
Input
3 4
101

Output
2
```

## Solution

For a permutation `p`, let `pos(x)` be the position of value `x`. For every
`k < n`, define

```text
Lk = min(pos(0), pos(1), ..., pos(k - 1))
Rk = max(pos(0), pos(1), ..., pos(k - 1))
```

Every subarray with MEX `k` must contain all values from `0` through `k - 1`,
so it must contain the interval `[Lk, Rk]`. Therefore:

- if `pos(k)` is inside `[Lk, Rk]`, every such subarray also contains `k`, so
  MEX `k` is impossible;
- if `pos(k)` is outside `[Lk, Rk]`, the interval `[Lk, Rk]` itself contains
  every smaller value and excludes `k`, so its MEX is `k`.

Consequently, MEX `k` exists exactly when `k` is outside the interval occupied
by all smaller values.

We can now construct a permutation by inserting the values in increasing order.
Before inserting `k`, the relative order contains `0, 1, ..., k - 1`. There are
two outer gaps and `k - 1` internal gaps:

- if `s[k - 1] == '1'`, insert `k` into either outer gap: `2` choices;
- if `s[k - 1] == '0'`, insert `k` into an internal gap: `k - 1` choices.

Every final permutation has exactly one such insertion history, so

```text
f(s) = product over k = 1..n-1 of:
       2,     if s[k - 1] == '1'
       k - 1, if s[k - 1] == '0'
```

The first bit must be `1`, because the subarray containing only `0` always has
MEX `1`. The last bit must also be `1`, because the entire permutation always
has MEX `n`. The last bit adds no factor: there is no value `n` to insert.

### Checking the `c` factor

The exact product can be far too large, and checking the product modulo
`1e9+7` cannot tell us whether it is divisible by `c`. Instead, start with

```text
remainingC = c
```

For every multiplication factor `x`, update

```text
remainingC /= gcd(remainingC, x)
```

This removes from `remainingC` exactly the prime factors supplied by this
occurrence of `x`. It is important to use one `gcd`: repeatedly dividing by
`x` could remove more copies of its prime factors than the product actually
contains.

More formally, for every prime `q`, after processing a partial product `P`, the
exponent of `q` in `remainingC` is

```text
max(0, exponent of q in c - exponent of q in P).
```

Dividing by `gcd(remainingC, x)` preserves this invariant after multiplying
`P` by `x`. Thus, after all factors:

- `remainingC == 1` exactly when `c` divides `f(s)`, so the answer is `-1`;
- otherwise `c` does not divide `f(s)`, so output `f(s) mod 1e9+7`.

For example, with `s = 10001` and `c = 8`, the factors are `2, 1, 2, 3`:

```text
remainingC: 8 -> 4 -> 4 -> 2 -> 2
```

It does not become `1`, correctly showing that `8` does not divide
`2 * 1 * 2 * 3 = 12`.

The algorithm takes `O(n log c)` time because of the GCD operations and uses
`O(1)` auxiliary space.
