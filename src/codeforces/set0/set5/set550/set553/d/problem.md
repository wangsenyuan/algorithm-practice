Nudist Beach is planning a military operation to attack the Life Fibers. In this operation, they will attack and capture several cities which are currently under the control of the Life Fibers.

There are n cities, labeled from 1 to n, and m bidirectional roads between them. Currently, there are Life Fibers in every city. In addition, there are k cities that are fortresses of the Life Fibers that cannot be captured under any circumstances. So, the Nudist Beach can capture an arbitrary non-empty subset of cities with no fortresses.

After the operation, Nudist Beach will have to defend the captured cities from counterattack. If they capture a city and it is connected to many Life Fiber controlled cities, it will be easily defeated. So, Nudist Beach would like to capture a set of cities such that for each captured city the ratio of Nudist Beach controlled neighbors among all neighbors of that city is as high as possible.

More formally, they would like to capture a non-empty set of cities S with no fortresses of Life Fibers. The strength of a city  is defined as (number of neighbors of x in S) / (total number of neighbors of x). Here, two cities are called neighbors if they are connnected with a road. The goal is to maximize the strength of the weakest city in S.

Given a description of the graph, and the cities with fortresses, find a non-empty subset that maximizes the strength of the weakest city.

Input
The first line of input contains three integers n, m, k (2  ≤  n  ≤ 100 000, 1 ≤ m ≤ 100 000, 1 ≤ k ≤ n - 1).

The second line of input contains k integers, representing the cities with fortresses. These cities will all be distinct.

The next m lines contain the roads. The i-th of these lines will have 2 integers ai, bi (1 ≤ ai, bi ≤ n, ai ≠ bi). Every city will have at least one road adjacent to it.

There is no more than one road between each pair of the cities.

Output
The first line should contain an integer r, denoting the size of an optimum set (1 ≤ r ≤ n - k).

The second line should contain r integers, denoting the cities in the set. Cities may follow in an arbitrary order. This line should not contain any of the cities with fortresses.

If there are multiple possible answers, print any of them.

Examples
InputCopy
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
OutputCopy
3
1 4 5
InputCopy
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
OutputCopy
8
1 5 4 8 10 6 3 7
Note
The first example case achieves a strength of 1/2. No other subset is strictly better.

The second example case achieves a strength of 1. Note that the subset doesn't necessarily have to be connected.


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

