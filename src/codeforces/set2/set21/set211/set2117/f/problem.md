# Problem Description

Yousef has a rooted tree* consisting of exactly $n$ vertices, which is rooted at vertex $1$. You would like to give Yousef an array $a$ of length $n$, where each $a_i$ $(1 \leq i \leq n)$ can either be $1$ or $2$.

Let $s_u$ denote the sum of $a_v$ where vertex $v$ is in the subtree† of vertex $u$. Yousef considers the tree special if all the values in $s$ are pairwise distinct (i.e., all subtree sums are unique).

Your task is to help Yousef count the number of different arrays $a$ that result in the tree being special. Two arrays $b$ and $c$ are different if there exists an index $i$ such that $b_i \neq c_i$.

As the result can be very large, you should print it modulo $10^9 + 7$.

---

**Definitions:**
- *A tree is a connected undirected graph with $n-1$ edges.
- †The subtree of a vertex $v$ is the set of all vertices that pass through $v$ on a simple path to the root. Note that vertex $v$ is also included in the set.

## Input

The first line contains an integer $t$ $(1 \leq t \leq 10^4)$ — the number of test cases.

Each test case consists of several lines:
- The first line of the test case contains an integer $n$ $(2 \leq n \leq 2 \cdot 10^5)$ — the number of vertices in the tree.
- Then $n-1$ lines follow, each of them contains two integers $u$ and $v$ $(1 \leq u, v \leq n, u \neq v)$ which describe a pair of vertices connected by an edge.

It is guaranteed that:
- The given graph is a tree and has no loops or multiple edges.
- The sum of $n$ over all test cases doesn't exceed $2 \cdot 10^5$.

## Output

For each test case, print one integer $x$ — the number of different arrays $a$ that result in the tree being special, modulo $10^9 + 7$.

## ideas
1. 如果有超过2个叶子节点 => 0
2. 如果只有一个叶子节点 => 就是一条线， pow(2, n)
3. 如果有两个叶子节点，=> 先找到它们的lcp, lcp上面的点的取值没有影响, pow(2, m)
4. 主要是这两条路径的问题
5. 假设这两条边的长度分别为a, b
6. 如果 a = b (1, 2, 2, 2), (2, 2, 2, 2)， 2 * pow(2, m)
7. 如果 a!= b, (1, 2, 2, ?, ?), (2, 2)
8.      (2, 2, 2, ?, ?), (1, 2, 2) 也是ok的