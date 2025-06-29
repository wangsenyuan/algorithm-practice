# Graph Cutting Contest

Little Chris is participating in a graph cutting contest. He's a pro. The time has come to test his skills to the fullest.

## Problem Description

Chris is given a simple undirected connected graph with $n$ vertices (numbered from 1 to $n$) and $m$ edges. The problem is to cut it into edge-distinct paths of length 2. Formally, Chris has to partition all edges of the graph into pairs in such a way that:

- The edges in a single pair are adjacent
- Each edge must be contained in exactly one pair

For example, the figure shows a way Chris can cut a graph. The first sample test contains the description of this graph.

You are given a chance to compete with Chris. Find a way to cut the given graph or determine that it is impossible!

## Input

The first line of input contains two space-separated integers $n$ and $m$ ($1 \leq n, m \leq 10^5$), the number of vertices and the number of edges in the graph.

The next $m$ lines contain the description of the graph's edges. The $i$-th line contains two space-separated integers $a_i$ and $b_i$ ($1 \leq a_i, b_i \leq n$; $a_i \neq b_i$), the numbers of the vertices connected by the $i$-th edge.

**Note:** The given graph is guaranteed to be simple (without self-loops and multi-edges) and connected.

**Important:** Since the size of the input and output could be very large, don't use slow output techniques in your language. For example, do not use input and output streams (`cin`, `cout`) in C++.

## Output

If it is possible to cut the given graph into edge-distinct paths of length 2:

- Output $\frac{m}{2}$ lines
- In the $i$-th line print three space-separated integers $x_i$, $y_i$ and $z_i$, the description of the $i$-th path
- The graph should contain this path, i.e., the graph should contain edges $(x_i, y_i)$ and $(y_i, z_i)$
- Each edge should appear in exactly one path of length 2
- If there are multiple solutions, output any of them

If it is impossible to cut the given graph, print `"No solution"` (without quotes).


## ideas
1. m需要是偶数。
2. 在例子中， 如果先选择(1, 2, 3), (1, 4, 3), 那么边（2， 4）就会被剩下，无法处理
3. 如果是一棵树，且有偶数条边，（有奇数个点），那么就肯定有答案
4. 所以，如果能够把图，先分成奇数个点组成的图（1个点，3个点），那么肯定是可以的
5. 但是要从graph出分出这样的树，似乎有点麻烦
6. 还是应该找不变量
7. 每次选择3个点，(a, b, c), 删除(a b), (b, c), 那么deg(a) - 1, deg(b) - 2, deg(c) - 1
8. 如果从所有点的deg的xor来看，似乎是没有变的 xor(deg(_) pairity) = 0
9. 然后每次都选择deg最大的点，它肯定是可以找到两个邻居 