# Expected Inversions After Random Reversals

You are given a permutation of `n` numbers `p1, p2, …, pn`. Perform `k` operations; in each operation, pick two indices `l` and `r` uniformly at random with `l ≤ r`, then reverse the subarray `pl, pl+1, …, pr`. Compute the expected number of inversions after all operations.

## Input
- The first line contains integers `n` and `k` (`1 ≤ n ≤ 100`, `1 ≤ k ≤ 10^9`).
- The second line contains `n` distinct integers `p1, p2, …, pn` forming a permutation of `[1, n]`.

## Subproblems
- **G1 (3 pts):** `1 ≤ n ≤ 6`, `1 ≤ k ≤ 4`
- **G2 (5 pts):** `1 ≤ n ≤ 30`, `1 ≤ k ≤ 200`
- **G3 (16 pts):** `1 ≤ n ≤ 100`, `1 ≤ k ≤ 10^9`

## Output
Print the expected number of inversions after `k` operations with absolute or relative error at most `1e-9`.

## Examples

```
Input
3 1
1 2 3
Output
0.833333333333333
```

```
Input
3 4
1 3 2
Output
1.458333333333334
```

## Note
In the first sample, reversing a single-element interval leaves the permutation unchanged. Reversing the first or last two elements produces one inversion, and reversing the whole array produces three inversions. Taking expectations over all equally likely intervals yields `0.833333333333333`.

