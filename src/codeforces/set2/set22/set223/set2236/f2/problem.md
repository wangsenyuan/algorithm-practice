# F2. Elections in Saransk (hard version)

[Problem link](https://codeforces.com/problemset/problem/2236/F2)

**Contest:** [Codeforces Round 1103 (Div. 3)](https://codeforces.com/contest/2236)

time limit per test: 3 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

This is the **hard version** of the problem. The only difference is that `1 <= x <= 5 * 10^5`.

On the way home after buying his favorite soda "Zola Cero", Egor saw that elections for the position of
"Best Number" are taking place in Saransk.

There are `n` people at the polling station. Each person brought a number `a_i`. When the `i`-th
person enters the voting booth, they choose a candidate `p_i` that divides `a_i`.

After everyone has voted, we get an array of votes `[p_1, p_2, ..., p_n]`.

Egor really likes the number `x` and considers the voting **ideal** if

```text
x · lcm(p_1, p_2, ..., p_n) = p_1 · p_2 · ... · p_n
```

Help him find the number of different arrays `p` modulo `10^9 + 7` that are ideal.

Two arrays are considered different if they differ in at least one position.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases.

Each test case contains:

- one line with two integers `n` and `x` (`1 <= n <= 10^5`, `1 <= x <= 5 * 10^5`);
- one line with `n` integers `a_1, a_2, ..., a_n` (`1 <= a_i <= 5 * 10^5`).

It is guaranteed that the sum of `n` over all test cases does not exceed `10^5`.

## Output

For each test case, print one integer — the number of ideal voting arrays modulo `10^9 + 7`.

## Example

### Input

```text
5
2 2
2 4
1 557
7
7 4
2 4 8 13 11 1 6
3 1000
1 2 3
3 3
4 8 10
```

### Output

```text
2
0
3600
0
0
```

## Solution Summary

Prime factors are independent, so handle one prime `pr` at a time. If
`e_i = v_pr(p_i)` and `u = v_pr(x)`, the condition

```text
max(e_i) + u = sum(e_i)
```

must hold for this prime.

### Primes not in x

If `u = 0`, then

```text
max(e_i) = sum(e_i)
```

This is possible only when `pr` appears in at most one chosen `p_i`. If the total exponent of `pr`
among all `a_i` is `c`, then we may either not use this prime, or choose one of those `c` exponent
slots. So this prime contributes:

```text
c + 1
```

This is the same counting as in F1.

### Primes in x

For a prime `pr | x`, let `u = v_pr(x)`. We need to count assignments of exponents
`0 <= e_i <= v_pr(a_i)` such that

```text
sum(e_i) = max(e_i) + u
```

A direct DP over all people works logically:

```text
dp[current maximum exponent][current sum of exponents]
```

but doing this for every person and every prime factor of `x` is too slow on large tests. It also has
many repeated transitions, because `v_pr(a_i)` is tiny: all values are at most `5 * 10^5`, so a prime
exponent is at most `18`.

The optimized solution groups people by this exponent instead.

For a fixed prime `pr | x`, compute:

```text
cnt[e] = number of i with v_pr(a_i) = e
```

Now enumerate the final maximum exponent `mx`. First count all assignments where every `e_i <= mx`.
For a person with `v_pr(a_i) = e`, the allowed chosen exponent is from `0` to `min(e, mx)`, so its
generating polynomial is:

```text
1 + z + z^2 + ... + z^min(e,mx)
```

For all people with the same exponent `e`, this becomes:

```text
(1 + z + z^2 + ... + z^min(e,mx)) ^ cnt[e]
```

Multiplying these small polynomials gives the number of assignments by total exponent sum. The
coefficient of `z^(mx + u)` counts assignments whose sum is correct and whose maximum is at most `mx`.

To require the maximum to be exactly `mx`, subtract the count for maximum at most `mx - 1`:

```text
ways_exact_mx =
    ways_at_most_mx[mx + u] - ways_at_most_(mx-1)[mx + u]
```

Summing this over all possible `mx` gives the contribution of `pr`.

### Performance improvement

The TLE version recalculated the DP by scanning all `n` people for every prime factor of `x`, and each
transition used per-person prime-exponent map lookups. In the worst case this does a large amount of
repeated work.

The AC version performs one factorization pass over the array, records only `cnt[e]` for primes that
divide `x`, and then works with very small polynomials. The degree only needs to go up to

```text
max v_pr(a_i) + v_pr(x)
```

which is below `40` for the constraints. Polynomial exponentiation by squaring handles large `cnt[e]`
without iterating once per person.

So the expensive part changes from roughly:

```text
number of prime factors of x * n * small DP state
```

to:

```text
factorization of all a_i + number of prime factors of x * tiny polynomial work
```

This removes the large-test bottleneck.

### Algorithm

1. Precompute least prime factors up to `5 * 10^5`.
2. Factor `x` and remember its distinct prime factors and their exponents.
3. Factor every `a_i` once.
   - Add each prime exponent to the global total count.
   - If the prime divides `x`, increment `cnt[e]` for that prime.
4. For primes not dividing `x`, multiply the answer by `totalExponent + 1`.
5. For each prime dividing `x`, use the grouped polynomial counting above and multiply its contribution.

### Complexity

Let `A = 5 * 10^5`. The sieve costs `O(A)`.

Each test case factors all `a_i` once. For primes dividing `x`, the remaining work is over at most a few
prime factors, exponent values up to `18`, and polynomial degree below `40`, so it is effectively
constant per such prime.
