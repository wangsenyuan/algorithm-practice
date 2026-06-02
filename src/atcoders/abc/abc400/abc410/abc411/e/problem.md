# E - E [max]

https://atcoder.jp/contests/abc411/tasks/abc411_e

**Time Limit:** 3 sec / **Memory Limit:** 1024 MiB

**Score:** 450 points

## Problem Statement

There are `N` six-sided dice. The dice are numbered from `1` to `N`, and the
numbers written on the faces of die `i` are:

```text
A_{i,1}, A_{i,2}, ..., A_{i,6}
```

All `N` dice are rolled simultaneously. For each die, the face that comes up is
chosen independently and uniformly at random.

Find the expected value, modulo `998244353`, of the maximum value among the
numbers shown on the top faces.

## Expected Value Modulo

The expected value is a rational number. If it is written as an irreducible
fraction:

```text
P / Q
```

then under the constraints, `Q` is not divisible by `998244353`. Therefore,
there is a unique integer `R` such that:

```text
R * Q == P (mod 998244353)
0 <= R < 998244353
```

Output this `R`.

## Constraints

- `1 <= N <= 10^5`
- `1 <= A_{i,j} <= 10^9`
- All input values are integers.


## Solution

Let the rolled values be independent, and let

```text
M = max(rolled values)
```

We need `E[M]`.

The useful observation is that the maximum is easier to count through the event
`M <= x`.

For a fixed value `x`, define

```text
cnt_i(x) = number of faces on die i whose value is <= x
```

Then

```text
P(M <= x) = product(cnt_i(x)) / 6^N
```

because `M <= x` means every die independently rolls one of its faces not larger
than `x`.

Now collect all face values, sort them, and remove duplicates:

```text
s[0] < s[1] < ... < s[k-1]
```

Between two adjacent values, nothing changes. For every `t` with
`s[i] <= t < s[i+1]`, the set of faces `<= t` is exactly the set of faces
`<= s[i]`, so `P(M <= t)` is constant on that whole interval.

Using the tail-sum form,

```text
E[M] = s[k-1] - sum over t < s[k-1] of P(M <= t)
```

After grouping equal-probability intervals, this becomes

```text
E[M] = s[k-1]
       - sum_{i=0}^{k-2} (s[i+1] - s[i]) * P(M <= s[i])
```

Substituting the probability formula:

```text
E[M] = s[k-1]
       - sum_{i=0}^{k-2} (s[i+1] - s[i]) * product(cnt_j(s[i])) / 6^N
```

This is exactly what the implementation computes.

### Maintaining `product(cnt_i)`

The code processes the sorted distinct values from small to large.

For every die `j`, it maintains:

```text
b[j] = number of faces on die j already processed
     = number of faces on die j with value <= current s[i]
```

It also maintains:

```text
prod = product of non-zero b[j]
zero = number of dice with b[j] == 0
```

When processing value `s[i]`, every occurrence of that value increments the
corresponding die count. If one die has duplicate faces with the same value, the
die index appears multiple times in `upd[i]`, so `b[j]` increases once per face.
This is correct because every face is a distinct equally likely outcome.

When `b[j]` changes from `old` to `old + 1`:

- if `old == 0`, this die stops being zero, so `zero--`;
- otherwise, remove the old factor from `prod` by multiplying by `inv(old)`;
- then multiply `prod` by `old + 1`.

After all updates for `s[i]`, if `zero == 0`, every die has at least one face
`<= s[i]`, so

```text
prod / 6^N = P(M <= s[i])
```

The implementation first accumulates the numerator:

```text
ans -= prod * (s[i+1] - s[i])
```

Then, after all intervals are processed, it divides once by `6^N` and adds the
largest possible value:

```text
ans = ans / 6^N + s[k-1]
```

All divisions are done with modular inverses under `998244353`.

The complexity is `O(N log N)` from sorting the `6N` face values and building the
update buckets.
