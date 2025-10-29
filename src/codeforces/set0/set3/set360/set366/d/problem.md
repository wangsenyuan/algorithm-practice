# Problem Description

Dima and Inna love spending time together. The problem is, Seryozha isn't too enthusiastic to leave his room for some reason. But Dima and Inna love each other so much that they decided to get criminal...

Dima constructed a trap graph. He shouted: "Hey Seryozha, have a look at my cool graph!" to get his roommate interested and kicked him into the first node.

A trap graph is an undirected graph consisting of n nodes and m edges. For edge number k, Dima denoted a range of integers from lk to rk (lk ≤ rk). In order to get out of the trap graph, Seryozha initially (before starting his movements) should pick some integer (let's call it x), then Seryozha must go some way from the starting node with number 1 to the final node with number n. At that, Seryozha can go along edge k only if lk ≤ x ≤ rk.

Seryozha is a mathematician. He defined the loyalty of some path from the 1-st node to the n-th one as the number of integers x, such that if he initially chooses one of them, he passes the whole path. Help Seryozha find the path of maximum loyalty and return to his room as quickly as possible!

## Input

The first line of the input contains two integers n and m (2 ≤ n ≤ 10³, 0 ≤ m ≤ 3·10³). Then follow m lines describing the edges. Each line contains four integers ak, bk, lk and rk (1 ≤ ak, bk ≤ n, 1 ≤ lk ≤ rk ≤ 10⁶). The numbers mean that in the trap graph the k-th edge connects nodes ak and bk, this edge corresponds to the range of integers from lk to rk.

Note that the given graph can have loops and multiple edges.

## Output

In a single line of the output print an integer — the maximum loyalty among all paths from the first node to the n-th one. If such paths do not exist or the maximum loyalty equals 0, print in a single line "Nice work, Dima!" without the quotes.

## Examples

### Example 1

**Input:**

```text
4 4
1 2 1 10
2 4 3 5
1 3 1 5
2 4 2 7
```

**Output:**

```text
6
```

### Example 2

**Input:**

```text
5 6
1 2 1 10
2 5 11 20
1 4 2 5
1 3 10 11
3 4 12 10000
4 5 6 6
```

**Output:**

```text
Nice work, Dima!
```

## Note

**Explanation of the first example:**

Overall, we have 2 ways to get from node 1 to node 4: first you must go along the edge 1-2 with range [1-10], then along one of the two edges 2-4.

One of them contains range [3-5], that is, we can pass through with numbers 3, 4, 5. So the loyalty of such path is 3.

If we go along edge 2-4 with range [2-7], then we can pass through with numbers 2, 3, 4, 5, 6, 7. The loyalty is 6. That is the answer.

The edge 1-2 have no influence on the answer because its range includes both ranges of the following edges.

### ideas

1. 这个题目理解不了啊。是描述上有问题吗？
2. 指的是x的数量，而不是x的大小
3. 给定一个x，看能否连通 1 ～ n, 如果能够 +1
4. 但是肯定TLE。 还需要再想想
5. x不一定是连续的
6. 比如 1 - 2 的范围是 [1, 10], [1 - 3] 的范围 是 [12 - 20]
7. 然后 2 - 4, 的范围是 [6 - 7], [3 - 4] 的范围是 [12 - 14]
8. 那么 x = 6, 7, 12, 13, 14 就可以
9. 但是， 一共m条边， 处理后，最多m段，可以按照每段去检查
10. 那么就可以了

**问题解释 (Problem Explanation):**

这是一个图论问题，关键在于理解"loyalty"的概念。

**核心概念:**

- 在开始之前，Seryozha 必须选择一个整数 x（只能选一次，选好后不能再改）
- 每条边 k 有一个范围 [lk, rk]
- 只有当 lk ≤ x ≤ rk 时，才能通过边 k
- 找到从节点 1 到节点 n 的一条路径
- **loyalty = 有多少个不同的 x 值可以让这条路径完全走通**

**具体例子 (Example 1 详解):**

图结构:

```text
1 --[1-10]--> 2 --[3-5]--> 4
1 --[1-10]--> 2 --[2-7]--> 4  (另一条边)
```

从节点 1 到节点 4 的两条路径:

- **路径1:** 1→2 (范围[1-10]) → 2→4 (范围[3-5])
  - 要完成这条路径，x 必须同时在 [1-10] 和 [3-5] 中
  - 交集: [3-5]，包含 3, 4, 5 三个整数
  - loyalty = 3

- **路径2:** 1→2 (范围[1-10]) → 2→4 (范围[2-7])
  - x 必须同时在 [1-10] 和 [2-7] 中
  - 交集: [2-7]，包含 2, 3, 4, 5, 6, 7 六个整数
  - loyalty = 6 ✓ (答案)

**算法思路:**

1. 枚举所有从 1 到 n 的路径（或使用 BFS/DFS）
2. 对于每条路径，计算所有边的范围的交集
3. 交集的长度（整数个数）就是这条路径的 loyalty
4. 返回最大 loyalty

**关键点:**

- loyalty 是路径上所有边范围的交集的大小
- 如果交集为空，loyalty = 0
- 如果没有路径存在，输出 "Nice work, Dima!"
