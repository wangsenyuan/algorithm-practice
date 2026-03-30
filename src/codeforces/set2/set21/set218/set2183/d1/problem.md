You are given a rooted tree∗
 consisting of 𝑛
 vertices numbered from 1
 to 𝑛
, where the root has index 1
, and each vertex is initially white. Define 𝑑𝑖
 as the distance from the root to the 𝑖
-th vertex. You can perform the following operations any number of times:

Select a subset 𝑆
 of white vertices such that no two nodes in 𝑆
 are connected by an edge, or have the same distance to node 1
. Formally, 𝑆
 should satisfy for all 𝑥,𝑦∈𝑆
 and 𝑥≠𝑦
, 𝑑𝑥≠𝑑𝑦
 and there are no edges between 𝑥
 and 𝑦
.
Color the vertices in 𝑆
 black.
Your job is to find the minimum number of operations required to color the full tree black.
∗
A tree is a connected graph without cycles. A rooted tree is a tree where one vertex is special and called the root.

## Input
Each test contains multiple test cases. The first line contains the number of test cases 𝑡
 (1≤𝑡≤104
). The description of the test cases follows.

The first line of each test case contains a single integer 𝑛
 (2≤𝑛≤2⋅105
) – the number of vertices in the tree.

The 𝑖
-th following 𝑛−1
 lines contain two integers 𝑢𝑖
 and 𝑣𝑖
 (1≤𝑢𝑖,𝑣𝑖≤𝑛
, 𝑢𝑖≠𝑣𝑖
) – the ends of the 𝑖
-th edge.

It is guaranteed that the given edges form a tree.

It is guaranteed that the sum of 𝑛
 over all test cases doesn't exceed 2⋅105
.

## Output
For each test case, output the minimum number of operations on a new line.

## Example
**Input**

```
10
5
3 1
1 2
5 1
4 1
5
3 2
2 4
2 5
1 2
5
3 4
4 1
5 1
1 2
5
2 5
3 1
2 1
3 4
5
1 3
1 5
4 3
2 4
13
2 1
3 2
4 2
5 4
6 3
7 1
8 5
9 6
10 4
11 7
12 8
13 10
10
5 7
8 1
1 10
2 8
8 4
9 4
6 1
5 3
7 8
10
7 6
3 7
6 9
7 1
9 8
5 1
3 10
9 2
1 4
10
10 6
2 8
4 10
7 5
1 2
7 10
10 9
9 1
7 3
10
6 8
9 7
4 10
5 9
4 2
3 8
6 5
1 5
1 10
```

**Output**

```
5
4
4
3
3
3
4
4
4
3
```

## Note
In the first test case, 𝑑1=1
 and 𝑑2=𝑑3=𝑑4=𝑑5=2
. We can show that we must perform at least 5
 operations because there are no two nodes that can be operated on simultaneously.

In the second test case, we can show that the least number of operations required to color the full tree is 4
. One way to do so is to color nodes 1
 and 3
 in the same operation, and all 3
 other nodes in their own operations.