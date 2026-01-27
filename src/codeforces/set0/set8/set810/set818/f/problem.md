Ivan is developing his own computer game. Now he tries to create some levels for his game. But firstly for each level he needs to draw a graph representing the structure of the level.

Ivan decided that there should be exactly nᵢ vertices in the graph representing level i, and the edges have to be bidirectional. When constructing the graph, Ivan is interested in special edges called bridges. An edge between two vertices u and v is called a bridge if this edge belongs to every path between u and v (and these vertices will belong to different connected components if we delete this edge). For each level Ivan wants to construct a graph where at least half of the edges are bridges. He also wants to maximize the number of edges in each constructed graph.

So the task Ivan gave you is: given q numbers n₁, n₂, ..., nq, for each i tell the maximum number of edges in a graph with nᵢ vertices, if at least half of the edges are bridges. Note that the graphs cannot contain multiple edges or self-loops.

## Input

The first line of input file contains a positive integer q (1 ≤ q ≤ 100,000) — the number of graphs Ivan needs to construct.

Then q lines follow, i-th line contains one positive integer nᵢ (1 ≤ nᵢ ≤ 2·10⁹) — the number of vertices in i-th graph.

Note that in hacks you have to use q = 1.

## Output

Output q numbers, i-th of them must be equal to the maximum number of edges in i-th graph.

## Example

### Input
```
3
3
4
6
```

### Output
```
2
3
6
```

## Note

In the first example it is possible to construct these graphs:

1 - 2, 1 - 3;
1 - 2, 1 - 3, 2 - 4;
1 - 2, 1 - 3, 2 - 3, 1 - 4, 2 - 5, 3 - 6.


### ideas
1. 好像很简单呐， h = (n + 1) / 2
2. h + (n - h) * (n - h - 1) / 2
3. when n = 4
4. h = 2, res = 2 + 1 = 3
5. 理解错了
6. 如果 n = 10, 如果有x个边组成桥，连接了一个叶子节点，那么剩下(n - x)个节点
7. 那么 n-x个节点，必须至少组成一个圈, 有 n - x 条边，且最多有 (n - x) * (n - x - 1) / 2
8. 所有有 n - x <= x, （n - x) * (n - x - 1) / 2 >= x
9. 