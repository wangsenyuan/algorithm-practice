# Codeforces 348B - Apple Tree

## Problem Description

You are given a rooted tree with n vertices. In each leaf vertex there's a single integer — the number of apples in this vertex.

The weight of a subtree is the sum of all numbers in this subtree leaves. For instance, the weight of a subtree that corresponds to some leaf is the number written in the leaf.

A tree is balanced if for every vertex v of the tree all its subtrees, corresponding to the children of vertex v, are of equal weight.

Count the minimum number of apples that you need to remove from the tree (specifically, from some of its leaves) in order to make the tree balanced. Notice that you can always achieve the goal by just removing all apples.

## Input

The first line contains integer n (2 ≤ n ≤ 10⁵), showing the number of vertices in the tree. The next line contains n integers a₁, a₂, ..., aₙ (0 ≤ aᵢ ≤ 10⁸), aᵢ is the number of apples in the vertex number i. The number of apples in non-leaf vertices is guaranteed to be zero.

Then follow n - 1 lines, describing the tree edges. Each line contains a pair of integers xᵢ, yᵢ (1 ≤ xᵢ, yᵢ ≤ n, xᵢ ≠ yᵢ) — the vertices connected by an edge.

The vertices are indexed from 1 to n. Vertex 1 is the root.

## Output

Print a single integer — the minimum number of apples to remove in order to make the tree balanced.

## Note

Please, do not write the %lld specifier to read or write 64-bit integers in С++. It is preferred to use the cin, cout streams or the %I64d specifier.


### ideas
1. 假设处理完了u的所有的子树，现在要处理u，如果u不是平衡的，那么就必须再处理v
2. 假设把u的子树都处理完了，然后v是（为了保证v平衡）数量最少
3. 那么其他子树必须也减少（至少到v）但是这样子，是否能达到平衡呢？
4. 或者这样子，在一棵已经平衡的子树中，要怎么样操作，才能重新平衡？
5. 假设子树u的子树, v1, v2, v3... vi
6. 分别需要减少f(1), f(2),... f(i)个apple，才能保证它们各自平衡
7. 那么如果现在要让它们再平衡，那么减少的数量貌似，必须是f(1), f(2), ... f(i)的lcm？
8. 考虑u有3个分叉，那么假设它达到平衡后，如果要继续减少，那么减少的部分，那么必须是3的倍数（只考虑u的情况）
9. 这样子，它才能平均分配到3个子树中去减少，并再次平衡
10. lcm * 3(这是u需要减少的量)
11. 