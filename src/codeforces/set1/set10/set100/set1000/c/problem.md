# Problem C: Segment Coverage

## Problem Statement

You are given $n$ segments on a coordinate line; each endpoint of every segment has integer coordinates. Some segments can degenerate to points. Segments can intersect with each other, be nested in each other or even coincide.

Your task is the following: for every $k \in [1..n]$, calculate the number of points with integer coordinates such that the number of segments that cover these points equals $k$. A segment with endpoints $l_i$ and $r_i$ covers point $x$ if and only if $l_i \leq x \leq r_i$.

## Input

The first line of the input contains one integer $n$ $(1 \leq n \leq 2 \cdot 10^5)$ — the number of segments.

The next $n$ lines contain segments. The $i$-th line contains a pair of integers $l_i, r_i$ $(0 \leq l_i \leq r_i \leq 10^{18})$ — the endpoints of the $i$-th segment.

## Output

Print $n$ space separated integers $cnt_1, cnt_2, \ldots, cnt_n$, where $cnt_i$ is equal to the number of points such that the number of segments that cover these points equals to $i$.

## Examples

### Example 1
**Input:**
```
3
0 3
1 3
3 8
```

**Output:**
```
6 2 1
```

### Example 2
**Input:**
```
3
1 3
2 4
5 7
```

**Output:**
```
5 2 0
```


### ideas
1. 