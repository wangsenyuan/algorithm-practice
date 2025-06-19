# Problem D: Full Binary Tree Queries

## Problem Description

You have a full binary tree having infinite levels.

Each node has an initial value. If a node has value `x`, then its left child has value `2·x` and its right child has value `2·x + 1`.

The value of the root is `1`.

You need to answer `Q` queries.

## Query Types

There are 3 types of queries:

1. **Type 1**: Cyclically shift the values of all nodes on the same level as node with value `X` by `K` units. (The values/nodes of any other level are not affected).

2. **Type 2**: Cyclically shift the nodes on the same level as node with value `X` by `K` units. (The subtrees of these nodes will move along with them).

3. **Type 3**: Print the value of every node encountered on the simple path from the node with value `X` to the root.

**Note**: Positive `K` implies right cyclic shift and negative `K` implies left cyclic shift.

It is guaranteed that at least one type 3 query is present.

## Input

- The first line contains a single integer `Q` (1 ≤ Q ≤ 10⁵).
- Then `Q` queries follow, one per line:
  - **Queries of type 1 and 2**: `T X K` (1 ≤ T ≤ 2; 1 ≤ X ≤ 10¹⁸; 0 ≤ |K| ≤ 10¹⁸), where `T` is the type of the query.
  - **Queries of type 3**: `3 X` (1 ≤ X ≤ 10¹⁸).

## Output

For each query of type 3, print the values of all nodes encountered in descending order.

## ideas
1. X所在的层级不会改变， 所以可以计算出X所在的层级H
2. 如果是操作1，只是这一层数字的shift，对于节点4, 5, 6, 7假设原始的是 4, 5, 6, 7, right shift 1后，变成了, [7, 4, 5, 6]
3. 但是4的parent还是2(值是7)
4. 如果是操作2，就比较麻烦了。需要知道旋转到的新位置，然后从该位置计算
5. 如果先操作1，再操作2呢，再操作1
6. 这些数字的相对顺序不会变，在操作2的时候，计算新的头部节点是什么
7. 如果这时候进行了操作1，计算一个offset
8. 如果再进行操作2，offset是不变的
9. 两个操作是独立的？
10. 始终是按照位置去找parent，但是要计算出当前位置的值