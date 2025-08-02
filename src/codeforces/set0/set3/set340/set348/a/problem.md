# Mafia Game Problem

## Problem Description

One day n friends gathered together to play "Mafia". During each round of the game some player must be the supervisor and other n - 1 people take part in the game. For each person we know in how many rounds he wants to be a player, not the supervisor: the i-th person wants to play ai rounds. What is the minimum number of rounds of the "Mafia" game they need to play to let each person play at least as many rounds as they want?

## Input

The first line contains integer n (3 ≤ n ≤ 10^5). The second line contains n space-separated integers a1, a2, ..., an (1 ≤ ai ≤ 10^9) — the i-th number in the list is the number of rounds the i-th person wants to play.

## Output

In a single line print a single integer — the minimum number of game rounds the friends need to let the i-th person play at least ai rounds.

**Note:** Please, do not use the %lld specifier to read or write 64-bit integers in С++. It is preferred to use the cin, cout streams or the %I64d specifier.

## Examples

### Example 1

**Input:**
```
3
3 2 2
```

**Output:**
```
4
```

### Example 2

**Input:**
```
4
2 2 2 2
```

**Output:**
```
3
```


## ideas
1. 假设共有m天可以参与完耍，那么m >= max(a[i])
2. a[i]升序排列
3. 考虑某个人当了w天的队长（这些天不能算做他玩的天数），那么m >= w[i] + a[i]
4. 也就是说，如果一共可以玩耍m天，那么对于a[1]来说，它可以在m - a[1]天内作为队长
5. a[2]来说，也是 m - a[2]天。。。
6. 那么就要有 m - a[1] + m - a[2] + ... m - a[i] >= m
7. n * m - sum >= m => (n - 1) * m >= sum
8. 