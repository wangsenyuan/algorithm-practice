# Problem G: Coprime Sequence

## Problem Description

Let's denote as $L(x, p)$ an infinite sequence of integers $y$ such that $\gcd(p, y) = 1$ and $y > x$ (where $\gcd$ is the greatest common divisor of two integer numbers), sorted in ascending order. The elements of $L(x, p)$ are 1-indexed.

**Example:** For $L(7, 22)$, the first three elements are:
- 1st element: 9
- 2nd element: 13  
- 3rd element: 15

## Task

You have to process $t$ queries. Each query is denoted by three integers $x$, $p$ and $k$, and the answer to this query is the $k$-th element of $L(x, p)$.

## Input

- **First line:** One integer $t$ $(1 \leq t \leq 30000)$ — the number of queries to process.
- **Next $t$ lines:** Each line contains three integers $x$, $p$ and $k$ for the $i$-th query $(1 \leq x, p, k \leq 10^6)$.

## Output

Print $t$ integers, where the $i$-th integer is the answer to the $i$-th query.

## Examples

### Example 1
**Input:**
```
3
7 22 1
7 22 2
7 22 3
```

**Output:**
```
9
13
15
```

**Explanation:** 
- Query 1: Find the 1st element of $L(7, 22)$ → 9
- Query 2: Find the 2nd element of $L(7, 22)$ → 13
- Query 3: Find the 3rd element of $L(7, 22)$ → 15

### Example 2
**Input:**
```
5
42 42 42
43 43 43
44 44 44
45 45 45
46 46 46
```

**Output:**
```
187
87
139
128
141
```

## Notes

- The sequence $L(x, p)$ contains all integers $y > x$ that are coprime with $p$
- Elements are sorted in ascending order
- Queries are 1-indexed
- All input values are within the range $[1, 10^6]$

## ideas
1. 对于L(x, p), 假设一个数字y， 那么有多少数字 z <= y， 且 gcd(p, z) > 1 呢？
2. 那么 y - x + 1 - count(z <= y) = k
3. 假设p的质因数arr, 那么可以用inclusion-exclusion来出来
4. 2 * 3 * 5 * 7 * 11 * 13 * 17 * 19
5. 最多有8个不同的质因数