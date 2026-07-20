# F. Vessels, Heights and Two Versions (Hard Version)

[Problem link](https://codeforces.com/problemset/problem/2234/F)

**Contest:** [Codeforces Round 1102 (Div. 2)](https://codeforces.com/contest/2234)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

This is the hard version of the problem. The difference between the versions is that in
this version, constraints on `n` and on the number of test cases are higher. You can hack
only if you solved all versions of this problem.

There are `n` communicating vessels of infinite height arranged in a circle. The base area
of each vessel is `1` cm², and between the `i`-th vessel and the `(i mod n) + 1`-th vessel
there is a connection of negligible volume at height `h_i` cm. For each vessel `i`, find the
maximum total volume of water in cm³ that can be placed in these vessels under the condition
that the `i`-th vessel remains empty.

Formally, you are given an array `h_1, h_2, ..., h_n`. A cyclic array of non-negative
integers `w_1, w_2, ..., w_n` is called good if the following holds:

- For every `i` from `1` to `n`, if `max(w_i, w_{(i mod n) + 1}) > h_i`, then
  `w_i = w_{(i mod n) + 1}`. In other words, if the maximum of two neighboring elements of
  the array `w` exceeds the corresponding element of the array `h`, then these two
  neighboring elements of the array `w` must be equal.

For each `i` from `1` to `n`, output the maximum possible sum
`w_1 + w_2 + ... + w_n` among all good arrays `w_1, w_2, ..., w_n`, under the condition that
`w_i = 0`.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t`
(`1 <= t <= 10^4`). The description of the test cases follows.

The first line of each test case contains one integer `n` (`3 <= n <= 2 * 10^5`) — the
number of vessels.

The second line of each test case contains `n` integers `h_1, h_2, ..., h_n`
(`1 <= h_i <= 10^9`) — the heights of the partitions between the vessels.

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output `n` integers — the `l`-th integer means the maximum total volume
of water in cm³ in the vessels under the condition that the `l`-th vessel remains empty.

## Example

### Input

```text
4
4
1 2 3 4
5
5 3 1 5 2
6
3 4 2 6 1 5
7
1 2 1 4 2 3 5
```

### Output

```text
6 6 7 9
17 16 14 14 17
21 21 20 20 21 21
17 17 17 17 21 21 22
```

## Note

Consider the first test case.

- To keep vessel `1` empty, one good array is `w = [0, 1, 2, 3]`, with a total of `6` cm³ of
  water.
- To keep vessel `2` empty, one good array is `w = [1, 0, 2, 3]`, with a total of `6` cm³ of
  water.
- To keep vessel `3` empty, one good array is `w = [2, 2, 0, 3]`, with a total of `7` cm³ of
  water.
- To keep vessel `4` empty, one good array is `w = [3, 3, 3, 0]`, with a total of `9` cm³ of
  water.

For example, the array `w = [2, 2, 0, 4]` is not good, because `max(w_3, w_4) > h_3`, and
therefore `w_3 = w_4` must hold.

It can be shown that each of the arrays above has the maximum possible sum among all
suitable options.

## Solution

Fix one vessel `x` and require `w_x = 0`. Then the remaining `n - 1` vessels form a path
between the two edges adjacent to `x`.

Consider walking from the empty vessel along one side of this path. Suppose the previous
vessel can have height `cur`, and the next edge has height `h`.

- If the next vessel gets a different height from `cur`, then both neighboring heights must
  be at most `h`.
- If the next vessel keeps height `cur`, then the edge condition is still valid even when
  `cur > h`, because the two neighboring vessels are equal.

So the best possible next height is:

```text
max(cur, h)
```

Starting from `cur = 0`, the heights along one side are therefore the prefix maxima of the
edge heights on that side. For a vessel in the middle of the path, there is a constraint
from both directions, so its final maximum height is the smaller of the two directional
limits. A convenient way to count the total is to split the circle at an edge with maximum
height: that edge is never the bottleneck, and every answer becomes the sum of two
independent side contributions ending at that maximum edge.

Duplicate the array:

```text
H = h + h
```

Choose one position `first` where `H[first]` is globally maximum, and set
`last = first + n`. This represents one full cycle cut at a maximum edge.

The code computes two arrays:

- `dp[i]`: total contribution of the side from `first` to `i`;
- `fp[i]`: total contribution of the side from `i` to `last`.

Both are sums of running maxima over many intervals. For example, `fp[i]` is the sum of the
prefix maxima when walking from edge `i` toward `last`.

These sums can be computed with a monotonic stack. When processing a new height `H[i]`,
all smaller previous heights are popped, because `H[i]` becomes the running maximum for all
intervals that used to be capped by those smaller heights. If `p` is the nearest previous
position with `H[p] >= H[i]`, then:

```text
dp[i] = dp[p] + H[i] * (i - p)
```

The reverse pass is symmetric:

```text
fp[i] = fp[p] + H[i] * (p - i)
```

where `p` is the nearest next position with `H[p] >= H[i]`.

Finally, if the empty vessel corresponds to position `i` in the doubled cycle, its answer is:

```text
left side  = dp[i - 1]   if i > first
right side = fp[i]
answer    = left side + right side
```

The loop over `i = first ... last` fills all `n` answers modulo the original indices.

## Correctness

For one fixed empty vessel, the first vessel on either side cannot exceed the height of the
edge adjacent to the empty vessel; otherwise it would differ from `0` across an edge whose
limit is too low. Inductively, if the previous vessel has best height `cur`, then the next
vessel can either stay equal to `cur`, or choose a different height not exceeding the edge
height `h`. Thus its best height is exactly `max(cur, h)`. Therefore each side's optimal
heights are running maxima of edge heights from the empty end.

The globally maximum edge can be used as the cut because no running maximum on either side
can exceed it. Splitting there does not remove any stricter constraint; it only separates the
two directional contributions for every possible empty vessel.

The monotonic-stack recurrences compute the same running-maximum sums. After removing all
smaller heights, the nearest remaining height `>= H[i]` is the last point where the running
maximum would not become `H[i]`. Every interval ending at `i` after that point has maximum
`H[i]`, giving the term `H[i] * (i - p)`. The reverse pass has the same argument in the
opposite direction. Therefore `dp` and `fp` give the correct side sums, and adding the two
side sums gives the maximum total for each empty vessel.

## Complexity

Each index of the doubled array is pushed to and popped from each monotonic stack at most
once. All passes are linear.

The time complexity is `O(n)` per test case, and the memory complexity is `O(n)`.
