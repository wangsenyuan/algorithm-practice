# 3957. Maximum Sum of M Non-Overlapping Subarrays II

[Problem link](https://leetcode.cn/problems/maximum-sum-of-m-non-overlapping-subarrays-ii/description/)

Difficulty: Hard

You are given an integer array `nums` of length `n`, and three integers `m`,
`l`, and `r`.

Select at least one and at most `m` non-overlapping subarrays from `nums` such
that:

- each selected subarray has length between `l` and `r`, inclusive;
- the total sum of all selected subarrays is maximized.

Return the maximum total sum.

A subarray is a contiguous non-empty sequence of elements within an array.

## Function signature

```go
func maximumSum(nums []int, m int, l int, r int) int64
```

## Examples

### Example 1

```text
Input: nums = [4,1,-5,2], m = 2, l = 1, r = 3
Output: 7
```

One optimal strategy is to select `[4, 1]` with sum `5`, and `[2]` with sum
`2`. The total is `7`.

### Example 2

```text
Input: nums = [1,0,3,4], m = 2, l = 1, r = 2
Output: 8
```

One optimal strategy is to select `[1]` with sum `1`, and `[3, 4]` with sum
`7`. The total is `8`.

### Example 3

```text
Input: nums = [-1,7,-4], m = 1, l = 2, r = 3
Output: 6
```

Select `[-1, 7]`, whose sum is `6`.

### Example 4

```text
Input: nums = [-3,-4,-1], m = 2, l = 1, r = 2
Output: -1
```

All subarray sums are negative, so the best choice is `[-1]`.

## Constraints

- `1 <= n == nums.length <= 10^5`
- `-10^5 <= nums[i] <= 10^5`
- `1 <= m <= n`
- `1 <= l <= r <= n`

## Detailed solution explanation

Let `pref[i]` be the sum of `nums[0:i]`. If a selected subarray starts at
position `j` and ends at position `i - 1`, its sum is:

```text
pref[i] - pref[j]
```

The length constraint `l <= length <= r` means:

```text
l <= i - j <= r
```

So when the right endpoint is `i`, the valid start positions are:

```text
i - r <= j <= i - l
```

The usual DP would be:

```text
dp[k][i] = maximum sum using at most k subarrays inside nums[0:i]
```

For a fixed `k`, if the last chosen subarray ends at `i - 1` and starts at `j`,
then:

```text
dp[k][i] = dp[k-1][j] + pref[i] - pref[j]
```

Rearrange the transition:

```text
dp[k][i] = pref[i] + max(dp[k-1][j] - pref[j])
```

where `j` is inside `[i-r, i-l]`.

That range maximum can be maintained by a monotonic queue, so one DP layer costs
`O(n)`. However, doing this for all `m` layers would be `O(nm)`, too slow when
both can be `10^5`.

### Penalty trick

Instead of directly limiting the number of chosen subarrays to at most `m`, add
a fixed penalty `cost` for every chosen subarray.

For a chosen set with original sum `S` and `cnt` subarrays, the adjusted value is:

```text
S - cost * cnt
```

For a fixed `cost`, we can compute the best adjusted value with a one-dimensional
DP:

```text
dp[i] = best adjusted result inside nums[0:i]
```

Each `dp[i]` stores two values:

```text
(sum, cnt)
```

- `sum`: best adjusted sum.
- `cnt`: how many subarrays were used to get that sum.

When two states have the same adjusted sum, prefer the one with larger `cnt`.
This tie-breaking makes the number of selected subarrays monotonic as `cost`
changes.

The transition at position `i` has two choices:

1. Do not end a subarray at `i - 1`:

```text
dp[i] = dp[i-1]
```

2. Choose a last subarray `[j, i)` with length in `[l, r]`:

```text
dp[i] = pref[i] + max(dp[j].sum - pref[j]) - cost
```

and the count becomes:

```text
dp[j].cnt + 1
```

The valid `j` range is still:

```text
i - r <= j <= i - l
```

### Monotonic queue

For every valid start `j`, the value needed by the transition is:

```text
dp[j].sum - pref[j]
```

So while scanning `i` from left to right:

1. Add `j = i - l` into the queue, because it has just become a valid start.
2. Remove starts `< i - r`, because they are now too far away.
3. Keep the queue decreasing by `(dp[j].sum - pref[j], dp[j].cnt)`.
4. The front of the queue is the best start for a subarray ending at `i - 1`.

This makes `calc(cost)` run in `O(n)`.

### Why binary search works

If `cost` is small, selecting more subarrays is attractive.

If `cost` is large, selecting fewer subarrays is attractive.

Therefore, the number of subarrays selected by `calc(cost)` is monotonic
non-increasing as `cost` increases. We binary search the largest `cost` such
that the best adjusted answer still uses at least `m` subarrays.

Suppose at this `cost`, `calc(cost)` returns:

```text
adjusted_sum = S - cost * cnt
```

The original sum for exactly `m` subarrays is recovered by adding back the
penalty for `m` subarrays:

```text
answer = adjusted_sum + cost * m
```

This is the standard Lagrangian relaxation / WQS binary search idea.

### Why this solves "at most m"

First run `calc(0)`, where no penalty is charged.

- If the best no-penalty result uses `cnt <= m` subarrays, it is already a valid
  answer for "at most `m`", so return its sum.
- If it uses more than `m` subarrays, use binary search on the penalty to force
  the result down to the best value under the `m` limit.

### At least one subarray

The DP allows selecting zero subarrays, whose adjusted sum is `0`. This is useful
internally, but the problem requires selecting at least one subarray.

If `calc(0)` selects zero subarrays, then every selectable subarray has
non-positive contribution compared with choosing nothing. In that case, the
answer must be the maximum single valid subarray sum. The helper `bestOne` finds
that with another monotonic queue:

```text
max(pref[i] - min(pref[j])) where i-r <= j <= i-l
```

### Complexity

Each `calc(cost)` call is `O(n)`. Binary search uses `O(log C)` calls, where
`C` is the value range of possible sums. With `|nums[i]| <= 10^5` and
`n <= 10^5`, using a wide `int64` range is safe.

Overall complexity:

```text
O(n log C)
```

Space complexity:

```text
O(n)
```
