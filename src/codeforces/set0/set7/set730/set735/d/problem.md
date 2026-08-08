# D. Taxes

[Problem link](https://codeforces.com/problemset/problem/735/D)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

## Problem

Mr. Funt has total income `n` burles. The tax for a number `x` is the maximum
proper divisor of `x`, that is, the largest divisor of `x` different from `x`.

He may split `n` into any number of parts:

```text
n_1 + n_2 + ... + n_k = n
```

where every part must be at least `2`. He pays tax for each part separately.

Find the minimum total tax he can pay.

## Constraints

- `2 <= n <= 2 * 10^9`

## Input

The input contains a single integer `n`.

## Output

Print one integer: the minimum possible tax.

## Examples

### Input 1

```text
4
```

### Output 1

```text
2
```

### Input 2

```text
27
```

### Output 2

```text
3
```

## Solution Summary

For any prime number `p`, its largest proper divisor is `1`, so using a prime
part costs exactly `1` tax. Therefore, minimizing total tax is equivalent to
using as few prime summands as possible.

The answer is always one of `1`, `2`, or `3`:

- If `n` is prime, use `n` itself, so the answer is `1`.
- Otherwise, if `n` is even, Goldbach's conjecture for the needed range lets us
  write `n` as a sum of two primes, so the answer is `2`.
- Otherwise `n` is odd and composite. It can be a sum of two primes only if one
  of them is `2`, so check whether `n - 2` is prime. If yes, the answer is `2`.
- If not, the answer is `3`: write `n = 3 + (n - 3)`, where `n - 3` is even and
  can be split into two primes.

So the implementation only needs primality checks up to `sqrt(n)`:

```text
if isPrime(n):       answer = 1
else if n is even:  answer = 2
else if isPrime(n-2): answer = 2
else:               answer = 3
```

Complexity: `O(sqrt(n))`.
