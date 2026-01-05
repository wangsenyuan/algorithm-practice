# Problem Description

You are given a functional graph. It is a directed graph, in which from each vertex goes exactly one arc. The vertices are numerated from 0 to n - 1.

Graph is given as the array f₀, f₁, ..., fₙ₋₁, where fᵢ — the number of vertex to which goes the only arc from the vertex i. Besides you are given array with weights of the arcs w₀, w₁, ..., wₙ₋₁, where wᵢ — the arc weight from i to fᵢ.

The graph from the first sample test.

Also you are given the integer k (the length of the path) and you need to find for each vertex two numbers sᵢ and mᵢ, where:

- sᵢ — the sum of the weights of all arcs of the path with length equals to k which starts from the vertex i;
- mᵢ — the minimal weight from all arcs on the path with length k which starts from the vertex i.

The length of the path is the number of arcs on this path.

## Input

The first line contains two integers n, k (1 ≤ n ≤ 10⁵, 1 ≤ k ≤ 10¹⁰). The second line contains the sequence f₀, f₁, ..., fₙ₋₁ (0 ≤ fᵢ < n) and the third — the sequence w₀, w₁, ..., wₙ₋₁ (0 ≤ wᵢ ≤ 10⁸).

## Output

Print n lines, the pair of integers sᵢ, mᵢ in each line.

## Examples

### Example 1

**Input:**

```text
7 3
1 2 3 4 3 2 6
6 3 1 4 2 2 3
```

**Output:**

```text
10 1
8 1
7 1
10 2
8 2
7 1
9 3
```

### Example 2

**Input:**

```text
4 4
0 1 2 3
0 1 2 3
```

**Output:**

```text
0 0
4 1
8 2
12 3
```

### Example 3

**Input:**

```text
5 3
1 2 3 4 0
4 1 2 14 3
```

**Output:**

```text
7 1
17 1
19 2
21 3
8 1
```


### ideas
1. 所有的节点，最后都指向了一个圈
2. 可以把这个圈变成一个独立的节点，然后就形成了一棵树
3. 如果是这个圈上的节点，分两种情况，如果k < 这个圈的长度，就需要找到它能k长度的边，否则，就是这个圈上的sum *  k / .
4. 所以，这里有个关键的点， 需要从圈上的一个点u，找到距离k1的边的sum, 还有其中的最小值