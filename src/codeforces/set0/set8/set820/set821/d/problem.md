# Problem Description

Okabe likes to be able to walk through his city on a path lit by street lamps. That way, he doesn't get beaten up by schoolchildren.

Okabe's city is represented by a 2D grid of cells. Rows are numbered from 1 to n from top to bottom, and columns are numbered 1 to m from left to right. Exactly k cells in the city are lit by a street lamp. It's guaranteed that the top-left cell is lit.

Okabe starts his walk from the top-left cell, and wants to reach the bottom-right cell. Of course, Okabe will only walk on lit cells, and he can only move to adjacent cells in the up, down, left, and right directions. However, Okabe can also temporarily light all the cells in any single row or column at a time if he pays 1 coin, allowing him to walk through some cells not lit initially.

Note that Okabe can only light a single row or column at a time, and has to pay a coin every time he lights a new row or column. To change the row or column that is temporarily lit, he must stand at a cell that is lit initially. Also, once he removes his temporary light from a row or column, all cells in that row/column not initially lit are now not lit.

Help Okabe find the minimum number of coins he needs to pay to complete his walk!

## Input

The first line of input contains three space-separated integers n, m, and k (2 ≤ n, m, k ≤ 10^4).

Each of the next k lines contains two space-separated integers ri and ci (1 ≤ ri ≤ n, 1 ≤ ci ≤ m) — the row and the column of the i-th lit cell.

It is guaranteed that all k lit cells are distinct. It is guaranteed that the top-left cell is lit.

## Output

Print the minimum number of coins Okabe needs to pay to complete his walk, or -1 if it's not possible.

## Examples

### Example 1

**Input:**

```text
4 4 5
1 1
2 1
2 3
3 3
4 3
```

**Output:**

```text
2
```

### Example 2

**Input:**

```text
5 5 4
1 1
2 1
3 1
3 2
```

**Output:**

```text
-1
```

### Example 3

**Input:**

```text
2 2 4
1 1
1 2
2 1
2 2
```

**Output:**

```text
0
```

### Example 4

**Input:**

```text
5 5 4
1 1
2 2
3 3
4 4
```

**Output:**

```text
3
```

## Note

In the first sample test, Okabe can take the path, paying only when moving to (2, 3) and (4, 4).

In the fourth sample, Okabe can take the path, paying when moving to (1, 2), (3, 4), and (5, 4).



### solution
First, let's make this problem into one on a graph. The important piece of information is the row and column we're on, so we'll create a node like this for every lit cell in the grid. Edges in the graph are 0 between 2 nodes if we can reach the other immediately, or 1 if we can light a row/column to get to it. Now it's a shortest path problem: we need to start from a given node, and with minimum distance, reach another node.

Only problem is, number of edges can be large, causing the algorithm to time out. There are a lot of options here to reduce number of transitions. The most elegant one I found is Benq's solution, which I'll describe here. From a given cell, you can visit any adjacent lit cells. In addition, you can visit any lit cell with difference in rows at most 2, and any lit cell with difference in columns at most 2. So from the cell (r,c), you can just loop over all those cells.

The only tricky part is asking whether the current lit row/column should be a part of our BFS state. Since we fill the entire row/col and can then visit anything on that row/col, it doesn't matter where we came from. This means that you can temporarily light each row/column at most once during the entire BFS search.

So complexity is O(n + m + k), with a log factor somewhere for map or priority queue. Interestingly enough, you can remove the priority queue log factor because the BFS is with weights 0 and 1 only, but it performs slower in practice.

You can see the code implementing this approach below.

Benq's code:
Another approach to this problem was using "virtual nodes". Virtual nodes are an easy way to put transitions between related states while keeping number of edges low. In this problem, we can travel to any lit cell if its row differs by <=2, or its column differs by at most 2, but naively adding edges would cause O(k^2) edges.

Instead, for every row, lets make a virtual node. For every lit cell in this row, put an edge between the lit cell and this virtual node with cost 1. We can do something similar for every column.

Now, it's easy to see that the shortest path in this graph suffices. A minor detail is that we should divide the answer by 2 since every skipping of a row or column ends up costing 2 units of cost.