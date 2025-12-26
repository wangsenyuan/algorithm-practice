# Problem Description

Bogdan has a birthday today and mom gave him a tree consisting of n vertecies. For every edge of the tree i, some number xi was written on it. In case you forget, a tree is a connected non-directed graph without cycles. After the present was granted, m guests consecutively come to Bogdan's party. When the i-th guest comes, he performs exactly one of the two possible operations:

1. Chooses some number yi, and two vertecies ai and bi. After that, he moves along the edges of the tree from vertex ai to vertex bi using the shortest path (of course, such a path is unique in the tree). Every time he moves along some edge j, he replaces his current number yi by ⌊yi / xj⌋, that is, by the result of integer division yi div xj.
2. Chooses some edge pi and replaces the value written in it xpi by some positive integer ci < xpi.

As Bogdan cares about his guests, he decided to ease the process. Write a program that performs all the operations requested by guests and outputs the resulting value yi for each i of the first type.

## Input

The first line of the input contains integers, n and m (2 ≤ n ≤ 200 000, 1 ≤ m ≤ 200 000) — the number of vertecies in the tree granted to Bogdan by his mom and the number of guests that came to the party respectively.

Next n - 1 lines contain the description of the edges. The i-th of these lines contains three integers ui, vi and xi (1 ≤ ui, vi ≤ n, ui ≠ vi, 1 ≤ xi ≤ 10^18), denoting an edge that connects vertecies ui and vi, with the number xi initially written on it.

The following m lines describe operations, requested by Bogdan's guests. Each description contains three or four integers and has one of the two possible forms:

- `1 ai bi yi` corresponds to a guest, who chooses the operation of the first type.
- `2 pi ci` corresponds to a guests, who chooses the operation of the second type.

It is guaranteed that all the queries are correct, namely 1 ≤ ai, bi ≤ n, 1 ≤ pi ≤ n - 1, 1 ≤ yi ≤ 10^18 and 1 ≤ ci < xpi, where xpi represents a number written on edge pi at this particular moment of time that is not necessarily equal to the initial value xpi, as some decreases may have already been applied to it. The edges are numbered from 1 to n - 1 in the order they appear in the input.

## Output

For each guest who chooses the operation of the first type, print the result of processing the value yi through the path from ai to bi.

## Examples

### Example 1

**Input:**

```text
6 6
1 2 1
1 3 7
1 4 4
2 5 5
2 6 2
1 4 6 17
2 3 2
1 4 6 17
1 5 5 20
2 4 1
1 5 1 3
```

**Output:**

```text
2
4
20
3
```

### Example 2

**Input:**

```text
5 4
1 2 7
1 3 3
3 4 2
3 5 5
1 4 2 100
1 5 4 1
2 2 2
1 1 3 4
```

**Output:**

```text
2
0
2
```

### ideas
1. 只需要记录那些不为1的边，那么每次计算操作，就不会超过100次