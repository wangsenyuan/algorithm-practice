# Problem

You've got array $a[1], a[2], \ldots, a[n]$, consisting of $n$ integers. Count the number of ways to split all the elements of the array into three contiguous parts so that the sum of elements in each part is the same.

More formally, you need to find the number of such pairs of indices $i, j$ ($2 \leq i \leq j \leq n - 1$), that $\sum_{k=1}^{i-1} a[k] = \sum_{k=i}^{j} a[k] = \sum_{k=j+1}^{n} a[k]$.

## Input

The first line contains integer $n$ ($1 \leq n \leq 5 \cdot 10^5$), showing how many numbers are in the array. The second line contains $n$ integers $a[1], a[2], \ldots, a[n]$ ($|a[i]| \leq 10^9$) — the elements of array $a$.

## Output

Print a single integer — the number of ways to split the array into three parts with the same sum.

## Examples

### Example 1

**Input:**

```text
5
1 2 3 0 3
```

**Output:**

```text
2
```

### Example 2

**Input:**

```text
4
0 1 -1 0
```

**Output:**

```text
1
```

### Example 3

**Input:**

```text
2
4 1
```

**Output:**

```text
0
```
