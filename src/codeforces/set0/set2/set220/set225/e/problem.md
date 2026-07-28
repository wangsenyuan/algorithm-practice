# Codeforces 225E — Unsolvable

[Problem link](https://codeforces.com/problemset/problem/225/E)

## Problem

Consider the following equation:

![The equation from the official statement](https://espresso.codeforces.com/e132675bc04c8df0b7e312e08476a7157b8f823d.png)

Here, `[a]` denotes the integer part of `a`.

Find every positive integer `z` for which the equation has no solution in
positive integers `x` and `y`. Write these values in increasing order:

`z₁, z₂, z₃, ...`

Given `n`, output `zₙ`.

## Input

The input contains one integer `n` (`1 ≤ n ≤ 40`).

## Output

Print `zₙ` modulo `1000000007`.

## Examples

### Example 1

```text
Input
1

Output
1
```

### Example 2

```text
Input
2

Output
3
```

### Example 3

```text
Input
3

Output
15
```

## Solution Summary

Let

\[
f(x,y)=\left\lfloor\frac{x}{2}\right\rfloor+y+xy.
\]

We must find the positive values of `z` that cannot be written as `f(x, y)`
with positive integers `x` and `y`.

### 1. Odd `x`

Write `x = 2a - 1`, where `a ≥ 1`. Then

\[
z=(a-1)+2ay,
\qquad
z+1=a(2y+1).
\]

The factor `2y + 1` is odd and at least `3`. Hence an odd-`x` solution
exists exactly when `z + 1` has an odd divisor greater than `1`:

- If the factorization is `z + 1 = a \cdot q` with odd `q ≥ 3`, choose
  `y = (q - 1) / 2` and `x = 2a - 1`.
- Conversely, any odd-`x` solution supplies this factorization.

Therefore, a value with no odd-`x` solution must have

\[
z+1=2^k,
\qquad
z=2^k-1.
\]

### 2. Even `x`

Write `x = 2a`, where `a ≥ 1`. Then

\[
z=a+(2a+1)y,
\qquad
2z+1=(2a+1)(2y+1).
\]

Both factors on the right are odd and at least `3`. Thus an even-`x`
solution exists exactly when `2z + 1` is composite.

For a candidate left by the odd case, `z = 2^k - 1`, so

\[
2z+1=2^{k+1}-1.
\]

It is unsolvable precisely when this Mersenne number is prime. If `p = k + 1`
is a Mersenne-prime exponent, the corresponding answer is

\[
z=2^{p-1}-1.
\]

The answers are increasing with `p`, so the required `z_n` uses the `n`-th
Mersenne-prime exponent. The first 40 such exponents are stored in
`mersennePrimeExponents`; this is sufficient because `n ≤ 40`.

### Implementation

`solve` selects `p := mersennePrimeExponents[n-1]` and computes

\[
2^{p-1}-1 \pmod {10^9+7}
\]

with the binary exponentiation helper `pow`.

### Correctness Proof

For any positive integer `z`:

1. By the odd-`x` derivation, `z` has no odd-`x` solution if and only if
   `z = 2^k - 1` for some positive `k`.
2. For such a `z`, the even-`x` derivation shows that an even-`x` solution
   exists if and only if `2^{k+1} - 1` is composite.
3. Therefore `z` has no solution at all if and only if
   `z = 2^{p-1} - 1` and `2^p - 1` is prime.
4. `mersennePrimeExponents[n-1]` is the `n`-th increasing exponent satisfying
   that condition, and `2^{p-1} - 1` increases with `p`.

Thus `solve` returns exactly `z_n` modulo `10^9 + 7`.

### Complexity Analysis

The lookup is constant time. Binary exponentiation takes `O(log p)` time and
`O(1)` auxiliary space, where `p` is the selected exponent.
