# F - Grid Clipping (ABC449)

**Contest:** [ABC449](https://atcoder.jp/contests/abc449) — AtCoder Beginner Contest 449  
**Task:** [https://atcoder.jp/contests/abc449/tasks/abc449_f](https://atcoder.jp/contests/abc449/tasks/abc449_f)

**Time limit:** 4 sec / **Memory limit:** 1024 MiB  
**Score:** 500 points

## Problem Statement

There is a grid with `H` rows and `W` columns. The cell at the `r`-th row from the
top and `c`-th column from the left is called cell `(r, c)`.

Each cell is painted black or white:

- cells `(R_k, C_k)` are black for `k = 1, 2, ..., N`
- the remaining `H * W - N` cells are white

Find the number of rectangular regions with height `h` and width `w` such that every
cell inside the rectangle is white.

More formally, count pairs of integers `(r_0, c_0)` satisfying all of the
following:

- `1 <= r_0 <= H - h + 1`
- `1 <= c_0 <= W - w + 1`
- cell `(r_0 + i, c_0 + j)` is white for all integers `(i, j)` with
  `0 <= i < h` and `0 <= j < w`

## Constraints

- `1 <= h <= H <= 10^9`
- `1 <= w <= W <= 10^9`
- `0 <= N <= 2 * 10^5`
- `1 <= R_k <= H`
- `1 <= C_k <= W`
- `(R_{k_1}, C_{k_1}) != (R_{k_2}, C_{k_2})` for `k_1 != k_2`
- All input values are integers.

## Input

The input is given from standard input in the following format:

```text
H W h w N
R_1 C_1
R_2 C_2
:
R_N C_N
```

## Output

Print the answer.

## Sample Input 1

```text
3 4 2 2 3
1 3
2 4
3 1
```

## Sample Output 1

```text
2
```

There are two valid `2 x 2` rectangles whose top-left corners are `(1, 1)` and
`(1, 2)`.

## Sample Input 2

```text
4 4 3 2 2
2 2
3 4
```

## Sample Output 2

```text
0
```

No rectangle of size `3 x 2` is entirely white.

## Sample Input 3

```text
449 449 3 14 0
```

## Sample Output 3

```text
194892
```

There are no black cells.

## Sample Input 4

```text
31 9 5 7 10
14 8
8 4
18 8
12 1
8 5
9 6
18 1
14 7
5 6
26 7
```

## Sample Output 4

```text
12
```


### ideas
1. 好难~
2. 如果不考虑black cell, 那么一共有 (H - h + 1) * (W - w + 1)个(h * w)的区域
3. 那么怎么排除掉这些black的区域呢?
4. 对于一个black cell, 那么包含它的区域,确实可以找出来
5. i >= r - h + 1, i <= r
6. j >= c - w + 1, j <= c
7. 但是,问题在于,还有重叠. 
8. 

## Solution Summary

Count possible top-left corners instead of cells in the original grid.

Without black cells, the number of `h x w` rectangles is:

```text
(H - h + 1) * (W - w + 1)
```

Let a rectangle's top-left corner be `(x, y)`. For a black cell `(r, c)`, all
top-left corners of rectangles containing it satisfy:

```text
r - h + 1 <= x <= r
c - w + 1 <= y <= c
```

After clipping to the valid top-left-corner ranges, this black cell forbids one
axis-aligned rectangle:

```text
x in [max(1, r - h + 1), min(H - h + 1, r)]
y in [max(1, c - w + 1), min(W - w + 1, c)]
```

So the answer is:

```text
all possible top-left corners - union area of forbidden rectangles
```

The forbidden rectangles may overlap, so subtracting them one by one is not
valid. Instead, compute their union area with a sweep line:

1. Convert every forbidden rectangle into two x-events:
   - add its y-interval at `x1`
   - remove its y-interval at `x2 + 1`
2. Compress all y-interval endpoints. Since the intervals are inclusive on
   integer coordinates, store `[y1, y2 + 1)` as a half-open interval.
3. Sweep events in increasing x. Between two neighboring event x-coordinates,
   the covered y-length is constant.
4. A segment tree over compressed y-coordinates maintains the total covered
   y-length under interval add/remove operations.

For a sweep gap of width `dx`, if the segment tree says `coveredY` y-coordinates
are forbidden, then this contributes:

```text
dx * coveredY
```

forbidden top-left corners.

No additional area needs to be added after the final event. Each forbidden
x-interval `[x1, x2]` is represented by an add event at `x1` and a remove event
at `x2 + 1`. Before processing the events at coordinate `x`, the sweep adds the
area of the half-open range `[prevX, x)` using the coverage produced by the
previous events. Therefore, the remove event at `x2 + 1` completes the counting
of exactly the integer x-coordinates `x1, x1 + 1, ..., x2`.

The last event is necessarily a removal boundary for the rectangles that extend
furthest to the right. Once it is processed, no forbidden rectangle remains
active to its right, so there is no remaining covered range to count.

For example, an inclusive x-interval `[2, 4]` creates an add event at `2` and a
remove event at `5`. The sweep counts the half-open interval `[2, 5)`, which
contains exactly the integer coordinates `2`, `3`, and `4`. Nothing after `5`
belongs to that rectangle.

Finally subtract the union area from the total number of top-left corners.

## Correctness

Every valid white rectangle is uniquely represented by its top-left corner.
There are `(H - h + 1) * (W - w + 1)` possible top-left corners in total.

For any black cell `(r, c)`, a rectangle with top-left corner `(x, y)` contains
that cell exactly when:

```text
x <= r <= x + h - 1
y <= c <= y + w - 1
```

Rearranging gives exactly the forbidden top-left-corner rectangle above. Thus a
top-left corner is invalid if and only if it belongs to at least one forbidden
rectangle.

The sweep line counts each integer top-left corner in the union of forbidden
rectangles exactly once. The segment tree stores whether each compressed
y-segment is covered by at least one active forbidden rectangle, so overlapping
rectangles do not cause double counting. Therefore the computed union area is
exactly the number of invalid top-left corners.

Subtracting invalid top-left corners from all top-left corners gives exactly the
number of all-white `h x w` rectangles.

## Complexity

There are at most `2N` sweep events and `2N` compressed y-endpoints. Sorting and
all segment-tree updates take `O(N log N)` time, and the memory usage is `O(N)`.
