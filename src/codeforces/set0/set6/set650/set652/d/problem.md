# Problem: Nested Segments

You are given n segments on a line. There are no ends of some segments that coincide. For each segment find the number of segments it contains.

## Input

The first line contains a single integer n (1 ≤ n ≤ 2·10⁵) — the number of segments on a line.

Each of the next n lines contains two integers li and ri (-10⁹ ≤ li < ri ≤ 10⁹) — the coordinates of the left and the right ends of the i-th segment. It is guaranteed that there are no ends of some segments that coincide.

## Output

Print n lines. The j-th of them should contain the only integer aj — the number of segments contained in the j-th segment.

## Examples

### Example 1
**Input:**
```
4
1 8
2 3
4 7
5 6
```

**Output:**
```
3
0
1
0
```

### Example 2
**Input:**
```
3
3 4
1 5
2 6
```

**Output:**
```
0
1
1
```


### ideas
1. 按照端点处理，在右节点的时候，计算有多少个左端点在它的中间