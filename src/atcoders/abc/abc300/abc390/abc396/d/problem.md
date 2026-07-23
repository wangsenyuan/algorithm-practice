# D - Minimum XOR Path

[Problem link](https://atcoder.jp/contests/abc396/tasks/abc396_d)

**Contest:** [AtCoder Beginner Contest 396](https://atcoder.jp/contests/abc396)

time limit: 2 sec

memory limit: 1024 MiB

score: 400 points

## Problem

You are given a simple connected undirected graph with `N` vertices
(`1..N`) and `M` edges. Edge `i` connects `u_i` and `v_i` and has label `w_i`.

Among all simple paths from vertex `1` to vertex `N` (paths that do not visit
the same vertex more than once), find the minimum XOR of the edge labels on
the path.

## Constraints

- `2 <= N <= 10`
- `N - 1 <= M <= N*(N-1)/2`
- `1 <= u_i < v_i <= N`
- `0 <= w_i < 2^60`
- The graph is simple, connected, and undirected
- All input values are integers

## Input

```text
N M
u_1 v_1 w_1
...
u_M v_M w_M
```

## Output

Print the answer.

## Samples

### Sample 1

```text
Input
4 4
1 2 3
2 4 5
1 3 4
3 4 7

Output
3
```

Paths `1 -> 2 -> 4` (XOR `6`) and `1 -> 3 -> 4` (XOR `3`); answer is `3`.

### Sample 2

```text
Input
4 3
1 2 1
2 3 2
3 4 4

Output
7
```

### Sample 3

```text
Input
7 10
1 2 726259430069220777
1 4 988687862609183408
1 5 298079271598409137
1 6 920499328385871537
1 7 763940148194103497
2 4 382710956291350101
3 4 770341659133285654
3 5 422036395078103425
3 6 472678770470637382
5 7 938201660808593198

Output
186751192333709144
```

## Solution Summary

Because `N <= 10`, we can enumerate every simple path from vertex `1` to
vertex `N` with depth-first search.

Build an adjacency list that stores edge indices. During DFS, maintain:

- `u`: the current vertex;
- `mask`: the set of vertices already used by the current path;
- `sum`: the XOR of all edge labels on the current path.

For an edge with endpoints `a` and `b`, when the current vertex is `u`, the
other endpoint can be obtained by

```text
a XOR b XOR u
```

because one of `a` and `b` equals `u` and the two equal values cancel. The
implementation performs this operation with the original one-based endpoint
numbers and converts the result back to a zero-based vertex.

For every edge incident to `u`, continue the DFS only if its other endpoint
has not been visited. The next state adds that vertex to `mask` and replaces
`sum` by `sum XOR edge_weight`.

When the DFS reaches vertex `N`, the current path is complete. Update the
answer with the minimum XOR value found, and stop extending that path.

### Correctness Proof

We prove that the algorithm returns the minimum XOR value of a simple path
from vertex `1` to vertex `N`.

At every DFS call, `mask` contains exactly the vertices on the current path,
and `sum` is exactly the XOR of the labels of its edges. This is true for the
initial state, whose path contains only vertex `1` and whose XOR is `0`. Each
transition chooses an incident edge, adds its unvisited opposite endpoint, and
XORs that edge's label into `sum`, so the invariant remains true.

The visited mask prevents the DFS from entering a vertex already on the
current path. Therefore, every path reaching vertex `N` is simple.

Conversely, consider any simple path from vertex `1` to vertex `N`. Starting
from vertex `1`, each next vertex on this path is unvisited and connected by
an incident edge, so the DFS includes the corresponding transition. Repeating
this argument shows that the DFS eventually enumerates the entire path.
Therefore, every valid simple path is considered.

When a path reaches vertex `N`, the invariant shows that `sum` equals that
path's XOR value. Since the algorithm takes the minimum over all enumerated
simple paths, and those are exactly all valid paths, the returned answer is
correct.

### Complexity

The number of simple paths can be factorial in `N`, so the worst-case time
complexity is `O(N!)`. This is feasible because `N <= 10`.

The graph and adjacency list use `O(N + M)` space, and the recursion stack and
visited mask use `O(N)` additional space.
