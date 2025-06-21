# D. Almost Acyclic Graph

You are given a directed graph consisting of `n` vertices and `m` edges. Each edge is directed, so it can be traversed in only one direction. You are allowed to remove at most one edge from it.

Can you make this graph acyclic by removing at most one edge? A directed graph is called **acyclic** if and only if it doesn't contain any cycle (a non-empty path that starts and ends in the same vertex).

## Input

The first line contains two integers, `n` and `m` (2 ≤ `n` ≤ 500, 1 ≤ `m` ≤ min(`n`*(`n` - 1), 100000)) — the number of vertices and the number of edges, respectively.

The next `m` lines describe the edges. Each line contains two integers, `u` and `v`, denoting a directed edge from vertex `u` to vertex `v` (1 ≤ `u`, `v` ≤ `n`, `u` ≠ `v`).

Each ordered pair (`u`, `v`) is listed at most once.

## Output

If it is possible to make this graph acyclic by removing at most one edge, print `YES`. Otherwise, print `NO`.

## ideas
1. 如果不存在强连通 -> ok
2. 如果存在多个强连通图 -> !ok
3. 只有一个的时候，它是一个环，才可行， 否则肯定不行。
4. 