# Tree Positioning Puzzle

## Problem Description

The tree is a non-oriented connected graph without cycles. In particular, there always are $n - 1$ edges in a tree with $n$ vertices.

The puzzle is to position the vertices at the points of the Cartesian plane with integral coordinates, so that the segments between the vertices connected by edges are parallel to the coordinate axes. Also, the intersection of segments is allowed only at their ends. Distinct vertices should be placed at different points.

Help Dasha to find any suitable way to position the tree vertices on the plane.

It is guaranteed that if it is possible to position the tree vertices on the plane without violating the condition which is given above, then you can do it by using points with integral coordinates which don't exceed $10^{18}$ in absolute value.

## Input

The first line contains single integer $n$ ($1 \leq n \leq 30$) — the number of vertices in the tree.

Each of next $n - 1$ lines contains two integers $u_i, v_i$ ($1 \leq u_i, v_i \leq n$) that mean that the $i$-th edge of the tree connects vertices $u_i$ and $v_i$.

It is guaranteed that the described graph is a tree.

## Output

If the puzzle doesn't have a solution then in the only line print "NO".

Otherwise, the first line should contain "YES". The next $n$ lines should contain the pair of integers $x_i, y_i$ ($|x_i|, |y_i| \leq 10^{18}$) — the coordinates of the point which corresponds to the $i$-th vertex of the tree.

If there are several solutions, print any of them.

## Examples

### Example 1
**Input:**
```
7
1 2
1 3
2 4
2 5
3 6
3 7
```

**Output:**
```
YES
0 0
1 0
0 1
2 0
1 -1
-1 1
0 2
```

### Example 2
**Input:**
```
6
1 2
2 3
2 4
2 5
2 6
```

**Output:**
```
NO
```

### Example 3
**Input:**
```
4
1 2
2 3
3 4
```

**Output:**
```
YES
3 3
4 3
5 3
6 3
```


## ideas
1. 如果有一个vertex的deg超过了4，那么就没有答案，否则是不是一定有答案呢？
2. 好像是的吧？
3. 以任何一个点为root，比如1
4. 然后处理它的第一颗子树，把它指向到右边（很远的地方）
5. 然后处理第二课颗子树，指向上方
6. 。。。
7. 然后考虑第一课子树的情况，它的第一课子树指向右方
8. 第二颗子树指向上方
9. 但是第三课子树，要指向下方
10. 所以每棵子树，需要指定一个范围（或者不指定也行，后续要能调整坐标）
11. 