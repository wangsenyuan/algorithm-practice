# G - Count Simple Paths 2 (ABC419)

**Contest:** [ABC419](https://atcoder.jp/contests/abc419) — AtCoder Beginner Contest 419  
**Task:** [https://atcoder.jp/contests/abc419/tasks/abc419_g](https://atcoder.jp/contests/abc419/tasks/abc419_g)

**Time limit:** 4 sec / **Memory limit:** 1024 MiB  
**Score:** 600 points

## Problem Statement

You are given a simple connected undirected graph with `N` vertices numbered
`1` to `N` and `M` edges. The `i`-th edge connects vertices `u_i` and `v_i`.

For each `k = 1, 2, ..., N - 1`, find the number of simple paths from vertex `1`
to vertex `N` that contain exactly `k` edges.

## Constraints

- `2 <= N <= 2 * 10^5`
- `N - 1 <= M <= N + 20`
- `1 <= u_i < v_i <= N`
- The given graph is a simple connected undirected graph.
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```text
N M
u_1 v_1
u_2 v_2
...
u_M v_M
```

## Output

Output the answers in the following format:

```text
ans_1 ans_2 ... ans_{N-1}
```

Here, `ans_i` is the number of simple paths from vertex `1` to vertex `N` that
contain exactly `i` edges.

## Sample Input 1

```text
5 6
1 2
1 3
2 4
3 4
3 5
4 5
```

## Sample Output 1

```text
0 1 2 1
```

For each `k = 1, 2, 3, 4`, the simple paths from vertex `1` to vertex `5` that
contain exactly `k` edges are as follows.

- `k = 1`: none
- `k = 2`: `1 -> 3 -> 5`
- `k = 3`: `1 -> 2 -> 4 -> 5` and `1 -> 3 -> 4 -> 5`
- `k = 4`: `1 -> 2 -> 4 -> 3 -> 5`

## Sample Input 2

```text
11 15
1 2
1 3
2 3
3 4
3 5
4 5
5 6
5 7
6 7
7 8
7 9
8 9
9 10
9 11
10 11
```

## Sample Output 2

```text
0 0 0 0 1 5 10 10 5 1
```

## Sample Input 3

```text
7 18
6 7
4 5
1 7
2 7
1 4
2 5
4 6
2 3
5 6
5 7
1 5
2 4
2 6
1 2
1 3
3 4
1 6
3 5
```

## Sample Output 3

```text
1 3 11 29 50 42
```


### ideas
1. 如果M = n - 1, 那么它就是一棵树，这种情况下，只能有一个答案（从1到n)
2. 在一个环上，如果给定了两个点，那么它们之间的距离，有两种可能性
3. 但是如果是在一个比较复杂的联通块上，入点和出点之间的距离（似乎就比较复杂了）
4. 这个貌似得用 mask进行处理（不能访问已经访问过的点）
5. 那些一条线，要被压缩成一个点（因为只有一种选择）
6. 仅仅压缩还不行，还需要知道长度
7. 假设是 1 <-> 2 <-> 3 <-> 1,  3 <-> 4 <-> 5 <-> 3
8. 感觉不大对。
9. 先构造出树， 然后对于额外的边，进行处理
10. 假设 dp[mask][?] 是使用了额外的mask表示的边的数量后，能够从1...n的距离
11. 新的边加入后，它连接了(u, v), 如果知道新产生的距离d, 加进去就可以了
12. 但是这里有个问题(1...u) 可能因为mask的加入，有可能会有多条距离。
13. 而这个状态，却无法保存下来。
14. 