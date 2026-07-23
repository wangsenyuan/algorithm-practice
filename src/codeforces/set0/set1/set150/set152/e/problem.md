# E. Garden

[Problem link](https://codeforces.com/problemset/problem/152/E)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Vasya's garden is an `n × m` grid. Square `(i, j)` has `a_{ij}` flowers. Covering
a square with concrete kills all flowers there.

There are `k` important squares (buildings). Cover some squares with concrete
so that:

1. every important square is covered;
2. every pair of important squares is connected via 4-adjacent concrete squares;
3. the total number of dead flowers is minimized.

## Constraints

- `1 <= n, m <= 100`
- `n * m <= 200`
- `1 <= k <= min(n * m, 7)`
- `1 <= a_{ij} <= 1000`
- Important squares have distinct coordinates

## Input

```text
n m k
a_11 ... a_1m
...
a_n1 ... a_nm
x_1 y_1
...
x_k y_k
```

## Output

First line: the minimum number of dead flowers.

Then `n` lines of `m` characters: `X` for concrete, `.` otherwise.

Any optimal plan is accepted.

## Examples

### Sample 1

```text
Input
3 3 2
1 2 3
1 2 3
1 2 3
1 2
3 3

Output
9
.X.
.X.
.XX
```

### Sample 2

```text
Input
4 5 4
1 4 5 1 2
2 2 2 2 7
2 4 1 4 5
3 2 1 7 1
1 1
1 5
4 1
4 4

Output
26
X..XX
XXXX.
X.X..
X.XX.
```


### ideas
1. 看起来这个是TSP问题, dp[state][i] 目前在哪个位置, 然后找到更新下一个状态
2. 但是问题是, 部分路径是可以被重用的
3. 比如在sample 2中, 第二行的前3列, 被(1, 5) (4, 4) 重用了, 只需要计算一次
4. dp[r][c][mask] 表示到(r, c)时,由mask表示的状态, 且上半部分的建造点,被覆盖时的最优解
5. 这样子不对的, 状态也恢复不出来
6. 如果只有两个点, 那么很容易, 如果存在3个点, 有没有可能那个*中心*点肯定在其中的一条最短路径上?
7. 证明不了(似乎就算是这样, 也没啥用)
8. 感觉要整行dp, dp[row][mask][?] 表示到row的时候, 这一行的状态是mask, 还必须知道, row行之上的所有点的连接情况(是否都连接在一起了)

### editorial
The solution of this problem is based on dynamic programming. dp[mask][v] — the value of the minimum correct concrete cover, if we consider as important elements only elements of the mask mask, and there are additionally covered the vertex v = (i, j) of the field.

There are two types of transfers.

First of all we can, as if to cut coverage on the vertex v. Then you need to go through subpattern of vertex submask, which will go to the left coverage and make an optimizing transfer. Update dp[mask][v] with the value dp[submask][v] + dp[mask ^ submask][v] - cost(v).

Second, perhaps in the vertex v in the optimal coverage mask mask, which covers the vertex v, you can not make the cut separating the set of vertices. In this case, this vertex forms something a kind of <>. And there a vertex u exists, on which we can make the cut, with the whole shortest path from a vertex u to v belongs to the optimal coverage. Let's precalculate the shortest paths between all pairs of cells. Now to make this transition, we should count the value of dynamics dp[mask][v] for all vertices v only on the basis of the first transition. Now you can make the second transition. For all u, dp[mask][v], update the value of dp[mask][u] + dist(v, u) - cost(u).

Let's process separately state in which exactly one bit in the mask, and the vertex which corresponding to this bit is equal to v. In this case the answer is equal to cost(v), of course.

## Solution Summary

This is a vertex-weighted Steiner tree problem with at most `k = 7` required
vertices.

Treat every grid cell as a graph vertex. Two vertices are adjacent when their
cells share a side. If vertex `v` represents cell `(r,c)`, then

```text
cost(v) = a[r][c].
```

Selecting that vertex means covering the cell and killing its flowers.

Let `V = n*m`, and flatten `(r,c)` into `v = r*m+c`.

### Vertex-weighted shortest paths

Define `dist[u][v]` as the minimum total cost of a path from `u` to `v`,
including the costs of both endpoints.

Initialize:

```text
dist[u][u] = cost(u)
dist[u][v] = cost(u) + cost(v)    if u and v are adjacent
```

All other distances start at infinity. Floyd-Warshall uses

```text
dist[i][j] =
    min(dist[i][j],
        dist[i][x] + dist[x][j] - cost(x)).
```

The two paths both include `x`, so `cost(x)` is subtracted once.

The implementation also stores:

```text
next[u][v]
```

which is the first vertex after `u` on a minimum-cost path from `u` to `v`.
For adjacent vertices, `next[u][v] = v`. If Floyd-Warshall improves the path
from `i` to `j` through `x`, then:

```text
next[i][j] = next[i][x].
```

Starting at `u` and repeatedly moving to `next[current][v]` reconstructs every
cell on the chosen shortest path to `v`.

### Subset DP

Let:

```text
dp[mask][v]
```

be the minimum cost of a connected concrete region that:

- contains every important cell represented by `mask`;
- also contains vertex `v`.

The region may contain additional non-important cells and may branch. Vertex
`v` is an anchor or meeting point; it does not need to be important.

For important vertex `terminal[i]`, initialize:

```text
dp[1<<i][terminal[i]] = cost(terminal[i]).
```

There are two transitions.

#### 1. Merge two terminal groups at the same vertex

Split `mask` into two nonempty disjoint parts `sub` and `mask^sub`. If their
optimal regions both contain `v`, their union is connected and covers all
terminals in `mask`:

```text
dp[mask][v] =
    min(dp[mask][v],
        dp[sub][v]
        + dp[mask^sub][v]
        - cost(v)).
```

Both states counted `v`, so its cost is subtracted once.

#### 2. Move the anchor through a shortest path

Suppose a connected region for `mask` already contains `u`. Attach a
minimum-cost path from `u` to `v`:

```text
dp[mask][v] =
    min(dp[mask][v],
        dp[mask][u]
        + dist[u][v]
        - cost(u)).
```

Vertex `u` is included in both terms, so its cost is subtracted once.

For each `mask`, perform all subset merges first and then apply the
shortest-path transition for every pair `(u,v)`.

### Why the transitions are sufficient

Because all vertex costs are positive, an optimal connected region can be
assumed to be a tree: a cycle can be removed without disconnecting the
required vertices.

Root an optimal tree at its anchor `v`.

- If the tree branches at `v`, its terminal-bearing branches can be divided
  into two nonempty terminal groups. The merge transition joins the optimal
  solutions for those groups at `v`.
- If the useful tree does not branch at `v`, follow its path to the next
  terminal or branching vertex `u`. The unbranched segment can be replaced by
  a minimum-cost path between `u` and `v`, which is exactly the
  shortest-path transition.

Repeating this decomposition eventually reaches singleton-terminal base
states. Therefore, the DP considers an optimal tree.

Conversely, every transition preserves connectivity and contains all
terminals represented by its mask. Thus every DP construction is feasible,
so the minimum DP value is exactly the optimum.

### Recovering the concrete cells

Each `state` stores how it was obtained:

```text
fromMask = -1
    singleton-terminal base state

fromMask = mask
    shortest-path transition from fromPos to the current vertex

0 < fromMask < mask
    merge of fromMask and mask^fromMask at the current vertex
```

For a shortest-path state, use `next` to mark every cell from `fromPos` to the
current vertex, then recursively recover the state at `fromPos`.

For a merge state, recursively recover both subset states at the same meeting
vertex. Marking the union of all recovered cells produces the required `X`
cells.

Let:

```text
all = (1<<k)-1.
```

Choose the vertex `v` minimizing `dp[all][v]`, reconstruct from
`(all,v)`, and print the resulting grid.

### Complexity

Let `V = n*m <= 200`.

- Floyd-Warshall: `O(V^3)`.
- Subset merging: `O(3^k * V)`.
- Shortest-path DP transitions: `O(2^k * V^2)`.
- Total space: `O(V^2 + 2^k * V)`.

These bounds fit `k <= 7` and `V <= 200`.
