# Tree Minimization Problem

## Problem Description

Vanya wants to minimize a tree. He can perform the following operation multiple times:

Choose a vertex v, and two disjoint (except for v) paths of equal length:
- Path 1: a₀ = v, a₁, ..., aₖ
- Path 2: b₀ = v, b₁, ..., bₖ

Additionally, vertices a₁, ..., aₖ, b₁, ..., bₖ must not have any neighbours in the tree other than adjacent vertices of corresponding paths.

After that, one of the paths may be merged into the other, that is, the vertices b₁, ..., bₖ can be effectively erased.

## Task

Help Vanya determine if it is possible to make the tree into a path via a sequence of described operations, and if the answer is positive, also determine the shortest length of such path.

## Input

- The first line contains the number of vertices n (2 ≤ n ≤ 2·10⁵)
- Next n-1 lines describe edges of the tree
- Each edge line contains two space-separated integers u and v (1 ≤ u, v ≤ n, u ≠ v) — indices of endpoints of the corresponding edge
- It is guaranteed that the given graph is a tree

## Output

- If it is impossible to obtain a path, print `-1`
- Otherwise, print the minimum number of edges in a possible path

## Examples

### Example 1

**Input:**
```
6
1 2
2 3
2 4
4 5
1 6
```

**Output:**
```
3
```

**Explanation:** A path of three edges is obtained after merging paths 2-1-6 and 2-4-5.

### Example 2

**Input:**
```
7
1 2
1 3
3 4
1 5
5 6
6 7
```

**Output:**
```
-1
```

**Explanation:** It is impossible to perform any operation. For example, it is impossible to merge paths 1-3-4 and 1-5-6, since vertex 6 additionally has a neighbour 7 that is not present in the corresponding path.

## ideas
1. 考虑一条直线，如果它有奇数个节点，那么就可以折半；否则就无法处理
2. 考虑这棵树的直径，如果它是偶数个节点组成的，同样没法折半（只有两条直径的边合并，才能减少直径）
3. 然后考虑其他所有非直径的边，能否合并掉？