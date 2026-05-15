# D. Billion Players Game (Codeforces 2157D)

**Limits:** 2 seconds per test, 256 MB  
**I/O:** standard input / standard output

Source: [https://codeforces.com/problemset/problem/2157/D](https://codeforces.com/problemset/problem/2157/D)

## Problem Statement

You are following the world championship of the Billion Players Game. There are `10^9` players competing, and you want to predict the final ranking `p` of Godflex, your favorite streamer.

After analyzing recent matches, you are sure that:

```text
l <= p <= r
```

but you do not know anything else.

There are `n` offers made by the in-game bookmaker. In the `i`-th offer, the bookmaker suggests an estimation `a_i` for Godflex's final ranking.

For each offer, you must choose exactly one of the following actions:

1. Ignore the offer.
2. Accept the offer by claiming that `p <= a_i`.
   - If you are right, you earn `|p - a_i|` robocoins.
   - Otherwise, you lose `|p - a_i|` robocoins.
3. Accept the offer by claiming that `p >= a_i`.
   - If you are right, you earn `|p - a_i|` robocoins.
   - Otherwise, you lose `|p - a_i|` robocoins.

After you decide how to deal with all the offers, the actual Billion Players Game is played. Godflex gets a position `p` in `[l, r]`, and then all the offers are settled up.

Your total score is the number of robocoins you are guaranteed to earn, that is, the minimum number of robocoins you earn over all possible values of `p` in `[l, r]`.

Find the maximum possible score you can guarantee.


## Solution

The main trick is to stop thinking about whether a claim is true or false.  For an offer with estimate `a`, the final payoff is always one of two linear expressions.

If we claim `p >= a`, then:

```text
p >= a  gives payoff  p - a
```

If we claim `p <= a`, then:

```text
p <= a  gives payoff  a - p
```

These formulas work even when the claim is wrong, because being wrong means losing the same absolute difference.

So:

- `p >= a` is better when `a` is small, because it subtracts `a`;
- `p <= a` is better when `a` is large, because it adds `a`.

### Pairing Intuition

Take one small estimate and one large estimate.

For the small one, claim `p >= small`.

For the large one, claim `p <= large`.

Their total payoff is:

```text
(p - small) + (large - p) = large - small
```

The unknown `p` disappears.  This pair gives the same profit no matter what the true ranking is.

Therefore, the safest profit comes from pairing the smallest estimates with the largest estimates.

For example, in sample 3:

```text
a = [1, 3, 5, 7, 9]
```

Pair `1` with `9`:

```text
(p - 1) + (9 - p) = 8
```

Pair `3` with `7`:

```text
(p - 3) + (7 - p) = 4
```

Ignore `5`.

The guaranteed score is:

```text
8 + 4 = 12
```

which matches the sample answer.

### What If the Counts Are Not Balanced?

Balanced pairs remove `p` completely.  But sometimes it is also useful to take extra one-sided offers.

If we take extra `p >= a` offers, then the score increases as `p` increases.  The worst case is:

```text
p = l
```

If we take extra `p <= a` offers, then the score decreases as `p` increases.  The worst case is:

```text
p = r
```

So the algorithm tries every possible imbalance:

- `k` extra small estimates used as `p >= a`;
- `k` extra large estimates used as `p <= a`.

Everything else is paired smallest-with-largest as much as possible.

### Why Sorting Is Enough

Sort the estimates:

```text
a_1 <= a_2 <= ... <= a_n
```

For any fixed numbers of offers:

- use the smallest estimates for `p >= a`;
- use the largest estimates for `p <= a`;
- ignore the middle estimates.

Using a larger estimate for `p >= a` only makes `p-a` worse.  Using a smaller estimate for `p <= a` only makes `a-p` worse.  So this greedy choice is forced once the counts are fixed.

Prefix sums let us get:

- the sum of the smallest `k` estimates;
- the sum of the largest `k` estimates.

### Case 1: Extra `p >= a` Offers

Let `k >= 0` be the number of extra `p >= a` offers.

After reserving these `k` extra small estimates, use as many balanced pairs as possible:

```text
pairs = floor((n - k) / 2)
```

We choose:

- smallest `pairs + k` estimates for `p >= a`;
- largest `pairs` estimates for `p <= a`.

The worst case is `p = l`, so the score is:

```text
score = sum(largest pairs)
      - sum(smallest pairs+k)
      + k * l
```

The code computes this as:

```go
pairs := (n - k) / 2
cur := sumLarge(pairs) - sumSmall(pairs+k) + int64(k)*l
```

### Case 2: Extra `p <= a` Offers

Let `k >= 1` be the number of extra `p <= a` offers.

Again:

```text
pairs = floor((n - k) / 2)
```

We choose:

- largest `pairs + k` estimates for `p <= a`;
- smallest `pairs` estimates for `p >= a`.

The worst case is `p = r`, so the score is:

```text
score = sum(largest pairs+k)
      - sum(smallest pairs)
      - k * r
```

The code computes this as:

```go
pairs := (n - k) / 2
cur := sumLarge(pairs+k) - sumSmall(pairs) - int64(k)*r
```

### Why Ignoring Middle Offers Is Fine

For a fixed imbalance `k`, after taking the useful smallest and largest estimates, there may be some middle estimates left.

Those are ignored.

This is optimal because middle values are worse than the chosen extremes:

- they are too large to be good `p >= a` offers;
- they are too small to be good `p <= a` offers;
- adding an unpaired middle value would tilt the score toward one endpoint without enough compensation.

Thus checking all `k` covers every optimal strategy.

### Complexity

For each test case:

```text
Sort estimates: O(n log n)
Prefix sums:    O(n)
Try all k:      O(n)
Memory:         O(n)
```

The sum of `n` over all test cases is at most `2 * 10^5`, so this is easily fast enough.
