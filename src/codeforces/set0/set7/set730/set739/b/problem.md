# Problem Description

Alyona has a tree with n vertices. The root of the tree is the vertex 1. In each vertex Alyona wrote a positive integer, in the vertex i she wrote ai. Moreover, the girl wrote a positive integer to every edge of the tree (possibly, different integers on different edges).

Let's define dist(v, u) as the sum of the integers written on the edges of the simple path from v to u.

The vertex v controls the vertex u (v ≠ u) if and only if u is in the subtree of v and dist(v, u) ≤ au.

Alyona wants to settle in some vertex. In order to do this, she wants to know for each vertex v what is the number of vertices u such that v controls u.

## Input

The first line contains single integer n (1 ≤ n ≤ 2·10⁵).

The second line contains n integers a₁, a₂, ..., aₙ (1 ≤ ai ≤ 10⁹) — the integers written in the vertices.

The next (n - 1) lines contain two integers each. The i-th of these lines contains integers pi and wi (1 ≤ pi ≤ n, 1 ≤ wi ≤ 10⁹) — the parent of the (i + 1)-th vertex in the tree and the number written on the edge between pi and (i + 1).

It is guaranteed that the given graph is a tree.

## Output

Print n integers — the i-th of these numbers should be equal to the number of vertices that the i-th vertex controls.
