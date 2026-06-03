# E - Subarray Sum Divisibility (ABC419)

**Contest:** [ABC419](https://atcoder.jp/contests/abc419) — AtCoder Beginner Contest 419  
**Task:** [https://atcoder.jp/contests/abc419/tasks/abc419_e](https://atcoder.jp/contests/abc419/tasks/abc419_e)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 475 points

## Problem Statement

You are given a length-`N` integer sequence `A = (A_1, A_2, ..., A_N)`.

Your goal is to perform the following operation repeatedly so that for every
length-`L` contiguous subarray of `A`, the sum is a multiple of `M`.

- Choose an integer `i` such that `1 <= i <= N`, and increase the value of `A_i`
  by `1`.

Find the minimum possible number of operations before achieving the goal.

## Constraints

- `1 <= N, M <= 500`
- `1 <= L <= N`
- `0 <= A_i < M`
- All input values are integers.

## Solution

Every length-`L` window must have total sum `0 (mod M)`. The key is that windows
of length `L` overlap in a very rigid way.

### Same residue class must share the same target mod `M`

Subtract the condition for window `[s, s+L-1]` from window `[s+1, s+L]`:

```text
(A[s] + x[s]) ≡ (A[s+L] + x[s+L])   (mod M)
```

So for all valid `s`, positions `s` and `s+L` must end with the same remainder mod
`M`. By chaining, every index with the same value of `pos % L` must finish with
the same remainder mod `M`.

Call that common remainder `t_r` for class `r = pos % L`.

### Cost for one class

If class `r` uses target remainder `j`, every index in that class pays the cost
to increment from `A[pos]` up to some value with remainder `j (mod M)`.

For one element with current remainder `w = A[pos] % M` (`w = M` when `A[pos] = 0`):

- if `j >= w`: cost is `j - w`
- otherwise: cost is `M - w + j`

Summing over all positions in class `r` gives `dp[r][j]`.

### All windows share one sum condition

A length-`L` window contains exactly one index from each class
`0, 1, ..., L-1`. Therefore every window sum mod `M` equals:

```text
t_0 + t_1 + ... + t_{L-1}   (mod M)
```

So the classes are independent in cost, but their chosen remainders must satisfy:

```text
t_0 + t_1 + ... + t_{L-1} ≡ 0   (mod M)
```

### Final DP

Run a knapsack-style DP over the `L` classes:

- `fp[s]` = minimum total cost after some prefix of classes, with remainder sum `s`
- Transition: pick `j` for the next class and update `fp[(s+j) % M]`

Start with `fp[0] = 0`, process all `L` classes, and return `fp[0]`.

### How the code maps to this

1. Build `dp[r][j]` by looping all positions and adding each element's increment
   cost into its class row.
2. Maintain `fp` / `nfp` arrays of size `M`.
3. For each class `i`, combine `dp[i][j]` with the previous `fp` table.
4. Output `fp[0]`.

### Complexity

- Building `dp`: `O(N * M)`
- Remainder-sum DP: `O(L * M^2)`
- Space: `O(L * M + M)`

With `N, M <= 500`, this fits easily.
