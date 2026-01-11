# Problem C

Kostya is a progamer specializing in the discipline of Dota 2. Valve Corporation, the developer of this game, has recently released a new patch which turned the balance of the game upside down. Kostya, as the captain of the team, realizes that the greatest responsibility lies on him, so he wants to resort to the analysis of innovations patch from the mathematical point of view to choose the best heroes for his team in every game.

A Dota 2 match involves two teams, each of them must choose some heroes that the players of the team are going to play for, and it is forbidden to choose the same hero several times, even in different teams. In large electronic sports competitions where Kostya's team is going to participate, the matches are held in the Captains Mode. In this mode the captains select the heroes by making one of two possible actions in a certain, predetermined order: pick or ban.

To pick a hero for the team. After the captain picks, the picked hero goes to his team (later one of a team members will play it) and can no longer be selected by any of the teams.

To ban a hero. After the ban the hero is not sent to any of the teams, but it still can no longer be selected by any of the teams.

The team captain may miss a pick or a ban. If he misses a pick, a random hero is added to his team from those that were available at that moment, and if he misses a ban, no hero is banned, as if there was no ban.

Kostya has already identified the strength of all the heroes based on the new patch fixes. Of course, Kostya knows the order of picks and bans. The strength of a team is the sum of the strengths of the team's heroes and both teams that participate in the match seek to maximize the difference in strengths in their favor. Help Kostya determine what team, the first one or the second one, has advantage in the match, and how large the advantage is.

## Input

The first line contains a single integer $n$ ($2 \leq n \leq 100$) — the number of heroes in Dota 2.

The second line contains $n$ integers $s_1, s_2, \ldots, s_n$ ($1 \leq s_i \leq 10^6$) — the strengths of all the heroes.

The third line contains a single integer $m$ ($2 \leq m \leq \min(n, 20)$) — the number of actions the captains of the team must perform.

Next $m$ lines look like "action team", where action is the needed action: a pick (represented as a "p") or a ban (represented as a "b"), and team is the number of the team that needs to perform the action (number 1 or 2).

It is guaranteed that each team makes at least one pick. Besides, each team has the same number of picks and the same number of bans.

## Output

Print a single integer — the difference between the strength of the first team and the strength of the second team if the captains of both teams will act optimally well.

## Examples

**Input:**

```text
2
2 1
2
p 1
p 2
```

**Output:**

```text
1
```

**Input:**

```text
6
6 4 5 4 5 5
4
b 2
p 1
b 1
p 2
```

**Output:**

```text
0
```

**Input:**

```text
4
1 2 3 4
4
p 2
b 2
p 1
b 1
```

**Output:**

```text
-2
```

### ideas
1. 这个题目有点奇怪
2. pick是不是就应该拿当时能拿到的最高值？
3. ban呢？好像不一定，因为有可能后续就轮到自己pick了。即使轮不到，也有可能让自己pick到更好的结果
4. 比如对于 7, 6, 4
5. 如果把7ban掉了，对方获得了6，自己只能获得4，差值是2
6. 如果不ban7，那么对方获得7，自己获得6；差值就是1
7. 如果n比较小的话，还可以用mask表示，哪些被ban掉了
8. 但是，似乎这个mask也不用特别大
9. 如果在位置i处，前面有w次ban，那么只可能在i后面（包括i）有w位被ban掉
10. 更远的ban没有意义
11. 反过来，当前是ban，如果后面有w次pick，那么超过w的ban也没有意义
12. dp[i][mask] mask[d] = 1, 表示 i+d是否被ban掉了