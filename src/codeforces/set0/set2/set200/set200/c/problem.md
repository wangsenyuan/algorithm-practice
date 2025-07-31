# Football Championship Problem

## Problem Description

The Berland National team takes part in the local Football championship which now has a group stage. Let's describe the formal rules of the local championship:

### Game Rules
- The team that kicked most balls in the enemy's goal area wins the game
- The victory gives 3 points to the team, the draw gives 1 point and the defeat gives 0 points

### Group Stage Rules
- A group consists of four teams
- The teams are ranked by the results of six games: each team plays exactly once with each other team
- The teams that get places 1 and 2 in the group stage results, go to the next stage of the championship

### Ranking Criteria
In the group stage the team's place is defined by the total number of scored points: the more points, the higher the place is. If two or more teams have the same number of points, then the following criteria are used (the criteria are listed in the order of falling priority, starting from the most important one):

1. **Goal Difference**: The difference between the total number of scored goals and the total number of missed goals in the championship: the team with a higher value gets a higher place
2. **Goals Scored**: The total number of scored goals in the championship: the team with a higher value gets a higher place
3. **Lexicographical Order**: The lexicographical order of the name of the teams' countries: the country with the lexicographically smaller name gets a higher place

## Task

The Berland team plays in the group where the results of 5 out of 6 games are already known. To be exact, there is the last game left. There the Berland national team plays with some other team. 

The coach asks you to find such score X:Y (where X is the number of goals Berland scored and Y is the number of goals the opponent scored in the game), that fulfills the following conditions:

1. **X > Y**: Berland is going to win this game
2. **Top 2 Placement**: After the game Berland gets the 1st or the 2nd place in the group
3. **Minimum Goal Difference**: If there are multiple variants, you should choose such score X:Y, where value X - Y is minimum
4. **Minimum Goals Conceded**: If it is still impossible to come up with one score, you should choose the score where value Y (the number of goals Berland misses) is minimum

## Input

The input has five lines.

Each line describes a game as `"team1 team2 goals1:goals2"` (without the quotes), what means that team `team1` played a game with team `team2`, besides, `team1` scored `goals1` goals and `team2` scored `goals2` goals. 

- The names of teams `team1` and `team2` are non-empty strings, consisting of uppercase English letters, with length of no more than 20 characters
- `goals1`, `goals2` are integers from 0 to 9
- The Berland team is called "BERLAND"
- It is guaranteed that the Berland team and one more team played exactly 2 games and the other teams played exactly 3 games

## Output

Print the required score in the last game as `X:Y`, where X is the number of goals Berland scored and Y is the number of goals the opponent scored. 

If the Berland team does not get the first or the second place in the group, whatever this game's score is, then print on a single line `"IMPOSSIBLE"` (without the quotes).

**Note**: The result score can be very huge, 10:0 for example.

## Examples

### Example 1

**Input:**
```
AERLAND DERLAND 2:1
DERLAND CERLAND 0:3
CERLAND AERLAND 0:1
AERLAND BERLAND 2:0
DERLAND BERLAND 4:0
```

**Output:**
```
6:0
```

**Explanation:**
In the first sample "BERLAND" plays the last game with team "CERLAND". If Berland wins with score 6:0, the results' table looks like that in the end:

- **AERLAND** (points: 9, goal difference: 4, goals scored: 5)
- **BERLAND** (points: 3, goal difference: 0, goals scored: 6)
- **DERLAND** (points: 3, goal difference: 0, goals scored: 5)
- **CERLAND** (points: 3, goal difference: -4, goals scored: 3)

### Example 2

**Input:**
```
AERLAND DERLAND 2:2
DERLAND CERLAND 2:3
CERLAND AERLAND 1:3
AERLAND BERLAND 2:1
DERLAND BERLAND 4:1
```

**Output:**
```
IMPOSSIBLE
```

**Explanation:**
In the second sample teams "AERLAND" and "DERLAND" have already won 7 and 4 points, respectively. The Berland team wins only 3 points, which is not enough to advance to the next championship stage. 

### ideas
1. 假设有A, B, C, D 4只队伍
2. 目前剩余A，B还没有比赛，B是我们的目标
3. 最后的结果 X > Y, 那么就要求B获胜，并得到分数3, 
4. score[0], score[1] + 3, score[2], score[3]
5. 如果发现, score[1] + 3 还是不够前两名, 那么 => IMPOSSIBLE
6. 如果发现, score[1] + 3 可以得到前两名，