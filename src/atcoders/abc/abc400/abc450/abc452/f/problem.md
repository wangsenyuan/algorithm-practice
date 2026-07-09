# F - Interval Inversion Count

[Problem link](https://atcoder.jp/contests/abc452/tasks/abc452_f)

**Contest:** [AtCoder Beginner Contest 452](https://atcoder.jp/contests/abc452)

time limit: 2 sec

memory limit: 1024 MiB

score: 500 points

You are given a positive integer `N` and a permutation `P = (P_1, P_2, ..., P_N)` of
`(1, 2, ..., N)`.

You are also given an integer `K`. Find the number of pairs of integers `(l, r)`
satisfying both of the following:

- `1 <= l <= r <= N`
- The inversion number of the sequence `(P_l, P_{l+1}, ..., P_r)` equals `K`

## Constraints

- `1 <= N <= 5 * 10^5`
- `0 <= K <= N*(N-1)/2`
- `1 <= P_i <= N`
- `P_i != P_j` for `i != j`
- All input values are integers

## Input

```text
N K
P_1 P_2 ... P_N
```

## Output

Print the number of pairs `(l, r)` that satisfy the conditions.

## Sample Input 1

```text
7 3
6 3 2 1 7 5 4
```

## Sample Output 1

```text
5
```

For example, the inversion number of `(P_1, P_2, P_3) = (6, 3, 2)` is `3`, so
`(l, r) = (1, 3)` is valid. The other valid pairs are `(2, 4)`, `(2, 5)`, `(4, 7)`,
and `(5, 7)`.

## Sample Input 2

```text
4 1
1 2 3 4
```

## Sample Output 2

```text
0
```

Every contiguous subsequence of `P` has inversion number `0`, so no pair works.

## Sample Input 3

```text
25 18
14 19 24 8 12 11 6 5 3 13 22 15 17 2 9 4 7 18 10 25 23 16 1 20 21
```

## Sample Output 3

```text
3
```

## Solution

Count subarrays with **exactly `K` inversions** using two sliding windows and Fenwick trees.

### Key fact

For a fixed right endpoint `r`, as the left endpoint `l` moves right, the inversion
count of `[l, r]` is **non-increasing** (removing the leftmost element only deletes
inversions). So for each `r`, the valid `l` form a contiguous range.

### Two pointers

Maintain two left borders for the same `r`:

| pointer | meaning after update |
| --- | --- |
| `l1` | smallest index with `inv([l1, r]) < K` |
| `l2` | smallest index with `inv([l2, r]) <= K` |

Then:

- `l < l2` → inversions `> K`
- `l2 <= l < l1` → inversions `= K`
- `l >= l1` → inversions `< K`

For each `r`, add `l1 - l2` when the window with left `l2` has exactly `K` inversions
(if the count jumps over `K`, that range is empty).

### Maintaining inversions with BIT

When appending `p[r]` into a window:

```text
new inversions = (#elements already in window) - (#elements <= p[r])
```

When deleting `p[l]` from the left:

```text
removed inversions = (#remaining elements < p[l])
```

Each window has its own Fenwick tree (`s1` / `s2`) storing the multiset of values
currently inside.

### Complexity

Each index enters/leaves each window at most once, so the time complexity is
`O(N log N)`, which fits `N <= 5 * 10^5`. The answer can be up to about `N^2 / 2`, so
use a 64-bit integer for the result.
