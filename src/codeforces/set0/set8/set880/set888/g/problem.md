# Problem G: Minimum Spanning Tree with XOR Weights

## Problem Statement

You are given a complete undirected graph with n vertices. A number ai is assigned to each vertex, and the weight of an edge between vertices i and j is equal to ai ⊕ aj.

Calculate the weight of the minimum spanning tree in this graph.

## Input

The first line contains n (1 ≤ n ≤ 200000) — the number of vertices in the graph.

The second line contains n integers a1, a2, ..., an (0 ≤ ai < 2^30) — the numbers assigned to the vertices.

## Output

Print one number — the weight of the minimum spanning tree in the graph.

## Examples

### Example 1
**Input:**
```
5
1 2 3 4 5
```

**Output:**
```
8
```

### Example 2
**Input:**
```
4
1 2 3 4
```

**Output:**
```
8
```