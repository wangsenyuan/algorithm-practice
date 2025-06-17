# Problem Statement

Berland SU holds yet another training contest for its students today. $n$ students came, each of them brought his laptop. However, it turned out that everyone has forgot their chargers!

## Problem Details

Let students be numbered from $1$ to $n$. Laptop of the $i$-th student has charge $a_i$ at the beginning of the contest and it uses $b_i$ of charge per minute (i.e. if the laptop has $c$ charge at the beginning of some minute, it becomes $c-b_i$ charge at the beginning of the next minute). The whole contest lasts for $k$ minutes.

Polycarp (the coach of Berland SU) decided to buy a single charger so that all the students would be able to successfully finish the contest. He buys the charger at the same moment the contest starts.

### Charger Rules

Polycarp can choose to buy the charger with any non-negative (zero or positive) integer power output. The power output is chosen before the purchase, it can't be changed afterwards. Let the chosen power output be $x$. At the beginning of each minute (from the minute contest starts to the last minute of the contest) he can plug the charger into any of the student's laptops and use it for some integer number of minutes.

If the laptop is using $b_i$ charge per minute then it will become $b_i-x$ per minute while the charger is plugged in. Negative power usage rate means that the laptop's charge is increasing. The charge of any laptop isn't limited, it can become infinitely large. The charger can be plugged in no more than one laptop at the same time.

The student successfully finishes the contest if the charge of his laptop never is below zero at the beginning of some minute (from the minute contest starts to the last minute of the contest, zero charge is allowed). The charge of the laptop of the minute the contest ends doesn't matter.

Help Polycarp to determine the minimal possible power output the charger should have so that all the students are able to successfully finish the contest. Also report if no such charger exists.

## Input

The first line contains two integers $n$ and $k$ ($1 \leq n \leq 2 \cdot 10^5$, $1 \leq k \leq 2 \cdot 10^5$) — the number of students (and laptops, correspondingly) and the duration of the contest in minutes.

The second line contains $n$ integers $a_1, a_2, \ldots, a_n$ ($1 \leq a_i \leq 10^{12}$) — the initial charge of each student's laptop.

The third line contains $n$ integers $b_1, b_2, \ldots, b_n$ ($1 \leq b_i \leq 10^7$) — the power usage of each student's laptop.

## Output

Print a single non-negative integer — the minimal possible power output the charger should have so that all the students are able to successfully finish the contest.

If no such charger exists, print -1.

## ideas
1. 符合二分的性质。
2. 对于x，如何检查它是否ok？
3. 在整个过程中，不应该有不充电的时刻（因为充电肯定更有利）
4. 在任意时刻t，应该选择哪个laptop去充电呢？电量最低的吗？
5. 应该是的，但是有个问题，就是不能够所有的都去更新
6. 假设在时刻t，那么所有的laptop的电量，都需要更新为 wi - bi (不充电的情况下)
7. 如果有充电的情况 wi - bi + x
8. 在这些值中，如果出现了负值 -> false
9. 如果不考虑x，那么在时刻t, wi = ai - t * bi
10. 在任意t时刻，如果出现了负值，那么就应该在足够早的时间内，给它补充(相当于增加ai)
11. 这里t是个变量，
12. ai - t * bi < 0 => let ti =  ai / bi, 那么ti就是i必须要充电的最晚时间
13. 如果ti >= k => good
14. 否则，就要给它充电 dx, 至少要能让ti增长
15. (ai + dx) / bi = ti + 1
16. dx = bi * (ti + 1) - ai
17. 进而算出dt, dt就是当前需要充电的最小值