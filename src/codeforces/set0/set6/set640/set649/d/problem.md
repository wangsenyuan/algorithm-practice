### Problem
Computer memory consists of $n$ cells arranged in a row (numbered $1$ to $n$ from left to right). Each cell is either free or assigned to a process.

For each process, the cells it owns form a contiguous segment. Using operations of the form "copy data from an occupied cell to a free cell, then mark the original cell free," arrange memory so that all occupied cells are to the left of all free cells. The relative order of processes must be preserved: if all cells of process $i$ were before all cells of process $j$ initially, this must remain true after the operations. The order of cells within a process may change.

At least one memory cell is occupied, and process numbers are unique.

### Input
- The first line contains an integer $n$ ($1 \le n \le 200\,000$) — the number of memory cells.
- The second line contains $n$ integers $a_1, a_2, \dots, a_n$ ($0 \le a_i \le n$), where $a_i = 0$ means the $i$-th cell is free, otherwise $a_i$ is the id of the process that owns the $i$-th cell. It is guaranteed that at least one $a_i \ne 0$. Process ids are integers from $1$ to $n$ in arbitrary order (not necessarily consecutive).

### Output
Print a single integer — the minimum number of operations needed to defragment the memory.

### Examples

Input
```
4
0 2 2 1
```
Output
```
2
```

Input
```
8
0 8 8 8 0 4 4 2
```
Output
```
4
```

### Note
In the first example, two operations are sufficient:
1. Copy the data from the third cell to the first → memory becomes: `2 2 0 1`.
2. Copy the data from the fourth cell to the third → memory becomes: `2 2 1 0`.
