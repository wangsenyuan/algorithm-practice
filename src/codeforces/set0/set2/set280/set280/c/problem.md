### Problem

You are given a rooted tree with `n` nodes, numbered from `1` to `n`. The root is node `1`.

You play the following game:

- Repeatedly choose one of the remaining nodes `v` uniformly at random (each remaining node is equally likely).
- Remove the entire subtree rooted at `v` (including `v` itself).
- The game ends when the tree becomes empty — equivalently, when the step that chooses node `1` is performed.

Find the expected number of steps in this game.

### Input

- The first line contains an integer `n` (`1 ≤ n ≤ 10^5`).
- The next `n - 1` lines each contain two integers `a_i b_i` (`1 ≤ a_i, b_i ≤ n`, `a_i ≠ b_i`), edges of the tree.
- It is guaranteed the graph is a tree.

### Output

Print a single real number — the expected number of steps.
Your answer will be accepted if the absolute or relative error does not exceed `1e-6`.

### Examples

**Input**
```text
2
1 2
```
**Output**
```text
1.50000000000000000000
```

**Input**
```text
3
1 2
1 3
```
**Output**
```text
2.00000000000000000000
```

### Note

- Sample 1: There are two possibilities — remove the root immediately (1 step), or remove the leaf first, then the root (2 steps), each with probability `1/2`, so expectation is `1*(1/2) + 2*(1/2) = 1.5`.
- Sample 2: Two of the three equally likely first choices reduce to Sample 1, and one removes the root directly:  
  `1*(1/3) + (1+1.5)*(2/3) = 2`.

### Summary

- Start with a rooted tree at node `1`.
- Repeatedly pick a remaining node uniformly at random and delete its whole subtree.
- The game ends when the root is picked (tree becomes empty).
- Output the expected number of steps, within `1e-6` error.

### 思路 / Key insights

把“每一步在剩余节点里等概率随机选一个”看成：给每个节点 `v` 分配一个独立的随机时间戳 `T_v ~ U(0,1)`，然后按 `T` 从小到大依次被“选中”（等价于随机排列）。

当某个节点 `v` 被选中时，会删除 `v` 的整棵子树，因此子树里之后的节点不会再被选中。于是**真正会发生删除操作的节点**恰好是那些在根到它的路径上时间戳最小的节点：

- `v` 会成为一次操作 ⟺ `T_v = min{ T_u | u 在 path(1 -> v) 上 }`

由于路径上这些时间戳是独立连续分布，最小值落在路径上任意一个节点的概率相同。设 `depth(v)` 为 `v` 到根的边数，则路径长度为 `depth(v)+1`，因此：

\[
P(v\ \text{被操作}) = \frac{1}{depth(v)+1}
\]

期望步数等于所有节点被操作概率之和：

\[
\mathbb{E}[\#steps] = \sum_{v=1}^{n} \frac{1}{depth(v)+1}
\]

实现上只需从根 `1` 做一次 DFS/BFS 求深度并累加上述式子即可，时间复杂度 `O(n)`。

