# D. The Child and Sequence

[Problem link](https://codeforces.com/problemset/problem/438/D)

**Contest:** [Codeforces Round 251 (Div. 2)](https://codeforces.com/contest/438)

time limit per test: 4 seconds

memory limit per test: 256 megabytes

## Problem

Given an integer array `a`, process `m` operations:

1. Print the sum of `a[l..r]`.
2. Replace every `a[i]` in `a[l..r]` by `a[i] % x`.
3. Set `a[k]` to `x`.

## Constraints

- `1 <= n, m <= 10^5`
- `1 <= a[i] <= 10^9`
- `1 <= l <= r <= n`
- `1 <= x <= 10^9`

## Input

The first line contains `n` and `m`.

The second line contains `n` integers `a[1..n]`.

Each of the next `m` lines describes one operation:

- `1 l r`
- `2 l r x`
- `3 k x`

## Output

For each type-1 operation, print the sum of the specified range.

## Solution

Use a segment tree. Each node stores:

- The sum of its segment, for type-1 queries.
- The maximum value in its segment, for pruning type-2 updates.

For an update `a[i] %= x`, any value smaller than `x` stays unchanged. Thus,
if a node's maximum is less than `x`, the entire node can be skipped. Otherwise
descend to its children; at a leaf, apply the modulo and update both the sum and
maximum to the new value. The strict comparison matters: a value equal to `x`
becomes zero, so `max == x` cannot be pruned.

After processing children, recompute the parent sum and maximum. Point
assignment updates the leaf and its ancestors in the usual way.

### Why modulo updates are efficient

A modulo operation can change many elements, so one update is not necessarily
`O(log n)`. However, each element's value drops by at least half whenever a
modulo operation changes it. For current value `v`:

- If `x <= v/2`, then `v % x < x <= v/2`.
- If `v/2 < x <= v`, then `v % x = v - x < v/2`.

So each element changes at most `O(log A)` times between assignments, where
`A` is the maximum value (`A <= 10^9`). A point assignment can reset a value,
so over the whole input there are at most `O((n + m) log A)` changed-element
events. Each changed leaf causes only its `O(log n)` path to be traversed;
range boundaries also add `O(log n)` overhead per operation. The total
complexity is `O((n + m) log A log n)`, with `O(n)` space.

Modulo operations must still be applied in order: in general, `(v % a) % b`
does not equal `v % b` when `a > b`. The segment tree does not reorder or
combine moduli; it skips only values that provably remain unchanged.

## Sample

```text
Input
5 5
1 2 3 4 5
2 3 5 4
3 3 5
1 2 5
2 1 3 3
1 1 3

Output
8
5
```

