# C. Triple Removal

[Problem link](https://codeforces.com/problemset/problem/2152/C)

**Contest:** [Squarepoint Challenge (Codeforces Round 1055, Div. 1 + Div. 2)](https://codeforces.com/contest/2152)

time limit per test: 2 seconds

memory limit per test: 1024 megabytes

## Problem

For a binary array, one operation chooses three equal elements at positions
`i < j < k`, removes them, and costs `min(k - j, j - i)` using the current
array indices. The cost of an array is the minimum total cost needed to remove
all its elements; it is `-1` when that cannot be done.

Given a binary array `a`, answer independent range queries: for every `[l, r]`,
find the cost of subarray `a[l..r]`.

## Input

The first line contains the number of test cases `t` (`1 <= t <= 10^4`).

For each test case:

- the first line contains `n` and `q` (`1 <= n, q <= 250000`);
- the next line contains `n` binary values;
- each of the next `q` lines contains a range `l, r` (`1 <= l <= r <= n`).

The sums of `n` and of `q` over all test cases are each at most `250000`.

## Output

For every query, print the corresponding minimum cost, or `-1` if the subarray
cannot be made empty.

## Sample

```text
Input
2
12 4
0 0 1 1 0 1 0 1 0 1 1 0
1 12
2 7
5 10
6 11
6 3
0 0 0 1 1 1
1 3
4 6
1 6

Output
4
2
3
-1
1
1
2
```

## Solution

Each operation removes three equal bits, so a subarray can be emptied only
when its length is a multiple of `3` and its number of `1`s (equivalently
`0`s) is also a multiple of `3`. Otherwise the answer is `-1`.

When those congruences hold, every operation costs at least `1`, and `len/3`
operations are required, so `len/3` is a lower bound. A pair of adjacent
equal bits lets every operation be realized at cost `1` (a tight triple
`000` or `111` can be formed and peeled). A fully alternating window
`0101…` or `1010…` has no adjacent equals: the first triple of equal bits
sits two steps apart and costs `2`, after which the remainder behaves like
the ordinary case. That is exactly one extra unit, so the cost is
`len/3 + 1`.

Precompute a prefix of values and the length of the strict alternating run
ending at each index. A query `[l, r]` is alternating iff that run at `r`
covers the whole window. Each query is then `O(1)` after an `O(n)` scan.

### Correctness sketch

Necessity of the two mod-`3` conditions is immediate from the operation.
Sufficiency of cost `len/3` when a window is not purely alternating follows
from the existence of an adjacent equal pair: removing a tight triple never
creates a new alternating obstruction that would force a later cost-`2`
step. A purely alternating window has no tight triple, so at least one
operation must cost `2`, matching the `+1`. Independent queries share the
same prefixes, so the scan answers all of them.

### Complexity

`O(n + q)` time and `O(n)` memory per test. The sums of `n` and of `q` are
each at most `2.5 · 10^5`.
