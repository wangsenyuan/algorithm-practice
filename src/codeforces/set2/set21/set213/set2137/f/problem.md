Given two arrays `x` and `y`, both of size `m`, let `z` be another array of size `m` such that the prefix maximum at each position of `z` is the same as the prefix maximum at each position of `x`.

Formally, for all `1 <= i <= m`:

`max(x1, x2, ..., xi) = max(z1, z2, ..., zi)`

Define `f(x, y)` to be the maximum number of positions where `zi = yi` over all possible arrays `z`.

You are given two sequences of integers `a` and `b`, both of size `n`. Find the value of:

`sum_{l=1..n} sum_{r=l..n} f([al, al+1, ..., ar], [bl, bl+1, ..., br])`

## Input

Each test contains multiple test cases.

- The first line contains the number of test cases `t` (`1 <= t <= 10^4`).
- The description of the test cases follows.

For each test case:

- The first line contains an integer `n` (`1 <= n <= 2 * 10^5`).
- The second line contains `n` integers `a1, a2, ..., an` (`1 <= ai <= 2 * n`).
- The third line contains `n` integers `b1, b2, ..., bn` (`1 <= bi <= 2 * n`).

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output:

`sum_{l=1..n} sum_{r=l..n} f([al, al+1, ..., ar], [bl, bl+1, ..., br])`

over all pairs `(l, r)`.

## Example

### Input

```text
6
3
5 3 1
4 2 1
5
1 2 3 4 5
1 2 3 4 5
6
7 1 12 10 5 8
9 2 4 3 6 5
1
1
2
6
5 1 2 6 3 4
3 1 6 5 2 4
2
2 2
1 1
```

### Output

```text
5
35
26
0
24
1
```

## Note

In the first test case, the answer is the sum of the following:

- `f([5], [4]) = 0`, using `z = [5]`.
- `f([3], [2]) = 0`, using `z = [3]`.
- `f([1], [1]) = 1`, using `z = [1]`.
- `f([5, 3], [4, 2]) = 1`, using `z = [5, 2]`.
- `f([3, 1], [2, 1]) = 1`, using `z = [3, 1]`.
- `f([5, 3, 1], [4, 2, 1]) = 2`, using `z = [5, 2, 1]`.


### ideas
1. 在给定(x, y)的情况下，可以先计算出w
2. w[0] = x[0], w[1] = max(w[0], x[1]), .. w[i] = max(w[i-1], x[i]), ...
3. 如果y[i] <= w[i-1], 那么+1, 否则+0, 这时候z[i] = y[i], 不会改变w的结果; 否则就不一致了
4. 

## Solution explanation

For one fixed subarray `x = a[l..r]`, define its required prefix maximum sequence:

```text
w[i] = max(x[0], x[1], ..., x[i])
```

We need to choose another array `z` such that its prefix maxima are exactly `w`, and we want as many positions as possible where `z[i] = y[i]`.

Consider one position `i`.

Before this position, the prefix maximum is:

```text
prev = max(x[0], ..., x[i-1])
```

If we set `z[i] = y[i]`, then the new prefix maximum of `z` becomes:

```text
max(prev, y[i])
```

The required prefix maximum from `x` is:

```text
max(prev, x[i])
```

So position `i` can match `y[i]` exactly when:

```text
max(prev, y[i]) = max(prev, x[i])
```

There are two cases.

### Case 1: `x[i] = y[i]`

Then the equality is always true:

```text
max(prev, y[i]) = max(prev, x[i])
```

So this position contributes `1` for every subarray containing it.

### Case 2: `x[i] != y[i]`

Then the only way the two maxima are equal is if both values are hidden below the previous maximum:

```text
prev >= max(x[i], y[i])
```

In this case, setting `z[i] = y[i]` does not change the prefix maximum, and the required prefix maximum also does not change.

If `prev < max(x[i], y[i])`, then one of `x[i]` or `y[i]` becomes the new prefix maximum while the other does not, so matching is impossible.

## Summing by position

Instead of evaluating every subarray separately, count the contribution of each global position `r`.

If position `r` contributes to a subarray `[l..R]`, then the right endpoint `R` can be anything from `r` to `n-1`. That gives:

```text
n - r
```

choices for the right endpoint.

So we only need to count how many left endpoints `l <= r` make position `r` matchable.

## When `a[r] = b[r]`

This is always matchable for every left endpoint.

There are `r + 1` choices for `l`, and `n - r` choices for `R`, so the contribution is:

```go
(r + 1) * (n - r)
```

This matches the code:

```go
if x[r] == y[r] {
	res += (r + 1) * (n - r)
}
```

## When `a[r] != b[r]`

We need:

```text
max(a[l..r-1]) >= max(a[r], b[r])
```

Let:

```text
need = max(a[r], b[r])
```

For a fixed `r`, suppose the latest index before `r` with `a[index] >= need` is `p`.

Then:

- if `l <= p`, the prefix `a[l..r-1]` contains that large enough value, so position `r` is matchable;
- if `l > p`, no previous value in the subarray is large enough, so it is not matchable.

Therefore the number of valid left endpoints is:

```text
p + 1
```

and the total contribution of position `r` is:

```go
(p + 1) * (n - r)
```

## Segment tree

While scanning `r` from left to right, the code maintains previous `a` values in a segment tree.

At value `v`, the tree stores the largest index seen so far with:

```text
a[index] = v
```

To find the latest previous index with value at least `need`, query the range:

```go
[need, 2*n+1)
```

This returns:

```go
l := tr.Get(max(x[r], y[r]), 2*n+1)
```

Then the contribution is:

```go
res += (l + 1) * (n - r)
```

After processing position `r`, insert `a[r]` into the segment tree:

```go
tr.Update(x[r], r)
```

This order is important: for the non-equal case, position `r` needs a previous maximum from `a[l..r-1]`, so `a[r]` itself must not be included in the query.

## Complexity

Each position does one range-maximum query and one point update on values up to `2n`.

Time complexity:

```text
O(n log n)
```

Memory complexity:

```text
O(n)
```
