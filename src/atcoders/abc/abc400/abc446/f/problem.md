# Problem F

You are given a directed graph $G$ with $N$ vertices and $M$ edges. The vertices of $G$ are numbered as vertex $1$, vertex $2$, …, vertex $N$, and the $i$-th edge ($1 \le i \le M$) goes from vertex $U_i$ to vertex $V_i$.

Solve the following problem for $k = 1, 2, \ldots, N$.

Takahashi's goal is to delete some (possibly zero) vertices from $G$, along with all edges having those vertices as an endpoint, so that the following condition is satisfied:

**The only vertices reachable from vertex $1$ by traversing zero or more edges are vertices $1, 2, \ldots, k$.**

Determine whether he can achieve his goal, and if so, find the minimum number of vertices he needs to delete to achieve it.

---

## Constraints

- $2 \le N \le 3 \times 10^5$
- $1 \le M \le 3 \times 10^5$
- $1 \le U_i, V_i \le N$
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```
N M
U_1 V_1
U_2 V_2
⋮
U_M V_M
```

## Output

Output $N$ lines. The $i$-th line ($1 \le i \le N$) should contain the answer to the problem for $k = i$.

For each $k$, output $-1$ if it is impossible for Takahashi to achieve his goal; otherwise output the minimum number of vertices that need to be deleted.

---

## Examples

### Sample 1

**Input:**

```
5 5
1 2
2 4
2 5
4 3
4 5
```

**Output:**

```
1
2
-1
1
0
```

**Explanation:**

- **$k = 1$:** The goal can be achieved by deleting vertex $2$. Vertices $3, 4, 5$ are already unreachable from vertex $1$ without deleting them. At least one vertex must be deleted, so output $1$.
- **$k = 2$:** The goal can be achieved by deleting vertices $4$ and $5$. It is impossible with one or fewer deletions, so output $2$.
- **$k = 3$:** Takahashi cannot achieve the goal no matter how he deletes vertices. Output $-1$.
- **$k = 4$:** The goal can be achieved by deleting vertex $5$. At least one vertex must be deleted, so output $1$.
- **$k = 5$:** The goal can be achieved without deleting any vertices. Output $0$.

### Sample 2

**Input:**

```
5 5
1 1
1 2
3 1
3 1
3 5
```

**Output:**

```
1
0
-1
-1
-1
```

---

**Note:** $G$ is not necessarily connected. $G$ may have multi-edges or self-loops.


### ideas
1. 假设处理到i, 那么此时把从 <= i 的节点，连接出去的边，都激活
2. 如果1...i都能从1可达，删除那些 > i 的节点