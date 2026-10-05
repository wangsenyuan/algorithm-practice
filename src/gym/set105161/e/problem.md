# E. Divide

[Problem link](https://codeforces.com/gym/105161/problem/E)

**Contest:** 2024 Jiangsu Collegiate Programming Contest

time limit per test: 6 seconds

memory limit per test: 1024 megabytes

## Problem

For each independent query on an integer subarray, repeatedly select its
leftmost maximum value and replace it with its floor half. Report the maximum
remaining after exactly `k` operations.

## Input

The first line contains `n` and `q` (`1 <= n, q <= 10^5`). The next line has
`n` values `a[i]` (`0 <= a[i] <= 10^5`). Each query gives `l`, `r`, and `k`
(`1 <= l <= r <= n`, `0 <= k <= 10^9`).

## Output

For each query, print the final maximum value.

## Solution

For a proposed final maximum `x`, each value `v` needs one operation for every
halving state above `x`. This is the count of powers of two `p` for which
`v >= (x + 1) * p`. A persistent segment tree answers every range-count
threshold, so binary search finds the smallest `x` whose required operations
do not exceed `k`.

Time complexity is `O((n + q log V log^2 V))`, where `V = 10^5`.

## Alternative `solve`: offline divide and conquer with a Fenwick tree

The current `solve` is a different implementation from `solve1`. It answers
all queries together instead of doing a separate binary search and persistent
tree query for every query.

### 1. Rephrase one reduction count

For one value `v`, define `D(v, x)` as the number of Reduce operations needed
to make it strictly smaller than `x`. Since every operation halves the value,

```text
D(v, x) = bits.Len(v / x)
```

for positive `x`. Integer division is intentional. For example, with `v = 9`:

```text
x = 6:  9 -> 4, so D(9, 6) = 1 = bits.Len(9 / 6)
x = 3:  9 -> 4 -> 2, so D(9, 3) = 2 = bits.Len(9 / 3)
```

Therefore, while narrowing an answer interval from `[low, high)` at its middle
`mid`, the extra operations contributed by `v` are

```text
D(v, mid) - D(v, high)
= bits.Len(v / mid) - bits.Len(v / high)
```

This is exactly the value added to the Fenwick tree in `play`.

### 2. Parallel binary search

`play(idx, qIdx, low, high)` processes every query whose answer is known to be
in the half-open integer interval `[low, high)`. The initial interval is
`[0, 100001)`, so `100001` is an exclusive sentinel above every possible
answer.

At `mid = (low + high) / 2`, the Fenwick tree temporarily stores the above
extra-operation contribution at each applicable original array index. A range
sum over a query's `[l, r]` is then the number of operations required to cross
the current boundary from `high` down to `mid` for that subarray.

- If `cnt <= k`, the query can pay for that transition. The code subtracts
  `cnt` from its remaining budget and sends it to `[low, mid)`.
- Otherwise, its answer is at least `mid`, so it is sent to `[mid, high)` with
  the same remaining budget.

When `high = low + 1`, only one integer answer remains; assigning `low` is
therefore correct.

### 3. Why `idx` splits into `b` and `c`

The same original value can matter in both children because its successive
halved values may cross the current midpoint. The code keeps only the values
that can still affect each child:

- `b` is for the lower child. `shift` finds the last value in `v, v/2, ...`
  that is still greater than `low`. If that value is below `mid`, the value can
  affect an answer in `[low, mid)`.
- `c` is for the upper child. Values no greater than `mid` cannot affect an
  answer in `[mid, high)`, so only `v > mid` is retained.

This filtering is important: it avoids repeatedly considering values that can
no longer change a query's answer.

### 4. Fenwick-tree lifecycle

The Fenwick tree is only scratch space for one `play` call. It is populated,
used to classify all of that call's queries, and then reset before the two
recursive calls. This permits one tree to be reused safely without any state
from a parent or sibling interval leaking into another calculation.

Compared with `solve1`, this version shares the answer-range partitioning and
range-count work among queries. It is an offline algorithm: all queries must
be read before it starts, which is fine because the problem's queries are
independent.

## `solve1`: per-query binary search with a persistent segment tree

`solve1` is the more direct version. It processes every query separately and
asks a monotone question:

> Can all values in this query range be reduced to at most `target` within
> `k` operations?

If the answer is true for a target, it remains true for every larger target.
Thus, ordinary binary search can find the smallest feasible target, which is
exactly the final maximum after `k` operations.

### Counting the operations for one target

Consider one value `v`. It contributes one required operation for each state
that is still greater than `target`:

```text
v, floor(v / 2), floor(v / 4), ...
```

The state after `j` previous reductions is greater than `target` exactly when

```text
v >= (target + 1) * 2^j
```

So `operationsNeeded` loops over these thresholds:

```go
for threshold := target + 1; threshold <= maxValue; threshold <<= 1
```

For every threshold it adds the number of array values in the query range that
are at least that threshold. A value such as `9`, with `target = 1`, is counted
at thresholds `2`, `4`, and `8`, representing `9 -> 4 -> 2 -> 1`: three
operations.

`limit` is the query's `k`. Once the running count exceeds it, the function
returns immediately because the binary-search predicate is already false.

### Persistent segment tree

`roots[i]` represents the frequency tree for the prefix `a[0:i]`. Inserting
`a[i]` creates a new root by copying only the `O(log V)` nodes on that value's
path; all untouched nodes are shared with the previous version.

For a query range `[left, right)`, subtracting the counts in `roots[left]`
from the corresponding counts in `roots[right]` gives the frequency
distribution of that subarray. `countLess` performs this subtraction while
walking the tree, returning how many values are smaller than a supplied bound.
Consequently,

```text
range count of values >= threshold
= range length - countLess(threshold)
```

is available in `O(log V)` time without rebuilding a data structure for each
query.

### Query flow and cost

For each query, `solve1` binary-searches `target` in `[0, 100000]`. Every
candidate invokes `operationsNeeded`; if its cost is at most `k`, the search
keeps the lower half, otherwise it keeps the upper half.

Building the persistent tree costs `O(n log V)` time and space. One feasibility
check uses `O(log^2 V)` time: `O(log V)` thresholds and an `O(log V)` tree
query for each. The binary search adds another `O(log V)`, so all queries take
`O(q log^3 V)` time after preprocessing.
