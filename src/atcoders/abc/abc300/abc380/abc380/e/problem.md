# Problem E

There are N cells in a row, numbered 1 to N.
For each 1≤i<N, cells i and i+1 are adjacent.

Initially, cell i is painted with color i.

You are given Q queries. Process them in order. Each query is of one of the following two types.

1 x c: Repaint the following to color c: all reachable cells reachable from cell x by repeatedly moving to an adjacent cell painted in the same color as the current cell.
2 c: Print the number of cells painted with color c.

## Constraints

- 1≤N≤5×10⁵
- 1≤Q≤2×10⁵
- In queries of the first type, 1≤x≤N.
- In queries of the first and second types, 1≤c≤N.
- There is at least one query of the second type.
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```text
N Q
query₁
⋮
queryQ
```

Each query is given in one of the following two formats:

```text
1 x c
2 c
```

## Output

Let q be the number of queries of the second type. Print q lines.

The i-th line should contain the answer to the i-th such query.

## Sample Input 1

```text
5 6
1 5 4
1 4 2
2 2
1 3 2
1 2 3
2 3
```

## Sample Output 1

```text
3
4
```
