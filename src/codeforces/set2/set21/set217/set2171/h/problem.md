# H. Shiori Miyagi and Maximum Array Score (Codeforces 2171H)

**Limits:** 4 seconds per test, 256 MB  
**I/O:** standard input / standard output

Source: [https://codeforces.com/problemset/problem/2171/H](https://codeforces.com/problemset/problem/2171/H)

## Problem Statement

For arbitrary integers `b >= 2` and `x >= 1`, define `v(b, x)` to be the maximum integer `k` such that:

```text
b^k divides x
```

Equivalently, `v(b, x)` is the largest `k` such that `x` is a multiple of `b^k`. This value is always a well-defined non-negative integer.

You are given integers `n` and `m` satisfying `n <= m`.

Find the maximum value of:

```text
sum_{i=2}^n v(i, a_i)
```

over all arrays `a` of length `n` satisfying:

- `a` is strictly increasing, meaning `a_i < a_{i+1}` for every `1 <= i < n`;
- `1 <= a_i <= m` for every `1 <= i <= n`.


## Solution

We need choose a strictly increasing array:

```text
1 <= a_1 < a_2 < ... < a_n <= m
```

and maximize:

```text
sum_{i=2}^n v(i, a_i)
```

Only `a_2, ..., a_n` affect the score, but `a_1` is still useful for the strict increasing condition.

### Feasible Range for `a_i`

For a fixed position `i`, the smallest possible value is:

```text
a_i >= i
```

because there must be at least `i - 1` positive distinct smaller values before it.

The largest possible value is:

```text
a_i <= m - n + i
```

because after `a_i`, we still need to place `n - i` larger values before reaching `m`.

So:

```text
i <= a_i <= m - n + i
```

### Remove Strictness by Shifting

Define:

```text
b_i = a_i - i
```

Then:

```text
0 <= b_i <= m - n
```

Also, since `a_i < a_{i+1}` and all values are integers:

```text
a_{i+1} >= a_i + 1
```

therefore:

```text
a_{i+1} - (i + 1) >= a_i - i
```

so:

```text
b_1 <= b_2 <= ... <= b_n
```

The strict increasing array `a` is now equivalent to a nondecreasing sequence of offsets `b`.

This is the key transformation used by the implementation.

### DP Definition

Let `q = a_i - i`.

We process positions `i = 2, 3, ..., n`.  Maintain:

```text
best[q] = maximum score after processing some prefix,
          with the last chosen offset <= q
```

Because offsets must be nondecreasing, if we choose `a_i = j`, where:

```text
q = j - i
```

then the previous offset must be at most `q`.  So the transition is:

```text
new score = max previous score with offset <= q + v(i, j)
```

The code stores prefix maxima in a Fenwick tree:

```go
bit.query(q)
```

returns the best score among offsets `<= q`, and:

```go
bit.update(q, value)
```

updates the best score ending at offset `q`.

### Only Multiples Matter

For fixed `i`, the value:

```text
v(i, j)
```

is positive only when `j` is divisible by `i`.

If `j` is not divisible by `i`, then:

```text
v(i, j) = 0
```

Choosing such a value never increases the score.  It can only serve as a spacer for the nondecreasing offset sequence.

But explicit spacer states are unnecessary:

- all scores are nonnegative;
- `bit.query(q)` already allows any previous offset `<= q`;
- a zero-contribution choice would only carry an existing score forward to a larger or equal offset, which future prefix queries already handle.

Therefore we only need to perform updates for values `j` that are multiples of `i`.

For each `i`, valid `j` must satisfy:

```text
i <= j <= m - n + i
j % i = 0
```

The implementation loops over these multiples.

### Why the Loop Goes Downward

For one fixed `i`, every transition must use states from previous positions only.  It must not use an update produced by the same `i`.

The Fenwick tree answers prefix maximum queries.  If we process `j` in decreasing order, then any update already made for this `i` has a larger offset, so it is not included in:

```go
bit.query(j - i)
```

Thus the descending loop prevents using the same position twice.

This is why the code does:

```go
for j := (m-n+i)/i*i; j >= i; j -= i {
    v := bit.query(j-i) + calc(i, j)
    bit.update(j-i, v)
}
```

### Computing `v(i, j)`

The helper:

```go
calc(i, j)
```

counts how many times `j` can be divided by `i`:

```text
while j % i == 0:
    answer++
    j /= i
```

This is exactly the definition of `v(i, j)`.

### Final Answer

All offsets are in:

```text
0 ... m - n
```

After processing all positions, the answer is the best score with final offset at most `m - n`:

```go
bit.query(m - n)
```

### Complexity

For each `i`, we iterate over multiples of `i` up to `m - n + i`, so the total number of transitions is:

```text
sum_{i=2}^n O(m / i) = O(m log m)
```

Each transition uses one Fenwick query and one update:

```text
O(log m)
```

So the total complexity per test case is:

```text
O(m log^2 m)
```

with memory:

```text
O(m - n + 1)
```

The sum of `m` over all test cases is at most `2 * 10^5`, so this is fast enough.

