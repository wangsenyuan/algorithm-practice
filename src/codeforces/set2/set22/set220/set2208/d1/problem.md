# Problem

You once had an undirected tree with `n` nodes. To make the tree look more interesting, you assigned an **arbitrary direction** to each of the `n - 1` edges.

Over time you forgot the tree structure. You found a note that records, **after** orienting the edges, for every ordered pair `(u, v)` with `1 ≤ u, v ≤ n`, whether `u` can reach `v` in the resulting directed graph.

You must recover **some** tree structure and edge directions consistent with the note, or report that none exists. If multiple solutions exist, print any.

For a directed graph, `x` can reach `y` if and only if there is a sequence of nodes `u1, u2, ..., uk` such that `u1 = x`, `uk = y`, and for every `i` from `2` to `k` the directed edge `u(i-1) → ui` exists. In particular, every node can reach itself.

## Input

The first line contains integer `t` (`1 ≤ t ≤ 10^4`) — the number of test cases.

Each test case:

- First line: integer `n` (`2 ≤ n ≤ 500`) — number of nodes.
- Next `n` lines: string `si` of length `n`, consisting only of `'0'` and `'1'`.
  The `j`-th character of `si` is `'1'` if and only if node `i` can reach node `j` after the edges are directed.

It is guaranteed that the sum of `n^3` over all test cases does not exceed `500^3`.

## Output

For each test case:

- If a solution exists, print `Yes` (case-insensitive accepted, e.g. `yEs`, `YES`).
- Then print `n - 1` lines, each with two integers `x` and `y`, meaning there is a directed edge `x → y` in your oriented tree.

If no solution exists, print `No` (case-insensitive accepted).

If there are multiple valid trees/orientations, print any.

## Example

### Input

```text
11
4
1000
1111
1010
0001
4
1111
0111
0010
0111
4
0011
0111
0011
0001
4
1000
0110
0010
1111
4
1000
0110
1010
1111
5
10000
01011
00111
00010
00001
5
10000
11000
10101
10111
00001
5
10000
01101
00100
01110
10001
4
1100
0100
0011
0001
4
1110
0100
0010
0101
3
100
111
101
```

### Output

```text
Yes
2 3
2 4
3 1
No
No
Yes
2 3
4 1
4 2
No
No
Yes
2 1
3 1
3 5
4 3
No
No
Yes
1 2
1 3
4 2
Yes
2 3
3 1
```

## Note

In the first test case, nodes `1` and `4` can only reach themselves, node `2` can reach every node, and node `3` can reach only nodes `1` and `3`. The printed edges satisfy these reachability constraints.

In the second test case, no valid tree and orientation exist.

## Ideas

1. `s[i][j] = 1` 表示从 i 能够到达 j，也就是存在一条 path: i -> ... -> j
2. 如果 i, j 中间不存在 u，使得 `s[i][u] = 1` 且 `s[u][j] = 1`，那么 i->j 成立
