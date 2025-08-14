# Problem E: Darth Vader's Tree Counting

## Problem Description

When Darth Vader gets bored, he sits down on the sofa, closes his eyes and thinks of an infinite rooted tree where each node has exactly $n$ sons, at that for each node, the distance between it and its $i$-th left child equals to $d_i$. The Sith Lord loves counting the number of nodes in the tree that are at a distance at most $x$ from the root. The distance is the sum of the lengths of edges on the path between nodes.

But he has got used to this activity and even grew bored of it. 'Why does he do that, then?' — you may ask. It's just that he feels superior knowing that only he can solve this problem.

Do you want to challenge Darth Vader himself? Count the required number of nodes. As the answer can be rather large, find it modulo $10^9 + 7$.

## Input

The first line contains two space-separated integers $n$ and $x$ ($1 \leq n \leq 10^5$, $0 \leq x \leq 10^9$) — the number of children of each node and the distance from the root within the range of which you need to count the nodes.

The next line contains $n$ space-separated integers $d_i$ ($1 \leq d_i \leq 100$) — the length of the edge that connects each node with its $i$-th child.

## Output

Print a single number — the number of vertices in the tree at distance from the root equal to at most $x$.

## Examples

### Example 1

**Input:**
```
3 3
1 2 3
```

**Output:**
```
8
```


### ideas
1. 如果root节点中能够在距离s中，访问到dp[s]个节点
2. 那么 dp[s+d[1]] += dp[s]
3.     dp[s+d[2]] += dp[s]
4.     ...
5.     dp[s+d[i]] += dp[s]
6. dp[0] = 1
7. 这里的问题，貌似是s很大
8. 假设在某个阶段得到 dp[?] 
9. 那么新的结果貌似是在它的上面 mul d[?] (貌似是加法)
10. fp[1]表示距离为1的节点数
11. fp[2]表示距离为2的节点数
12. fp[100] 是最大的数
13. fp[i] * dp[j] => dp[i+j]
14. 那么 pow(fp, s) 次？
15. 如果加入 fp[0] = 1, 但是怎么限定最后的距离 x <= s 呢？
16. 