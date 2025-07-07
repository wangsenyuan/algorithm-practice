# Problem C: p-Interesting Graph

## Problem Description

Let's call an undirected graph of n vertices **p-interesting**, if the following conditions are fulfilled:

1. The graph contains exactly **2n + p** edges
2. The graph doesn't contain self-loops and multiple edges
3. For any integer k (1 ≤ k ≤ n), any subgraph consisting of k vertices contains at most **2k + p** edges

**Note:** A subgraph of a graph is some set of the graph vertices and some set of the graph edges. At that, the set of edges must meet the condition: both ends of each edge from the set must belong to the chosen set of vertices.

Your task is to find a p-interesting graph consisting of n vertices.

## Input

- The first line contains a single integer **t** (1 ≤ t ≤ 5) — the number of tests in the input
- Next **t** lines each contains two space-separated integers: **n, p** (5 ≤ n ≤ 24; p ≥ 0) — the number of vertices in the graph and the interest value for the appropriate test

It is guaranteed that the required graph exists.

## Output

For each of the **t** tests print **2n + p** lines containing the description of the edges of a p-interesting graph:

- The i-th line must contain two space-separated integers **ai, bi** (1 ≤ ai, bi ≤ n; ai ≠ bi) — two vertices, connected by an edge in the resulting graph
- Consider the graph vertices numbered with integers from 1 to n

Print the answers to the tests in the order the tests occur in the input. If there are multiple solutions, you can print any of them.

### ideas
1. 2 * n + p <= n * (n - 1) / 2
2. 那么是不是一定存在一个m <= n, 2 * m + p = m * (m - 1) / 2 呢？
3. 也就是在一个完全子图中，正好有额外p条边
4. 4 * m + 2 * p = m * (m - 1)
5. 似乎也不一定。有可能差个1
6. 如果差个1，最后在两个叶子节点间，连起来就可以了？