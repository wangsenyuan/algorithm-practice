# MST with Required Edges

## Problem Description

For a connected undirected weighted graph $G$, MST (minimum spanning tree) is a subgraph of $G$ that contains all of $G$'s vertices, is a tree, and sum of its edges is minimum possible.

You are given a graph $G$. If you run a MST algorithm on the graph, it would give you only one MST and it causes other edges to become jealous. You are given some queries, each query contains a set of edges of graph $G$, and you should determine whether there is a MST containing all these edges or not.

## Input

- **First line**: Two integers $n, m$ ($2 \leq n, m \leq 5 \cdot 10^5$, $n - 1 \leq m$) — the number of vertices and edges in the graph.
- **Next $m$ lines**: Each line contains three integers $u_i, v_i, w_i$ ($u_i \neq v_i$, $1 \leq w_i \leq 5 \cdot 10^5$) — the endpoints and weight of the $i$-th edge.
  - There can be more than one edge between two vertices.
  - It's guaranteed that the given graph is connected.
- **Next line**: A single integer $q$ ($1 \leq q \leq 5 \cdot 10^5$) — the number of queries.
- **Next $q$ lines**: Each line contains a query starting with an integer $k_i$ ($1 \leq k_i \leq n - 1$) — the size of edges subset, followed by $k_i$ distinct space-separated integers from 1 to $m$ — the indices of the edges.
  - It is guaranteed that the sum of $k_i$ for $1 \leq i \leq q$ does not exceed $5 \cdot 10^5$.

## Output

For each query, print:
- `"YES"` (without quotes) if there's a MST containing the specified edges
- `"NO"` (without quotes) otherwise

## Example

### Input
```
5 7
1 2 2
1 3 2
2 3 1
2 4 1
3 4 1
3 5 2
4 5 2
4
2 3 4
3 3 4 5
2 1 7
2 1 2
```

### Output
```
YES
NO
YES
NO
```

### Explanation

This is the graph of the sample:

![Graph visualization would go here]

**Weight of minimum spanning tree on this graph is 6.**

- **Query 1**: MST with edges $(1, 3, 4, 6)$ contains all edges from the first query, so answer is `"YES"`.
- **Query 2**: Edges from the second query form a cycle of length 3, so there is no spanning tree including these three edges. Thus, answer is `"NO"`.
- **Query 3**: MST can include edges 1 and 7, so answer is `"YES"`.
- **Query 4**: MST cannot include both edges 1 and 2 as they would form a cycle, so answer is `"NO"`.

## ideas
1. 如果一条边，在没有被选中的情况下，但是可以通过交换，能够被替换某条已经被选中的边
2. 那么也是ok的
3. 假设这条边的端点是(u, v), 那么找到它们的lca, p， 如果path (p..u), (p...v)中间存在一条边，它的weight和这条边相同，那么就可以用它来替换
4. 但是这样子，似乎很难搞
5. 能在构造的过程中判断吗？好像可以，这时候，只需要判断路径中的最大值，是否相同就可以了
6.  