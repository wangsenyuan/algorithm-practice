# Problem E - Game Levels

## Problem Description

In one secret game developed deep in the ZeptoLab company, the game universe consists of n levels, located in a circle. You can get from level i to levels i - 1 and i + 1, also you can get from level 1 to level n and vice versa. The map of the i-th level description size is ai bytes.

In order to reduce the transmitted traffic, the game gets levels as follows. All the levels on the server are divided into m groups and each time a player finds himself on one of the levels of a certain group for the first time, the server sends all levels of the group to the game client as a single packet. Thus, when a player travels inside the levels of a single group, the application doesn't need any new information. Due to the technical limitations the packet can contain an arbitrary number of levels but their total size mustn't exceed b bytes, where b is some positive integer constant.

Usual situation is that players finish levels one by one, that's why a decision was made to split n levels into m groups so that each group was a continuous segment containing multiple neighboring levels (also, the group can have two adjacent levels, n and 1). Specifically, if the descriptions of all levels have the total weight of at most b bytes, then they can all be united into one group to be sent in a single packet.

Determine, what minimum number of groups do you need to make in order to organize the levels of the game observing the conditions above?

As developing a game is a long process and technology never stagnates, it is yet impossible to predict exactly what value will take constant value b limiting the packet size when the game is out. That's why the developers ask you to find the answer for multiple values of b.

## Input

The first line contains two integers n, q (2 ≤ n ≤ 10^6, 1 ≤ q ≤ 50) — the number of levels in the game universe and the number of distinct values of b that you need to process.

The second line contains n integers ai (1 ≤ ai ≤ 10^9) — the sizes of the levels in bytes.

The next q lines contain integers bj (1 ≤ bj ≤ 10^15), determining the values of constant b, for which you need to determine the answer.

## Output

For each value of bj from the input print on a single line integer mj (1 ≤ mj ≤ n), determining the minimum number of groups to divide game levels into for transmission via network observing the given conditions.

## Examples

### Input
```
6 3
2 4 2 1 3 2
7
4
6
```

### Output
```
2
4
3
```

## Note

In the test from the statement you can do in the following manner:

- at b = 7 you can divide into two segments: 2|421|32 (note that one of the segments contains the fifth, sixth and first levels);
- at b = 4 you can divide into four segments: 2|4|21|3|2;
- at b = 6 you can divide into three segments: 24|21|32|.

## ideas
1. 考虑单个b如何快速的计算。
2. 转换成前缀和后, s1, s2, s3, .... sn
3. 那么就是 si - s0 <= b, sj - s(i) <= b, ... sn - sk <= b
4. dp[i] = min(dp[j] + 1, s[i] - s[j] <= b)
5. 如果没有环型，那么通过使用双端队列，在O(n)时间内，就可以处理
6. 但是环型要怎么处理呢？
7. 怎么感觉是(s[n] - 0 + b - 1) / b ?
8. 4 2 1 3 2 2 and b = 7
9. 0,4,6,7,10,12,14 = 2
10. (sum + b - 1) / b ? 
11. 比如 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2  b = 7
12. 这个答案 = 6
13. sum = (5 * 6 + 6) / 7 = 6 
14. [sum / b]这个是最好的结果，不可能比这个更小（因为每一个数字都 <= b, 所以不会出现和超出的段）
15. 但是一定能得到这个结果吗？
16. 如果可以，似乎就太简单了。
17. 2 4 2 1 3 2。 应该是不成立的，如果成立，那么这个顺序的串(在b = 7)的时候，答案也应该是2
18. 2，2，2，2，2，2，2, 这个答案 = 3
19. 而且这个队列越长，答案越偏离sum/b
20. 在确定起点的时候，肯定是贪心的进行分割是最优的
21. 
