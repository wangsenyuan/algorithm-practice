# C. Double-Rift Dial

[Problem link](https://codeforces.com/problemset/problem/2249/C)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

A permutation `p` of length `n` is written clockwise on a circular dial.

Choose a starting position `s` and read one full clockwise circle:

```text
q = [p_s, p_{s+1}, ..., p_n, p_1, ..., p_{s-1}]
```

For every non-empty prefix of `q`, take the set of values in that prefix. Split
that set into maximal segments of consecutive integers; these segments are
called blocks.

A starting position is good if every non-empty prefix produces at most two
blocks. Count the number of good starting positions.

## Constraints

- `1 <= t <= 10^4`
- `1 <= n <= 2 * 10^5`
- `p` is a permutation of `1..n`
- Sum of `n` over all test cases does not exceed `2 * 10^5`

## Input

The first line contains `t`, the number of test cases.

For each test case:

- The first line contains `n`.
- The second line contains `p_1, p_2, ..., p_n`.

## Output

For each test case, print one integer: the number of good starting positions.

## Example

### Input

```text
4
1
1
5
1 3 5 2 4
6
1 2 4 5 3 6
7
1 3 5 7 2 4 6
```

### Output

```text
1
4
3
0
```

## Note

In the first case, the only starting position is good.

In the second case, starting from position `1` is bad because after reading
three values the set is `{1, 3, 5}`, which has three blocks. The other four
starting positions are good.

## Solution

For a fixed rotation `q`, define `f[i]` as the number of blocks in the prefix
`q[0..i]`. The rotation is good exactly when `max(f) <= 2`.

The useful way to count blocks is:

```text
blocks(S) = count of values v in S where v - 1 is not in S
```

Each block has exactly one left endpoint, so this count is the number of
blocks.

When scanning one fixed rotation from left to right, suppose the next value is
`v`:

- if `v - 1` is not already present, a new block starts, so the block count
  increases by `1`;
- if `v + 1` is already present, two adjacent parts connect, so the block count
  decreases by `1`.

This change first appears at the prefix that includes `v`, and it remains true
for every longer prefix. Therefore, when `v` is at index `i`, we apply those
changes to the range `[i, n)`. After all values are processed, the segment tree
stores all values `f[0], f[1], ..., f[n-1]`; the rotation is good iff the tree
maximum is at most `2`.

Now rotate the sequence by one position. If

```text
q = [v, q_1, q_2, ..., q_{n-1}]
```

then the next rotation is:

```text
r = [q_1, q_2, ..., q_{n-1}, v]
```

Most prefix sets are almost the same as before, but the contribution of `v`
moves from the beginning to the end. Only relationships involving `v` can
change:

1. `v` was present from every old prefix, but in the new rotation it appears
   only in the final prefix.
2. If `v - 1` exists, then `v` may stop connecting to `v - 1` for prefixes that
   contain exactly one of them.
3. If `v + 1` exists, then removing `v` may make `v + 1` become a new block
   start for prefixes that contain `v + 1` but not `v`.

Because `q` is a permutation, the position of every value is known. So each of
these affected prefix groups is a continuous interval on the cyclic prefix
index line. If an affected interval wraps around the end, split it into two
ordinary ranges. The implementation uses `addCircular` for exactly this.

The concrete updates for moving old first value `v = q[i-1]` while considering
the rotation starting at index `i` are:

- For `v - 1`: prefixes between the old position of `v` and the position of
  `v - 1` lose one block-start contribution, so add `-1` on that circular
  range. If `v == 1`, there is no `v - 1`; instead, old prefixes other than
  the first had counted `1` as a block start, so add `-1` on `[1, n)`.
- For `v + 1`: prefixes that contain `v + 1` but no longer contain `v` gain
  one block-start contribution, so add `+1` on the corresponding range.
- The final prefix now includes all numbers again, including moved `v`, so add
  `+1` to the single position `[i, i+1)`.

The segment tree supports range addition and range maximum. After each rotation,
if the maximum prefix block count is at most `2`, this starting position is
good.

Rotate the permutation first so value `1` is at position `0`; this removes one
boundary case from the first construction and does not change the answer.

Complexity: `O(n log n)` per test case.
