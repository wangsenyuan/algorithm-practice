### Problem

Tree is a connected acyclic graph. Suppose you are given a tree consisting of \(n\) vertices. The vertex of this tree is called **centroid** if the size of each connected component that appears if this vertex is removed from the tree doesn't exceed .

You are given a tree of size \(n\) and can perform no more than one **edge replacement**. Edge replacement is the operation of removing one edge from the tree (without deleting incident vertices) and inserting one new edge (without adding new vertices) in such a way that the graph remains a tree. For each vertex you have to determine if it's possible to make it centroid by performing no more than one edge replacement.

### Input

The first line of the input contains an integer \(n\) \((2 \le n \le 400\,000)\) — the number of vertices in the tree. Each of the next \(n - 1\) lines contains a pair of vertex indices \(u_i\) and \(v_i\) \((1 \le u_i, v_i \le n)\) — endpoints of the corresponding edge.

### Output

Print \(n\) integers. The \(i\)-th of them should be equal to `1` if the \(i\)-th vertex can be made centroid by replacing no more than one edge, and should be equal to `0` otherwise.

### Examples

#### Example 1

**Input**

```text
3
1 2
2 3
```

**Output**

```text
1 1 1 
```

#### Example 2

**Input**

```text
5
1 2
1 3
1 4
1 5
```

**Output**

```text
1 0 0 0 0 
```

### Note

In the first sample each vertex can be made a centroid. For example, in order to turn vertex 1 to centroid one has to replace the edge \((2, 3)\) with the edge \((1, 3)\).


### ideas
1. 看起来比较费时
2. 考虑一个节点，它原来不是centroid，那么必然存在一个子树的size > n / 2
3. 所以，假设要去掉一条边，那么必然是从这个子树中，找到一条边e，把它去掉，然后接到当前节点上
4. 假设是u，子树是v，断掉的是w, 那么 sz[v] > n / 2, sz[w] <= n / 2, 且 sz[v] - sz[w] < n / 2
5. 那么w似乎可以与u无关，就是在v中，最大的子树w，且 sz[w] <= n / 2.
6. 要找到v的一个最大的子树，是比较容易的。但是麻烦的地方在于，这个子树，有可能是在u这边的
7. 所以，需要两个 w1, w2