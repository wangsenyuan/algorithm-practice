# D. Storming Arasaka

[Problem link](https://codeforces.com/problemset/problem/2238/D)

**Contest:** [Codeforces Round 1106 (Div. 2)](https://codeforces.com/contest/2238)

## Problem

There is a secret number `n`. Consider all of its positive divisors except `1`
(the divisor equal to `n` is included), and partition them into several
nonempty layers `L1, L2, ..., Lk`. A partition is called **good** if:

1. For any divisor `x` from layer `Li`, all of its divisors except `1` and `x`
   lie only in the layers `L1, L2, ..., L(i-1)`.
2. In each layer, all numbers can be ordered into a chain so that any two
   neighboring numbers in this chain have GCD greater than `1`.

The length of the password is the number of layers `k`. For each candidate
value of `n`, find the minimum possible number of layers.

## Constraints

- `1 <= t <= 10^4`
- `2 <= n <= 10^6`

## Input

The first line contains `t` — the number of test cases.

Each test case consists of a single line containing one integer `n`.

## Output

For each test case, output a single integer — the minimum number of layers.

## Sample

```text
Input
8
2
4
8
16
32
67
120
33

Output
1
2
3
4
5
1
7
3
```

### Note

In the first 5 test cases, the given number has the form `2^k`. The answer for
them is `k`. The positive divisors except `1` are `2^1, 2^2, ..., 2^k`. No two
of them can lie in the same layer, so they must occupy different layers. One
valid arrangement is `Li = {2^i}`, which uses exactly `k` layers.

## Summary

Write the prime factorization as

```text
n = p1^e1 * p2^e2 * ... * pc^ec
```

Let `sum = e1 + e2 + ... + ec` and let `cnt = c`, the number of distinct prime
factors. The answer is:

```text
sum + cnt - 1
```

Why this many layers are necessary:

- For every prime `p_i`, the chain `p_i, p_i^2, ..., p_i^ei` must appear in
  strictly increasing layers, because every proper divisor of a number must be
  in an earlier layer.
- The distinct primes `p1, p2, ..., pc` cannot share one layer, since their
  pairwise GCD is `1`, so they cannot be ordered as neighboring elements in a
  valid same-layer chain.

Together these constraints force `sum` layers for the powers and `cnt - 1`
extra separations between the first prime layers, so at least
`sum + cnt - 1` layers are needed.

This lower bound is also achievable. Put the prime factors in an order, for
example `p1, p2, ..., pc`. The first `cnt` layers introduce these primes one by
one. After that, each next layer increases the total number of prime factors in
the divisor by one. Divisors in the same later layer can be connected into a
chain by GCD because they share at least one prime with a neighboring divisor.

So the problem reduces to factorizing `n`. The implementation precomputes the
least prime factor for every number up to `10^6`. For each test case it divides
`n` by each distinct prime factor:

- `sum` counts all prime factors with multiplicity.
- `cnt` counts how many distinct prime factors appear.

Then it returns `sum + cnt - 1`.

The sieve costs `O(N)` preprocessing. Each test case is factorized in
`O(number of prime factors of n)` time, with `O(N)` memory for the least-prime
factor table.
