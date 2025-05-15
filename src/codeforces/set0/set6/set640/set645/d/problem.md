While Farmer John rebuilds his farm in an unfamiliar portion of Bovinia, Bessie is out trying some alternative jobs. In her new gig as a reporter, Bessie needs to know about programming competition results as quickly as possible. When she covers the 2016 Robot Rap Battle Tournament, she notices that all of the robots operate under deterministic algorithms. In particular, robot i will beat robot j if and only if robot i has a higher skill level than robot j. And if robot i beats robot j and robot j beats robot k, then robot i will beat robot k. Since rapping is such a subtle art, two robots can never have the same skill level.

Given the results of the rap battles in the order in which they were played, determine the minimum number of first rap battles that needed to take place before Bessie could order all of the robots by skill level.


### ideas
1. 好像理解错了
2. 如果整个图，不满足条件，是否意味着没有答案？
3. 考虑一棵树（不是一条线的情况），全部的情况都不行
4. 那么省掉任何一部分，肯定也不能得到一个唯一的rag
5. 确实是可以二分的
6. 然后就是判断如何得到一个唯一的dag
7. 先找到trunk（如果有多个头、或者多个尾， false）
8. 然后把这个trunk移除掉，剩余的是否能找到一个trunk（多个trunk）
9. 