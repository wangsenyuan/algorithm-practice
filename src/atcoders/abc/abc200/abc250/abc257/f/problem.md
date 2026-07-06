# F - Teleporter Setting

[Problem link](https://atcoder.jp/contests/abc257/tasks/abc257_f)

**Contest:** [AtCoder Beginner Contest 257](https://atcoder.jp/contests/abc257)

time limit: 2 sec

memory limit: 1024 MiB

score: 500 points

There are N towns numbered 1 through N and M teleporters. Each teleporter connects two towns
bidirectionally; using one takes 1 minute.

The i-th teleporter connects towns U_i and V_i. If U_i = 0, one endpoint is town V_i and the other
endpoint is undetermined.

For each i = 1, 2, ..., N, answer the following:

> If every undetermined teleporter is connected to town i, what is the minimum number of minutes
> needed to travel from town 1 to town N using teleporters only? Print -1 if impossible.

## Constraints

- 2 <= N <= 3 * 10^5
- 1 <= M <= 3 * 10^5
- 0 <= U_i < V_i <= N
- (U_i, V_i) != (U_j, V_j) for i != j
- All input values are integers

## Input

```text
N M
U_1 V_1
...
U_M V_M
```

## Output

Print N integers separated by spaces. The k-th integer is the answer when i = k.

## Sample Input 1

```text
3 2
0 2
1 2
```

## Sample Output 1

```text
-1 -1 2
```

When undetermined teleporters connect to town 1 or town 2, town 3 is unreachable from town 1.
When they connect to town 3, town 1 -> town 2 -> town 3 takes 2 minutes.

Note that undetermined endpoints may create self-loops or multiple teleporters between the same pair
of towns.

## Sample Input 2

```text
5 5
1 2
1 3
3 4
4 5
0 2
```

## Sample Output 2

```text
3 3 3 3 2
```

## Solution

Ignore the undetermined endpoint first. Build the graph using only edges whose endpoints are both
normal towns, and remember every town `v` that appears in an edge `(0, v)`.

Run BFS twice on this normal graph:

1. `dp1[x]`: shortest distance from town `1` to town `x`.
2. `dp2[x]`: shortest distance from town `N` to town `x`.

For a query town `u`, every undetermined edge becomes an edge between `u` and some remembered town
`v`. A shortest path from `1` to `N` can be one of these forms:

1. Use no undetermined edge: `dp1[N]`.
2. Use one undetermined edge from the left side:
   `1 -> ... -> v -> u -> ... -> N`, cost `dp1[v] + 1 + dp2[u]`.
3. Use one undetermined edge from the right side:
   `1 -> ... -> u -> v -> ... -> N`, cost `dp1[u] + 1 + dp2[v]`.
4. Use two undetermined edges:
   `1 -> ... -> v1 -> u -> v2 -> ... -> N`, cost `dp1[v1] + 2 + dp2[v2]`.

The chosen special endpoint should not be `u`, because edge `(0, u)` becomes a self-loop for query
`u` and cannot help. To answer all queries quickly, keep the best few candidates among remembered
towns by `dp1[v]` and by `dp2[v]`. Then for each `u`, pick the best candidate whose town is not `u`.

The implementation stores up to three best candidates for each side. This is enough because only one
town, the current query `u`, may need to be skipped.

For each query, take the minimum over all valid candidate costs. If every candidate is unreachable,
print `-1`.

### Correctness

The two BFS runs compute shortest distances in the graph that does not use any undetermined edge, so
any segment between normal towns in a final path is measured correctly by `dp1` or `dp2`.

For a fixed query town `u`, every undetermined edge is incident to `u`. Therefore any path that uses
undetermined edges can enter or leave the normal graph only through `u` and one remembered endpoint.
A shortest path never needs more than two useful undetermined edges: one to reach `u` from the side of
town `1`, and one to leave `u` toward the side of town `N`. Extra visits through `u` only add cycles
and cannot improve the distance.

Thus every optimal path belongs to one of the four forms listed above. The algorithm checks all four
forms using the best available remembered endpoints, excluding only endpoints equal to `u` because they
create self-loops. Hence it finds a distance no larger than the true optimum.

Conversely, every candidate distance used by the algorithm corresponds to an actual path in the graph
for query `u`: normal shortest-path segments plus one or two valid undetermined edges. Therefore the
algorithm never reports a value smaller than the true shortest distance.

The reported value is exactly the shortest distance for each query, or `-1` if no candidate path
exists.

### Complexity

Let `N` be the number of towns and `M` the number of edges. The two BFS runs take `O(N + M)` time.
Keeping the best special endpoints and answering all `N` queries are linear.

The total time complexity is `O(N + M)`, and the memory usage is `O(N + M)`.
