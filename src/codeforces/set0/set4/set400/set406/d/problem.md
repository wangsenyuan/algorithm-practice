# Hill Climbers Problem

This problem has nothing to do with Little Chris. It is about hill climbers instead (and Chris definitely isn't one).

## Problem Description

There are n hills arranged on a line, each in the form of a vertical line segment with one endpoint on the ground. The hills are numbered with numbers from 1 to n from left to right. The i-th hill stands at position xi with its top at height yi. 

For every two hills a and b, if the top of hill a can be seen from the top of hill b, their tops are connected by a rope. Formally, the tops of two hills are connected if the segment connecting their top points does not intersect or touch any of the other hill segments. Using these ropes, the hill climbers can move from hill to hill.

There are m teams of climbers, each composed of exactly two members. The first and the second climbers of the i-th team are located at the top of the ai-th and bi-th hills, respectively. They want to meet together at the top of some hill. 

Now, each of two climbers move according to the following process:

- if a climber is at the top of the hill where the other climber is already located or will come eventually, the former climber stays at this hill;
- otherwise, the climber picks a hill to the right of his current hill that is reachable by a rope and is the rightmost possible, climbs this hill and continues the process (the climber can also climb a hill whose top is lower than the top of his current hill).

For each team of climbers, determine the number of the meeting hill for this pair!

## Input

The first line of input contains a single integer n (1 ≤ n ≤ 10⁵), the number of hills. 

The next n lines describe the hills. The i-th of them contains two space-separated integers xi, yi (1 ≤ xi ≤ 10⁷; 1 ≤ yi ≤ 10¹¹), the position and the height of the i-th hill. The hills are given in the ascending order of xi, i.e., xi < xj for i < j.

The next line of input contains a single integer m (1 ≤ m ≤ 10⁵), the number of teams. 

The next m lines describe the teams. The i-th of them contains two space-separated integers ai, bi (1 ≤ ai, bi ≤ n), the numbers of the hills where the climbers of the i-th team are located. It is possible that ai = bi.

## Output

In a single line output m space-separated integers, where the i-th integer is the number of the meeting hill for the members of the i-th team.


## ideas
1. 这是一个关于lcp的问题
2. 问题的关键在于如何构造这棵树
3. 对于j来说，就是找到最近的i，在j...i中间，
4. 把它们连起来后，它们中间的点, 都在这条线的下方, 
5. 这条线可以用 y = x * a + b
6. yi = xi * a + b
7. yj = xj * a + b
8. a = (yi -yj) / (xi - xj)
9. b ...
10. yk < xk * a + b (满足这个条件)
11. 感觉能够看到的是一个convex hull的上半边