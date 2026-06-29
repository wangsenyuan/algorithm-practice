# G - Range Set Modifying Query

[Problem link](https://atcoder.jp/contests/abc430/tasks/abc430_g)

**Contest:** [AtCoder Beginner Contest 430](https://atcoder.jp/contests/abc430)

time limit: 4 sec

memory limit: 1024 MiB

score: 625 points

There are `N` sets `S_1, ..., S_N`. Initially, all sets are empty.

You are given `Q` queries in the following formats. Process them in order.

- **Type 1:** `1 L R x` — for each `S_i` with `L <= i <= R`, add `x`.
- **Type 2:** `2 L R x` — for each `S_i` with `L <= i <= R`, remove `x`.
- **Type 3:** `3 L R` — among `S_i` with `L <= i <= R`, find the maximum number of elements and the
  number of sets that achieve this maximum.

## Constraints

- `1 <= N <= 3 * 10^5`
- `1 <= Q <= 3 * 10^5`
- For each query, `1 <= L <= R <= N`.
- For type 1 and 2 queries, `1 <= x <= 60`.
- All input values are integers.

## Input

```text
N Q
query_1
...
query_Q
```

Each query is one of:

```text
1 L R x
2 L R x
3 L R
```

## Output

Let `q` be the number of type 3 queries. Print `q` lines. The `i`-th line should contain two integers
`x` and `y` separated by a space, where `x` is the maximum number of elements and `y` is the number
of sets that achieve this maximum for the `i`-th type 3 query.

## Sample Input 1

```text
4 7
1 1 2 10
1 2 4 20
3 1 3
2 1 2 20
1 2 3 10
3 1 2
3 1 4
```

## Sample Output 1

```text
2 1
1 2
2 1
```

### Note

- After the first two updates, `(S_1, S_2, S_3, S_4) = ({10}, {10, 20}, {20}, {20})`.
- The first type 3 query on `S_1, S_2, S_3` has maximum size `2`, achieved by one set (`S_2`).
- After removing `20` from `S_1, S_2` and adding `10` to `S_2, S_3`, the final type 3 query on all
  four sets again has maximum size `2`, achieved by `S_3` only.

## Solution

The important detail is that adding or removing an element from a set is **idempotent**:

- adding `x` to a set that already contains `x` changes nothing;
- removing `x` from a set that does not contain `x` changes nothing.

So we should update the answer only on positions where the membership of `x` really changes.

The solution keeps two kinds of data structures:

1. For every value `x` from `1` to `60`, store the positions `i` where `x` is currently in `S_i`.
   These positions are stored as disjoint intervals in a treap.
2. A global lazy segment tree stores `|S_i|`, the size of every set. It supports range `+1`, range
   `-1`, and range query for maximum size and count.

### Global Segment Tree

For every segment tree node, store:

```text
best = maximum |S_i| inside this segment
cnt  = number of positions in this segment whose |S_i| equals best
lazy = pending range addition
```

Initially all sets are empty, so every `best` is `0`, and each leaf has `cnt = 1`.

When merging two children:

- the parent `best` is the larger child maximum;
- the parent `cnt` is the sum of counts from children whose maximum equals that larger value.

For example:

```text
left:  best = 3, cnt = 2
right: best = 5, cnt = 4
parent: best = 5, cnt = 4
```

If both sides have the same maximum:

```text
left:  best = 5, cnt = 2
right: best = 5, cnt = 4
parent: best = 5, cnt = 6
```

This segment tree answers type 3 queries directly.

### Intervals for One Value

For a fixed value `x`, consider all indices `i` where `x` belongs to `S_i`.

Instead of storing every index separately, store maximal disjoint intervals:

```text
[l1, r1], [l2, r2], ...
```

For example, if `x` is present in sets:

```text
S_2, S_3, S_4, S_7, S_8
```

then the intervals for `x` are:

```text
[2, 4], [7, 8]
```

The implementation stores these intervals in a treap ordered by the left endpoint. The treap gives
expected `O(log M)` insertion, deletion, predecessor, and lower-bound search, where `M` is the
number of intervals for this value.

### Type 1: Add `x` on `[L, R]`

We need to add `x` only to positions where `x` is not already present.

Look at the current intervals of `x` that intersect or touch `[L, R]`.

Example:

```text
current intervals for x: [2, 4], [8, 10]
query add x to:          [3, 8]
```

The old intervals already cover:

```text
[3, 4] and [8, 8]
```

So the only positions whose set size really increases are:

```text
[5, 7]
```

The global segment tree receives:

```text
add +1 on [5, 7]
```

Then all touched intervals are merged into one new interval:

```text
[2, 10]
```

The implementation does this by:

1. Checking the previous interval whose left endpoint is at most `L`; it may overlap or touch
   `[L, R]`.
2. Iterating intervals whose left endpoint is at most `R + 1`.
3. For every gap not already covered, applying `+1` to the global segment tree.
4. Removing the old touched intervals and inserting the merged interval.

Touching intervals are merged too, because after adding `[L, R]`, intervals such as `[1, 3]` and
`[4, 6]` become one continuous interval `[1, 6]`.

### Type 2: Remove `x` on `[L, R]`

Now we need to subtract only positions where `x` is actually present.

Example:

```text
current intervals for x: [2, 10]
query remove x from:     [4, 7]
```

Only `[4, 7]` changes, so the global segment tree receives:

```text
add -1 on [4, 7]
```

The old interval is split into:

```text
[2, 3], [8, 10]
```

The implementation iterates all intervals that overlap `[L, R]`:

1. Remove the old interval from the treap.
2. Apply `-1` to the overlap with `[L, R]`.
3. Reinsert the remaining left piece, if any.
4. Reinsert the remaining right piece, if any.

If a remove query touches positions where `x` is absent, nothing is applied there.

### Why This Handles Repeated Operations

Suppose `x` is already present on `[2, 5]`, and we add `x` to `[3, 4]`.

There is no uncovered gap, so the global segment tree receives no `+1`. This is correct because set
membership did not change.

Similarly, if `x` is absent on `[10, 20]`, and we remove `x` from `[10, 20]`, no interval overlaps
the query, so the global segment tree receives no `-1`.

This is why interval sets are useful: they let us update only the positions whose membership changes.

### Correctness

For each value `x`, the treap stores exactly the maximal intervals of indices whose sets contain
`x`. The add operation fills only uncovered gaps in `[L, R]`, then merges all touched intervals, so
after it finishes, the interval representation is again exact. The remove operation subtracts only
the overlap between `[L, R]` and existing intervals, then reinserts the pieces outside the removed
range, so it also preserves the exact representation.

Whenever a position starts containing a new value, its set size increases by `1`; whenever it stops
containing a value, its set size decreases by `1`. The global segment tree receives exactly these
range changes and no others, so it always stores the correct value of `|S_i|` for every position.

A type 3 query asks for the maximum `|S_i|` and the number of positions that achieve it in `[L, R]`.
The segment tree node merge rule maintains exactly that information, so the range query returns the
required answer.

### Complexity

Each changed interval is removed or inserted once in a treap, and each changed gap/overlap causes one
range update in the global segment tree.

Across all operations for one value, interval splits and merges are charged to interval changes. The
expected cost of each treap operation is `O(log Q)`, and each segment tree range operation costs
`O(log N)`.

Memory usage is `O(N + total number of stored intervals)`.
