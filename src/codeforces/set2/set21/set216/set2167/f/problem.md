# Problem

Behruzbek is given a tree with `n` nodes. For a chosen root `r`, define:

- Consider every set of `k` distinct nodes.
- For each set, compute its LCA in the tree rooted at `r`.
- Let `Sr` be the set of all distinct nodes that appear as these LCAs.

The **cuteness** for root `r` is `|Sr|`.

The **kawaiiness** is:

`sum_{r=1..n} |Sr|`.

Find this value for each test case.

A tree is a connected graph without cycles.

## Input

The first line contains integer `t` (`1 <= t <= 10^4`) — the number of test cases.

For each test case:

- The first line contains two integers `n` and `k` (`2 <= k <= n <= 2 * 10^5`) — number of vertices and subset size.
- The next `n - 1` lines each contain two integers `u` and `v` (`1 <= u, v <= n`, `u != v`) — an edge of the tree.

It is guaranteed that these edges form a tree.

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output one integer — `sum_{r=1..n} |Sr|`.

## Example

**Input**

```text
4
2 2
1 2
5 3
1 2
1 3
1 4
1 5
6 3
1 2
1 3
2 4
2 5
3 6
10 5
5 6
4 9
3 9
2 6
2 8
8 9
6 10
1 6
4 7
```

**Output**

```text
2
9
17
35
```

## Note

Let `f(i) = |Si|`.

For the third example:

- Root `1`: only nodes `1` and `2` can be obtained, so `f(1)=2`.  
  Examples: `LCA(4,5,6)=1`, `LCA(2,4,5)=2`.
- Root `2`: only nodes `1` and `2` can be obtained, so `f(2)=2`.  
  Examples: `LCA(1,3,6)=1`, `LCA(1,4,5)=2`.
- Root `3`: `f(3)=3`. Example: `LCA(2,4,6)=3`.
- Root `4`: `f(4)=3`. Example: `LCA(1,3,5)=2`.
- Root `5`: `f(5)=3`. Example: `LCA(3,4,6)=2`.
- Root `6`: `f(6)=4`. Example: `LCA(3,4,5)=3`.

Hence total kawaiiness is `2 + 2 + 3 + 3 + 3 + 4 = 17`.

### ideas
1. 一个节点u能成为k个节点的lca, sz[u] >= k? 好像是成立的（选择其中一个节点是自己，再选择k-1个子节点即可）
2. 所以，就是看一个节点有多少贡献？
3. 