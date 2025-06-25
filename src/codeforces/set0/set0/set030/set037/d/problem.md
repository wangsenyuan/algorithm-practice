# Problem D: Timetable

## Problem Description

When Petya has free time from computer games, he attends university classes. Every day the lessons on Petya's faculty consist of two double classes. The floor where the lessons take place is a long corridor with M classrooms numbered from 1 to M, situated along it.

All the students of Petya's year are divided into N groups. Petya has noticed recently that these groups' timetable has the following peculiarity: the number of the classroom where the first lesson of a group takes place does not exceed the number of the classroom where the second lesson of this group takes place.

Once Petya decided to count the number of ways in which one can make a lesson timetable for all these groups. The timetable is a set of 2N numbers: for each group the number of the rooms where the first and the second lessons take place. Unfortunately, he quickly lost track of his calculations and decided to count only the timetables that satisfy the following conditions:

1. On the first lesson in classroom i exactly Xi groups must be present.
2. In classroom i no more than Yi groups may be placed.

Help Petya count the number of timetables satisfying all those conditions. As there can be a lot of such timetables, output modulo $10^9 + 7$.

## Input

The first line contains one integer M (1 ≤ M ≤ 100) — the number of classrooms.

The second line contains M space-separated integers — Xi (0 ≤ Xi ≤ 100) the amount of groups present in classroom i during the first lesson.

The third line contains M space-separated integers — Yi (0 ≤ Yi ≤ 100) the maximal amount of groups that can be present in classroom i at the same time.

It is guaranteed that all the Xi ≤ Yi, and that the sum of all the Xi is positive and does not exceed 1000.

## Output

In the single line output the answer to the problem modulo $10^9 + 7$.

## ideas
1. 真一头雾水
2. 这里有多少个人（N）并没有给定，所以也不知道有多少个groups
3. 但是N = sum(X[i]) 因为第一堂课程，大家应该是同时参加的
4. 然后考虑第二节课程， 这里Y[i]有点歧义，按照 *In classroom i no more than Yi groups may be placed* 这个定义，那么Y[i]是两堂课程同时算入？
5. 比如第一堂课，a在教室1，第二堂课他还在教室1，那么他应该计入1，如果他第二堂课程在教室2，那么Y[1], Y[2]都要计入？
6. 但是按照 *the maximal amount of groups that can be present in classroom i at the same time* 那么这个就是指在教室i，两堂课程，各自都不能超过Y[i]； 似乎这个理解更好处理
7. dp[i][w] 表示在满足前w个教室要求的情况下的计数, w表示从前i个教室中参加完第一堂课程后，第二堂课程要在i后面参加
8. dp[i][w] => 在第i个教室增加了x[i]个人，w + x[i], 其中有u个人留下来参加了第二堂课程，那么状态就变成了 dp[i+1][w+x[i] - u]
9. 其中u <= y[i]
10. m * W * Y = 100 * 100 * W (W是可以到n的)，也就是1000?
11. dp[n][0]
12. YEAH