# E - Alphabet Tiles (ABC358)

**Contest:** [ABC358](https://atcoder.jp/contests/abc358) — AtCoder Beginner Contest 358  
**Task:** [https://atcoder.jp/contests/abc358/tasks/abc358_e](https://atcoder.jp/contests/abc358/tasks/abc358_e)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 475 points

## Problem Statement

AtCoder Land sells tiles with English letters written on them. Takahashi is thinking of making a nameplate by arranging these tiles in a row.

Find the number, modulo 998244353, of strings consisting of uppercase English letters with a length between 1 and K, inclusive, that satisfy the following conditions:

For every integer i satisfying 1 ≤ i ≤ 26, the following holds:

- Let a_i be the i-th uppercase English letter in lexicographical order. For example, a_1 = `A`, a_5 = `E`, a_26 = `Z`.
- The number of occurrences of a_i in the string is between 0 and C_i, inclusive.

## Constraints

- 1 ≤ K ≤ 1000
- 0 ≤ C_i ≤ 1000
- All input values are integers.

## Solution

Count every uppercase string whose length is between 1 and K (inclusive) and
whose i-th letter appears at most C_i times. Letters with C_i = 0 never appear.
The answer is taken modulo 998244353.

### Key idea

Process the 26 letters in lexicographic order (A, B, …, Z). Maintain a DP over
**total string length**:

```text
dp[d] = number of distinct strings of length exactly d
        using only letters processed so far, respecting their caps
```

Start with `dp[0] = 1` (the empty string). When introducing a letter that may be
used at most `x` times, every existing string of length `d` with count `v` can be
extended by inserting `i` copies of the new letter (`0 ≤ i ≤ x`, `i + d ≤ K`).
Those `i` new symbols can be placed among the `d + i` positions in
`C(d + i, i)` ways, so:

```text
fp[d + i] += v * C(d + i, i)   (mod 998244353)
```

After all 26 letters, sum `dp[1] + … + dp[K]` — lengths 1 through K, excluding
the empty string at `dp[0]`.

The binomial coefficient is the standard “multiset permutation” factor: if a
string already uses certain counts of earlier letters, choosing where the new
letter goes among all positions counts exactly `C(d + i, i)`.

### Precomputation (matching `init()` and `nCr`)

`K` and each `C_i` are at most 1000, so factorials up to about 2000 suffice.
The code precomputes in `init()`:

| Symbol      | Meaning                                                                      |
| ----------- | ---------------------------------------------------------------------------- |
| `F[i]`      | `i!` mod 998244353                                                           |
| `I[i]`      | `(i!)^{-1}` mod 998244353 via Fermat: `pow(F[X-1], mod-2)` and backward pass |
| `nCr(n, r)` | `F[n] * I[r] * I[n-r]`; returns 0 if `r < 0` or `n < r`                      |

`X = 1010` covers all binomial arguments encountered in the DP loops.

### Algorithm (matching `solve`)

**Step 0 — initialize.**

```go
dp := make([]int, k+1)
fp := make([]int, k+1)
dp[0] = 1
```

**Step 1 — iterate over letter caps.** For each `x` in `c` (26 values):

1. Scan `d = 0 … k`. Skip if `dp[d] == 0`.
2. For `i = 0 … x` with `i + d ≤ k`, update the next layer:

   ```go
   fp[i+d] = add(fp[i+d], mul(v, nCr(i+d, i)))
   ```

3. Copy `fp` into `dp`, then `clear(fp)` for the next letter.

Note: the inner loop uses `nCr(i+d, i)` — in code the loop variable `i` is the
**count of the current letter**, and `d` is the length built from previous
letters. The binomial index is `(i+d, i)`.

**Step 2 — aggregate.** Sum non-empty lengths:

```go
for d := 1; d <= k; d++ {
    ans = add(ans, dp[d])
}
```

**Step 3 — modular arithmetic.** `add` and `mul` keep all values in
`[0, mod)` with `mod = 998244353`.

### Why it works

1. **Correctness by letter order.** After processing the first `t` letter types,
   `dp[d]` counts exactly the strings of length `d` whose per-letter usage respects
   `C_1, …, C_t`. The empty string (`dp[0] = 1`) is the unique base case.

2. **Transition is a bijection.** Fix a string counted in `dp[d]` before adding
   letter `a_t`. Appending `i` copies of `a_t` and choosing their positions among
   `d + i` slots yields a unique longer string; every valid extension arises once.
   Multiplying by `C(d + i, i)` counts those interleavings.

3. **Caps are enforced** by restricting `i ≤ x` for the current letter and by
   never exceeding total length `K` via `i + d ≤ k`.

4. **Empty string excluded** deliberately: only `d ≥ 1` enter the final sum, as
   required by “length between 1 and K”.

### Walkthrough: Sample 1 (`K = 2`, `C = [2, 1, 1, 0, …]`)

Only A, B, C may appear (caps 2, 1, 1; rest zero).

**After letter A** (`x = 2`): from `dp[0] = 1`, choose `i ∈ {0, 1, 2}` →
`dp = [1, 1, 1]` (lengths 0, 1, 2 each have one A-only string).

**After letter B** (`x = 1`):

| From `dp[d]` | `i` | Contribution to `fp[d+i]` |
| ------------ | --- | ------------------------- |
| `d=0, v=1`   | 0   | `fp[0] += 1`              |
| `d=0, v=1`   | 1   | `fp[1] += C(1,1)=1`       |
| `d=1, v=1`   | 0   | `fp[1] += 1`              |
| `d=1, v=1`   | 1   | `fp[2] += C(2,1)=2`       |
| `d=2, v=1`   | 0   | `fp[2] += 1`              |

Result: `dp = [1, 2, 3]` — e.g. length 1 has `A, B`; length 2 has `AA, AB, BA`.

**After letter C** (`x = 1`): same transition → `dp = [1, 3, 7]`.

**Answer:** `dp[1] + dp[2] = 3 + 7 = 10`, matching the sample list
(`A, B, C, AA, AB, AC, BA, BC, CA, CB`).

### Complexity

- **Time:** `O(26 · K · max(C_i))` — 26 letters, each scans `dp[0…K]` and up to
  `C_i + 1` choices for `i`. With `K, C_i ≤ 1000`, this is well within limits.
- **Space:** `O(K)` for the two DP arrays plus `O(X)` for factorial tables.

### ideas
