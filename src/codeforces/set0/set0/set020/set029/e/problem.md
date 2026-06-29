# E. Quarrel

[Problem link](https://codeforces.com/problemset/problem/29/E)

**Contest:** [Codeforces Beta Round 29 (Div. 2, Codeforces format)](https://codeforces.com/contest/29)

time limit per test: 2 seconds

memory limit per test: 64 megabytes

input: standard input

output: standard output

Friends Alex and Bob live in Bertown. There are `n` crossroads in the town, and some pairs of
crossroads are connected by bidirectional roads of equal length.

Bob starts at crossroad `1` and needs to reach crossroad `n`. Alex starts at crossroad `n` and needs
to reach crossroad `1`. After a quarrel, they do not want to appear at the same crossroad at the
same time.

They move simultaneously with equal speeds. Until both reach their destinations, neither of them may
stay still. It is allowed for them to meet in the middle of a road while moving in opposite
directions, but they must never be at the same crossroad at the same time. It is also possible that
one person reaches a crossroad exactly when the other leaves it.

Find two routes with the same number of roads, one for Bob from `1` to `n` and one for Alex from
`n` to `1`, such that they never appear at the same crossroad at the same time. Among all valid
route pairs, minimize this common length. If no such routes exist, print `-1`.

## Input

The first line contains two integers `n` and `m` (`2 <= n <= 500`, `1 <= m <= 10000`) — the number
of crossroads and roads.

Each of the next `m` lines contains two integers `u` and `v` (`1 <= u, v <= n`) — a bidirectional
road between crossroads `u` and `v`.

It is guaranteed that there are no loops and no multiple edges.

## Output

If there is no valid pair of routes, output `-1`.

Otherwise, on the first line output an integer `k` — the minimum common route length.

On the second line output `k + 1` integers: Bob's route from crossroad `1` to crossroad `n`.

On the third line output `k + 1` integers: Alex's route from crossroad `n` to crossroad `1`.

At each time moment `i` (`0 <= i <= k`), the two printed routes must have different vertices.

## Solution

Think about the two people together. At some integer time, Bob is at vertex `u`, and Alex is at
vertex `v`. This pair `(u, v)` is one BFS state.

The start state is:

```text
(1, n)
```

The target state is:

```text
(n, 1)
```

From state `(u1, v1)`, in one second:

- Bob must move to some neighbor `u2` of `u1`;
- Alex must move to some neighbor `v2` of `v1`;
- after moving, they must not stand at the same crossroad, so `u2 != v2`.

Meeting in the middle of an edge is allowed, so we do not forbid the case where Bob moves
`u1 -> v1` and Alex moves `v1 -> u1`. Only equality at integer-time vertices matters.

So the graph of states has at most `n * n` states. A normal BFS on this state graph gives the
minimum route length, and storing the previous state lets us restore both routes.

### Basic BFS transition

The direct transition is:

```text
for each u2 in adj[u1]:
    for each v2 in adj[v1]:
        if u2 != v2 and (u2, v2) is not visited:
            visit (u2, v2)
```

This is easy to understand, but it is too slow on dense graphs. If `u1` and `v1` both have degree
around `140`, one popped state tries almost `140 * 140` pairs. With many states, this becomes too
large.

### Bitset optimization

The optimized code keeps the same BFS, but speeds up the inner loop.

For each vertex `x`, store its neighbors as a bitset:

```text
adjBits[x]
```

If bit `t` is `1`, then `t` is a neighbor of `x`.

For visited states, store another bitset:

```text
vis[u]
```

If bit `v` is `1` in `vis[u]`, then state `(u, v)` has already been visited.

Now suppose we are expanding state `(u1, v1)`, and Bob chooses a next vertex `u2`. The possible
Alex next vertices are all neighbors of `v1`, except:

- vertices where state `(u2, v2)` was already visited;
- vertex `u2` itself, because `u2 == v2` means they meet at a crossroad.

In bitset form:

```text
cand = adjBits[v1] without vis[u2]
cand = cand without bit u2
```

In Go, `&^` means "and not", so:

```go
cand := adjBits[v1][w] &^ vis[u2][w]
```

This line means: among the vertices represented by word `w`, keep Alex's possible next vertices
that are neighbors of `v1` and have not been visited together with Bob at `u2`.

Then:

```go
if u2/64 == w {
    cand &^= 1 << (u2 % 64)
}
```

removes the forbidden vertex `v2 = u2`.

### Small bitset example

Use zero-based vertex ids in this example: `0, 1, 2, 3, 4, 5, 6, 7`.

Suppose we are at state:

```text
(u1, v1) = (0, 7)
```

Bob chooses:

```text
u2 = 3
```

Alex is currently at `v1 = 7`, and assume:

```text
adj[7] = {1, 2, 3, 5}
```

Then the neighbor bitset of vertex `7` is:

```text
vertex:  7 6 5 4 3 2 1 0
bit:     0 0 1 0 1 1 1 0
```

So Alex could move to `1, 2, 3, 5` before applying restrictions.

Now assume that for Bob's next vertex `u2 = 3`, we already visited:

```text
(3, 2), (3, 5)
```

Then:

```text
vis[3] has bits {2, 5}
```

Subtracting visited states:

```text
adjBits[7]      = {1, 2, 3, 5}
vis[3]          = {2, 5}
candidate left  = {1, 3}
```

But `v2 = 3` is forbidden because Bob is also moving to `3`. After removing bit `3`:

```text
final candidates = {1}
```

So for this fixed `u2 = 3`, the optimized loop will only create one next state:

```text
(3, 1)
```

The slow nested loop would still check all four Alex neighbors `1, 2, 3, 5`; the bitset loop removes
the already visited and invalid choices at once.

### Enumerating set bits

After computing `cand`, every `1` bit is one valid `v2`.

```go
for cand > 0 {
    b := bits.TrailingZeros64(cand)
    v2 := w*64 + b
    ...
    cand &= cand - 1
}
```

`bits.TrailingZeros64(cand)` returns the index of the lowest `1` bit. For example:

```text
cand = 00101000
```

The lowest set bit is at position `3`, so `b = 3`.

Then:

```go
cand &= cand - 1
```

removes that lowest `1` bit:

```text
00101000 -> 00100000
```

The loop continues until no candidate bits remain.

### Path reconstruction

Each state is packed as one integer:

```text
id(u, v) = u * n + v
```

For every visited state, store:

```text
dist = BFS distance
prev = previous state id
```

When the target state `(n, 1)` is reached, follow `prev` pointers back to `(1, n)`. Each restored
state `(u, v)` gives one vertex in Bob's route and one vertex in Alex's route. The restored order is
backwards, so reverse both arrays before printing.

### Correctness

Every valid pair of routes corresponds to a walk in the state graph: at time `i`, the state is
`(BobPosition_i, AlexPosition_i)`. The rule `u2 != v2` exactly enforces that they are never at the
same crossroad at the same integer time, and crossing the same edge in opposite directions remains
allowed.

Conversely, every path in the state graph from `(1, n)` to `(n, 1)` gives two routes of equal
length satisfying the problem rules.

BFS explores this unweighted state graph by increasing number of moves, so the first time it reaches
`(n, 1)`, the distance is the minimum possible common route length. The bitset optimization changes
only how outgoing states are enumerated; it enumerates exactly the same unvisited valid states as the
basic nested loop.

### Complexity

There are at most `n^2` states. The optimized transition uses bitsets of `ceil(n / 64)` words.

In practice this avoids the dense-graph `deg(u) * deg(v)` blow-up, because already visited or
invalid Alex positions are filtered word by word before individual candidates are enumerated.

Memory usage is `O(n^2)` for states and visited bitsets.
