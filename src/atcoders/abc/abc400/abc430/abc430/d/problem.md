# D - Neighbor Distance

[Problem link](https://atcoder.jp/contests/abc430/tasks/abc430_d)

**Contest:** [AtCoder Beginner Contest 430](https://atcoder.jp/contests/abc430)

time limit: 4 sec

memory limit: 1024 MiB

score: 400 points

There is a number line, and initially person 0 is standing alone at coordinate 0.

From now on, persons `1, 2, ..., N` will arrive in this order and stand on the number line. Person `i`
stands at coordinate `X_i`. Here, `X_i >= 1`, and `X_i` is distinct for all persons.

Each time a person arrives, answer the following question.

- Suppose that currently `r + 1` persons `0, 1, ..., r` are standing on the number line.
- Define `d_i` as the distance to the nearest other person from person `i`. More formally,

  ```text
  d_i = min_{0 <= j <= r, j != i} |X_i - X_j|
  ```

- Find the sum `sum_{i=0}^r d_i`.

## Constraints

- All input values are integers.
- `1 <= N <= 5 * 10^5`
- `1 <= X_i <= 10^9`
- `X_i != X_j` if `i != j`.

## Input

```text
N
X_1 X_2 ... X_N
```

## Output

Print `N` lines. The `i`-th line should contain the answer when person `i` arrives.

## Sample Input 1

```text
10
5 2 7 4 108728325 390529120 597713292 322456626 845148281 812604915
```

## Sample Output 1

```text
10
7
8
8
108728326
390529121
523096670
452057486
699492475
517144218
```

### Note

When person 1 arrives, there are persons at coordinates `0` and `5`, so the answer is `5 + 5 = 10`.

When person 2 arrives, there are persons at coordinates `0`, `2`, and `5`, so the answer is
`2 + 2 + 3 = 7`.

When person 3 arrives, there are persons at coordinates `0`, `2`, `5`, and `7`, so the answer is
`2 + 2 + 2 + 2 = 8`.

When person 4 arrives, there are persons at coordinates `0`, `2`, `4`, `5`, and `7`, so the answer is
`2 + 2 + 1 + 1 + 2 = 8`.

## Solution

Maintain the current set of occupied coordinates in sorted order.

When a new coordinate `v` is inserted, only a few nearest-neighbor distances can change:

- the new person at `v`;
- the closest existing coordinate on the left of `v`, call it `l`;
- the closest existing coordinate on the right of `v`, call it `r`.

Every other coordinate keeps the same nearest neighbor distance, because inserting `v` cannot become
closer to a point unless that point is adjacent to `v` in sorted order after the insertion.

So each insertion can be handled by:

1. Find `l`, the maximum existing coordinate `< v`.
2. Find `r`, the minimum existing coordinate `> v`.
3. Remove the old contribution of `l` and `r`.
4. Add the new contribution of `l`, `v`, and `r`.
5. Insert `v` into the ordered set.

The coordinate `0` is inserted before processing all arrivals.

### Contribution of One Coordinate

For an existing coordinate `x`, let:

- `left(x)` be the closest existing coordinate smaller than `x`;
- `right(x)` be the closest existing coordinate larger than `x`.

Then its contribution is:

```text
min(x - left(x), right(x) - x)
```

If only one side exists, use the distance to that side.

When inserting `v`, the old contribution of `l` is based on `left(l)` and `r`. After insertion, it
is based on `left(l)` and `v`.

Similarly, the old contribution of `r` is based on `l` and `right(r)`. After insertion, it is based
on `v` and `right(r)`.

The new coordinate `v` contributes:

```text
min(v - l, r - v)
```

again using the only existing side if one side is missing.

This is exactly what the implementation updates in `sum`.

### Ordered Set

The algorithm needs predecessor and successor queries:

- `upperBound(v - 1)` returns the largest coordinate `< v`;
- `lowerBound(v + 1)` returns the smallest coordinate `> v`.

Go has no built-in tree set, so the solution implements a red-black tree.

### Red-Black Tree

A red-black tree is a binary search tree with an extra color on each node. The colors enforce that
the tree height stays `O(log N)`, so search, predecessor, successor, and insertion are all
`O(log N)`.

The implementation uses these rules:

- every newly inserted non-root node starts red;
- the root is always black;
- a red node cannot have a red parent;
- nil children are treated as black.

After inserting a red node, the only possible violation is that its parent is also red. Let:

- `n` be the newly inserted node;
- `p` be its parent;
- `g` be its grandparent;
- `u` be its uncle, the sibling of `p`.

There are two main cases.

#### Case 1: Red Uncle

If `u` exists and is red:

```text
      g(B)                 g(R)
     /   \       =>       /   \
   p(R)  u(R)           p(B)  u(B)
```

Recolor `p` and `u` to black, recolor `g` to red, and continue fixing from `g`. This may push the
red-red conflict upward.

#### Case 2: Black or Missing Uncle

If the uncle is black or nil, recoloring alone cannot fix the height balance. Rotations are needed.

For the straight shape:

```text
      g(B)
     /
   p(R)
   /
 n(R)
```

recolor `p` black, recolor `g` red, and rotate around `g`.

For the bent shape:

```text
      g(B)
     /
   p(R)
     \
     n(R)
```

first rotate around `p` to convert it into the straight shape, then apply the straight-shape fix.

The right-side cases are symmetric.

The important implementation details are:

- new nodes must be red, otherwise no balancing happens and sorted insertions make a chain;
- the first root must be forced black;
- `uncle == nil` must be handled as black, not dereferenced;
- rotations must reconnect the rotated subtree to the old grandparent/root.

### Correctness

For any coordinate, its nearest other person must be one of its immediate neighbors in sorted order.
When inserting `v`, sorted order changes only around the gap where `v` is inserted. Therefore only
`l`, `v`, and `r` can have changed nearest-neighbor distances; all other coordinates keep the same
two immediate neighbors as before.

The algorithm subtracts the old contributions of the affected old coordinates, adds their new
contributions, adds the contribution of `v`, and leaves all unchanged coordinates untouched. Thus
after every insertion, `sum` equals the required total.

The red-black tree maintains the same set of coordinates as a normal binary search tree, while its
balancing rules keep the height logarithmic. Therefore predecessor and successor queries return the
correct adjacent coordinates efficiently.

### Complexity

Each insertion performs a constant number of predecessor/successor queries and one tree insertion.
All of them are `O(log N)` with the red-black tree.

Total complexity is `O(N log N)`, and memory usage is `O(N)`.
