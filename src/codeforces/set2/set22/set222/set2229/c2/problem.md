# C2. We Be Flipping (Hard Version)

[Problem link](https://codeforces.com/problemset/problem/2229/C2)

**Contest:** [Spectral::Cup 2026 Round 2 (Codeforces Round 1100, Div. 1 + Div. 2)](https://codeforces.com/contest/2229)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

This is the hard version of the problem. The difference between the versions is that in this version, you must maximise the sum. You can hack only if you solved all versions of this problem.

You have an array `a` of length `n` which consists of non-zero (but possibly negative) integers. You will perform the following operation at most `n` times (possibly none):

- select an index `i` (`1 <= i <= n`) such that `a_i > 0`
- then for each `j` where `1 <= j <= i` do `a_j := -a_j`

Output a valid sequence of operations of length at most `n` which maximises the sum of `a` at the end.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 <= t <= 10^4`). The description of the test cases follows.

The first line of each test case contains an integer `n` (`2 <= n <= 2 * 10^5`) — the length of the array `a`.

The second line of each test case contains `n` integers `a_1, a_2, …, a_n` (`-10^9 <= a_i <= 10^9`, `a_i != 0`).

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output a single integer `k` (`0 <= k <= n`) — the number of operations you will perform.

Then output `k` integers `b_1, …, b_k` where `b_i` is the index you perform the `i`-th operation on.

After performing the operations the sum of `a` should be maximal.

## Solution

If no operation is performed, the final sum is simply `sum(a)`.

Otherwise, let `p` be the largest index on which an operation is performed.
No later operation can change `a_p`, so `a_p` must be initially positive and
becomes `-a_p` after its operation. All elements after `p` are unchanged. The
best possible outcome for every element before `p` is its absolute value, so
every non-empty optimal result has the form:

```text
[|a_1|, |a_2|, ..., |a_{p-1}|, -a_p, a_{p+1}, ..., a_n]
```

for an initially positive `a_p`. Its sum is:

```text
sum(abs(a[1:p])) - a[p] + sum(a[p+1:n])
```

Evaluate this expression for every positive pivot `p` using prefix sums of
absolute values and suffix sums of original values, while also considering the
no-operation sum.

### Construction

For the chosen pivot `p`, first make all elements before it negative. Scan
from `p - 1` down to `1`; whenever the current value is positive, operate on
that index. Each such operation flips the still-unprocessed prefix, so a
single parity flag is enough to determine the current sign without modifying
the array. Finally, operate on `p`. This turns the negative prefix positive
and makes `a_p` negative, producing the target form above. Every selected
index is positive at the moment it is used, and there are at most `p <= n`
operations.

### Correctness

The pivot argument gives an upper bound: after the largest operation, its own
value must be `-a_p`, positions to its right cannot change, and no position to
its left can contribute more than its absolute value. The construction reaches
that upper bound for each eligible pivot. Taking the best pivot (or no
operation) therefore maximises the final sum.

### Complexity

Prefix/suffix evaluation and operation construction both take `O(n)` time and
use `O(n)` auxiliary space.

## Example

### Input

```text
5
5
-1 -2 -3 -5 -4
4
5 7 10 19
5
1 -3 2 -1 10
4
16 -13 -18 -16
11
2 -10 -11 3 -10 15 7 18 16 17 -9
```

### Output

```text
0
0
2
1 3
0
6
6 3 1 5 4 7
```

### Note

In the first test case, no operations are possible.

In the second test case, the sum is already maximal.

In the third test case, operations are made as follows:

- `[1, -3, 2, -1, 10]` with `i = 1` becomes `[-1, -3, 2, -1, 10]`
- `[-1, -3, 2, -1, 10]` with `i = 3` becomes `[1, 3, -2, -1, 10]`

This has sum `11`, which can be proven to be maximal.

## Status

Implemented with the pivot construction above; official samples are covered by
the package tests.
