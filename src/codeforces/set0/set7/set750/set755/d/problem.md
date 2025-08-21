# PolandBall and Polygon

## Problem Description

PolandBall has such a convex polygon with n vertices that no three of its diagonals intersect at the same point. PolandBall decided to improve it and draw some red segments.

He chose a number k such that gcd(n, k) = 1. Vertices of the polygon are numbered from 1 to n in a clockwise way. PolandBall repeats the following process n times, starting from the vertex 1:

Assume you've ended last operation in vertex x (consider x = 1 if it is the first operation). Draw a new segment from vertex x to k-th next vertex in clockwise direction. This is a vertex x + k or x + k - n depending on which of these is a valid index of polygon's vertex.

Your task is to calculate number of polygon's sections after each drawing. A section is a clear area inside the polygon bounded with drawn diagonals or the polygon's sides.

## Input

There are only two numbers in the input: n and k (5 ≤ n ≤ 10^6, 2 ≤ k ≤ n - 2, gcd(n, k) = 1).

## Output

You should print n values separated by spaces. The i-th value should represent number of polygon's sections after drawing first i lines.

## Examples

### Example 1
**Input:**
```
5 2
```

**Output:**
```
2 3 5 8 11
```

### Example 2
**Input:**
```
10 3
```

**Output:**
```
2 3 4 6 9 12 16 21 26 31
```

## Note

The greatest common divisor (gcd) of two integers a and b is the largest positive integer that divides both a and b without a remainder.

For the first sample testcase, you should output "2 3 5 8 11". Pictures below correspond to situations after drawing lines.


### solution

It can be easily visualized by drawing a polygon by hand.

In the initial round around the polygon going from 1 to n, the diagonals will not intersect. So every diagonal will produce 1 extra section. The final diagonal around which cross the point 1 will produce +2 sections because it will intersect the diagonal which starts at point 1 to (k+1). Every subsequent diagonal will intersect with two other diagonals (one starting from a point before their starting point and one between the starting and ending point) producing +3 sections each time. Suppose x sections were introduced with every diagonal. The final diagonal which crosses point 1 will introduce x+1 sections. And every subsequent diagonal in the next round will produce x+2 sections. Complexity: O(n) because every point becomes a starting point for a diagonal exactly once.