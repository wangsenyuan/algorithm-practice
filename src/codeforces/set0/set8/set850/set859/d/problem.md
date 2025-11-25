# Problem Description

The annual college sports-ball tournament is approaching, which for trademark reasons we'll refer to as Third Month Insanity. There are a total of $2^N$ teams participating in the tournament, numbered from 1 to $2^N$. The tournament lasts $N$ rounds, with each round eliminating half the teams. The first round consists of $2^{N-1}$ games, numbered starting from 1. In game $i$, team $2 \cdot i - 1$ will play against team $2 \cdot i$. The loser is eliminated and the winner advances to the next round (there are no ties). Each subsequent round has half as many games as the previous round, and in game $i$ the winner of the previous round's game $2 \cdot i - 1$ will play against the winner of the previous round's game $2 \cdot i$.

Every year the office has a pool to see who can create the best bracket. A bracket is a set of winner predictions for every game. For games in the first round you may predict either team to win, but for games in later rounds the winner you predict must also be predicted as a winner in the previous round. Note that the bracket is fully constructed before any games are actually played. Correct predictions in the first round are worth 1 point, and correct predictions in each subsequent round are worth twice as many points as the previous, so correct predictions in the final game are worth $2^{N-1}$ points.

For every pair of teams in the league, you have estimated the probability of each team winning if they play against each other. Now you want to construct a bracket with the maximum possible expected score.

## Input

Input will begin with a line containing $N$ ($2 \leq N \leq 6$).

$2^N$ lines follow, each with $2^N$ integers. The $j$-th column of the $i$-th row indicates the percentage chance that team $i$ will defeat team $j$, unless $i = j$, in which case the value will be 0. It is guaranteed that the $i$-th column of the $j$-th row plus the $j$-th column of the $i$-th row will add to exactly 100.

## Output

Print the maximum possible expected score over all possible brackets. Your answer must be correct to within an absolute or relative error of $10^{-9}$.

Formally, let your answer be $a$, and the jury's answer be $b$. Your answer will be considered correct if $\frac{|a - b|}{\max(1, |b|)} \leq 10^{-9}$.

## Examples

**Example 1:**

**Input:**

```text
2
0 40 100 100
60 0 40 40
0 60 0 45
0 60 55 0
```

**Output:**

```text
1.75
```

**Example 2:**

**Input:**

```text
3
0 0 100 0 100 0 0 0
100 0 100 0 0 0 100 100
0 0 0 100 100 0 0 0
100 100 0 0 0 0 100 100
0 100 0 100 0 0 100 0
100 100 100 100 100 0 0 0
100 0 100 0 0 100 0 0
100 0 100 0 100 100 100 0
```

**Output:**

```text
12
```

**Example 3:**

**Input:**

```text
2
0 21 41 26
79 0 97 33
59 3 0 91
74 67 9 0
```

**Output:**

```text
3.141592
```

## Note

In the first example, you should predict teams 1 and 4 to win in round 1, and team 1 to win in round 2. Recall that the winner you predict in round 2 must also be predicted as a winner in round 1.


### ideas
1. 对于第一轮的比赛，对于每场比赛，预测其中某个team是winner；这样子，可以有， $2^^{n-2}$ 个指标
2. dp[state][h] 表示在第h层（及以上），如果由state表示的team获胜的期望分
3. 但是这个state有点大， 1 << 32, 搞不定～
4. 可不可以单独考虑每个人的贡献？ 似乎不大行，因为不知道对手的情况下，无法计算胜率
5. 考虑树，dp[l....r][s] = 子树u，最后胜出的是s时的分数
6. 那么在dp[1....n][?], 根据？的总得分，进行排序，
7. n = 1 << 6