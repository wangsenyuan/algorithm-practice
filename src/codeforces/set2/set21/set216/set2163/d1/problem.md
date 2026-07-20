# D1. Diadrash (Easy Version)

[Problem link](https://codeforces.com/problemset/problem/2163/D1)

time limit per test: 3 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

**This is an interactive problem.**

## Problem

There is a hidden permutation `p` of the integers `0, 1, ..., n-1`.
You are also given `q` ranges `[l_i, r_i]` (`1 ≤ l_i ≤ r_i ≤ n`).

Find

```text
max_{i=1..q} MEX(p[l_i], p[l_i+1], ..., p[r_i])
```

You may make at most `max(300, ceil(n/2) + 2)` queries of the form:

- `? l r` — the jury returns `MEX(p[l], ..., p[r])`

When you know the answer, output:

- `! x`

`MEX` of a multiset is the smallest non-negative integer that does not appear
in it.

## Constraints

- `1 ≤ t ≤ 100`
- `4 ≤ n ≤ 10^4`
- `1 ≤ q ≤ 3 · 10^5`
- Sum of `n` over all tests ≤ `10^4`
- Sum of `q` over all tests ≤ `3 · 10^5`

## Interaction

For each test case, read `n`, `q`, then `q` ranges.

To query, print `? l r` and flush; then read one integer.

To answer, print `! x` and flush.

## Sample (interaction sketch)

Input (jury side interleaved):

```
3
4 3
1 2
2 4
1 3
...
```

One valid interaction sequence ends with answers `2`, `2`, `1` for the three
test cases (exact queries depend on strategy).


### ideas
1. 可以用二分, 找到包含0的位置?
2. 然后, 找到包含0的区间, 那些被完全包含的, 可以直接去掉
3. 这时候, 貌似最多只有一半的区间包含0, l1 ...0 .. r1,
4. l1 < l2 ... < lk
5. r1 < r2 ... < rk
6. 这样的区间数, 最多还是有n/2个, 如果加上2分查询 => 超过了限制
7. 搞不下去了~
8. 剩下的区间, 感觉是会组成一个平台, 两边小, 中间大
9. 假设f(i) = w 是最大的, 它应该包含0, 1, 2 ... w - 1
10. 这些数字分布在0的两边, 那么当l[i] > pos[x]的时候, f(i) <= x, 
11. 同样的r[i] < pos[y]的时候, f(i) <= y
12. 

## Solution Summary

The key is that we do not need the exact position of `0` and do not need to
search for a peak among the ranges.

### 1. Only ranges containing `0` matter

For any range:

```text
MEX(range) > 0  if and only if  range contains 0.
```

A range not containing `0` has MEX `0`, so it cannot improve an answer that
is initialized to `0`.

The code uses zero-based indices and sets

```text
mid = (n - 1) / 2.
```

It first queries `[0, mid]`:

- If the returned value `w > 0`, then `0` is in the first half
  `[0, mid]`.
- If `w == 0`, then `0` is in the second half `[mid+1, n-1]`.

This one query tells us which endpoints can belong to a range with positive
MEX.

### 2. Keep one widest range for each endpoint

MEX is monotone under inclusion. If one interval is contained in another,
then

```text
[l1,r1] contained in [l2,r2]
=> MEX(p[l1..r1]) <= MEX(p[l2..r2]).
```

Therefore, among ranges sharing one endpoint, only the widest one can be
useful.

The implementation builds two arrays:

```text
R[l] = the maximum right endpoint of an input range starting at l
L[r] = the minimum left endpoint of an input range ending at r
```

Thus `[l,R[l]]` contains every input range starting at `l`, and `[L[r],r]`
contains every input range ending at `r`.

### 3. Query the relevant half

#### `0` is in the first half

If an input range has positive MEX, it contains the position of `0`.
Therefore its left endpoint satisfies

```text
l <= position(0) <= mid.
```

The code visits every `l` from `0` through `mid` and, when such a range
exists, queries `[l,R[l]]`.

For any omitted range `[l,r]`, the queried range `[l,R[l]]` is an input range
that contains `[l,r]`. Its MEX is therefore at least as large.

#### `0` is in the second half

Symmetrically, every range with positive MEX has

```text
r >= position(0) > mid.
```

The code visits every `r` from `mid+1` through `n-1` and queries `[L[r],r]`
when an input range ending at `r` exists. This widest range dominates every
other input range with the same right endpoint.

The answer is the maximum value returned by these queries.

### Correctness Proof

We prove that the algorithm returns the maximum MEX among all given ranges.

Consider any input range `I`.

- If `I` does not contain `0`, its MEX is `0`, which is already covered by
  the initial value of `ans`.
- Suppose `I` contains `0`. If `0` is in the first half, the left endpoint of
  `I` is at most `mid`. The algorithm queries the widest input range with
  that same left endpoint, which contains `I` and therefore has MEX at least
  `MEX(I)`. If `0` is in the second half, the symmetric argument applies to
  the right endpoint and `[L[r],r]`.

Hence every input range is dominated by a range queried by the algorithm, so
the maximum queried MEX is at least the required answer. Conversely, every
range queried after the half query is itself one of the given input ranges,
so its MEX cannot exceed the required answer. The two values are equal, and
the algorithm is correct.

### Query Count

The first-half side contains `mid+1 = ceil(n/2)` possible left endpoints. The
second-half side contains at most `floor(n/2)` possible right endpoints.
Therefore the algorithm makes at most

```text
1 + ceil(n/2)
```

queries: one query to locate the half containing `0`, followed by at most one
query per relevant endpoint. This is within the allowed limit
`max(300, ceil(n/2)+2)`.

### Complexity

- Local computation: `O(n+q)` per test case.
- Local memory: `O(n)`.
- Interactive queries: at most `1+ceil(n/2)`.
