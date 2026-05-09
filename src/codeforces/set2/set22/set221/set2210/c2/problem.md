# Codeforces 2210C2 - A Simple GCD Problem (Hard Version)

https://codeforces.com/problemset/problem/2210/C2

## Statement

This is the hard version of the problem. The difference between the versions is
that in this version:

- `1 <= n <= 5 * 10^4`;
- `1 <= bi <= 10^9` for `1 <= i <= n`;
- the time limit is `6` seconds.

Note that a solution for one version does not necessarily solve the other
version. You can hack only if you solved all versions of this problem.

You are given two arrays `a` and `b` of length `n`.

For each index `i` (`1 <= i <= n`) of array `a`, you can perform the following
operation at most once:

- choose an arbitrary integer `m` (`m != ai`) such that `1 <= m <= bi`, and set
  `ai := m`.

Let the array after performing all operations be `a'`. You can only perform
operations in such a way that the following condition holds:

For all `1 <= l < r <= n`,

```text
gcd(a[l], a[l+1], ..., a[r]) = gcd(a'[l], a'[l+1], ..., a'[r])
```

Here, `gcd(x, y)` denotes the greatest common divisor of integers `x` and `y`.

Determine the maximum number of operations that can be performed while ensuring
that the condition remains satisfied.

## Input

Each test contains multiple test cases.

The first line contains the number of test cases `t` (`1 <= t <= 10^4`).

For each test case:

- The first line contains an integer `n` (`2 <= n <= 5 * 10^4`) — the length of
  `a`.
- The second line contains `n` space-separated integers
  `a1, a2, ..., an` (`1 <= ai <= 10^9`).
- The third line contains `n` space-separated integers
  `b1, b2, ..., bn` (`1 <= bi <= 10^9`).

It is guaranteed that the sum of `n` over all test cases does not exceed
`5 * 10^4`.

## Output

For each test case, output the maximum number of operations that can be done.

## Sample

```text
9
7
1 2 3 4 5 6 7
100 6 12 20 5 2 7
3
67 67 67
1000000000 1000000000 1000000000
6
8 10 10 12 12 14
1 1 1 1 1 1
8
2010 330 550 2 210 385 1001 323
2010 1000 1200 30 500 1000 2000 1000
4
1 2 4 3
1 1 1 1
6
2 3 4 5 1 6
10 15 20 25 5 30
2
1 2
2 100
5
2860 2860 143 9009 9009
2860 2860 1430 9009 9009
3
2 60 60
10 60 60
```

```text
7
3
0
7
1
6
2
0
0
```

## Note

In the first test case, since the GCD of all subarrays is `1`, we can change the
array to:

```text
a' = [67, 5, 7, 11, 1, 2, 1]
```

Here, `7` operations are performed, and since there are only `7` elements in the
array, this is the maximum possible.

In the second test case, we can change the array to:

```text
a' = [938, 2613, 4489]
```

It can be shown that after this transformation, the GCD of all subarrays remains
the same. Hence the answer is `3`.

In the third test case, no operation can be performed, so the answer is `0`.


### ideas
1. 在保证gcd(c[l..r]) = gcd(a[l...r])的前提下，使用m <= b[i]，c[i] = m
2. 先考虑两个相邻的位置, gcd(c[i], c[i+1]) = gcd(a[i], a[i+1])
3. 如果可以，是不是将c[i], c[i+1] = g，就可以了（因为这里要求m <= b[i], b[i+1], 那么使用最小的g似乎是一个更优的选择)
4. 那么对于i来说，只需要考虑i-1，i, i+1?
5. c[i] = lcm(g(a[i-1], a[i]), g(a[i], a[i+1]))?
6. 

## Solution

Let the final array be `c`.

The condition asks to preserve the GCD of every subarray with length at least
`2`. The first important simplification is that it is enough to preserve every
adjacent pair GCD.

Define:

```text
g[i] = gcd(a[i], a[i+1])     for 0 <= i+1 < n
```

For any interval `[l, r]` with `l < r`,

```text
gcd(a[l], a[l+1], ..., a[r]) = gcd(g[l], g[l+1], ..., g[r-1])
```

Why this is true: consider any prime `p` and only look at the exponent of `p` in
each number. The exponent in the interval GCD is the minimum exponent among all
positions `l..r`. The exponent in `gcd(g[l], ..., g[r-1])` is the minimum of
`min(e[i], e[i+1])` over adjacent pairs. These two minima are equal. The minimum
position is included in at least one adjacent pair because the interval length is
at least `2`.

So we only need:

```text
gcd(c[i], c[i+1]) = g[i]
```

for every adjacent pair.

## The mandatory base value

For each position `i`, `c[i]` must be divisible by the GCD it forms with the
left neighbor and also by the GCD it forms with the right neighbor.

So the smallest required divisor of `c[i]` is:

```text
base[i] = gcd(a[i], a[i+1])                         if i = 0
base[i] = gcd(a[i-1], a[i])                         if i = n-1
base[i] = lcm(gcd(a[i-1], a[i]), gcd(a[i], a[i+1])) otherwise
```

Every valid final value can be written as:

```text
c[i] = base[i] * k[i]
```

Also, the original value `a[i]` is always divisible by `base[i]`, because it is
divisible by both adjacent pair GCDs.

If `base[i] > b[i]`, then we cannot change position `i`, because every changed
value must be at most `b[i]`. In that case the only possible value is the
original `a[i]`.

Otherwise, possible changed values are `base[i] * k` where:

```text
base[i] * k <= b[i]
base[i] * k != a[i]
```

## Why only adjacent DP is needed

After choosing `base[i] * k[i]`, the only condition involving `k[i]` and
`k[i+1]` is:

```text
gcd(base[i] * k[i], base[i+1] * k[i+1]) = g[i]
```

This is local to the adjacent pair `(i, i+1)`, so the whole array can be handled
with a left-to-right DP.

Let the DP state remember which multiplier candidate was chosen for `i-1`.
When trying a multiplier candidate for `i`, check the pair GCD above. If it is
equal to `g[i-1]`, the transition is valid. Add `1` to the score if the value at
`i` is changed.

## Why a few prime multipliers are enough

The hard part is proving that we do not need to try all possible multipliers
`k`.

For a fixed position `i`, suppose there exists some valid multiplier `k > 1`.
Take any prime divisor `p` of `k`. Since:

```text
base[i] * k
```

does not create an extra GCD with either neighbor, the smaller value:

```text
base[i] * p
```

also does not create an extra GCD with either neighbor. Any bad prime factor
would already have appeared in `k` and would already have broken the pair GCD.

Also:

```text
base[i] * p <= base[i] * k <= b[i]
```

So whenever a non-trivial multiplier works, one prime multiplier also works.
If that prime multiplier happens to give the original value `a[i]`, then using
multiplier `1` gives `base[i]`, which is also valid: `base[i]` divides `a[i]`,
so replacing `a[i]` by the smaller required base cannot increase any adjacent
GCD, and it still contains the required adjacent GCD factors. Since
`base[i] != a[i]` in this case and `base[i] <= b[i]`, it is a valid operation.

Now consider how many small primes can be forbidden by fixed neighbors. For a
candidate prime `p`, it is bad only if it appears in one of the two neighbor
values after removing the required GCD part; otherwise multiplying by `p` cannot
increase the adjacent GCD.

Each relevant neighbor value is at most `10^9`. A number not exceeding `10^9`
has at most `9` distinct prime factors, because:

```text
2 * 3 * 5 * 7 * 11 * 13 * 17 * 19 * 23 = 223092870
2 * 3 * 5 * 7 * 11 * 13 * 17 * 19 * 23 * 29 > 10^9
```

So the two neighbors together can forbid at most `18` distinct primes. Among the
first `20` primes, at least one prime is not forbidden whenever we need an
available prime replacement. This is why it is enough to test:

```text
1, 2, 3, 5, 7, 11, ..., 71
```

The value `1` represents choosing the mandatory base itself.

## DP

The implementation uses:

```text
factors = [1, first 20 primes]
```

For every position `i`, first compute `base[i]`. If `base[i] > b[i]`, the code
sets `base[i] = a[i]`, which forces the unchanged original value to be the only
usable value at that position.

Then:

```text
dp[j] = maximum operations up to current position
        when current value is base[i] * factors[j]
```

For a transition from multiplier `x` at `i-1` to multiplier `y` at `i`, check:

```text
gcd(base[i-1] * x, base[i] * y) == gcd(a[i-1], a[i])
```

The candidate is allowed if it is either unchanged or within the bound:

```text
base[i] * y == a[i]  or  base[i] * y <= b[i]
```

If the candidate value differs from `a[i]`, it contributes one operation.

The answer is the maximum DP value after the last position.

## Complexity

There are only `21` multiplier candidates per position, so each adjacent
transition costs `21 * 21`.

The total complexity is:

```text
O(n * 21^2)
```

with `O(21)` DP memory.
