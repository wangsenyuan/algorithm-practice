# B. Siga ta Kymata (Codeforces 2163B)

**Limits:** 1.5 seconds per test, 256 MB  
**I/O:** standard input / standard output

Source: [https://codeforces.com/problemset/problem/2163/B](https://codeforces.com/problemset/problem/2163/B)

## Problem Statement

You are given a permutation `p` of every integer from `1` to `n`. You also own a binary string `s` of size `n`, initially with `s_i = 0` for all `1 <= i <= n`.

You may do the following operation at most `5` times:

- Choose two integers `l` and `r` such that `1 <= l <= r <= n`.
- Then, for every `i` satisfying both:
  - `l < i < r`;
  - `min(p_l, p_r) < p_i < max(p_l, p_r)`;

  set `s_i` to `1`.

You are also given a binary string `x` of size `n`.

After performing operations, it must hold for every `1 <= i <= n` that if `x_i = 1`, then `s_i = 1`. If `x_i = 0`, then `s_i` may be either `0` or `1`.

Find any sequence of at most `5` operations satisfying this condition, or report that it is impossible.

You do not need to minimize the number of operations.

## Definitions

A permutation `p` of every integer from `1` to `n` is a sequence of elements from `1` to `n` where every element appears exactly once.

A string is binary if every character is `0` or `1`.


## Solution

Let:

```text
pos1 = position of value 1
posN = position of value n
```

There are four positions that can never be set to `1`:

1. position `1`, the first array index;
2. position `n`, the last array index;
3. `pos1`, the position of the minimum value;
4. `posN`, the position of the maximum value.

The first and last positions cannot be set because an operation only affects indices strictly between `l` and `r`:

```text
l < i < r
```

The positions of values `1` and `n` cannot be set because an affected value must be strictly between the endpoint values:

```text
min(p_l, p_r) < p_i < max(p_l, p_r)
```

No value is smaller than `1`, and no value is larger than `n`, so neither `1` nor `n` can be strictly between two permutation values.

Therefore, if any of these four positions has `x_i = 1`, the answer is impossible.

### Five Operations Are Enough Otherwise

Assume all four impossible positions have `x_i = 0`.

We will output these five operations:

```text
(pos1, posN)
(1, pos1)
(1, posN)
(pos1, n)
(posN, n)
```

In the implementation, indices are stored as `0`-based internally and converted to `1`-based when printed.

Some of these operations may have `l > r` in internal order, so the code first uses the smaller index as `l` and the larger as `r` where needed.  Operations where an endpoint pair gives an empty affected set are harmless, because we only need to use at most five operations.

Now prove that every other position can be covered by at least one of these operations.

### Middle Positions

Consider any index `i` strictly between `pos1` and `posN`.

The operation:

```text
(pos1, posN)
```

uses endpoint values `1` and `n`.  Since `p_i` is neither `1` nor `n`, we have:

```text
1 < p_i < n
```

So every index strictly between `pos1` and `posN` is set to `1`.

### Positions Left of Both `pos1` and `posN`

Now consider an index `i` to the left of both `pos1` and `posN`.

Since `i` is not the first index, it lies strictly between index `1` and both `pos1`, `posN`.

Let the first value be:

```text
a = p_1
```

There are two cases:

- if `p_i < a`, then `p_i` lies strictly between `1` and `a`, so operation `(1, pos1)` covers it;
- if `p_i > a`, then `p_i` lies strictly between `a` and `n`, so operation `(1, posN)` covers it.

The equality case `p_i = a` cannot happen because `p` is a permutation and `i` is not the first index.

So every valid target position left of both extremes is covered.

### Positions Right of Both `pos1` and `posN`

The right side is symmetric.

Consider an index `i` to the right of both `pos1` and `posN`.

Since `i` is not the last index, it lies strictly between both `pos1`, `posN` and index `n`.

Let the last value be:

```text
b = p_n
```

There are two cases:

- if `p_i < b`, then `p_i` lies strictly between `1` and `b`, so operation `(pos1, n)` covers it;
- if `p_i > b`, then `p_i` lies strictly between `b` and `n`, so operation `(posN, n)` covers it.

Again, `p_i = b` cannot happen unless `i` is the last index, which we already excluded.

So every valid target position right of both extremes is covered.

### Conclusion

Every position except:

```text
first index, last index, pos1, posN
```

is covered by at least one of the five operations above.

Thus:

- if `x` requires any of those four impossible positions, print `-1`;
- otherwise, print the five operations.

The problem allows setting extra positions with `x_i = 0` to `1`, so it is fine that these operations may cover more positions than required.

### Complexity

For each test case, we only need to find the positions of `1` and `n`, then print at most five operations.

```text
Time:   O(n)
Memory: O(n)
```
