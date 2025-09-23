# Problem C: Alyona and the Tree

## Problem Statement

Alyona decided to go on a diet and went to the forest to get some apples. There she unexpectedly found a magic rooted tree with root in the vertex 1, every vertex and every edge of which has a number written on.

The girl noticed that some of the tree's vertices are sad, so she decided to play with them. Let's call vertex v **sad** if there is a vertex u in subtree of vertex v such that `dist(v, u) > au`, where:
- `au` is the number written on vertex u
- `dist(v, u)` is the sum of the numbers written on the edges on the path from v to u

Leaves of a tree are vertices connected to a single vertex by a single edge, but the root of a tree is a leaf if and only if the tree consists of a single vertex — root.

Thus Alyona decided to remove some of tree leaves until there will be no any sad vertex left in the tree. What is the minimum number of leaves Alyona needs to remove?

## Input

- In the first line of the input integer `n` (1 ≤ n ≤ 10⁵) is given — the number of vertices in the tree.
- In the second line the sequence of n integers `a₁, a₂, ..., aₙ` (1 ≤ aᵢ ≤ 10⁹) is given, where aᵢ is the number written on vertex i.
- The next n - 1 lines describe tree edges: ith of them consists of two integers `pᵢ` and `cᵢ` (1 ≤ pᵢ ≤ n, -10⁹ ≤ cᵢ ≤ 10⁹), meaning that there is an edge connecting vertices i + 1 and pᵢ with number cᵢ written on it.

## Output

Print the only integer — the minimum number of leaves Alyona needs to remove such that there will be no any sad vertex left in the tree.

## Example

**Input:**
```
9
88 22 83 14 95 91 98 53 11
3 24
7 -8
1 67
1 64
9 65
5 12
6 -80
3 8
```

**Output:**
```
5
```

## Note

The following image represents possible process of removing leaves from the tree:



### ideas
1. 我们考虑节点v，如果要满足它的限制, dist(v, ?) <= a[v]
2. 那么就需要把所有超过限制的节点都给删除掉。
3. 在暴力的情况下，必须一层层的从下往上删除，肯定是TLE的。
4. 反过来，考虑u，如果它有一个祖先节点限制它必须被删除掉，那么就没必要往下处理了
5. 所以，dist[v, u] = prev[u] - prev[v] > a[u]
6. prev[u] - a[u] > prev[v]
8.  