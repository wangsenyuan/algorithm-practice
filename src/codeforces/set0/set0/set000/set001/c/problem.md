# C. Ancient Berland Circus

[Problem link](https://codeforces.com/problemset/problem/1/C)

**Contest:** [Codeforces Beta Round #1](https://codeforces.com/contest/1)

time limit per test: 2 seconds

memory limit per test: 64 megabytes

input: standard input

output: standard output

Modern Berland circuses use a round arena with diameter 13 meters. In Ancient Berland, arenas were
regular (equiangular) polygons. Each corner had a pillar, and ropes between pillars marked the edges.

Scientists found the remains of an ancient arena: only three pillars remain. You are given their
coordinates. Find the smallest possible area of the original arena.

## Input

Three lines. Each line contains two numbers — the coordinates of one pillar. Every coordinate has
absolute value at most 1000 and is given with at most six digits after the decimal point.

## Output

Print the smallest possible area of the ancient arena. The answer should be accurate to at least 6
digits after the decimal point. It is guaranteed that the optimal polygon has at most 100 vertices.

## Example

### Input

```text
0.000000 0.000000
1.000000 1.000000
0.000000 1.000000
```

### Output

```text
1.00000000
```

## Solution

All vertices of a regular polygon lie on one circle. The three remaining pillars are three vertices of
the original regular polygon, so they lie on the same circumcircle. Because three non-collinear points
determine exactly one circle, the center `O` and radius `R` of that circle are fixed no matter how many
vertices the original polygon had.

First compute the circumcenter `O` of the three points. One geometric way to see this is to intersect
the perpendicular bisectors of two triangle sides; the implementation uses the equivalent coordinate
formula. Then compute the radius `R` as the distance from `O` to any given point.

Next compute the polar angle of each point around `O` with `atan2`, normalize it into `[0, 2*pi)`, and
sort the three angles. The three arcs between consecutive given vertices are:

```text
arc1 = angle[1] - angle[0]
arc2 = angle[2] - angle[1]
arc3 = 2*pi - angle[2] + angle[0]
```

For a regular `n`-gon, one central step is:

```text
step = 2*pi / n
```

The three given vertices can belong to this `n`-gon exactly when each arc is an integer number of
steps:

```text
arc1 = w * step
arc2 = v * step
arc3 = u * step
```

for positive integers `w`, `v`, and `u`. Since the three arcs cover the whole circle, these integers
also satisfy `w + v + u = n`.

The statement guarantees that the optimal polygon has at most `100` vertices, so try `n` from `3` to
`100` and take the first one whose three arc ratios are all close to integers. The first valid `n`
gives the smallest possible area, because for a fixed circumradius the area of an inscribed regular
polygon increases as `n` increases.

For that `n`, split the polygon into `n` isosceles triangles with vertex at the center. Each triangle
has area:

```text
R^2 * sin(2*pi/n) / 2
```

so the total area is:

```text
n * R^2 * sin(2*pi/n) / 2
```

### Correctness

The circumcenter calculation gives the unique circle through the three given pillars. Any valid
original regular polygon must use this same circle, because all of its vertices are concyclic and it
contains the three given points.

For a fixed `n`, the vertices of a regular `n`-gon divide the circle into equal central angles
`2*pi/n`. Therefore two chosen vertices of that polygon are separated by an integer number of those
steps. The algorithm checks exactly this condition for the three circular arcs between the given
points, so it accepts precisely the values of `n` for which all three points can be vertices of the
same regular `n`-gon.

The algorithm tries `n` in increasing order. The first accepted `n` is therefore the valid polygon with
the fewest vertices. With fixed circumradius, an inscribed regular polygon's area increases as the
number of vertices increases, so this first valid polygon has the smallest possible area.

Finally, the area formula sums the `n` congruent triangles formed by the center and adjacent polygon
vertices, so it returns the area of that polygon. Thus the algorithm returns the required minimum area.

### Complexity

Only `n = 3..100` are tested, and each check uses three arcs. The running time is `O(1)`, and the
memory usage is `O(1)`.
