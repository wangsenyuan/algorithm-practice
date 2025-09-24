# Problem E: Tree Construction

## Story

Iahub and Iahubina went to a picnic in a forest full of trees. Less than 5 minutes passed before Iahub remembered of trees from programming. Moreover, he invented a new problem and Iahubina has to solve it, otherwise Iahub won't give her the food.

## Problem Statement

Iahub asks Iahubina: can you build a rooted tree, such that:

- Each internal node (a node with at least one son) has at least two sons
- Node i has ci nodes in its subtree

Iahubina has to guess the tree. Being a smart girl, she realized that it's possible no tree can follow Iahub's restrictions. In this way, Iahub will eat all the food. You need to help Iahubina: determine if there's at least one tree following Iahub's restrictions. The required tree must contain n nodes.

## Input

The first line of the input contains integer n (1 ≤ n ≤ 24). Next line contains n positive integers: the i-th number represents ci (1 ≤ ci ≤ n).

## Output

Output on the first line "YES" (without quotes) if there exist at least one tree following Iahub's restrictions, otherwise output "NO" (without quotes).

## Examples

### Example 1

**Input:**
```
4
1 1 1 4
```

**Output:**
```
YES
```

### Example 2

**Input:**
```
5
1 1 5 2 1
```

**Output:**
```
NO
```

## Solution Ideas

1. 假设存在，那么c[i]最大的肯定是root
2. 然后它必须至少有2个sons， 那么肯定第一个大的c[?]肯定是它的第一个son
3. 假设是c[j]（但是第2个就不一定了）但是不管怎么样， c[j] 是一个子问题
4. 然后，在剩余的n-2节点中，选择c[j]个节点，去组成，一个子树，如果可行，继续处理 
5. 因为每个节点至少有2个内部节点，所以叶子节点的数量 >= n / 2,至少有一半是叶子节点
6. 考虑这样一种情况， 5 = 1 + 1 + 1 + 1 + (1自己)， 也可以是 5 = 3 + 1 + (1)
7. 当多个更小的节点，合并成一个更大的节点的时候，其实是产生了一个额外的节点
8. 考虑1, 1, 1, 1, 1, 1, 1..... 10 (root = 10, 9个叶子节点)
9. 所有的叶子节点都接到root上，然后如果有一个节点1变成了3, 两个叶子节点接到它上面， 然后它在接到root
10. 如果再出现一个节点5，如果有足够多的（这个时候，不能用原来的两个1）
11. 应该找最小的节点去合并