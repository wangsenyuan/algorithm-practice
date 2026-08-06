# D. Goods on the Shelf

[Problem link](https://codeforces.com/problemset/problem/2233/D)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

## Problem

A supermarket shelf is described by an array `a` of length `n`, where `a_i` is
the type of the good at position `i`.

The shelf is arranged correctly if all goods of the same type form one
contiguous block. Formally, for every `1 <= i < j <= n` with `a_i = a_j`, every
position `k` from `i` to `j` must also satisfy `a_k = a_i`.

You may choose two different positions at most once and swap the goods at those
positions. You may also choose not to swap anything.

Determine whether it is possible to make the shelf arranged correctly after at
most one swap.

## Constraints

- `1 <= t <= 10^4`
- `2 <= n <= 2 * 10^5`
- `1 <= a_i <= 10^9`
- The sum of `n` over all test cases does not exceed `2 * 10^5`.

## Input

The first line contains a single integer `t`, the number of test cases.

For each test case:

- The first line contains `n`.
- The second line contains `n` integers `a_1, a_2, ..., a_n`.

## Output

For each test case, output `YES` if the shelf can be arranged correctly with at
most one swap. Otherwise, output `NO`.

The answer is case-insensitive.

## Example

### Input

```text
7
3
1 2 1
2
7 7
6
1 2 3 1 2 3
6
1 1 2 3 2 3
7
1 2 3 1 2 3 4
6
1 2 1 2 1 1
6
1 2 2 3 3 1
```

### Output

```text
YES
YES
NO
YES
NO
YES
NO
```

## Notes

- In the first sample, swapping positions `1` and `2` gives `[2, 1, 1]`.
- In the second sample, the shelf is already arranged correctly.
- In the third sample, one swap is not enough.
- In the sixth sample, swapping positions `1` and `4` makes the shelf arranged
  correctly.

## Solution

First count the total frequency of every value. In a correctly arranged shelf,
each maximal block must have length exactly equal to the total frequency of
that block's value.

Scan the array from left to right and find the first block whose length is too
short. Let its value be `x`, and let `[L, R]` be the full span from the first to
the last occurrence of `x`. The one useful swap, if it exists, must either move
an adjacent outside value into this span or move one of the non-`x` values inside
the span to an endpoint. The implementation tries these candidate swaps and
checks whether all blocks then have their required lengths.

If every block already has the required length, the answer is `YES`.
