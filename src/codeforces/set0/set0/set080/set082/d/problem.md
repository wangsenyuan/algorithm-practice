Vasya has recently developed a new algorithm to optimize the reception of customer flow and he considered the following problem.

Let the queue to the cashier contain n people, at that each of them is characterized by a positive integer ai — that is the time needed to work with this customer. What is special about this very cashier is that it can serve two customers simultaneously. However, if two customers need ai and aj of time to be served, the time needed to work with both of them customers is equal to max(ai, aj). Please note that working with customers is an uninterruptable process, and therefore, if two people simultaneously come to the cashier, it means that they begin to be served simultaneously, and will both finish simultaneously (it is possible that one of them will have to wait).

Vasya used in his algorithm an ingenious heuristic — as long as the queue has more than one person waiting, then some two people of the first three standing in front of the queue are sent simultaneously. If the queue has only one customer number i, then he goes to the cashier, and is served within ai of time. Note that the total number of phases of serving a customer will always be equal to ⌈n / 2⌉.

Vasya thinks that this method will help to cope with the queues we all hate. That's why he asked you to work out a program that will determine the minimum time during which the whole queue will be served using this algorithm.


### ideas
1. n <= 1000, n * n 的算法will work
2. dp[i][j] 表示前i个人，但是留下的是j (j <= i)
3. 那么j肯定在目前的头部, 再增加两个人
4. dp[i+2][j] = dp[i][j] + max(a[i+1], a[i+2]) (j又被留下来了)
5. dp[i+2][i+1] = dp[i][j] + max(a[j], a[i+2]) (i+1被留下来了)
6. dp[i+2][i+2] = dp[i][j] + max(a[j], a[i+1]) (i+2被留下来了)
7. i可以变成i/2