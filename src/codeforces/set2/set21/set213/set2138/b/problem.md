# B. Antiamuny Wants to Learn Swap

For an array `b` of length `m`, you may perform the following two operations:

1. Select an index `1 ≤ i ≤ m - 1`. Then, swap the values of `b_i` and `b_{i+1}`.
2. Select an index `1 ≤ i ≤ m - 2`. Then, swap the values of `b_i` and `b_{i+2}`.

However, you can only perform operation 2 at most once.

We define `f(b)` as the minimum number of operations (using both operation 1 and operation 2) required to sort array `b` in non-decreasing order, and `g(b)` as the minimum number of operations required to sort array `b` in non-decreasing order using only operation 1.

The array `b` is **perfect** if `f(b) = g(b)`. In other words, the ability to use operation 2 does not reduce the number of operations required to sort array `b` compared to using only adjacent swaps.

You are given a permutation `a` of length `n`, and must answer `q` queries. Each query consists of two integers `l` and `r` (`1 ≤ l ≤ r ≤ n`), representing the subarray `a[l … r]`. For each query, determine whether the subarray `a[l … r]` is perfect.

> **\*** A permutation of length `n` is an array consisting of `n` distinct integers from `1` to `n` in arbitrary order. For example, `[2, 3, 1, 5, 4]` is a permutation, but `[1, 2, 2]` is not (2 appears twice), and `[1, 3, 4]` is also not a permutation of length 3 (4 is out of range).

> **†** The subarray `a[l … r]` includes all elements `a_l, a_{l+1}, …, a_r`.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 ≤ t ≤ 5 · 10^4`). The description of the test cases follows.

The first line of each test case contains two integers `n`, `q` (`1 ≤ n, q ≤ 5 · 10^5`) — the length of array `a` and the number of queries.

The second line of each test case contains `n` integers `a_1, a_2, …, a_n` (`1 ≤ a_i ≤ n`) — the elements in permutation `a`.

Each of the next `q` lines contains two integers `l` and `r` (`1 ≤ l ≤ r ≤ n`) — the left and right endpoints of the queried subarray.

It is guaranteed that the sum of `n` and the sum of `q` over all test cases do not exceed `5 · 10^5`.

## Output

For each test case, output `q` lines. For each query, print `"YES"` if the queried subarray `a[l … r]` is perfect, and `"NO"` otherwise.

You can output the answer in any case (upper or lower). For example, the strings `"yEs"`, `"yes"`, `"Yes"`, and `"YES"` will be recognized as positive responses.

## Example

### Input

```text
2
5 5
1 5 4 3 2
1 2
1 5
3 5
1 4
2 5
5 5
3 2 1 4 5
1 1
4 5
1 4
2 5
3 4
```

### Output

```text
YES
NO
NO
NO
NO
YES
YES
NO
YES
YES
```

## Note

In the first test case:

- Query 1: `a[1 … 2] = [1, 5]` is already sorted in increasing order. Thus, `f(a[1 … 2]) = g(a[1 … 2]) = 0` and the subarray `a[1 … 2]` is perfect.
- Query 2: `a[1 … 5] = [1, 5, 4, 3, 2]`. `f(a[1 … 5]) = 4` as we can sort the array using the following sequence of operations:

  `[1, 5, 4, 3, 2] → op2 → [1, 3, 4, 5, 2] → op1 → [1, 3, 4, 2, 5] → op1 → [1, 3, 2, 4, 5] → op1 → [1, 2, 3, 4, 5]`

  On the other hand, `g(a[1 … 5]) = 6` as it requires at least 6 adjacent swaps to sort `a[1 … 5]`. Since `f(a[1 … 5]) ≠ g(a[1 … 5])`, the subarray `a[1 … 5]` is not perfect.
- Query 3: `a[3 … 5] = [4, 3, 2]`. `f(a[3 … 5]) = 1` as we can sort the array using the following sequence of operations:

  `[4, 3, 2] → op2 → [2, 3, 4]`

  On the other hand, `g(a[3 … 5]) = 3` as it requires at least 3 adjacent swaps to sort `a[3 … 5]`. Since `f(a[3 … 5]) ≠ g(a[3 … 5])`, the subarray `a[3 … 5]` is not perfect.

## Solution

For each query `[l … r]`, decide whether the subarray is **perfect**: using at most
one distance-2 swap does not beat adjacent swaps alone (`f(b) = g(b)`).

### Short recap

- Operation 1: adjacent swap (any number of times).
- Operation 2: swap positions `i` and `i + 2` (at most once).
- `g(b)` = minimum adjacent swaps to sort (inversion count for a permutation).
- `f(b)` = minimum operations when one distance-2 swap is allowed.
- Output `YES` iff `f(b) = g(b)` for the queried subarray.

### Key observations

1. **When does operation 2 help?** If the subarray contains three indices
   `i < j < k` with values `x > y > z` (a decreasing triple in value order), one
   distance-2 swap can replace adjacent swaps that would otherwise fix that
   pattern. Example: `[4, 3, 2]` needs 3 adjacent swaps but one op2 at the middle
   yields `[2, 3, 4]`.

2. **Perfect ⟺ no decreasing triple inside the subarray.** The subarray is perfect
   exactly when it has **no** indices `p < q < r` (all in `[l … r]`) with
   `a[p] > a[q] > a[r]`. Equivalently: no value pattern `x > y > z` appearing
   left-to-right within the range.

3. **Offline queries by right endpoint.** Bucket each query `(l, r)` into
   `todo[r - 1]`. Scan the array once left-to-right; when index `i` is processed,
   answer every query whose right endpoint is `i + 1`.

4. **Prefix witness `last`.** While scanning, maintain `last`: the maximum left
   index `p` among all decreasing triples `a[p] > a[q] > a[r]` with
   `p < q < r ≤ i` seen in prefix `[0 … i]`. For a query `(l, r)` handled at
   index `i = r - 1`:
   - `last >= l` → some triple lies entirely in `[l … r]` → **NO**
   - `last < l` → no triple is fully contained → **YES**

5. **Why `last := -inf`?** Before any element is seen there is no triple.
   `last` must be strictly less than every valid query left bound `l` (0-indexed
   `l - 1 ≥ 0`). Initializing to `-inf` (i.e. `-inf` where `inf = 1 << 60` in
   code) makes `last < l` true for all early prefixes. Segment trees are also
   seeded with `-inf` meaning “no prior index recorded”.

### Data structures (two segment trees on value domain)

Values are reindexed to `0 … n - 1` (`a[i]--` in the loop). Both trees have size
`n` over value indices.

| Tree | `Update(v, x)` | `Get(L, R)` |
| --- | --- | --- |
| `s1` | latest index where value `v` appeared | max stored index in `[L, R)` |
| `s2` | auxiliary chain pointer at value `v` | max stored value in `[L, R)` |

At index `i` with value `v`:

```text
w  = s1.Get(v, n)           // rightmost index < i with value > v
s2.Update(v, w)             // link v to that larger predecessor
last = max(last, s2.Get(v+1, n))   // extend witness using values > v
s1.Update(v, i)             // record current position of v
```

Intuition for one step: if some larger value `u > v` was seen earlier at index
`j`, then `s2[u]` stores the latest index before `j` where a value `> u` was seen.
Taking the max over all `u > v` yields the leftmost index of a triple
`(x, u, v)` with `x > u > v` ending at the current position.

### Algorithm (matching `solve`)

**Step 0 — bucket queries.**

```go
todo[r] = append(todo[r], query_id)   // r = r_1based - 1
```

**Step 1 — initialize.**

```go
last := -inf
s1, s2 := all positions set to -inf
```

**Step 2 — left-to-right scan** (`i = 0 … n - 1`):

1. `v := a[i] - 1`
2. `w := s1.Get(v, n)` — latest index before `i` with value strictly greater than `v`
3. `s2.Update(v, w)`
4. `last = max(last, s2.Get(v+1, n))` — if positive, a triple ending at `i` exists
5. `s1.Update(v, i)`
6. For each query `id` in `todo[i]` with left bound `l := l_1based - 1`:
   - `ans[id] = (last < l)` → `YES` when true

**Step 3 — return** `ans` in input order.

### Why it works

Consider the moment index `i` with value `v` is processed. Any decreasing triple
ending at `i` must look like `x > u > v` at positions `p < q < i`. The middle
value `u` satisfies `u > v`; when `u` was processed, `s2[u]` was set to the
latest index before `q` holding a value `> u`, namely `p`. Thus
`s2.Get(v+1, n) = max_{u > v} s2[u]` is exactly the maximum `p` over all such
triples ending at `i`. Updating `last = max(last, …)` keeps the strongest witness
seen anywhere in the prefix.

For query `[l … r]`, a triple lies **inside** the subarray iff all three indices
are ≥ `l`, which happens exactly when its left index `p` satisfies `p ≥ l`. The
scan keeps `last = max p` over triples in the prefix ending at `r`. If
`last < l`, every triple has `p < l` and cannot be fully contained; if
`last ≥ l`, some triple lies entirely in `[l … r]` → not perfect.

### Walkthrough: `a = [1, 5, 4, 3, 2]` (sample 1)

0-indexed values after `--`: `[0, 4, 3, 2, 1]`.

| `i` | `v` | `w` | `last` after step | note |
| --- | --- | --- | --- | --- |
| 0 | 0 | `-inf` | `-inf` | |
| 1 | 4 | `-inf` | `-inf` | |
| 2 | 3 | 1 | `-inf` | saw 4 at index 1 |
| 3 | 2 | 2 | 1 | triple (5,4,3) at indices 1,2,3 |
| 4 | 1 | 2 | 2 | triple (4,3,2) at indices 2,3,4 |

Queries at `i = 4` (`r = 5`): `last = 2`. For `(l, r) = (1, 5)` → `l = 0`,
`last ≥ l` → **NO**. For `(3, 5)` → `l = 2`, `last ≥ l` → **NO**. For
`(1, 2)` at `i = 1`, `last = -inf < 0` → **YES**.

### Walkthrough: `a = [3, 2, 4, 1]` (test 3)

Values: `[2, 1, 3, 0]`. At `i = 3`, `v = 0`, `w = 2`, `s2[1] = 0` from earlier,
so `last = 0`. Query `[1 … 4]` has `l = 0`, `last ≥ l` → **NO** (triple
`3 > 2 > 1` at indices 0, 1, 3).

### Complexity

- **Time:** `O((n + q) log n)` per test case — one scan, each step does `O(log n)`
  segment-tree updates/queries; answering queries at each index is `O(1)` amortized
  over bucket lists.
- **Space:** `O(n + q)` for trees, buckets, and answers.

### ideas

1. `g(b)` is the inversion count (minimum adjacent swaps).
2. A single op2 can save work only when some `x > y > z` appears in order in the
   subarray.
3. Perfect subarray ⟺ no decreasing triple fully inside `[l … r]`.
4. Scan left-to-right; `last` is the witness left index; answer `YES` iff
   `last < l` at position `r - 1`.