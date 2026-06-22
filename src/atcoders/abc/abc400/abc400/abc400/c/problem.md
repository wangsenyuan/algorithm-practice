# C - 2^a b^2 (ABC400)

**Contest:** [ABC400](https://atcoder.jp/contests/abc400) — AtCoder Beginner Contest 400  
**Task:** [https://atcoder.jp/contests/abc400/tasks/abc400_c](https://atcoder.jp/contests/abc400/tasks/abc400_c)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 350 points

## Problem Statement

A positive integer X is called a good integer if and only if it satisfies the following condition:

- There exists a pair of positive integers (a, b) such that X = 2^a × b^2.

For example, 400 is a good integer because 400 = 2^2 × 10^2.

Given a positive integer N, find the number of good integers between 1 and N, inclusive.

## Constraints

- 1 ≤ N ≤ 10^18
- N is an integer.

## Input

The input is given from Standard Input in the following format:

```text
N
```

## Output

Print the number of good integers between 1 and N, inclusive.

## Sample Input 1

```text
20
```

## Sample Output 1

```text
5
```

There are five good integers between 1 and 20: 2, 4, 8, 16, and 18. Thus, print 5.

## Sample Input 2

```text
400
```

## Sample Output 2

```text
24
```

## Sample Input 3

```text
1234567890
```

## Sample Output 3

```text
42413
```

Note that the input might not fit in a 32-bit integer type.

## Solution

The exponent `a` is positive. Split it according to its parity.

- If `a` is odd, write `a = 2k + 1`. Then
  `X = 2^(2k+1) * b^2 = 2 * (2^k * b)^2`, so `X` has the form `2y^2`.
- If `a` is even, write `a = 2k` with `k >= 1`. Then
  `X = 2^(2k) * b^2 = 4 * (2^(k-1) * b)^2`, so `X` has the form `4y^2`.

Conversely, every number of the form `2y^2` or `4y^2` is good. These two sets are
disjoint: if `2x^2 = 4y^2`, then `x^2 = 2y^2`, which has no positive integer
solution because the exponent of `2` has different parity on the two sides.

Therefore, the answer is

```text
floor(sqrt(N / 2)) + floor(sqrt(N / 4)).
```

The implementation finds both square roots with binary search. For a divisor `d`,
it searches for the first integer `x` such that `x^2 > N / d`; subtracting one gives
`floor(sqrt(N / d))`.

### Correctness

Every good integer belongs to exactly one of the two disjoint forms `2y^2` and
`4y^2`, according to whether `a` is odd or even. There are
`floor(sqrt(N / 2))` values of the first form not exceeding `N`, and
`floor(sqrt(N / 4))` values of the second form. Their sum therefore counts every
good integer in `[1, N]` exactly once.

### Complexity

Each binary search takes `O(log N)` time. The total time complexity is `O(log N)`,
and the space complexity is `O(1)`.
