# Problem

**题意：** 给定一棵 $n$ 个点的树。可以删掉任意条边（可以不删）。删边后图会变成若干连通块，定义得分为这些连通块大小的**乘积**。求可能得到的最大得分。

## Input

- The first line contains an integer $n$ ($1 \le n \le 700$) — the number of vertices.
- The next $n - 1$ lines describe the edges. Each line has two integers $a_i, b_i$ ($1 \le a_i, b_i \le n$) — the indices of the vertices joined by that edge. The input describes a tree.

## Output

Output a single integer — the maximum product of the sizes of the connected components after removing some edges.

## Examples

### Example 1

**Input:**

```
5
1 2
2 3
3 4
4 5
```

**Output:**

```
6
```

### Example 2

**Input:**

```
8
1 2
1 3
2 4
2 5
3 6
3 7
6 8
```

**Output:**

```
18
```

### Example 3

**Input:**

```
3
1 2
1 3
```

**Output:**

```
3
```

### editorial
Lemma. In one of optimal solutions there are no simple paths of length 3.
Proof. We can remove the middle edge from such a path. The connected component will split into two components of sizes a and b, where a ≥ 2, b ≥ 2, and therefore ab ≥ a + b.

We'll root the tree and calculate recursively the numbers hv = the solution of the problem for the subtree with the root v, and fv = the product of hu for all children of v. If v is a leaf, then hv = fv = 1.

We show how to calculate hv, given the solution for all subtrees of v. Consider the connected component of v in an optimal solution. It follows from the lemma that the component has one of the following types:
1. The single vertex v.
2. The vertex v and several children of v.
3. The vertex v, one child of v — w, and several children of w.

In the first case the result of the game is fv.

In the second case it equals Πfi · Πhj · k = Π(fi / hi) · fv · k, where i iterates over children belonging to the connected component, j iterates over the rest children, and k is the size of the component. Since we want to maximize the result, we're interested in children with the largest value of fi / hi. Therefore, the largest possible result in this case equals the maximum value of Πi ≤ s  (fi / hi) · fv · (s + 1), where it's supposed that the children are sorted in descending order of fi / hi.

In the third case we can use a similar reasoning for every child w. The best result will be the maximum of the expression fv · (fw / hw) · Πi ≤ s (fi / hi) · (s + 2) as w iterates over children of v, and s iterates from 1 to the number of children of w; note that the children of w have already been sorted in the previous step.

Therefore, the number of operations necessary to calculate fv is proportional to the total number of children and grandchildren of v, which is less than n. The complexity of the algorithm, then, is O(n2) (ignoring operations with long numbers).

## Explanation

下面按我们代码里的状态重新解释一遍。

### 1. 先看最关键的结构性质

把树任意删边后，会得到若干连通块，答案是这些连通块大小的乘积。

题解里的引理很重要：

- 在某个最优方案里，不会存在长度至少为 `3` 的简单路径完全留在同一个连通块里。

原因是：

- 如果一个连通块里有一条长度 `3` 的路径，那么这条路径中间那条边删掉以后，这个连通块会裂成两个大小分别为 `a` 和 `b` 的部分。
- 因为两边至少都含有 `2` 个点，所以 `a >= 2, b >= 2`。
- 原来这个连通块对答案的贡献是 `a + b`，删开后变成 `a * b`。
- 而 `a * b >= a + b`。

所以，在最优解里，每个保留下来的连通块都非常“浅”，最多只会长成下面三种样子之一：

1. 只有当前点 `u` 自己。
2. `u` 加上一些儿子。
3. `u` 加上某个儿子 `v`，再加上 `v` 的一些儿子。

这就是为什么 DP 只需要看“儿子”和“孙子”，不需要继续向下展开。

### 2. 状态定义

把树以 `1` 号点为根。

对每个点 `u`，定义 3 个量：

- `f[u]`
  - `u` 的所有儿子子树各自独立处理时，得到的乘积。
  - 公式：`f[u] = ∏ h[child]`

- `h[u]`
  - 只看 `u` 的整棵子树，并且要求 `u` 所在的连通块在这棵子树内部“封口”，也就是 `u` 不再和父亲连。
  - 这是 `solve()` 最后真正要的量，根节点答案就是 `h[root]`。

- `g[u]`
  - 只看 `u` 的整棵子树，但这次假设 `u` 已经和父亲连在同一个连通块里了。
  - 也就是说，`u` 这个连通块里已经提前占了一个来自父亲的点，所以它的大小基数比 `h[u]` 多 `1`。
  - 这个状态是为了处理“父亲 `p` 选择把 `u` 接进来，同时再从 `u` 的儿子里挑一些点”的情况。

代码里正是：

- `f[u] = product(h[v])`
- `h[u] = u` 不向父亲延伸时的最优值
- `g[u] = u` 已经接到父亲上时的最优值

### 3. 为什么要按 `f[v] / h[v]` 排序

假设当前在处理点 `u`。

如果不把某个儿子 `v` 接进 `u` 的连通块，那么 `v` 子树给出的贡献是：

- `h[v]`

如果把 `v` 接进来，那么 `v` 这个点就不能作为一个独立连通块的根去算了，它对子树贡献改成：

- `f[v]`

所以，“把 `v` 接进来”这件事，相当于把原来的贡献 `h[v]` 换成了 `f[v]`。

收益倍率是：

- `f[v] / h[v]`

如果要从若干儿子里挑一些接进来，显然应该优先挑这个倍率更大的。  
所以代码里会把儿子按 `f[v] / h[v]` 从大到小排序。

比较两个分数时，为了避免浮点误差，代码用交叉乘法比较：

- 比较 `f[a] / h[a]` 和 `f[b] / h[b]`
- 等价于比较 `f[a] * h[b]` 和 `f[b] * h[a]`

### 4. `h[u]` 的转移

先记：

- `f[u] = ∏ h[v]`

这对应“`u` 不和任何儿子连在一起”，所有儿子子树都各自独立。

#### 情况 A：连通块只有 `u` 自己

那么 `u` 这个连通块大小是 `1`，总乘积就是：

- `f[u]`

所以初始化：

- `h[u] = f[u]`

#### 情况 B：`u` 和前 `k` 个最优儿子连成一个块

如果把排序后前 `k` 个儿子接进来，那么：

- 原来这 `k` 棵子树贡献是 `h[v]`
- 现在变成 `f[v]`
- 连通块大小从 `1` 变成 `k + 1`

总乘积变成：

- `f[u] * ∏ (f[v] / h[v]) * (k + 1)`

代码里用一个滚动变量 `cur` 维护：

- 初始 `cur = f[u]`
- 每加入一个儿子 `v`，做一次 `cur = cur / h[v] * f[v]`

于是：

- 更新 `h[u]` 时尝试 `cur * (k + 1)`
- 更新 `g[u]` 时尝试 `cur * (k + 2)`

这里 `g[u]` 多出来的 `1`，就是“父亲那个点已经和 `u` 连上了”。

#### 情况 C：`u` 和某个儿子 `v` 连，同时继续吃进 `v` 的一些儿子

这是引理允许的第三种形态：

- 连通块形状是 “`u - v - (v 的若干儿子)`”

如果固定选儿子 `v`，那么在 `u` 这一层：

- 其他儿子还是各自独立，所以先从 `f[u]` 出发
- 但 `v` 不再作为独立根处理，所以把 `h[v]` 换成“`v` 已经接到父亲时的最优值” `g[v]`

所以候选值是：

- `f[u] / h[v] * g[v]`

代码里就是：

```go
updateMax(&h[u], divMulBig(f[u], h[v], g[v]))
```

这一步非常关键。  
如果没有 `g[v]`，就没法把“跨过一层继续接孙子”的情况正确表示出来。

### 5. 为什么 `g[u]` 这样定义就够了

`g[u]` 表示：

- `u` 已经和父亲连上了
- 在 `u` 的子树内部，还能怎么扩展，才能使总乘积最大

这时允许的连通块形状只有两种：

1. `parent - u`
2. `parent - u - 若干儿子`

不能再继续往下延伸两层，否则整条路径就会超过引理允许的深度。

所以 `g[u]` 只需要考虑：

- 一个基点 `2`，表示 `parent + u`
- 再按 `f[v] / h[v]` 从大到小接若干个儿子

因此：

- 初始值是 `f[u] * 2`
- 每加入一个儿子后，候选变成 `cur * (已接儿子数 + 2)`

这正是代码中的：

```go
g[u] = mulInt(f[u], 2)
...
updateMax(&g[u], mulInt(cur, i+3))
```

因为 `i` 从 `0` 开始，所以 `i + 3 = (i + 1) + 2`。

### 6. 用样例 1 手推

样例 1 是一条链：

```text
1 - 2 - 3 - 4 - 5
```

以 `1` 为根后，结构是：

```text
1
└── 2
    └── 3
        └── 4
            └── 5
```

我们从叶子往上算。

#### 点 5

叶子点最简单：

- `f[5] = 1`
- `h[5] = 1`
- `g[5] = 2`

解释：

- `h[5] = 1`：自己单独一个块
- `g[5] = 2`：如果它已经和父亲连上，那这个块大小就是 `2`

#### 点 4

它只有一个儿子 `5`。

先算：

- `f[4] = h[5] = 1`

初始化：

- `h[4] = 1`
- `g[4] = 2`

考虑把儿子 `5` 接进来：

- `cur = 1 / h[5] * f[5] = 1`

更新：

- `h[4] = max(1, 1 * 2) = 2`
- `g[4] = max(2, 1 * 3) = 3`

再考虑“4 和 5 连，同时继续往 5 的下面延伸”：

- `f[4] / h[5] * g[5] = 1 * 2 = 2`

所以最终：

- `h[4] = 2`
- `g[4] = 3`

这很合理：

- 在子树 `{4,5}` 中，最优是连成一个大小 `2` 的块，所以 `h[4] = 2`
- 如果 `4` 还和父亲连着，那么 `{parent,4,5}` 大小为 `3`，所以 `g[4] = 3`

#### 点 3

它只有一个儿子 `4`。

先算：

- `f[3] = h[4] = 2`

初始化：

- `h[3] = 2`
- `g[3] = 4`

把儿子 `4` 接进来：

- `cur = 2 / h[4] * f[4] = 2 / 2 * 1 = 1`

更新：

- `h[3] = max(2, 1 * 2) = 2`
- `g[3] = max(4, 1 * 3) = 4`

再考虑“3 连 4，再向下吃进 5”：

- `f[3] / h[4] * g[4] = 2 / 2 * 3 = 3`

于是：

- `h[3] = 3`
- `g[3] = 4`

这个 `3` 对应的方案就是：

- 连通块 `{3,4,5}` 大小为 `3`

#### 点 2

它只有一个儿子 `3`。

先算：

- `f[2] = h[3] = 3`

初始化：

- `h[2] = 3`
- `g[2] = 6`

把儿子 `3` 接进来：

- `cur = 3 / 3 * f[3] = 2`

更新：

- `h[2] = max(3, 2 * 2) = 4`
- `g[2] = max(6, 2 * 3) = 6`

再考虑“2 连 3，再继续往下”：

- `f[2] / h[3] * g[3] = 3 / 3 * 4 = 4`

所以：

- `h[2] = 4`
- `g[2] = 6`

#### 点 1

它只有一个儿子 `2`。

先算：

- `f[1] = h[2] = 4`

初始化：

- `h[1] = 4`

把儿子 `2` 接进来：

- `cur = 4 / 4 * f[2] = 3`

更新：

- `h[1] = max(4, 3 * 2) = 6`

再考虑“1 连 2，再继续往下”：

- `f[1] / h[2] * g[2] = 4 / 4 * 6 = 6`

最终：

- `h[1] = 6`

答案就是 `6`。

### 7. 样例 1 的最优方案到底是什么

链 `1-2-3-4-5` 的最优方案之一是：

- 保留边 `(1,2)`
- 删掉边 `(2,3)`
- 保留边 `(3,4)`
- 保留边 `(4,5)`

这样得到两个连通块：

- `{1,2}`，大小 `2`
- `{3,4,5}`，大小 `3`

乘积：

- `2 * 3 = 6`

这就是输出 `6` 的原因。

### 8. 代码里的关键几行对应什么含义

#### 先算 `f[u]`

```go
f[u] = big.NewInt(1)
for _, v := range adj[u] {
    if v == p {
        continue
    }
    dfs(u, v)
    children[u] = append(children[u], v)
    f[u] = mulBig(f[u], h[v])
}
```

含义：

- 所有儿子先各自独立做最优，乘起来就是 `f[u]`

#### 儿子按收益排序

```go
sort.Slice(children[u], func(i, j int) bool {
    a, b := children[u][i], children[u][j]
    left := mulBig(f[a], h[b])
    right := mulBig(f[b], h[a])
    return left.Cmp(right) > 0
})
```

含义：

- 按 `f[v] / h[v]` 从大到小排序

#### 枚举把前若干个儿子接进来

```go
h[u] = cloneBig(f[u])
g[u] = mulInt(f[u], 2)

cur := cloneBig(f[u])
for i, v := range children[u] {
    cur = divMulBig(cur, h[v], f[v])

    updateMax(&h[u], mulInt(cur, i+2))
    updateMax(&g[u], mulInt(cur, i+3))
}
```

含义：

- `h[u]` 里的 `i+2` 表示：`u` 本身加上前 `i+1` 个儿子
- `g[u]` 里的 `i+3` 表示：父亲 + `u` + 前 `i+1` 个儿子

#### 枚举“选一个儿子继续向下延伸”

```go
for _, v := range children[u] {
    updateMax(&h[u], divMulBig(f[u], h[v], g[v]))
}
```

含义：

- 先把 `v` 从“独立子树最优值 `h[v]`”中拿出来
- 再换成“`v` 已和父亲 `u` 相连时的最优值 `g[v]`”

这就是第三种结构：

- `u - v - (v 的若干儿子)`

### 9. 复杂度

每个点 `u`：

- 要处理自己的所有儿子
- 还要在转移里考虑每个儿子对应的 `g[v]`

整体复杂度是：

- `O(n^2)`

这里 `n <= 700`，可以通过。

由于答案可能非常大，代码用 `big.Int` 保存结果。
