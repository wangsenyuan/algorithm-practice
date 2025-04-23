There is a game where you need to move through a labyrinth. The labyrinth consists of 𝑛
 platforms, connected by 𝑚
 passages.

Each platform is at some level 𝑙𝑖
, an integer number from 0
 to 𝐻−1
. In a single step, if you are currently on platform 𝑖
, you can stay on it, or move to another platform 𝑗
. To move to platform 𝑗
 they have to be connected by the passage, and their levels have to be the same, namely 𝑙𝑖=𝑙𝑗
.

After each step, the levels of all platforms change. The new level of platform 𝑖
 is calculated as 𝑙′𝑖=(𝑙𝑖+𝑠𝑖)mod𝐻
, for all 𝑖
.

You start on platform 1
. Find the minimum number of steps you need to get to platform 𝑛
.

Input
The first line of input contains a single integer 𝑡
 (1≤𝑡≤104
) — the number of test cases. Then the descriptions of the test cases follow.

The first line of each test case contains three integers 𝑛
, 𝑚
, and 𝐻
 (2≤𝑛≤105
, 1≤𝑚≤105
, 1≤𝐻≤109
).

The second line contains 𝑛
 integers 𝑙𝑖
, the initial level of each platform (0≤𝑙𝑖≤𝐻−1
).

The third line contains 𝑛
 integers 𝑠𝑖
, the change of level for each platform (0≤𝑠𝑖≤𝐻−1
).

Next 𝑚
 lines contain a description of the passages. Each passage is described as a pair of integers — the platforms, connected by the passage. There is at most one passage connecting each pair of platforms, and there is no passage connecting a platform to itself.

The sum of 𝑛
 for all tests does not exceed 105
, the sum of 𝑚
 for all tests does not exceed 105
.

### ideas
1. 假设在时刻x，那么可以计算此时，所有处在同一高度的平台
2. 且如果知道用户所在的平台（集合）那么就可以找到新的集合（通过通道连接的集合）
3. 对于一个通道它连接的平台为u，v
4. (l[u] + s[u] * t) % H = (l[v] + s[v] % t) % H,可以移动
5. 可以计算出最早的t（基于当前时刻）
6. 这个应该可以算出来，但我不会算～～ (如果s[u] == s[v]), 那么永远都相遇不了
7. 把这个看作直线，且l[u] < l[v], 如果s[u] > s[v]， 
8. 假设经过时间t后， l[u] + s[u] * t = l[v] + s[v] * t
9. 如果s[v] > s[u]， 咋算呢？
10. (l[u] + s[u] * t) % H = (l[v] + s[v] * t) % H
11. (s[v] - s[u]) * t == (l[v] - l[u]) % H
12. 