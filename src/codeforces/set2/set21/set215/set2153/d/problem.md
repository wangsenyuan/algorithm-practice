# D. Not Alone (Codeforces 2153D)

**Folder:** `solution.go` / `solution_test.go` — minimum operations to make a circular array “nice” under L1 cost; linear-time DP with three wrap alignments.

**Statement (abridged):** A circular array `b` of length `n` is **nice** if every index `i` has at least one **adjacent** index (including `n` next to `1`) with the same value. Given integer circular array `a`, you may change entries by ±1 per operation. Minimize `Σ |b_i - a_i|`.

Constraints: `3 ≤ n ≤ 2·10^5` per test, `Σ n ≤ 2·10^5`, `1 ≤ a_i ≤ 10^9`.

---

## Structure of nice arrays

On a circle, indices with the same value form **contiguous arcs**. If some value appeared at a **single** index with different values on both sides, that index would have no equal neighbor — forbidden. So every maximal constant run has length **at least 2**.

---

## Key observation: only blocks of length 2 and 3

In an optimal `b`, we can assume the circle is partitioned into blocks of lengths **2 or 3** only:

- Any length `L ≥ 2` can be tiled by segments of size 2 and 3 (every integer `≥ 2` is a non-negative combination of 2 and 3).
- **Cost of a block:** if positions `i..j` must all become one value `t`, the minimum `Σ |a_k - t|` over `t` is achieved at a **median** of the `a` values on that block.
  - Length **2** `{x,y}`: optimal cost is `|x - y|` (any `t` between `x` and `y`).
  - Length **3** `{x,y,z}`: optimal cost is `max(x,y,z) - min(x,y,z)` (median of three numbers).

The implementation uses `calc` for these two cases (and handles `n = 2` or `n = 3` trivially).

---

## From circle to line: three “cut” alignments

If we **cut** the circle between two indices, we get a line `0 .. n-1`. On the line, a valid partition uses only blocks of size 2 or 3 **unless** the cut falls **inside** a block that **wraps** around the end (that block uses a suffix of the line and a prefix).

Such a wrapping block has length at most **3** (only sizes 2 and 3 are allowed). So the cut can lie:

- **between** two blocks (`w = 1` in code — standard linear partition from the start), or
- inside a **2-block** that straddles `(n-1, 0)`, or
- inside a **3-block** that straddles the wrap.

Trying **three** alignments (`play(1)`, `play(2)`, `play(3)` in `solution.go`) is enough: they correspond to reserving `w ∈ {1,2,3}` positions at the end of the index range so the DP state correctly accounts for the initial wrap block (see initial `dp` seeds for each `w`).

---

## DP on one alignment

For a fixed `w`, let `dp[k]` be the minimum cost to cover indices `0..k` consistently with that alignment. The transition tries the last block length `d ∈ {2,3}` ending at the current position:

```text
dp[i+d-1] = min(dp[i+d-1], dp[i-1] + calc(a[i : i+d]))
```

(`calc` = pair or triple cost as above.)

The loop runs for `i` from `1` up to `n - (w - 1) - 1` so that only feasible prefixes are updated; the answer for this alignment is `dp[n - w]`.

**Initial states** differ per `w` so that the first one or two blocks match how the wrap-around block meets index `0`:

- **`w == 1`:** no extra wrap in the seed — `dp[1] = |a[0]-a[1]|`, `dp[2] = calc(a[0:3])`.
- **`w == 2`:** wrap uses `a[n-1]` with `a[0]` (pair) or extends to `a[1]` (triple); see `dp[0]`, `dp[1]`, `dp[2]` in code.
- **`w == 3`:** wrap uses three indices around the seam `n-2, n-1, 0` or combinations with the start of the array; again encoded in `dp[0]`, `dp[1]`, `dp[2]`.

The final answer is `min(play(1), play(2), play(3))`.

---

## Complexity

`O(n)` time and `O(n)` memory per test case (constant number of DP passes).

---

## Reference

[Problem 2153D on Codeforces](https://codeforces.com/problemset/problem/2153/D)
