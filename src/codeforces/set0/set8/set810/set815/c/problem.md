# C. Karen and Supermarket

[Problem link](https://codeforces.com/problemset/problem/815/C)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

## Problem

On the way home, Karen decided to stop by the supermarket to buy some
groceries.

She can spend up to `b` dollars. The supermarket sells `n` goods. The `i`-th
good costs `c_i` dollars and can be bought at most once.

Karen has `n` coupons. If she purchases the `i`-th good, she may use the
`i`-th coupon to decrease its price by `d_i`. A coupon cannot be used without
buying the corresponding good.

Constraint: for all `i ≥ 2`, to use the `i`-th coupon, Karen must also use
the `x_i`-th coupon (which may require further coupons).

What is the maximum number of goods she can buy without exceeding budget `b`?

## Constraints

- `1 ≤ n ≤ 5000`
- `1 ≤ b ≤ 10^9`
- `1 ≤ d_i < c_i ≤ 10^9`
- For `i ≥ 2`: `1 ≤ x_i < i`

## Input

```
n b
c_1 d_1
c_2 d_2 x_2
...
c_n d_n x_n
```

The first item has no `x_1`. For `i ≥ 2`, `x_i` means the `x_i`-th coupon
must be used before coupon `i` can be used.

## Output

One integer — the maximum number of different goods Karen can buy within
budget `b`.

## Samples

### Sample 1

Input:

```
6 16
10 9
10 5 1
12 2 1
20 18 3
10 2 3
2 1 5
```

Output:

```
4
```

One optimal purchase: goods `1, 3, 4, 6` with cost `1 + 10 + 2 + 2 = 15`
(coupons on `1, 3, 4`; no coupon on `6`).

### Sample 2

Input:

```
5 10
3 1
3 1 1
3 1 2
3 1 3
3 1 4
```

Output:

```
5
```

### ideas
1. 假设 x[1] = 1, 那么根据x[i], i -> x[i] -> x[x[i]] ... 最终肯定会指向1
2. 以1为root, 可以构成一个tree, 所以是在这个tree上的dp
3. dp[u][i] 表示在子树u上, 购买i个产品时,花费的最小值
4. dp[u][i] = min over [v, j] dp[v][j] + dp[u][i - j]
5. 但是咋感觉是 n * n * n 的复杂性呢?
6. 上面的分析不对, 可以在购买u (但是不够 x[u]) 的情况下, 只是放弃了d[u]
7. dp[u][0/1][i] 表示在子树u上, 根据parent状态(0 表示不购买, 1表示购买), 在子树u中购买i个item的最优解
8. 如果不购买u, dp[u][0/1][i] = dp[v][0][j] + dp[u][0/1][i - j] 的最小值
9. 如果购买u, dp[u][state][i] = dp[v][1][j] + dp[u][state][i-j] + c[u] - state * d[u]
10. 一个节点u上似乎有4个状态, (x[u], u) = (0/1, 0/1)

## Solution Summary

### 把优惠券依赖建成树

对于每个 `i >= 2`，使用优惠券 `i` 的前提是使用优惠券 `x_i`，并且 `x_i < i`。
因此，从 `x_i` 向 `i` 连边后，所有商品组成一棵以商品 `1` 为根的树。

如果使用了节点 `u` 的优惠券，那么才有可能使用它的孩子的优惠券。反过来，如果不使用
`u` 的优惠券，那么整个子树中都不能使用任何优惠券；但是仍然可以按原价购买子树中的
任意商品。

算法分别处理这两种情况。

### 不使用 `u` 的优惠券：`g[u]`

`best[u]` 收集子树 `u` 内所有商品的原价 `c`，排序后令：

```text
g[u][k] = 子树 u 中按原价购买 k 件商品的最小费用
```

显然只需要购买原价最便宜的 `k` 件，所以 `g[u]` 就是排好序的 `best[u]` 的前缀和。

这个状态覆盖了“不使用优惠券 `u`”的所有方案。此时后代优惠券都不可用，因此没有其他
优惠选择需要考虑。

### 可以使用父亲优惠券时的 DP

`dfs2(u)` 返回数组 `res`，其中：

```text
res[k] = 在父亲优惠券已经使用、因而优惠券 u 可用的前提下，
         从子树 u 购买 k 件商品的最小费用
```

对于使用优惠券 `u` 的方案，必须购买商品 `u`，费用为 `c[u] - d[u]`：

```text
res[0] = 0
res[1] = c[u] - d[u]
```

`res[0]` 表示整棵子树什么都不买。对于所有 `k >= 1` 的优惠方案，商品 `u` 一定已经使用
优惠券购买。

### 合并孩子

先递归计算最大子树孩子 `big[u]`。如果从这个孩子购买 `i-1` 件，再使用优惠券购买 `u`，
得到：

```text
res[i] = dfs2(big[u])[i-1] + c[u] - d[u]
```

选择最大孩子先初始化，可以避免对最大的一块做一次完整卷积。

对于其他孩子 `v`，令 `tmp = dfs2(v)`。当前已经处理的部分购买 `j` 件，孩子 `v` 购买
`k` 件，则进行树上背包转移：

```text
next[j+k] = min(next[j+k], res[j] + tmp[k])
```

这里 `j` 必须从 `1` 开始。因为 `tmp[k]` 允许使用优惠券 `v`，而使用它的前提是优惠券
`u` 已经使用；`res[j]` 在 `j >= 1` 时恰好保证商品 `u` 已经用优惠券购买。

`k` 需要枚举完整的 `0..size[v]`：

- `k = 0` 表示不从孩子子树购买商品；
- `k = size[v]` 表示购买孩子子树中的全部商品。

合并后，新的最大购买数量是 `cur + size[v]`，所以必须复制下标 `0..cur+size[v]`，包括
右端点。否则会丢失“购买所有已处理商品”的状态。

### 使用或放弃优惠券 `u`

所有孩子合并完成后，当前 `res[k]` 表示使用优惠券 `u` 的最优费用。还有一种选择是完全
不使用优惠券 `u`，此时费用为 `g[u][k]`。因此最终取：

```text
res[k] = min(res[k], g[u][k])
```

这样，`dfs2(u)` 同时包含：

- 使用优惠券 `u`，并允许后代继续使用优惠券；
- 不使用优惠券 `u`，整个子树只按原价购物。

根节点的优惠券不需要父亲，所以 `dfs2(1)` 就包含所有合法方案。最后从 `N` 到 `1`
寻找最大的 `k`，满足 `res[k] <= b`。

### Correctness Proof

我们对子树大小进行归纳。

叶子 `u` 只有两种相关方案：什么都不买，或者以 `c[u]-d[u]` 使用优惠券购买 `u`；再与
按原价购买的 `g[u]` 取最小值，显然正确。

假设所有孩子的 `dfs2` 都正确。若使用优惠券 `u`，则商品 `u` 必须被购买，并且每个孩子
子树可以独立选择一个由 `dfs2(v)` 表示的合法方案。树上背包卷积枚举了各孩子购买数量的
所有分配方式，并把费用相加，因此得到使用优惠券 `u` 时的最小费用。

若不使用优惠券 `u`，任何后代优惠券都不可能合法使用，只能从子树中选择若干商品按原价
购买；`g[u]` 恰好给出每个购买数量的最小费用。两类方案取最小后，`dfs2(u)` 对所有 `k`
都正确。由归纳可知根节点的数组包含全部合法购买方案，因此算法返回的最大可购买数量正确。

### Complexity

树上背包合并的总复杂度为 `O(N^2)`。

当前实现还为每个节点保存并排序整棵子树的原价列表，以构造 `g[u]`。这部分最坏需要：

- 时间：`O(N^2 log N)`；
- 空间：`O(N^2)`。

因此，按照当前代码实现，总时间复杂度为 `O(N^2 log N)`，空间复杂度为 `O(N^2)`。
