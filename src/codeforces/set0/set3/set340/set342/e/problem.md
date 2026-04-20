# Problem

Xenia the programmer has a tree consisting of `n` nodes. The tree nodes are indexed from `1` to `n`. Initially, node `1` is painted red, and all other nodes are painted blue.

The distance between two tree nodes `v` and `u` is the number of edges in the shortest path between them.

Xenia needs to process queries of two types:

1. Paint a specified blue node in red.
2. For a given node, find the closest red node and print the shortest distance to it.

Your task is to write a program that executes these queries.

## Input

The first line contains two integers `n` and `m` (`2 ≤ n ≤ 10^5`, `1 ≤ m ≤ 10^5`) — the number of nodes in the tree and the number of queries.

The next `n - 1` lines contain tree edges. The `i`-th line contains integers `ai`, `bi` (`1 ≤ ai, bi ≤ n`, `ai ≠ bi`) — an edge of the tree.

The next `m` lines contain queries. Each query is a pair `ti`, `vi` (`1 ≤ ti ≤ 2`, `1 ≤ vi ≤ n`):

- If `ti = 1`, paint blue node `vi` in red.
- If `ti = 2`, print the shortest distance from node `vi` to any red node.

It is guaranteed that the graph is a tree and all queries are valid.

## Output

For each query of type `2`, print the answer on a separate line.

## Example

**Input**

```text
5 4
1 2
2 3
2 4
4 5
2 1
2 5
1 2
2 5
```

**Output**

```text
0
3
2
```

### ideas
1. 操作1，只会有n-1次（因为1已经是red了）
2. 假设查询v, 那么有两种情况，一种是v子树中的红色节点（这些应该是高度最小的那些）
3. 另外一种是v外面的节点，假设u, 那么dist[u, v] = height[u] + height[v] - 2 * height[lca]
4. 那么就是 height[u] - 2 * height[lca] 最小的那些点
5. 如果用heavy-light分解，从1...v的路径就是所有的祖先节点，只要知道这些祖先节点上的，最小的height就可以了
6. 所以，现在问题的关键是mark一个red节点的时候，需要更新一个访问的最小值（从1到v）
7. 然后反过来在从v到1查一遍即可
8. heavy-light分解 + 区间更新/查询
9. 还是有问题～
10. 比如mark(v), 对于某个父节点u，它可能已经有一个更近的子节点(w), 那么此时就不能再继续更新上去
11. 同时，即使能更新（这个时候怎么表示呢？）
12. 对于u来说，如果v是它第一个被mark的节点，那么u得贡献 = height[v]  - 2 * height[u]
13. 但是在此之前u的贡献 = inf?
14. 

### solution

这份做法是：

- 先做一次 heavy-light decomposition
- 再在线段树上维护一类形如 `depth(red) - 2 * depth(ancestor)` 的最小值

核心恒等式是：

- 对任意红点 `r` 和查询点 `v`
- 设 `a = lca(r, v)`
- 那么

```text
dist(r, v) = depth(r) + depth(v) - 2 * depth(a)
```

如果固定查询点 `v`，那么 `a` 一定是 `v` 的某个祖先。

所以可以把答案写成：

```text
answer(v) = depth(v) + min over ancestors a of v (
    min over red nodes r whose path to root passes through a (
        depth(r) - 2 * depth(a)
    )
)
```

也就是说：

- 对每个祖先 `a`，我们只需要知道所有经过 `a` 的红点里，`depth(r)` 的最小值；
- 然后在查询 `v` 时，沿着 `v -> root` 的祖先链取
  `depth(r) - 2 * depth(a)` 的最小值即可。

#### 红点更新

当某个点 `u` 被染成红色时，令 `h = depth(u)`。

对 `u` 到根路径上的每个祖先 `a`，它都可以提供一个候选值：

```text
h - 2 * depth(a)
```

所以本质上是：

- 把从 `root` 到 `u` 的整条路径上的所有点，
- 都用值 `h` 做一次 `min` 更新。

代码里借助 HLD，把这条路径拆成若干条重链区间，在每段区间上做线段树区间取最小。

#### 为什么线段树里能这样维护

线段树按 HLD 的 dfs 序建。

对位置对应的节点 `x`，它的深度是 `depth(x)`。  
如果一次更新传入某个红点深度 `h`，那么这个位置应该尝试变成：

```text
h - 2 * depth(x)
```

所以线段树节点里需要知道：

- 这一段里点的最大深度 `arr`
- 这一段当前维护到的最优值 `val`

这里用了一个很巧的性质：

- HLD 的一段链区间在 dfs 序里是连续的；
- 同一棵线段树节点覆盖的这段区间里，最深的点深度就是右端点，也就是这段的最大深度；
- 当我们对整段统一做一次更新 `h` 时，这段内最优候选会落在最深的位置，
  因为 `h - 2 * depth(x)` 随深度单调变小。

因此对整段可以直接记：

```text
val = min(val, h - 2 * maxDepthOfThisSegment)
```

再靠懒标记把这个 `h` 继续下传给子区间。

这就是代码里：

```go
tr.val[i] = min(tr.val[i], v - 2*tr.arr[i])
```

的含义，其中：

- `v` 是红点的深度
- `arr[i]` 是该线段覆盖范围内的最大深度

#### 查询

查询点 `u` 时：

1. 用 HLD 把 `u -> root` 路径拆成若干条链区间
2. 在线段树上查询这些区间里的最小 `val`
3. 最后再加上 `depth(u)`

即：

```text
answer = depth(u) + minValOnPath(root, u)
```

#### 初始条件

节点 `1` 一开始就是红色，所以先对它做一次更新。

#### 复杂度

- 预处理 HLD：`O(n)`
- 每次更新：`O(log^2 n)`
- 每次查询：`O(log^2 n)`

总复杂度：

- `O((n + m) log^2 n)`

对 `n, m <= 1e5` 足够通过。
