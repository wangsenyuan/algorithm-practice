# Problem

Crazy Town is a plane on which there are $n$ infinite line roads. Each road is defined by the equation $a_i x + b_i y + c_i = 0$, where $a_i$ and $b_i$ are not both equal to zero. The roads divide the plane into connected regions, possibly of infinite extent. Each such region is called a **block**. An **intersection** is a point where at least two different roads intersect.

Your home is in one block and you need to get to the University, which is in another block. In one step you can move from one block to another only if the length of their common border is nonzero (in particular, if two blocks are adjacent at one intersection but have no shared nonzero boundary segment, you cannot move from one to the other in one step).

Determine the minimum number of steps to get from your home to the block containing the university. It is guaranteed that neither your home nor the university lies on any road.

## Input

- The first line contains two space-separated integers $x_1$, $y_1$ ($-10^6 \le x_1, y_1 \le 10^6$) — the coordinates of your home.
- The second line contains two space-separated integers $x_2$, $y_2$ ($-10^6 \le x_2, y_2 \le 10^6$) — the coordinates of the university.
- The third line contains an integer $n$ ($1 \le n \le 300$) — the number of roads.
- The next $n$ lines each contain three space-separated integers $a_i$, $b_i$, $c_i$ ($-10^6 \le a_i, b_i, c_i \le 10^6$; $|a_i| + |b_i| > 0$) — the coefficients of the line $a_i x + b_i y + c_i = 0$ for the $i$-th road.

It is guaranteed that no two roads coincide, and that neither your home nor the university lies on any road.

## Output

Output the minimum number of steps.

## Examples

### Example 1

**Input:**

```
1 1
-1 -1
2
0 1 0
1 0 0
```

**Output:**

```
2
```

### Example 2

**Input:**

```
1 1
-1 -1
3
1 0 0
0 1 0
1 1 -3
```

**Output:**

```
2
```

## Note

Pictures for the samples are presented below (A is your home, B is the university; different blocks are filled with different colors).


### ideas
1. 假设A在B的上方，就是找有多少条线，在B的上方，且在A的下方？
2. 