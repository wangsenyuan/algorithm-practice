# Codeforces 2207D - Boxed Like a Fish

https://codeforces.com/problemset/problem/2207/D

## Statement

Let `n` and `k` be positive integers. You are given a tree with `n` vertices
numbered `1..n`.

Cyndaquil is traversing this tree and is trying to reach one of its leaves.
Initially, he starts at a non-leaf vertex `v`. In one turn, he may either stay
still or move from his current vertex along an edge to any neighboring vertex.

Snorlax is trying to stop Cyndaquil by sleeping on an edge. When Snorlax chooses
an edge, Cyndaquil is blocked from traversing that edge until Snorlax moves
again. Only one edge may be blocked at a time, so only the most recently chosen
edge is blocked.

Snorlax is slow and has a cooldown timer, initially `0`. He may choose a new edge
only when the cooldown is `0` or lower, though he does not have to choose one
immediately. When he moves to a new edge, the timer is reset to `k`. After each
of Cyndaquil's turns, even if Cyndaquil stays still, the cooldown timer
decreases by `1`.

They take turns as described, with Snorlax acting first. Initially, Snorlax is
not sleeping on any edge.

Assuming both play optimally, determine whether Cyndaquil can always reach a
leaf after a finite number of turns.

A tree is a connected graph without cycles.

A leaf is a vertex with degree `1`.

## Input

Each test contains multiple test cases.

The first line contains the number of test cases `t` (`1 <= t <= 10^4`).

For each test case:

- The first line contains three integers `n`, `k`, and `v`
  (`3 <= n <= 5 * 10^5`, `1 <= k, v <= n`) — the number of vertices, Snorlax's
  cooldown timer, and Cyndaquil's starting vertex.
- The next `n - 1` lines each contain two integers `a` and `b`
  (`1 <= a, b <= n`, `a != b`), describing an edge between vertices `a` and `b`.

It is guaranteed that:

- the edges form a tree;
- `v` is not a leaf;
- the sum of `n` over all test cases does not exceed `5 * 10^5`.

## Output

For each test case, print `YES` if Cyndaquil can always reach a leaf in a finite
number of turns. Otherwise, print `NO`.

You may print the answer in any letter case.

## Sample

```text
6
6 2 1
1 2
2 3
2 4
1 5
5 6
7 1 4
1 2
2 3
3 4
4 5
5 6
6 7
3 1 3
1 3
2 3
4 1 4
1 3
3 4
4 2
9 3 5
4 5
5 6
4 7
9 8
8 7
1 2
2 3
3 4
9 4 5
4 5
5 6
4 7
9 8
8 7
1 2
2 3
3 4
```

```text
YES
NO
YES
NO
NO
YES
```

## Note

In the first test case, Cyndaquil starts at vertex `1`. If Snorlax blocks the
edge from `1` to `2`, then Cyndaquil can reach vertex `6` after two turns, which
is a leaf. Otherwise, Cyndaquil can advance to vertex `2`, from which Snorlax
cannot stop him from reaching at least one of vertices `3` and `4`.

In the second test case, Snorlax can block Cyndaquil indefinitely from reaching
either leaf `1` or leaf `7`.

## Solution

Root the tree at Cyndaquil's starting vertex `v`.

We classify vertices as **winning** for Cyndaquil in the following sense:

If Cyndaquil can reach such a vertex while Snorlax's cooldown is `0`, then
Cyndaquil can force reaching some leaf in a finite number of turns.

Every leaf is winning, because if Cyndaquil is already at a leaf, he has
succeeded.

For an internal vertex `u`, the question is whether Cyndaquil has enough
different child directions to outrun Snorlax's cooldown.

## Why only two best child directions matter

Suppose Cyndaquil is at vertex `u` and Snorlax's cooldown is `0`.

For each child subtree of `u`, there are two closely related distances:

- `dp[c]`: distance from child `c` to the nearest winning vertex inside `c`'s
  subtree;
- `dp[c] + 1`: distance from `u` to that same winning vertex, because we first
  cross the edge `u -> c`.

In the implementation, the two smallest child values are usually kept before
adding this extra edge:

```text
first  <= second       // these are child dp values
```

In the explanation below, let:

```text
d1 = first + 1
d2 = second + 1
```

These are the two smallest distances measured from `u`.

If there is only one such direction, Snorlax can always block that unique escape
direction when ready, so it cannot make `u` winning by itself.

Now suppose there are two directions. Snorlax can block one edge or wait and
block the final edge toward the first direction Cyndaquil commits to. If
Cyndaquil turns back and tries the second direction, Snorlax needs cooldown time
before he can move the block there.

The important quantity is the time needed to go from the first winning direction
to the second through `u`.

If the nearest winning vertices are at distances `d1` and `d2` from `u`, then
the time pressure is:

```text
d1 + d2 - 1
```

The `-1` appears because Snorlax blocks when Cyndaquil is one edge away from the
first winning point; from that moment, Cyndaquil turns around and starts using
the other direction.

So if:

```text
d1 + d2 - 1 <= k
```

then Snorlax's cooldown is too slow: after blocking one direction, he cannot
recover fast enough to also block the other before Cyndaquil reaches a winning
position. Thus `u` itself is winning.

Equivalently:

```text
d1 + d2 <= k + 1
```

Substituting `d1 = first + 1` and `d2 = second + 1` gives:

```text
first + second + 1 <= k
```

With integer values, this is the same as:

```text
first + second < k
```

This is why an implementation may use `first + second < k`, even though the
distance-from-`u` explanation naturally writes `best1 + best2 - 1 <= k`.

If this condition is false, Snorlax can keep waiting and blocking the currently
threatened direction, forcing Cyndaquil to return before reaching a winning
position.

## Tree DP

Let:

```text
dp[u] = shortest distance from u to a winning vertex inside u's rooted subtree
```

If `u` itself is winning, then:

```text
dp[u] = 0
```

For a leaf:

```text
dp[u] = 0
```

For an internal vertex, compute all child values:

```text
child value from c = dp[c]
```

Let `first` and `second` be the two smallest child values. If both exist and:

```text
first + second < k
```

then `u` is winning, so:

```text
dp[u] = 0
```

If instead you define:

```text
best1 = first + 1
best2 = second + 1
```

as the two smallest distances from `u`, the same condition is:

```text
best1 + best2 - 1 <= k
```

These two forms are equivalent.

Otherwise, `u` is not immediately winning, and the nearest winning vertex in
`u`'s subtree is reached through the best child:

```text
dp[u] = first + 1
```

After the DFS finishes, Cyndaquil can force a win if and only if:

```text
dp[v] = 0
```

## Why this is sufficient

The recurrence captures both players' optimal strategies.

For Cyndaquil:

- If he reaches a winning vertex with cooldown `0`, by definition he can force a
  leaf.
- If a vertex has two sufficiently close winning child directions, Snorlax can
  block at most one of them before cooldown prevents switching quickly enough.

For Snorlax:

- If a vertex has fewer than two winning directions, there is no fork to exploit.
- If the two nearest winning directions are still too far apart, then every pair
  of directions is too far apart. Snorlax can wait, block the direction
  Cyndaquil is about to complete, and recover in time before Cyndaquil can
  convert the detour into another winning position.

Thus the two smallest child distances completely determine whether `u` becomes
winning.

## Complexity

Root the tree once at `v` and run one DFS. Each edge is processed once.

The total complexity is:

```text
O(n)
```

per test case, with `O(n)` memory.


### ideas
1. 如果v已经是leaf， yes
2. Snorlax如果要阻拦，肯定是阻拦Cyndaquil下一个要选择的边（但不一定是最近的那条边）
3. dp[v]表示在子树v中阻止Cyndaquil到达叶子节点，需要的最小的cooldown时间
4. dp[v] = 0 if v is leaf
5. else, 如果u只有一个叶子节点dp[u] = dp[v] + 1 ? 只需要在从v中到达叶子节点前，进行阻拦即可
6. 如果u有多个叶子节点， sort, dp[v]; 如果存在任何两个以上相同的dp[v], dp[u] = dp[v] + 1 (这是因为，Snorlax)只能防治一个
7. 否则 dp[u] = max(dp[v]) + 1 (前面的依次预防都可以做到，只有最后一个才能稳定到达)
8. 
