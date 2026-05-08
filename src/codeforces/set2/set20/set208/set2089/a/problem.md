# A. Simple Permutation

Source: <https://codeforces.com/problemset/problem/2089/A>

Given an integer `n`, construct a permutation:

```text
p_1, p_2, ..., p_n
```

of length `n`.

For each `1 <= i <= n`, define:

```text
c_i = ceil((p_1 + p_2 + ... + p_i) / i)
```

Among:

```text
c_1, c_2, ..., c_n
```

there must be at least:

```text
floor(n / 3) - 1
```

prime numbers.

It is guaranteed that such a permutation always exists.

## Input

The first line contains an integer `t`:

```text
1 <= t <= 10
```

Each test case contains one integer:

```text
n
```

Constraints:

```text
2 <= n <= 100000
```

## Output

For each test case, output a permutation:

```text
p_1 p_2 ... p_n
```

of length `n` satisfying the condition.

If there are multiple valid answers, output any of them.

## Example

### Input

```text
3
2
3
5
```

### Output

```text
2 1
2 1 3
2 1 3 4 5
```

## Note

For `n = 2`:

```text
c_1 = ceil(2 / 1) = 2
c_2 = ceil((2 + 1) / 2) = 2
```

Both values are prime.

For `n = 5`, the sample permutation gives:

```text
c_1 = 2
c_2 = 2
c_3 = 2
c_4 = 3
c_5 = 3
```

All of them are prime.

## Solution

We need many prefixes whose ceiling average is prime. A direct way is to choose one prime number `q`, then make a long prefix where the average always stays close enough to `q` so that:

```text
ceil(prefix_sum / prefix_len) = q
```

The code picks `q` as the first prime in this range:

```text
floor(n / 3) <= q <= ceil(2n / 3)
```

Then it places numbers around `q` in this order:

```text
q, q - 1, q + 1, q - 2, q + 2, ...
```

while all numbers are still inside `[1, n]`. After that, it appends every unused number in increasing order.

For the centered part, after taking symmetric pairs around `q`:

```text
(q - d) + (q + d) = 2q
```

So the prefix average keeps returning to `q`. For odd prefix lengths, the prefix is:

```text
q, q - 1, q + 1, ..., q - d, q + d
```

Its sum is exactly:

```text
(2d + 1) * q
```

and therefore the average is exactly `q`.

For even prefix lengths, the prefix has one extra smaller value `q - d` before its matching `q + d` appears. The sum is:

```text
2d * q - d
```

Its average is:

```text
q - 1/2
```

so the ceiling is still `q`.

Therefore every prefix inside this centered block contributes a prime `c_i = q`.

Because `q` is chosen between about `n / 3` and `2n / 3`, there are enough valid numbers on both sides of `q` to create at least:

```text
floor(n / 3) - 1
```

such prefixes. The rest of the permutation can be filled with any unused numbers because the requirement has already been satisfied.

Time complexity:

```text
O(n + sqrt(n))
```

The `sqrt(n)` factor comes from checking primality while searching for `q`, and the permutation construction is linear.
