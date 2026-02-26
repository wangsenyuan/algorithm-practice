# Problem

Inna and Dima want to surprise Sereja with a huge candy matrix. Rows are numbered $1$ to $n$ (top to bottom) and columns $1$ to $m$ (left to right). Cell $(i, j)$ is the intersection of row $i$ and column $j$. There are $p$ candies; the $k$-th candy is at $(x_k, y_k)$.

Sereja then applies three operations:

1. **Rotate clockwise** $x$ times (each time by $90°$).
2. **Horizontal flip** $y$ times (mirror the matrix left–right).
3. **Rotate counterclockwise** $z$ times (each time by $90°$).

The candies are not damaged; only their coordinates change. For each candy, find its **new coordinates** after all operations.

---

## Input

- Line 1: six integers $n$, $m$, $x$, $y$, $z$, $p$ ($1 \le n, m \le 10^9$; $0 \le x, y, z \le 10^9$; $1 \le p \le 10^5$).
- Next $p$ lines: two integers $x_k$, $y_k$ ($1 \le x_k \le n$, $1 \le y_k \le m$) — initial coordinates of the $k$-th candy. Multiple candies may share a cell.

## Output

For each of the $p$ candies, output one line with its new coordinates (two space-separated integers).

---

## Examples

**Input**

```
3 3 3 1 1 9
1 1
1 2
1 3
2 1
2 2
2 3
3 1
3 2
3 3
```

**Output**

```
1 3
1 2
1 1
2 3
2 2
2 1
3 3
3 2
3 1
```

---

## Note

**Horizontal flip** means mirroring the matrix (e.g. column $j$ becomes column $m+1-j$). Rotations are as usual (clockwise: row/column roles swap and one dimension is flipped).
