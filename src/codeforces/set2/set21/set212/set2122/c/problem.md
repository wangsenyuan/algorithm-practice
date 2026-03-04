# Problem

- **Setting:** $n$ points $(x_i, y_i)$ on the plane, $n$ even.
- **Task:** Partition the $n$ points into $\frac{n}{2}$ disjoint pairs (each point in exactly one pair).
- **Goal:** Maximize the sum of Manhattan distances over all pairs, i.e. $\sum_{i=1}^{n/2} |x_{a_i}-x_{b_i}| + |y_{a_i}-y_{b_i}|$.
- **Input:** Multiple test cases; each gives $n$ and $n$ point coordinates; total $n$ over all cases $\le 2 \cdot 10^5$.
- **Output:** For each case, $\frac{n}{2}$ lines of two integers each (indices of the paired points); any valid solution if multiple exist.

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 10^4$). The description of the test cases follows.

The first line of each test case contains a single even integer $n$ ($2 \le n \le 2 \cdot 10^5$) — the number of points.

The $i$-th of the next $n$ lines contains two integers $x_i$ and $y_i$ ($-10^6 \le x_i, y_i \le 10^6$) — the coordinates of the $i$-th point.

It is guaranteed that the sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

## Output

For each test case, output $\frac{n}{2}$ lines, the $i$-th line containing two integers $a_i$ and $b_i$ — the indices of points in the $i$-th pair.

If there are multiple solutions, print any of them.

## Examples

### Example 1

**Input:**

```
2
4
1 1
3 0
4 2
3 4
10
-1 -1
-1 2
-2 -2
-2 0
0 2
2 -3
-4 -4
-4 -2
0 1
-4 -2
```

**Output:**

```
4 1
2 3
8 1
9 10
7 5
2 3
6 4
```

## Note

In the first test case, an optimal solution is to select the pairs $(1, 4)$ and $(2, 3)$, which achieves a distance sum of $5 + 3 = 8$.

In the second test case, an optimal solution is to select the pairs $(1, 8)$, $(9, 10)$, $(5, 7)$, $(2, 3)$, $(4, 6)$, which achieves a distance sum of $4 + 7 + 10 + 5 + 7 = 33$.

## ideas
1. 没有想法～
2. 