# C. Incremental Stay

[Problem link](https://codeforces.com/problemset/problem/2151/C)

**Contest:** [Codeforces Round 1053 (Div. 2)](https://codeforces.com/contest/2151)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

## Problem

A museum door records `2n` distinct activity times in increasing order. Each
event is either an entrance or an exit, but the direction and visitor identity
are unknown. A visitor's stay is its exit time minus entrance time.

For every simultaneous-visitor limit `k` from `1` to `n`, find the maximum
possible sum of all visitor stay times, assuming at most `k` visitors may be in
the museum at once. The museum is empty at time `0` and after all events.

## Input

The first line contains the number of test cases `t` (`1 <= t <= 10^4`).

For each test case:

- the first line contains `n` (`1 <= n <= 2 * 10^5`);
- the second line contains `2n` strictly increasing times `a_i`
  (`1 <= a_i <= 10^9`).

The sum of `n` over all test cases is at most `2 * 10^5`.

## Output

For every test case, print `n` integers: the maximum total stay time for each
`k` from `1` to `n`.

## Sample

```text
Input
3
1
32 78
2
4 5 6 9
4
6149048 26582657 36124499 43993239 813829899 860114890 910238130 913669539

Output
46
4 6
78018749 1737022233 1845329695 3385003015
```

## Status

Implemented with the optimal occupancy profile and prefix sums of consecutive
time gaps. The tests cover each official test case independently.

## Ideas

Fix a capacity `k`. To maximize the total stay time, maximize the number of
visitors inside during every gap between consecutive sensor events. The greedy
event assignment is therefore:

1. let a visitor enter at each of the first `k` events;
2. let a visitor leave at each of the last `k` events;
3. in between, make a visitor leave when the museum already contains `k`
   people, and otherwise make one enter.

Thus the occupancy alternates between `k` and `k - 1` in the middle. Let
`d_i = a_{i+1} - a_i` be the one-based gap after event `i`.

- For `k = 1`, the answer is the sum of odd-indexed gaps.
- For `k = 2`, it is that same sum plus twice the sum of even-indexed gaps.
- For `k >= 3`, compare capacities `k` and `k - 2`: occupancy increases by
  `2` exactly on gaps `k - 1` through `2n - k + 1`.

With prefix sums of `d_i`, this gives:

\[
ans_k = ans_{k-2} + 2 \cdot \sum_{i=k-1}^{2n-k+1} d_i.
\]

All answers are computed in `O(n)` time per test case.

## Summary

The maximum total stay time is the weighted area under the largest valid
occupancy profile. Two parity-specific base profiles plus the `k` versus
`k - 2` recurrence reduce every capacity answer to an `O(1)` prefix-sum query.
