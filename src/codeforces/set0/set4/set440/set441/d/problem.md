# Permutation Transformation

## Problem Description

A **permutation** `p` of length `n` is a sequence of distinct integers `p[1], p[2], ..., p[n]` (1 ≤ p[i] ≤ n). A permutation is an **identity permutation** if for any `i` the following equation holds: `p[i] = i`.

A **swap** `(i, j)` is the operation that swaps elements `p[i]` and `p[j]` in the permutation. Let's assume that $f(p)$ is the minimum number of swaps needed to make the permutation `p` an identity permutation.

Valera wonders how he can transform permutation `p` into any permutation `q`, such that $f(q) = m$, using the minimum number of swaps. Help him do that.

## Input

- **First line**: Integer `n` (1 ≤ n ≤ 3000) — the length of permutation `p`
- **Second line**: `n` distinct integers `p[1], p[2], ..., p[n]` (1 ≤ p[i] ≤ n) — Valera's initial permutation
- **Third line**: Integer `m` (0 ≤ m < n)

## Output

- **First line**: Integer `k` — the minimum number of swaps
- **Second line**: `2k` integers `x[1], x[2], ..., x[2k]` — the description of the swap sequence

The printed numbers show that you need to consecutively make swaps `(x[1], x[2]), (x[3], x[4]), ..., (x[2k-1], x[2k])`.

**Note**: If there are multiple sequences of swaps with the minimum length, print the lexicographically minimum one.

## Examples

### Example 1

**Input:**
```
5
1 2 3 4 5
2
```

**Output:**
```
2
1 2 1 3
```

### Example 2

**Input:**
```
5
2 1 4 5 3
2
```

**Output:**
```
1
1 2
```

## Note

A sequence `x[1], x[2], ..., x[s]` is **lexicographically smaller** than sequence `y[1], y[2], ..., y[s]` if there exists an integer `r` (1 ≤ r ≤ s) such that:

- `x[1] = y[1], x[2] = y[2], ..., x[r-1] = y[r-1]`
- and `x[r] < y[r]`

## Solution

In this task you should represent permutation as graph with `n` vertexes, and from every vertex `i` exists exactly one edge to vertex `p[i]`. It's easy to understand that such graph consists of simple cycles only.

If we make swap `(i, j)`, edges `(i, p[i])` and `(j, p[j])` will become edges `(i, p[j])` and `(j, p[i])` respectively. Then if `i` and `j` are in the same cycle, this cycle will break:

```
Before: i → p[i] → ... → j → p[j] → ...
After:  i → p[j] → ... → j → p[i] → ...
```

But if they are in different cycles, these cycles will merge into one:

```
Before: i → p[i] → ... → i  (cycle 1)
        j → p[j] → ... → j  (cycle 2)
After:  i → p[j] → ... → j → p[i] → ... → i  (merged cycle)
```

This means that every swap operation increases number of cycles by one, or decreases it by one.

Assuming all above, to get permutation `q` from permutation `p`, we need to increase (or decrease) number of cycles in `p` to $n - m$. Let $c$ be the number of cycles in `p`. Then $k$ always equals $|(n - m) - c|$.

For satisfying lexicographical minimality we will review three cases:

### Case 1: $n - m < c$

It's easy to understand that in this case you must decrease cycles number by merging cycles one by one with cycle containing vertex 1. This way every swap has form `(1, v)`, where $v > 1$. Because every cycle vertex is bigger than previous cycle vertex, this case can be solved with $O(n)$.

### Case 2: $n - m > c$

In this case you should break cycle for every vertex, making swap with smallest possible vertex (it should be in this cycle too). This could be done if represent cycle by line. As soon as every cycle is broken with linear asymptotics, this case solution works with $O(n^2)$.

**Bonus**: this way of representing cycle lets us optimize solution to $O(n \log n)$ asymptotics, you may think how.

### Case 3: $n - m = c$

Because in this case $k = 0$, there is nothing need to be swapped.