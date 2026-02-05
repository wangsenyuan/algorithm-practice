Given a tree\* with $n$ vertices numbered from $1$ to $n$. Each vertex has an initial color $a_i$.

Each edge of the tree is defined by three numbers: $u_i$, $v_i$, and $c_i$, where $u_i$ and $v_i$ are the endpoints of the edge, and $c_i$ is the edge parameter. The **cost** of the edge is defined as follows:

- if the colors of vertices $u_i$ and $v_i$ are the same, the cost is $0$;  
- otherwise, the cost is $c_i$.

You are also given $q$ queries. Each query has the form: repaint vertex $v$ to color $x$. The queries depend on each other (after each query, the color change is preserved). After each query, you need to output the sum of the costs of all edges in the tree.

\*A tree is a connected graph without cycles.

## Input

The first line contains an integer $t$ ($1 \le t \le 10^4$) — the number of test cases.

The first line of each test case contains two integers $n$ and $q$ ($1 \le n, q \le 2 \cdot 10^5$) — the number of vertices and the number of queries, respectively.

The second line contains $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \le a_i \le n$), where the $i$-th number specifies the initial color of vertex $i$.

The next $n - 1$ lines describe the edges of the tree. Each line contains three integers $u$, $v$, and $c$, denoting an edge between vertices $u$ and $v$ with parameter $c$ ($1 \le u, v \le n$, $1 \le c \le 10^9$).

The following $q$ lines contain the queries. Each query contains two integers $v$ and $x$ — repaint vertex $v$ to color $x$ ($1 \le v, x \le n$).

It is guaranteed that the sum of $n$ and the sum of $q$ across all test cases do not exceed $2 \cdot 10^5$.

## Output

For each query, output a single integer on a separate line — the sum of the costs of all edges in the tree after applying the corresponding query.

## Example

### Input

```
4
1 1
1
1 1
2 3
1 1
1 2 10
1 2
2 2
1 1
5 4
1 2 1 2 3
1 2 5
2 3 3
2 4 4
4 5 7
3 2
5 2
1 2
2 3
4 3
1 1 2 2
1 2 2
2 3 6
2 4 8
3 1
4 1
2 2
```

### Output

```
0
10
0
10
12
5
0
12
8
0
16
```

## Note

- **First test:** $n = 1$, one vertex — no edges. Query: repaint $a_1$ to $1$, the sum of costs is $0$.

- **Second test:** $n = 2$, edge $1 - 2$ (with $c = 10$). Queries:
  - $a_1 = 2$: colors $[2, 1]$, cost is $10$;
  - $a_2 = 2$: colors $[2, 2]$, cost $0$;
  - $a_1 = 1$: colors $[1, 2]$, cost $10$.

- **Third test:** $n = 5$, edges: $1 - 2$ ($c = 5$), $2 - 3$ ($c = 3$), $2 - 4$ ($c = 4$), $4 - 5$ ($c = 7$). Initial colors $[1, 2, 1, 2, 3]$. Queries:
  - $a_3 = 2 \rightarrow [1, 2, 2, 2, 3]$: edges $1 - 2$ ($c = 5$) and $4 - 5$ ($c = 7$) give $12$;
  - $a_5 = 2 \rightarrow [1, 2, 2, 2, 2]$: only edge $1 - 2$ ($c = 5$) contributes, so the cost is $5$;
  - $a_1 = 2 \rightarrow [2, 2, 2, 2, 2]$: all colors equal, cost is $0$;
  - $a_2 = 3 \rightarrow [2, 3, 2, 2, 2]$: edges $1 - 2$ ($5$), $2 - 3$ ($3$), $2 - 4$ ($4$) give $12$.



### ideas
1. 对于每个节点，存储它对应颜色的贡献（和自己无关，和邻居有关）dp[u][c]
2. 然后在更新自己颜色的时候，那么贡献 = edge_cost_sum - dp[u][new_color] 
3. 但是这个时候，会对邻居的状态也产生影响
4. 如果每个邻居都去处理，那么TLE
5. 如果邻居的数量比较少（似乎去更新也没有）如果邻居很多。那么反过来，这个邻居的邻居就比较少了
6. 所以，似乎可以采取不同的策略，就是如果改变u得时候，发现它的邻居较少,就直接去更新
7. 如果发现它的邻居比较多（那么标记它变化了）根据某个结构算出变化，在处理到它的邻居时，在反过来处理；
8. 似乎可行。