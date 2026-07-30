# E - Maximize Rating

[Problem link](https://atcoder.jp/contests/abc327/tasks/abc327_e)

**Contest:** [AtCoder Beginner Contest 327](https://atcoder.jp/contests/abc327)

time limit: 2 sec

memory limit: 1024 MiB

score: 475 points

## Problem Statement

Takahashi participated in `N` contests and earned a performance `P_i` in the
`i`-th contest. He wants to choose some (at least one) contests from these and
maximize his rating calculated from the results of those contests.

Find the maximum possible rating he can achieve by optimally choosing the
contests.

Here, Takahashi's rating `R` is calculated as follows, where `k` is the number
of chosen contests and `(Q_1, Q_2, ..., Q_k)` are the performances in the chosen
contests **in the order he participated**:

```text
R = (sum_{i=1}^{k} (0.9)^{k-i} Q_i) / (sum_{i=1}^{k} (0.9)^{k-i}) - 1200 / sqrt(k)
```

## Constraints

- `1 <= N <= 5000`
- `1 <= P_i <= 5000`
- All input values are integers.

## Input

```text
N
P_1 P_2 ... P_N
```

## Output

Print the maximum possible rating that Takahashi can achieve.

Your output will be considered correct if the absolute or relative error from
the true value is at most `10^{-6}`.

## Sample Input 1

```text
3
1000 600 1200
```

## Sample Output 1

```text
256.735020470879931
```

Choosing the first and third contests gives:

```text
R = (0.9 * 1000 + 1.0 * 1200) / (0.9 + 1.0) - 1200 / sqrt(2) = 256.73502...
```

## Sample Input 2

```text
3
600 1000 1200
```

## Sample Output 2

```text
261.423219407873376
```

Selecting all three contests maximizes the rating.

## Sample Input 3

```text
1
100
```

## Sample Output 3

```text
-1100.000000000000000
```

The rating can also be negative.


## ideas
1. 考虑选中了k个Q, 第二项是固定值
2. 考虑第一项如何最大化
3. (0.9) ^^ (k - i) 
4. 当i从1..k, i越小, 0.9的pow次越小
5. 所以, 这种情况下, 为了让贡献最大化, 应该使用最小的Q
6. 所以Q应该递增排列
7. 剩下是个简单的dp
8. 题目不能排序
9. 貌似有精度问题

## Solution Summary

For a fixed number `k` of chosen contests, the denominator

```text
1 + 0.9 + ... + 0.9^(k-1)
```

and the penalty `1200 / sqrt(k)` are fixed.  We only need the largest possible
weighted average for every `k`.

Let `fp[k]` be that denominator, and let `dp[k]` be the maximum value of the
first term of the rating after choosing exactly `k` contests from the prefix
processed so far:

```text
dp[k] = max (0.9^(k-1) * Q1 + ... + Qk) / fp[k]
```

When the current performance is `v`, appending it to an optimal choice of
`k-1` contests changes its numerator from `dp[k-1] * fp[k-1]` to

```text
0.9 * dp[k-1] * fp[k-1] + v.
```

Therefore the in-place transition is:

```go
tmp := (dp[k-1]*fp[k-1])*0.9 + float64(v)
tmp /= fp[k]
dp[k] = max(dp[k], tmp)
```

Process `k` in decreasing order so the source state still belongs to the
previous prefix.  The answer is then:

```text
max over k = 1..N of dp[k] - 1200 / sqrt(k).
```

### Important reachable-state bound

Before processing `p[i]`, at most `i` contests can have been chosen.  Thus the
current contest can only create a state up to `k = i + 1`:

```go
for i, v := range p {
    for k := i + 1; k > 0; k-- {
        // transition from k-1 to k
    }
}
```

Do **not** iterate `k` down from `N` while using a finite sentinel such as
`-1e16`.  Those unreachable states would still be transitioned repeatedly.
Although each transition multiplies the old value by roughly `0.9`, after a
few hundred transitions the finite sentinel is diluted and can become a fake
competitive state.  It then claims an impossible large `k`, receiving an
incorrectly small penalty.  This is a state-validity bug, not a floating-point
precision problem.  Limiting `k` to `i+1` ensures every transition starts from
a genuinely reachable state.

The time complexity is `O(N^2)` and the memory complexity is `O(N)`.
