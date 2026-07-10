# G. Greedy Subsequences

[Problem link](https://codeforces.com/problemset/problem/1132/G)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

For an array `c`, a **greedy subsequence** is a sequence of indices
`p_1 < p_2 < ... < p_l` such that for each `i`, `p_{i+1}` is the **smallest**
index greater than `p_i` with `c[p_{i+1}] > c[p_i]`.

You are given an array `a_1, ..., a_n`. For each contiguous subsegment of length
`k`, find the length of its **longest** greedy subsequence.

## Constraints

- `1 <= k <= n <= 10^6`
- `1 <= a_i <= n`

## Input

```text
n k
a_1 a_2 ... a_n
```

## Output

Print `n - k + 1` integers — the answers for subsegments `a[1..k]`,
`a[2..k+1]`, ..., `a[n-k+1..n]`.

## Example 1

```text
Input
6 4
1 5 2 5 3 6

Output
2 2 3
```

- `[1, 5, 2, 5]` — longest greedy length `2` (e.g. `1,5` or `2,5`)
- `[5, 2, 5, 3]` — length `2` (`2,5`)
- `[2, 5, 3, 6]` — length `3` (`2,5,6`)

## Example 2

```text
Input
7 6
4 5 2 5 3 6 6

Output
3 3
```

- `[4, 5, 2, 5, 3, 6]` — length `3`
- `[5, 2, 5, 3, 6, 6]` — length `3`


## Solution

For every position `i`, define `fa[i]` as the first index to the right with
`a[fa[i]] > a[i]`. If no such index exists, attach `i` to a virtual root `n`.
This is exactly the next step of a greedy subsequence starting at `i`, as long
as `fa[i]` still lies inside the current window.

The `fa` forest can be built with a monotonic decreasing stack. When a new value
`a[i]` is larger than stack top `a[v]`, `i` is the first greater position for
`v`, so `fa[v] = i`. Remaining stack elements after the scan are attached to
the virtual root.

After that, turn this forest into Euler order. The important property is that
every subtree is a contiguous interval. For a fixed window `[l, l+k)`, let
`dp[x]` be the greedy subsequence length starting from `x`, but only while
following `fa` edges that remain inside the window.

Process windows from right to left:

1. When position `r = l+k` leaves the window, every active node in the subtree
   of `r` loses one valid ancestor step, so subtract `1` from the whole Euler
   interval of `r`.
2. When inserting `l`, if `fa[l] < l+k`, then `dp[l] = dp[fa[l]] + 1`;
   otherwise `dp[l] = 1`.
3. The answer for the window is the maximum `dp` value among all active nodes.

A lazy segment tree over Euler order supports subtree add, point query, point
assignment, and global maximum query.

### Correctness notes

The hidden pitfall is that the range subtract for an exiting right endpoint can
touch descendants that have not been inserted yet. Therefore insertion must be a
point assignment, not a point addition. Otherwise a stale negative lazy value may
remain on the leaf and break cases such as:

```text
2 1
1 2
```

The correct answer is `1 1`. With point addition, the first window can become
`0`; with point assignment, the newly inserted node receives its real
`dp[l]` value.

The forest invariant proves the DP:

- If `fa[x]` is inside the window, the greedy subsequence starting at `x` must
  next go to `fa[x]`, so its length is `1 + dp[fa[x]]`.
- If `fa[x]` is outside the window or is the virtual root, the subsequence stops
  at `x`, so its length is `1`.
- Taking the maximum over inserted positions gives the longest greedy
  subsequence in the current window.

### Memory and implementation notes

The constraints allow `n = 10^6`, so the representation matters.

- Do not store children as `[][]int`: one million slice headers alone cost too
  much memory. The implementation uses `firstChild` and `nextSibling` arrays
  instead.
- Avoid recursive DFS because the tree can be a chain. Euler order and subtree
  sizes are computed iteratively.
- Segment tree values are at most `k`, so `int32` is enough for `val` and
  `lazy`, cutting the two largest arrays in half.
- Print answers with a buffered writer. Building `fmt.Sprintf("%v", res)` first
  creates an extra large string and can push memory over the limit.

### Complexity

The monotonic stack and Euler traversal are `O(n)`. Each index is inserted once,
and each window shift performs `O(log n)` segment tree work, so the total time is
`O(n log n)`.

The memory usage is `O(n)`.
