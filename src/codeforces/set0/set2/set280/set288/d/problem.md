# D — Vertex-disjoint path pairs in a tree

## Problem description

Little penguin Polo has a **tree**: an undirected, connected graph with **n** vertices and **n − 1** edges. Vertices are numbered **1 … n**.

He wants the number of **quadruples** **(a, b, c, d)** describing two paths in the tree, such that:

1. **1 ≤ a < b ≤ n**
2. **1 ≤ c < d ≤ n**
3. There is **no** vertex that lies on **both** the (unique) shortest path from **a** to **b** **and** the shortest path from **c** to **d**.

Here “shortest path” means fewest edges (in a tree, there is exactly one simple path between any two vertices).

## Input

- The first line contains an integer **n** (1 ≤ n ≤ 8·10⁴) — the number of nodes.
- Each of the next **n − 1** lines contains two integers **uᵢ**, **vᵢ** (1 ≤ uᵢ, vᵢ ≤ n, uᵢ ≠ vᵢ) — an edge of the tree.

The graph is guaranteed to be a tree.

## Output

Print one integer — the answer.

**C++ note (legacy):** avoid `%lld` for 64-bit I/O; prefer `cin`/`cout` or `%I64d` on older judges.

## Example

**Input:**

```
4
1 2
2 3
3 4
```

**Output:**

```
2
```

## Ideas（草稿）

1. 一共有 **w = n(n − 1)/2** 条 path（无序顶点对对应一条树上的路）。
2. 所以一共有 **w(w − 1)/2** 对 path 的组合（朴素太大）。
3. 能不能 **divide and conquer**？
4. 固定一个点 **r**，把图分成若干块，每块单独算再 merge？
5. merge 具体怎么做还不清楚。
6. 换思路：按某条边 **e(u, v)** 把 tree 分成两部分。
7. 然后 merge 两边的结果 + 两边的 path 数量相乘 + …
8. **(u, v)** 这条边怎么处理？若以 **u** 为 root，延伸到子树 **v**，这样的结构里和 **sz(v)** 有关。
9. **dp[u][0]**：子树 **u** 里的答案，且用到的路径都不经过 **u**（草稿含义）。
10. **dp[u][1]**：子树 **u** 里的答案，且用到的路径可能经过 **u**（草稿含义）。
11. **dp[u][0] = Σ dp[v][1]**（对 **u** 的儿子 **v**）+ 与边数相关的交叉项（未写完）。
12. **dp[u][1]**：经过 **u** 的路径分两类：**u** 是端点，或 **u** 是中间点。
13. 中间点情况要看这条路径把 **u** 的邻域分成多少组；若某组有 **w** 个点，则组内 path 数 **w(w − 1)/2**。
14. **u** 的贡献 = “包含 **u** 的这条边/路径”与所有其他合法路径组合（草稿）。
15. 也许可以反过来：**总数 − 不合法对数**？
16. 若子树 **v** 的大小为 **sz[v]**，则 **g(v) = sz[v](sz[v] − 1)/2**（子树内无序顶点对 / path 数的一种计数）。
17. 经过 **u**、但不进入 **v** 的路径数量似乎可写成某个 **f(u)**（草稿）。
18. 猜测子树 **v** 的贡献形如 **f(u) · g(v)**？（待验证。）

## Editorial (excerpt; formulas omitted)

Root the tree (e.g. at vertex **1**). After fixing the **first** path and removing all its vertices (and incident edges), the remaining graph is a forest; let the component sizes be **c₁, …, cₖ**. The number of admissible **second** paths can be written as a **closed form in c₁, …, cₖ** (see the official Codeforces PDF — one step in the original text is only rendered as “equal to .” here).

That yields an **O(n²)** baseline: enumerate the first path and evaluate the formula for the second. The full solution speeds this up to **O(n)** with DFS: for a fixed vertex as the “last” vertex of the first path, aggregate contributions from inside the current subtree and from the complement; use subtree sizes and **partial sums** along adjacency lists to keep the recurrence linear.

**Alternatives mentioned in the original note:** count **bad** pairs of paths and subtract from the total; or explore a divide-and-conquer approach.
