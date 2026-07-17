# E - Path Decomposition of a Tree

[Problem link](https://atcoder.jp/contests/abc397/tasks/abc397_e)

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 475 points

## Problem Statement

You are given a tree with `NK` vertices. The vertices are numbered
`1, 2, ..., NK`, and the `i`-th edge connects vertices `u_i` and `v_i`
bidirectionally.

Determine whether this tree can be decomposed into `N` paths, each of length
`K`. More precisely, determine whether there exists an `N × K` matrix `P`
satisfying:

- `P_{1,1}, ..., P_{1,K}, P_{2,1}, ..., P_{N,K}` is a permutation of
  `1, 2, ..., NK`.
- For each `i = 1..N` and `j = 1..K-1`, there is an edge connecting
  `P_{i,j}` and `P_{i,j+1}`.

## Constraints

- `1 <= N`
- `1 <= K`
- `NK <= 2 * 10^5`
- `1 <= u_i < v_i <= NK`
- The given graph is a tree
- All input values are integers

## Input

```
N K
u_1 v_1
u_2 v_2
...
u_{NK-1} v_{NK-1}
```

## Output

Print `Yes` if the tree can be decomposed into `N` paths each of length `K`,
otherwise print `No`.

## Samples

### Sample 1

Input:

```
3 2
1 2
2 3
3 4
2 5
5 6
```

Output:

```
Yes
```

Paths: `(1,2)`, `(3,4)`, `(5,6)`.

### Sample 2

Input:

```
3 2
1 2
2 3
3 4
2 5
3 6
```

Output:

```
No
```


### ideas
1. 考虑一个节点u, 它有两个子节点, v1, v2, path(v1) 表示v1子树中, 剩余的从v1开始的一个路径,
2. 如果 path(v1) + 1 + path(v2) = k => 那么可以切掉u子树(path(u) = 0)
3. 如果 path(v1) > 0, path(v2) > 0 => false (因为这两两条路径都需要u)

## Solution Summary

Root the tree at vertex `0` and process it with DFS from the leaves upward.
After decomposing as much of the subtree of `u` as possible, at most one path
may remain unfinished. This unfinished path must:

- contain `u`;
- have `u` as one endpoint;
- continue toward the parent of `u` to eventually reach `K` vertices.

The return value of `dfs(parent, u)` describes this state:

- `-1`: the subtree cannot be decomposed correctly;
- `0`: the whole subtree has been decomposed into paths of `K` vertices;
- `r`, where `1 <= r < K`: one unfinished path with `r` vertices remains,
  with `u` as its endpoint.

The positive return values from the children are collected in `todo`. Each
such value represents an unfinished child path that can only be connected
through `u`.

### DFS Transitions

#### No unfinished child path

If `todo` is empty, all child subtrees are already complete. Vertex `u` starts
a new unfinished path by itself, so DFS returns `1`.

```go
if len(todo) == 0 {
    return 1
}
```

#### One unfinished child path

Suppose the child path contains `x` vertices. Adding `u` extends it to
`x + 1` vertices:

```go
w := todo[0] + 1
return w % k
```

- If `x + 1 < K`, the path is still unfinished, so its new length is returned.
- If `x + 1 = K`, the path is complete and `0` is returned.

A child never returns `K`, so `x + 1` cannot exceed `K` in this case.

#### Two unfinished child paths

Suppose their lengths are `l` and `r`. Both paths need to use `u`, so the only
possible choice is to connect them through `u`:

```text
child path of length l -- u -- child path of length r
```

The resulting path contains `l + 1 + r` vertices. It must be a complete path
of exactly `K` vertices; otherwise this subtree is impossible.

The code checks

```go
if (l+1+r)%k != 0 {
    return -1
}
return 0
```

Because `1 <= l,r < K`, the sum is between `3` and `2K-1`. Therefore, if it is
divisible by `K`, it must equal `K`; it cannot equal `2K`. The two unfinished
paths and `u` then form one complete path.

This corrects the third rough idea above: two positive child paths do not
always mean failure. They are valid precisely when joining them through `u`
produces a path of `K` vertices.

#### More than two unfinished child paths

A simple path can use at most two edges incident to `u`. If three or more child
paths all need `u`, they cannot all be included without branching or reusing
`u`. Thus DFS returns `-1`.

### Why the Greedy Merge Is Forced

Every completed path contains exactly `K` vertices. Hence, after removing all
completed paths from a child subtree, the number of vertices in its unfinished
part is fixed modulo `K`. The parent cannot choose a different unfinished
length.

Furthermore, all unfinished child paths can connect to the rest of the tree
only through `u`. Consequently:

- zero unfinished paths must start a new path at `u`;
- one unfinished path must be extended with `u`;
- two unfinished paths must be joined and completed through `u`;
- three or more unfinished paths are impossible.

So the DFS does not discard any alternative valid decomposition.

### Root Condition

The root has no parent, so it cannot pass an unfinished path upward. The entire
tree is decomposable exactly when the root returns `0`:

```go
res := dfs(-1, 0)
if res == 0 {
    return "Yes"
}
return "No"
```

When `K = 1`, every vertex alone is already a valid path, so the code returns
`Yes` immediately.

### Correctness Proof

We prove by induction on subtree size that `dfs(parent, u)` returns the state
described above.

For a leaf, there are no unfinished child paths. DFS returns `1`, correctly
representing the one-vertex unfinished path containing only the leaf.

For an internal vertex, assume every child return value correctly describes
its subtree. A child returning `0` is already completely decomposed. Every
child returning a positive value has exactly one unfinished path whose only
possible connection outside that child subtree is through `u`.

The four transitions above consider every possible number of unfinished child
paths. With zero or one, `u` respectively starts or extends the only possible
unfinished path. With two, both must be joined through `u`, and the transition
accepts exactly when this creates a complete `K`-vertex path. With more than
two, no simple path can contain all required branches through `u`. Therefore
the return value is correct for the subtree of `u`.

By induction, the root return value is correct. Since the root cannot leave an
unfinished path, the algorithm returns `Yes` exactly when the whole tree can be
partitioned into `N` paths of `K` vertices.

### Complexity

Let `M = NK` be the number of vertices. Every vertex and edge is processed
once.

- Time: `O(M) = O(NK)`
- Space: `O(M) = O(NK)` for the adjacency list and DFS stack
