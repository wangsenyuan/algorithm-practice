# Problem D - Sponsored by Declan Akaba

You are given a simple, connected, undirected graph with $n$ vertices and $m$ edges. The graph contains no self-loops or multiple edges. You are also given a multiset $A$ consisting of $\ell$ elements:

$$
A = \{A_1, A_2, \ldots, A_\ell\}
$$

Starting from vertex $1$, you may perform the following move any number of times, as long as the multiset $A$ is not empty:

- Select an element $k \in A$ and remove it from the multiset. You must remove exactly one occurrence of $k$ from $A$.
- Traverse any walk* of exactly $k$ edges to reach some vertex (possibly the same one you started from).

For each $i$ ($1 \leq i \leq n$), determine whether there exists a sequence of such moves that starts at vertex $1$ and ends at vertex $i$, using the original multiset $A$.

Note that the check for each vertex $i$ is independent — you restart from vertex $1$ and use the original multiset $A$ for each case.

---

*\* A walk of length $k$ is a sequence of vertices $v_0, v_1, \ldots, v_{k-1}, v_k$ such that each consecutive pair of vertices $(v_i, v_{i+1})$ is connected by an edge in the graph. The sequence may include repeated vertices.*

---

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \leq t \leq 10^4$). The description of the test cases follows.

The first line of each test case contains three integers $n$, $m$, and $\ell$ ($2 \leq n \leq 2 \cdot 10^5$, $n-1 \leq m \leq 4 \cdot 10^5$, $1 \leq \ell \leq 2 \cdot 10^5$) — the number of vertices, the number of edges, and the size of the multiset, respectively.

The second line of each test case contains $\ell$ integers $A_1, A_2, \ldots, A_\ell$ ($1 \leq A_i \leq 10^4$) — the elements of the multiset.

Each of the following $m$ lines contains two integers $u$ and $v$ ($1 \leq u < v \leq n$) — the endpoints of an edge in the graph.

It is guaranteed that the edges form a simple, connected graph without self-loops or multiple edges.

It is guaranteed that the sum of $n$, the sum of $m$, and the sum of $\ell$ over all test cases does not exceed $2 \cdot 10^5$, $4 \cdot 10^5$, and $2 \cdot 10^5$, respectively.

## Output

For each test case, output a binary string of length $n$, where the $i$-th character is $1$ if there exists a sequence of moves ending at vertex $i$, and $0$ otherwise.


## ideas
1. 对于u，就是看是否存在一条路径，从1到u，可以经过$sum(A)$ 的路径？
2. 所以A可以直接变成S(求和)
3. 因为是一个graph，所以从1到u可能有多条路径
4. 只有奇偶性有关系
5. 假设到达u的时候，还剩余偶数条边，那么就是ok的（因为这偶数条边，可以通过到邻居，马上返回的方式消耗掉）
6. 但是如果到达u的时候，只有奇数条边，那就不行了
7. 理解错了，A中的数字，不需要全部用完