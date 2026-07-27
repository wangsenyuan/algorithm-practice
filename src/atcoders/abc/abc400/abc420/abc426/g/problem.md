# G - Range Knapsack Query

[Problem link](https://atcoder.jp/contests/abc426/tasks/abc426_g)

**Contest:** [AtCoder Beginner Contest 426](https://atcoder.jp/contests/abc426)

time limit: 2 sec

memory limit: 1024 MiB

score: 575 points

## Problem

There are `N` items numbered `1` to `N`. Item `i` has weight `W_i` and value
`V_i`.

Process `Q` queries. The `j`-th query gives integers `L_j`, `R_j`, `C_j`
(`1 <= L_j <= R_j <= N`). Choose some (possibly zero) items from
`L_j .. R_j` whose total weight does not exceed `C_j`, and find the maximum
possible total value.

## Constraints

- `1 <= N <= 2 * 10^4`
- `1 <= Q <= 2 * 10^5`
- `1 <= W_i <= 500`
- `1 <= V_i <= 10^9`
- `1 <= L_j <= R_j <= N`
- `1 <= C_j <= 500`
- All input values are integers

## Input

```text
N
W_1 V_1
...
W_N V_N
Q
L_1 R_1 C_1
...
L_Q R_Q C_Q
```

## Output

Print `Q` lines. The `i`-th line is the answer for the `i`-th query.

## Sample 1

```text
Input
4
3 4
5 8
1 2
2 3
3
1 4 7
2 4 10
1 2 2

Output
11
13
0
```

## Sample 2

```text
Input
8
167 430302156
22 623690081
197 476190629
176 24979445
22 877914575
247 211047202
232 822804784
25 628894325
8
6 8 176
3 5 80
1 7 310
4 8 368
4 5 218
3 4 431
4 6 228
1 1 239

Output
628894325
877914575
2324409440
2329613684
902894020
501170074
902894020
430302156
```

## Solution

Let `K = max(C_j)`, so `K <= 500`.

A segment tree is not efficient here if two knapsack tables are combined with
a max-plus convolution. Merging two tables requires

```text
merged[c] = max(left[x] + right[c-x]) for 0 <= x <= c,
```

which costs `O(K^2)` per merge. This is already too expensive while building
the tree, and a query would require several such merges.

Instead, process all queries offline with divide and conquer on the item
indices.

### Divide the item interval

Consider a recursive call for the half-open interval `[left, right)`, and let

```text
mid = (left + right) / 2.
```

Every query assigned to this call belongs to exactly one of three groups:

1. `R <= mid`: the query is completely inside the left half.
2. `L >= mid`: the query is completely inside the right half.
3. `L < mid < R`: the query crosses the split.

The first two groups are passed to the corresponding recursive calls. Crossing
queries are answered at the current split.

Thus, every query is eventually answered exactly once: at the first recursive
split separating its left and right endpoints. A one-item interval is the base
case.

### Knapsack tables around the split

For crossing queries, precompute the following tables:

```text
dp[i][c] = maximum value with total weight at most c
```

Their represented item ranges depend on which side of `mid` they are on:

```text
i < mid: dp[i] uses items [i, mid)
i > mid: dp[i] uses items [mid, i)
dp[mid]: the empty range
```

Starting with `dp[mid][c] = 0`, build the left tables from right to left and
the right tables from left to right. Each new table adds exactly one item.

If that item has weight `w` and value `v`, the ordinary 0/1-knapsack
transition is

```text
current[c] = previous[c]                              if c < w
current[c] = max(previous[c], previous[c-w] + v)     if c >= w
```

One table therefore takes `O(K)` time to construct. The implementation only
computes capacities up to the maximum capacity of the crossing queries at the
current split.

### Answer a crossing query

For a query `[L, R)` with capacity `C`, the two relevant item ranges are
disjoint:

```text
[L, mid) and [mid, R).
```

Suppose `x` capacity is reserved for the left part. Then `C-x` capacity remains
for the right part, giving

```text
dp[L][x] + dp[R][C-x].
```

Trying every possible split of the capacity gives

```text
answer = max(dp[L][x] + dp[R][C-x]) for 0 <= x <= C.
```

This costs `O(K)` per query.

### Correctness

For every left table, the knapsack transition considers both possibilities
for the newly added item: exclude it and retain `previous[c]`, or include it
and add its value to an optimal solution with capacity `c-w`. Therefore
`dp[i][c]` is optimal for `[i, mid)`. The same argument proves that every right
table is optimal for `[mid, i)`.

For a crossing query, take any feasible selection and let `x` be the total
weight selected from its left part. Its left-side value is at most
`dp[L][x]`, and its right-side value is at most `dp[R][C-x]`. Hence its total
value is no greater than one of the candidates considered by the formula.

Conversely, the two DP solutions used by any candidate come from disjoint item
ranges and have combined weight at most `C`, so their union is a feasible
selection for the query. The maximum candidate is therefore exactly the
optimal answer.

### Complexity

At one recursion depth, each item participates in at most one `O(K)` knapsack
update. There are `O(log N)` depths, so all table construction costs

```text
O(K N log N).
```

Each query is answered once in `O(K)`, while routing query IDs through the
recursion costs `O(Q log N)`. The total time is

```text
O(K(N log N + Q) + Q log N),
```

usually written as `O(K(N log N + Q))` because `K = 500`.

The reusable DP table contains `N+1` rows of `K+1` values, requiring `O(NK)`
memory. The recursively partitioned query-ID slices require at most
`O(Q log N)` additional memory in this implementation. Thus, the total memory
complexity is

```text
O(NK + Q log N).
```
