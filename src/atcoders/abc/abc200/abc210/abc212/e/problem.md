# Problem E - Safety Journey

The Republic of AtCoder has N cities, called City 1, City 2, ..., City N. Initially, there was a bidirectional road between every pair of different cities, but M of these roads have become unusable due to deterioration over time. More specifically, for each 1≤i≤M, the road connecting City U_i and City V_i has become unusable.

Takahashi will go for a K-day trip that starts and ends in City 1. Formally speaking, a K-day trip that starts and ends in City 1 is a sequence of K+1 cities (A_0, A_1, ..., A_K) such that A_0 = A_K = 1 holds and for each 0≤i≤K−1, A_i and A_i+1 are different and there is still a usable road connecting City A_i and City A_i+1.

Print the number of different K-day trips that start and end in City 1, modulo 998244353. Here, two K-day trips (A_0, A_1, ..., A_K) and (B_0, B_1, ..., B_K) are said to be different when there exists an i such that A_i ≠ B_i.

 ### ideas
 1. 感觉是矩阵乘法
 2. 好像也不对， n * n * n * log k， 还是太慢了
 3. 这个还不如， dp[i][j] 到达节点i时，且进行了j天后的计数
 4. 这个是 n * n * k 的算法