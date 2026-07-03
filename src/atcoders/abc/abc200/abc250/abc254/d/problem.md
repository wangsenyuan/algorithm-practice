# D - Together Square

[Problem link](https://atcoder.jp/contests/abc254/tasks/abc254_d)

**Contest:** [AtCoder Beginner Contest 254](https://atcoder.jp/contests/abc254)

time limit: 2 sec

memory limit: 1024 MiB

score: 400 points

You are given an integer `N`. Count the number of pairs of positive integers `(i, j)` with `i <= N`
and `j <= N` such that `i * j` is a perfect square.

## Constraints

- `1 <= N <= 2 * 10^5`
- `N` is an integer

## Input

```text
N
```

## Output

Print the answer.

## Sample Input 1

```text
4
```

## Sample Output 1

```text
6
```

The six pairs are `(1,1), (1,4), (2,2), (3,3), (4,1), (4,4)`.

For example, `(2,3)` does not count because `2 * 3 = 6` is not a square.

## Sample Input 2

```text
254
```

## Sample Output 2

```text
896
```

## Solution

Write the prime factorization of a positive integer `x` as:

```text
x = p1^e1 * p2^e2 * ... * pk^ek
```

An integer is a perfect square iff every exponent in its prime factorization is even. Therefore
`i * j` is a square iff, for every prime `p`,

```text
exponent of p in i + exponent of p in j is even
```

Only the parity of each exponent matters. For example:

```text
12 = 2^2 * 3^1  -> odd-exponent part is 3
27 = 3^3        -> odd-exponent part is 3
12 * 27 = 324 = 18^2
```

The even powers do not matter, because they already form a square. What remains is the product of
the primes whose exponents are odd. This product is square-free; call it the square-free key of the
number.

Examples:

```text
1  -> key 1
2  -> key 2
3  -> key 3
4  = 2^2       -> key 1
8  = 2^3       -> key 2
12 = 2^2 * 3   -> key 3
18 = 2 * 3^2   -> key 2
```

Now consider two numbers:

```text
i = key(i) * a^2
j = key(j) * b^2
```

Then:

```text
i * j = key(i) * key(j) * (a*b)^2
```

The square part `(a*b)^2` is harmless. The product is a square exactly when `key(i) * key(j)` is a
square. Since each key is square-free, this happens iff:

```text
key(i) = key(j)
```

So the problem reduces to grouping numbers `1..N` by their square-free key. If one key appears `c`
times, it contributes `c * c` ordered pairs `(i, j)`, because both `i` and `j` can be any number in
that group.

### Computing the key

For each `num` from `1` to `N`, factor it by trial division:

1. Let `x = num` and `key = 1`.
2. For every divisor `p`, count how many times `p` divides `x`.
3. If the count is odd, multiply `key` by `p`.
4. After the loop, if `x > 1`, the remaining factor is prime and appears once, so multiply `key` by
   `x`.

Store `cnt[key]++`.

Finally:

```text
answer = sum over all keys cnt[key]^2
```

### Correctness

For any prime `p`, the exponent of `p` in `i*j` is the sum of its exponents in `i` and `j`. This sum
is even exactly when the two exponents have the same parity. Thus `i*j` is a square iff every prime
has matching exponent parity in `i` and `j`.

The square-free key records exactly the primes with odd exponent. Therefore two numbers have
matching exponent parity for every prime iff their keys are equal. Every ordered pair selected from
the same key group has square product, and every ordered pair with square product must come from one
such group. Summing `cnt[key]^2` over all groups counts exactly all valid ordered pairs.

### Complexity

The implementation factors each number by trial division up to the square root of the remaining
value. This is fast enough for `N <= 2 * 10^5` in Go. The memory usage is `O(N)` in the worst case
for the key counts.
