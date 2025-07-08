# Tree Land Officer Assignment

## Problem Description

Now Fox Ciel becomes a commander of Tree Land. Tree Land, like its name said, has n cities connected by n - 1 undirected roads, and for any two cities there always exists a path between them.

Fox Ciel needs to assign an officer to each city. Each officer has a rank — a letter from 'A' to 'Z'. So there will be 26 different ranks, and 'A' is the topmost, so 'Z' is the bottommost.

There are enough officers of each rank. But there is a special rule must obey: if x and y are two distinct cities and their officers have the same rank, then on the simple path between x and y there must be a city z that has an officer with higher rank. The rule guarantee that a communications between same rank officers will be monitored by higher rank officer.

Help Ciel to make a valid plan, and if it's impossible, output "Impossible!".

## Input

The first line contains an integer n (2 ≤ n ≤ 10^5) — the number of cities in Tree Land.

Each of the following n - 1 lines contains two integers a and b (1 ≤ a, b ≤ n, a ≠ b) — they mean that there will be an undirected road between a and b. Consider all the cities are numbered from 1 to n.

It guaranteed that the given graph will be a tree.

## Output

If there is a valid plan, output n space-separated characters in a line — i-th character is the rank of officer in the city with number i.

Otherwise output "Impossible!".


## ideas
1. 如果 n <= 26, 每个只分配一个就可以了
2. 所有的叶子节点，可以分配'Z'，然后第二层分配'Y', 然后第三层分配'X', ...
3. 但是这样子有可能不够分配，比如超过26层的，就无法分配了
4. 考虑一条直线, 1,2 ..... n
5. ZYX...BAZYX....BZ....CZ....D...ZYZ 这个？不能是Y，否则不满足两个Y中间要有比Y大的rank
6. 貌似是这个pattern
7. 26 + 25 + ... +  1  = (26 + 1) * 26 / 2 (这个是高度)
8. 答案不对。漏掉了什么吗？ 
9. DCBA  DCB  DC  D
10. DC(D)BA  DC(D)B  DC D
11. DC(D)B(D)A 
12. DC(D)B(DC)A
13. F(X) 表示 A...X 能够使用的最多的长度
14. F(X + 1) = F(X) * 2 + 1
15. A => BAB => CBCACBC => F(26) 那就足够大了 
16. 一条线的问题解决了。
17. 但是树要怎么处理？
18. 假设在左子树中，有w个节点， 右子树中有u个节点
19. 理论上他们需要配色是根据f(x)给确定下来，然后给节点v设置一个大于这两个节点的数量rank？
20. 