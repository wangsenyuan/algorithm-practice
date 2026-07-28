# C. Partial Sums

[Problem link](https://codeforces.com/problemset/problem/223/C)

**Contest:** [Codeforces Round #136 (Div. 1)](https://codeforces.com/contest/223)

time limit per test: 4 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

You have an array `a` of `n` integers (1-indexed). One operation replaces `a`
with its prefix-sum array modulo `10^9 + 7`:

```text
s_i = (a_1 + a_2 + ... + a_i) mod (10^9 + 7)
a := s
```

Find the array after exactly `k` such operations.

## Constraints

- `1 <= n <= 2000`
- `0 <= k <= 10^9`
- `0 <= a_i <= 10^9`

## Input

The first line contains two integers `n` and `k`.

The second line contains `n` integers `a_1, a_2, ..., a_n`.

## Output

Print `n` integers — the array after `k` operations, separated by spaces.

## Sample 1

```text
Input
3 1
1 2 3

Output
1 3 6
```

## Sample 2

```text
Input
5 0
3 14 15 92 6

Output
3 14 15 92 6
```

## ideas

1. One prefix-sum step is a lower-triangular transform; `k` steps compose to
   binomial weights.
2. After `k` ops, original `a[j]` contributes `C(k - 1 + d, d)` to index
   `j + d` (0-based), with `C(-1, 0) := 1` so `k = 0` is the identity.
3. Precompute `C[d]` by `C[0] = 1`, `C[d+1] = C[d] · (k + d) / (d + 1)`
   modulo `10^9+7`, then each answer entry is an `O(n)` convolution → `O(n^2)`.

## summary

### Goal

Apply the prefix-sum map `a ↦ s` with `s_i = (a_1+…+a_i) mod (10^9+7)` exactly
`k` times (`k` up to `10^9`, `n ≤ 2000`).

### Why not simulate

One step is `O(n)`, but `k` is far too large. Need a closed form for the
`k`-fold composition.

### Linear algebra view

One prefix-sum is multiplication by a lower-triangular matrix of all ones on
and below the diagonal. Its `k`-th power has binomial entries: the coefficient
from original index `j` to final index `i` (`0`-based, `j ≤ i`) is

```text
C(k - 1 + (i - j), i - j)
```

(with the convention `C(-1, 0) = 1`, so `k = 0` leaves the array unchanged).

So

```text
answer[i] = Σ_{j=0..i} a[j] · C(k - 1 + i - j, i - j)  (mod 10^9+7)
```

### Computing the binomials

Let `C[d] = C(k - 1 + d, d)`. Then

```text
C[0] = 1
C[d+1] = C[d] · (k + d) / (d + 1)   (mod 10^9+7)
```

Only `n` coefficients are needed; use modular inverses for the divisions.

### Complexity

Build `C` in `O(n log MOD)` (or `O(n)` with precomputed inverses), then each of
`n` outputs is an `O(n)` sum → `O(n^2)` total, fine for `n = 2000`.

### Tiny check

`a = [1,2,3]`, `k = 1`: `C[d] = C(d, d) = 1`, so
`answer = [1, 1+2, 1+2+3] = [1,3,6]`.
