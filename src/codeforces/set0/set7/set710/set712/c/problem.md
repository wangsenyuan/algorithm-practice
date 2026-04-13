# C. Memory and De-Evolution

Memory is now interested in the **de-evolution** of objects, specifically **triangles**. He starts with an **equilateral** triangle of side length `x`, and he wishes to perform operations to obtain an equilateral triangle of side length `y`.

In a **single second**, he can modify the length of **a single side** of the current triangle so that it remains a **non-degenerate** triangle (positive area). At every moment, each side length must be an **integer**.

What is the **minimum number of seconds** required for Memory to obtain the equilateral triangle of side length `y`?

## Input

The first and only line contains two integers `x` and `y` (`3 ≤ y < x ≤ 100000`) — the starting and target equilateral side lengths.

## Output

Print a single integer: the minimum number of seconds needed to reach an equilateral triangle of side length `y` starting from side length `x`.

## Examples

### Example 1

**Input**

```
6 3
```

**Output**

```
4
```

### Example 2

**Input**

```
8 5
```

**Output**

```
3
```

### Example 3

**Input**

```
22 4
```

**Output**

```
6
```

## Note

In the first sample, Memory starts from side length `6` and aims for `3`. If we write a triangle as `(a, b, c)` for its three side lengths, an optimal sequence uses **four** seconds (each second changes one side by an integer amount while keeping integer sides and positive area). The second and third samples are answered in **3** and **6** seconds respectively.

### ideas
1. 如果 2 * y > x => 3
2. 2 * y <= x, (x, x, x) => (x, x, y) => y + z > x => z > x - y
3. (x, x - y + 1, y), y + z > x - y + 1 => y > x - 2 * y + 1
4. 