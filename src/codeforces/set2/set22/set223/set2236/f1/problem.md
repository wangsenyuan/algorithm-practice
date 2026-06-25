# F1. Elections in Saransk (easy version)

[Problem link](https://codeforces.com/problemset/problem/2236/F1)

**Contest:** [Codeforces Round 1103 (Div. 3)](https://codeforces.com/contest/2236)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

This is the **easy version** of the problem. The only difference is that `x = 1`.

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

In this version, `x = 1`, so the condition becomes `lcm(p_1, ..., p_n) = p_1 · ... · p_n`, i.e. all
`p_i` are pairwise coprime.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases.

Each test case contains:

- one line with two integers `n` and `x` (`1 <= n <= 10^5`, `x = 1`);
- one line with `n` integers `a_1, a_2, ..., a_n` (`1 <= a_i <= 5 * 10^5`).

It is guaranteed that the sum of `n` over all test cases does not exceed `10^5`.

## Output

For each test case, print one integer — the number of ideal voting arrays modulo `10^9 + 7`.

## Example

### Input

```text
4
4 1
2 3 1 4
2 1
2 4
6 1
3 9 1 6 4 5
7 1
1 2 3 67 13 8 8
```

### Output

```text
8
4
40
64
```

## Solution

Because this is the easy version, `x = 1`, so the required condition is

```text
lcm(p_1, p_2, ..., p_n) = p_1 * p_2 * ... * p_n
```

Look at one prime `q`. In the product on the right, the exponent of `q` is the sum of its exponents
inside all chosen values `p_i`. In the LCM on the left, the exponent of `q` is the maximum of those
exponents.

These two values are equal if and only if `q` appears in at most one chosen `p_i`. Therefore, the whole
condition is equivalent to saying that the chosen numbers `p_1, ..., p_n` are pairwise coprime.

Now count choices prime by prime.

Suppose prime `q` appears in `a_i` with exponent `e_i`. Since `p_i` must divide `a_i`, voter `i` may
contribute any exponent from `0` to `e_i` for this prime. But `q` can be used by at most one voter, so
the valid choices for `q` are:

1. nobody uses `q`;
2. choose exactly one occurrence slot among all prime powers of `q` in all `a_i`.

If the total exponent of `q` across the entire array is

```text
c_q = e_1 + e_2 + ... + e_n
```

then this prime contributes `c_q + 1` independent choices.

Different primes do not interact, so the final answer is

```text
prod over all primes q of (c_q + 1) mod 1e9+7
```

### Algorithm

1. Precompute the least prime factor `lpf[v]` for every `v <= 5 * 10^5` with a linear sieve.
2. For every number `a_i`, repeatedly divide it by `lpf[a_i]` and increment that prime's total exponent.
3. Multiply `(count + 1)` for every prime that appeared.

### Correctness

For any fixed prime `q`, the equality between the exponent in the LCM and the exponent in the product
holds exactly when `q` is present in at most one selected divisor `p_i`. If two different selected
divisors both contain `q`, the product exponent is larger than the LCM exponent, so the array is invalid.
If at most one selected divisor contains `q`, the sum and maximum exponents are the same, so `q` causes no
violation.

For prime `q`, choosing a valid contribution means either choosing exponent `0` for every voter, or
choosing one voter `i` and one positive exponent from `1` to `e_i`. This gives exactly
`1 + sum e_i = c_q + 1` choices.

Since divisibility by different primes is independent in prime factorization, multiplying these counts
over all primes counts every valid array once and only once. Therefore the algorithm returns the number of
ideal voting arrays.

### Complexity

The sieve costs `O(A)` time and memory, where `A = 5 * 10^5`.

For each test case, factorization takes `O(total number of prime factors with multiplicity)` over the
array, and the final multiplication is linear in the number of distinct primes that appeared.
