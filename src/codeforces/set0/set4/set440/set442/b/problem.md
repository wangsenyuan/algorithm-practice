# Problem B: Andrey's Programming Contest

## Problem Description

Andrey needs one more problem to conduct a programming contest. He has n friends who are always willing to help. He can ask some of them to come up with a contest problem. Andrey knows one value for each of his friends — the probability that this friend will come up with a problem if Andrey asks him.

Help Andrey choose people to ask. As he needs only one problem, Andrey is going to be really upset if no one comes up with a problem or if he gets more than one problem from his friends. You need to choose such a set of people that maximizes the chances of Andrey not getting upset.

## Input

The first line contains a single integer **n** (1 ≤ n ≤ 100) — the number of Andrey's friends.

The second line contains **n** real numbers **p_i** (0.0 ≤ p_i ≤ 1.0) — the probability that the i-th friend can come up with a problem. The probabilities are given with at most 6 digits after decimal point.

## Output

Print a single real number — the probability that Andrey won't get upset at the optimal choice of friends. The answer will be considered valid if it differs from the correct one by at most 10^-9.

## Examples

### Example 1
**Input:**
```
4
0.1 0.2 0.3 0.8
```

**Output:**
```
0.800000000000
```

**Explanation:** The best strategy for Andrey is to ask only one of his friends, the most reliable one.

### Example 2
**Input:**
```
2
0.1 0.2
```

**Output:**
```
0.260000000000
```

**Explanation:** The best strategy for Andrey is to ask all of his friends to come up with a problem. Then the probability that he will get exactly one problem is 0.1 × 0.8 + 0.9 × 0.2 = 0.26.

## ideas
1. 按照概率降序排列？
2. dp[i][x] 表示到i为止，得到x个问题的概率
3. dp[i][x] = dp[i-1][x] * (1 - pi) + dp[i-1][x-1] * pi
4. 那么1的概率 = sum(1...) - sum(2...) ?
5. dp[i][0] = 前面的人都不提供 = (1 - pi) *
6. dp[i][1] = 这个人提供 = 