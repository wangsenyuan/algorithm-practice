# G - Catch All Apples (ABC457)

**Contest:** [ABC457](https://atcoder.jp/contests/abc457) — AtCoder Beginner Contest 457  
**Task:** [https://atcoder.jp/contests/abc457/tasks/abc457_g](https://atcoder.jp/contests/abc457/tasks/abc457_g)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 625 points

## Problem Statement

N apples fall on a number line. Apple i falls at coordinate X_i at time T_i.

You want to place some robots on the number line to collect all N apples. The robots can be placed at any coordinates.

Each robot starts operating from time 0 and can move freely along the number line at a speed of at most 1. Multiple robots may occupy the same coordinate at the same time. Each robot can collect apple i if and only if it is at coordinate X_i at time T_i.

Find the minimum number of robots needed to collect all apples.

## Constraints

- 1 ≤ N ≤ 3 × 10^5
- 0 ≤ T_i ≤ 3 × 10^5
- 0 ≤ X_i ≤ 3 × 10^5
- (T_i, X_i) ≠ (T_j, X_j) (i ≠ j)
- All input values are integers.

## Input

The input is given from Standard Input in the following format:

```text
N
T_1 X_1
T_2 X_2
⋮
T_N X_N
```

## Output

Output the answer.

## Sample Input 1

```text
4
0 2
1 0
2 1
2 3
```

## Sample Output 1

```text
2
```

All apples can be collected with two robots by moving them as follows:

- Place robot 1 at coordinate 0 and robot 2 at coordinate 2.
- Time 0: Robot 2 collects apple 1.
- Time 1: Robot 1 collects apple 2. Move both robots in the positive direction at speed 1 until time 2.
- Time 2: Robot 1 collects apple 3 and robot 2 collects apple 4.

It is impossible to collect all apples with fewer than two robots, so output 2.

## Sample Input 2

```text
5
0 1
0 2
0 3
0 4
0 5
```

## Sample Output 2

```text
5
```

## Sample Input 3

```text
8
10 4
4 2
7 10
5 3
1 9
0 6
3 8
0 9
```

## Sample Output 3

```text
2
```

### ideas
1. 最差的情况是使用n个机器人，如果i,j， abs(x[i] - x[j]) = abs(t[i] - t[j]) 那么它们可以使用同一个机器人
2. 按照T升序处理，那么 abs(x[i] - x[j]) = t[i] - t[j]
3. 如果x[i] > x[j], x[i] - t[i] = x[j] - t[j]
4. 如果x[i] < x[j], x[j] + t[j] = t[i] + x[i]
5. 按照t升序处理，如果它左边有一个j，满足第一个条件，那么就可以省掉一个机器人
6. 会存在多个吗？比如 x[i] - x[j] = t[i] - t[j]
7. x[i] - x[k] = t[i] - t[k] 
8. x[j] - x[k] = x[j] - t[k]
9. 看起来似乎是j在(k, i)的中间
10. 但是，k在i的右边的时候，问题就变得复杂了（既可以和j匹配，也可以和k匹配）
11. 但是如果和j匹配，那么后期j的左边（本来可以和j匹配的，就无法被匹配了）
12. 但是如果i和k匹配，那么j就可以和l匹配，从而省下一个机器人
13. 这个时候似乎又必须从整体考虑，要从左到右考虑，
14. 把它们当作平面坐标系种的点，就是用斜率为1，或者-1（可以直角拐弯）线去覆盖这些点
15. 但是不能回退（只能沿着y轴向上）
16. 机器人可以停止不动，也就是说，它可以往上
17.
