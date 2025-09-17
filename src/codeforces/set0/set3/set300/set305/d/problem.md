# Problem D

## Problem Statement

Olya has got a directed non-weighted graph, consisting of n vertexes and m edges. We will consider that the graph vertexes are indexed from 1 to n in some manner. Then for any graph edge that goes from vertex v to vertex u the following inequation holds: v < u.

Now Olya wonders, how many ways there are to add an arbitrary (possibly zero) number of edges to the graph so as the following conditions were met:

1. You can reach vertexes number i + 1, i + 2, ..., n from any vertex number i (i < n).
2. For any graph edge going from vertex v to vertex u the following inequation fulfills: v < u.
3. There is at most one edge between any two vertexes.
4. The shortest distance between the pair of vertexes i, j (i < j), for which j - i ≤ k holds, equals j - i edges.
5. The shortest distance between the pair of vertexes i, j (i < j), for which j - i > k holds, equals either j - i or j - i - k edges.

We will consider two ways distinct, if there is the pair of vertexes i, j (i < j), such that first resulting graph has an edge from i to j and the second one doesn't have it.

Help Olya. As the required number of ways can be rather large, print it modulo 1000000007 (10^9 + 7).

## Input

The first line contains three space-separated integers n, m, k (2 ≤ n ≤ 10^6, 0 ≤ m ≤ 10^5, 1 ≤ k ≤ 10^6).

The next m lines contain the description of the edges of the initial graph. The i-th line contains a pair of space-separated integers ui, vi (1 ≤ ui < vi ≤ n) — the numbers of vertexes that have a directed edge from ui to vi between them.

It is guaranteed that any pair of vertexes ui, vi has at most one edge between them. It also is guaranteed that the graph edges are given in the order of non-decreasing ui. If there are multiple edges going from vertex ui, then it is guaranteed that these edges are given in the order of increasing vi.

## Output

Print a single integer — the answer to the problem modulo 1000000007 (10^9 + 7).

## Examples

### Example 1
**Input:**
```
7 8 2
1 2
2 3
3 4
3 6
4 5
4 7
5 6
6 7
```

**Output:**
```
2
```

### Example 2
**Input:**
```
7 0 2
```

**Output:**
```
12
```

### Example 3
**Input:**
```
7 2 1
1 3
3 5
```

**Output:**
```
0
```

## Note

In the first sample there are two ways: the first way is not to add anything, the second way is to add a single edge from vertex 2 to vertex 5.


## ideas
1. 阅读理解～
2. 条件1，基本就是说，在所有(i, i + 1)之间要有一条边
3. 主要是条件4和5
4. 在两个节点(l < r) 如果 r - l <= k, 那么它们的距离 = r - l 
5. 如果 r - l > k, 那么距离要么 = r - l, 要么 = r - l - k
6. 比如 1, 2, 3, 4, 5, 6, 7, 如果 k = 2, 那么 (1, 4) 之间的距离 = 3 or 1?
7. 如果在(2, 5)中间加一条边, 似乎是ok的， 那么 (1, 5) = 5 - 1 - 2 = 2
8. 1和4中间也可以加一条边，
9. 可以在(1, 5)中间加一条边吗？ 不可以
10. (1, 7)中间可以加一条边吗？似乎也不行
11. 所以条件4和提交5就是在说，在i和i+k+1中间可以有一条边
12. 但是，如果在(1, 4)中间有一条边，那么在（4, 7)中间就不能有，否则(1, 7)的距离 = 2， 就不符合条件5了
13. 所以(1, 4, 7)只能2选择1
14. 差不多就是这样？
15. 