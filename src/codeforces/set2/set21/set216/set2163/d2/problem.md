# D2. Diadrash (Hard Version)

[Problem link](https://codeforces.com/problemset/problem/2163/D2)

time limit per test: 3 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

**This is an interactive problem.**

The only difference from the easy version is the query limit: this version
allows at most **30 queries** per test case.

## Problem

There is a hidden permutation `p` of the integers `0, 1, ..., n-1`. You are
also given `q` distinct ranges `[l_i,r_i]`, where
`1 <= l_i <= r_i <= n`.

Find

```text
max MEX(p[l_i], p[l_i+1], ..., p[r_i])
```

over all given ranges.

You may make at most 30 queries of the form:

```text
? l r
```

The jury returns `MEX(p[l], ..., p[r])`. When the answer is known, print:

```text
! x
```

The interactor is not adaptive: the hidden permutation is fixed before any
queries are made.

## Constraints

- `1 <= t <= 100`
- `4 <= n <= 10^4`
- `1 <= q <= 3 * 10^5`
- Sum of `n` over all test cases does not exceed `10^4`
- Sum of `q` over all test cases does not exceed `3 * 10^5`
- No range is repeated within one test case

## Interaction

For each test case, read `n`, `q`, followed by the `q` ranges.

After every query, print a newline and flush the output. Exceeding 30 queries
causes a Wrong Answer verdict. Printing the final answer does not count as a
query.

## Sample Interaction

The sample is an interaction transcript, not a fixed input/output execution.
For its first test case, the hidden permutation is `[0,3,1,2]`, the ranges are
`[1,2]`, `[2,4]`, and `[1,3]`, and the required answer is `2`.

## Solution Summary

The easy version can query every relevant maximal range. With only 30
queries, the hard version instead uses two monotone MEX functions and binary
searches for their crossing point.

### 1. Remove ranges that cannot be optimal

MEX is monotone under inclusion:

```text
[l1,r1] contained in [l2,r2]
=> MEX(p[l1..r1]) <= MEX(p[l2..r2]).
```

Thus a range contained in another input range can be discarded.

The implementation first records

```text
R[l] = maximum right endpoint among input ranges starting at l
L[r] = minimum left endpoint among input ranges ending at r.
```

When scanning `l` from left to right, `next` is the right endpoint of the
last retained range. The range `[l,R[l]]` is appended to `todo` only when

```text
R[l] > next
and
L[R[l]] == l.
```

The first condition removes ranges contained in an earlier, wider range. The
second confirms that no input range with the same right endpoint begins
earlier. Consequently, if

```text
todo[i] = [l_i,r_i],
```

then both endpoint sequences are strictly increasing:

```text
l_0 < l_1 < ...
r_0 < r_1 < ...
```

Every discarded range is contained in a retained range, so the answer over
`todo` equals the answer over all input ranges.

If `[1,n]` itself is present, its MEX is `n` because `p` is a permutation of
`0..n-1`. This is the largest possible answer, so the code returns `n`
immediately.

### 2. Express an interval MEX using a prefix and a suffix

For a retained range `[l,r]`, define

```text
prefix(r) = MEX(p[1..r])
suffix(l) = MEX(p[l..n]).
```

Then

```text
MEX(p[l..r]) = min(prefix(r), suffix(l)).
```

To see why, let `x = MEX(p[l..r])`. Every value smaller than `x` occurs
inside `[l,r]`, so both the prefix and suffix contain all values below `x`.
The value `x` occurs exactly once in the whole permutation and is absent from
`[l,r]`. It must lie either before `l` or after `r`. In the first case the
suffix has MEX `x`; in the second case the prefix has MEX `x`. Hence the
minimum of the two is exactly `x`.

For `todo[i] = [l_i,r_i]`, write

```text
A(i) = prefix(r_i)
B(i) = suffix(l_i)
F(i) = min(A(i), B(i)).
```

Because `r_i` increases, the prefix only grows, so `A(i)` is nondecreasing.
Because `l_i` increases, the suffix only shrinks, so `B(i)` is
nonincreasing.

### 3. Binary search the crossing point

At index `mid`, the code obtains

```text
a = ask(1, r_mid) = A(mid)
b = ask(l_mid, n) = B(mid)
```

and records `min(a,b)`, which is the MEX of the original range
`[l_mid,r_mid]`.

- If `a < b`, then `F(mid)=a`. For every `i <= mid`,
  `F(i) <= A(i) <= a`, so no index on the left can beat the value already
  recorded. The search continues to the right.
- If `a >= b`, then `F(mid)=b`. For every `i >= mid`,
  `F(i) <= B(i) <= b`, so no index on the right can beat the recorded value.
  The search continues to the left.

Therefore every discarded part of the search interval is bounded by a value
already included in `ans`. When the binary search finishes, `ans` is the
maximum MEX among all retained ranges, and hence among all original ranges.

## Correctness Proof

We prove that `solve` returns the required maximum MEX.

First, the construction of `todo` removes only a range contained in another
input range. By MEX monotonicity, removing such a range cannot decrease the
maximum. Thus it is sufficient to find the maximum over `todo`.

For every retained range `todo[i]`, the prefix/suffix identity proves that
its MEX is exactly `F(i)=min(A(i),B(i))`. The retained right endpoints are
increasing, so `A` is nondecreasing; the retained left endpoints are
increasing, so `B` is nonincreasing.

At every binary-search step, if `A(mid)<B(mid)`, all candidates to the left
have value at most `F(mid)`; otherwise all candidates to the right have value
at most `F(mid)`. Since `F(mid)` is added to `ans` before that side is
discarded, no possible optimum is lost. Eventually every candidate has
either been queried or discarded beneath a queried upper bound. Therefore
`ans` is exactly the maximum MEX over `todo`, which is also the answer over
all input ranges.

## Query Count and Complexity

There are at most `n` retained ranges. Binary search performs at most

```text
ceil(log2(n+1))
```

iterations and uses two queries per iteration. Since `n <= 10^4`, this is at
most `2 * 14 = 28` queries, below the limit of 30.

- Local time: `O(n+q)` to build the endpoint arrays and retained ranges, plus
  `O(log n)` binary-search iterations.
- Local space: `O(n)`.
- Interactive queries: at most 28.
