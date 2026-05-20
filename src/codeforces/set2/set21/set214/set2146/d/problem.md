# D2. Max Sum OR (Hard Version)

https://codeforces.com/problemset/problem/2146/D2

**Time Limit:** 2 seconds

**Memory Limit:** 256 megabytes

## Problem

This is the hard version of the problem. The difference is that here
`0 <= l <= r < 2^30`.

You are given two integers `l` and `r` with `l <= r`.

Let:

```text
n = r - l + 1
```

Create two arrays `a` and `b`, both of length `n`. Initially, both arrays are:

```text
[l, l + 1, ..., r]
```

You may reorder array `a` arbitrarily. Array `b` remains in increasing order.

Maximize:

```text
sum(a_i | b_i)
```

where `|` is the bitwise OR operation.

You also need to output one possible reordered array `a` that achieves the
maximum value.


## Solution

For each fixed bit, `x | y` has that bit equal to `1` if at least one of `x`
and `y` has it. Therefore, when pairing numbers from `[l, r]`, we want the two
numbers in each pair to have opposite bits as often as possible.

The current implementation uses a binary trie as a greedy matcher.

First, insert every number from `l` to `r` into a trie. Each node stores how
many unused numbers remain in its subtree. The code stores bits from low to
high. When matching a number `i`, try to choose an unused number with the
opposite bit at the current trie level. If that side is empty, choose the same
bit. After choosing a branch, decrement the count in that subtree.

This constructs a permutation `a`, because every inserted number is removed
exactly once. For each `i` in `[l, r]`, the chosen number becomes `a[i-l]`.
Then the answer is computed directly:

```text
sum += i | a[i-l]
```

The package also keeps `solve1`, the earlier divide-and-conquer construction.
It works by splitting an interval `[lo, hi]` at the highest bit boundary. When
both bit values exist, it pairs values symmetrically across the boundary:

```text
mid     <-> mid+1
mid-1   <-> mid+2
...
```

Those pairs make the current high bit contribute to both OR values, and the
leftover parts are solved recursively on lower bits.

Both views are trying to maximize opposite-bit pairings. The trie version is the
wired `solve` function, while `solve1` is a useful alternative explanation of
why pairing across bit boundaries is optimal.

The complexity is `O(n log V)` per test case, where
`n = r - l + 1` and `V < 2^30`.

