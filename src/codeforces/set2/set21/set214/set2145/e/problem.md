# E. Predicting Popularity

[Problem link](https://codeforces.com/problemset/problem/2145/E)

time limit per test: 3 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Berflix has `n` users. User `i` has action preference `a_i` and drama preference
`d_i`. The movie has action level `a_c` and drama level `d_r`.

If both `a_i <= a_c` and `d_i <= d_r`, user `i` will definitely watch the movie.

Otherwise user `i` may still watch if enough others have already watched. Let `p`
be the number of viewers so far (`p = 0` initially). User `i` finds the movie
suitable when:

```text
max(a_i - a_c, 0) + max(d_i - d_r, 0) <= p
```

As long as some not-yet-watching user finds the movie suitable, they watch it,
increasing `p` by 1. This repeats until everyone has watched or no remaining user
finds it suitable.

There are `m` updates; each changes `(a_k, d_k)` for one user. After each update,
output the final total number of viewers.

## Input

- Line 1: `a_c`, `d_r` (`1 <= a_c, d_r <= 10^6`)
- Line 2: `n` (`1 <= n <= 5 * 10^5`)
- Line 3: `a_1, …, a_n` (`1 <= a_i <= 10^6`)
- Line 4: `d_1, …, d_n` (`1 <= d_i <= 10^6`)
- Line 5: `m` (`1 <= m <= 3 * 10^5`)
- Next `m` lines: `k_j`, `na_j`, `nd_j` — set `a_{k_j} = na_j`, `d_{k_j} = nd_j`
  (`1 <= k_j <= n`; `1 <= na_j, nd_j <= 10^6`)

## Output

Print `m` lines. The `j`-th line is the total number of viewers after the `j`-th
update.

## Example

### Input

```text
20 25
4
1 22 1 30
1 22 50 30
5
3 1 25
2 23 22
4 10 27
1 21 21
3 20 26
```

### Output

```text
3
2
4
4
0
```

## Note

After the first update, users 1 and 3 watch immediately (`p = 2`), then user 2
becomes suitable and watches (`p = 3`). User 4 still has
`max(30 - 20, 0) + max(30 - 25, 0) = 15 > 3`, so they do not watch.


## Solution

For each user define the required number of existing viewers:

```text
w_i = max(a_i - a_c, 0) + max(d_i - d_r, 0)
```

Users with `w_i = 0` watch immediately. Let their number be `p0`. Sort all other
currently active users by `w`. If their costs are
`w_0 <= w_1 <= ...`, then user `k` can watch after the previous `k` users when:

```text
w_k <= p0 + k
```

Equivalently, `w_k - k <= p0`. Therefore the answer is `p0` plus the longest
prefix whose maximum `w_k - k` does not exceed `p0`.

### Offline positions

Read all updates first. Collect every positive cost that can occur: costs from
the initial state and costs introduced by updates. Sort pairs `(cost, event)`;
the unique event index distinguishes equal costs. Each possible positive state
now owns one fixed position in this order.

Only some positions are active at a given time. For an active position `j`, the
segment-tree leaf stores:

```text
cost[j] - number of active positive-cost positions before j
```

Activating position `j` sets its cost and subtracts one from every later
position. Deactivating it clears the leaf and adds one to every later position.
Both operations are a point assignment plus a suffix range addition.

### Segment tree

Each node stores:

- `val`: the maximum adjusted cost in its interval;
- `cnt`: the number of active positions in its interval;
- `lazy`: a pending range addition.

To answer an update, search for the first active position whose adjusted cost is
greater than `p0`. If a node's maximum is at most `p0`, its entire active count
belongs to the valid prefix. Otherwise descend left first, adding the left
count only when the whole left interval is valid.

For every update:

1. Remove the user's old state, updating either `p0` or its positive-cost slot.
2. Insert the new state in the same way.
3. Query the valid positive-cost prefix and add `p0`.

The segment tree is created with at least one leaf so the case with no positive
costs is also valid.

### Complexity

Let `K <= n + m` be the number of collected positive-cost states. Sorting costs
takes `O(K log K)`. Each insertion, deletion, and answer query takes `O(log K)`,
so the total complexity is `O((n + m) log(n + m))` with `O(n + m)` memory.
