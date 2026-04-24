# Problem (Hard Version)

This is the hard version of the problem. The difference between the versions is that in this version, the constraint on `n` is higher. You can hack only if you solved all versions of this problem.

You once had an undirected tree with `n` nodes. To make the tree look more interesting, you assigned an **arbitrary direction** to each of the `n - 1` edges.

Over time you forgot the tree structure. You found a note that records, **after** orienting the edges, for every ordered pair `(u, v)` with `1 ≤ u, v ≤ n`, whether `u` can reach `v` in the resulting directed graph.

You must recover **some** tree structure and edge directions consistent with the note, or report that none exists. If multiple solutions exist, print any.

For a directed graph, `x` can reach `y` if and only if there is a sequence of nodes `u1, u2, ..., uk` such that `u1 = x`, `uk = y`, and for every `i` from `2` to `k` the directed edge `u(i-1) → ui` exists. In particular, every node can reach itself.

## Input

The first line contains integer `t` (`1 ≤ t ≤ 10^4`) — the number of test cases.

Each test case:

- First line: integer `n` (`2 ≤ n ≤ 8000`) — number of nodes.
- Next `n` lines: string `si` of length `n`, consisting only of `'0'` and `'1'`.
  The `j`-th character of `si` is `'1'` if and only if node `i` can reach node `j` after the edges are directed.

It is guaranteed that the sum of `n^2` over all test cases does not exceed `8000^2`.

## Output

For each test case:

- If a solution exists, print `Yes` (case-insensitive accepted, e.g. `yEs`, `YES`).
- Then print `n - 1` lines, each with two integers `x` and `y`, meaning there is a directed edge `x → y` in your oriented tree.

If no solution exists, print `No` (case-insensitive accepted).

If there are multiple valid trees/orientations, print any.

## Example

### Input

```text
11
4
1000
1111
1010
0001
4
1111
0111
0010
0111
4
0011
0111
0011
0001
4
1000
0110
0010
1111
4
1000
0110
1010
1111
5
10000
01011
00111
00010
00001
5
10000
11000
10101
10111
00001
5
10000
01101
00100
01110
10001
4
1100
0100
0011
0001
4
1110
0100
0010
0101
3
100
111
101
```

### Output

```text
Yes
2 3
2 4
3 1
No
No
Yes
2 3
4 1
4 2
No
No
Yes
2 1
3 1
3 5
4 3
No
No
Yes
1 2
1 3
4 2
Yes
2 3
3 1
```

## Note

In the first test case, nodes `1` and `4` can only reach themselves, node `2` can reach every node, and node `3` can reach only nodes `1` and `3`. The printed edges satisfy these reachability constraints.

In the second test case, no valid tree and orientation exist.

## Ideas

1. `s[i][j] = 1` 表示从 i 能够到达 j，也就是存在一条 path: i -> ... -> j
2. 如果 i, j 中间不存在 u，使得 `s[i][u] = 1` 且 `s[u][j] = 1`，那么 i->j 成立
3. 关键约束：`n` 最大为 8000，sum of `n^2` ≤ `8000^2`，需要 O(n^2) 算法

## Key Insights

### 1. Validity check: diagonal must be all '1'
`s[i][i]` must be `'1'` for every node (every node reaches itself). Fail immediately otherwise.

### 2. Reachability defines a DAG of "reachability sets"
Define the **out-degree** of node `i` as the number of nodes `j ≠ i` with `s[i][j] = '1'` (i.e., how many nodes `i` can reach). The reachability relation on a directed tree is a partial order — it must be acyclic (no two distinct nodes can mutually reach each other unless they are the same node).

### 3. Topological sort as the processing order
Compute in-degree in the reachability graph (how many nodes can reach `i`, excluding `i` itself). Nodes with in-degree 0 are "roots" — nothing reaches them except themselves. Run a standard BFS-based topological sort. If not all `n` nodes are processed, a cycle exists in the reachability relation → output `No`.

The topological order `que[]` gives a safe DFS traversal order: when we visit `u`, all nodes that can reach `u` have already been processed.

### 4. Spanning tree construction via Union-Find
Traverse nodes in topological order. For each node `u`, iterate over all `v` with `s[u][v] = '1'`. If `u` and `v` are in **different components** (Union-Find), add the directed edge `u → v` and merge their components. This greedily builds a spanning tree using only edges that exist in the reachability graph.

The key invariant: we only add edge `u → v` when `u` can reach `v` (`s[u][v] = '1'`), and we avoid cycles by checking Union-Find membership.

### 5. Final verification
After constructing the candidate tree (exactly `n-1` edges, all nodes connected), run BFS from every node on the candidate directed tree and compare the reachability string against the input `s[i]`. This catches cases where the greedy construction produced a structurally valid tree but with wrong reachability (e.g., missing transitive edges). The verify step is O(n²) total.

### 6. Complexity
- Topological sort: O(n²)
- Spanning tree DFS + Union-Find: O(n²)  
- Verification (n BFS runs): O(n²)  
- Total: **O(n²)** per test case, fitting within the `sum of n² ≤ 8000²` constraint.
