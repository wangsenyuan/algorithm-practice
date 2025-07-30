# Problem D: Bankopolis

## Problem Description

Bankopolis is an incredible city in which all the n crossroads are located on a straight line and numbered from 1 to n along it. On each crossroad there is a bank office.

The crossroads are connected with m oriented bicycle lanes (the i-th lane goes from crossroad ui to crossroad vi), the difficulty of each of the lanes is known.

Oleg the bank client wants to gift happiness and joy to the bank employees. He wants to visit exactly k offices, in each of them he wants to gift presents to the employees.

The problem is that Oleg don't want to see the reaction on his gifts, so he can't use a bicycle lane which passes near the office in which he has already presented his gifts (formally, the i-th lane passes near the office on the x-th crossroad if and only if min(ui, vi) < x < max(ui, vi)). Of course, in each of the offices Oleg can present gifts exactly once. Oleg is going to use exactly k - 1 bicycle lane to move between offices. Oleg can start his path from any office and finish it in any office.

Oleg wants to choose such a path among possible ones that the total difficulty of the lanes he will use is minimum possible. Find this minimum possible total difficulty.

## Input

The first line contains two integers n and k (1 ≤ n, k ≤ 80) — the number of crossroads (and offices) and the number of offices Oleg wants to visit.

The second line contains single integer m (0 ≤ m ≤ 2000) — the number of bicycle lanes in Bankopolis.

The next m lines contain information about the lanes.

The i-th of these lines contains three integers ui, vi and ci (1 ≤ ui, vi ≤ n, 1 ≤ ci ≤ 1000), denoting the crossroads connected by the i-th road and its difficulty.

## Output

In the only line print the minimum possible total difficulty of the lanes in a valid path, or -1 if there are no valid paths.

## Examples

### Example 1

**Input:**
```
7 4
4
1 6 2
6 2 2
2 4 2
2 7 1
```

**Output:**
```
6
```

**Explanation:**
In the first example Oleg visiting banks by path 1 → 6 → 2 → 4.

Path 1 → 6 → 2 → 7 with smaller difficulty is incorrect because crossroad 2 → 7 passes near already visited office on the crossroad 6.

### Example 2

**Input:**
```
4 3
4
2 1 2
1 3 2
3 4 2
4 1 1
```

**Output:**
```
3
```

**Explanation:**
In the second example Oleg can visit banks by path 4 → 1 → 3.


## ideas
1. 先理解一下这个题目，有向图，从某个点s开始，使用k-1条有向边，访问k个office，
2.  min(ui, vi) < x < max(ui, vi)， 这里x是指当前oleg所在的office编号，
3.  所以，这里的意思应该是如果oleg选择访问了x, 那么那些所有cross x（跨过x的边）就不能被使用
4.  1 -> 8 -> 20 -> 19 -> 9
5.  dp[x][y][0/1][j]表示最后一条边是(0, x -> y， 1, y -> x) 时的最优解（使用了j条边）
6.  dp[u][y][1][j+1] = dp[x][y][0][j] + cost(y -> u) x < u < y
7.  dp[y][u][0][j+1] = dp[x][y][0][j] + cost(y -> u), y < u
8.  dp[u][x][1][j+1] = dp[x][y][1][j] + cost(x -> u), u < x
9.  dp[u][y][0][j+1] = dp[x][y][1][j] + cost(x -> u), x < u < y
10.  这里不大对，假设到达节点x后，它其实有两个选择，往前还是往后，如果往后，它不能超过它的范围y，往后，也不能超过x
11.  dp[x][y][u][j]表示，在范围[x...y]且落在节点u上，且使用了j条边时的最优解
12.  dp[x][y][u][j] => dp[x][u][v][j+1] (v 在 x, u中间), dp[u][y][v][j+1], v在(u, y中间)
13.  如果 x = u, 那么它还可以往前
14.  这样是ok的 
15.  如何在[0][2][2]的时候，知道不能继续往右边去（而只能去左边呢？）
16.  