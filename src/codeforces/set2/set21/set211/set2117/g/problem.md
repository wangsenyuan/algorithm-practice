# Problem G: Minimum Cost Path

You are given an undirected connected weighted graph. Define the cost of a path of length $k$ to be as follows:

Let the weights of all the edges on the path be $w_1, w_2, \ldots, w_k$.

The cost of the path is $\min_{i=1}^{k} w_i + \max_{i=1}^{k} w_i$, or in other words, the **maximum edge weight + the minimum edge weight**.

Across all paths from vertex $1$ to $n$, report the cost of the path with minimum cost. Note that the path is not necessarily simple.

## Input

The first line contains an integer $t$ $(1 \leq t \leq 10^4)$ — the number of test cases.

The first line of each test case contains two integers $n$ and $m$ $(2 \leq n \leq 2 \cdot 10^5, n-1 \leq m \leq \min(2 \cdot 10^5, \frac{n(n-1)}{2}))$.

The next $m$ lines each contain integers $u, v$ and $w$ $(1 \leq u, v \leq n, 1 \leq w \leq 10^9)$ representing an edge from vertex $u$ to $v$ with weight $w$. 

**Constraints:**
- The graph does not contain self-loops or multiple edges
- The resulting graph is connected
- The sum of $n$ over all test cases does not exceed $2 \cdot 10^5$
- The sum of $m$ over all test cases does not exceed $2 \cdot 10^5$

## Output

For each test case, output a single integer — the minimum cost path from vertex $1$ to $n$.

## ideas
1. 按照w升序处理，那么就是最早把1...n连起来的时候的最大值 + 能够从1或者n到达的最小值
2. 