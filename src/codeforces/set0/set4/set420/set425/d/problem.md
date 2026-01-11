# Problem Description

Sereja has painted n distinct points on the plane. The coordinates of each point are integers. Now he is wondering: how many squares are there with sides parallel to the coordinate axes and with points painted in all its four vertexes? Help him, calculate this number.

## Input

The first line contains integer n (1 ≤ n ≤ 10^5). Each of the next n lines contains two integers xi, yi (0 ≤ xi, yi ≤ 10^5), the integers represent the coordinates of the i-th point. It is guaranteed that all the given points are distinct.

## Output

In a single line print the required number of squares.

## Examples

### Example 1

**Input:**
```
5
0 0
0 2
2 0
2 2
1 1
```

**Output:**
```
1
```

### Example 2

**Input:**
```
9
0 0
1 1
2 2
0 1
1 0
0 2
2 0
1 2
2 1
```

**Output:**
```
5
```


### ideas
1. 先排序
2. 想不到少于 n * n 的算法呐～
3. 假设在处理(x1, y1)的时候，能够把(x1后面，y1上面)距离到(x1, y1)相等的点的距离记录下来
4. 