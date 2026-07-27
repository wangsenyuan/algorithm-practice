# D. Cutting a Fence

[Problem link](https://codeforces.com/problemset/problem/212/D)

**Contest:** [Codeforces Round #131 (Div. 1)](https://codeforces.com/contest/212)

time limit: 5 seconds

memory limit: 256 megabytes

## Problem

Vasya has a fence of `n` planks in a line with heights `a1, a2, ..., an`.

He paints a fuchsia rectangle covering `k` consecutive planks starting at
position `x` (`1 <= x <= n - k + 1`). The rectangle height is
`min(a[x .. x+k-1])`.

Given `m` widths `k1, ..., km`, for each `ki` find the expected height when `x`
is chosen uniformly among all `n - ki + 1` positions.

## Constraints

- `1 <= n <= 10^6`
- `1 <= ai <= 10^9`
- `1 <= m <= 10^6`
- `1 <= ki <= n`

## Input

```text
n
a1 a2 ... an
m
k1 k2 ... km
```

## Output

Print `m` real numbers — the expected heights for each `ki`. Absolute or
relative error at most `1e-9` is accepted.

## Sample 1

```text
Input
3
3 2 1
3
1 2 3

Output
2.000000000000000
1.500000000000000
1.000000000000000
```

For `k = 1`: heights `3, 2, 1`, expectation `2`.
For `k = 2`: heights `2, 1`, expectation `1.5`.
For `k = 3`: height `1`, expectation `1`.

## Sample 2

```text
Input
2
1 1
3
1 2 1

Output
1.000000000000000
1.000000000000000
1.000000000000000
```


### ideas
1. 给定k的时候, f(1) = 从1...k的最低高度, f(2) = 2...k+1的最低高度
2. sum(f(1), ... f(x)) / x , x = n - k + 1
3. L[i], R[i], a[L[i]] < a[i] >= a[R[i]]
4. 那么如果有一个k, 它cover了i,但是没有cover a[i]的时候, a[i]就对它作出了贡献
5. k, 以及次数, 可以被计算出来
6. k = R[i] - L[i] - 1 是最大的k, (也就是在区间1...k)
7. k = 1的时候, 贡献1次, 2的时候, 贡献2次, w = min(i - L[i], R[i] - i)
8. 在1.j.w, a[i]的贡献 = j次
9. v = max(i - L[i], R[i] - i), 在[w+1, v], 这个阶段, 它的贡献就是w次
10. 在 [v+1, k],它的贡献 = k - j + 1 次

## Solution

For every width `k`, let `sum[k]` be the sum of the minimum heights of all
subarrays of length `k`. The requested expectation is

```text
sum[k] / (n-k+1).
```

The task is therefore to calculate `sum[k]` for every `1 <= k <= n` in linear
time.

### Assign each subarray to one minimum

Use a monotonic stack to find:

- `L[i]`: the nearest position to the left with `a[L[i]] < a[i]`;
- `R[i]`: the nearest position to the right with `a[R[i]] <= a[i]`.

The asymmetric comparisons are important when equal heights occur. They assign
each subarray to exactly one occurrence of its minimum instead of counting it
multiple times.

Index `i` is the assigned minimum precisely for subarrays that contain `i` and
stay inside `(L[i], R[i])`.

Define

```text
left  = i - L[i]
right = R[i] - i
w = min(left, right)
v = max(left, right)
d = left + right - 1 = R[i] - L[i] - 1
```

For a fixed length `k`, the number of such subarrays is

```text
count(i, k) =
    k           if 1 <= k <= w
    w           if w < k <= v
    d-k+1       if v < k <= d
    0           otherwise.
```

Thus, the contribution of `i` to `sum[k]` is
`a[i] * count(i, k)`. Across increasing values of `k`, this contribution forms
a trapezoid.

### The `change` array

Adding the complete trapezoid separately for every index would take
`O(n^2)`. Instead, record only the points where its slope changes.

For one index with height `h = a[i]`, its contribution has the following
slopes:

```text
lengths 1 ... w:       +h
lengths w+1 ... v:      0
lengths v+1 ... d+1:   -h
lengths d+2 onward:     0
```

Therefore, its four slope-change events are

```text
change[1]   += h
change[w+1] -= h
change[v+1] -= h
change[d+2] += h
```

The last event is at `d+2`, not `d+1`: the contribution is still decreasing
from `1*h` at length `d` to zero at length `d+1`. Its slope returns from `-h`
to zero only at the next position.

After adding these four events for every index, use two prefix accumulations:

```text
slope += change[k]
sum[k] = sum[k-1] + slope
```

The first prefix sum reconstructs

```text
slope[k] = sum[k] - sum[k-1],
```

and the second reconstructs `sum[k]`. Because addition is linear, combining
all indices in one `change` array reconstructs the total contribution of all
subarray minima for every length.

Finally, answer every query `k` with

```text
float64(sum[k]) / float64(n-k+1).
```

All contribution sums and slope events are stored as `int64`.

### Correctness

The asymmetric monotonic-stack boundaries assign each subarray to exactly one
index whose height equals its minimum. For a fixed index, `count(i, k)` counts
exactly the length-`k` subarrays assigned to it, so
`a[i] * count(i, k)` is exactly its contribution to `sum[k]`.

The four updates in `change` are the second differences of this contribution
sequence. Two prefix accumulations reconstruct that sequence without changing
any value. Summing these events for all indices therefore produces the sum of
the minima of every length-`k` subarray exactly once. Dividing by the number
`n-k+1` of such subarrays gives the required expectation.

### Complexity

The monotonic stack, contribution events, prefix reconstruction, and query
answers are all linear:

```text
Time:  O(n + m)
Space: O(n + m)
```
