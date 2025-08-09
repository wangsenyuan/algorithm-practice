# Alice and Bob and Triangles

Alice and Bob study properties of graphs. Alice takes a complete undirected graph with $n$ vertices, chooses some $m$ edges and keeps them. Bob gets the remaining edges.

Alice and Bob are fond of triangles in graphs, i.e., cycles of length 3. They wonder: what is the total number of triangles in the two graphs formed by Alice's and Bob's edges, respectively?

## Input

- The first line contains two space-separated integers $n$ and $m$ ($1 \le n \le 10^6$, $0 \le m \le 10^6$) — the number of vertices in the initial complete graph and the number of edges in Alice's graph, respectively.
- Then $m$ lines follow: the $i$-th line contains two space-separated integers $a_i$, $b_i$ ($1 \le a_i, b_i \le n$, $a_i \ne b_i$) — the endpoints of an edge in Alice's graph.
- Alice's graph contains no multiple edges and no self-loops. The initial complete graph also contains no multiple edges and no self-loops.
- Vertices are indexed from 1 to $n$.

## Output

Print a single integer — the total number of cycles of length 3 in Alice's and Bob's graphs together.

## Examples

### Example 1

Input

```text
5 5
1 2
1 3
2 3
2 4
3 4
```

Output

```text
3
```

### Example 2

Input

```text
5 3
1 2
2 3
1 3
```

Output

```text
4
```

## Notes

- In the first sample Alice has 2 triangles: (1, 2, 3) and (2, 3, 4). Bob's graph has 1 triangle: (1, 4, 5). Total = 3.
- In the second sample Alice has 1 triangle: (1, 2, 3). Bob's graph has 3 triangles: (1, 4, 5), (2, 4, 5), (3, 4, 5). Total = 4.


### ideas
1. 在完全图中，任选3个点，就可以组成一个三角形；
2. 