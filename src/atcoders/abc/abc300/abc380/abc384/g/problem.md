# G - Abs Sum

Problem: [AtCoder Beginner Contest 384 G](https://atcoder.jp/contests/abc384/tasks/abc384_g)

## Problem Statement

You are given two integer sequences `A = (A_1, A_2, ..., A_N)` and
`B = (B_1, B_2, ..., B_N)`, both of length `N`, and `K` pairs
`(X_1, Y_1), (X_2, Y_2), ..., (X_K, Y_K)`.

For every `k = 1, 2, ..., K`, calculate

```text
sum_{i=1}^{X_k} sum_{j=1}^{Y_k} |A_i - B_j|.
```

## Constraints

- `1 <= N <= 10^5`
- `0 <= A_i, B_i <= 2 * 10^8`
- `1 <= K <= 10^4`
- `1 <= X_k, Y_k <= N`
- All input values are integers.

## Input

```text
N
A_1 A_2 ... A_N
B_1 B_2 ... B_N
K
X_1 Y_1
X_2 Y_2
...
X_K Y_K
```

## Output

Print `K` lines. The `k`-th line must contain the answer for `(X_k, Y_k)`.

## Sample 1

```text
Input
2
2 4
3 5
4
1 1
1 2
2 1
2 2

Output
1
4
2
6
```

## Sample 2

```text
Input
5
1163686 28892 1263085 2347878 520306
1332157 1202905 2437161 1291976 563395
5
5 3
1 5
2 3
1 2
5 5

Output
13331322
2209746
6366712
207690
20241215
```

## Solution

Define

```text
F(x, y) = sum_{i=1}^{x} sum_{j=1}^{y} |A_i - B_j|.
```

Every query asks for one value `F(X_k, Y_k)`. Computing it independently takes
`O(X_k * Y_k)` time, which can be `O(N^2)` for one query and is far too slow for
`K <= 10^4` queries.

The important observation is that two nearby states differ by only one row or one
column of the conceptual `N * N` matrix

```text
M[i][j] = |A_i - B_j|.
```

For example,

```text
F(x + 1, y) - F(x, y) = sum_{j=1}^{y} |A_{x+1} - B_j|,
F(x, y + 1) - F(x, y) = sum_{i=1}^{x} |A_i - B_{y+1}|.
```

If these deltas can be evaluated quickly, the queries can be reordered with Mo's
algorithm so that the current `(x, y)` changes only a limited number of times.

### Current State

While processing the reordered queries, maintain the prefixes

```text
A[0:x] and B[0:y].
```

The code maintains the following invariant:

- `x` and `y` are the lengths of the two active prefixes;
- `asum` is the sum of `A[0:x]`;
- `bsum` is the sum of `B[0:y]`;
- `now` is exactly
  `sum_{i=0}^{x-1} sum_{j=0}^{y-1} |A[i] - B[j]|`;
- the `fa` Fenwick trees contain the active values from `A[0:x]`;
- the `fb` Fenwick trees contain the active values from `B[0:y]`.

There are two Fenwick trees for each side:

```text
fa0: frequency of each active A value
fa1: sum of each active A value
fb0: frequency of each active B value
fb1: sum of each active B value
```

Counts alone are insufficient: to sum absolute differences, both the number and
the sum of values on each side of a threshold are needed.

### Deriving the Cost of Adding One Value

Suppose the current active `B` prefix contains `y` values with total sum `bsum`,
and a new value `v = A[x]` is added to the active `A` prefix. Its contribution is

```text
sum over active B[j] of |v - B[j]|.
```

Split the active `B` values into two groups:

```text
low:  B[j] <= v
high: B[j] > v
```

Let:

```text
C = number of values in low
S = sum of values in low
```

For the low group, every absolute value opens as `v - B[j]`, so its contribution
is

```text
C * v - S.
```

The high group contains `y - C` values and has sum `bsum - S`. Every absolute
value opens as `B[j] - v`, so its contribution is

```text
(bsum - S) - (y - C) * v.
```

Adding the two parts gives

```text
C*v - S + (bsum - S) - (y-C)*v
= (2*C - y)*v + bsum - 2*S
= y*v + bsum - 2*((y-C)*v + S).
```

The last form is the one used in the code:

```go
now += y*A[x] + bsum - 2*((y-c)*A[x]+s)
```

The values `c` and `s` are obtained from `fb0` and `fb1` in `O(log N)` time.

Adding a new `B[y]` is symmetric. With

```text
C = number of active A values <= B[y]
S = sum of active A values <= B[y]
```

its contribution is

```text
x*B[y] + asum - 2*((x-C)*B[y] + S).
```

### Why Including Equal Values Is Correct

An equal value may be placed in either the low or high group because its absolute
difference is zero. This implementation puts equal values in the low group by
querying values `<= v`.

If one value equal to `v` moves from the high group to the low group, then both
`C` and `S` increase by `1` and `v`, respectively. In

```text
(2*C - y)*v + totalSum - 2*S,
```

the change is `2*v - 2*v = 0`. Therefore both `< v` and `<= v` partitions lead
to the same final contribution, provided the Fenwick boundary matches the chosen
definition consistently.

### Coordinate Compression and the Fenwick Boundary

Fenwick trees need integer positions in a small range, while values can be as
large as `2 * 10^8`. Combine all values from `A` and `B`, sort them, remove
duplicates, and replace every original value by its rank.

Coordinate compression preserves order:

```text
u <= v  if and only if  rank(u) <= rank(v).
```

Thus a Fenwick prefix query ending at `rank(v)` returns the count or sum of all
active values at most `v`.

The local Fenwick implementation uses inclusive endpoints:

```go
func (tr fenwick) queryRange(l, r int) int {
    return tr.query(r) - tr.query(l-1)
}
```

Therefore the correct call for values `<= v` is

```go
queryRange(0, rank(v))
```

and not `queryRange(0, rank(v)+1)`. The latter would include the next larger
compressed value and could even make an absolute-difference contribution
negative.

### Adding and Removing Prefix Elements

To move from the current state to a query `(q.x, q.y)`, adjust the two prefix
lengths one step at a time.

#### Increase `x`

1. Query the active `B` trees for the count and sum of values `<= A[x]`.
2. Add the resulting contribution to `now`.
3. Add `A[x]` to `asum`, `fa0`, and `fa1`.
4. Increment `x`.

#### Decrease `x`

1. Decrement `x`, making `A[x]` the element to remove.
2. Compute its contribution against the still-active `B` prefix.
3. Subtract that contribution from `now`.
4. Remove `A[x]` from `asum`, `fa0`, and `fa1`.

Increasing and decreasing `y` use exactly the symmetric operations with the
active `A` trees. Removal is the inverse of insertion because the opposite prefix
does not change between computing and subtracting the contribution.

### Mo's Algorithm Ordering

Processing queries in input order could move both prefix endpoints by `O(N)` for
every query. Instead, choose a block size `M` and sort by:

1. `floor(X_k / M)`;
2. `Y_k` ascending in even `X` blocks and descending in odd `X` blocks.

The alternating direction is sometimes called the snake ordering. At the boundary
between two adjacent `X` blocks, it avoids moving `y` from one extreme all the way
back to the other extreme.

Inside one `X` block:

- all `x` values lie in an interval of length at most `M`, so total `x` movement
  over all queries is `O(KM)`;
- `y` moves monotonically from one end toward the other, so it moves `O(N)` times
  per `X` block.

There are `O(N/M)` possible `X` blocks, giving

```text
total pointer movements = O(KM + N^2/M).
```

Balancing the two terms gives

```text
M = N / sqrt(K).
```

At the maximum constraints, this is `10^5 / sqrt(10^4) = 1000`, which is the
constant used by the implementation.

The answers are written into `ans[q.id]`, so reordering the queries does not alter
the required output order.

### Correctness Proof

We prove that the algorithm returns the correct answer for every query.

**Lemma 1:** When a value `v` is added to the active `A` prefix, the computed delta
equals the sum of `|v-B[j]|` over every active `B[j]`.

**Proof:** Partition the active `B` values into those at most `v` and those greater
than `v`. If the first group has count `C` and sum `S`, its contribution is
`C*v-S`. The second group has count `y-C` and sum `bsum-S`, so its contribution is
`(bsum-S)-(y-C)*v`. Their sum is exactly the formula added to `now`. ∎

**Lemma 2:** The symmetric formula for adding a value to the active `B` prefix is
correct.

**Proof:** Apply Lemma 1 after exchanging the roles of `A` and `B`. Since
`|A[i]-B[j]| = |B[j]-A[i]|`, the same derivation applies. ∎

**Lemma 3:** After every single movement of `x` or `y`, `now` equals `F(x,y)` and
the Fenwick trees contain exactly the active prefixes.

**Proof:** Initially both prefixes are empty, all trees and sums are zero, and
`now = F(0,0) = 0`. For an insertion, Lemma 1 or Lemma 2 adds exactly all newly
created pairs, then the inserted value is recorded in the corresponding trees.
For a removal, the same contribution is computed before removing the value and is
subtracted, deleting exactly all pairs involving that value. No other pair changes.
Therefore the invariant is preserved by every movement. ∎

**Theorem:** For every input query `(X_k,Y_k)`, the algorithm stores `F(X_k,Y_k)`
in its original answer position.

**Proof:** The algorithm moves from the preceding state until `x = X_k` and
`y = Y_k`. By Lemma 3, `now = F(X_k,Y_k)` at that point. It stores `now` using the
query's original identifier, so every query receives the correct value regardless
of the processing order. ∎

### Complexity

Let `M` be the Mo block size.

- Coordinate compression: `O(N log N)` time.
- Sorting the queries: `O(K log K)` time.
- Pointer movements: `O(KM + N^2/M)`.
- Each movement performs `O(log N)` Fenwick work.

Thus the total time is

```text
O(N log N + K log K + (KM + N^2/M) log N).
```

With `M = Theta(N/sqrt(K))`, this becomes

```text
O(N * sqrt(K) * log N).
```

The four Fenwick trees, compressed arrays, reordered queries, and answers use
`O(N + K)` space.

The maximum possible answer is `N^2 * 2*10^8 = 2*10^18`, so the accumulated
answer must use a signed 64-bit integer. The Go judge runs on a 64-bit target, on
which the implementation's `int` accumulators are 64-bit; `solve` exposes the
answers as `[]int64`.

