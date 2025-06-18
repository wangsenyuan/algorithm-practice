# Tree Leaf Removal Problem

## Problem Description
You are given a tree (connected graph without cycles) consisting of $n$ vertices. The tree is unrooted — it is just a connected undirected graph without cycles.

## Rules
In one move, you can:
- Choose exactly $k$ leaves (a leaf is a vertex connected to only one other vertex)
- All chosen leaves must be connected to the same vertex
- Remove these leaves along with their incident edges

More formally, you choose leaves $u_1, u_2, ..., u_k$ that have edges $(u_1,v), (u_2,v), ..., (u_k,v)$ and remove these leaves and edges.

Your task is to find the maximum number of moves you can perform if you remove leaves optimally.

## Input Format
The first line contains one integer $t$ $(1 \leq t \leq 2\cdot10^4)$ — the number of test cases.

For each test case:
1. First line contains two integers $n$ and $k$ $(2 \leq n \leq 2\cdot10^5; 1 \leq k < n)$
   - $n$ = number of vertices in the tree
   - $k$ = number of leaves to remove in one move
2. The next $n-1$ lines describe edges, each containing two integers $x_i$ and $y_i$ $(1 \leq x_i,y_i \leq n)$
   - Each line represents an edge connecting vertices $x_i$ and $y_i$
   - The given edges are guaranteed to form a tree

### Constraints
- Sum of $n$ across all test cases does not exceed $2\cdot10^5$ $(\sum n \leq 2\cdot10^5)$

## Output Format
For each test case, print the answer — the maximum number of moves you can perform if you remove leaves optimally.

## ideas
1. 对于节点v，如果它刚好有k个叶子邻居，那么删除它的k个叶子后，叶子总数 -k + 1 (v变成了新的叶子)
2. 如果叶子邻居超过k，那么叶子总数 -k (没有产生新的叶子)
3. 对于一v来说，它的操作步骤 = leaf_count_of(v) / k, 剩余 leaf_count_of(v) % k
4. 如果剩余的为0，那么v自己变成了叶子，如果剩余的 < k - 1, 那么v这里就stuck住了，不用继续处理它
5. 如果 = k - 1, 且v有parent，那么如果能处理完parent，那么有可能这里会贡献1
6. 那是不是就是尽量的删除呢？就根据leaf_count_of(v) ?
7.  