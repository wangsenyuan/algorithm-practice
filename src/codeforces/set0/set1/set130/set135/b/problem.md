# B. Rectangle and Square

[Problem link](https://codeforces.com/problemset/problem/135/B)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

## Problem

Given 8 pairwise distinct points on the plane, split them into two sets of 4
points each:

- one set must be the vertices of a square;
- the other set must be the vertices of a rectangle.

The rectangle is allowed to be a square. Both figures must have non-zero area,
and their sides do not have to be parallel to the axes.

If several valid partitions exist, output any one of them.

## Constraints

- There are exactly 8 input points.
- `|x_i|, |y_i| <= 10^4`.
- No two points coincide.

## Input

The input contains 8 lines. Each line contains two integers `x_i` and `y_i`,
the coordinates of point `i`.

## Output

If no valid partition exists, print:

```text
NO
```

Otherwise print:

```text
YES
<four indexes forming the square>
<four indexes forming the rectangle>
```

Indexes are 1-based and refer to the input order. The indexes may be printed in
any order, but all 8 printed indexes must be pairwise distinct.

## Example

### Input 1

```text
0 0
10 11
10 0
0 11
1 1
2 2
2 1
1 2
```

### Output 1

```text
YES
5 6 7 8
1 2 3 4
```

### Input 2

```text
0 0
1 1
2 2
3 3
4 4
5 5
6 6
7 7
```

### Output 2

```text
NO
```

### Input 3

```text
0 0
4 4
4 0
0 4
1 2
2 3
3 2
2 1
```

### Output 3

```text
YES
1 2 3 4
5 6 7 8
```

## Note

In the third sample, the valid square and rectangle do not have to be parallel
to the coordinate axes.
