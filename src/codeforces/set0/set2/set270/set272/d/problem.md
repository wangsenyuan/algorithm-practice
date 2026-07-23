# D. Dima and Two Sequences

[Problem link](https://codeforces.com/problemset/problem/272/D)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Dima has two sequences of points:

- `(a_1, 1), (a_2, 2), ..., (a_n, n)`
- `(b_1, 1), (b_2, 2), ..., (b_n, n)`

Count the number of distinct sequences of length `2n` that can be assembled by
using each of these `2n` points exactly once, such that the x-coordinates are
non-decreasing.

Two assembled sequences are distinct if they differ in at least one position as
a point pair.

Print the answer modulo `m`.

## Constraints

- `1 <= n <= 10^5`
- `1 <= a_i, b_i <= 10^9`
- `2 <= m <= 10^9 + 7`

## Input

```text
n
a_1 a_2 ... a_n
b_1 b_2 ... b_n
m
```

## Output

Print a single integer — the answer modulo `m`.

## Examples

### Sample 1

```text
Input
1
1
2
7

Output
1
```

Only sequence: `(1, 1), (2, 1)`.

### Sample 2

```text
Input
2
1 2
2 3
11

Output
2
```

Sequences:

- `(1, 1), (2, 2), (2, 1), (3, 2)`
- `(1, 1), (2, 1), (2, 2), (3, 2)`

## Solution

Any valid sequence is a merge of all `2n` points with nondecreasing x. Points
with different x-values cannot swap past each other, so the order is fully
determined except inside groups that share the same x.

Let `cnt(x)` be how many points have x-coordinate `x` (from both arrays).
Those `cnt(x)` points may appear in any order, contributing `cnt(x)!` ways.
Groups for different x are independent, so the naive product is:

```text
∏_x  cnt(x)!
```

When `a_i = b_i`, the two points `(a_i, i)` and `(b_i, i)` are identical, so
every permutation of a group double-counts by swapping them. If there are `d`
such indices, divide by `2^d`:

```text
answer = (∏_x cnt(x)!) / 2^d   (mod m)
```

`m` may not be prime, so `2` may have no modular inverse. Instead, while
multiplying the factors `2, 3, ..., cnt(x)` of each factorial, peel off factors
of `2` against the remaining duplicate count `d` before multiplying modulo `m`.

### Complexity

- Time: `O(n)` (build frequencies; sum of factorial loops is `O(n)`)
- Space: `O(n)` for the frequency map
