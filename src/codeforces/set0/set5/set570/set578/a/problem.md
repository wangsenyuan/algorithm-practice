# A Problem about Polyline

There is a polyline going through points (0, 0) – (x, x) – (2x, 0) – (3x, x) – (4x, 0) – ... – (2kx, 0) – (2kx + x, x) – ....

We know that the polyline passes through the point (a, b). Find minimum positive value x such that it is true or determine that there is no such x.

## Input

Only one line containing two positive integers a and b (1 ≤ a, b ≤ 10^9).

## Output

Output the only line containing the answer. Your answer will be considered correct if its relative or absolute error doesn't exceed 10^(-9). If there is no such x then output -1 as the answer.

## Examples

### Example 1

**Input:**

```
3 1
```

**Output:**

```
1.000000000000
```

### Example 2

**Input:**

```
1 3
```

**Output:**

```
-1
```

### Example 3

**Input:**

```
4 1
```

**Output:**

```
1.250000000000
```

## Note

You can see following graphs for sample 1 and sample 3.


### ideas
1. x >= b must hold
2. 假设(a, b)是在上升阶段（包括顶点部分）
3. (2 * kx, 0) -> (2 * k * x + x, x)
4. (a - 2 * kx) = b
5. 2 * k * x = a - b
6. 从（a, b) 画一条45度斜线，和x轴相交的位置 x0 = a - b
7. 假设周期 = x, 那么 x0和x需要满足什么条件呢？
8. x >= b and x0 = 2 * k * x
9. x越小，k就越大，如果 x = b
10. k = x0 / (2 * b)
11. 