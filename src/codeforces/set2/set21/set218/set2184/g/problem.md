# G. Nastiness of Segments

https://codeforces.com/problemset/problem/2184/G

## Problem

There are `n` blocks numbered from `1` to `n`. Initially, block `i` contains an
integer `a_i`, and the blocks are arranged in increasing order of their numbers.

For a segment `[l, r]`, an integer `d` is called nasty if:

```text
0 <= d <= r - l
min(a_l, a_{l+1}, ..., a_{l+d}) = d
```

The nastiness of `[l, r]` is the number of nasty values of `d` for that segment.

You need to process `q` operations:

1. `1 i x` -- change `a_i` to `x`.
2. `2 l r` -- output the nastiness of segment `[l, r]`.

## Input

The first line contains `t` (`1 <= t <= 10^4`) -- the number of test cases.

For each test case:

* The first line contains `n` and `q`
  (`1 <= n, q <= 2 * 10^5`).
* The second line contains `n` integers `a_1, a_2, ..., a_n`
  (`1 <= a_i <= 2 * 10^5`).
* The next `q` lines contain one operation each.
  * If `idx = 1`, the line is `1 i x`
    (`1 <= i <= n`, `1 <= x <= 2 * 10^5`).
  * If `idx = 2`, the line is `2 l r` (`1 <= l <= r <= n`).

It is guaranteed that the sum of `n` over all test cases is at most
`2 * 10^5`, and the sum of `q` over all test cases is at most `2 * 10^5`.

## Output

For each operation of type `2`, output the nastiness of the requested segment.

## Example

```text
Input
1
5 5
1 2 3 4 5
2 1 5
1 1 5
1 2 5
1 3 1
2 1 5

Output
1
0
```


## Solution summary (from `solution.go`)

### Nastiness is 0 or 1

For a fixed segment `[l, r]`, `d` is **nasty** when `min(a_l, …, a_{l+d}) = d` (with `0 ≤ d ≤ r − l`).

As you extend the prefix, the running minimum never increases. So at most **one** value `d` can satisfy the equality — the answer to a type-2 query is always **`0` or `1`**.

Equivalently (0-based indices): nastiness is **`1`** iff there exists an index `i` with `l ≤ i ≤ r` such that

```text
min(a[l], …, a[i]) = i − l
```

If no such `i` exists in `[l, r]`, the answer is **`0`**.

### Data structure

Build a **segment tree** on `a` that supports:

- **Point update** — operation `1 i x`.
- **Range minimum** on `[L, R]` — used inside each query.

### Answering `2 l r` — `play(l, r)`

Binary search on the right end `i` of the prefix `[l, i]` (search range `[l, n)`):

- Let `v = min(a[l], …, a[mid])`.
- If `v ≤ mid − l`, move left (`hi = mid`); else move right (`lo = mid + 1`).

This finds the **smallest** `i` such that `min(a[l], …, a[i]) ≤ i − l`. Then:

- If `i ≤ r` (in 0-based: `hi ≤ r`) **and** `min(a[l], …, a[i]) = i − l`, output **`1`**;
- otherwise output **`0`**.

So a query is one binary search (`O(log n)`) plus `O(log n)` segment-tree queries per step.

### Processing operations

Loop over queries:

- Type `1`: `update` position `i − 1` to `x`.
- Type `2`: append `play(l − 1, r − 1)` to the answer list.

### Complexity

- Build: `O(n)`.
- Each update / range-min query: `O(log n)`.
- Each type-2 query: `O(log² n)` from binary search on the tree.
- Memory: `O(n)` for the segment tree.

### Ideas (original notes)

1. Nasty values come from prefixes where `min` equals the offset `d`.
2. Extending a prefix can only lower (or keep) the minimum, so at most one nasty `d` per segment.
3. For query `[l, r]`, answer is `1` or `0` — find `i` with `min(a[l], …, a[i]) = i − l` inside the range.