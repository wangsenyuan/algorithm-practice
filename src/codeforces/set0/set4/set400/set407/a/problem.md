# Problem 407A - Triangle

## Problem Statement

There is a right triangle with legs of length `a` and `b`. Your task is to determine whether it is possible to locate the triangle on the plane in such a way that none of its sides is parallel to the coordinate axes. All the vertices must have integer coordinates. If there exists such a location, you have to output the appropriate coordinates of vertices.

## Input

The first line contains two integers `a`, `b` (1 ≤ a, b ≤ 1000), separated by a single space.

## Output

In the first line print either "YES" or "NO" (without the quotes) depending on whether the required location exists. 

If it does, print in the next three lines three pairs of integers — the coordinates of the triangle vertices, one pair per line. The coordinates must be integers, not exceeding 10^9 in their absolute value.


## ideas
1. 固定一个顶点在(0, 0)
2. 然后画两个圆r = a, r = b
3. 这样子，就最多有1000 * 1000个点，然后，检查这些点的组合，是否能产生一个合法的三角形
4. 