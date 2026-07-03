# B. Sereja and Tree

[Problem link](https://codeforces.com/problemset/problem/380/B)

**Contest:** [Codeforces Round #228 (Div. 2)](https://codeforces.com/contest/380)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Sereja builds a special binary tree with `n` levels. Each vertex is `(level, position)` where
`level` is `1..n` and `position` is `1..cnt[level]`. The root is `(1, 1)`.

Tree construction (pseudo code):

```text
cnt[1] = 1
fill left[][], right[][] with -1
for level = 1 .. n-1:
    cnt[level + 1] = 0
    for position = 1 .. cnt[level]:
        if position is a power of two:  // 1, 2, 4, 8, ...
            left[level][position]  = cnt[level + 1] + 1
            right[level][position] = cnt[level + 1] + 2
            cnt[level + 1] += 2
        else:
            right[level][position] = cnt[level + 1] + 1
            cnt[level + 1] += 1
```

After construction, `left[level][position]` / `right[level][position]` is the child index on
level `level + 1`, or `-1` if missing.

Each vertex `(level, position)` has a set `A(level, position)` (distinct integers, like C++
`std::set`). Process `m` operations:

- `1 t l r x` — for every vertex `(t, position)` with `l <= position <= r`, insert `x` into its set
- `2 t v` — for vertex `(t, v)`, take the union of sets over all vertices in its subtree; print
  the union size

## Input

The first line contains `n` and `m` (`1 <= n, m <= 7000`).

Each of the next `m` lines is an operation:

- `1 t l r x` (`1 <= t <= n`, `1 <= l <= r <= cnt[t]`, `1 <= x <= 10^6`)
- `2 t v` (`1 <= t <= n`, `1 <= v <= cnt[t]`)

## Output

For each operation of type 2, print the answer on its own line.

## Example

### Input

```text
4 5
1 4 4 7 1
1 3 1 2 2
2 1 1
2 4 1
2 3 3
```

### Output

```text
2
0
1
```

## Solution

There are at most `7000` operations, so it is acceptable to compare every update with every query.
The hard part is testing whether an update range on one level intersects the subtree of a query
vertex without explicitly building the tree. The full tree can contain hundreds of millions of
vertices when `n = 7000`, so materializing vertices, edges, or Euler tours is not viable.

### Key observation

For a fixed level, the descendants of any vertex form a contiguous interval of positions. Therefore,
an update

```text
1 t l r x
```

contributes to a later query

```text
2 tq v
```

iff `tq <= t` and the query vertex `(tq, v)` is an ancestor of at least one vertex in the updated
range `[l, r]` on level `t`.

Instead of expanding the query vertex downward, the solution lifts the update interval upward. If
the update covers positions `[l, r]` on level `t`, then on the previous level it covers exactly the
parents of these positions:

```text
[ parent(l), parent(r) ]
```

This remains a contiguous interval because positions preserve left-to-right order between adjacent
levels. Repeating this step gives, for every upper level, the interval of vertices whose subtrees
intersect the original update range.

Thus for one update we compute arrays `L[level]` and `R[level]`:

```text
L[t] = l, R[t] = r
L[level-1] = parent(L[level])
R[level-1] = parent(R[level])
```

Then a query `(tq, v)` is affected by this update exactly when:

```text
tq <= t and L[tq] <= v <= R[tq]
```

### Computing `parent(pos)`

Suppose a vertex is at position `p` on some level. Every vertex contributes one right child, and
every power-of-two position contributes one extra left child before its right child.

The number of powers of two in `1..p` is:

```text
bits.Len(p)
```

So the right child of `p` appears at next-level position:

```text
p + bits.Len(p)
```

If `p` is a power of two, its left child appears at:

```text
p + bits.Len(p) - 1
```

To invert this mapping, `parent(pos)` searches for the smallest `p` such that:

```text
p + bits.Len(p) >= pos
```

The implementation starts from the close estimate `pos - bits.Len(pos)`, moves right until the
candidate covers `pos`, then moves left while the previous candidate also covers `pos`. This gives
the unique parent position.

Example:

```text
p = 1  -> children 1, 2
p = 2  -> children 3, 4
p = 3  -> child       5
p = 4  -> children 6, 7
```

So `parent(1)=1`, `parent(2)=1`, `parent(3)=2`, `parent(4)=2`, `parent(5)=3`, and so on.

### Processing operations

First split the operations into updates and queries, preserving their original operation index
`time`. Also compress all values `x`, because only values that appear in update operations can ever
be counted.

For every update:

1. Build its ancestor intervals `L` and `R` on all levels from its level up to the root.
2. Scan all queries.
3. Ignore queries that happened before the update, and queries below the update level.
4. If the query vertex lies inside the lifted interval on its level, this update contributes value
   `x` to that query.

Each query asks for the size of a union of sets, so the same value `x` may be inserted many times
but must be counted once. The solution keeps a compressed bitset `seen[query]`; when an update with
compressed value id `k` reaches a query, it increments the answer only if bit `k` was not set yet.

### Correctness

For any update range `[l, r]` on level `t`, the set of its ancestors on any upper level is
contiguous. This follows from the tree construction preserving the left-to-right order of children:
all descendants of a vertex occupy a continuous block, so the parents of a continuous block also
form a continuous block. Therefore lifting `[l, r]` by repeatedly applying `parent` to both ends
produces exactly the vertices whose subtrees intersect the original update range.

A later query `(tq, v)` should include value `x` exactly when some updated vertex lies inside the
subtree of `(tq, v)`. By the lifted-interval property, that is equivalent to `(tq, v)` being one of
the ancestors of the update range, i.e. `L[tq] <= v <= R[tq]`. The algorithm uses precisely this
condition, so it detects exactly the updates that affect the query.

Finally, the bitset for each query guarantees set semantics: the first affecting update with value
`x` increments the answer, while later affecting updates with the same value only find the bit
already set. Hence the final answer is the number of distinct values in the union over the query
subtree.

### Complexity

Let `m <= 7000`. For each update, lifting the range costs `O(n)` in the worst case, and comparing it
with all queries costs `O(m)`. Since `n <= 7000`, the total time is `O(m * (n + m))`, which is fine
for the constraints. The bitsets store one bit per distinct updated value for every query, using
`O(q * d / 64)` memory where `q` is the number of queries and `d` is the number of distinct update
values.
