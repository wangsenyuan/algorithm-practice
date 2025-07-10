# Problem D: Complete Binary Tree Navigation

## Problem Description

T is a complete binary tree consisting of n vertices. It means that exactly one vertex is a root, and each vertex is either a leaf (and doesn't have children) or an inner node (and has exactly two children). All leaves of a complete binary tree have the same depth (distance from the root). So n is a number such that n + 1 is a power of 2.

In the picture you can see a complete binary tree with n = 15.

## Vertex Numbering

Vertices are numbered from 1 to n in a special recursive way:
- We recursively assign numbers to all vertices from the left subtree (if current vertex is not a leaf)
- Then assign a number to the current vertex
- Then recursively assign numbers to all vertices from the right subtree (if it exists)

In the picture vertices are numbered exactly using this algorithm. It is clear that for each size of a complete binary tree exists exactly one way to give numbers to all vertices. This way of numbering is called **symmetric**.

## Task

You have to write a program that for given n answers q queries to the tree.

Each query consists of:
- An integer number ui (1 ≤ ui ≤ n) - the number of vertex
- A string si - represents the path starting from this vertex

String si doesn't contain any characters other than:
- 'L' - traverse to the left child
- 'R' - traverse to the right child  
- 'U' - traverse to the parent

Characters from si have to be processed from left to right, considering that ui is the vertex where the path starts. If it's impossible to process a character (for example, to go to the left child of a leaf), then you have to skip it.

**The answer is the number of vertex where the path represented by si ends.**

### Example
If ui = 4 and si = "UURL", then the answer is 10.

## Input

- The first line contains two integer numbers n and q (1 ≤ n ≤ 10^18, q ≥ 1)
- n is such that n + 1 is a power of 2
- The next 2q lines represent queries; each query consists of two consecutive lines:
  - First line: ui (1 ≤ ui ≤ n)
  - Second line: non-empty string si (contains only 'L', 'R', 'U')
- It is guaranteed that the sum of lengths of si (for each i such that 1 ≤ i ≤ q) doesn't exceed 10^5

## Output

Print q numbers, i-th number must be the answer to the i-th query.

## Example

### Input
```
15 2
4
UURL
8
LRLLLLLLLL
```

### Output
```
10
5
```

### ideas
1. 如果u是一个左子树，往左边走，相当于除以2，向上走，相当于乘2，（向右边走，等会再分析）
2. 如果u是一个右子树，这时候需要知道它的parent的节点编号（假设是x）
3. 那么向左边走 = (u - x) / 2, 向上走 = x, 
4. 往右边的时候，x + (x - 1) / 2 + 1