# Problem B: Tree Operations

## Problem Description

A tree is a graph with $n$ vertices and exactly $n - 1$ edges; this graph should meet the following condition: there exists exactly one shortest (by number of edges) path between any pair of its vertices.

A subtree of a tree $T$ is a tree with both vertices and edges as subsets of vertices and edges of $T$.

You're given a tree with $n$ vertices. Consider its vertices numbered with integers from 1 to $n$. Additionally, an integer is written on every vertex of this tree. Initially, the integer written on the $i$-th vertex is equal to $v_i$.

In one move, you can apply the following operation:
- Select the subtree of the given tree that includes the vertex with number 1.
- Increase (or decrease) by one all the integers which are written on the vertices of that subtree.

**Task:** Calculate the minimum number of moves that is required to make all the integers written on the vertices of the given tree equal to zero.

## Input

- The first line contains an integer $n$ ($1 \leq n \leq 10^5$).
- Each of the next $n - 1$ lines contains two integers $a_i$ and $b_i$ ($1 \leq a_i, b_i \leq n$; $a_i \neq b_i$) indicating there's an edge between vertices $a_i$ and $b_i$.
- The last line contains $n$ space-separated integers $v_1, v_2, \ldots, v_n$ ($|v_i| \leq 10^9$).

**Note:** It's guaranteed that the input graph is a tree.

## Output

Print the minimum number of operations needed to solve the task.

## Notes

- Please, do not write the `%lld` specifier to read or write 64-bit integers in C++.
- It is preferred to use the `cin`, `cout` streams or the `%I64d` specifier.

## Examples

### Example 1

**Input:**
```
3
1 2
1 3
1 -1 1
```

**Output:**
```
3
```

**Explanation:**
The tree has 3 vertices with edges (1,2) and (1,3). The initial values are [1, -1, 1]. We need 3 operations to make all values zero.


## ideas
1. 没有给定root， 所以操作1，是指包括节点1的子树，进行操作？
2. 也可以单独选择节点1
3. 首先应该考虑对叶子节点进行操作（所有相同值的叶子节点，应该同时处理）然后处理下一层？
4. 