# Problem E - Payment Required

https://atcoder.jp/contests/abc402/tasks/abc402_e

**Time Limit:** 2 sec / **Memory Limit:** 1024 MiB

**Score:** 450 points

## Problem Statement

In the year 30XX, each submission to a certain contest requires a payment.

The contest has $N$ problems. Problem $i$ is worth $S_i$ points, and each
submission to it costs $C_i$ yen.

When Takahashi submits to problem $i$, he solves it with probability $P_i$%.
Each submission's result is independent of all previous submissions.

Takahashi has $X$ yen. He cannot make a submission if doing so would make his
total spending exceed $X$ yen.

Find the maximum possible expected value of the total score of solved problems,
assuming Takahashi may choose the next problem to submit to after seeing the
result of the previous submission.

Solving the same problem multiple times does not give additional score.

## Solution

Use dynamic programming over the set of problems that are still unsolved and
the remaining amount of money.

Let `dp[mask][w]` be the maximum expected score obtainable when:

- `mask` is the set of unsolved problems.
- `w` yen remains.

The answer is `dp[(1 << N) - 1][X]`.

From state `(mask, w)`, choose an unsolved problem `i` in `mask`. If its cost is
`C_i`, we can submit only when `C_i <= w`.

Let `p = P_i / 100`. After spending `C_i`:

- With probability `p`, problem `i` is solved. We gain `S_i` points and move to
  `mask ^ (1 << i)`.
- With probability `1 - p`, problem `i` is not solved. The score does not
  change, and the unsolved set remains `mask`.

So the transition is:

```text
dp[mask][w] =
  max over i in mask:
    p * (S_i + dp[mask ^ (1 << i)][w - C_i])
  + (1 - p) * dp[mask][w - C_i]
```

The value `dp[mask][w]` depends only on states with smaller money
`w - C_i`, so iterate `w` from `0` to `X`. For each budget, process all masks
and try all still-unsolved problems.

In the code, the candidate problem is obtained from a non-empty submask of
`mask` using `bits.TrailingZeros`. This may visit the same candidate more than
once, but it is still small enough for `N <= 8`; conceptually it is just taking
the maximum over every `i` in `mask`.

The complexity is `O(X * 3^N)` as written, or `O(X * N * 2^N)` with direct
iteration over the set bits. Both are easily within the limits for `N <= 8` and
`X <= 5000`.


