# Problem Description

You are given n distinct points on a plane with integral coordinates. For each point you can either draw a vertical line through it, draw a horizontal line through it, or do nothing.

You consider several coinciding straight lines as a single one. How many distinct pictures you can get? Print the answer modulo 10^9 + 7.

## Input

The first line contains single integer n (1 ≤ n ≤ 10^5) — the number of points.

n lines follow. The (i + 1)-th of these lines contains two integers xi, yi (-10^9 ≤ xi, yi ≤ 10^9) — coordinates of the i-th point.

It is guaranteed that all points are distinct.

## Output

Print the number of possible distinct pictures modulo 10^9 + 7.

## Examples

### Example 1

**Input:**
```
4
1 1
1 2
2 1
2 2
```

**Output:**
```
16
```

### Example 2

**Input:**
```
2
-1 -1
0 1
```

**Output:**
```
9
```

## Note

In the first example there are two vertical and two horizontal lines passing through the points. You can get pictures with any subset of these lines. For example, you can get the picture containing all four lines in two ways (each segment represents a line containing it).

The first way: The second way:

In the second example you can work with two points independently. The number of pictures is 3^2 = 9.

## Ideas

1. 没想法～～～

## Solution

Let's build graph on points. Add edge from point to left, right, top and bottom neighbours (if such neighbour exist). Note that we can solve problem independently for each connected component and then print product of answer for components. So we can consider only connected graphs without loss of generality.

Let's define X as number of different x-coords, Y as number of different y-coords.

### Case 1: Graph Contains a Cycle

What if graph contains some cycle? Let's consider this cycle without immediate vertices (vertices that lie on the same line with previous and next vertices of cycle). Draw a line from each such vertex to the next vertex of cycle (and from last to the first). We got all lines that are corresponding to x-coords and y-coords of vertices of cycle. Let's prove by induction that we can got all such lines from the whole graph.

Run depth-first search from vertices of cycle. Let we enter in some vertex that is not from cycle. It must have at least one visited neighbour. By induction for graph consisting of visited vertices we can get all lines. So there is line from visited neighbour. Draw line in another direction and continue depth-first search. Sooner or later we will get all lines for the whole graph. Please note that intermediate vertices of cycle will be processed correctly too.

If we can get all lines then we can get all subsets of lines. Answer for cyclic graph is 2^(X + Y).

### Case 2: Acyclic Graph (Tree)

Now look at another case — acyclic graph or tree. We can prove that we can get any incomplete subset of lines.

Let's fix subset and some line not from this subset. Just draw this line without restriction. By similar induction as in cyclic graph case we can prove that we can get all lines (but fixed line doesn't exist really).

Now let's prove that it is impossible to take all lines. For graph consisting of only one vertex it is obvious. Else fix some leaf. We must draw a line which are not directed to any neighbour because it is the only way to draw this line. But now we have tree with less number of vertices. So our statement is correct by induction.

Answer for tree is 2^(X + Y) - 1.

### Summary

So the problem now is just about building graph on points and checking each component on having cycles.
