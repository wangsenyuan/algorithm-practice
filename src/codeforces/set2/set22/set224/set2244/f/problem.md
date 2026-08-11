# F. Anya Loves Trees!

[Problem link](https://codeforces.com/problemset/problem/2244/F)

**Contest:** [Codeforces Round 1109 (Div. 3)](https://codeforces.com/contest/2244)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

## Problem

Anya considers a rooted tree with the root at vertex `1`. The leaves of the tree
contain integers from `1` to `k`, where `k` is the number of leaves. Each leaf
contains exactly one number. For every vertex, its children are ordered from
left to right in increasing order of their indices.

Anya noticed that if we list the leaves from left to right, their values form a
sequence. She wants this sequence to become strictly increasing.

To achieve this, Anya can perform the following operation: choose any vertex and
cyclically shift its children to the left. For example, if a vertex has children
in the order `[1, 2, 3]`, after the shift the order becomes `[2, 3, 1]`. This
operation can be applied to any vertices any number of times.

A cyclic left shift moves every child one position left, and the first child
becomes the last.

Help Yura determine whether Anya can make the sequence of integers in the leaves
strictly increasing from left to right using such operations.

## Constraints

- `1 ≤ t ≤ 10^4`
- `1 ≤ n ≤ 2 · 10^5`
- `1 ≤ p_i ≤ n`
- `0 ≤ a_i ≤ n`
- If vertex `i` is not a leaf, then `a_i = 0`
- If vertex `i` is a leaf, then `a_i > 0`
- All positive values `a_i` form a permutation
- Sum of `n` over all test cases does not exceed `2 · 10^5`

## Input

The first line contains a single integer `t` — the number of test cases.

The first line of each test case contains a single integer `n` — the number of
vertices.

The next line contains `n − 1` integers `p_2, p_3, …, p_n`, where `p_i` is the
parent of vertex `i`. The children of each vertex are ordered from left to right
in increasing order of their indices.

The third line contains `n` integers `a_1, a_2, …, a_n`. If vertex `i` is not a
leaf, then `a_i = 0`. If vertex `i` is a leaf, then `a_i > 0` and represents the
number written in that vertex.

## Output

For each test case, output `YES` if it is possible to make the leaf values
strictly increasing from left to right, and `NO` otherwise.

You may output each letter in any case.

## Example

### Input

```text
4
2
1
0 1
4
1 2 2
0 0 2 1
5
5 5 2 1
0 0 1 2 0
4
1 1 1
0 2 1 3
```

### Output

```text
YES
YES
YES
NO
```

## Solution

Goal: leaf order must be exactly `1…k`. Operations only cyclically rotate
children, so every subtree always occupies a contiguous segment of the leaf
sequence. In a YES configuration that segment must be a value interval
`[L, R]` in increasing order — so **every** node ends with `|dp[u]| = 1`.

### 1. Model a subtree as intervals

For node `u`, `dp[u]` is a list of pairs `(L, R)` — contiguous increasing
value blocks among leaves under `u` (after chosen shifts).

- Leaf `u` with value `a_u`: `dp[u] = [(a_u, a_u)]`.
- Internal node: recurse on children, concatenate their (single) intervals.

### 2. Cyclic shift = rotate to the minimum start

Concatenate children’s intervals, find the block with smallest `L`, and rotate
so it comes first. If the subtree is arrangeable as `[L, R]`, that sorted order
must start with `L = min`, so this is the only candidate rotation.

### 3. Merge consecutive blocks

Walk the rotated list and merge:

| relation to previous `(L', R')` | action |
|---|---|
| `L == R' + 1` | extend previous block to `R` |
| `L > R' + 1` | hole inside this position block → `NO` |
| `L ≤ R'` | overlap / disorder → `NO` |

After merging, require `|dp[u]| = 1` (one interval `[L, R]`). A second residual
run would mean a hole in a contiguous leaf segment, which can never appear
inside the global order `1…k`. (The code’s `|dp[u]| ≤ 2` is a looser early
filter; the tight invariant for YES is `|dp[u]| = 1` at every node.)

### 4. Root check

Answer `YES` iff every node merged successfully and the root’s unique interval
is `[1, k]`.

### One-liner

Bottom-up: rotate each node’s child blocks to start at the subtree minimum,
glue abutting ranges; every node (including the root) must become a single
interval.
