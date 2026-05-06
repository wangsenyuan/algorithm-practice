# D. Sum of LDS

You are given a permutation `p_1, …, p_n` such that for all `1 ≤ i ≤ n - 2`:

`max(p_i, p_{i+1}) > p_{i+2}`.

For each pair `1 ≤ l ≤ r ≤ n`, consider the subarray `p_l, p_{l+1}, …, p_r` and let **LDS** be the length of its **longest decreasing subsequence**. Compute the sum of LDS over **all** `(l, r)`.

> **\*** A **permutation** of length `n` is an array of `n` distinct integers from `1` to `n` in some order. For example, `[2, 3, 1, 5, 4]` is a permutation; `[1, 2, 2]` is not (duplicate `2`); `[1, 3, 4]` is not a permutation of length `3` (value `4` is out of range).

> **†** For an array `b`, a **decreasing subsequence** of length `k` is indices `i_1 < i_2 < … < i_k` such that `b_{i_1} > b_{i_2} > … > b_{i_k}`.

## Input

Each test contains multiple test cases. The first line contains `t` (`1 ≤ t ≤ 10^4`).

For each test case:

- One line: `n` (`3 ≤ n ≤ 5 · 10^5`).
- One line: `n` distinct integers `p_1, …, p_n` (`1 ≤ p_i ≤ n`).

It is guaranteed that `max(p_i, p_{i+1}) > p_{i+2}` for all `1 ≤ i ≤ n - 2`.

The sum of `n` over all test cases does not exceed `5 · 10^5`.

## Output

For each test case, print one integer: the sum over all subarrays of the LDS length of that subarray.

## Example

**Input**

```text
4
3
3 2 1
4
4 3 1 2
6
6 1 5 2 4 3
3
2 3 1
```

**Output**

```text
10
17
40
8
```

## Note

For any array `a`, write `LDS(a)` for the length of its longest decreasing subsequence.

In the first test case, every subarray is (strictly) decreasing.

In the second test case:

- `LDS([4]) = LDS([3]) = LDS([1]) = LDS([2]) = 1`
- `LDS([4, 3]) = LDS([3, 1]) = 2`, `LDS([1, 2]) = 1`
- `LDS([4, 3, 1]) = 3`, `LDS([3, 1, 2]) = 2`
- `LDS([4, 3, 1, 2]) = 3`

So the answer is `1 + 1 + 1 + 1 + 2 + 2 + 1 + 3 + 2 + 3 = 17`.


### ideas
1. L[j] 表示 p[i] > p[i+1] .. > p[j]
2. p[i-1] < p[i], 那么 p[i-2] > p[i] must hold

## Solution explanation

We sum by the right endpoint.

For a fixed right endpoint `r`, define:

`dp[r] = sum LDS(p[l..r])` over all `1 <= l <= r`.

The final answer is:

`sum dp[r]`.

The code does not store the whole `dp` array. It keeps only the current value `dp`, and accumulates it into `res`.

## What the condition gives us

The condition is:

`max(p_i, p_{i+1}) > p_{i+2}`.

Equivalently, there cannot be three consecutive increasing elements:

`p_i < p_{i+1} < p_{i+2}`.

So whenever we see an adjacent increase `p_{i-1} < p_i`, and `i >= 2`, the previous element two positions back must satisfy:

`p_{i-2} > p_i`.

This local restriction is the whole reason the LDS can be updated by looking only at the last adjacent pair.

## Key lemma

For every subarray ending at `i`, compared with the same subarray ending at `i-1`:

- if `p_{i-1} > p_i`, the LDS increases by exactly `1`;
- if `p_{i-1} < p_i`, the LDS stays the same.

In formula form, for every `l <= i - 1`:

```text
if p[i-1] > p[i]:
    LDS(l, i) = LDS(l, i-1) + 1
else:
    LDS(l, i) = LDS(l, i-1)
```

There is also one new subarray `[i, i]`, whose LDS is `1`.

## Why a descent increases every old subarray by 1

Consider `p_{i-1} > p_i`.

For any fixed left endpoint `l`, under the problem condition there is a longest decreasing subsequence of `p[l..i-1]` that ends at position `i-1`. Then we can append `p_i` to it, because:

`p_{i-1} > p_i`.

So:

`LDS(l, i) >= LDS(l, i-1) + 1`.

Adding one new element can increase an LDS by at most `1`, therefore:

`LDS(l, i) = LDS(l, i-1) + 1`.

Intuitively, a descending adjacent pair means the new last element can extend the best decreasing subsequence for every previous left endpoint.

## Why an ascent does not increase old subarrays

Consider `p_{i-1} < p_i`.

The new element `p_i` cannot be appended after `p_{i-1}`. Also, because three consecutive increases are forbidden, if `i >= 2` then:

`p_{i-2} > p_i`.

So `p_i` is not creating a strictly better tail than what was already available before it. Any decreasing subsequence using `p_i` can be matched by a decreasing subsequence of the same length inside `p[l..i-1]`.

Therefore adding `p_i` does not increase the LDS of any old subarray:

`LDS(l, i) = LDS(l, i-1)`.

For the boundary case `i = l`, the new singleton is handled separately.

## Deriving the code transition

Assume we already know:

`dp = sum LDS(p[l..i-1])` for all `l <= i-1`.

When moving to position `i`, there are `i` old subarrays if we use zero-based indexing:

```text
[0..i-1], [1..i-1], ..., [i-1..i-1]
```

There is also one new singleton subarray `[i..i]`.

### Case 1: `p[i-1] > p[i]`

Every old subarray gains `1`, and there are `i` old subarrays.

So:

```text
new_dp = old_dp + i + 1
```

The `+i` is from increasing all old subarrays by `1`.

The `+1` is from the new singleton `[i..i]`.

This is exactly:

```go
if p[i-1] > p[i] {
	dp += i
}
dp++
```

### Case 2: `p[i-1] < p[i]`

No old subarray changes. Only the new singleton contributes `1`.

So:

```text
new_dp = old_dp + 1
```

This is also covered by:

```go
dp++
```

because the `if p[i-1] > p[i]` block is skipped.

## Meaning of `res`

At the start:

```go
res := 1
dp := 1
```

This corresponds to the only subarray ending at index `0`:

`[0..0]`, whose LDS is `1`.

Then for every later index `i`, the code updates `dp` to become the sum of LDS values over all subarrays ending at `i`, and adds it to `res`:

```go
res += dp
```

So after the loop, `res` is the sum over all right endpoints, which is exactly the required answer.

## Example

For `p = [4, 3, 1, 2]`:

- `i = 0`: subarrays ending here: `[4]`, sum is `1`, so `dp = 1`, `res = 1`.
- `i = 1`: `4 > 3`, old subarrays gain `1`; `dp = 1 + 1 + 1 = 3`. Contributions are `[4,3] = 2`, `[3] = 1`.
- `i = 2`: `3 > 1`, old subarrays gain `1`; `dp = 3 + 2 + 1 = 6`. Contributions are `[4,3,1] = 3`, `[3,1] = 2`, `[1] = 1`.
- `i = 3`: `1 < 2`, old subarrays do not change; `dp = 6 + 1 = 7`. Contributions are `[4,3,1,2] = 3`, `[3,1,2] = 2`, `[1,2] = 1`, `[2] = 1`.

Total:

`1 + 3 + 6 + 7 = 17`.

## Complexity

The solution scans the permutation once.

Time complexity: `O(n)` per test case.

Memory complexity: `O(1)` besides the input array.
