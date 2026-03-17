### Problem

You have a rectangular table of length `a` and width `b`, and infinitely many circular plates of radius `r`.

Two players take turns placing a plate on the table. Rules:

- Plates must not overlap (they may touch).
- Every point of each plate must lie inside the table.
- Once placed, a plate cannot be moved.

The player who cannot make a legal move loses.

Determine who wins when both play optimally: the first player or the second?

### Input

- A single line with three space-separated integers: `a`, `b`, `r`  
  - `1 ≤ a, b, r ≤ 100` — table length, table width, and plate radius.

### Output

- If the **first** player wins, print `First`.
- Otherwise print `Second`.

### Examples

**Input**
```text
5 5 2
```

**Output**
```text
First
```

**Input**
```text
6 7 4
```

**Output**
```text
Second
```

### ideas
1. 考虑平面太复杂了，但是貌似可以按照x轴和y轴考虑
2. 