### Problem

You are given an `n × m` grid. Each cell has a lowercase letter (plot type), except:
- **S** — start (visited exactly once at the beginning)
- **T** — target (visited exactly once at the end)

Rules:
- Move only to side-adjacent cells; time per step is 1 minute.
- You may use **at most k different plot types** along the path (same type can be used any number of times). S and T do not count as types.
- Find a path from S to T that:
  1. **Minimizes total time** (shortest path).
  2. Among all shortest paths, is **lexicographically minimal** when the path is written as the sequence of plot-type letters (in order of visit).

If no valid path exists, output `-1`.

### Input

- First line: three integers `n`, `m`, `k`  
  - `1 ≤ n, m ≤ 50`, `n·m ≥ 2`, `1 ≤ k ≤ 4`
- Next `n` lines: the grid, each line exactly `m` characters (lowercase letters, `S`, `T`). Exactly one `S` and one `T` on the grid.

### Output

- If a valid path exists: print the path as the **sequence of plot-type letters** (no S at start, no T at end).
- If no valid path: print `-1`.
- The sequence may be empty (e.g. S and T adjacent); printing nothing or a newline is accepted.

### Examples

**Input**
```
5 3 2
Sba
ccc
aac
ccc
abT
```

**Output**
```
bcccc
```

**Input**
```
3 4 1
Sxyy
yxxx
yyyT
```

**Output**
```
xxxx
```

**Input**
```
1 3 3
TyS
```

**Output**
```
y
```

**Input**
```
1 4 1
SxyT
```

**Output**
```
-1
```

### Summary

- **BFS/state**: state = (row, col, set of plot types used so far). At most 4 types ⇒ state space is manageable.
- **Goal**: shortest path; ties broken by lexicographically smallest sequence of letters.
- **Output**: string of plot types along the path (excluding S and T), or `-1` if impossible.
