A couple of friends, Axel and Marston are travelling across the country of Bitland. There are n towns in Bitland, with some pairs of towns connected by one-directional roads. Each road in Bitland is either a pedestrian road or a bike road. There can be multiple roads between any pair of towns, and may even be a road from a town to itself. However, no pair of roads shares the starting and the destination towns along with their types simultaneously.

The friends are now located in the town 1 and are planning the travel route. Axel enjoys walking, while Marston prefers biking. In order to choose a route diverse and equally interesting for both friends, they have agreed upon the following procedure for choosing the road types during the travel:

The route starts with a pedestrian route.
Suppose that a beginning of the route is written in a string s of letters P (pedestrain road) and B (biking road). Then, the string  is appended to s, where  stands for the string s with each character changed to opposite (that is, all pedestrian roads changed to bike roads, and vice versa).
In the first few steps the route will look as follows: P, PB, PBBP, PBBPBPPB, PBBPBPPBBPPBPBBP, and so on.

After that the friends start travelling from the town 1 via Bitlandian roads, choosing the next road according to the next character of their route type each time. If it is impossible to choose the next road, the friends terminate their travel and fly home instead.

Help the friends to find the longest possible route that can be travelled along roads of Bitland according to the road types choosing procedure described above. If there is such a route with more than 1018 roads in it, print -1 instead.

Input
The first line contains two integers n and m (1 ≤ n ≤ 500, 0 ≤ m ≤ 2n2) — the number of towns and roads in Bitland respectively.

Next m lines describe the roads. i-th of these lines contains three integers vi, ui and ti (1 ≤ vi, ui ≤ n, 0 ≤ ti ≤ 1), where vi and ui denote start and destination towns indices of the i-th road, and ti decribes the type of i-th road (0 for a pedestrian road, 1 for a bike road). It is guaranteed that for each pair of distinct indices i, j such that 1 ≤ i, j ≤ m, either vi ≠ vj, or ui ≠ uj, or ti ≠ tj holds.

Output
If it is possible to find a route with length strictly greater than 1018, print -1. Otherwise, print the maximum length of a suitable path.

Examples
InputCopy
2 2
1 2 0
2 2 1
OutputCopy
3
InputCopy
2 3
1 2 0
2 2 1
2 2 0
OutputCopy
-1
Note
In the first sample we can obtain a route of length 3 by travelling along the road 1 from town 1 to town 2, and then following the road 2 twice from town 2 to itself.


### ideas
1. 考虑一个路径长度为w(它肯定是某个2倍数, 0, 1, 2, 4, 8, 16, ...)
2. 对于固定的长度，路径上的类型是确定的 0110100110010110, ... 
3. 不知道怎么考虑这个问题。
4. 一个点会被访问多次，所以需要确定，什么情况下，能够判断出这是一次新的路径，什么时候会导致循环
5. 假设长度为d的处理掉了，知道了某些信息,
6. 目前要处理长度为2*d的路径
7. 那么首先需要知道，有哪些点是d长度的终点，假设其中的某个点为u
8. 然后从u出发，如果能找到一条翻转的路径，经过了 d / 2 长度后，到达了位置v
9. 然后从v经过d/2的长度且翻转的路径，能够到达位置w
10. 比如长度为 d = 4的路径处理过了，知道能够到达u，现在要处理d = 8的路径
11. 已知到达u访问了 0110, 从u出发访问长度为2的翻转路径 10, 到达位置v, 然后从v出发访问长度为2的不翻转路径01到达w
12. 那么拼接起来就是 0110 1001 从而找到了一条长度为8的路径，且到达了位置w
13. dp[x][d][0/1] 表示从u出发，经过长度为 1 << d长度的路径，且是否翻转，能够到达的位置集合？
14. dp[x][d + 1][0] = dp[x][d][0] && dp[u][d-1][1] && dp[v][d-1][0]
15. dp[x][d+1][1] = dp[x][d][1] && dp[u][d-1][0] && dp[v][d-1][1]
16. 至少表示了一些信息出来; u, v要怎么确定？
17. dp[x][d+1][0] = dp[x][d][0] && dp[u][d][1] ?
18. dp[x][d+1][1] = dp[x][d][1] && dp[u][d][0] 
19. 似乎也是成立的，这样子，就只需要确定u了
20. fp[x][d][0] 表示通过长度为 1 << d, 且不翻转的路径能够到达的节点集合
21. fp[x][d][1] 表示通过长度为 1 << d, 且不翻转的路径能够到达的节点集合
22. fp[x][d+1][0] = merge fp[u][d][1] for each u from fp[x][d][0]
23. 这样子，甚至可以不用处理dp
24. 但是这样子，潜在的似乎是 O(n  * n)的，不考虑d的影响；
25. 因为 fp[x][d][0] 可能包含所有的节点， fp[u][d][1] 也包含所有的节点
26. 但是有一个猜想是，如果能在出现循环路径的情况是，提前返回，这个复杂性不会到n * n
27. 什么情况下，会出现循环呢？fp[x][d+1][0], x属于fp[x][d][0] && fp[x][d][1] => -1
28. 这个似乎是成立的，但是只是一个猜想（但似乎是显然的）


### editorial

Let us write Ai for the binary string obtained after i inverse-append steps, for example, A0 = 0, A1 = 01, and so on. Let us also write . By definition we must have , and .

Let us store matrices Pk and Qk, with entries Pk / Qk(v, u) equal to 1 for pairs of vertices v, u such that there is a Ak/Bk-path from v to u. Note that P0 and Q0 are exactly the adjacency matrices with 0- and 1-arcs respectively.

Next, note that Pk + 1(v, u) = 1 if and only if there is a vertex w such that Pk(v, w) = Qk(w, u) = 1, and a similar condition can be written for Qk + 1(v, u). It follows that Pk + 1 and Qk + 1 can be computed using Pk and Qk in O(n3) time (the method is basically boolean matrix multiplication: , ).

To use the matrices Pk and Qk to find the answer, let us store L — the largest answer found, and S — the set of vertices reachable from the vertex 1 in exactly L steps. Let's process k by decreasing from a certain value k0, and see if L can be increased by 2k. The next 2k characters after L-th position will form the string Ak or Bk depending on the popcount parity of L. Let's denote S' the set of vertices reachable from S following Ak / Bk. If S' is non-empty, we can increase L by 2k, and assign S = S', otherwise, we don't change anything. In the end, L will be the maximal path length as long as it at less than 2k0.

Note that we can take k0 = 60 since we don't care about exact value of answer if it is greater than 260. This results in an O(k0n3) solution, which is too slow. However, optimizing boolean multiplication with bitsets cuts the working time  times, and the solution is now fast enough.

Complexity:  time, and  memory. Here , and w = 64 is the word length in bits.In the second sample we can obtain an arbitrarily long route by travelling the road 1 first, and then choosing road 2 or 3 depending on the necessary type.
