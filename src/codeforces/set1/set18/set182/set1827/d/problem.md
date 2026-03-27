You are given a tree (an undirected connected acyclic graph) which initially only contains vertex `1`. There will be several queries to the given tree. In the `i`-th query, vertex `i + 1` will appear and be connected to vertex `pi` (`1 <= pi <= i`).

After each query, find the **minimum** number of operations required so that the **current** tree has **two** centroids. In one operation, you can add one vertex and one edge to the tree such that it remains a tree.

A vertex is called a **centroid** if its removal splits the tree into subtrees with at most `floor(n / 2)` vertices each, where `n` is the number of vertices of the tree. For example, the centroid of the following tree is `3` because the largest subtree after removing the centroid has `2` vertices.

In the next tree, vertices `1` and `2` are both centroids.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 <= t <= 10^4`). The description of the test cases follows.

The first line of each test case contains a single integer `n` (`2 <= n <= 5 * 10^5`) — the number of nodes of the final tree.

The second line of each test case contains `n - 1` integers `p1, p2, ..., p(n-1)` (`1 <= pi <= i`) — the index of the vertex that is connected to vertex `i + 1`.

It is guaranteed that the sum of `n` over all test cases does not exceed `5 * 10^5`.

## Output

For each test case, output `n - 1` integers. The `i`-th integer is the answer to the `i`-th query — the minimum number of operations required to make the current tree have two centroids.

We can show that an answer always exists.

## Example

**Input**

```
5
2
1
3
1 1
4
1 2 3
7
1 2 3 2 5 2
10
1 2 2 4 5 5 7 8 9
```

**Output**

```
0
0 1
0 1 0
0 1 0 1 2 3
0 1 2 1 0 1 0 1 2
```


### ideas

本题会用到的性质：
1. 当树的大小是偶数时，才有两个重心，且这两个重心一定相邻。断开两个重心之间的边，可以得到两个大小为 n/2 的连通块。
2. 在树上添加一个叶子，重心要么不动，要么移动到某个邻居上。

设现在有 n 个节点，设我们在某棵大小为 sz 的子树内添加了 x 个点，同时在子树外添加了 y 个点。
我们希望这棵子树的大小 * 2 = 节点总数。
即 (sz + x) * 2 = n + x + y。
化简得 x = n - sz*2 + y。

由上式可知，y 越小，x 也越小。那么 y=0 是最优的，换句话说，只在子树内加点是最优的。
所以需要加 n - sz*2 个点。

核心思路：往树上添加节点，如果某棵子树的大小恰好等于节点总数的一半（性质 1），即符合要求。
此时有两个重心，一个是子树的根，另一个是子树的根的父节点。或者反过来说，这棵子树是重心的某个儿子子树。
于是，我们只需考察以重心为根时，重心的儿子子树。

由 n - sz*2 可知，sz 越接近 n/2 越好。注意 sz 不能大于 n/2，否则与 sz 是重心的儿子子树的大小相矛盾。
因此，我们除了需要维护重心 ct，还需要维护 ct 的最大儿子子树的大小 maxSonSize。
答案就是 n - maxSonSize*2。

如何维护 maxSonSize？
添加一个点后，有可能 ct 的另一棵儿子子树的大小变大了，maxSonSize 可能换成另一棵儿子子树的大小了。如何判断？
为方便描述，下面把添加节点描述成激活这棵大小为 n（n 是输入的 n）的树中的节点。一开始只有 1 被激活。
分类讨论：
- 情况一：如果以 1 为根时，激活的点 i 在 ct 子树中，那么计算 ct 的儿子 son，使得 i 在 son 子树中。计算 son 子树的大小 size。
- 情况二：如果以 1 为根时，激活的点 i 不在 ct 子树中，那么 ct 上面的子树大小为 size = (当前被激活的节点总数) - (ct 子树的大小)。
然后用 size 更新 maxSonSize 的最大值。

如果更新后，maxSonSize > (当前被激活的节点总数) / 2，那么必须移动重心。
根据性质 2：
- 如果是情况一，那么把 ct 更新为 son。
- 如果是情况二，那么把 ct 更新为 ct 的父节点。
此外，由于我们只加了一个点，所以 maxSonSize > (当前被激活的节点总数) / 2 实际上只比 (当前被激活的节点总数) / 2 多一，那么移动重心后，maxSonSize 就刚好变成 (当前被激活的节点总数) / 2 了。

最后剩下两个实现细节：
1. 如何求出 ct 的儿子 son，使得 i 在 son 子树中？这可以用 LC1483. 树节点的第 K 个祖先 的倍增模板，计算从 i 往上跳到深度 depth[ct] + 1 的位置，即为 ct 的儿子。如果这个节点的父节点不是 ct，说明 i 不在 son 子树中（在 ct 的上面）。
2. 如何求出子树中的被激活的节点个数？计算整棵树的 DFS 时间戳，子树对应一段连续的时间戳，把未被激活的点视作 0，被激活的点视作 1，子树中的被激活的点的个数即为子数组中的 1 的个数。激活节点是单点修改，查询子树中的被激活的节点个数是区间查询，这可以用树状数组维护。
