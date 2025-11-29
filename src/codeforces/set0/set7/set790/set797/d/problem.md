Let T be arbitrary binary tree — tree, every vertex of which has no more than two children. Given tree is rooted, so there exists only one vertex which doesn't have a parent — it's the root of a tree. Every vertex has an integer number written on it. Following algorithm is run on every value from the tree T:

1. Set pointer to the root of a tree.
2. Return success if the value in the current vertex is equal to the number you are looking for
3. Go to the left child of the vertex if the value in the current vertex is greater than the number you are looking for
4. Go to the right child of the vertex if the value in the current vertex is less than the number you are looking for
5. Return fail if you try to go to the vertex that doesn't exist

Here is the pseudo-code of the described algorithm:

```cpp
bool find(TreeNode t, int x) {
    if (t == null)
        return false;
    if (t.value == x)
        return true;
    if (x < t.value)
        return find(t.left, x);
    else
        return find(t.right, x);
}
find(root, x);
```

The described algorithm works correctly if the tree is binary search tree (i.e. for each node the values of left subtree are less than the value in the node, the values of right subtree are greater than the value in the node). But it can return invalid result if tree is not a binary search tree.

Since the given tree is not necessarily a binary search tree, not all numbers can be found this way. Your task is to calculate, how many times the search will fail being running on every value from the tree.

If the tree has multiple vertices with the same values on them then you should run algorithm on every one of them separately.

## Input

First line contains integer number n (1 ≤ n ≤ 105) — number of vertices in the tree.

Each of the next n lines contains 3 numbers v, l, r (0 ≤ v ≤ 109) — value on current vertex, index of the left child of the vertex and index of the right child of the vertex, respectively. If some child doesn't exist then number -1 is set instead. Note that different vertices of the tree may contain the same values.

## Output

Print number of times when search algorithm will fail.

## Examples

### Example 1

**Input:**
```
3
15 -1 -1
10 1 3
5 -1 -1
```

**Output:**
```
2
```

### Example 2

**Input:**
```
8
6 2 3
3 4 5
12 6 7
1 -1 8
4 -1 -1
5 -1 -1
14 -1 -1
2 -1 -1
```

**Output:**
```
1
```

## Note

In the example the root of the tree in vertex 2. Search of numbers 5 and 15 will return fail because on the first step algorithm will choose the subtree which doesn't contain numbers you are looking for.


### ideas 
1. 有意思的题目
2. 对于node v，如果要正确的找到它，它必须在正确的路径上
3. 如果v的parent是u，即使u是在错误的位置，v有可能还是正确的
4. 对于v来说，如果它在某些位置向左的，那么这些地方的值 > a[v]
5. 所以，要维护两个队列，一个是从左来的parent，一个是从有来的parent
6. 不大对～
7. 正常情况下，要找到这个v，在整个路径上，所有比它大的数，都应该是往左运行的位置
8. 所有比它小的数，都是往右运动的位置