# Problem

For a tree `T` with root `r`, where each node `u` has a value `au`, the **cost** of the tree is defined as:

`sum over u in T of (au * d(r, u))`

where `d(r, u)` is the number of edges on the shortest path from `r` to `u` in `T`.

You are given a tree with `n` nodes, rooted at node `1`. Each node `i` has value `ai`.

For each `r` from `1` to `n`, solve the following **independently**:

Consider the **subtree of node `r`** with respect to root `1`. Formally, it is the tree consisting of all nodes `u` such that the shortest path from `1` to `u` contains `r`.

Find the **maximum** cost of this subtree after performing **at most one** operation of the following type:

- Choose any node `u` with `u ≠ r`.
- Remove the edge from `u` to its parent `p` (the unique node with `d(u, p) = 1` and `d(u, r) = d(p, r) + 1`).
- Add an edge from `u` to any node `v` that remains reachable from `r` in the current subtree.

After the operation, the graph must remain a tree.

## Input

The first line contains integer `t` (`1 ≤ t ≤ 10^4`) — the number of test cases.

Each test case:

- First line: integer `n` (`1 ≤ n ≤ 2 * 10^5`) — number of nodes.
- Second line: `n` integers `a1, a2, ..., an` (`1 ≤ ai ≤ 2 * 10^5`).
- Next `n - 1` lines: two integers `u` and `v` (`1 ≤ u, v ≤ n`) — an edge of the tree.

It is guaranteed that the edges form a tree.

The sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, print `n` integers — the answers for `r = 1, 2, ..., n`, in order.

## Example

**Input**

```text
3
5
1 3 2 1 2
1 2
2 3
3 4
3 5
7
1 2 3 1 3 2 1
1 2
2 3
3 4
4 5
4 6
3 7
5
5 4 3 2 1
1 2
2 3
3 4
4 5
```

**Output**

```text
18 10 5 0 0
40 28 18 8 0 0 0
20 10 4 1 0
```

## Note

In the first test case, for `r = 1`, it is optimal to choose `u = 5` and `v = 4`. The cost becomes:

`1 * 0 + 3 * 1 + 2 * 2 + 1 * 3 + 2 * 4 = 18`

For `r = 4`, the subtree has only one node, so no operation is possible and the cost is `0`.

## key insights

1. The code runs one DFS on the global tree rooted at node `1` (index `0`). For each node `u` it computes values that are combined into the printed answer for `r = u + 1` as `fp[u] + dp[u]`.

2. **Base contribution (`fp`)**  
   While DFSing, it maintains:
   - `sum[u]` — total `a` over the subtree of `u` (with respect to root `1`),
   - `dep[u]` — depth from root `1`,
   - `fp[u]` — accumulates the weighted depth sum for the subtree using the recurrence  
     `fp[u] += fp[v] + sum[v]` over children `v` (equivalently: extending the subtree through edge `(u, v)` adds `sum[v]` to the depth contribution of everything below `v`).

3. **One-operation extra gain (`dp`)**  
   The operation “rehang one child subtree” is modeled as attaching a child subtree `v` under the **deepest** part of a different child branch to maximize depth gain.

4. For each node `u`, it tracks the two child subtrees with the largest `most_dep` (deepest node depth inside that child subtree):
   - `best[0]` — child with maximum `most_dep`,
   - `best[1]` — second maximum.

5. For each child `v` of `u`, it evaluates a candidate gain:
   - if `v` is not `best[0]`, hang `v` under the deepest branch of `best[0]`:
     - `sum[v] * (most_dep[best[0]] - dep[v] + 1)`
   - else hang `v` under `best[1]` similarly:
     - `sum[v] * (most_dep[best[1]] - dep[v] + 1)`

   Then `dp[u] = max over children of (dp[v], and these candidate tmp values)`.

6. If `u` has fewer than two nontrivial child branches (`best[1] < 0`), no cross-branch rehang is possible at `u`, so the extra gain from this pattern is `0` at that node.

7. Final per-node output is `fp[i] + dp[i]` for all `i` (answers for `r = 1..n` in order).

8. Complexity: `O(n)` per test case (one DFS, constant work per edge), which matches `sum n ≤ 2 * 10^5`.
