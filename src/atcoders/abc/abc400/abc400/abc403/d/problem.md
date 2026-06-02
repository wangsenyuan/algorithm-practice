# Problem D - Forbidden Difference

https://atcoder.jp/contests/abc403/tasks/abc403_d

**Time Limit:** 2 sec / **Memory Limit:** 1024 MiB

**Score:** 425 points

## Problem Statement

You are given a length-$N$ integer sequence
$A=(A_1,A_2,\dots,A_N)$ and a non-negative integer $D$.

We wish to delete as few elements as possible from $A$ to obtain a sequence $B$
that satisfies the following condition:

- $|B_i-B_j| \neq D$ for all $i,j$ $(1 \leq i < j \leq |B|)$.

Find the minimum number of deletions required.

## Constraints

- $1 \leq N \leq 2 \times 10^5$
- $0 \leq D \leq 10^6$
- $0 \leq A_i \leq 10^6$
- All input values are integers.


## Solution

Instead of directly deciding which elements to delete, decide which values to
keep.

For a value `x`, all occurrences of `x` are equivalent:

- if we keep value `x`, we should keep every occurrence of `x`;
- if value `x` conflicts with another kept value, we may need to delete all
  occurrences of one of them.

So first count how many times each distinct value appears. The weight of value
`x` is:

```text
cnt[x] = number of occurrences of x
```

Then the problem becomes:

```text
choose values with maximum total count,
such that no two chosen values differ by D
```

If the maximum number of kept elements is `keep`, the answer is:

```text
N - keep
```

### Case `D = 0`

When `D = 0`, two equal values conflict because `|x - x| = 0`.

Therefore, each value can appear at most once in the remaining sequence. For
each distinct value, keep one occurrence and delete the rest.

So:

```text
answer = N - number of distinct values
```

This is why the code sorts the array, compacts equal values, and returns
`n - len(compacted)`.

### Case `D > 0`

For `D > 0`, value `x` only conflicts with values:

```text
x - D
x + D
```

Values with different remainders modulo `D` never conflict, because if
`|x-y| = D`, then `x` and `y` have the same remainder modulo `D`.

So we can split all values into independent chains by `x % D`.

For one remainder class, after sorting distinct values, only consecutive values
whose difference is exactly `D` are connected:

```text
r, r+D, r+2D, r+3D, ...
```

If there is a gap larger than `D`, for example:

```text
x, x+D, x+3D
```

then `x+D` and `x+3D` do not conflict, so this starts a new independent chain.

The code stores each such chain as:

```go
[]pair{{value, count}, ...}
```

and calls `play` to compute the maximum number of elements that can be kept from
that chain.

### DP on One Chain

Inside one chain, adjacent values differ by exactly `D`, so adjacent values
cannot both be kept. Non-adjacent values can both be kept.

This is the classic maximum-weight independent set on a path.

Let:

```text
weight[i] = count of the i-th value in the chain
dp[i] = maximum kept elements using values 0..i
```

For value `i`, there are two choices:

1. Do not keep value `i`: contribution is `dp[i-1]`.
2. Keep value `i`: then value `i-1` cannot be kept, contribution is
   `dp[i-2] + weight[i]`.

So:

```text
dp[i] = max(dp[i-1], dp[i-2] + weight[i])
```

The base cases are:

```text
dp[0] = weight[0]
dp[1] = max(weight[0], weight[1])
```

The implementation keeps only the last two DP values:

```go
x, y := arr[0].second, arr[1].second
y = max(x, y)
for i := 2; i < len(arr); i++ {
    z := max(y, arr[i].second+x)
    x, y = y, z
}
```

Here `x` is `dp[i-2]`, `y` is `dp[i-1]`, and `z` becomes `dp[i]`.

The important detail is the initialization of `y`:

```go
y = max(arr[0].second, arr[1].second)
```

For the first two adjacent values, we cannot keep both, so the best answer is
the larger of their counts.

### Complexity

The array is sorted once, and each distinct value belongs to exactly one chain.

The total complexity is:

```text
O(N log N)
```

from sorting, with `O(N)` additional processing.
