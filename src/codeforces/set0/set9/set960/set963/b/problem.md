# Problem Statement

You are given a tree (a graph with n vertices and n-1 edges in which it's possible to reach any vertex from any other vertex using only its edges).

A vertex can be destroyed if this vertex has even degree. If you destroy a vertex, all edges connected to it are also deleted.

Destroy all vertices in the given tree or determine that it is impossible.

## Input
- The first line contains integer n (1 ≤ n ≤ 2 × 10^5) — number of vertices in a tree.
- The second line contains n integers p₁, p₂, ..., pₙ (0 ≤ pᵢ ≤ n). If pᵢ ≠ 0 there is an edge between vertices i and pᵢ. It is guaranteed that the given graph is a tree.

## Output
- If it's possible to destroy all vertices, print "YES" (without quotes), otherwise print "NO" (without quotes).
- If it's possible to destroy all vertices, in the next n lines print the indices of the vertices in order you destroy them. If there are multiple correct answers, print any.

## ideas
1. 