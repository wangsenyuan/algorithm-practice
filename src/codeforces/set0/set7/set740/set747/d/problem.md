# Winter Tires Problem

## Problem Description

The winter in Berland lasts $n$ days. For each day we know the forecast for the average air temperature that day.

Vasya has a new set of winter tires which allows him to drive safely no more than $k$ days at any average air temperature. After $k$ days of using it (regardless of the temperature of these days) the set of winter tires wears down and cannot be used more. It is not necessary that these $k$ days form a continuous segment of days.

Before the first winter day Vasya still uses summer tires. It is possible to drive safely on summer tires any number of days when the average air temperature is non-negative. It is impossible to drive on summer tires at days when the average air temperature is negative.

Vasya can change summer tires to winter tires and vice versa at the beginning of any day.

**Find the minimum number of times Vasya needs to change summer tires to winter tires and vice versa to drive safely during the winter.** At the end of the winter the car can be with any set of tires.

## Input

The first line contains two positive integers $n$ and $k$ ($1 \leq n \leq 2 \cdot 10^5$, $0 \leq k \leq n$) — the number of winter days and the number of days winter tires can be used. It is allowed to drive on winter tires at any temperature, but no more than $k$ days in total.

The second line contains a sequence of $n$ integers $t_1, t_2, \ldots, t_n$ ($-20 \leq t_i \leq 20$) — the average air temperature in the $i$-th winter day.

## Output

Print the minimum number of times Vasya has to change summer tires to winter tires and vice versa to drive safely during all winter. If it is impossible, print $-1$.

## Examples

### Example 1
**Input:**
```
4 3
-5 20 -3 0
```

**Output:**
```
2
```

**Explanation:** Before the first winter day Vasya should change summer tires to winter tires, use it for three days, and then change winter tires to summer tires because he can drive safely with the winter tires for just three days. Thus, the total number of tires' changes equals two.

### Example 2
**Input:**
```
4 2
-5 20 -3 0
```

**Output:**
```
4
```

**Explanation:** Before the first winter day Vasya should change summer tires to winter tires, and then after the first winter day change winter tires to summer tires. After the second day it is necessary to change summer tires to winter tires again, and after the third day it is necessary to change winter tires to summer tires. Thus, the total number of tires' changes equals four.

### Example 3
**Input:**
```
10 6
2 -5 1 3 0 0 -4 -3 1 0
```

**Output:**
```
3
```

### ideas
1. 在零下温度出现时，必须换上冬季轮胎
2. 冬季轮胎最多使用k天，（如果零下温度的天数超过了k => -1)
3. 最小交换次数
4. 那么在现在使用夏季轮胎的情况下，只有在零下温度出现的情况下，才需要更换
5. 如果在使用冬季轮胎的情况下，只有当现在不更换（夏季），那么后续无法满足条件2的情况下，才需要更换
6. fp[i]表示在第i天如果使用冬季轮胎时，后面的零下温度天数
7. 那么也就是是当 fp[i] > m (剩余天数)时，且a[i] >= 0时，更换为夏季轮胎
8. 似乎不对。存在一种情况下，现在在尽量使用冬季轮胎。也满足条件，但是因为冬季轮胎k天用完后，必须更换
9. 但是如果交换一下，现在用夏季轮胎，有可能最后不用交换就可以
10. dp[i][0] 表示在i处使用夏季轮胎，切能够安全驶出时的最小交换次数
11. dp[i][1] 表示目前使用冬季胎
12. dp[i][0] = (a[i] >= 0) min(dp[i+1][0], dp[i+1] + 1)
13. dp[i][1] = min(dp[i+1][0] + 1, dp[i+1]) (但是必须满足一个k的条件)
14.    那怎么满足k的条件呢？
15. 现在使用夏季胎的情况下，尽量使用夏季胎，是没有问题的
16. dp[i][0] = dp[j][1] + 1 (j是next a[j] < 0)
17. dp[i][1] = min(dp[i+1][0] + 1, dp[j][1]) 如果这段内的天数，能够满足天气要求
18. 冬季轮胎，肯定是从某个 a[i] < 0 开始使用，
19. 如果a[j] 是下一个冷天，要么冬季轮胎一直用到j， 要么在下一天就更换为夏季轮胎