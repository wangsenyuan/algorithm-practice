# G. Wafu!

To help improve her math, Kudryavka is given a set `S` consisting of `n` distinct positive integers.

Initially, her score is `1`. While `S` is not empty, she may perform the following operation:

1. Let `m` be the minimum value in `S`.
2. Multiply her score by `m`.
3. Remove `m` from `S`.
4. For every integer `i` such that `1 <= i < m`, add `i` to `S`.
   It can be shown that this step adds no duplicates.

After exactly `k` operations, help her determine her score modulo `10^9 + 7`.

## Input

Each test contains multiple test cases.
The first line contains the number of test cases `t` (`1 <= t <= 10^4`).

For each test case:

- The first line contains two integers `n` and `k` (`1 <= n <= 2 * 10^5`, `1 <= k <= 10^9`).
- The second line contains `n` integers `s_1, s_2, ..., s_n`
  (`1 <= s_i <= 10^9`, `s_i != s_j`) — the elements of the initial set `S`.

It is guaranteed that `S` is non-empty before each of the `k` operations.

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output one integer — the answer modulo `10^9 + 7`.

## Example

**Input**

```text
4
2 3
1 3
3 6
5 1 4
2 100
2 100
5 15
1 2 3 4 5
```

**Output**

```text
3
24
118143737
576
```

## Note

In the first test case, the process is:

`{1,3} -> {3} -> {1,2} -> {2}`

- Remove `1`, then remove `3` (after adding `1,2`), then remove `1`.
- So the removed values are `1, 3, 1`, and the score is `1 * 3 * 1 = 3`.

In the second test case, the score is `1 * 4 * 1 * 2 * 1 * 3 = 24`.

## Solution Summary (from `solution.go`)

1. Sort the set and shift each value by `-1`.
   - The code works with `v = s_i - 1`, because deleting `s_i` creates the canonical block `1..v`.

2. Precompute `dp[j]` for `j = 0..30`.
   - `dp[j]` is the product contribution of fully consuming the block `1..(j+1)`.
   - Recurrence in code:
     - `dp[j] = (j + 1) * product(dp[0..j-1]) (mod M)`, `M = 1e9+7`.
   - Since `k <= 1e9`, only powers up to `2^30` matter.

3. Key operation count fact.
   - Fully consuming block `1..v` takes `2^v` operations.
   - So in the main loop, if `k >= 2^v` (and `v <= 30`), multiply answer by `dp[v]` and do `k -= 2^v`.

4. First unfinished block is handled by divide-and-conquer `f(v, k)`.
   - `f(v, k)` returns the contribution of the next `k` operations starting from minimum value `v+1`.
   - Base: `k == 0 => 1`.
   - Take current minimum once: multiply by `v+1`, then consume the remaining `k-1` using block decomposition by powers of two.
   - For each level `j` (small to large):
     - if `k >= 2^j`, take whole block via `dp[j]`;
     - otherwise recurse into `f(j, k)` and stop.

5. Why this is fast.
   - Main loop processes sorted original elements once.
   - Each full block subtraction removes a power-of-two chunk.
   - Recursive depth is bounded by ~31.
   - Overall per test case: `O(n log n + 31^2)` (dominated by sorting).