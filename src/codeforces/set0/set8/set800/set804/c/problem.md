# Problem Description

Isart and Modsart were trying to solve an interesting problem when suddenly Kasra arrived. Breathless, he asked: "Can you solve a problem I'm stuck at all day?"

We have a tree T with n vertices and m types of ice cream numerated from 1 to m. Each vertex i has a set of si types of ice cream. Vertices which have the i-th (1 ≤ i ≤ m) type of ice cream form a connected subgraph. We build a new graph G with m vertices. We put an edge between the v-th and the u-th (1 ≤ u, v ≤ m, u ≠ v) vertices in G if and only if there exists a vertex in T that has both the v-th and the u-th types of ice cream in its set. The problem is to paint the vertices of G with minimum possible number of colors in a way that no adjacent vertices have the same color.

Please note that we consider that empty set of vertices form a connected subgraph in this problem.

As usual, Modsart don't like to abandon the previous problem, so Isart wants you to solve the new problem.

## Input

The first line contains two integer n and m (1 ≤ n, m ≤ 3×10^5) — the number of vertices in T and the number of ice cream types.

n lines follow, the i-th of these lines contain single integer si (0 ≤ si ≤ 3×10^5) and then si distinct integers, each between 1 and m — the types of ice cream in the i-th vertex. The sum of si doesn't exceed 5×10^5.

n - 1 lines follow. Each of these lines describes an edge of the tree with two integers u and v (1 ≤ u, v ≤ n) — the indexes of connected by this edge vertices.

## Output

Print single integer c in the first line — the minimum number of colors to paint the vertices in graph G.

In the second line print m integers, the i-th of which should be the color of the i-th vertex. The colors should be between 1 and c. If there are some answers, print any of them.

## Examples

### Example 1

**Input:**

```text
3 3
1 1
2 2 3
1 2
1 2
2 3
```

**Output:**

```text
2
1 1 2
```

### Example 2

**Input:**

```text
4 5
0
1 1
1 3
3 2 4 5
2 1
3 2
4 3
```

**Output:**

```text
3
1 1 1 2 3
```

## Note

- In the first example the first type of ice cream is present in the first vertex only, so we can color it in any color. The second and the third ice cream are both presented in the second vertex, so we should paint them in different colors.

- In the second example the colors of the second, the fourth and the fifth ice cream should obviously be distinct.


### ideas
1. 由ice creams作为节点组成的G，如果(u, v) 在某个T的vertex上同时存在，就在它们之间连一条边
2. 然后用最少的颜色给节点着色，使的相邻的节点，不能是相同的颜色
3. 图着色应该是NP-hard的。这个题目出在这里，应该是可以通过某种方式解决。
4. 有个疑问，为什么要提供一个tree？还有就是，包含相同类型ice creame的节点，是在T上连通的
5. 感觉是个构造型的问题，最少的颜色数，应该是最多ice creame节点上的数量
6. 找到出现次数最多的ice cream，给他编号1
7. 假设已经处理了一组，现在是新的ice cream；这里主要的问题是如何判断，它是否和已有编号，处在相同的节点上
8. 要找到mex？
9. 这里似乎，是不是会用到联通的性质了？
10. 假设每个节点维护一个值，mex（就是它目前最小的未出现的编号）
11. 那么当前ice的编号 = max(mex of 这堆的节点)
12. 当更新了一个ice 的编号后，就可以反过来去更新这些节点
13. 但是复杂性是 n * m，还得再想想
14. 假设有一个节点有 k个ice creame，在这个节点上处理后，移动到下一个节点
15. 可能会增加几个新的进来，同时减少几个
16. 新增的似乎，比较好处理，减少的怎么处理呢？
17. 好像可以搞