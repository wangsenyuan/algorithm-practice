# C. 0mar and Alternating Sums

[Problem link](https://codeforces.com/problemset/problem/2246/C)

**Contest:** [Codeforces Round 1108 (Div. 2)](https://codeforces.com/contest/2246)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Define the alternating sum of an array `b` of length `k` to be

```text
∑_{i=1}^{k} (−1)^{i+1} b_i
```

You are given a non-decreasing\* array `a` of length `n` such that for all
`1 ≤ i ≤ n`, either `a_i = −1` or `a_i` is a positive integer. Find the number
of sequences `1 ≤ i_1 < i_2 < … < i_k ≤ n` such that the alternating sum of the
sequence `a_{i_1}, a_{i_2}, …, a_{i_k}` is `0`. Since this number may be large,
output it modulo `10^9 + 7`.

Two sets of indices are considered different if they have different lengths or
differ in at least one position.

\*A sequence `a_1, …, a_n` is non-decreasing if `a_1 ≤ a_2 ≤ … ≤ a_n`.

## Constraints

- `1 ≤ t ≤ 10^4`
- `1 ≤ n ≤ 2 · 10^5`
- `a_i = −1` or `1 ≤ a_i ≤ 10^9`
- The array `a` is non-decreasing
- Sum of `n` over all test cases does not exceed `2 · 10^5`

## Input

Each test contains multiple test cases. The first line contains the number of
test cases `t`.

The first line of each test case contains a single integer `n` — the length of
the array.

The second line contains `n` integers `a_1, a_2, …, a_n`.

## Output

For each test case, output a single integer — the number of subsequences that
have an alternating sum of zero, modulo `10^9 + 7`.

A subsequence of length `0` is considered to have an alternating sum of zero.

## Example

### Input

```text
4
5
-1 1 1 2 3
3
1 2 3
4
1 3 5 7
14
-1 -1 -1 1 2 2 3 3 3 5 5 5 5 5
```

### Output

```text
6
1
1
1536
```

## Note

In the first example, the following subsequences have an alternating sum of
zero:

- `[]`
- `[a_2, a_3] = [1, 1]`
- `[a_1, a_2, a_4] = [−1, 1, 2]`
- `[a_1, a_3, a_4] = [−1, 1, 2]`
- `[a_1, a_4, a_5] = [−1, 2, 3]`
- `[a_1, a_2, a_3, a_4, a_5] = [−1, 1, 1, 2, 3]`

In the second example, only the empty subsequence `[]` has an alternating sum
of zero.

## Solution

Group equal values together. Suppose a group contains `s` equal elements.
There are

```text
2^(s-1)
```

ways to choose an even number of its elements, and the same number of ways to
choose an odd number. This follows by pairing every subset that does not use a
fixed element with the subset obtained by adding that element; the two subset
sizes have opposite parity.

Let `d` be the number of distinct values in the whole array, including `-1`
when it occurs. Multiplying `2^(s-1)` over all groups gives

```text
product 2^(s-1) = 2^(sum(s)-d) = 2^(n-d).
```

Call this product `base`.

### Choosing an even number of `-1`s

The selected `-1`s cancel in pairs. The remaining selected positive elements
must contain an even number of elements from every value group.

To see why, write the selected positive elements in non-decreasing order as
`b_1, ..., b_m`. If `m` is even, their alternating sum can be paired as

```text
(b_1 - b_2) + (b_3 - b_4) + ... + (b_(m-1) - b_m).
```

Every term is non-positive, so the sum is zero exactly when every pair is
equal. Because equal values are consecutive, this is equivalent to selecting
an even number from every positive group. If `m` is odd, regrouping the sum as

```text
b_1 + (b_3 - b_2) + (b_5 - b_4) + ...
```

makes it strictly positive, so it cannot be zero.

Therefore this case contributes exactly `base = 2^(n-d)` subsequences. It also
covers the empty subsequence.

### Choosing an odd number of `-1`s

This case exists only when the array contains `-1`. An odd number of selected
`-1`s contributes `-1`. Because the following positive part begins at an even
position of the complete subsequence, its usual alternating sum

```text
b_1 - b_2 + b_3 - b_4 + ...
```

must equal `-1`.

For an even-length positive part, every paired difference `b_(2j-1)-b_(2j)`
is non-positive. Their sum is `-1` exactly when:

- one pair is `(v, v+1)`, contributing `-1`;
- every other pair contains two equal values.

Equivalently, the groups for two consecutive values `v` and `v+1` are chosen
an odd number of times, while every other positive group is chosen an even
number of times. An odd-length positive part cannot have a negative usual
alternating sum, by the same regrouping used above.

Let `ell` be the number of distinct positive values `v` for which `v+1` also
occurs. For each such adjacent pair, changing its two groups from even choices
to odd choices does not change the number of possibilities: both parities have
`2^(s-1)` choices. Hence every adjacent pair contributes another `base` ways.

The final formula is therefore

```text
without -1: answer = base
with -1:    answer = base * (ell + 1)
```

where `base = 2^(n-d)`.

### How the implementation computes it

During the group scan, `res1` accumulates `base`. For positive groups, `ways`
stores the prefix product of their `2^(s-1)` factors.

If `-1` occurs, `res2` is the number of ways to select an odd number of them.
The reverse scan maintains `suf`, the corresponding product for positive
groups to the right. Whenever two neighboring positive groups have values
`v` and `v+1`, `tmp` combines:

- odd-size choices for those two groups;
- even-size choices from the prefix before them;
- even-size choices from the suffix after them.

Their sum is `res3`. Multiplying it by `res2` supplies the odd choices for the
`-1` group, and adding `res1` includes the even-`-1` case.

### Correctness

The two cases partition all subsequences by the parity of the number of
selected `-1`s. The first case counts exactly the subsequences selecting every
value group evenly. In the second case, the pairing argument proves that a
zero alternating sum is equivalent to choosing exactly one consecutive
positive-value pair oddly and every other positive group evenly. Each valid
subsequence belongs to exactly one such pair, so the cases are disjoint and
complete. Thus the algorithm returns the required count.

### Complexity

The groups are scanned a constant number of times, so the time complexity is
`O(n)`. The positive group records use `O(d)` space, which is `O(n)` in the
worst case.
