## Problem Statement

You are given an undirected graph that consists of n vertices and m edges. Initially, each edge is colored either red or blue.

Each turn a player picks a single vertex and switches the color of all edges incident to it. That is, all red edges with an endpoint in this vertex change the color to blue, while all blue edges with an endpoint in this vertex change the color to red.

Find the minimum possible number of moves required to make the colors of all edges equal.

## Input

The first line of the input contains two integers n and m (1 ≤ n, m ≤ 100 000) — the number of vertices and edges, respectively.

The following m lines provide the description of the edges, as the i-th of them contains two integers ui and vi (1 ≤ ui, vi ≤ n, ui ≠ vi) — the indices of the vertices connected by the i-th edge, and a character ci () providing the initial color of this edge. If ci equals 'R', then this edge is initially colored red. Otherwise, ci is equal to 'B' and this edge is initially colored blue. It's guaranteed that there are no self-loops and multiple edges.

## Output

If there is no way to make the colors of all edges equal output -1 in the only line of the output.

Otherwise first output k — the minimum number of moves required to achieve the goal, then output k integers a1, a2, ..., ak, where ai is equal to the index of the vertex that should be used at the i-th move.

If there are multiple optimal sequences of moves, output any of them.

## Examples

### Example 1

Input:
```
3 3
1 2 B
3 1 R
3 2 B
```

Output:
```
1
2
```

### Example 2

Input:
```
6 5
1 3 R
2 3 R
3 4 B
4 5 R
4 6 R
```

Output:
```
2
3 4
```

### Example 3

Input:
```
4 5
1 2 R
1 3 R
2 3 B
3 4 B
1 4 B
```

Output:
```
-1
```


### ideas
1. 一条边的两端都switch的话，对于这条边来讲，相当于没有变化
2. 所以，假设要把red变成blue，那么red边的一端要进行奇数次变化，另外一端要进行偶数次变化
3. 那么利用这个进行分组，只要能够让red边的两端，分到不同的组，blue的两端，分到同一组，如果能做到，就有答案
4. 所以，red边赋值为1，blue边赋值为0，然后检查能否2分
5. 这样子，只能是检查有效性，但是如何找到最小值呢？找那个小的comp去switch就好了
6. done～