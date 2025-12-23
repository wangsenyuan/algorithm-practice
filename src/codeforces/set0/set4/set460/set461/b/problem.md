# Problem Description

Appleman has a tree with n vertices. Some of the vertices (at least one) are colored black and other vertices are colored white.

Consider a set consisting of k (0 ≤ k < n) edges of Appleman's tree. If Appleman deletes these edges from the tree, then it will split into (k + 1) parts. Note, that each part will be a tree with colored vertices.

Now Appleman wonders, what is the number of sets splitting the tree in such a way that each resulting part will have exactly one black vertex? Find this number modulo 1000000007 (10^9 + 7).

## Input

The first line contains an integer n (2 ≤ n ≤ 10^5) — the number of tree vertices.

The second line contains the description of the tree: n - 1 integers p0, p1, ..., pn - 2 (0 ≤ pi ≤ i). Where pi means that there is an edge connecting vertex (i + 1) of the tree and vertex pi. Consider tree vertices are numbered from 0 to n - 1.

The third line contains the description of the colors of the vertices: n integers x0, x1, ..., xn - 1 (xi is either 0 or 1). If xi is equal to 1, vertex i is colored black. Otherwise, vertex i is colored white.

## Output

Output a single integer — the number of ways to split the tree modulo 1000000007 (10^9 + 7).

## Examples

### Example 1

**Input:**

```text
3
0 0
0 1 1
```

**Output:**

```text
2
```

### Example 2

**Input:**

```text
6
0 1 1 0 4
1 1 0 0 1 0
```

**Output:**

```text
1
```

### Example 3

**Input:**

```text
10
0 1 2 1 4 4 4 0 8
0 0 0 1 0 1 1 0 0 1
```

**Output:**

```text
27
```


### ideas
1. 如果是一条直线，那么两个黑点之间的线，必须有一条被剪短，*k（k是黑点之间线的数量）
2. 考虑树的情况； 如果点u是黑色，且子树v中也有个黑色节点，那么(u...v)中间的边，恰好要有一条边被删除
3. 考虑星型结构, a, b, c都连接到o, (a, b, c)都是黑色的， 但是o是白色的
4. 一共3条边，只要删除其中的2条边，就可以完成(C(3, 2))
5. 但是如果每条边的长度不一致呢？包含的边的数量不一致
6. 如果选择o和a连接, 那么 = kb * kc, 如果选择和b连接， ka * kc, 如果选择和c连接， ka * kb
7. 肯定要dp，但是状态似乎有点乱; 
8. 