# E - Distribute Bunnies

[Problem link](https://atcoder.jp/contests/abc434/tasks/abc434_e)

**Contest:** [Sky Inc, Programming Contest 2025 (AtCoder Beginner Contest 434)](https://atcoder.jp/contests/abc434)

time limit: 2 sec

memory limit: 1024 MiB

score: 500 points

There are `N` rabbits numbered `1` to `N` on a number line. Rabbit `i` is at coordinate `X_i`.
Multiple rabbits may share a coordinate.

Each rabbit has **jumping power** `R_i`. Every rabbit jumps exactly once. A rabbit at coordinate `x`
with jumping power `r` moves to `x + r` or `x - r`.

Choose the direction of each jump freely. Find the maximum possible number of distinct coordinates
where rabbits are present after all jumps.

## Constraints

- `1 <= N <= 2 * 10^5`
- `-10^9 <= X_i <= 10^9`
- `1 <= R_i <= 10^9`
- All input values are integers.

## Input

```text
N
X_1 R_1
X_2 R_2
...
X_N R_N
```

## Output

Print the maximum possible number of distinct coordinates after jumping.

## Sample Input 1

```text
3
4 1
2 3
4 5
```

## Sample Output 1

```text
3
```

One optimal assignment:

- Rabbit 1: `4 - 1 = 3`
- Rabbit 2: `2 + 3 = 5`
- Rabbit 3: `4 - 5 = -1`

## Sample Input 2

```text
6
2 1
3 2
6 1
5 2
4 3
4 1
```

## Sample Output 2

```text
4
```

One optimal assignment:

- Rabbit 1: `2 - 1 = 1`
- Rabbit 2: `3 + 2 = 5`
- Rabbit 3: `6 + 1 = 7`
- Rabbit 4: `5 + 2 = 7`
- Rabbit 5: `4 - 3 = 1`
- Rabbit 6: `4 - 1 = 3`

## Sample Input 3

```text
10
1000000000 1000000000
1000000000 1
-1000000000 1000000000
-1000000000 1
0 1
2 1
1 2
4 1
3 2
4 3
```

## Sample Output 3

```text
9
```

## Solution

For rabbit `i`, the final coordinate is either:

```text
X_i - R_i
X_i + R_i
```

Treat every possible coordinate as a vertex. Then each rabbit becomes an edge connecting its two possible
coordinates. Choosing the rabbit's jump direction is the same as orienting this edge toward one of its endpoints;
the endpoint is the coordinate occupied by that rabbit.

Now consider one connected component of this graph. If it has `E` edges and `V` vertices, the contribution of this
component is:

```text
min(E, V)
```

The answer cannot exceed `E`, because each rabbit occupies only one coordinate. It also cannot exceed `V`, because
there are only `V` coordinates in the component.

These upper bounds are achievable:

- If `E < V`, the component is a tree-like component with enough structure to orient each edge toward a different
  vertex, covering `E` distinct vertices.
- If `E >= V`, the component contains at least one cycle. Choose a cycle and orient its edges around the cycle to
  cover all vertices on it. For every tree branch attached to the already covered part, orient each branch edge
  away from the covered side, so every new vertex is covered. Thus all `V` vertices can be occupied.

So we only need to build this graph, find connected components, count edges and vertices in each component, and add
`min(E, V)`.

The implementation compresses all coordinates `X_i - R_i` and `X_i + R_i`, builds the incidence lists, and runs BFS
over coordinate vertices. During BFS, each rabbit edge is counted once with a `marked` array.

## Correctness

Each rabbit has exactly two possible final coordinates, so representing it as an edge between those two coordinate
vertices preserves all choices. Selecting one endpoint of each edge is exactly selecting each rabbit's jump
direction.

For any connected component, no assignment can occupy more than the number of rabbits in it or more than the number
of coordinate vertices in it, so `min(E, V)` is an upper bound.

If `E < V`, a connected component has no cycle surplus; by rooting a spanning tree, each edge can be assigned to a
distinct child-side vertex, covering `E` vertices. If `E >= V`, take one cycle to cover its vertices, then orient
all edges in attached tree parts outward from the cycle; every remaining vertex is covered exactly when first
entered. Therefore the upper bound `min(E, V)` is always attainable.

The algorithm sums this value over all connected components. Components are independent because they have disjoint
coordinate vertices and disjoint rabbits, so the sum is the global maximum.

## Complexity

There are at most `2N` coordinate vertices and `N` edges. Coordinate compression costs `O(N log N)`, and the BFS
over the compressed graph is `O(N)`. The memory usage is `O(N)`.
