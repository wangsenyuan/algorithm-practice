# Codeforces 123E - Maze

https://codeforces.com/problemset/problem/123/E

## Statement

A maze is represented by a tree: an undirected graph where exactly one path
exists between each pair of vertices.

The entrance vertex and the exit vertex are chosen with some probability. The
exit is searched for by Depth First Search. If there are several possible ways to
move, the next move is chosen equiprobably. Consider the following pseudo-code:

```text
DFS(x)
    if x == exit vertex then
        finish search
    flag[x] <- TRUE
    random shuffle the vertices' order in V(x)
        // all permutations are equally likely
    for i <- 1 to length[V] do
        if flag[V[i]] = FALSE then
            count++;
            DFS(y);
    count++;
```

`V(x)` is the list of vertices adjacent to `x`. Initially, every value in the
`flag` array is `FALSE`. DFS starts with the entrance vertex as its parameter.
When the search is finished, `count` contains the number of moves.

Your task is to compute the mathematical expectation of the number of moves
needed to exit the maze.

## Input

The first line contains an integer `n` (`1 <= n <= 100000`) — the number of
vertices in the graph.

The next `n - 1` lines contain pairs of integers `ai` and `bi`, denoting an edge
between vertices `ai` and `bi` (`1 <= ai, bi <= n`). It is guaranteed that the
given graph is a tree.

The next `n` lines contain pairs of non-negative numbers `xi` and `yi`, which
describe the probability of choosing the `i`-th vertex as the entrance and exit,
respectively.

The probability that vertex `i` is chosen as the entrance is:

```text
xi / (x1 + x2 + ... + xn)
```

The probability that vertex `i` is chosen as the exit is:

```text
yi / (y1 + y2 + ... + yn)
```

The sum of all `xi` and the sum of all `yi` are positive and do not exceed
`10^6`.

## Output

Print the expectation of the number of moves.

The absolute or relative error must not exceed `10^-9`.

## Samples

```text
2
1 2
0 1
1 0
```

```text
1.00000000000000000000
```

```text
3
1 2
1 3
1 0
0 2
0 3
```

```text
2.00000000000000000000
```

```text
7
1 2
1 3
2 4
2 5
3 6
3 7
1 1
1 1
1 1
1 1
1 1
1 1
1 1
```

```text
4.04081632653
```

## Note

In the first sample, the original Codeforces note says the entrance is always
`1` and the exit is always `2`, but that contradicts the probability formula and
the sample input. Since `(x1, y1) = (0, 1)` and `(x2, y2) = (1, 0)`, the entrance
vertex is always `2` and the exit vertex is always `1`. The expected number of
moves is still `1`, because the tree has only one edge.

In the second sample, the entrance vertex is always `1`. The exit vertex is `2`
with probability `2/5`, and `3` with probability `3/5`. The mathematical
expectations for exit vertices `2` and `3` are equal by symmetry. During the
first move, DFS goes directly to the exit with probability `0.5`; otherwise it
goes to the other vertex first with probability `0.5`. The number of moves is
`1` in the first case and `3` in the second case, so the total expectation is:

```text
2/5 * (1 * 0.5 + 3 * 0.5) + 3/5 * (1 * 0.5 + 3 * 0.5)
```

## ideas
1. 先不考虑从哪个点enter，如果从u开始处理，且它的parent = p时，从u开始，exit前，count的贡献为dp[u][p?]
2. 如果u是一个exit的概率是p[u], dp[u] = (1 - p[u]) * sum(dp[v] * (1/children count))
3. sum部分好像不大对？因为进入子树v，如果从v离开了，那么就是dp[v], 如果没有离开，那么这个count是子节点的数量
4. 有点乱。感觉这个思路不对。
5. 考虑每天边的贡献，是否可行？
6. 对于边(u, v) 如果在u这边是entrance的概率，和v这边是exit的概率知道了，ent(u) * ext(v) 是不是就是这条边的贡献？
7. 不对，这里没有那个randomshuffle的贡献

## Solution

The useful view is to fix the entrance `s` and exit `t`, and look only at the
path from `s` to `t`.

When DFS is at a path vertex `u`, suppose the next vertex on the path to `t` is
`v`. The DFS randomly orders all currently unvisited neighbors of `u`. For any
off-path neighbor `w`, the event "`w` is tried before `v`" has probability
`1/2`, because in a random permutation the relative order of `w` and `v` is
equally likely.

If DFS enters an off-path component of size `sz`, it will completely explore
that component and come back to `u`. The cost is:

```text
2 * sz
```

because every vertex in that component contributes one move entering it and one
move returning from it.

Therefore the expected contribution of that off-path component is:

```text
(1/2) * 2 * sz = sz
```

The path edge `u -> v` itself always contributes `1`.

So for a fixed ordered pair `(s, t)`:

```text
answer(s, t)
  = distance(s, t)
    + sum of sizes of off-path components that may be explored
```

where the off-path component at a path vertex excludes:

- the component containing the previous path vertex, because it is already
  visited;
- the component containing the next path vertex, because choosing it continues
  toward the exit.

If `s == t`, the answer is `0`.

## Aggregating over probabilities

Normalize the entrance and exit weights first:

```text
enter[i] = xi / sum(x)
exit[i]  = yi / sum(y)
```

For an oriented edge `u -> v`, define:

```text
S(u, v) = number of vertices in the component containing v after removing u-v
X(u, v) = sum of enter probabilities in that component
Y(u, v) = sum of exit probabilities in that component
```

These values are available in `O(1)` per oriented edge after one arbitrary tree
rooting and a subtree DP.

### Path-edge contribution

An edge separates the entrance and exit exactly when the entrance is on one side
and the exit is on the other side. Summing over oriented edges:

```text
sum over oriented edges u -> v:
    X(u, v) * (1 - Y(u, v))
```

This counts one move for each edge on the entrance-exit path.

### Off-path contribution at the entrance

If `u` is the entrance and the next path vertex is `v`, the previous side is
empty. The off-path size is:

```text
n - 1 - S(u, v)
```

So the contribution is:

```text
sum over oriented edges u -> v:
    enter[u] * Y(u, v) * (n - 1 - S(u, v))
```

### Off-path contribution at internal path vertices

Now consider an internal path vertex `u`. Let `p` be the neighbor toward the
entrance and `v` be the neighbor toward the exit. They must be two distinct
neighbors of `u`.

The entrance is in component `p`, and the exit is in component `v`. The off-path
size is:

```text
n - 1 - S(u, p) - S(u, v)
```

The direct sum would be:

```text
sum over u
sum over distinct neighbors p, v of u:
    X(u, p) * Y(u, v) * (n - 1 - S(u, p) - S(u, v))
```

Doing this naively may be quadratic for a high-degree vertex. Instead, for each
fixed `u`, precompute:

```text
totalY     = sum over neighbors q of u: Y(u, q)
totalYSize = sum over neighbors q of u: Y(u, q) * S(u, q)
```

Then for each possible previous neighbor `p`, sum all valid next neighbors
`v != p` in `O(1)`:

```text
X(u, p) *
(
    (n - 1 - S(u, p)) * (totalY - Y(u, p))
    - (totalYSize - Y(u, p) * S(u, p))
)
```

This makes the whole solution linear in the number of tree edges.

## Complexity

Rooting the tree, computing subtree sizes and probability sums, and aggregating
all contributions each take `O(n)`.

The memory usage is also `O(n)`.
