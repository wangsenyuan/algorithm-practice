# E. Ancient Tree

[Problem link](https://codeforces.com/problemset/problem/2127/E)

**Contest:** [Atto Round 1 (Codeforces Round 1041, Div. 1 + Div. 2)](https://codeforces.com/contest/2127)

time limit per test: 4 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given a rooted tree of `n` vertices, where vertex `1` is the root. Each vertex has an integer weight `w_i` and a color `c_i`, where the colors are integers between `1` and `k`. Some vertices have lost their colors, represented by `c_i = 0`.

Vertex `v` is **cutie** if there exist two vertices `x` and `y` such that:

- `lca(x, y) = v`
- `c_x = c_y`
- `c_x != c_v`

The cost of the tree is the sum of weights of all cutie vertices.

Assign colors between `1` and `k` to all vertices that have lost their colors. Find the minimum possible cost among all valid colorings and provide one coloring with that cost.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases.

Each test case contains:

- one line with two integers `n` and `k` (`3 <= n <= 2 * 10^5`, `2 <= k <= n`) — the number of vertices and the number of colors;
- one line with `n` integers `w_1, w_2, ..., w_n` (`1 <= w_i <= 10^9`) — the vertex weights;
- one line with `n` integers `c_1, c_2, ..., c_n` (`0 <= c_i <= k`) — the vertex colors, where `c_i = 0` means vertex `i` has lost its color;
- `n - 1` lines, each with two integers `u` and `v` (`1 <= u, v <= n`) — an edge of the tree.

The given edges form a tree. The sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, print two lines:

- one integer — the minimum possible cost;
- `n` integers `c'_1, c'_2, ..., c'_n` — a coloring with that cost.

The coloring must satisfy `c'_i = c_i` if `c_i != 0`, and `1 <= c'_i <= k` if `c_i = 0`. If there are multiple colorings with the minimum cost, print any of them.

## Example

### Input

```text
4
4 4
5 5 5 5
1 0 2 3
1 2
1 3
1 4
5 2
3 1 4 1 5
1 2 1 2 2
1 4
2 1
3 4
4 5
11 3
3 1 4 3 1 4 3 1 4 5 6
0 0 0 2 1 2 1 2 2 1 1
1 2
2 3
2 4
2 5
2 6
1 7
7 8
7 9
7 10
7 11
4 3
2 3 2 3
2 1 0 0
3 1
1 2
2 4
```

### Output

```text
0
1 4 2 3
3
1 2 1 2 2
7
2 3 1 2 1 2 1 2 2 1 1
0
2 1 3 1
```

### Note

In the first test case, the missing color of vertex `2` has four choices:

- `c_2 = 1` makes no vertex cutie, so the cost is `0`;
- `c_2 = 2` makes vertex `1` cutie, since `c_2 = c_3 = 2`, `lca(2, 3) = 1` and `c_1 != 2`. The cost is `w_1 = 5`;
- `c_2 = 3` makes vertex `1` cutie, since `c_2 = c_4 = 3`, `lca(2, 4) = 1` and `c_1 != 3`. The cost is `w_1 = 5`;
- `c_2 = 4` makes no vertex cutie, so the cost is `0`.

The minimum cost is `0`.

In the second test case every vertex already has a color. Since `c_5 = c_2 = 2`, `lca(2, 5) = 1` and `c_1 != 2`, vertex `1` is cutie and the cost is `w_1 = 3`.

In the third test case, one coloring with cost `7` is `c = [2, 3, 1, 2, 1, 2, 1, 2, 2, 1, 1]`. Other valid colorings include:

- `[3, 1, 2, 2, 1, 2, 1, 2, 2, 1, 1]` with cutie vertices `1, 2, 3, 7` and cost `11`;
- `[3, 2, 1, 2, 1, 2, 1, 2, 2, 1, 1]` with cutie vertices `1, 2, 7` and cost `7`.

No coloring with cost smaller than `7` exists.

## Solution

A vertex `v` is cutie exactly when some color `x != c_v` appears in at least two different child-subtrees of `v`. Only already-colored vertices can force this: an uncolored vertex can always take a color already present in its subtree (or its parent's color), and that does not create any new cutie vertex. So the minimum cost is the sum of weights of vertices that are cutie no matter how the missing colors are filled.

If every `c_i = 0`, paint the whole tree with color `1`. Every vertex then has the same color as its descendants, so the cost is `0`.

### Forced cuties from each color

Root the tree at vertex `1` and build binary lifting. Record DFS discovery times.

For each color `x`, take the vertices already painted `x` and sort them by DFS time. A standard lemma says that a vertex `p` is the LCA of some pair of these vertices if and only if `p` is the LCA of two **consecutive** vertices in that order. Those LCAs are exactly the internal nodes of the virtual tree of color `x`.

For each such `p`:

- If `p` is already colored and `c_p != x`, then `p` is always cutie.
- If `p` is uncolored, count how many distinct colors use `p` as such an LCA, and remember one of them.

This yields three kinds of uncolored vertices:

- **one forcing color** — paint `p` with that color, so `p` itself is not cutie;
- **two or more forcing colors** — `p` is always cutie (it cannot match every such color);
- **none** — not forced.

The answer is the total weight of every always-cutie vertex (precolored ones with a conflicting color, and uncolored ones with two or more forcing colors).

### Fill the rest

After the forced assignments, walk down from the root. A still-uncolored vertex copies the color of its nearest colored ancestor (or any existing color, if the root is still empty). Painting with the parent color cannot make the parent cutie, because the two vertices then share a color.

### Complexity

Binary lifting and the per-color consecutive-LCA pass are `O(n log n)`. Extra memory is `O(n log n)` for the lifting table. The sum of `n` over tests is `2 * 10^5`.
