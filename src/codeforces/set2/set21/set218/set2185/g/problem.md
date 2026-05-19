# G. Mixing MEXes (Codeforces 2185G)

**Limits:** 2 seconds per test, 256 MB  
**I/O:** standard input / standard output

Source: [https://codeforces.com/problemset/problem/2185/G](https://codeforces.com/problemset/problem/2185/G)

## Problem Statement

You are given `n` arrays `a[1], a[2], ..., a[n]`.

Exactly one move is performed:

1. Choose a source array `a[i]`.
2. Choose one element of that array, say its `j`-th element `a[i][j]`.
3. Choose a different destination array `a[k]` where `k != i`.
4. Move that selected element from `a[i]` to the back of `a[k]`.

The value of such a move is the sum of the MEX values of all arrays after the
move is finished.

You must evaluate the sum of the values over all distinct moves. Two moves are
different if their ordered triples `(i, j, k)` are different.

For an array `b`, `MEX(b)` is the smallest non-negative integer not present in
`b`.

Examples:

- `MEX([1, 2, 0, 5]) = 3`;
- `MEX([1, 2, 4, 9]) = 0`.

## Solution summary (from `solution.go`)

### Key observation

A move is a triple `(i, j, k)`: remove `a[i][j]` from array `i`, append it to array `k` (`k ≠ i`).

Only **`MEX(a[i])`** and **`MEX(a[k])`** can change; every other array keeps its MEX. So the value of one move is:

`sum_after = S - mex1[i] - mex1[k] + new_mex(i) + new_mex(k)`

where `S = Σ_r mex1[r]` is the total MEX before any move, and `mex1[r]` is the initial MEX of array `r` (after sorting each array).

### Preprocessing per test case

Let `m = max_i |a[i]|`.

1. **Initial MEX.** For each array `r`, sort `a[r]`, scan to get `mex1[r]`, and set `S += mex1[r]`.

2. **Extended MEX / `dp`.** While scanning, also compute how far MEX could grow if we kept adding the next missing integers at the end of the sorted list (the “second segment” after the first gap). Store  
   `dp[mex1[r]] += mex_extended - mex1[r]`.  
   Intuitively, `dp[v]` collects how much **destination** `k` gains when its current MEX is exactly `v` and we append a value that **fills** that hole (so MEX jumps forward by `dp[v]`).

### Classifying a removed value `v = a[i][j]`

Fix source array `i` and element `v`. There are `n - 1` choices of destination `k`. The code adds the total contribution over all those moves without iterating every `k` explicitly.

**Case A — `v > mex1[i]`**  
Removing `v` does not lower `MEX(a[i])`; source side stays `mex1[i]`.

For the destination, appending `v` only matters when `v` equals the current MEX of `k` (otherwise `MEX(a[k])` unchanged). Summing over all `k` gives a uniform part `S · (n - 1)` plus an extra `dp[v]` when `v` is a valid index.

**Case B — `v < mex1[i]`** (note: `v ≠ mex1[i]` because `mex1[i]` is missing from `a[i]`)

Use a temporary frequency map on values in `a[i]` while processing index `i`:

- **Duplicate / another copy:** if `freq[v] > 1`, or `v` appears twice consecutively in the sorted array, removing one copy does **not** change `MEX(a[i])`. Contribution pattern like Case A: `S · (n - 1) + dp[v]`.

- **Unique at this level:** removing `v` drops `MEX(a[i])` from `mex1[i]` down to `v`. Source contribution becomes `v` instead of `mex1[i]`, so use  
  `(S - mex1[i] + v) · (n - 1) + dp[v]`  
  for the aggregate over destinations (again with the same `dp[v]` correction when the appended value closes the MEX gap at `k`).

After handling all elements of `a[i]`, clear `freq` for the next `i`.

### Answer

`ans` is the sum over all `(i, j)` of the formulas above; each already counts the `n - 1` destinations. Return `ans`.

### Complexity

- Sorting all array elements: `O(Σ l[i] log l[i])`.
- Second pass over all elements with `freq`: `O(Σ l[i])`.
- Memory: `O(n + m)` for `mex1`, `dp`, and `freq`.

### Ideas (original notes)

1. `(i, j, k)` only changes `mex(a[i])` and `mex(a[k])`.
2. If removing `a[i][j]` does not touch the current MEX gap, `a[i]` is unchanged.
3. After moving to `k`, if `mex(a[k]) > a[i][j]`, destination MEX is often unchanged.
4. If `mex(a[k]) = a[i][j]`, we need how far MEX can increase — handled by `dp`.
5. Aggregate per value `v` (up to array length scale) instead of per triple.