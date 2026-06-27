# D. Root was Built by Love, Broken by Destiny

[Problem link](https://codeforces.com/problemset/problem/2127/D)

**Contest:** [Atto Round 1 (Codeforces Round 1041, Div. 1 + Div. 2)](https://codeforces.com/contest/2127)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Heartfall River runs horizontally through Destinyland and divides it into northern and southern sides.

Engineer Root wants to build `n` houses along the river, numbered `1` to `n`. All houses on the northern side and all houses on the southern side must lie on straight lines parallel to the river.

There are `m` bridges; the `i`-th bridge connects houses `u_i` and `v_i` (`u_i != v_i`). All `n` houses are connected by these bridges, and no two bridges connect the same pair of houses.

Count the number of ways to arrange the `n` houses along the river, modulo `10^9 + 7`, such that:

- every bridge connects houses on opposite sides of the river;
- when each bridge is drawn as a straight segment between its two houses, no two bridges cross.

Two arrangements are different if at least one house is on a different side, or if on some side the relative order of two houses is reversed.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases.

Each test case contains:

- one line with two integers `n` and `m` (`2 <= n <= 2 * 10^5`, `n - 1 <= m <= min(n(n - 1) / 2, 2 * 10^5)`);
- `m` lines, each with two integers `u_i` and `v_i` (`1 <= u_i, v_i <= n`, `u_i != v_i`).

The graph is connected and simple. The sum of `n` and the sum of `m` over all test cases are each at most `2 * 10^5`.

## Output

For each test case, print one integer — the number of valid arrangements modulo `10^9 + 7`.

## Example

### Input

```text
4
2 1
1 2
3 3
1 2
1 3
2 3
5 4
1 2
1 3
3 4
3 5
4 3
1 2
1 3
1 4
```

### Output

```text
2
0
8
12
```

### Note

- In the first test case, house 1 is north and house 2 is south, or the opposite.
- In the second test case, every pair of houses is connected, so some bridge would connect two houses on the same side; the answer is `0`.
- In the third test case, one valid arrangement is shown in the statement figure.

## Solution

Call a vertex a **leaf** if its degree in the original graph is one. Remove all such leaves once, and call the remaining vertices the **core**. The key property of a graph drawable with straight non-crossing bridges between two parallel lines is:

> The core must be a single path. A core vertex may additionally have any number of original leaves attached to it.

The core is allowed to have one vertex, as in a star. It cannot be empty except for the graph consisting of one edge.

### Why the core is a path

Look at the non-leaf vertices in their order near the river. A non-leaf vertex can continue the bridge structure toward at most one non-leaf vertex on each side. If it had three non-leaf neighbors, two bridge segments would have to interleave and cross. Hence every core degree is at most two.

The graph is connected. After removing original leaves, all remaining vertices must still be connected: otherwise the original graph would need a bridge between the separated parts, and its endpoint would also be in the core. Thus the core is connected and has maximum degree two, so it is a path. Any violation means the answer is zero.

The implementation builds the graph induced by non-leaves, checks that each degree is at most two, and walks from an endpoint to verify that it visits every core vertex. This also rejects a cycle, where every induced degree is two.

### Count arrangements of a valid core

Suppose a core vertex `u` has `rem[u]` original leaf neighbors. Once the core path is fixed, these leaves occupy the available positions belonging to `u`. They do not create crossings with one another, so their labels may be permuted freely:

```text
rem[u]!
```

Therefore the contribution of all leaves is:

```text
product(rem[u]!) over core vertices u
```

There are also choices for the core itself:

- If the core has at least two vertices, choose which side contains one endpoint (`2` choices), and choose which endpoint is first along the river (`2` choices). This contributes `4`.
- If the core has exactly one vertex, only its side is chosen, contributing `2`.
- If every vertex is a leaf, connectivity forces `n = 2`; the one bridge has exactly `2` side assignments.

Thus, for a non-empty core, the answer is:

```text
(4 if core length >= 2, otherwise 2) * product(rem[u]!)
```

All factorials up to `2 * 10^5` are precomputed modulo `10^9 + 7`.

### Correctness

If the algorithm returns zero, the induced core is either disconnected, has a vertex of degree greater than two, or is a cycle. None of these can be arranged without crossings, so zero is correct.

Otherwise the core is a path. Choosing its direction and one side determines the alternating placement of every core vertex. Each original leaf can then be placed in the positions adjacent to its core parent; permuting leaves of the same parent never introduces a crossing. This constructs every counted arrangement.

Conversely, every valid arrangement has the same path-shaped core, one of the counted core orientations, and one permutation of the leaves attached to each core vertex. The choices are independent and unique, so the product formula counts every valid arrangement exactly once.

### Complexity

For each test case, graph construction, checking the core, and the path walk take `O(n + m)` time. The extra memory usage is `O(n + m)`. Factorials are precomputed once in `O(2 * 10^5)` time.
