# Problem Description

Sereja has an n × m rectangular table a, each cell of the table contains a zero or a number one. Sereja wants his table to meet the following requirement: each connected component of the same values forms a rectangle with sides parallel to the sides of the table. Rectangles should be filled with cells, that is, if a component form a rectangle of size h × w, then the component must contain exactly hw cells.

A connected component of the same values is a set of cells of the table that meet the following conditions:

- every two cells of the set have the same value;
- the cells of the set form a connected region on the table (two cells are connected if they are adjacent in some row or some column of the table);
- it is impossible to add any cell to the set unless we violate the two previous conditions.

Can Sereja change the values of at most k cells of the table so that the table met the described requirement? What minimum number of table cells should he change in this case?

## Input

The first line contains integers n, m and k (1 ≤ n, m ≤ 100; 1 ≤ k ≤ 10). Next n lines describe the table a: the i-th of them contains m integers ai1, ai2, ..., aim (0 ≤ ai,j ≤ 1) — the values in the cells of the i-th row.

## Output

Print -1, if it is impossible to meet the requirement. Otherwise, print the minimum number of cells which should be changed.

## Examples

### Example 1

**Input:**

```text
5 5 2
1 1 1 1 1
1 1 1 1 1
1 1 0 1 1
1 1 1 1 1
1 1 1 1 1
```

**Output:**

```text
1
```

### Example 2

**Input:**

```text
3 4 1
1 0 0 0
0 1 1 1
1 1 1 0
```

**Output:**

```text
-1
```

### Example 3

**Input:**

```text
3 4 1
1 0 0 1
0 1 1 0
1 0 0 1
```

**Output:**

```text
0
```

## Ideas

1. 修改次数最少的情况下，组成一个chessboard
2. 没法dp啊。 状态是 C(n, k)的
3. 假如分成多片，每个区域的h, w 不超过20？
4. 最后组合出来？
5. 不会～

## Solution

**Key Observation:**

If we have two arrays x[1..n], 0 ≤ xi ≤ 1 and y[1..m], 0 ≤ yi ≤ 1, then the described matrix can be represented as: ai,j = xi XOR yj.

**Algorithm:**

- If n ≤ k, then we can backtrack array x and use greedy to find the best y.
- Otherwise, there will be at least one i such that we will not change any cell in row number i. So we can simply bruteforce some row and use it as x. Then we use greedy to find y.
- From all possible rows, we choose the most optimal one. Such a row will exist because the number of mistakes is lower than the number of rows, so it isn't possible to have at least one mistake in each row.

**Greedy Algorithm:**

For every element of y, we determine whether it's better to choose 0 or 1:

1. Count the number of different bits between x and the current column j
2. If the number of differences is lower than the count of same cells, set yj = 0
3. Otherwise, set yj = 1
