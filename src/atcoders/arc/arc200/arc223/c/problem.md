# C - Whole Product of Pairwise Distances

[Problem link](https://atcoder.jp/contests/arc223/tasks/arc223_c)

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 500 points

## Problem Statement

You are given a sequence `A` of `N` positive integers.
Find the remainder when

```
∏_{1 ≤ i < j ≤ N} |A_i - A_j|
```

is divided by `N`.

Solve `T` test cases per input.

## Constraints

- `1 <= T <= 10^5`
- `2 <= N <= 2 * 10^5`
- `1 <= A_i <= 10^9`
- The sum of `N` over all test cases is at most `2 * 10^5`
- All input values are integers

## Input

```
T
case_1
case_2
...
case_T
```

Each test case:

```
N
A_1 A_2 ... A_N
```

## Output

Output `T` lines. The `t`-th line should contain the answer for the `t`-th
test case.

## Samples

### Sample 1

Input:

```
1
3
1 9 5
```

Output:

```
2
```

`|1-9| × |1-5| × |9-5| = 8 × 4 × 4 = 128`, and `128 mod 3 = 2`.

### Sample 2

Input:

```
1
3
2 2 3
```

Output:

```
0
```

### Sample 3

Input:

```
1
5
11 33 22 55 44
```

Output:

```
3
```

## Solution

Let

```text
P = product(|A[i] - A[j]|) for all i < j.
```

We need to compute `P modulo N`.

### Duplicate remainders make the answer zero

For every pair:

```text
A[i] - A[j] = (A[i] mod N) - (A[j] mod N)  (mod N).
```

If two values have the same remainder modulo `N`, their difference is
divisible by `N`. That factor makes the entire product divisible by `N`, so the
answer is `0`.

Otherwise, there are `N` distinct remainders, each between `0` and `N-1`.
Therefore they form a permutation of:

```text
0, 1, ..., N-1.
```

This is the duplicate check performed while constructing `b`.

### Sort to remove the absolute values

Sort the original values and call the sorted sequence `c`:

```text
c[0] < c[1] < ... < c[N-1].
```

The values are indeed distinct in the nonzero case, because equal values have
the same remainder. After sorting, every pairwise difference is positive:

```text
P = product(c[j] - c[i]) for all i < j.
```

Define:

```text
b[i] = c[i] mod N.
```

Modulo `N`, we can replace every `c[i]` by `b[i]`:

```text
P = product(b[j] - b[i])  (mod N).
```

The sequence `b` is a permutation of `0, 1, ..., N-1`.

### Product for the identity permutation

First consider the identity order:

```text
b = [0, 1, ..., N-1].
```

Its product is:

```text
D = product(j - i) for all 0 <= i < j < N.
```

For a fixed right endpoint `j`, its contribution is:

```text
(j-0)(j-1)...(j-(j-1)) = j!.
```

Therefore:

```text
D = 1! * 2! * ... * (N-1)!  (mod N).
```

The code computes the factorials and their cumulative product together:

```go
res, fact := 1, 1
for i := 1; i < n; i++ {
	fact = fact * i % n
	res = res * fact % n
}
```

### Sign of the residue permutation

The expression

```text
product(b[j] - b[i])
```

is a Vandermonde product. Swapping two elements of `b` changes its sign.
Consequently, for an arbitrary permutation `b`:

```text
P = sign(b) * D  (mod N).
```

It remains to determine whether `b` is an even or odd permutation.

If a permutation of `N` elements contains `cycles` disjoint cycles, it can be
formed using:

```text
N - cycles
```

swaps. A cycle of length `L` contributes `L-1` swaps, and summing this over all
cycles gives the formula.

Thus:

- if `N-cycles` is even, the sign is positive and the answer is `D`;
- if `N-cycles` is odd, the sign is negative and the answer is `-D mod N`.

The code converts the negative value into a nonnegative remainder with:

```go
res = (n - res) % n
```

### Example

For the first sample, sorting gives:

```text
c = [1, 5, 9]
b = [1, 2, 0].
```

The permutation `b` is one cycle `0 -> 1 -> 2 -> 0`, so:

```text
N - cycles = 3 - 1 = 2,
```

which is even. Its sign is positive. The fixed product is:

```text
D = 1! * 2! = 2 mod 3,
```

so the answer is `2`.

### Correctness Proof

If two elements have equal remainders modulo `N`, their difference is
divisible by `N`, so the algorithm correctly returns `0`.

Otherwise, the sorted remainder sequence `b` is a permutation of
`0, 1, ..., N-1`. Sorting makes every original distance equal to
`c[j]-c[i]`, and reducing each value modulo `N` transforms the required product
into the Vandermonde product of `b` without changing its remainder.

For the identity permutation, this product equals
`1! * 2! * ... * (N-1)!`. Every permutation changes the Vandermonde product by
exactly its permutation sign. The cycle decomposition computes that sign as
`(-1)^(N-cycles)`. Therefore the algorithm returns exactly the required
pairwise-distance product modulo `N`.

### Complexity

Sorting dominates the running time. All remaining passes are linear.

- Time: `O(N log N)` per test case
- Space: `O(N)`
