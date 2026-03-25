You have a rooted tree consisting of `n` vertices. Each vertex of the tree has some color. We will assume that the tree vertices are numbered by integers from `1` to `n`. Then we represent the color of vertex `v` as `c_v`. The tree root is a vertex with number `1`.

In this problem you need to answer to `m` queries. Each query is described by two integers `v_j`, `k_j`. The answer to query `v_j, k_j` is the number of colors of vertices `x` such that the subtree of vertex `v_j` contains at least `k_j` vertices of color `x`.



## Input

The first line contains two integers `n` and `m` (`2 <= n <= 10^5`; `1 <= m <= 10^5`). The next line contains a sequence of integers `c1, c2, ..., c_n` (`1 <= c_i <= 10^5`). The next `n - 1` lines contain the edges of the tree. The `i`-th line contains the numbers `a_i, b_i` (`1 <= a_i, b_i <= n; a_i != b_i`) — the vertices connected by an edge of the tree.

Next `m` lines contain the queries. The `j`-th line contains two integers `v_j, k_j` (`1 <= v_j <= n; 1 <= k_j <= 10^5`).

## Output

Print `m` integers — the answers to the queries in the order the queries appear in the input.

## Examples

### Example 1

**Input**

```
8 5
1 2 2 3 3 2 3 3
1 2
1 5
2 3
2 4
5 6
5 7
5 8
1 2
1 3
1 4
2 3
5 3
```

**Output**

```
2
2
1
0
1
```

### Example 2

**Input**

```
4 1
1 2 3 4
1 2
2 3
3 4
1 1
```

**Output**

```
4
```

## Note

A subtree of vertex `v` in a rooted tree with root `r` is a set of vertices `{u : dist(r, v) + dist(v, u) = dist(r, u)}`. Where `dist(x, y)` is the length (in edges) of the shortest path between vertices `x` and `y`.


