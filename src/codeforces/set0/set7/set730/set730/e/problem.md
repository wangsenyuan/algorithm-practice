All-Berland programming contest comes to an end. In total, n teams participated in it. Like in ACM-ICPC, current results stopped refreshing one hour before the contest ends. So at the Award Ceremony, results are partially known. For each team the value ai is given — the number of points the i-th team has earned before the last hour of the contest. Besides that, the Jury has evaluated all submissions sent during the last hour and knows values di — the number of points earned by the i-th team during the last hour (these values can be negative, which means that a team can lose points).

Before the contest, each team got unique id from 1 to n. According to the contest rules, a team with more points takes a higher place. If two or more teams have equal number of points, the team with lower id will take the higher place. So no two teams can share the same place.

The Award Ceremony proceeds in the following way. At the beginning of the ceremony, a large screen shows the results for the time moment "one hour before the end", which means that the i-th team has ai points. Then the Jury unfreezes results of the teams one by one in some order. When result of the j-th team is unfrozen, its score changes from aj to aj + dj. At this time the table of results is modified and the place of the team can change. The unfreezing of the j-th team is followed by the applause from the audience with duration of |xj - yj| seconds, where xj is the place of the j-th team before unfreezing and yj is the place right after the unfreezing. For example, if the team does not change the place, there is no applause from the audience. As you can see, during the Award Ceremony, each team will be unfrozen exactly once.

Your task is to find such an order to unfreeze all the teams that the total duration of applause is maximum possible.

## Input

The first line of the input file contains a single integer n (1 ≤ n ≤ 100) — the number of teams.

Each of the next n lines contains two integers ai and di (1 ≤ ai ≤ 100, -100 ≤ di ≤ 100) — the number of points the i-th team has earned before the last hour of the contest and the number of points earned by this team during the last hour. It is possible that after unfreezing a team will have a negative score.

## Output

Print the only integer — maximal total applause duration in seconds if the Jury can choose any order of the teams to unfreeze.

## Examples

**Input:**
```
4
17 -14
52 -5
1 52
6 0
```

**Output:**
```
4
```

**Input:**
```
5
4 5
3 2
5 -3
6 -2
4 3
```

**Output:**
```
14
```

## Note

In the first example the initial standings are:

- Team 2, 52 points
- Team 1, 17 points
- Team 4, 6 points
- Team 3, 1 point

Here any order of unfreezing the teams leads to 4 seconds of applause in total. For example, let's unfreeze teams in their order from the Team 1 to the Team 4.

After the Team 1 became unfrozen the standings are:

- Team 2, 52 points
- Team 4, 6 points
- Team 1, 3 points
- Team 3, 1 point

So there is 1 second of applause, because the difference between old and new places |2 - 3| = 1.

After the Team 2 became unfrozen the standings are:

- Team 2, 47 points
- Team 4, 6 points
- Team 1, 3 points
- Team 3, 1 point

The place of the Team 2 has not changed, so no applause during unfreezing.

After the Team 3 became unfrozen the standings are:

- Team 3, 53 points
- Team 2, 47 points
- Team 4, 6 points
- Team 1, 3 points

The place of the Team 3 has changed from 4 to 1, so the duration of applause is |4 - 1| = 3.

The unfreezing of the Team 4 has not changed any place because d4 = 0.

Therefore, the total duration of applause is 1 + 0 + 3 + 0 = 4 seconds.


### ideas
1. 没有遇到过这类问题～
2. 那是不是每次都选择（贪心）那个改变名次最大的人解冻？
3. 假设现在选择了i，并且它的名次是往下掉了（比如从第一名到了第x位）收获是x-1
4. 但是后面有某个j，会让它的名次继续下掉一位？
5. 那么这个j，必须满足的时候，它的现有名次，在x后面，新的名次在x前面（且在1的后面）
6. 如果有很多个这样的j，似乎让i在后面解冻是更好的选择
7. 假设x[i], y[i]， 和 x[j], y[j]有重叠的区域
8. 是不是什么样的顺序都可以（无所谓的）