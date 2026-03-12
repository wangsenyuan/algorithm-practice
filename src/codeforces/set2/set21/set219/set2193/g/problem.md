# Problem

**题意简述**：交互题。给定一棵 `n` 个点的无环连通无向图。隐藏两个点 `x`、`y`（可相同）。每次可询问两个点 `a, b`，裁判返回 1 表示路径 (x–y) 与路径 (a–b) 至少有一个公共点，否则返回 0。目标是用不超过 `⌊n/2⌋ + 1` 次询问，找到 x–y 路径上的任意一个点。注意 interactor 是自适应的。



## Input

- The first line contains one integer `t` (`1 ≤ t ≤ 10^4`) — the number of test cases.
- The first line of each test case contains one integer `n` (`2 ≤ n ≤ 2 * 10^5`) — the number of vertices.
- The next `n - 1` lines each contain two integers `v`, `u` (`1 ≤ v, u ≤ n`) — an edge between `v` and `u`.

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Interaction

- You may use at most `⌊n/2⌋ + 1` queries per test case. Each query: output `? a b`.
- After each query, read one integer `0` or `1` — the response.
- When you have found a vertex on the path, output `! v` (`1 ≤ v ≤ n`).
- If you make more than `⌊n/2⌋ + 1` queries, the response will be `-1`; your program should terminate immediately to avoid an incorrect verdict.


### ideas
1. 如果 x != y，倒是比较有个办法处理。但是如果 x = y, 要怎么处理？
2. 如果x = y, 那么就必须找到这个点
3. 假设现在处理子树u
4. 如果u是一个叶子节点，那么不处理，退回到上级节点
5. 如果目前这个子树有children, 那么就两个一组，进行询问
6. 有两种情况发生，所有的都问过以后，都没有发现（那么可以把u这个子树直接删掉）
7. 如果最后剩下一个叶子节点，（children的个数为奇数），那么返回这个叶子节点（而不是自己）
8. 如果在某个(a, b)询问的过程中，出现了1（那么必然是(a, u, b)中的一个
9. 如果u确定不是（前面已经包含u，且答案为0）那么再多问一次a, 
10. 如果u不确定, 那么需要再问两次？似乎多了一次；
11. 比如正好是 (1 <-> 2 <-> 3) 这种结构，那么只能问2次（这个能处理的）
12. 似乎到了最后的时候，还是得特殊处理一下