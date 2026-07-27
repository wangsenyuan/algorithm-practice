# D2. Little String (Hard Version)

[Problem link](https://codeforces.com/problemset/problem/2189/D2)

**Contest:** [Codeforces contest 2189](https://codeforces.com/contest/2189)

## Problem

This is the hard version. Here the string `s` may contain `?`.

For a binary string `w = w_1...w_n`, define `f(w)` as the number of
permutations `p` of `[0, 1, ..., n-1]` such that for every `i` in `1..n`:

- if `w_i = 1`, then some subarray of `p` has MEX equal to `i`;
- if `w_i = 0`, then no subarray of `p` has MEX equal to `i`.

Given `s` (characters `0`, `1`, and/or `?`) and a positive integer `c`,
consider all strings `w` obtained by replacing each `?` with `0` or `1`. Find
the smallest `f(w)` that is **not** divisible by `c`, or determine that none
exists. Output the answer modulo `10^9+7`, or `-1`.

## Constraints

- `1 <= t <= 10^4`
- `3 <= n <= 2 * 10^5`
- `1 <= c <= 10^9`
- `s` has length `n` and consists of `0`, `1`, and `?`
- Sum of `n` over all test cases `<= 2 * 10^5`

## Input

```text
t
case_1
...
case_t
```

Each test case:

```text
n c
s
```

## Output

For each test case, print one integer.

## Sample 1

```text
Input
3 3
00?

Output
-1
```

## Sample 2

```text
Input
3 1
???

Output
-1
```

## Sample 3

```text
Input
4 100
1001

Output
4
```

## Sample 4

```text
Input
3 3
???

Output
2
```

## Sample 5

```text
Input
6 100
111001

Output
96
```

## Sample 6

```text
Input
6 100
111101

Output
64
```

## Sample 7

```text
Input
5 8
100?1

Output
12
```

## Sample 8

```text
Input
4 100
1??0

Output
-1
```

## Sample 9

```text
Input
20 253034496
10001100011000??????

Output
833286105
```

## Sample 10

```text
Input
3 4
1?1

Output
2
```

## Solution

### Counting permutations for a fixed binary string

For every `k < n`, consider the smallest interval of the permutation containing
all values `0, 1, ..., k - 1`.

- MEX `k` exists if and only if value `k` is outside this interval.
- If we insert values in increasing order, placing `k` outside the current
  sequence gives `2` choices.
- Placing `k` inside the sequence gives `k - 1` choices.

Therefore, using one-based string positions,

```text
factor(k) = 2,     if w[k] = 1
            k - 1, if w[k] = 0

f(w) = product of factor(k), for k = 1..n-1
```

The first bit must be `1`, because the single value `0` gives a subarray with
MEX `1`. The last bit must also be `1`, because the entire permutation has MEX
`n`. The last bit contributes no factor because value `n` is not in the
permutation.

### First build the minimum product

For a `?` at position `k`, compare its two possible factors:

| Position | Choose `0` | Choose `1` | Minimum choice |
|---|---:|---:|---|
| `k = 1` | invalid | `2` | force `1` |
| `k = 2` | `1` | `2` | choose `0` |
| `k = 3` | `2` | `2` | either |
| `k >= 4` | `k - 1` | `2` | choose `1` |

Choosing the smaller factor independently at every `?` produces the smallest
possible product; call it `B`.

While constructing `B`, its value modulo `1e9+7` is maintained for the answer.
To test divisibility by `c`, maintain:

```text
remainingC = c
remainingC /= gcd(remainingC, factor)
```

After all factors, `remainingC == 1` exactly when `c` divides the complete
product. This works because each GCD removes exactly the prime factors of `c`
supplied by the current factor.

If `remainingC != 1`, then `B` is already the smallest product not divisible by
`c`, so it is the answer.

### When the minimum product is divisible by `c`

Suppose `c` divides `B`. We must change some `?` away from its minimum choice.

For `k >= 4`, the minimum choice is `1`, whose factor is `2`. Changing it to
`0` replaces that factor with `k - 1`.

- If `k` is odd, then `k - 1` is even. Write `k - 1 = 2q`. The new product is
  `B / 2 * (k - 1) = B * q`. It remains divisible by `c`, so this change cannot
  help.
- If `k` is even, then `k - 1` is odd. The replacement removes exactly one
  factor of `2` from the product and adds only odd factors. This is the only
  kind of change that can make the product stop being divisible by `c`.

Let `v2(x)` denote the exponent of `2` in `x`. Because `c` divides `B`,
`v2(B) >= v2(c)`. If we make `r` useful changes, the new power of `2` is
`v2(B) - r`. We need

```text
v2(B) - r < v2(c)
```

so the minimum required number of changes is

```text
r = v2(B) - v2(c) + 1
```

If there are fewer than `r` even positions `k >= 4` with `s[k] == '?'`, no
assignment can work and the answer is `-1`.

Otherwise, choose the first `r` eligible positions. Replacing position `k`
changes factor `2` into `k - 1`. For a fixed number of replacements, choosing
the smallest positions minimizes the product.

### Example

Consider:

```text
n = 5
c = 8
s = 1???1
```

The minimum choices give factors:

```text
2, 1, 2, 2
```

Thus `B = 8`, which is divisible by `8`. Here:

```text
v2(B) = 3
v2(c) = 3
r = 3 - 3 + 1 = 1
```

The first eligible position is `k = 4`. Change it from `1` to `0`, replacing
factor `2` with factor `3`:

```text
2 * 1 * 2 * 3 = 12
```

Now `8` does not divide `12`, and `12` is the minimum valid answer.

### Complexity

The implementation uses two linear passes. GCD operations take `O(log c)`
time, so the total complexity is `O(n log c)` with `O(1)` auxiliary space.
