# Problem Description

Today at the lesson Vitya learned a very interesting function — **mex**. Mex of a sequence of numbers is the minimum non-negative number that is not present in the sequence as an element. For example, $\text{mex}([4, 33, 0, 1, 1, 5]) = 2$ and $\text{mex}([1, 2, 3]) = 0$.

Vitya quickly understood all tasks of the teacher, but can you do the same?

You are given an array consisting of $n$ non-negative integers, and $m$ queries. Each query is characterized by one number $x$ and consists of the following consecutive steps:

1. Perform the bitwise addition operation modulo 2 (XOR) of each array element with the number $x$.
2. Find mex of the resulting array.

Note that after each query the array changes.

## Input

First line contains two integer numbers $n$ and $m$ ($1 \leq n, m \leq 3 \cdot 10^5$) — number of elements in array and number of queries.

Next line contains $n$ integer numbers $a_i$ ($0 \leq a_i \leq 3 \cdot 10^5$) — elements of the array.

Each of next $m$ lines contains a query — one integer number $x$ ($0 \leq x \leq 3 \cdot 10^5$).

## Output

For each query print the answer on a separate line.

## Examples

### Example 1

**Input:**

```text
2 2
1 3
1
3
```

**Output:**

```text
1
0
```

### Example 2

**Input:**

```text
4 3
0 1 5 6
1
2
4
```

**Output:**

```text
2
0
0
```

### Example 3

**Input:**

```text
5 4
0 1 5 6 7
1
1
4
5
```

**Output:**

```text
2
2
0
2
```
