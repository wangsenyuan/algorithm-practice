# G - Graph Problem 2026 (ABC461)

**Contest:** [ABC461](https://atcoder.jp/contests/abc461) — AtCoder Beginner Contest 461  
**Task:** [https://atcoder.jp/contests/abc461/tasks/abc461_g](https://atcoder.jp/contests/abc461/tasks/abc461_g)

**Time limit:** 4 sec / **Memory limit:** 1024 MiB  
**Score:** 625 points

## Problem Statement

There is a simple undirected graph with `N` vertices and `M` edges.  
The vertices are numbered `1` through `N` and the edges are numbered `1` through `M`; edge `i`
connects vertices `u_i` and `v_i`.

Assign a non-negative integer weight `W_j` not greater than `2026` to each vertex `j` so that
the following condition is satisfied:

- `W_{u_i} + W_{v_i} <= 2026` for each edge `i`.

Find the maximum possible value of the sum of all vertex weights (that is, `sum_{j=1}^{N} W_j`).

## Constraints

- `1 <= N <= 5 × 10^4`
- `0 <= M <= 5 × 10^4`
- `1 <= u_i < v_i <= N`
- `(u_1, v_1), (u_2, v_2), ..., (u_M, v_M)` are pairwise distinct.
- All input values are integers.

## Input

```text
N M
u_1 v_1
u_2 v_2
⋮
u_M v_M
```

## Output

Print the answer on a single line.

## Sample Input 1

```text
3 2
1 2
2 3
```

## Sample Output 1

```text
4052
```

By assigning weights `2026, 0, 2026` to vertices `1, 2, 3`, the sum of all vertex weights
becomes `4052`, and it is impossible to make it larger.

## Sample Input 2

```text
4 6
1 2
2 3
1 4
2 4
1 3
3 4
```

## Sample Output 2

```text
4052
```

## Sample Input 3

```text
2 1
1 2
```

## Sample Output 3

```text
2026
```

## Solution

Let

```text
x_v = W_v / 2026
```

Then each `x_v` must satisfy `0 <= x_v <= 1`, and every edge `(u, v)` must satisfy

```text
x_u + x_v <= 1
```

Maximizing `sum W_v` is the same as maximizing `sum x_v` and multiplying the result
by `2026`.

This is the linear-programming relaxation of maximum independent set:

```text
maximize   sum x_v
subject to x_u + x_v <= 1  for every edge (u, v)
           0 <= x_v <= 1
```

It is more convenient to use the complementary variables

```text
y_v = 1 - x_v
```

The edge condition becomes

```text
y_u + y_v >= 1
```

and maximizing `sum x_v` becomes minimizing `sum y_v`. This is the fractional vertex
cover problem:

```text
minimize   sum y_v
subject to y_u + y_v >= 1  for every edge (u, v)
           0 <= y_v <= 1
```

By linear-programming duality, the minimum fractional vertex cover value equals the
maximum fractional matching value. In a general graph, a maximum fractional matching can
use values `0`, `1/2`, and `1`. The doubled bipartite graph lets us compute twice this
fractional matching value as an ordinary integral max flow.

### Doubled Bipartite Graph

Create two copies of every original vertex:

- left copy `v`;
- right copy `N + v`.

Add the following directed edges, all with capacity `1`:

- source `S -> left(v)` for every vertex `v`;
- right(v) `-> T` for every vertex `v`;
- for every original undirected edge `(u, v)`, add `left(u) -> right(v)` and
  `left(v) -> right(u)`.

Any integral matching in this bipartite double cover corresponds to choosing directed
uses of original edges. Since every original vertex has capacity `1` on the left side
and capacity `1` on the right side, the flow value is exactly twice the best fractional
matching value in the original graph:

- an original matching edge of value `1` can be represented by the two directed edges
  `left(u) -> right(v)` and `left(v) -> right(u)`;
- an odd cycle with value `1/2` on every edge becomes an alternating integral matching
  in the double cover.

So if the computed max flow is `F`, then

```text
minimum fractional vertex cover = F / 2
maximum sum of x_v = N - F / 2
```

Finally, multiply by `2026`:

```text
answer = 2026 * (N - F / 2)
       = 2026 * N - 1013 * F
```

This is exactly what the implementation returns.

### Why Integer Weights Are Enough

The original problem asks for integer vertex weights. The fractional vertex-cover /
fractional-matching optimum of a graph is half-integral, so each optimal `x_v` can be
chosen from `{0, 1/2, 1}`. Therefore each optimal weight can be chosen from
`{0, 1013, 2026}`, all of which are integers. Solving the fractional problem therefore
also solves the original integer-weight problem.

## Correctness

We prove that the algorithm returns the maximum possible total vertex weight.

First, scaling every weight by `2026` transforms the original constraints into
`0 <= x_v <= 1` and `x_u + x_v <= 1` for every edge. This scaling preserves the order of
all feasible solutions, so maximizing the original weight sum is equivalent to maximizing
`sum x_v`.

Second, define `y_v = 1 - x_v`. For every edge `(u, v)`, the constraint
`x_u + x_v <= 1` is equivalent to `y_u + y_v >= 1`. Also,
`sum x_v = N - sum y_v`. Therefore maximizing `sum x_v` is equivalent to minimizing
`sum y_v` over all fractional vertex covers.

Third, by linear-programming duality, the minimum fractional vertex cover value equals
the maximum fractional matching value of the original graph. The doubled bipartite graph
constructed by the algorithm has an integral max flow whose value is twice that maximum
fractional matching value. Each unit of original fractional matching contributes two
units in the double cover, and every integral matching in the double cover projects back
to a feasible fractional matching in the original graph by halving the selected directed
edge uses. Hence, if the max flow is `F`, the minimum fractional vertex cover value is
`F / 2`.

Combining these facts, the maximum value of `sum x_v` is `N - F / 2`. Multiplying by
`2026`, the maximum original weight sum is
`2026 * N - 1013 * F`, which is exactly the value returned by the algorithm.

Finally, the fractional optimum is half-integral, so the corresponding original weights
can be chosen as integers. Thus the value computed from the fractional formulation is
attainable in the original problem. Therefore the algorithm is correct.

## Complexity

The flow network has `2N + 2` vertices. It has `2M` edges between the two vertex copies,
plus `2N` source/sink edges, so `O(N + M)` edges in total.

The implementation uses Dinic's algorithm on this unit-capacity bipartite network.

Time complexity: polynomial in `N + M`; fast enough for `N, M <= 5 * 10^4` in this
unit-capacity bipartite graph.

Memory complexity: `O(N + M)`.
