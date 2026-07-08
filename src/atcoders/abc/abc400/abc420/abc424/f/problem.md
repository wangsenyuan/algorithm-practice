# F - Adding Chords

[Problem link](https://atcoder.jp/contests/abc424/tasks/abc424_f)

**Contest:** [AtCoder Beginner Contest 424](https://atcoder.jp/contests/abc424)

time limit: 3 sec

memory limit: 1024 MiB

score: 525 points

There are `N` points equally spaced on a circle, numbered `1, 2, ..., N` clockwise.

Process `Q` queries in order. Each query is:

- Draw the line segment connecting points `A_i` and `B_i`. If this segment intersects a
  segment that has already been drawn, do not draw it.

It is guaranteed that the `2Q` integers `A_1, ..., A_Q, B_1, ..., B_Q` are pairwise
distinct.

For each query, output whether the segment was drawn.

## Constraints

- `2 <= N <= 10^6`
- `1 <= Q <= 3 * 10^5`
- `1 <= A_i < B_i <= N`
- The `2Q` integers `A_1, ..., A_Q, B_1, ..., B_Q` are pairwise distinct
- All input values are integers

## Input

```text
N Q
A_1 B_1
A_2 B_2
...
A_Q B_Q
```

## Output

Print `Q` lines. On the i-th line, print `Yes` if the i-th segment was drawn, and `No`
otherwise.

## Sample Input 1

```text
8 3
1 5
2 7
3 4
```

## Sample Output 1

```text
Yes
No
Yes
```

- Query 1: draw the segment connecting points 1 and 5
- Query 2: do not draw the segment connecting points 2 and 7 (it intersects the first
  segment)
- Query 3: draw the segment connecting points 3 and 4

## Sample Input 2

```text
999999 4
123456 987654
888888 999999
1 3
2 777777
```

## Sample Output 2

```text
Yes
No
Yes
No
```


### Solution

Write every accepted chord as an interval `(l, r)` with `l < r`.

For a new chord `(l, r)`, an old accepted chord `(a, b)` intersects it iff the four
endpoints are interleaved on the circle:

- `l < a < r < b`, or
- `a < l < b < r`.

The other cases do not intersect:

- both endpoints of the old chord are inside `(l, r)`;
- both endpoints are outside `(l, r)`;
- one interval fully contains the other.

So it is enough to know whether there is an accepted chord with exactly one endpoint in
the open interval `(l, r)`.

Maintain two segment trees over the point positions:

- `s1[x] = y` if an accepted chord starts at `x` and ends at `y`; the tree stores range
  maximums.
- `s2[y] = x` if an accepted chord ends at `y` and starts at `x`; the tree stores range
  minimums.

For a query `(l, r)`:

1. Check accepted chords whose left endpoint is inside `(l, r)`. If
   `max s1[l+1 ... r-1] > r`, then some chord is `(a, b)` with `l < a < r < b`, so the
   new chord intersects it.
2. Check accepted chords whose right endpoint is inside `(l, r)`. If
   `min s2[l+1 ... r-1] < l`, then some chord is `(a, b)` with `a < l < b < r`, so the
   new chord intersects it.
3. If neither condition holds, output `Yes` and insert the chord into both trees.
   Otherwise output `No` and do not insert it.

The guarantee that all query endpoints are pairwise distinct means no two chords share an
endpoint, so the open interval checks cover all possible intersections.

Each query performs two range queries and, only when accepted, two point updates.
Therefore the time complexity is `O(Q log N)`, and the memory complexity is `O(N)`.
