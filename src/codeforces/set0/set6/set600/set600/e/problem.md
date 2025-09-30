# Problem E - Dominating Colours

## Problem Statement

You are given a rooted tree with root in vertex 1. Each vertex is coloured in some colour.

Let's call colour c **dominating** in the subtree of vertex v if there are no other colours that appear in the subtree of vertex v more times than colour c. So it's possible that two or more colours will be dominating in the subtree of some vertex.

The subtree of vertex v is the vertex v and all other vertices that contains vertex v in each path to the root.

For each vertex v find the sum of all dominating colours in the subtree of vertex v.

## Input

The first line contains integer n (1 ≤ n ≤ 10^5) — the number of vertices in the tree.

The second line contains n integers ci (1 ≤ ci ≤ n), ci — the colour of the i-th vertex.

Each of the next n - 1 lines contains two integers xj, yj (1 ≤ xj, yj ≤ n) — the edge of the tree. The first vertex is the root of the tree.

## Output

Print n integers — the sums of dominating colours for each vertex.

## Examples

### Example 1

**Input:**
```
4
1 2 3 4
1 2
2 3
2 4
```

**Output:**
```
10 9 3 4
```

### Example 2

**Input:**
```
15
1 2 3 1 2 3 3 1 1 3 2 2 1 2 3
1 2
1 3
1 4
1 14
1 15
2 5
2 6
2 7
3 8
3 9
3 10
4 11
4 12
4 13
```

**Output:**
```
6 5 4 3 2 3 3 1 1 3 2 2 1 2 3
```


### ideas
1. 考虑一条直线，从左到右，要怎么处理？
2. 假设当前节点的color为x, freq[x]++, 如果freq[x] > freq[old] (old是之前的)， 那么更新为x，如果 freq[x] = freq[old], +x, 否则 + old
3. 怎么快速的找到这个节点里面最多的颜色的个数呢？
4. 在u子树中， 一种颜色比如x,它可以不是某个子节点v中最多的那个，但是它可以是u最多的那个，比如x在v1中是3个，在v2中是2个，都少于各子树最多的那个，但是加起来却可能是最多的
5. 所以，即使算出了v的结果，它也无法直接作为u的结果来使用
6. 如果把整棵树排平，对于u，在它的进入和离开节点中间，包括所有的子节点，那么它内部节点的贡献 = 那些出现了两次的节点（一次进入，一次离开）+ 自己
7. 这样子似乎就可以用segment tree了？
8. 对于节点u，就是找一个区间内，出现最多的（且每个节点必须出现两次）
9. 不需要计算两次，计算一次就可以了。然后就可以用merge来实现