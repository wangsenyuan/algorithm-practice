# Problem C

Katya studies in a fifth grade. Recently her class studied right triangles and the Pythagorean theorem. It appeared, that there are triples of positive integers such that you can construct a right triangle with segments of lengths corresponding to triple. Such triples are called Pythagorean triples.

For example, triples (3, 4, 5), (5, 12, 13) and (6, 8, 10) are Pythagorean triples.

Here Katya wondered if she can specify the length of some side of right triangle and find any Pythagorean triple corresponding to such length? Note that the side which length is specified can be a cathetus as well as hypotenuse.

Katya had no problems with completing this task. Will you do the same?

## Input

The only line of the input contains single integer n (1 ≤ n ≤ 10^9) — the length of some side of a right triangle.

## Output

Print two integers m and k (1 ≤ m, k ≤ 10^18), such that n, m and k form a Pythagorean triple, in the only line.

In case if there is no any Pythagorean triple containing integer n, print -1 in the only line. If there are many answers, print any of them.

## Examples

### Example 1

**Input:**

```text
3
```

**Output:**

```text
4 5
```

### Example 2

**Input:**

```text
6
```

**Output:**

```text
8 10
```

### Example 3

**Input:**

```text
1
```

**Output:**

```text
-1
```

### Example 4

**Input:**

```text
17
```

**Output:**

```text
144 145
```

### Example 5

**Input:**

```text
67
```

**Output:**

```text
2244 2245
```


### ideas
1. w = a, w = b, or w = c
2. 如果 w = a, 那么 w = (m2 - n2) * k
3. w = (m - n) * (m + n) * k
4. 那么m-n, m+n, k都是w的因子
5. 假设 u = m - n, v = m + n
6. u + v = 2 * m
7. v - u = 2 * n （所以，要相同parity的数）
8. 那么只要有这样的两个数，就能找到m, n(随便两个数，都可以)
9. 假设不存在这样的数， w = b = 2 * m * n
10. 那么 w必须是偶数，（貌似只要是偶数就行) 让 n = 1就可以了, m = w/2
11. 现在w是奇数，切不满足条件a
12. w = c = m * m + n * n（那就简单了）迭代所有的m * m 就可以了