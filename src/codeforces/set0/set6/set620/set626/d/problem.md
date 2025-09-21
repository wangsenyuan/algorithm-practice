# Problem D: Ball Game Probability

## Problem Statement

Andrew and Jerry are playing a game with Harry as the scorekeeper. The game consists of three rounds. In each round, Andrew and Jerry draw randomly without replacement from a jar containing n balls, each labeled with a distinct positive integer. Without looking, they hand their balls to Harry, who awards the point to the player with the larger number and returns the balls to the jar. The winner of the game is the one who wins at least two of the three rounds.

Andrew wins rounds 1 and 2 while Jerry wins round 3, so Andrew wins the game. However, Jerry is unhappy with this system, claiming that he will often lose the match despite having the higher overall total. 

**What is the probability that the sum of the three balls Jerry drew is strictly higher than the sum of the three balls Andrew drew?**

## Input

The first line of input contains a single integer **n** (2 ≤ n ≤ 2000) — the number of balls in the jar.

The second line contains **n** integers **ai** (1 ≤ ai ≤ 5000) — the number written on the ith ball. It is guaranteed that no two balls have the same number.

## Output

Print a single real value — the probability that Jerry has a higher total, given that Andrew wins the first two rounds and Jerry wins the third. Your answer will be considered correct if its absolute or relative error does not exceed 10^-6.

Namely: let's assume that your answer is **a**, and the answer of the jury is **b**. The checker program will consider your answer correct, if |a - b| / max(1, |b|) ≤ 10^-6.

## Examples

### Example 1
**Input:**
```
2
1 2
```

**Output:**
```
0.0000000000
```

### Example 2
**Input:**
```
3
1 2 10
```

**Output:**
```
0.0740740741
```

## Explanation

**Example 1:** There are only two balls. In the first two rounds, Andrew must have drawn the 2 and Jerry must have drawn the 1, and vice versa in the final round. Thus, Andrew's sum is 5 and Jerry's sum is 4, so Jerry never has a higher total.

**Example 2:** Each game could've had three outcomes — 10-2, 10-1, or 2-1. Jerry has a higher total if and only if Andrew won 2-1 in both of the first two rounds, and Jerry drew the 10 in the last round. This has probability 1/27 ≈ 0.0740740741.


### ideas
1. 前两轮Andrew要得到更高的分数，第三轮Jerry得到更多的分数，且Jerry的总分数大于Andrew的分数
2. 将a升序排列; dp[x] 在一轮结束， Andrew的分数 - Jerry 的分数 = x的概率
3. dp2[x] = 两轮后差为x的概率；
4. 结果就等于 dp2[x] * dp[y] where y > x
5. dp[x] 比较好搞, 这个可以通过 n * n 的处理得到
6. 但是 dp2[x] 咋搞？
7. a[i] <= 5000, 所以，x的范围不会超过5000, 那么好搞了