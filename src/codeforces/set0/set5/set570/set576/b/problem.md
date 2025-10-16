A tree of size n is an undirected connected graph consisting of n vertices without cycles.

Consider some tree with n vertices. We call a tree invariant relative to permutation p = p1p2... pn, if for any two vertices of the tree u and v the condition holds: "vertices u and v are connected by an edge if and only if vertices pu and pv are connected by an edge".

You are given permutation p of size n. Find some tree size n, invariant relative to the given permutation.

## Input

The first line contains number n (1 ≤ n ≤ 10^5) — the size of the permutation (also equal to the size of the sought tree).

The second line contains permutation pi (1 ≤ pi ≤ n).

## Output

If the sought tree does not exist, print "NO" (without the quotes).

Otherwise, print "YES", and then print n - 1 lines, each of which contains two integers — the numbers of vertices connected by an edge of the tree you found. The vertices are numbered from 1, the order of the edges and the order of the vertices within the edges does not matter.

If there are multiple solutions, output any of them.

## Examples

### Input
```
4
4 3 2 1
```

### Output
```
YES
4 1
4 2
1 3
```

### Input
```
3
3 1 2
```

### Output
```
NO
```

## Note

In the first sample test a permutation transforms edge (4, 1) into edge (1, 4), edge (4, 2) into edge (1, 3) and edge (1, 3) into edge (4, 2). These edges all appear in the resulting tree.

It can be shown that in the second sample test no tree satisfies the given condition.



### ideas
1. 有点费脑子～
2. 如果p = 1,2,3,...n
3. 是不是一条直线就满足条件了？
4. p = 1,2,3...n的特点，就是它们每个点都是独立的
5. 把p分成了很多组,每组是一个cycle
6. (u, v) 之间有一条边, (p[u], p[v])之间也必须有一条边
7. 如果这个分组只有两个(u, v)