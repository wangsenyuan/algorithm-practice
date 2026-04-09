# Statement

You are on an undirected connected graph of `n` vertices and `m` weighted edges.  
The edges are indexed from `1` to `m`. The `i`-th edge connects vertices `ui` and `vi`, and has weight `wi`.

You decided to take a wonderful journey around the graph.

Suppose you are currently at vertex `x`. You can do the following operations any number of times:

1. Mark an edge connecting `x` and `y`, take photos along the edge, and move to `y`.  
   This costs exactly the edge weight.
2. Transfer by train to another arbitrary vertex `z != x`.  
   You may choose any path `x ⇝ z` (not necessarily simple).  
   If the path uses edges with indices `e1, e2, ..., ek`, then the transfer cost is:
   `w[max(e1, e2, ..., ek)]`.

Formally, there exists a vertex sequence `p1, p2, ..., p(k+1)` such that:

- `x = p1`
- `z = p(k+1)`
- for each `i in [1, k]`, edge `ei` connects `pi` and `p(i+1)`

You are now at vertex `1`, and you need to mark every edge at least once and return to vertex `1`.  
Calculate the minimum total cost.

Please note: transfer cost is **not** the maximum edge weight on the path, and **not** the maximum index itself; it is the **weight of the edge whose index is maximal** on that path.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `T` (`1 <= T <= 10^4`).

For each test case:

- The first line contains two integers `n` and `m` (`1 <= n <= 10^6`, `0 <= m <= 10^6`).
- Each of the next `m` lines contains `ui, vi, wi` (`1 <= ui, vi <= n`, `1 <= wi <= 10^9`), describing edge `i`.

It is guaranteed that:

- the graph is connected,
- the graph may contain self-loops and multiple edges,
- the sums of `n` and `m` over all test cases do not exceed `10^6` each.

## Output

For each test case, output one integer — the minimum cost.

## Example

### Input

```text
5
5 6
2 4 15
2 5 4
1 3 6
2 3 9
1 2 10
3 4 7
4 3
1 2 3
1 3 2
1 4 1
2 3
1 2 1
2 1 3
1 1 4
6 6
2 3 10
1 3 10
5 6 10
6 6 1
4 5 10
3 4 10
5 5
1 2 4
5 1 5
4 3 6
2 4 10
1 4 7
```

### Output

```text
58
8
8
71
43
```

## Note

Visualizer link

Let `u --e--> v` denote going from vertex `u` to vertex `v` using edge `e`.

In the first test case, one possible solution is:

- Initially at vertex `1`.
- Mark edge `3` and move to vertex `3`, cost `6`.
- Mark edge `6` and move to vertex `4`, cost `7`.
- Mark edge `1` and move to vertex `2`, cost `15`.
- Mark edge `4` and move to vertex `3`, cost `9`.
- Transfer to vertex `5` using path `3 --6--> 4 --1--> 2 --2--> 5`, cost `7`  
  (max edge index on path is `6`, so cost is `w6 = 7`).
- Mark edge `2` and move to vertex `2`, cost `4`.
- Mark edge `5` and move to vertex `1`, cost `10`.

Total cost: `6 + 7 + 15 + 9 + 7 + 4 + 10 = 58`.

In the second test case, one possible solution is:

- Mark edge `1` and move to vertex `2`, cost `3`.
- Transfer to vertex `3` by path `2 --1--> 1 --3--> 4 --3--> 1 --2--> 3`, cost `1`.  
  (The transfer path is allowed to be non-simple.)
- Mark edge `2` and move to vertex `1`, cost `2`.
- Mark edge `3` and move to vertex `4`, cost `1`.
- Transfer to vertex `1` by path `4 --3--> 1`, cost `1`.

## Algorithm

Let the answer be:

- the sum of all edge weights, because every edge must be marked at least once;
- plus the minimum extra cost needed to fix parity, so that the whole route can start at `1`, end at `1`, and use marked edges as an Euler tour with some transfers inserted.

So the real task is: how do we pair the odd-degree vertices using the special transfer cost?

### 1. Why only odd-degree vertices matter

If we only look at marked-edge traversals, every time we enter a vertex through marked edges we must also leave it through marked edges, except for route endpoints. Since the route starts and ends at vertex `1`, after adding transfers every vertex must effectively have even degree in the marked-edge multiset.

That means:

- vertices with even degree need nothing;
- vertices with odd degree must be matched by transfers.

So we only need to compute the cheapest way to pair odd-degree vertices.

### 2. Transfer cost between two vertices

For a transfer from `x` to `y`, we may choose any path, and the price is:

`weight of the edge with maximum index on that path`.

This means the path cost is determined by edge index order, not by edge weight.

A key reformulation is:

- if we process edges in index order `1..m`,
- then the first moment when `x` and `y` become connected
- determines the transfer cost candidate;
- however, later edges inside the same connected component may reduce that cost further, because a later edge may have smaller weight.

So a normal MST is not enough. We need a structure that preserves how connected components evolve in index order.

### 3. Component-evolution tree

Process edges in original order. Maintain DSU connected components.

Create a tree node for every original vertex and for every edge.

When processing edge `(u, v, w)`:

- if `u` and `v` are in different DSU components, create a new internal node with label `w`, and connect it to the current roots of those two component trees; then merge the DSU components, and the new node becomes the root of the merged component;
- if `u` and `v` are already in the same DSU component, create a new unary node with label `w` whose child is the current component root; this records that a later edge appeared inside the component and may lower future transfer cost.

After all edges are processed, each DSU component is represented by one rooted tree. Since the graph is connected, we get one final root.

Interpretation:

- every node represents some connected subgraph that already exists after processing a prefix of edges;
- its label is the weight of the edge that created this state;
- for any two vertices, the first tree node whose subtree contains both vertices represents the earliest component containing both;
- the best transfer cost inside that component is the minimum label on the path from that node upward through unary wrappers created later.

### 4. Best transfer cost for each merge node

Run one DFS from the final root.

Define `best[u]` as the minimum label on the path from the global root down to `u`.

Then for every binary merge node `u`, `best[u]` is exactly the minimum transfer cost that can be used to connect one odd vertex from its left child subtree and one odd vertex from its right child subtree.

Why?

- the two odd vertices become connectable when their components first merge at `u`;
- after that, any later unary wrappers above `u` correspond to extra edges added inside the same larger component, which may reduce the transfer cost;
- taking the minimum label on that ancestor chain captures the cheapest possible transfer after they are already in one component.

### 5. Tree DP for pairing odd vertices

Now do a bottom-up DP on the component tree.

Let `unmatched[u]` be the number of odd-degree vertices in the subtree of `u` that are still unpaired after processing everything inside this subtree as cheaply as possible.

Transitions:

- for an original vertex node: `unmatched[u] = degree(u) mod 2`;
- for a unary node: `unmatched[u] = unmatched[child]`;
- for a binary merge node with children `L` and `R`:
  - let `cnt = unmatched[L] + unmatched[R]`;
  - inside this merged component, we can pair as many cross-subtree odd vertices as possible;
  - each such pair costs `best[u]`;
  - number of pairs is `cnt / 2`;
  - so add `(cnt / 2) * best[u]` to the answer;
  - keep `unmatched[u] = cnt mod 2`.

This greedy pairing is optimal because `best[u]` is the first cost level where odd vertices from the two child components can interact. Any odd vertices not paired here must go upward and can only be paired with cost no smaller than what is available above.

In a connected graph, the total number of odd-degree vertices is even, so the final root leaves `0` unmatched.

### 6. Full procedure

1. Add all edge weights to the answer.
2. Compute parity of each vertex degree, ignoring self-loops for parity because they contribute `2` to degree.
3. Build the component-evolution tree with DSU in original edge order.
4. DFS the tree to compute `best[u]`.
5. Run bottom-up DP to add the minimum extra transfer cost.

### 7. Complexity

- DSU construction: `O(m alpha(n))`
- tree traversal and DP: `O(n + m)`
- memory: `O(n + m)`

This fits the constraint that total `n` and total `m` over all test cases are at most `10^6`.

## Official Editorial 

Since every edge must be marked at least once, a baseline cost is the **sum of all edge weights** (add every `wi` for `i = 1 .. m`).

- If the graph is **Eulerian** (every vertex has even degree), that sum alone suffices: walk an Euler tour that uses each edge once.
- Otherwise, think of adding **virtual edges** whose cost is the **minimum train-transfer cost** between their endpoints. Only **odd-degree** vertices need pairing; optimally you add **one** virtual edge per matched pair.

### Transfer cost and reconstruction tree

To compute transfer costs, build a **reconstruction tree** `T` of `G` by inserting real edges in index order `1 .. m` (Kruskal / DSU style).

For a transfer from `u` to `v`, let `o = LCA(u, v)` in `T`. The edge with **maximum index** on any `u`–`v` path in `G` corresponds to `o` or one of its **ancestors** in `T` (standard reconstruction-tree property).

For each internal node `i` of `T` (nodes that correspond to edges of `G`), precompute:

- `f[i] =` minimum edge weight among **`i` and all ancestors of `i`** (on the path to the root).

Then `f[i] <= f[parent(i)]`, so when two odd vertices first share a subtree rooted at internal node `p`, it is optimal to pair them **as early as possible** using cost **`f[p]`**.

### Greedy DFS matching

DFS the reconstruction tree from the root. For each node `x`, for each child `y`, recurse and track whether an **unmatched leaf** (original graph vertex) of **odd degree** remains in `y`’s subtree. Whenever two such leaves can be paired at `x`, pay **`f[x]`** immediately. When finishing `x`, at most **one** odd leaf in `x`’s subtree stays unmatched and is passed upward.

This greedy pairing is optimal with the above `f[·]`.

### Complexity (editorial)

- **Time:** `O(m * α(n))` for DSU while building the tree, plus linear tree work.
- **Space:** `O(n + m)` for the tree and auxiliary arrays (implementation-dependent; the original note mentions `O(n)` in a looser sense).