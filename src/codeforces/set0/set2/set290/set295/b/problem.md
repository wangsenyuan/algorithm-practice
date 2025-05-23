# Problem Statement

Greg has a weighted directed graph, consisting of $n$ vertices. In this graph, any pair of distinct vertices has an edge between them in both directions. Greg loves playing with the graph and now he has invented a new game:

The game consists of $n$ steps.
On the $i$-th step, Greg removes vertex number $x_i$ from the graph. As Greg removes a vertex, he also removes all the edges that go in and out of this vertex.
Before executing each step, Greg wants to know the sum of lengths of the shortest paths between all pairs of the remaining vertices. The shortest path can go through any remaining vertex. In other words, if we assume that $d(i, v, u)$ is the shortest path between vertices $v$ and $u$ in the graph that formed before deleting vertex $x_i$, then Greg wants to know the value of the following sum:

$$
\sum_{v \neq u} d(i, v, u)
$$

Help Greg, print the value of the required sum before each step.

---

## Input

- The first line contains integer $n$ ($1 \leq n \leq 500$) — the number of vertices in the graph.
- Next $n$ lines contain $n$ integers each — the graph adjacency matrix: the $j$-th number in the $i$-th line $a_{ij}$ ($1 \leq a_{ij} \leq 10^5$, $a_{ii} = 0$) represents the weight of the edge that goes from vertex $i$ to vertex $j$.
- The next line contains $n$ distinct integers: $x_1, x_2, ..., x_n$ ($1 \leq x_i \leq n$) — the vertices that Greg deletes.

---

## Output

Print $n$ integers — the $i$-th number equals the required sum before the $i$-th step.

---

**Note:**

Please, do not use the `%lld` specifier to read or write 64-bit integers in C++. It is preferred to use the `cin`, `cout` streams or the `%I64d` specifier.

## ideas
1. 反过来处理？
2. 加入一个点前，后面的节点，
3. 都可以看作是一个以u为root的树
4. 所以有m棵树，每个树有m个节点
5. 然后添加节点u
6. 在树上要怎么处理呢？
7. 先做出以u为root的新树
8. 只要找出，离u最近的v(u->v最短)，然后把v接上来
9. 然后去更新v子树的状态（如果u->w更短，就更新上去）
10. 然后更新u在其他子树中的距离
11. 假设v是进入u最短的边a[v][u]最小 （v是后面的节点）
12. dp[r][w]是以r为根，到v的最近距离
13. dp[r][w] = min(dp[r][w], dp[r][v] + dp[v][u] + dp[u][w], dp[r][u] + dp[u][w])
14. 这样子，貌似是可以在n * n * n 的时间内完成 