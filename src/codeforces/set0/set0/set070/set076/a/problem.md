The kingdom of Olympia consists of N cities and M bidirectional roads. Each road connects exactly two cities and two cities can be connected with more than one road. Also it possible that some roads connect city with itself making a loop.

All roads are constantly plundered with bandits. After a while bandits became bored of wasting time in road robberies, so they suggested the king of Olympia to pay off. According to the offer, bandits want to get a gift consisted of gold and silver coins. Offer also contains a list of restrictions: for each road it is known gi — the smallest amount of gold and si — the smallest amount of silver coins that should be in the gift to stop robberies on the road. That is, if the gift contains a gold and b silver coins, then bandits will stop robberies on all the roads that gi ≤ a and si ≤ b.

Unfortunately kingdom treasury doesn't contain neither gold nor silver coins, but there are Olympian tugriks in it. The cost of one gold coin in tugriks is G, and the cost of one silver coin in tugriks is S. King really wants to send bandits such gift that for any two cities there will exist a safe path between them. Your task is to find the minimal cost in Olympian tugriks of the required gift.

Input
The first line of the input contains two integers N and M (2 ≤ N ≤ 200, 1 ≤ M ≤ 50 000) — the number of cities and the number of roads, respectively. The second line contains two integers G and S (1 ≤ G, S ≤ 109) — the prices of gold and silver coins in tugriks. The following M lines contain information about the offer. Each of the records in list is given as four integers xi, yi, gi, si, where xi and yi are the numbers of cities that the road connects and gi, si are minimal gold and silver coins requirements for the i-th road (1 ≤ xi, yi ≤ N, 1 ≤ gi, si ≤ 109). Cities are numbered from 1 to N. It is possible that there are more than one road between a pair of cities. It is possible that a road connects the city with itself.

Output
The output should contain the minimal cost of the gift in Olympian tugriks. If there is no gift that satisfies the given requirements output 

### ideas
1. 假设使用 x, y( 那么总共 x * G + y * S)
2. 那么所有 g[i] <= x and s[i] <= y 的边都可以被保留下来
3. 且保证剩余的这些边要联通整个图
4. 这里就是由 ax + by = c 组成的线，要在这个线上，找到一个点，在它的左下角，存在n个点的一个连通图
5. 假设2分c (可以log的吧)
6. 然后在这条线上找是否存在一个点，它的左下角的区域内，包括一个n的联通图
7. 从最左边向右边移动，在移动的过程中，添加新的节点（满足x条件，y条件）并移除过期的条件
8. 但移除怎么处理呢？
9. 自环可以直接移除掉。但似乎对答案没有帮助。
10. 可以用交集的概念吗？
11. 假设在线a*x + by = c 这样的一个线的下方
12. 只考虑x（不管y，但必须在c的下方）在每个x处记录它左边的dsu
13. 同样的对y，记录它下方的dsu
14. 那么在(x, y)处，是不是就是两个集合的交集？
15. 这里，可以对(u, v)进行判断，是否既在x的集合中，也在y的集合中
16. 好像是可以的
17.( 200 * 200 / 2 + 5w) * logc