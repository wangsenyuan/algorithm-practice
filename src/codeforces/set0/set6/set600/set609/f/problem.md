# F. Frogs and mosquitoes

[Problem link](https://codeforces.com/problemset/problem/609/F)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

There are `n` frogs sitting on the coordinate axis `Ox`. For each frog two
values `x_i`, `t_i` are known: the position and the initial length of the tongue
of the `i`-th frog. It is guaranteed that all positions `x_i` are different.

There are `m` mosquitoes landing on the coordinate axis one by one. For each
mosquito two values are known: `p_j`, the coordinate where the `j`-th mosquito
lands, and `b_j`, the size of the `j`-th mosquito.

Frogs and mosquitoes are represented as points on the coordinate axis.

A frog can eat a mosquito if the mosquito is at the same position as the frog or
to its right, and the distance between them is not greater than the length of
the frog's tongue.

If at some moment several frogs can eat a mosquito, the leftmost frog, the one
with minimum `x_i`, eats it. After eating a mosquito, the length of that frog's
tongue increases by the size of the eaten mosquito. It is possible that after
this, the frog becomes able to eat some other mosquitoes; in that case, it must
eat them.

For each frog, print two values: the number of eaten mosquitoes and the length
of its tongue after all mosquitoes have landed and all possible mosquitoes have
been eaten.

Each mosquito lands only after frogs have eaten all possible mosquitoes that
landed before it. Mosquitoes are given in landing order.

## Input

The first line contains two integers `n`, `m`
(`1 <= n, m <= 2 * 10^5`) -- the number of frogs and mosquitoes.

Each of the next `n` lines contains two integers `x_i`, `t_i`
(`0 <= x_i, t_i <= 10^9`) -- the position and the initial tongue length of the
`i`-th frog. It is guaranteed that all `x_i` are different.

The next `m` lines contain two integers each: `p_j`, `b_j`
(`0 <= p_j, b_j <= 10^9`) -- the position and size of the `j`-th mosquito.

## Output

Print `n` lines. The `i`-th line should contain two integer values `c_i`,
`l_i`: the number of mosquitoes eaten by the `i`-th frog and the final length
of the tongue of the `i`-th frog.

## Examples

### Input 1

```text
4 6
10 2
15 0
6 1
0 1
110 10
1 1
6 0
15 10
14 100
12 2
```

### Output 1

```text
3 114
1 10
1 1
1 2
```

### Input 2

```text
1 2
10 2
20 2
12 1
```

### Output 2

```text
1 3
```


## Solution

Process mosquitoes in landing order. A mosquito that cannot be eaten immediately
stays on the line and may be eaten later after some frog's tongue grows.

For a mosquito at position `p`, the frog that eats it is:

```text
the leftmost frog with x <= p and reach >= p
```

where:

```text
reach = x + current tongue length
```

So the main problem is maintaining frog reaches under updates, and answering
leftmost-reachable queries quickly.

### Sort Frogs by Position

Sort frogs by `x`. Keep:

```go
xs[i]    = sorted position of frog i
reach[i] = current rightmost point frog i can reach
```

The original frog index is stored in `arr[i].id`, because output must be in the
input order.

For a mosquito at `p`, only frogs with `x <= p` can eat it. Let:

```go
pos = last index with xs[pos] <= p
```

Among indices `0..pos`, we need the first index whose `reach[i] >= p`.

This is exactly what the segment tree stores: each segment keeps the maximum
`reach` in that range. To find the leftmost frog, recursively search the left
child first if its maximum reach is at least `p`.

```go
id := seg.FirstAtLeast(pos, p)
```

If no such frog exists, this mosquito cannot be eaten now.

### Waiting Mosquitoes

Uneaten mosquitoes are stored in an AVL tree keyed by their position.

For each position, the tree node stores:

```go
cnt = number of waiting mosquitoes at this position
val = total size of those mosquitoes
```

This aggregation is valid because if a frog can reach a position, it will eat
all mosquitoes waiting at that same position, and all their sizes simply add to
the tongue length.

### Eating a New Mosquito

When a new mosquito `(p, b)` lands:

1. Query the segment tree for the leftmost frog that can eat `p`.
2. If none exists, insert `(p, b)` into the waiting AVL tree.
3. Otherwise, that frog eats it:

   ```go
   reach[id] += b
   ans[orig][0]++
   ans[orig][1] += b
   seg.Update(id, reach[id])
   ```

After growing, the same frog may now be able to eat old waiting mosquitoes.
Because this frog only eats mosquitoes in:

```text
xs[id] <= position <= reach[id]
```

the code repeatedly takes the largest waiting position `<= reach[id]`:

```go
tmp := FindPrevNode(wait, reach[id])
```

If `tmp.key < xs[id]`, then every remaining waiting mosquito that is reachable
from the right is still left of the frog, so this frog cannot eat it.

Otherwise, the frog eats all mosquitoes at `tmp.key`, grows by `tmp.val`, and
the process repeats:

```go
ans[orig][0] += tmp.cnt
ans[orig][1] += tmp.val
reach[id] += tmp.val
wait = Delete(wait, tmp.key)
seg.Update(id, reach[id])
```

This loop implements the statement's rule that after a frog eats and grows, it
must immediately eat any newly reachable mosquitoes before the next mosquito
lands.

### Why Not Store Active Ownership Intervals?

An alternative idea is to maintain artificial intervals saying which frog owns
which positions. That is fragile.

A frog's real interval is always:

```text
[x, reach]
```

But an ownership structure may store a trimmed interval:

```text
[activeLeft, reach]
```

where `activeLeft > x` because an earlier frog already covers the left part.

When a left frog grows, it may affect many later frogs, not only the first one:

```text
A.x < B.x < C.x < D.x
```

If `A.reach` grows far enough, then:

- some later frogs may become fully dominated and should be removed;
- some may be partially covered and need their active left boundary moved;
- more frogs after that may also need updates.

Stopping after the first partial overlap leaves stale active intervals, and then
a later frog may incorrectly eat a mosquito that the earlier grown frog should
eat.

The segment tree avoids all of this. It stores only real reaches and directly
answers the actual rule:

```text
find the leftmost frog with x <= p and reach >= p
```

No trimmed ownership intervals are needed.

### Complexity

Each mosquito is inserted into the waiting tree at most once and deleted at most
once. Each frog reach update costs `O(log n)`, and each waiting-tree operation
costs `O(log m)`.

The total complexity is:

```text
O((n + m) log(n + m))
```
