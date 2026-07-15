# D. Mahmoud and Ehab and another array construction task

[Problem link](https://codeforces.com/problemset/problem/959/D)

**Contest:** Codeforces Round 473 (Div. 2)

Time limit: 3 seconds

Memory limit: 256 megabytes

## Problem Statement

You are given an integer array `a` of length `n`. Construct another array `b`
of the same length satisfying all of the following conditions:

- `b` is lexicographically greater than or equal to `a`;
- every `b_i` is at least `2`;
- `b` is pairwise coprime: `gcd(b_i, b_j) = 1` for every `i < j`.

Among all valid arrays, output the lexicographically smallest one.

An array `x` is lexicographically greater than an array `y` if, at the first
position where they differ, the value in `x` is greater than the value in `y`.
Equal arrays are also allowed.

## Constraints

- `1 <= n <= 100000`
- `2 <= a_i <= 100000`

## Input

```text
n
a_1 a_2 ... a_n
```

## Output

Print `n` space-separated integers forming the required array `b`.

## Samples

### Sample 1

Input:

```text
5
2 3 5 4 13
```

Output:

```text
2 3 5 7 11
```

### Sample 2

Input:

```text
3
10 3 7
```

Output:

```text
10 3 7
```

The second sample is already pairwise coprime, so it can be printed unchanged.

## Solution

Two positive integers are coprime exactly when they share no prime factor.
Therefore, while constructing `b`, maintain which prime factors have already
appeared.

The implementation uses:

```text
vis[p] = true
```

when prime `p` divides one of the values already placed in `b`.

It first builds `lpf[x]`, the least prime factor of every `x < 2000000`, with
a linear sieve. This allows each candidate to be factorized quickly.

### Keep the longest possible equal prefix

Process the array from left to right. Before position `i`, maintain this
invariant:

- `b[0 ... i-1]` equals `a[0 ... i-1]`;
- those values are pairwise coprime;
- `vis` contains exactly the prime factors used by this prefix.

The helper `check(v)` factorizes `v`. It returns `true` if none of its prime
factors is already marked in `vis`.

If `check(a[i])` is true, set:

```text
b[i] = a[i].
```

Keeping the same value is always lexicographically better than increasing this
position, so it is optimal to preserve the equal prefix for as long as
possible. All prime factors of `a[i]` are then marked as used.

### The first changed position

Suppose `a[i]` shares a prime factor with the prefix. We cannot keep it.
Because every previous position is still equal to `a`, position `i` is the
first place where `b` differs from `a`. To ensure `b >= a`, we must choose:

```text
b[i] > a[i].
```

The code checks `a[i]+1`, `a[i]+2`, and so on, and selects the first number
whose prime factors are all unused. This number does not have to be prime; it
only has to be coprime with the constructed prefix.

Choosing the first valid number makes this first changed position as small as
possible, which is necessary for lexicographic minimality.

After setting this value, `b` is already lexicographically greater than `a`.
The remaining positions no longer have any lower bounds coming from the
corresponding `a[j]` values.

### Fill the suffix with unused primes

For every later position, we only need the smallest value at least `2` that is
coprime with all earlier choices.

Let `p` be the smallest unused prime. It is valid. Moreover, no smaller valid
integer exists: if an integer `x < p` were valid, every prime factor of `x`
would also be unused, giving an unused prime smaller than `p`, a contradiction.

Therefore the suffix is filled with the unused primes in increasing order.
After using a prime, it is marked and removed from the front of `primes`.

### Why the sieve bound is sufficient

Before the first change, every retained value is at most `100000`, so all used
prime factors are also at most `100000`. The replacement at the first changed
position starts from at most `100001`; an unused prime provides a valid upper
bound for this search.

Afterward, at most `N-1` additional primes are needed. There are `148933`
primes below `2000000`, more than enough for the prefix factors and at most
`100000` suffix positions. Thus every value accessed by `lpf` and every prime
needed by the construction lies below the constant `X = 2000000`.

### Correctness Proof

Before the first change, the algorithm keeps `b[i] = a[i]` whenever possible.
This preserves the smallest possible prefix, and marking all prime factors
ensures the chosen prefix remains pairwise coprime.

At the first incompatible position, equality is impossible. The algorithm
tries every larger integer in increasing order and chooses the first one that
shares no used prime factor. Hence it chooses the smallest possible value at
the first position where `b` must exceed `a`.

Once that position is increased, the suffix cannot affect the comparison with
`a`. At each suffix position, the smallest valid integer is exactly the
smallest unused prime, so choosing unused primes in increasing order makes the
suffix lexicographically minimal. Since no prime factor is ever reused, all
elements of the resulting array are pairwise coprime.

Therefore the constructed array is valid, is lexicographically at least `a`,
and is the lexicographically smallest array with these properties.

### Complexity

Let `X = 2000000`.

- The linear sieve takes `O(X)` time and `O(X)` space.
- Factoring the input and tested candidates takes `O((N+X) log X)` time in a
  conservative worst-case bound.
- The suffix scan over primes is linear in the number of skipped and selected
  primes.

Overall, the implementation fits within `O((N+X) log X)` time and `O(N+X)`
space.
