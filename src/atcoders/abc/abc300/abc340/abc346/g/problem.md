# G - Alone (ABC346)

**Contest:** [ABC346](https://atcoder.jp/contests/abc346) — AtCoder Beginner Contest 346  
**Task:** [https://atcoder.jp/contests/abc346/tasks/abc346_g](https://atcoder.jp/contests/abc346/tasks/abc346_g)

**Time limit:** 3 sec / **Memory limit:** 1024 MiB  
**Score:** 575 points

## Problem Statement

You are given an integer sequence `A = (A_1, A_2, …, A_N)`.

Count pairs of integers `(L, R)` such that:

- `1 <= L <= R <= N`
- among `A_L, A_{L+1}, …, A_R`, some value appears exactly once in that subarray
  (there exists `x` with exactly one index `i` in `[L, R]` where `A_i = x`)

## Constraints

- `2 <= N <= 2 × 10^5`
- `1 <= A_i <= N`
- All input values are integers

## Input

```text
N
A_1 A_2 … A_N
```

## Output

Print the answer.

## Sample Input 1

```text
5
2 2 1 2 1
```

## Sample Output 1

```text
12
```

Valid pairs: `(1,1), (1,3), (1,4), (2,2), (2,3), (2,4), (3,3), (3,4), (3,5),
`(4,4), (4,5), (5,5)`.

## Sample Input 2

```text
4
4 4 4 4
```

## Sample Output 2

```text
4
```

## Sample Input 3

```text
10
1 2 1 4 3 3 3 2 2 4
```

## Sample Output 3

```text
47
```

## Solution

Fix the right endpoint `R` and count the left endpoints `L` for which `[L, R]`
contains a value appearing exactly once.

For every possible `L`, maintain a coverage count: the number of values that occur
exactly once in `[L, R]`. The subarray is valid precisely when this count is positive.

Suppose the current value `A_R = x` previously appeared at positions `p1` and `p2`,
where `p1` is the latest occurrence, `p2` is the second latest occurrence, and both
are `0` when they do not exist.

- Before adding `A_R`, `x` occurs exactly once for `L` in `[p2 + 1, p1]`. If `p1 > 0`,
  remove this contribution by adding `-1` to that interval.
- After adding `A_R`, `x` occurs exactly once for `L` in `[p1 + 1, R]`. Add `1` to
  that interval.

Thus every step needs two range additions and a count of positions whose coverage is
zero. A lazy segment tree stores:

- `val`: the minimum coverage in a segment;
- `sz`: how many positions attain that minimum;
- `lazy`: the pending range addition.

All coverage values are non-negative. Position `0` is an unused sentinel whose
coverage stays zero, so the minimum over `[0, R]` is always zero. Consequently,
`Query(0, R)` returns the number of zero-coverage positions. There are `R + 1`
positions in this range, hence the number of valid left endpoints is

```text
R + 1 - Query(0, R)
```

Adding this value for every `R` gives the answer.

### Correctness

For a fixed value `x`, let its latest and second latest occurrences after processing
position `R` be `q1` and `q2`. The value `x` appears exactly once in `[L, R]` if and
only if `q1` is included and `q2` is excluded, which is exactly `q2 < L <= q1`. The
two range updates replace the interval for the old latest occurrence with the interval
for the new latest occurrence.

Therefore, the coverage at every left endpoint `L` equals the number of values that
appear exactly once in `[L, R]`. It is positive exactly when `[L, R]` satisfies the
problem condition. The segment tree counts all such left endpoints, and summing them
over every right endpoint counts every valid pair `(L, R)` exactly once.

### Complexity

Each position performs at most two range updates and one query, all in `O(log N)`
time. The total complexity is `O(N log N)` time and `O(N)` memory.
