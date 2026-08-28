# B. Permutation Cuts

[Problem link](https://codeforces.com/problemset/problem/2249/B)

**Contest:** [Codeforces Round 1112 (Div. 1)](https://codeforces.com/contest/2249)

time limit per test: 4 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

You are given an integer `n` and an array `a` of length `n-1`.

For a permutation `p` of length `n`, define

```text
v_i = min(max(p_1, ..., p_i), max(p_{i+1}, ..., p_n))
```

for each `1 <= i <= n-1`.

In other words, cut the permutation between positions `i` and `i+1`, take the
maximum element on each side, and let `v_i` be the smaller of these two values.

Count the number of permutations `p` of length `n` such that `v_i = a_i` for
every `1 <= i <= n-1`. Output the answer modulo `998244353`.

A permutation of length `n` is an array of `n` distinct integers from `1` to `n`
in arbitrary order.

## Constraints

- `1 <= t <= 10^4`
- `2 <= n <= 10^6`
- `1 <= a_i <= n`
- Sum of `n` over all test cases does not exceed `10^6`

## Input

The first line contains `t`, the number of test cases.

For each test case:

- The first line contains `n`.
- The second line contains `n-1` integers `a_1, a_2, ..., a_{n-1}`.

## Output

For each test case, print one integer: the number of suitable permutations,
modulo `998244353`.

## Example

### Input

```text
11
2
1
3
2 2
3
1 1
4
2 3 2
5
3 3 4 2
2
2
3
1 2
4
3 3 3
5
4 4 4 4
4
2 1 2
6
3 3 5 5 5
```

### Output

```text
2
2
0
0
2
0
2
4
12
0
8
```

## Note

In the first test case, both permutations `[1, 2]` and `[2, 1]` satisfy `v_1 = 1`.

In the second test case, the only suitable permutations are `[2, 1, 3]` and
`[3, 1, 2]`.

In the third test case, no suitable permutation exists.

## Solution

Let `p_k = n`. For a cut before `k`, the right side contains `n`, so its
smaller maximum is exactly the maximum of the left prefix. For a cut at or
after `k`, symmetrically, it is the maximum of the right suffix. Therefore:

- values of `a` to the left of `n` are non-decreasing prefix maxima;
- values of `a` to the right of `n` are non-increasing suffix maxima.

After merging equal adjacent values into blocks, the block values must thus
strictly increase and then strictly decrease. A value cannot occur in two
different blocks: it would have to be a maximum on both sides of `n`, which
would require two copies of the same value in the permutation. Also, `n`
cannot occur in `a`, because the side opposite `n` has maximum at most `n-1`.

The code scans `a` once to check precisely these conditions. It remembers
whether the scan has already started falling, and rejects either an increase
after that point or the start of a second block for a previously seen value.

### Counting

Sort the values conceptually (the implementation uses their frequencies, so
it stays linear). Consider occurrences in non-decreasing value order.

The first occurrence of a value `x` is forced: that position is where the
corresponding prefix/suffix maximum first becomes `x`, so the permutation must
place `x` there. For every later occurrence of `x`, its position may contain
any unused value smaller than `x`. If `used` permutation values have already
been assigned, there are exactly `x - used` choices. Multiply these choices
for all repeated occurrences.

Finally, the roles of `n` and the first forced `n-1` can be exchanged, giving
the initial factor `2`.

For example, for `a = [4, 4, 4, 4]`, the first `4` is forced. The remaining
three occurrences have `3`, `2`, and `1` choices, so the answer is
`2 * 3 * 2 * 1 = 12`.

### Correctness sketch

The structural check is necessary by the prefix/suffix-maximum observation
above. It is also sufficient: a valid rising side and falling side specify the
first occurrence of every maximum, while every repeated maximum can be filled
by an arbitrary unused smaller number without changing that maximum. The
counting loop enumerates exactly those choices once each. The two orientations
of the endpoints containing `n` and `n-1` are distinct and are the only
remaining choice, hence the factor `2`.

### Complexity

Each test case uses `O(n)` time and `O(n)` auxiliary memory. The total `n`
over all test cases is at most `10^6`.
