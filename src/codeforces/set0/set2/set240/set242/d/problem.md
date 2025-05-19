Valera has n counters numbered from 1 to n. Some of them are connected by wires, and each of the counters has a special button.

Initially, all the counters contain number 0. When you press a button on a certain counter, the value it has increases by one. Also, the values recorded in all the counters, directly connected to it by a wire, increase by one.

Valera and Ignat started having a dispute, the dispute is as follows. Ignat thought of a sequence of n integers a1, a2, ..., an. Valera should choose some set of distinct counters and press buttons on each of them exactly once (on other counters the buttons won't be pressed). If after that there is a counter with the number i, which has value ai, then Valera loses the dispute, otherwise he wins the dispute.

Help Valera to determine on which counters he needs to press a button to win the dispute.

Input
The first line contains two space-separated integers n and m (1 ≤ n, m ≤ 105), that denote the number of counters Valera has and the number of pairs of counters connected by wires.

Each of the following m lines contains two space-separated integers ui and vi (1 ≤ ui, vi ≤ n, ui ≠ vi), that mean that counters with numbers ui and vi are connected by a wire. It is guaranteed that each pair of connected counters occurs exactly once in the input.

The last line contains n space-separated integers a1, a2, ..., an (0 ≤ ai ≤ 105), where ai is the value that Ignat choose for the i-th counter.

Output
If Valera can't win the dispute print in the first line -1.

Otherwise, print in the first line integer k (0 ≤ k ≤ n). In the second line print k distinct space-separated integers — the numbers of the counters, where Valera should push buttons to win the dispute, in arbitrary order.

If there exists multiple answers, you are allowed to print any of them.

### ideas
1. 先理解题目A给所有的counter设置一个值
2. B选中其中的部分counter，然后每个选中的press一下，每个节点按照规则，产生一个值b[i]
3. 如果存在一个i, a[i] = b[i], 那么A获胜，否则B获胜
4. 假设所有的点都选了，这时候b[i] = deg[i] + 1
5. 假设这时候存在某个u, b[u] = deg[u], 那么必须选择它的一个邻居v, 不选中
6. 但是如果关闭v（不选中),假设v除了u还连到了其他的节点w, 且deg[w] = a[w] + 1
7. 那么关闭v的时候，就会使的w出现了a[w] = b[w]
8. 假设只有两个节点，(1, 2, 3) 两两相连; a[1] = 0, a[2] = 1, a[3] = 2
9. 显然没法都不选，如果只选a[1], 那么不能满足a[2], 如果选择(1, 2) 似乎是满足条件的（全部选择也是可以的）
10. 所以，感觉始终是有答案的？
11. 感觉要从目前a[i] = b[i], 且a[i]最大的值开始
12. 这时候，不去关闭v，而是关闭i。这样子i在后续的操作中，肯定是满足条件的（因为后续指挥减少）
13. 同时如果i的邻居中，出现了某个v，a[v] = b[v], 那么再把它加入进去
14. 这样子每个点只会被加入一次。所以时间性能是ok的。
15. 但是是否会出现无解的情况？
16. 不可能出现。因为每次关闭，影响的始终是邻居中，开启的点。如果影响到了，顶多是把它关闭掉。所以始终是可以操作的