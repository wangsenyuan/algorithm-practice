# F. Subtree Minimum Query

## Problem Statement

You are given a rooted tree consisting of $n$ vertices. Each vertex has a number written on it; number $a_i$ is written on vertex $i$.

Let's denote $d(i, j)$ as the distance between vertices $i$ and $j$ in the tree (that is, the number of edges in the shortest path from $i$ to $j$). 

Also let's denote the **k-blocked subtree** of vertex $x$ as the set of vertices $y$ such that both these conditions are met:

- $x$ is an ancestor of $y$ (every vertex is an ancestor of itself);
- $d(x, y) \leq k$.

You are given $m$ queries to the tree. The $i$-th query is represented by two numbers $x_i$ and $k_i$, and the answer to this query is the minimum value of $a_j$ among such vertices $j$ such that $j$ belongs to $k_i$-blocked subtree of $x_i$.

Write a program that would process these queries quickly!

**Note:** The queries are given in a modified way.

## Input

The first line contains two integers $n$ and $r$ ($1 \leq r \leq n \leq 100000$) — the number of vertices in the tree and the index of the root, respectively.

The second line contains $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq 10^9$) — the numbers written on the vertices.

Then $n - 1$ lines follow, each containing two integers $x$ and $y$ ($1 \leq x, y \leq n$) and representing an edge between vertices $x$ and $y$. It is guaranteed that these edges form a tree.

Next line contains one integer $m$ ($1 \leq m \leq 10^6$) — the number of queries to process.

Then $m$ lines follow, $i$-th line containing two numbers $p_i$ and $q_i$, which can be used to restore $i$-th query ($1 \leq p_i, q_i \leq n$).

The $i$-th query can be restored as follows:

Let $last$ be the answer for previous query (or $0$ if $i = 1$). Then $x_i = ((p_i + last) \mod n) + 1$, and $k_i = (q_i + last) \mod n$.

## Output

Print $m$ integers. The $i$-th of them has to be equal to the answer to $i$-th query.

## Example

### Input
```
5 2
1 3 2 3 5
2 3
5 1
3 4
4 1
2
1 2
2 3
```

### Output
```
2
5
```

### ideas
1. 如果k <= sqrt(n), 貌似先计算出来是更好的
2. 如果k > sqrt(n), 它是颗树，所以v有超过距离为k的节点时，u也会有这样的节点
3. 当k比较大的时候，整颗树有可能会被包括进去
4. dp[u][i] 表示在子树u中，距离不超过 1 << i的距离的最小值？然后进行合并？
5. 当k = 1 << i的时候，应该是可以的
6. 但是其他情况呢？这时候就不知道去哪里了
7. 如果用heavy-light分解呢？
8. 在主分支上，是一个区间查询
9. 但是，在极端情况下，仍然是o(n)的
10. 所以，假设一个节点u，它的子节点数量比较少（ < sqrt(n))， 那么就每个节点查询一边
11. 如果超过sqrt(n) (平均每个子节点的sz < sqrt(n))
12. 但是也只是平均值，仍然有个别节点会超过这个
13. 所以，需要知道每个节点sqrt(n)距离处的节点