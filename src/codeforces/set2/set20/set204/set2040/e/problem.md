# E. Control of Randomness

[Problem link](https://codeforces.com/problemset/problem/2040/E)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

You are given a tree with `n` vertices.

Place a robot in some vertex `v != 1`, and suppose you initially have `p` coins.
Consider the following process, where in the `i`-th step (starting from `i = 1`):

- If `i` is odd, the robot moves to an adjacent vertex in the direction of
  vertex `1`.
- If `i` is even, you can either pay one coin (if there are some left) and then
  the robot moves toward vertex `1`, or not pay, and then the robot moves to an
  adjacent vertex chosen uniformly at random.

The process stops as soon as the robot reaches vertex `1`. Let `f(v, p)` be the
minimum possible expected number of steps if coins are spent optimally.

Answer `q` queries: for each query `(v_i, p_i)`, find `f(v_i, p_i)` modulo
`998244353`.

Formally, if the answer is an irreducible fraction `P/Q`, output
`P * Q^{-1} mod 998244353`.

## Constraints

- `1 <= t <= 10^3`
- `2 <= n <= 2 * 10^3`
- `1 <= q <= 2 * 10^3`
- Tree edges: `1 <= u_i, v_i <= n`, `u_i != v_i`
- Queries: `2 <= v_i <= n`, `0 <= p_i <= n`
- Sum of `n` over all test cases does not exceed `2 * 10^3`
- Sum of `q` over all test cases does not exceed `2 * 10^3`

## Input

The first line contains the number of test cases `t`.

For each test case:

- The first line contains two integers `n` and `q`.
- The next `n - 1` lines contain the edges of the tree.
- The next `q` lines each contain two integers `v_i` and `p_i`.

## Output

For each test case, print `q` integers: the values of `f(v_i, p_i)` modulo
`998244353`.

## Example

```text
Input
2
4 4
1 2
2 3
2 4
2 0
3 0
4 0
3 1
12 10
1 2
2 3
2 4
1 5
5 6
6 7
6 8
6 9
8 10
10 11
10 12
6 0
9 0
10 0
11 0
3 1
7 1
10 1
12 1
12 2
11 12

Output
1
6
6
2
4
9
8
15
2
3
6
9
5
5
```

## Solution

First solve the problem without queries and without forced (coin) moves.

Consider the path structure. The current vertex `v` has grandparent `u`. On an
odd move, the robot goes to the parent of `v`. On the following even move from
that parent:

- with probability `1/(x+1)` it goes to `u`;
- otherwise it goes to a brother of `v`.

Here `x` is the number of brothers of `v`, including itself (so the parent of
`v` has degree `x+1`). Landing on a brother changes nothing — the next odd step
returns to the same parent and the situation repeats. Thus `v` and all its
brothers share the same answer.

Let `d[v]` be the expected number of steps from `v` (with no coins). Then:

```text
d[v] = 2 + (1/(x+1)) * d[u] + (x/(x+1)) * d[v]
```

which rearranges to:

```text
d[v] = d[u] + 2 * (x + 1)
```

The path therefore consists of **blocks of height 2**: the robot repeatedly
tries to clear the next block until it succeeds, then moves on to the next one.
Clearing a block with parameter `x` costs `2 * (x + 1)` expected steps.

With coins, paying once **skips a block**: the robot clears it on the first try,
spending `2` instead of `2 * (x + 1)`. The saving is `2 * x`. So for a query
with `p` coins, greedily skip the `p` blocks with the largest `x` on the path
from `v` to the root.

Because odd and even steps alternate, maintain two multisets of the relevant
`x` values along the path to the root — one for odd depths and one for even
depths. Answer queries offline while DFS-traversing the tree: when a query is
encountered at the current vertex, take the first `p` (largest) elements of the
corresponding multiset and subtract their savings from the no-coin answer.

- Trivial per-query walk to the root: `O(n * q)`.
- Offline DFS with multisets:
  `O(n + sum_i p_i * (multiset cost))`.

### Detailed explanation

Root the tree at vertex `1`. For every non-root vertex, its parent is the next
vertex on the path toward `1`.

#### Divide the path into two-level blocks

Suppose the robot is at vertex `v`, its parent is `w`, and the parent of `w` is
`u`:

```text
v -> w -> u
```

Let `x` be the number of children of `w`. Thus `x` includes `v` and all of
`v`'s siblings. Since `w` is not the root, it also has the parent `u`, so:

```text
degree(w) = x + 1
```

Starting from `v`, the next two moves form one attempt to cross this block:

1. The odd move is forced, so the robot moves from `v` to `w`.
2. If no coin is used on the even move, one of the `x + 1` neighbors of `w`
   is chosen uniformly.

With probability `1/(x+1)`, the second move goes to `u`, so the robot advances
two levels toward the root. With probability `x/(x+1)`, it goes to one of the
children of `w`.

All children of `w` have the same state for this process: their next move is
odd and forced back to `w`. Consequently, after a failed attempt the robot is
in exactly the same probabilistic situation as before the attempt. The
particular child on which it landed does not matter.

Let `d[v]` be the expected number of remaining moves when there are no coins.
For a vertex whose grandparent is `u`, the recurrence is:

```text
d[v] = 2 + d[u] / (x + 1) + x * d[v] / (x + 1)
```

Multiplying by `x + 1` and rearranging gives:

```text
d[v] = d[u] + 2 * (x + 1)
     = d[u] + 2 * degree(w)
```

Equivalently, each attempt costs two moves and succeeds with probability
`1/(x+1)`. The expected number of attempts is `x + 1`, hence the expected cost
of the block is `2 * (x + 1)`.

There are two base cases:

- `d[1] = 0`.
- If `v` is a direct child of the root, then `d[v] = 1`, because the first odd
  move reaches the root and the process immediately stops.

Therefore, if `depth(v) = D`, the blocks have lower endpoints at depths
`D, D-2, D-4, ...`. If `D` is odd, one final forced move remains after all
complete blocks. Another way to write the no-coin answer is:

```text
d[v] = D % 2 + sum(2 * degree(center))
```

where `center` ranges over the parents in the relevant two-level blocks.

#### Effect of spending a coin

If a coin is spent on the even move from `w`, the robot is forced to move to
`u`. The block is then crossed on the first attempt and costs exactly two
moves. Its saving is:

```text
2 * (x + 1) - 2 = 2 * x
```

A failure without a coin returns the process to the same state, so it reveals
no useful new information. For an optimal policy, a block can consequently be
treated as one of two independent choices:

- do not spend a coin, paying expected cost `2 * (x + 1)`;
- spend one coin, paying cost `2` and saving `2 * x`.

The costs of different blocks add together. Thus allocating `p` coins is the
same as selecting at most `p` block savings. An exchange argument proves the
greedy choice: if a selected block has `x=a` and an unselected block has
`x=b>a`, exchanging them increases the saving by `2 * (b-a)`. Hence an optimal
selection consists of the `p` largest `x` values, or all blocks if fewer than
`p` exist.

#### Why two path multisets are enough

For a query at depth `D`, the centers of its blocks occur at depths:

```text
D-1, D-3, D-5, ...
```

They all have parity opposite to `D`. During a DFS, the implementation keeps
two multisets for the current root-to-vertex path:

- `tr[0]` stores values belonging to even-depth vertices;
- `tr[1]` stores values belonging to odd-depth vertices.

For every non-root vertex `w`, the stored value is:

```text
x[w] = degree(w) - 1
```

which is exactly the number of children of `w` when the tree is rooted at `1`.
The root is never inserted because reaching it ends the process; it cannot be
the center of a block.

At vertex `u` with depth `dep`, queries are answered before `u` is inserted.
The multisets therefore contain exactly the proper ancestors of `u`. The
relevant block centers are selected from:

```text
tr[(dep & 1) ^ 1]
```

because their depth parity is opposite to that of `u`. After answering the
queries at `u`, `x[u]` is inserted so that descendants can use `u` as a block
center. It is removed when DFS leaves the subtree.

#### Segment tree operations

All possible `x` values are sorted and deduplicated for coordinate
compression. Each node of the segment tree stores:

- `cnt`: how many active path vertices have values in this interval;
- `sum`: the sum of those active `x` values.

Larger compressed coordinates correspond to larger `x`. To compute the sum of
the largest `k` active values, `sumKmax` examines the right child first:

- if the right child contains at least `k` values, recurse only into it;
- otherwise, take its entire sum and obtain the remaining values from the left
  child.

For a query `(u, p)`, let `k` be the smaller of `p` and the number of relevant
blocks. The final answer is:

```text
answer = dp[u] - 2 * sumKmax(k)
```

Here `dp[u]` is computed during the same DFS:

```text
dp[u] = 1                                      if parent(u) = 1
dp[u] = dp[grandparent(u)] + 2 * degree(parent(u)) otherwise
```

The derived expectations and savings are integers, so the implementation only
needs to normalize the final result modulo `998244353`.

#### Correctness proof

**Lemma 1.** A block centered at `w`, with `x` children, costs
`2 * (x + 1)` expected moves when no coin is used.

**Proof.** Every attempt takes two moves. It succeeds by choosing the parent of
`w`, which happens with probability `1/(x+1)`. Otherwise the robot lands on a
child of `w` and returns to the same state for the next attempt. Therefore the
expected number of attempts is `x+1`, and the expected cost is
`2 * (x+1)`. **End of proof.**

**Lemma 2.** Spending a coin on a block with parameter `x` saves exactly
`2 * x` expected moves.

**Proof.** By Lemma 1, the block costs `2 * (x+1)` without a coin. A coin forces
the successful second move, making the cost exactly `2`. Their difference is
`2*x`. **End of proof.**

**Lemma 3.** For `p` available coins, spending them on the blocks with the
largest `x` values is optimal.

**Proof.** Block costs are additive, and by Lemma 2 a selected block contributes
saving `2*x`. If a solution selects `a` while leaving a larger `b` unselected,
exchanging them increases the total saving. Repeating this exchange produces
exactly the set of largest values without making the solution worse.
**End of proof.**

**Lemma 4.** At a query vertex of depth `D`,
`tr[(D & 1) ^ 1]` contains exactly the centers of all relevant blocks.

**Proof.** The relevant centers have depths `D-1, D-3, ...`, all opposite in
parity to `D`. DFS inserts a vertex after answering its own queries and removes
it after finishing its subtree, so at query time the trees contain precisely
the proper ancestors on the current path. Selecting the opposite-parity tree
therefore selects exactly the block centers. **End of proof.**

**Theorem.** The algorithm returns `f(v, p)` for every query.

**Proof.** The recurrence for `dp[v]` sums the no-coin expected cost of every
two-level block, together with the possible final forced step. By Lemma 4, the
chosen segment tree contains exactly those blocks. Its `sumKmax` operation
selects their largest `x` values, which is optimal by Lemma 3, and the algorithm
subtracts the corresponding `2*x` savings given by Lemma 2. Thus the resulting
value is the minimum possible expected number of moves. **End of proof.**

#### Complexity

Sorting the values for coordinate compression costs `O(n log n)`. Every
non-root vertex is inserted into and removed from one segment tree once, and
each query performs one largest-`k` sum operation. Each such operation costs
`O(log n)`, so the total complexity per test case is:

```text
O((n + q) log n) time
O(n + q) memory
```
