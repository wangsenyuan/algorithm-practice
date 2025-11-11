# Problem D - Triangles with Non-Zero Area

Vanya got bored and he painted $n$ distinct points on the plane. After that he connected all the points pairwise and saw that as a result many triangles were formed with vertices in the painted points. He asks you to count the number of the formed triangles with the non-zero area.

## Input

The first line contains integer $n$ ($1 \leq n \leq 2000$) — the number of the points painted on the plane.

Next $n$ lines contain two integers each $x_i$, $y_i$ ($-100 \leq x_i, y_i \leq 100$) — the coordinates of the $i$-th point. It is guaranteed that no two given points coincide.

## Output

In the first line print an integer — the number of triangles with the non-zero area among the painted points.

## Examples

### Example 1

**Input:**
```
4
0 0
1 1
2 0
2 2
```

**Output:**
```
3
```

### Example 2

**Input:**
```
3
0 0
1 1
2 0
```

**Output:**
```
1
```

### Example 3

**Input:**
```
1
1 1
```

**Output:**
```
0
```

## Note

**Example 1:** There are 3 triangles formed: $(0, 0) - (1, 1) - (2, 0)$; $(0, 0) - (2, 2) - (2, 0)$; $(1, 1) - (2, 2) - (2, 0)$.

**Example 2:** There is 1 triangle formed: $(0, 0) - (1, 1) - (2, 0)$.

**Example 3:** A single point doesn't form a single triangle.
