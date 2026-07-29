# E. Cake Trial

[Problem link](https://codeforces.com/problemset/problem/2238/E)

**Contest:** [Codeforces Round 1106 (Div. 2)](https://codeforces.com/contest/2238)

## Problem

GLaDOS arranged `n` cakes in a row. Each cake is real (`T`) or fake (`F`).
Chell can see the truth, but her answer must declare all fake cakes as one
contiguous segment `[l, r]` (possibly empty — all cakes declared real).

A mistake is a cake whose declared type differs from the truth.

Some cakes are not placed yet. You are given a string `s` of length `n`:

- `T` — fixed real
- `F` — fixed fake
- `N` — GLaDOS may set it to `T` or `F`

GLaDOS fills all `N` first (adversarially). Then Chell, seeing the final string,
chooses a segment to minimize mistakes. Find the maximum number of mistakes
GLaDOS can force.

## Constraints

- `1 <= t <= 10^4`
- `1 <= n <= 500`
- `s` consists of `T`, `F`, `N`
- Sum of `n^3` over all test cases does not exceed `500^3`

## Input

The first line contains `t`.

For each test case:

- The first line contains `n`.
- The second line contains the string `s`.

## Output

For each test case, print one integer — the maximum mistakes GLaDOS can
guarantee.

## Sample 1

```text
Input
4
FTFF

Output
1
```

## Sample 2

```text
Input
5
TNFTT

Output
0
```

## Sample 3

```text
Input
6
TFTTTN

Output
1
```

## Sample 4

```text
Input
6
TNNFTF

Output
2
```

## Sample 5

```text
Input
7
TNFNTNF

Output
2
```

## Sample 6

```text
Input
6
NNFFNN

Output
2
```

## Sample 7

```text
Input
7
TNTFNTN

Output
2
```

## Sample 8

```text
Input
1
N

Output
0
```

## Sample 9

```text
Input
5
NNNNN

Output
2
```

## Sample 10

```text
Input
10
NNNTTNNNFN

Output
3
```

## ideas
1. 只有T/F的时候, Chell的策略是什么呢?
2. 假设她选择了一段l...r, 那么她的正确度 = T[1...n] - T[l...r] + F[l...r]
3. F[l...r] = r - l + 1 - (T[r] - T[l-1])
4. T[n] - (T[r] - T[l-1]) + r - T[r] - (l - 1 - T[l-1])
5. T[n] + r - 2 * T[r] - (l - 1 - 2 * T[l-1])
6. 修改N => T\F, 要使得f(l..r)的最大值最小
7. 有个的感觉, 就是应该尽量的把两头的改成F, 而把中间的变成T
8. 假设把, l...r中间的都改成T, 然后能快速的知道现在的结果的话, 就能搞
9.

## Solution Summary

### Rewriting Chell's objective

Fix one final assignment of all `N` cakes. Suppose it contains `f` fake cakes
in total, and Chell declares `[l, r]` fake. Her mistakes are the fake cakes
outside the segment plus the real cakes inside it:

\[
f-F(l,r)+T(l,r)=f-(F(l,r)-T(l,r)).
\]

Give each cake a weight:

```text
F -> +1
T -> -1
```

Then `F(l, r) - T(l, r)` is the segment sum. Chell chooses the segment with
the maximum sum, and the empty segment is allowed, so her minimum number of
mistakes is

\[
f-M,
\]

where `M` is the non-negative maximum subarray sum of the final `+1/-1`
array. GLaDOS wants to maximize this value while choosing every `N`.

### Kadane state

For ordinary maximum subarray sum, Kadane's algorithm keeps:

- `suffix`: the maximum non-negative segment sum ending at the current index;
- `best`: the maximum segment sum seen anywhere so far.

If the next weight is `+1`, then:

\[
suffix'=suffix+1,
\qquad
best'=\max(best,suffix').
\]

If the next weight is `-1`, then:

\[
suffix'=\max(suffix-1,0),
\qquad
best'=best.
\]

The code runs these updates for every possible way to assign the `N` cakes.

### DP definition

After processing a prefix, `dp[f][suffix]` is the smallest possible `best`
among all assignments of that prefix such that:

- exactly `f` processed cakes are fake;
- Kadane's current non-negative ending sum is `suffix`.

For the same `(f, suffix)`, GLaDOS prefers the smaller `best`, because the
final number of forced mistakes is `f - best`. Thus all other assignments with
the same state can be discarded.

Initially the prefix is empty:

```text
dp[0][0] = 0
```

and every other state is `inf`.

### Transitions

If the current character can become `F` (`F` or `N`), it contributes `+1`:

```go
ndp[f+1][suffix+1] = min(
    ndp[f+1][suffix+1],
    max(dp[f][suffix], suffix+1),
)
```

The fake count and the ending sum both increase by one. The new global maximum
is either the previous one or the new ending segment.

If the current character can become `T` (`T` or `N`), it contributes `-1`:

```go
nextSuffix := max(suffix-1, 0)
ndp[f][nextSuffix] = min(ndp[f][nextSuffix], dp[f][suffix])
```

The fake count does not change. A negative cake cannot create a larger maximum
subarray sum, and a negative ending segment is replaced by the empty segment.

For `N`, both transitions are available; for a fixed `F` or `T`, only its own
transition is available.

### Final answer

After all cakes are processed, every reachable state describes one class of
final assignments. Its value is `f - dp[f][suffix]`, so the answer is:

```go
ans = max(ans, f-dp[f][suffix])
```

For example, GLaDOS can assign `NNNNN` as `FTFTF`. Its weights are
`+1, -1, +1, -1, +1`; it has `f = 3` fake cakes and `M = 1`, so it forces
`3 - 1 = 2` mistakes.

### Correctness sketch

The weight transformation proves that Chell's optimal response to every final
assignment has value `f - M`. Kadane's two updates maintain exactly the
maximum ending sum and global maximum subarray sum for that assignment. The DP
branches over every legal choice for every `N`, while retaining the minimum
possible `M` for each equal `(f, suffix)` state. Therefore the final maximum
of `f - M` is exactly the number of mistakes GLaDOS can force.

### Complexity

There are `O(n^2)` states and two constant-time transitions per state for each
of `n` positions. The time complexity is `O(n^3)` and the memory complexity is
`O(n^2)`.
