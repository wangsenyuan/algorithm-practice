# A. Two Squares

[Problem link](https://codeforces.com/problemset/problem/993/A)

**Contest:** [Codeforces Round #493 (Div. 2)](https://codeforces.com/contest/993)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given two squares:

1. One with sides parallel to the coordinate axes.
2. One with sides at 45 degrees to the coordinate axes.

Determine whether the two squares intersect. The interior counts as part of each square, so one
square fully inside the other still counts as intersection. A single common point also counts.

## Input

Two lines, one per square. Each line contains four pairs of integers — the coordinates of the square's
vertices in clockwise or counterclockwise order.

The first line is the axis-aligned square; the second is the rotated square.

All coordinates are integers between -100 and 100.

## Output

Print `Yes` if the squares intersect, otherwise print `No`. Letter case does not matter.

## Example

### Input

```text
0 0 6 0 6 6 0 6
1 3 3 5 5 3 3 1
```

### Output

```text
YES
```

### Input

```text
0 0 6 0 6 6 0 6
7 3 9 5 11 3 9 1
```

### Output

```text
NO
```

### Input

```text
6 0 6 6 0 6 0 0
7 4 4 7 7 10 10 7
```

### Output

```text
YES
```

### Note

In the first example, the rotated square lies entirely inside the axis-aligned square.

In the second example, the squares share no points.

## Solution

Let the first square be the axis-aligned one. From its four vertices we can get
`minX`, `maxX`, `minY`, and `maxY`, so checking whether a point is inside it is just a
range test on both coordinates. Boundary points count, so the comparisons are inclusive.

First, test all four vertices of the rotated square. If any of them is inside the
axis-aligned square, the two squares intersect.

This alone is not enough, because the axis-aligned square may lie inside the rotated
square, or the two squares may cross while no tested vertex of the rotated square is
inside the first square. To handle the opposite direction with the same helper, use the
coordinate transform:

```text
u = x + y
v = x - y
```

Under this transform, a square whose sides are at 45 degrees to the original axes becomes
axis-aligned in `(u, v)` coordinates, while the original axis-aligned square becomes a
45-degree rotated square. Therefore, after transforming both squares, we can run the same
intersection test again with their roles swapped.

The helper also checks the center of the second square. The center is the average of the
four vertices:

```text
centerX = (x1 + x2 + x3 + x4) / 4
centerY = (y1 + y2 + y3 + y4) / 4
```

The implementation keeps this as integer arithmetic by comparing the sums directly
against `4 * minX`, `4 * maxX`, `4 * minY`, and `4 * maxY`. This avoids losing precision
when the center has half-integer coordinates after transformation.

If either direction finds a vertex or center inside the axis-aligned square, the answer is
`YES`; otherwise, the squares are disjoint and the answer is `NO`.

The algorithm uses only a constant number of point checks, so the time complexity is
`O(1)` and the memory usage is `O(1)`.
