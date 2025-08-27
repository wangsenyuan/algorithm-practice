# Problem C

## Description

You are given an undirected graph with $n$ vertices. There are no edge-simple cycles with the even length in it. In other words, there are no cycles of even length that pass each edge at most once. Let's enumerate vertices from 1 to $n$.

You have to answer $q$ queries. Each query is described by a segment of vertices $[l; r]$, and you have to count the number of its subsegments $[x; y]$ ($l \leq x \leq y \leq r$), such that if we delete all vertices except the segment of vertices $[x; y]$ (including $x$ and $y$) and edges between them, the resulting graph is bipartite.

## Input

The first line contains two integers $n$ and $m$ ($1 \leq n \leq 3 \cdot 10^5$, $1 \leq m \leq 3 \cdot 10^5$) — the number of vertices and the number of edges in the graph.

The next $m$ lines describe edges in the graph. The $i$-th of these lines contains two integers $a_i$ and $b_i$ ($1 \leq a_i, b_i \leq n$; $a_i \neq b_i$), denoting an edge between vertices $a_i$ and $b_i$. It is guaranteed that this graph does not contain edge-simple cycles of even length.

The next line contains a single integer $q$ ($1 \leq q \leq 3 \cdot 10^5$) — the number of queries.

The next $q$ lines contain queries. The $i$-th of these lines contains two integers $l_i$ and $r_i$ ($1 \leq l_i \leq r_i \leq n$) — the query parameters.

## Output

Print $q$ numbers, each in new line: the $i$-th of them should be the number of subsegments $[x; y]$ ($l_i \leq x \leq y \leq r_i$), such that the graph that only includes vertices from segment $[x; y]$ and edges between them is bipartite.

## Examples

### Example 1

**Input:**
```
6 6
1 2
2 3
3 1
4 5
5 6
6 4
3
1 3
4 6
1 6
```

**Output:**
```
5
5
14
```

### Example 2

**Input:**
```
8 9
1 2
2 3
3 1
4 5
5 6
6 7
7 8
8 4
7 2
3
1 8
1 4
3 8
```

**Output:**
```
27
8
19
```


### ideas
1. 每条边，只会（最多一次）出现在一个圈中，且这个圈是一个奇数长度的圈
2. 如果一个圈，处在l...r中间，设这个圈里面的最小下标为x，最大下标为y
3. 那么 (x - l + 1) * (r - y + 1) 的子区间，都是bad的
4. 所以可以预先计算 
5. dp[r] 表示1...r中间的bad的子区间数量
6. dp[r] - dp[l-1] - (左端点在1...l-1, 右端点在 l....r 中间的bad区间) = 只在区间l....r中间的bad区间
7. 跨越的部分，x < l and y >= l 的圈的数量 (它们的贡献 = x * (r - y + 1))
8. 但是，还是不大对，假设有两个圈(x1, y1), (x2, y2)
9. 都满足条件，难道加两次？肯定是不对的