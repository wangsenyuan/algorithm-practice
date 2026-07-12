# C. Polygon for the Angle

[Problem link](https://codeforces.com/problemset/problem/1096/C)

**Contest:** [Educational Codeforces Round 57 (Rated for Div. 2)](https://codeforces.com/contest/1096)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given an angle `ang` (in degrees).

For each query, find the minimum number of vertices `n` in a regular `n`-gon such that three vertices
`a`, `b`, and `c` (not necessarily consecutive) exist with `∠abc = ang`. If no such polygon exists,
print `-1`.

If several answers exist, print the smallest. It is guaranteed that when an answer exists, it does not
exceed `998244353`.

## Input

The first line contains an integer `T` (`1 <= T <= 180`) — the number of queries.

Each of the next `T` lines contains one integer `ang` (`1 <= ang < 180`).

## Output

For each query, print one integer `n` (`3 <= n <= 998244353`) — the minimum possible number of
vertices, or `-1` if no such polygon exists.

## Example

### Input

```text
4
54
50
2
178
```

### Output

```text
10
18
90
180
```

### Note

For `ang = 54`, the answer is `10`.

For `ang = 50`, the answer is `18` (for example, `∠v2 v1 v6 = 50°` on a regular 18-gon).

For `ang = 2`, one valid triple gives angle `2°`.

For `ang = 178`, the minimum is `180`, not `90`.

## Solution

Place the regular `n`-gon on its circumcircle. An inscribed angle equals half the central angle of the
opposite arc. If the arc from `a` to `c` passing through `b` spans `k` edges, then:

```text
ang = k * 180 / n
```

Let `g = gcd(ang, 180)`. Write `ang = g * k` and `180 = g * n0`, so a candidate is `n = n0 = 180 / g`.

This candidate works unless the arc uses `n - 1` edges, leaving only one vertex on the other side —
then no triangle can be formed with that angle. That happens exactly when `k + 1 = n`. In that case,
double `n`.

For every query in the constraints, this construction yields a valid minimum `n`.

### Complexity

`O(T log 180)` time, `O(T)` memory.

## Summary

**Key facts about a regular `n`-gon inscribed in its circumcircle.**

- The central angle between two adjacent vertices is `360 / n`.
- The interior angle at a vertex is `(n - 2) * 180 / n`.
- Any three vertices `a`, `b`, `c` form an inscribed angle `∠abc` that equals **half** the central
  angle of the arc `ac` not containing `b`. Since that arc is made of whole edges, if it spans `k`
  edges its central angle is `k * 360 / n`, so:

```text
∠abc = k * 180 / n,   for some integer 1 <= k <= n - 2
```

The picture to keep in mind is:

```text
                  b
               .-' '-.
            .-'   |   '-.
          a ----- O ----- c
            \           /
             \  k edges /
              '-------'

O is the center of the regular polygon.
The lower arc from a to c is the arc that does not contain b.
If that arc contains k polygon edges, its central angle is k * 360 / n.
The inscribed angle at b is half of that, so angle abc = k * 180 / n.
```

**Reduce the target.** We need `k * 180 / n = ang` for the smallest `n`. Let `g = gcd(ang, 180)` and
write `ang = g * k`, `180 = g * n`. This is the unique reduced fraction `ang / 180 = k / n`, so the
smallest polygon that can realize `ang` at all has `n = 180 / g` vertices, using `k = ang / g` edges on
the arc.

**Fix the degenerate case.** The arc must leave at least one vertex for `b` on each side, i.e.
`k <= n - 2`. The reduction only fails when `k = n - 1` (equivalently `k + 1 = n`): the arc wraps all
but one vertex, so no valid `b` remains. Doubling both `k` and `n` keeps the same angle
(`2k * 180 / 2n = k * 180 / n`) while making `2k <= 2n - 2`, which always holds. So the fix is:

```text
n = 180 / gcd(ang, 180)
k = ang / gcd(ang, 180)
if k + 1 == n: n *= 2
```

Visually, the bad case looks like this:

```text
Reduced n-gon when k = n - 1

       b
      / \
     /   \
    c --- a
     \   /
      \ /
   long arc a -> c uses n - 1 edges

The intercepted arc leaves only one edge on the other side.
There is no room to choose three distinct vertices with b outside that arc.
After doubling n, the same angle uses 2k edges out of 2n, so there is enough room.
```

**Why no `-1`.** For any `1 <= ang < 180`, `g = gcd(ang, 180) >= 1` gives a finite `n <= 360`, so an
answer always exists within the constraints; the `-1` branch is never reached.

**Sample check.**

| `ang` | `g = gcd(ang,180)` | `n = 180/g` | `k = ang/g` | `k+1 == n`? | answer |
|-------|--------------------|-------------|-------------|-------------|--------|
| 54    | 18                 | 10          | 3           | no          | 10     |
| 50    | 10                 | 18          | 5           | no          | 18     |
| 2     | 2                  | 90          | 1           | no          | 90     |
| 178   | 2                  | 90          | 89          | yes         | 180    |

The last row is the degenerate case: `k + 1 = 90 = n`, so `n` doubles to `180`.
