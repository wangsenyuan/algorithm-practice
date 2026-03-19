### Problem

Sereja has `m` non-empty sets of integers `A1, A2, ..., Am`. These sets form a **partition** of `{1, 2, ..., n}` — every integer `v (1 ≤ v ≤ n)` belongs to **exactly one** set.

Sereja also has an integer `d`.

He chooses some sets. Let the chosen set indices be:

`i1 < i2 < ... < ik`.

Let `b` be the sorted (in increasing order) array that contains the union of the chosen sets:

`b = sort(A_{i1} ∪ A_{i2} ∪ ... ∪ A_{ik})`.

Sereja’s choice is **correct** if the following holds for `b`:

1. `b1 ≤ d`
2. `b_{j+1} - b_j ≤ d` for all `1 ≤ j < |b|`
3. `b_{|b|} ≥ n - d + 1`

Find the minimum possible number of chosen sets `k`.

### Input

- First line: `n, m, d` (`1 ≤ d ≤ n ≤ 10^5`, `1 ≤ m ≤ 20`)
- Next `m` lines:
  - First number is `s_i` — size of set `A_i`
  - Then `s_i` distinct integers in `[1, n]` — elements of `A_i`

It is guaranteed that the sets form a partition of `{1, ..., n}`.

### Output

Print the minimum value `k` such that there is a correct choice.

### Examples

**Input**
```text
3 2 2
1 2
2 1 3
```

**Output**
```text
1
```

**Input**
```text
5 1 1
5 4 5 3 2 1
```

**Output**
```text
1
```

**Input**
```text
7 3 1
4 1 3 5 7
2 2 6
1 4
```

**Output**
```text
3
```

### Thoughts

Let `S` be the union of the chosen sets. The three conditions on the sorted array `b` are equivalent to saying:

- every segment of `d` consecutive values on the number line contains at least one element of `S`
- because the first chosen value must be at most `d`, the segment `[1, d]` must be hit
- because the last chosen value must be at least `n - d + 1`, the segment `[n - d + 1, n]` must be hit

So the problem becomes: choose the minimum number of sets such that every interval `[x, x + d - 1]` contains at least one chosen number.

Since the given sets form a partition of `{1, 2, ..., n}`, each value `v` belongs to exactly one set. Assign each value its set id. Then every length-`d` interval corresponds to a bitmask of the set ids appearing inside that interval.

If we look at the sets that are **not** chosen, then a choice is invalid exactly when some interval is composed only of unchosen sets. In other words, if an interval-mask `mask` is fully contained in the unchosen-set mask `U`, then that interval has no chosen value, so the choice fails.

Example:

- `d = 3`
- `A1 = {1, 6}`
- `A2 = {2, 3}`
- `A3 = {4, 5}`

Suppose we choose only `A1`. Then the unchosen-set mask is `U = {2, 3}`.

Now look at the interval `[2, 4]`, which contains the numbers `{2, 3, 4}`.

- `2` belongs to `A2`
- `3` belongs to `A2`
- `4` belongs to `A3`

So the interval-mask is exactly `{2, 3}`.

Since `{2, 3} ⊆ U`, every number in this interval belongs to an unchosen set. That means `[2, 4]` contains no chosen value at all, so this choice is invalid.

Therefore:

- compute all masks of length-`d` intervals with a sliding window
- mark those masks as forbidden
- run subset DP so that every superset of a forbidden mask is also forbidden
- among all masks `U` that are not forbidden, maximize `popcount(U)`

If `best` is the maximum number of unchosen sets, then the answer is:

`m - best`

The sliding window is `O(n)`, and the subset DP is `O(m * 2^m)`, which is fast enough because `m <= 20`.
