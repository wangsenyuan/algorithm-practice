You are given a directed graph G with n vertices and m arcs (multiple arcs and self-loops are allowed). You have to paint each vertex of the graph into one of the k (k ≤ n) colors in such way that for all arcs of the graph leading from a vertex u to vertex v, vertex v is painted with the next color of the color used to paint vertex u.

The colors are numbered cyclically 1 through k. This means that for each color i (i < k) its next color is color i + 1. In addition, the next color of color k is color 1. Note, that if k = 1, then the next color for color 1 is again color 1.

Your task is to find and print the largest possible value of k (k ≤ n) such that it's possible to color G as described above with k colors. Note that you don't necessarily use all the k colors (that is, for each color i there does not necessarily exist a vertex that is colored with color i).

### ideas
1. k=1时，肯定是成立的；且如果有self-loop，k只能等于1
2. 感觉和环的长度有关系
3. 如果没有环，那么答案，可以为n（最大的）
4. 如果有环的情况下，只有一个环，那么这个环的大小，就是k的最大值
5. 如果有多个环呢？这些环的gcd
6. 怎么找出所有环的大小？ dfs？
7. u如果有多条路径到v，那么k = gcd(这些路径的长度) 
8. 这个要怎么计算出来呢？ 