# Problem Statement

Almost everything has density, even a graph. We define the density of a non-directed graph (nodes and edges of the graph have some values) as follows:

$$density = \frac{v}{e}$$

where $v$ is the sum of the values of the nodes, $e$ is the sum of the values of the edges.

Once DZY got a graph $G$, now he wants to find a connected induced subgraph $G'$ of the graph, such that the density of $G'$ is as large as possible.

An induced subgraph $G'(V', E')$ of a graph $G(V, E)$ is a graph that satisfies:

- $V' \subseteq V$;
- edge $(u, v) \in E'$ if and only if $u \in V'$, $v \in V'$, and edge $(u, v) \in E$;
- the value of an edge in $G'$ is the same as the value of the corresponding edge in $G$, so as the value of a node.

Help DZY to find the induced subgraph with maximum density. Note that the induced subgraph you choose must be connected.

## Input

The first line contains two space-separated integers $n$ ($1 \le n \le 500$), $m$ ($0 \le m \le \frac{n(n-1)}{2}$). Integer $n$ represents the number of nodes of the graph $G$, $m$ represents the number of edges.

The second line contains $n$ space-separated integers $x_i$ ($1 \le x_i \le 10^6$), where $x_i$ represents the value of the $i$-th node. Consider the graph nodes are numbered from $1$ to $n$.

Each of the next $m$ lines contains three space-separated integers $a_i$, $b_i$, $c_i$ ($1 \le a_i < b_i \le n$; $1 \le c_i \le 10^3$), denoting an edge between node $a_i$ and $b_i$ with value $c_i$. The graph won't contain multiple edges.

## Output

Output a real number denoting the answer, with an absolute or relative error of at most $10^{-9}$.

## Examples

### Input
```
1 0
1
```

### Output
```
0.000000000000000
```

### Input
```
2 1
1 2
1 2 1
```

### Output
```
3.000000000000000
```

### Input
```
5 6
13 56 73 98 17
1 2 56
1 3 29
1 4 42
2 3 95
2 4 88
3 4 63
```

### Output
```
2.965517241379311
```

## Note

In the first sample, you can only choose an empty subgraph, or the subgraph containing only node 1.

In the second sample, choosing the whole graph is optimal.


### ideas
1. induced graph 肯定是一棵树，这个可以反证出来, 多出来的那些边在被删除的情况下，会得到更优的结果
2. 在选中一组节点的情况下，使用最小的m-1条边是最优的结果
3. 但是问题在于，这组节点，不一定能 connect在一起
4. 按照edge 从小到大的的处理
5. 如果把 g[a] ~ g[b] 连接起来后, 密度更大，就连接起来