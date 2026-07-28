# C. Village Guilds

[Problem link](https://codeforces.com/problemset/problem/2238/C)

**Contest:** [Codeforces Round 1106 (Div. 2)](https://codeforces.com/contest/2238)

## Problem

There are `n` houses connected by bidirectional paths. The graph is a rooted
tree rooted at vertex `1` (the town hall).

For an arbitrary house `v` and a non-negative integer `h`, consider the set of
houses that lie in the subtree of `v` at distance exactly `h` from `v`. Such a
set is called a **guild**. For `h = 0`, the guild is just `{v}`.

Two guilds are different if there exists a house that belongs to one but not
the other. Count the number of different non-empty guilds in the tree.

## Constraints

- `1 <= t <= 10^4`
- `2 <= n <= 2 · 10^5`
- `1 <= p_i < i` for each `i = 2..n`
- Sum of `n` over all test cases `<= 2 · 10^5`

## Input

The first line contains `t` — the number of test cases.

For each test case:

- The first line contains `n` — the number of houses.
- The second line contains `n - 1` integers `p_2, p_3, ..., p_n`, where `p_i`
  is the parent of house `i`.

## Output

For each test case, output a single integer — the number of different guilds.

## Sample 1

```text
Input
5
1 2 3 4

Output
5
```

### Note

There are 5 guilds, each consisting of a single house.

## Sample 2

```text
Input
3
1 1

Output
4
```

### Note

Besides the single-house guilds, there is also a guild consisting of vertices
`2` and `3` (for example from `v = 1`, `h = 1`).

## Sample 3

```text
Input
7
1 2 1 3 5 5

Output
9
```

![Sample 3 tree](sample3-tree.png)

## Sample 4

```text
Input
10
1 1 3 2 2 4 4 4 3

Output
15
```

## Sample 5

```text
Input
15
1 2 1 3 3 4 3 7 3 10 6 7 1 9

Output
22
```

## ideas

1. 每个点的 `h = 0` guild 都是单点集 `{v}`，彼此不同，先贡献 `n`。
2. 若 `u` 只有 0/1 个子树，子树里已经出现过的多层 guild 会沿唯一最深链
   “上移”复现，不会在 `u` 产生新的非单点 guild。
3. 若 `u` 至少两个子树，设子树高度（到叶子的最长距离）的最大、次大值为
   `best1 >= best2`，则距离 `1 .. best2+1` 的 guild 是新的（至少两条链都能
   支撑到该深度，集合由多条分支拼成，不会等于某个纯子树 guild）。
4. 因此额外贡献为 `best2 + 1`；答案
   `n + Σ_u (best2(u) + 1)`（仅对有至少两个儿子的 `u` 求和）。

## summary

### Goal

Count distinct non-empty **guilds**: for each vertex `v` and `h >= 0`, the set of
nodes in `v`'s subtree at distance exactly `h` from `v` (so `h = 0` gives `{v}`).

### Key observation

- All `n` singletons `{v}` are distinct → base answer `n`.
- Along a single root-to-leaf chain, a deeper guild often equals a shallower one
  taken from a descendant, so a node with **at most one child** adds **no** new
  multi-node guild.
- New multi-node guilds appear only when a node has **≥ 2** child subtrees: the
  second-longest subtree height `best2` limits how far you can go while still
  involving at least two branches. That yields `best2 + 1` new guilds at that
  node (distances `1, 2, ..., best2 + 1`).

### Formula

Let `height[u]` = longest distance from `u` down to a leaf in its subtree
(`height[leaf] = 0`). For each `u`, let `best1 >= best2` be the two largest
`height[child]` (missing values treated as `-1`). Then

```text
answer = n + Σ_u [best2(u) + 1]     // only where best2 >= 0
```

### Code (`solve`)

1. Build children lists from parents `p` (0-based).
2. DFS from the root; each call returns the largest child-height seen as
   `first` (with leaf sentinel `-1`, so `dfs(child) + 1` equals that child's
   `height`).
3. Track `first` / `second` (max / second-max of those values).
4. Always `res++` (the singleton); then `res += second + 1` (extra guilds;
   when `second == -1` this adds `0`).
5. Return `first` so the parent can treat it as this node's height contribution.

Equivalent to `n + Σ (best2 + 1)` over branching nodes; `O(n)` time.

### Sample 2 check

Tree: `1` with children `2, 3` (both leaves).

- Leaves: each adds `1`.
- Root: `best2 = 0` → add `1 + (0 + 1) = 2`.
- Total `4` = `3` singletons + guild `{2, 3}`.
