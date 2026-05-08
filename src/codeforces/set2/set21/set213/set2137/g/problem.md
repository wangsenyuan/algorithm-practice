There is a directed acyclic graph with `n` nodes and `m` edges. Each node is initially colored blue.

## Fun graph game

A token starts on node `s`. Cry and River take turns moving the token along a directed edge to an adjacent node, with Cry going first.

- **Cry wins** if the token ever reaches a node with no outgoing edges (after either player's turn).
- **River wins** if the token reaches a red node (after either player's turn).
- If the token reaches a node that is both red and has no outgoing edges, **River wins**.

Since the graph is acyclic, the game always ends in a finite number of turns.

Cry or River can win immediately if the starting node `s` has no outgoing edges, or is red, respectively.

## Queries

You must handle `q` queries:

- `1 u`: update the color of node `u` to red. It is guaranteed that node `u` was blue before this update.
- `2 u`: if a fun graph game is played with the token initially at node `u`, and both players play optimally, does Cry win?

## Input

Each test contains multiple test cases.

- The first line contains the number of test cases `t` (`1 <= t <= 10^4`). The description of the test cases follows.

For each test case:

- The first line contains three integers `n`, `m`, `q` (`2 <= n <= 2 * 10^5`, `1 <= m, q <= 2 * 10^5`).
- The next `m` lines each contain two integers `u` and `v` (`1 <= u, v <= n`), meaning there is a directed edge from `u` to `v`.
- The next `q` lines each contain two integers `x` and `u` (`1 <= x <= 2`, `1 <= u <= n`) — the query type and the node.

It is guaranteed that the graph is a directed acyclic graph, and that no edge appears more than once.

It is guaranteed that the sum of `n`, the sum of `m`, and the sum of `q` over all test cases each do not exceed `2 * 10^5`.

## Output

For each query of the second type, output `YES` if Cry wins; otherwise output `NO`. Letters may be in uppercase or lowercase.

## Example

### Input

```text
1
7 8 10
1 2
1 3
1 4
2 5
3 6
5 7
2 3
3 4
2 1
1 3
1 4
2 1
2 2
2 3
2 4
2 5
2 6
2 7
```

### Output

```text
YES
NO
YES
NO
NO
YES
YES
YES
YES
```

## Note

Below is the graph from the sample.

Initially, all nodes are blue. Cry cannot lose, and eventually the token is moved to a node without outgoing edges.

After nodes `3` and `4` are painted red, nodes `1`, `3`, and `4` are a win for River under optimal play. If the game starts at nodes `3` or `4`, River wins immediately. If the game starts at node `1`, one possible play is:

- Cry moves the token to node `2`.
- River moves the token to node `3`, which is red, so River wins.

It can be shown that Cry still wins with optimal play for all other starting nodes.

### ideas
1. River获胜的条件是，进入一个Red节点，
2. 如果River不能获胜，那么是Cry获胜。所以问题的关键就是看从一个节点u出发，无论Cry怎么选择，都会进入一个Red节点。
3. 这是一个DAG, 
4. 那么对于u，什么情况下，肯定会走到red节点呢？
5. 如果一个节点u已经是一个失败节点（river胜出），那么它就一直是一个失败节点
6. 但是如果是一个胜利节点，后面有可能还是会变成失败节点
7. 那么如果从u出发，只能到达失败节点，那么u也是失败节点。或者到达v节点后，可以到达失败节点，（v不一定是失败节点）但是u肯定是失败节点
8. 想不出来～

## Detailed solution

The graph is a DAG, so every game ends. Since River wins immediately when the token reaches a red node, and Cry wins immediately when the token reaches a sink, every state is either winning for Cry or winning for River.

The updates only paint blue nodes red. Red nodes never become blue again. So as time goes on, the game can only become worse for Cry. This monotonicity is the key: a state can change from "Cry wins" to "Cry loses", but never in the other direction.

## Two DP states

For each node `u`, maintain two booleans:

```text
dp1[u] = whether Cry wins if the token is at u and it is Cry's turn
dp2[u] = whether Cry wins if the token is at u and it is River's turn
```

The query `2 u` asks about a game starting at `u` with Cry moving first, so the answer is exactly:

```text
dp1[u]
```

Initially every node is blue. In that initial graph, Cry can always win, because the game must eventually reach a sink and there are no red nodes. Therefore:

```text
dp1[u] = true
dp2[u] = true
```

for every node `u`.

When nodes become red, some of these states turn false.

## Game transitions

Let there be an edge:

```text
v -> u
```

Think about how losing states propagate backwards through this edge.

### If `dp1[u]` becomes false

This means:

```text
if the token is at u and it is Cry's turn, Cry loses.
```

Now consider predecessor `v` when it is River's turn. River can move from `v` to `u`. After that move, the token is at `u` and it is Cry's turn, which is losing for Cry.

So River has a move that makes Cry lose. Therefore:

```text
dp2[v] = false
```

This is why the code handles a popped state `(u, 1)` by setting all reverse-neighbors' `dp2` to false:

```go
if w == 1 {
	for _, v := range adj[u] {
		if dp2[v] {
			dp2[v] = false
			que[head] = data{v, 2}
			head++
		}
	}
}
```

Here `adj[u]` stores predecessors of `u`, because the input edges were reversed when building `adj`.

### If `dp2[u]` becomes false

This means:

```text
if the token is at u and it is River's turn, Cry loses.
```

Now consider predecessor `v` when it is Cry's turn. Cry chooses one outgoing edge from `v`.

Cry wins from `v` only if he has at least one move to a child where Cry still wins on River's turn. In other words:

```text
dp1[v] = true if there exists child u with dp2[u] = true
```

So `dp1[v]` becomes false only when every outgoing child of `v` has `dp2[child] = false`.

The code maintains:

```text
cnt[v] = number of outgoing children child of v with dp2[child] = false
deg[v] = total outdegree of v
```

When a child `u` changes to `dp2[u] = false`, all predecessors `v` increment `cnt[v]`. If now:

```text
cnt[v] == deg[v]
```

then every move from `v` goes to a losing state for Cry, so:

```text
dp1[v] = false
```

This is the second propagation rule in the code:

```go
if w == 2 {
	for _, v := range adj[u] {
		cnt[v]++
		if cnt[v] == deg[v] && dp1[v] {
			dp1[v] = false
			que[head] = data{v, 1}
			head++
		}
	}
}
```

## Painting a node red

When query `1 s` paints node `s` red, River wins immediately if the token is at `s`, regardless of whose turn it is. Therefore both states become losing for Cry:

```text
dp1[s] = false
dp2[s] = false
```

Then the two reverse propagation rules above may make more states losing.

The code performs this with a queue:

```go
que := make([]data, n*2)
```

Each queue item is:

```go
type data struct {
	id    int
	state int
}
```

where:

- `state = 1` means `dp1[id]` just became false;
- `state = 2` means `dp2[id]` just became false.

The `flip(s)` function starts by enqueueing the states of `s` that newly became false, then processes the queue until no more states change.

## Why the guard in `flip` is necessary

This part is subtle.

Even if node `s` was still blue, one or both of its DP states may already be false because of previously painted red nodes reachable from it.

For example, a blue node can already be losing for Cry if River can force the token into an earlier red node.

So when `s` is painted red, we must not blindly enqueue both states. We should enqueue only the states that actually change from true to false:

```go
if dp1[s] {
	dp1[s] = false
	que[head] = data{s, 1}
	head++
}
if dp2[s] {
	dp2[s] = false
	que[head] = data{s, 2}
	head++
}
```

This is required for correctness, not only for speed.

If `dp2[u]` is propagated twice, then every predecessor `v` increments `cnt[v]` twice for the same child `u`. Then `cnt[v] == deg[v]` may become true even though not all children of `v` are losing. That would incorrectly set `dp1[v] = false`.

So the invariant is:

```text
each state dp1[u] / dp2[u] is enqueued exactly once, when it changes true -> false
```

With that invariant, `cnt[v]` really counts distinct outgoing children whose `dp2` state is false.

## Why this is efficient

Every state can change only once:

```text
dp1[u]: true -> false
dp2[u]: true -> false
```

There are only `2n` such changes.

When a state changes, the algorithm scans the reverse adjacency list of that node. Across the whole test case:

- each edge is scanned at most once due to a `dp1` change;
- each edge is scanned at most once due to a `dp2` change.

Therefore the total propagation cost is:

```text
O(n + m)
```

Handling all queries adds `O(q)`, so the total complexity is:

```text
O(n + m + q)
```

The memory usage is:

```text
O(n + m)
```
