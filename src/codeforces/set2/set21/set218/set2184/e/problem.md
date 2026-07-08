# E. Exquisite Array

[Problem link](https://codeforces.com/problemset/problem/2184/E)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

An array is **k-exquisite** if it has at least two elements and any two adjacent
elements differ by at least `k`.

You are given a permutation `p` of length `n`. For each `k` from `1` to `n - 1`,
count how many **k-exquisite subarrays** of `p` exist.

A subarray is a sequence of one or more consecutive elements of the array.

## Input

The first line contains `t` (`1 <= t <= 25000`) — the number of test cases.

For each test case:

- The first line contains `n` (`2 <= n <= 10^5`).
- The second line contains `n` distinct integers `p_1, ..., p_n`
  (`1 <= p_i <= n`) — the permutation.

The sum of `n` over all test cases is at most `2 * 10^5`.

## Output

For each test case, print `n - 1` integers on one line: for each `k` from `1` to
`n - 1`, the number of k-exquisite subarrays.

## Example

```text
Input
3
5
5 1 4 2 3
3
3 2 1
4
3 1 2 4

Output
10 6 3 1
3 0
6 2 0
```


## Solution

For a fixed `k`, define an adjacent gap `i` as active if:

```text
abs(p[i+1] - p[i]) >= k
```

A subarray with at least two elements is `k`-exquisite iff every gap inside it is active.
So if there is a consecutive run of `m` active gaps, it contributes:

```text
1 + 2 + ... + m = m * (m + 1) / 2
```

subarrays. The answer for `k` is the sum of this value over all active-gap components.

When `k` decreases, active gaps are only added, never removed. Therefore process `k` from
`n-1` down to `1`.

Precompute `todo[d]`: all gap positions `i` where `abs(p[i+1] - p[i]) = d`.

Maintain connected components of active gaps. The code represents a component by its left
and right endpoints in vertex indices:

- a component `[l, r]` contains active gaps `l, l+1, ..., r-1`;
- its number of active gaps is `r-l`;
- its contribution is `count(l, r) = (r-l) * (r-l+1) / 2`.

Two DSU-like parent arrays are used:

- `L.Find(x)` returns the left endpoint of the active component containing vertex `x`.
- `R.Find(x)` returns the right endpoint of the active component containing vertex `x`.

When adding a new active gap between `l` and `r = l+1`:

1. Find the existing component ending at `l`: `[l1, l]`.
2. Find the existing component starting at `r`: `[r, r1]`.
3. Remove their old contributions from the running total.
4. Add the merged component contribution `count(l1, r1)`.
5. Link `l -> r` in the right DSU and `r -> l` in the left DSU.

This running total `sum` is the full answer for the current `k`, including components
that were created by larger gaps earlier. This is necessary because a lower threshold may
activate one component while older active components remain separate.

After processing all gaps with value exactly `k`, set `ans[k] = sum`.

The total number of gap insertions is `n-1`, and every insertion does only inverse-Ackermann
DSU work. The time complexity is `O(n alpha(n))` per test case, and the memory complexity
is `O(n)`.
