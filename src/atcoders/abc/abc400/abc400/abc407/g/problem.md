# G - Domino Covering SUM

[Problem link](https://atcoder.jp/contests/abc407/tasks/abc407_g)

**Contest:** [AtCoder Beginner Contest 407](https://atcoder.jp/contests/abc407)

time limit: 2 sec

memory limit: 1024 MiB

score: 600 points

There is a grid with `H` rows and `W` columns. Cell `(i, j)` contains an integer `A_{i,j}`.

Place zero or more dominoes on the grid. Each domino covers two adjacent cells:

- `(i, j)` and `(i, j+1)` for `1 <= i <= H`, `1 <= j < W`, or
- `(i, j)` and `(i+1, j)` for `1 <= i < H`, `1 <= j <= W`.

No cell may be covered by more than one domino.

For a domino placement, define its **score** as the sum of all integers in cells **not** covered by any
domino.

Find the maximum possible score.

## Constraints

- `1 <= H`
- `1 <= W`
- `H * W <= 2000`
- `-10^12 <= A_{i,j} <= 10^12` (`1 <= i <= H`, `1 <= j <= W`)
- All input values are integers

## Input

```text
H W
A_{1,1} A_{1,2} ... A_{1,W}
A_{2,1} A_{2,2} ... A_{2,W}
...
A_{H,1} A_{H,2} ... A_{H,W}
```

## Output

Print the answer.

## Sample Input 1

```text
3 4
3 -1 -4 1
-5 9 -2 -6
-5 3 -5 8
```

## Sample Output 1

```text
23
```

## Sample Input 2

```text
5 5
-70 11 -45 -54 -30
-99 39 -83 -69 -77
-48 -21 -43 -96 -24
-54 -65 21 -88 -44
-90 -33 -67 -29 -62
```

## Sample Output 2

```text
39
```

## Sample Input 3

```text
8 9
-74832 16944 58683 32965 97236 -52995 43262 -51959 40883
-58715 13846 24919 65627 -11492 -63264 29966 -98452 -75577
40415 77202 15542 -50602 83295 85415 -35304 46520 -38742
37482 56721 -38521 63127 55608 95115 42893 10484 70510
53019 40623 25885 -10246 70973 32528 -33423 19322 52097
79880 74931 -58277 -33783 91022 -53003 11085 -65924 -63548
78622 -77307 81181 46875 -81091 63881 11160 -82217 -55492
62770 39530 -95923 92440 -69899 77737 89392 -14281 84899
```

## Sample Output 3

```text
2232232
```

## Solution

Let `S` be the sum of every cell before placing dominoes. If a domino covers adjacent cells `u` and `v`, the final score changes by `-(A[u] + A[v])`: those two values are removed from `S`.

Equivalently, give this domino a cost of `A[u] + A[v]`. For a set of dominoes, the result is:

```text
S - (sum of costs of chosen dominoes)
```

The chosen dominoes may not share a cell, so they are a matching in the grid graph. We therefore need a matching with the minimum total cost.

### Ignore non-negative dominoes

A domino with cost at least zero cannot improve the answer. Removing it does not decrease the score and only frees its two cells. Thus an optimal answer uses only edges with:

```text
A[u] + A[v] < 0
```

This also means every selected domino is beneficial, so we may stop as soon as no negative-cost augmentation remains.

### Turn the matching into flow

Color the grid like a chessboard. Every adjacent pair consists of one black cell and one white cell, so the graph is bipartite.

Build a flow network as follows:

1. Add a source `s` and sink `t`.
2. For each black cell `u`, add `s -> u` with capacity `1` and cost `0`.
3. For every adjacent black/white pair `(u, v)` whose sum is negative, add `u -> v` with capacity `1` and cost `A[u] + A[v]`.
4. For each white cell `v`, add `v -> t` with capacity `1` and cost `0`.

One unit of flow along `s -> u -> v -> t` selects exactly one domino. Capacities one on the source and sink edges ensure no cell is selected twice. Conversely, every legal domino placement gives such a flow. Therefore, a minimum-cost flow represents the best matching.

### Successive shortest augmenting paths

The solver repeatedly finds the cheapest residual path from `s` to `t`.

- Initially, negative edge costs exist, so shortest distances are calculated once to initialize vertex potentials.
- Thereafter, reduced costs are non-negative; Dijkstra finds each next shortest augmenting path.
- Reverse residual edges allow a later augmentation to replace earlier dominoes, which is necessary for a globally optimal matching.
- If the cheapest path has cost `>= 0`, taking another domino cannot improve the score, so the algorithm stops.

If the total cost of all selected flow paths is `C` (which is non-positive), the answer is `S - C`.

All values use `int64`: the sum of up to 2000 cell values can be as large as `2 * 10^15` in magnitude.

### Correctness

Every unit of flow corresponds to one valid domino because it uses one black cell, one adjacent white cell, and capacity one prevents either endpoint from being reused. Every valid placement of negative-cost dominoes similarly maps to a feasible flow of the same total cost.

The successive shortest-path algorithm maintains a minimum-cost flow for its current value. Its residual graph includes reverse edges, so it considers both adding a domino and exchanging previously selected dominoes. When its cheapest possible augmentation is non-negative, no feasible matching change can lower the matching cost. Hence the obtained flow has the minimum total cost among all useful domino matchings, and `S - C` is the maximum possible uncovered-cell sum.

### Complexity

Let `V = H * W + 2` and let `E = O(H * W)`, since every cell has at most four neighbors. At most `floor(H * W / 2)` augmentations are performed. With potentials and Dijkstra, the running time is `O(VE + F * E log V)`, where `F <= 1000`, and memory usage is `O(V + E)`.
