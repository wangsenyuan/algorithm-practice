# Vertex Cactus Paths

## Problem Description

A connected undirected graph is called a **vertex cactus** if each vertex of this graph belongs to at most one simple cycle.

### Definitions

1. **Simple Cycle**: A sequence of distinct vertices v1, v2, ..., vt (t > 2), such that:
   - For any i (1 ≤ i < t), there exists an edge between vertices vi and vi + 1
   - There exists an edge between vertices v1 and vt

2. **Simple Path**: A sequence of not necessarily distinct vertices v1, v2, ..., vt (t > 0), such that:
   - For any i (1 ≤ i < t), there exists an edge between vertices vi and vi + 1
   - Each edge occurs no more than once
   - The path starts at vertex v1 and ends at vertex vt

## Task

Given a vertex cactus graph and a list of vertex pairs, count the number of distinct simple paths between each pair. Two simple paths are considered distinct if their sets of edges are different.

## Input Format

- First line: Two space-separated integers `n, m` (2 ≤ n ≤ 105; 1 ≤ m ≤ 105)
  - n: number of vertices
  - m: number of edges

- Next m lines: Two space-separated integers `ai, bi` (1 ≤ ai, bi ≤ n)
  - Describes an edge connecting vertices ai and bi

- Next line: Single integer `k` (1 ≤ k ≤ 105)
  - Number of interesting vertex pairs

- Next k lines: Two space-separated integers `xi, yi` (1 ≤ xi, yi ≤ n; xi ≠ yi)
  - Pair of vertices to find paths between

## Constraints

- The graph is guaranteed to be a vertex cactus
- No loops or multiple edges exist
- Vertices are numbered from 1 to n

## Output Format

Print k lines, where the i-th line contains a single integer:
- The number of distinct simple paths from xi to yi
- Output each number modulo 1000000007 (10<sup>9</sup> + 7)

