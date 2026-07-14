# D - Cutting Woods

[Problem link](https://atcoder.jp/contests/abc217/tasks/abc217_d)

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 400 points

## Problem Statement

We have a long piece of timber with a length of `L` meters. For each
`x = 1, 2, ..., L - 1`, there is a mark called Mark `x` at `x` meters from the
left end of the piece.

You are given `Q` queries, the `i`-th of which is represented as a pair of
numbers `(c_i, x_i)`. Process the queries in ascending order of `i` as
described below.

- If `c_i = 1`: cut the piece at Mark `x_i` into two.
- If `c_i = 2`: choose the piece with Mark `x_i` on it and print its length.

For both kinds of queries, it is guaranteed that there will have been no cut
at Mark `x_i` when the query is to be processed.

## Constraints

- `1 <= L <= 10^9`
- `1 <= Q <= 2 * 10^5`
- `c_i = 1` or `2`
- `1 <= x_i <= L - 1`
- For every `i`, there is no earlier cut query at the same mark `x_i`
- All values in input are integers

## Input

```
L Q
c_1 x_1
c_2 x_2
...
c_Q x_Q
```

## Output

Print one line per query with `c_i = 2`: the length of the piece containing
Mark `x_i`.

## Samples

### Sample 1

Input:

```
5 3
2 2
1 3
2 2
```

Output:

```
5
3
```

### Sample 2

Input:

```
5 3
1 2
1 4
2 3
```

Output:

```
2
```

### Sample 3

Input:

```
100 10
1 31
2 41
1 59
2 26
1 53
2 58
1 97
2 93
1 23
2 84
```

Output:

```
69
31
6
38
38
```

## Solution

### 1. What a length query needs

At any moment, the existing cuts divide the timber into pieces. For a query at
position `x`, let:

```text
l = the greatest existing cut strictly smaller than x
r = the smallest existing cut strictly greater than x
```

Then `x` lies in the piece `[l,r]`, whose length is:

```text
r - l
```

The two ends of the timber can be treated as cuts that always exist at
positions `0` and `L`.

Therefore, each query only requires three operations:

1. insert a new cut;
2. find the nearest existing cut to the left;
3. find the nearest existing cut to the right.

### 2. Why coordinate compression is needed

The timber length can be as large as `10^9`, so we cannot build a data
structure with one entry for every possible position.

However, a position can matter only if it appears in a query. The code first
collects:

```text
every query position x, together with 0 and L
```

It sorts them and removes duplicates:

```go
slices.Sort(markers)
markers = slices.Compact(markers)
```

The index of `x` in this sorted array is found by:

```go
i := sort.SearchInts(markers, x)
```

Compression preserves the order of all positions. Thus, searching to the left
or right among compressed indices is equivalent to searching to the left or
right on the original timber.

Knowing future coordinates does not mean that future cuts are active. The
segment trees described below store only cuts that have already been made.

### 3. Two segment trees

The implementation maintains two segment trees over the compressed indices.

#### Maximum tree `t1`

At the index of every active cut, `t1` stores its original coordinate. Every
inactive index stores `-inf`. Its operation is `max`.

For a query position at index `i`:

```go
l := t1.Get(0, i)
```

The half-open range `[0,i)` contains exactly the compressed coordinates
strictly smaller than `x`. Taking their maximum returns the nearest active cut
to the left.

#### Minimum tree `t2`

At the index of every active cut, `t2` also stores its original coordinate.
Every inactive index stores `inf`. Its operation is `min`.

The query:

```go
r := t2.Get(i, n)
```

searches all compressed coordinates greater than or equal to `x` and returns
the nearest active cut on the right.

The range includes index `i`, but the problem guarantees that `x` has not
already been cut when either kind of query is processed. Consequently, index
`i` is inactive during a type-2 query, so the returned active cut is strictly
greater than `x`.

Initially, only the timber boundaries are active:

```go
t1.Update(0, 0)
t2.Update(n-1, L)
```

These sentinels guarantee that every type-2 query has both a left and a right
boundary.

### 4. Processing the queries

For a cut query `1 x`, activate `x` in both trees:

```go
t1.Update(i, x)
t2.Update(i, x)
```

For a length query `2 x`, find the two boundaries and append their distance:

```go
r := t2.Get(i, n)
l := t1.Get(0, i)
ans = append(ans, r-l)
```

Although all coordinates were compressed in advance, queries are still
processed in their original order. A cut affects only the queries that occur
after its update.

### 5. Correctness proof

We prove that every answer produced by the algorithm is correct.

#### Lemma 1: the two segment trees contain exactly the active cuts

Initially, the algorithm activates only `0` and `L`, which are precisely the
two permanent timber boundaries. Whenever a type-1 query cuts at `x`, the
algorithm updates the compressed index of `x` in both trees. No other update
is performed. Therefore, after every processed query, a finite value is stored
at an index in the two trees if and only if that coordinate is an active cut.

#### Lemma 2: `t1.Get(0,i)` returns the nearest cut to the left of `x`

Because the compressed coordinates are sorted, `[0,i)` contains exactly the
coordinates smaller than `x`. By Lemma 1, inactive coordinates contribute
`-inf`, while every active cut contributes its coordinate. The range maximum
is therefore the greatest active cut smaller than `x`, which is the nearest
cut on its left.

#### Lemma 3: `t2.Get(i,n)` returns the nearest cut to the right of `x`

The range `[i,n)` contains the coordinates greater than or equal to `x`.
During a type-2 query, `x` is not an active cut, so its own index contributes
`inf`. By Lemma 1, the range minimum is therefore the smallest active cut
strictly greater than `x`, which is the nearest cut on its right.

#### Theorem: every type-2 answer is the length of the piece containing `x`

Let `l` and `r` be the cuts returned by the two range queries. By Lemmas 2 and
3, there is no active cut between `l` and `x`, nor between `x` and `r`.
Therefore, `l` and `r` are exactly the boundaries of the unique timber piece
containing `x`. The algorithm outputs `r-l`, which is precisely that piece's
length.

### 6. Complexity

There are at most `Q+2` compressed coordinates.

- Sorting and deduplicating them costs `O(Q log Q)` time.
- Each query performs one binary search and at most two segment-tree
  operations, each in `O(log Q)` time.
- The marker array and the two segment trees use `O(Q)` space.

Thus the total complexity is:

```text
Time:  O(Q log Q)
Space: O(Q)
```
