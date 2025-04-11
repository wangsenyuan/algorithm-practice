In this task Anna and Maria play the following game. Initially they have a checkered piece of paper with a painted n × m rectangle (only the border, no filling). Anna and Maria move in turns and Anna starts. During each move one should paint inside the last-painted rectangle a new lesser rectangle (along the grid lines). The new rectangle should have no common points with the previous one. Note that when we paint a rectangle, we always paint only the border, the rectangles aren't filled.

Nobody wins the game — Anna and Maria simply play until they have done k moves in total. Count the number of different ways to play this game.

Input
The first and only line contains three integers: n, m, k (1 ≤ n, m, k ≤ 1000).

Output
Print the single number — the number of the ways to play the game. As this number can be very big, print the value modulo 1000000007 (109 + 7).

### ideas
1. dp[h][w][k] 表示在h * w 的一个grid中，完成k步时的数量
2. dp[h][w][0] = 1
3. dp[h][w][k] = 就是找 dp[i][j][k-1] * (i, j大小的矩形出现的次数)
4. 后面一部分似乎是这样的， (n  - 2 - i + 1) * (m - 2 - j + 1)
5. 但是(i, j, k)的迭代还是太多了
6. 但是，有个点是，因为不能有公共点， 所以 n * m, i * j 之间至少差 2 * (n + m) - 2
7. 所以，这样子的话，衰减的似乎还是很快的 
8. dp[h][w][k] = 迭代i ( from 1...h - 2) (j from 1 to w - 2) and i * j <= n * m - (2 * (n + m) - 2 )
9. k也是有限制的 k <= min(n / 2, m / 2)