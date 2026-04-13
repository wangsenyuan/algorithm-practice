# E. Dividing Kingdom

A country called Flatland is an infinite two-dimensional plane. Flatland has `n` cities; each city is a point on the plane.

Flatland is ruled by king Circle IV. Circle IV has 9 sons. He wants to give each of his sons part of Flatland to rule. For that, he wants to draw **four distinct straight lines**, such that **two** of them are parallel to the `Ox` axis, and **two** others are parallel to the `Oy` axis. No straight line may pass through any city. Thus, Flatland is divided into **9** parts, and each son is given exactly one of these parts. Circle IV decided that the `i`-th son should get the part of Flatland that has exactly `a_i` cities.

Help Circle find such four straight lines so that, after dividing Flatland into 9 parts by these lines, the resulting parts can be assigned to the sons so that son `i` receives a part that contains exactly `a_i` cities.

## Input

The first line contains an integer `n` (`9 ≤ n ≤ 10^5`) — the number of cities in Flatland.

The next `n` lines each contain two space-separated integers `x_i`, `y_i` (`-10^9 ≤ x_i, y_i ≤ 10^9`) — the coordinates of the `i`-th city. No two cities coincide.

The last line contains **nine** space-separated integers `a_1, a_2, …, a_9`.

## Output

If there is no solution, print a single integer `-1`.

Otherwise, print on the **first** line two distinct real numbers: `x_1`, `x_2` — the abscissas of the lines parallel to the `Oy` axis. On the **second** line print two distinct real numbers: `y_1`, `y_2` — the ordinates of the lines parallel to the `Ox` axis. If there are multiple solutions, print any.

When the answer is checked, a city is considered to lie on a line if the distance from the city to that line does not exceed `10^-6`. Two lines are considered the same if the distance between them does not exceed `10^-6`.

## Examples

### Example 1

**Input**

```
9
1 1
1 2
1 3
2 1
2 2
2 3
3 1
3 2
3 3
1 1 1 1 1 1 1 1 1
```

**Output**

```
1.5000000000 2.5000000000
1.5000000000 2.5000000000
```

### Example 2

**Input**

```
15
4 4
-1 -3
1 5
3 -4
-4 4
-1 1
3 -3
-4 -5
-3 3
3 2
4 1
-4 2
-2 -5
-3 4
-1 4
2 1 2 1 2 1 3 2 1
```

**Output**

```
-3.5000000000 2.0000000000
3.5000000000 -1.0000000000
```

### Example 3

**Input**

```
10
-2 10
6 0
-16 -6
-4 13
-4 -2
-17 -10
9 15
18 16
-5 2
10 -5
2 1 1 1 1 1 1 1 1
```

**Output**

```
-1
```

## Note

There is no solution for the third sample test.


## Thoughts 1: Persistent Tree

The key observation is that the exact labels of the 9 regions do not matter at first; only the multiset of city counts matters.

So we can:

1. decide which 3 sons belong to the left column,
2. which 3 belong to the middle column,
3. and the remaining 3 belong to the right column.

That is only:

- `C(9, 3)` choices for the left column,
- `C(6, 3)` choices for the middle column.

So the vertical grouping is small enough to brute force.

After fixing the 3 columns, each column still has 3 cells, and inside each column we only need to decide the order from bottom to top. That gives `3!` permutations per column, which is also small.

Then the real problem becomes:

- for a chosen vertical split,
- and for a chosen horizontal order inside each column,
- can we quickly count how many points lie in each prefix rectangle?

The persistent-tree solution precomputes all prefixes by `x`:

- sort points by `x`,
- for each prefix `1..k`, build a persistent segment tree over `y`,
- then for any candidate horizontal cut we can query:
  - among the first `k` points by `x`,
  - how many have `y <= bound`.

This makes checking a candidate layout fast:

- first prefix = left column,
- second prefix = left + middle columns,
- query two horizontal boundaries,
- compare the resulting rectangle counts with the chosen permutation of `a`.

So the persistent tree is just a data structure for answering many prefix-by-`x`, threshold-by-`y` counting queries.

## Thoughts 2: Simpler Fenwick Sweep

The persistent tree is not strictly necessary.

What we really need is still the same query:

- among the first `k` points in `x` order,
- count how many have `y` rank at most `r`.

Instead of storing all prefixes persistently, we can collect all needed queries offline:

1. enumerate the same vertical column partitions,
2. enumerate the same `3! * 3! * 3!` row orders,
3. for each candidate layout, record the four needed prefix queries:
   - left column with first horizontal cut,
   - left column with second horizontal cut,
   - left+middle columns with first horizontal cut,
   - left+middle columns with second horizontal cut.

Now sort all these queries by `k`, sweep points from left to right in `x` order, and maintain one Fenwick tree over compressed `y`:

- when the sweep reaches prefix `k`, add the first `k` points into the Fenwick tree,
- answer every query with this `k` using a Fenwick prefix sum on `y`.

This is simpler conceptually:

- no persistence,
- one BIT,
- one offline sweep.

The search over column partitions and row permutations stays exactly the same; only the counting backend changes.

So both solutions share the same combinational idea:

- brute force the assignment of the 9 target counts into a `3 x 3` grid shape,
- and use a data structure to verify each candidate efficiently.

The difference is only the query engine:

- `solve1`: persistent segment tree over all `x` prefixes,
- `solve`: offline Fenwick tree sweep over collected queries.
