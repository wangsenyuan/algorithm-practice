 # Black Boar Painting Problem

## Problem Description

I see a pink boar and I want it painted black. Black boars look much more awesome and mighty than the pink ones. Since Jaggy became the ruler of the forest, he has been trying his best to improve the diplomatic relations between the forest region and the nearby ones.

Some other rulers, however, have requested too much in return for peace between their two regions, so he realized he has to resort to intimidation. Once a delegate for diplomatic relations of a neighboring region visits Jaggy's forest, if they see a whole bunch of black boars, they might suddenly change their mind about attacking Jaggy. Black boars are really scary, after all.

Jaggy's forest can be represented as a tree (connected graph without cycles) with n vertices. Each vertex represents a boar and is colored either black or pink. Jaggy has sent a squirrel to travel through the forest and paint all the boars black. The squirrel, however, is quite unusually trained and while it traverses the graph, it changes the color of every vertex it visits, regardless of its initial color: pink vertices become black and black vertices become pink.

Since Jaggy is too busy to plan the squirrel's route, he needs your help. He wants you to construct a walk through the tree starting from vertex 1 such that in the end all vertices are black. A walk is a sequence of vertices, such that every consecutive pair has an edge between them in a tree.

## Input

The first line of input contains integer n (2 ≤ n ≤ 200,000), denoting the number of vertices in the tree.

The following n lines contains n integers, which represent the color of the nodes:
- If the i-th integer is 1, the i-th vertex is black
- If the i-th integer is -1, the i-th vertex is pink

Each of the next n - 1 lines contains two integers, which represent the indexes of the vertices which are connected by the edge. Vertices are numbered starting with 1.

## Output

Output path of a squirrel: output a sequence of visited nodes' indexes in order of visiting. In case of all the nodes are initially black, you should print 1. Solution is guaranteed to exist. If there are multiple solutions to the problem you can output any of them provided length of sequence is not longer than 10^7.

## Example

**Input:**
```
5
1
1
-1
1
-1
2 5
4 3
2 4
4 1
```

**Output:**
```
1 4 2 5 2 4 3 4 1 4 1
```

**Explanation:**
At the beginning squirrel is at node 1 and its color is black. Next steps are as follows:

1. From node 1 we walk to node 4 and change its color to pink.
2. From node 4 we walk to node 2 and change its color to pink.
3. From node 2 we walk to node 5 and change its color to black.
4. From node 5 we return to node 2 and change its color to black.
5. From node 2 we walk to node 4 and change its color to black.
6. We visit node 3 and change its color to black.
7. We visit node 4 and change its color to pink.
8. We visit node 1 and change its color to pink.
9. We visit node 4 and change its color to black.
10. We visit node 1 and change its color to black.

### ideas
1. 如果一个vertex是pink的，那么它需要被访问奇数次
2. 如果是black的，那么需要被访问偶数次
3. 假设一个节点u，有一个子节点v，它是pink的，那么必须访问它u，然后访问v
4. 如果u是black的，那么此时它变成了pink，回来的时候，就变成了black
5. 如果u原来是pink的，那么它就变成了black，回来时，变成了pink。所以它就是那个新的v
6. 保证离开u的时候，u必须是black的
7. 如果离开u的时候，它是pink的，先进入parent节点，然后返回，然后离开
8. 但是，可能存在到最后时刻1是pink的
9. 比如两个节点，1-2， 都是pink的
10. 1 -> 2 -> 1
11. 如果 1是pink的，但是2是black的, 1 -> 2 -> 1 -> 2
12. 所以是ok的