# D — Bicycle race on Lake Lucerne

## Problem description

Maria is in a bicycle race on a speedway that follows the **shoreline** of Lake Lucerne. The shore is a closed polygonal path made only of **axis-aligned** segments (north, south, east, or west).

Use a standard grid: **Ox** from west to east, **Oy** from south to north. The **start** is the **southernmost** point of the track; if several are tied, choose the **westernmost** among them. Racers begin moving **north**. On each straight section they move in one of the four cardinal directions and may change direction only at **vertices** between sections. They never make a **180°** turn (no north↔south or east↔west reversals).

Maria calls a **turn** (at some vertex) **dangerous** if, **when that turn is ignored** (she keeps going straight in the direction of the incoming segment), she would **immediately** end up **in the water** (off the safe track).

Determine the **number of dangerous turns** on the track.

## Input

- The first line contains an integer **n** (4 ≤ n ≤ 1000) — the number of **straight sections** (edges) of the track.
- The next **n + 1** lines each contain two integers **xᵢ**, **yᵢ** (−10⁴ ≤ xᵢ, yᵢ ≤ 10⁴). The points are given in order along the track: the **first** point is the start, the **i**-th segment runs from **(xᵢ, yᵢ)** to **(xᵢ₊₁, yᵢ₊₁)** for i = 1 … n, and the path **closes** (last point equals the first).

**Guarantees:**

- The first segment is directed **north**.
- The first point is the southernmost point of the track (westernmost tie-break).
- The last point coincides with the first (closed contour).
- Non-adjacent segments do not share points; consecutive segments share exactly one vertex.
- No repeated vertices except the first and last being the same.
- Two consecutive segments are neither **collinear** (same direction) nor **opposite** directions.

## Output

Print one integer — the number of **dangerous** turns.

## Examples

### Example 1

**Input:**

```
6
0 0
0 1
1 1
1 2
2 2
2 0
0 0
```

**Output:**

```
1
```

### Example 2

**Input:**

```
16
1 1
1 5
3 5
3 7
2 7
2 9
6 9
6 7
5 7
5 3
4 3
4 4
3 4
3 2
5 2
5 1
1 1
```

**Output:**

```
6
```
