# D — Nudist Beach operation

## Problem description

Nudist Beach is planning a military operation against the Life Fibers. They will attack and capture several cities that are currently under Life Fiber control.

There are **n** cities labeled **1 … n**, and **m** bidirectional roads. Life Fibers occupy every city. Among them, **k** cities are fortresses and **cannot** be captured. Nudist Beach may capture any **non-empty** subset of cities that **does not** include a fortress.

After the operation, captured cities must be defended. If a captured city has many neighbors still under Life Fiber control, it is vulnerable. So Nudist Beach wants a set **S** of captured cities such that, for every **x ∈ S**, the fraction of neighbors of **x** that also lie in **S** is as large as possible in the worst case.

More formally: choose a non-empty set **S** with no fortresses. For each **x ∈ S**, define the **strength** of **x** as

**(number of neighbors of x in S) / (total number of neighbors of x)**

(two cities are neighbors if they are connected by a road). Maximize the **minimum** strength over all **x ∈ S** (the “weakest” captured city).

Given the graph and the fortress cities, find any non-empty **S** that achieves this optimum.

## Input

- The first line contains three integers **n**, **m**, **k** (2 ≤ n ≤ 10⁵, 1 ≤ m ≤ 10⁵, 1 ≤ k ≤ n − 1).
- The second line contains **k** distinct integers — the fortress cities.
- Each of the next **m** lines contains two integers **aᵢ**, **bᵢ** (1 ≤ aᵢ, bᵢ ≤ n, aᵢ ≠ bᵢ) — a road. There is at most one road between any pair of cities. Every city has at least one incident road.

## Output

- The first line: an integer **r** — the size of an optimal set (1 ≤ r ≤ n − k).
- The second line: **r** distinct integers — the cities in **S**, in any order. None of them may be a fortress.

If several answers exist, print any.

## Examples

### Example 1

**Input:**

```
9 8 4
3 9 6 8
1 2
1 3
1 4
1 5
2 6
2 7
2 8
2 9
```

**Output:**

```
3
1 4 5
```

### Example 2

**Input:**

```
10 8 2
2 9
1 3
2 9
4 5
5 6
6 7
7 8
8 10
10 4
```

**Output:**

```
8
1 5 4 8 10 6 3 7
```

## Note

In the first example, the achieved minimum strength is **1/2**; no other feasible subset is strictly better.

In the second example, the minimum strength is **1**. The chosen set does **not** need to be connected.

## Summary

- **Model:** Undirected simple graph on **n** nodes, **m** edges; **k** forbidden nodes (fortresses) that cannot appear in the answer.
- **Choice:** Pick any **non-empty** set **S** of nodes with no fortresses (**S** need not be connected).
- **Score:** For each **x ∈ S**, strength(**x**) = (# neighbors of **x** in **S**) / (degree of **x**). The objective value is **min** over **x ∈ S** of strength(**x**).
- **Goal:** Maximize that minimum (a max–min / “bottleneck ratio” problem on the graph).
- **Output:** Print **|S|** and then the list of nodes in one optimal **S** (any optimal solution is allowed).

## Solution Summary

The key idea is to think in terms of a target ratio.

For a fixed value `x`, ask:

- is there a non-empty set `S` of non-fortress cities such that every `u in S` satisfies
  `inside(u) / deg(u) >= x`?

Here:

- `deg(u)` is the original degree of `u`,
- `inside(u)` is the number of neighbors of `u` that also stay inside `S`.

### 1. Safe deletions for a fixed ratio

Suppose we are testing some threshold `x`.

If a non-fortress vertex currently has:

- `inside(u) / deg(u) < x`,

then this vertex can never belong to any valid answer for this `x`, because its ratio is already too small.

So for a fixed `x` we can repeatedly delete every such vertex.

This is exactly like a peeling process:

1. fortresses are forbidden from the start, so remove them conceptually,
2. among the remaining vertices, if some vertex has too few surviving neighbors relative to its original degree, delete it too,
3. continue until no more deletions are possible.

If anything remains, then the remaining set is feasible for `x`.

### 2. Why the best answer appears during peeling

Now forget the fixed `x` for a moment and run the peeling process in the natural greedy order:

- always remove the vertex with the smallest current ratio
  `inside(u) / deg(u)`.

At any moment, the vertices that are still alive form some candidate set `S`.
The minimum ratio inside this set is exactly the ratio of the current heap minimum.

So while we peel vertices one by one, we are implicitly scanning all critical threshold values that can matter.

The best achievable answer is therefore:

- the maximum value of the current minimum ratio over the whole peeling process.

This is what the first pass in `solve` computes:

- `d1` = current number of alive non-fortress neighbors,
- `d2` = original degree,
- the heap is ordered by `d1 / d2`,
- whenever the current minimum ratio improves, we record it.

### 3. Why the second pass reconstructs the optimal set

After we know the optimal ratio `best = d1 / d2`, we run the same peeling process again, but this time we stop exactly when the current minimum ratio reaches that threshold.

At that moment:

- every remaining vertex has ratio at least `best`,
- and by maximality of `best`, this remaining set is optimal.

So the heap contents at the stop moment are exactly one valid optimal answer.

That is what the second call to `play` does:

- it peels vertices while `current_ratio < best`,
- then stops and returns all still-alive non-fortress vertices.

### 4. How the degrees are maintained

For each non-fortress vertex:

- `d2` never changes and is just the original degree,
- `d1` starts as the number of non-fortress neighbors,
- whenever a non-fortress vertex is removed, each alive non-fortress neighbor loses one internal neighbor, so its `d1` decreases by `1`.

This is why a priority queue works well here:

- the heap always exposes the weakest currently alive vertex,
- and each edge causes at most one decrease-key when one endpoint is removed.

### 5. Complexity

Each vertex is inserted once and removed at most once.
Each edge is processed when one of its endpoints is deleted.

So the complexity is:

- `O((n + m) log n)`

which fits the constraints `n, m <= 1e5`.

