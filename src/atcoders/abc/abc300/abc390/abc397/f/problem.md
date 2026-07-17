# F - Variety Split Hard

[Problem link](https://atcoder.jp/contests/abc397/tasks/abc397_f)

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 550 points

## Problem Statement

> This problem is a harder version of Problem C. Here, the sequence is split
> into three subarrays.

You are given an integer sequence of length `N`: `A = (A_1, A_2, ..., A_N)`.

When splitting `A` at two positions into three non-empty (contiguous)
subarrays, find the maximum possible sum of the counts of distinct integers
in those subarrays.

More formally, find the maximum sum of the following three values for a pair
of integers `(i, j)` such that `1 <= i < j <= N-1`:

- the count of distinct integers in `(A_1, ..., A_i)`
- the count of distinct integers in `(A_{i+1}, ..., A_j)`
- the count of distinct integers in `(A_{j+1}, ..., A_N)`

## Constraints

- `3 <= N <= 3 * 10^5`
- `1 <= A_i <= N`
- All input values are integers

## Input

```
N
A_1 A_2 ... A_N
```

## Output

Print the answer.

## Samples

### Sample 1

Input:

```
5
3 1 4 1 5
```

Output:

```
5
```

For `(i, j) = (2, 4)`, the parts `(3, 1)`, `(4, 1)`, `(5)` have distinct
counts `2, 2, 1` (sum `5`).

### Sample 2

Input:

```
10
2 5 6 4 4 1 1 3 1 4
```

Output:

```
9
```

## Solution

Use 0-based indices in the explanation. Suppose the second cut is after
position `i`, so the three parts are

```text
[0 ... k] | [k+1 ... i] | [i+1 ... N-1]
```

where `k` is the position of the first cut.

We sweep `i` from left to right. During the sweep:

- `len(last)` is the number of distinct values in `[0 ... i]`;
- `len(suf)` is the number of distinct values in `[i+1 ... N-1]`;
- the segment tree stores one value for every possible first cut `k`.

The only difficult part is quickly finding the best sum of distinct counts in
the first two subarrays.

### Turning Two Distinct Counts into an Overlap Count

For a fixed `i` and `k`, let

- `L` be the set of values in `[0 ... k]`;
- `M` be the set of values in `[k+1 ... i]`.

Then

```text
|L| + |M| = |L union M| + |L intersection M|
```

The union is exactly the set of distinct values in `[0 ... i]`, whose size is
`len(last)`. Therefore,

```text
distinct(0 ... k) + distinct(k+1 ... i)
= len(last) + overlap(k)
```

where `overlap(k)` is the number of values that occur on both sides of cut
`k` inside `[0 ... i]`.

So, for the current `i`, the answer is

```text
len(last) + max(overlap(k)) + len(suf)
```

The segment tree maintains all `overlap(k)` values and returns their maximum.

### How Processing One Value Changes the Cuts

Assume the current value is `v = A[i]`, and its previous occurrence was at
position `j`.

For every cut `k` in `[j, i-1]`:

- the occurrence at `j` lies in `[0 ... k]`;
- the occurrence at `i` lies in `[k+1 ... i]`.

Thus `v` now occurs on both sides of each such cut, so their overlap counts
increase by `1`:

```go
tr.update(j, i-1, 1)
```

No other cut changes:

- if `k < j`, the new occurrence at `i` does not create a new crossing of the
  cut; any crossing by `v` was already represented;
- if `k >= i`, the current position is not to the right of the cut.

After the update, `last[v] = i` records the latest occurrence. This is a range
addition, and we need the maximum over a range, so a lazy segment tree supports
both operations in `O(log N)` time.

### Maintaining the Suffix

At the start, `suf` contains the frequency of every value in the whole array.
Before evaluating position `i`, remove `A[i]` from it. If its frequency becomes
zero, delete the key. Consequently, `len(suf)` is always the number of distinct
values strictly after `i`.

### Why the Code Also Checks Boundary States

For a literal three-part split, valid indices satisfy `k < i` and `i < N-1`.
The implementation also queries `k = i`, processes `i = N-1`, and initializes
the answer with the distinct count of the whole array. These states correspond
to temporarily allowing an empty part.

This does not make the answer too large. Since `N >= 3`, any partition with an
empty part can be refined into three non-empty contiguous parts. Splitting a
non-empty segment cannot decrease the sum of distinct counts, because every
value present before the split remains present in at least one of the two new
segments. Therefore every boundary-state score is no greater than some valid
three-part score.

### Correctness Proof

We prove that the algorithm returns the maximum possible sum.

For a fixed sweep position `i`, the segment-tree value at cut `k` equals the
number of distinct values occurring in both `[0 ... k]` and `[k+1 ... i]`.
Initially all values are zero. When processing `A[i]`, if its previous
occurrence is `j`, exactly the cuts `j <= k < i` newly place this value on both
sides, and the range update adds one to exactly those cuts. Hence the invariant
holds after every iteration.

By the set identity above, for every first cut `k`, `len(last) + overlap(k)` is
exactly the sum of distinct counts in the first two parts. The segment-tree
query chooses the maximum such value. Adding `len(suf)`, which is exactly the
distinct count of the third part, gives the best score for the current second
cut `i`.

The sweep considers every possible second cut, so it considers every valid
pair of cuts. The extra boundary states cannot exceed the score of a valid
split, as shown above. Therefore the maximum recorded by the algorithm is
exactly the required answer.

### Complexity

- Time: `O(N log N)`
- Space: `O(N)`
