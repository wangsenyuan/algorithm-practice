# F - Rotated Inversions

[Problem link](https://atcoder.jp/contests/abc396/tasks/abc396_f)

**Contest:** [AtCoder Beginner Contest 396](https://atcoder.jp/contests/abc396)

time limit: 2 sec

memory limit: 1024 MiB

score: 500 points

## Problem

You are given integers `N`, `M` and a length-`N` sequence of non-negative
integers `A = (A_1, A_2, ..., A_N)`.

For each `k = 0, 1, ..., M-1`, define

```text
B_i = (A_i + k) mod M
```

and find the inversion number of `B`.

The inversion number of a sequence is the number of pairs `(i, j)` with
`1 <= i < j <= N` and `B_i > B_j`.

## Constraints

- `1 <= N, M <= 2 * 10^5`
- `0 <= A_i < M`
- All input values are integers

## Input

```text
N M
A_1 A_2 ... A_N
```

## Output

Print `M` lines. The `i`-th line (`1 <= i <= M`) should contain the answer for
`k = i - 1`.

## Samples

### Sample 1

```text
Input
3 3
2 1 0

Output
3
1
1
```

- `k = 0`: `B = (2, 1, 0)`, inversions `3`
- `k = 1`: `B = (0, 2, 1)`, inversions `1`
- `k = 2`: `B = (1, 0, 2)`, inversions `1`

### Sample 2

```text
Input
5 6
5 3 5 0 1

Output
7
3
3
1
1
5
```

### Sample 3

```text
Input
7 7
0 1 2 3 4 5 6

Output
0
6
10
12
12
10
6
```

## Solution Summary

Let `inv(k)` be the inversion number after adding `k` modulo `M` to every
element. We first compute `inv(0)`, then derive every following answer from the
previous one.

### Initial inversion number

Scan `A` from left to right while storing the frequencies of values already
seen in a Fenwick tree.

For the current value `v`, the number of previous values greater than `v` is

```text
number of previous values - number of previous values <= v.
```

Adding this quantity for every position gives `inv(0)` in `O(N log M)` time.

### Change caused by one rotation

When the shift increases from `k` to `k+1`, exactly the elements whose current
value is `M-1` wrap around to `0`. In terms of the original array, these are
the elements with value

```text
v = M - 1 - k.
```

Consider an occurrence of `v` at zero-based position `i`:

- before it wraps, it is greater than every element from another value class;
- after it wraps, it is smaller than every element from another value class.

Therefore, its pairs with earlier positions become inversions, while its pairs
with later positions stop being inversions. If all positions are temporarily
counted, its contribution to the change is

```text
i - (N - 1 - i) = 2*i - N + 1.
```

Pairs of equal values must remain non-inversions because both endpoints wrap
at the same time. They do not affect the total formula: for two equal elements
at positions `i < j`, the formula contributes `-1` from the pair when handling
`i` and `+1` when handling `j`, so the two contributions cancel.

Hence we can precompute

```text
delta[v] = sum(2*i - N + 1) over all positions i with A[i] = v.
```

If the current inversion number is known and value `v` is the next value to
wrap, the next inversion number is simply

```text
inv(next) = inv(current) + delta[v].
```

Starting with `inv(0)`, process `v = M-1, M-2, ..., 0`. Append the current
answer before applying `delta[v]`; the appended values are exactly
`inv(0), inv(1), ..., inv(M-1)`.

### Correctness Proof

We prove that the algorithm outputs the inversion number for every shift.

First, the Fenwick-tree scan computes `inv(0)` correctly. For each right
endpoint `i`, it counts exactly the earlier elements whose values are greater
than `A[i]`. Thus every inversion is counted once, at its right endpoint, and
no non-inversion is counted.

Next, consider increasing a shift by one and let `v` be the original value
class that wraps from `M-1` to `0`. The relative order of two elements outside
this class does not change. Two elements inside this class remain equal, so
their pair is never an inversion. For a pair with exactly one endpoint in this
class, the relative order reverses: an occurrence of `v` gains one inversion
with each earlier outside element and loses one inversion with each later
outside element. Summed over all occurrences, this change is exactly
`delta[v]`; contributions mistakenly attributed to pairs inside the class
cancel in equal and opposite pairs.

Finally, between shifts `k` and `k+1`, the class
`v = M-1-k` is precisely the class that wraps. Processing original values from
`M-1` down to `0` therefore applies the correct change after every answer.
Starting from the correct value `inv(0)`, induction shows that every emitted
value is the correct `inv(k)`.

### Complexity

- Time: `O(N log M + M)`.
- Space: `O(M)`.
