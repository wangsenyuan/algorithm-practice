# Problem B - Gerald's Game

## Problem Description

Gerald plays the following game. He has a checkered field of size n × n cells, where m various cells are banned. Before the game, he has to put a few chips on some border (but not corner) board cells. Then for n - 1 minutes, Gerald every minute moves each chip into an adjacent cell. He moves each chip from its original edge to the opposite edge.

## Losing Conditions

Gerald loses in this game in each of the three cases:

- At least one of the chips at least once fell to the banned cell.
- At least once two chips were on the same cell.
- At least once two chips swapped in a minute (for example, if you stand two chips on two opposite border cells of a row with even length, this situation happens in the middle of the row).

In that case he loses and earns 0 points. When nothing like that happened, he wins and earns the number of points equal to the number of chips he managed to put on the board. Help Gerald earn the most points.

## Input

The first line contains two space-separated integers n and m (2 ≤ n ≤ 1000, 0 ≤ m ≤ 10^5) — the size of the field and the number of banned cells. 

Next m lines each contain two space-separated integers. Specifically, the i-th of these lines contains numbers xi and yi (1 ≤ xi, yi ≤ n) — the coordinates of the i-th banned cell. All given cells are distinct.

Consider the field rows numbered from top to bottom from 1 to n, and the columns — from left to right from 1 to n.

## Output

Print a single integer — the maximum points Gerald can earn in this game.

## ideas
1. 有点不好理解
2. He moves each chip from its original edge to the opposite edge 这个表示啥？
3. 也就是说，左边的只能一路向右，上边的只能一路向下
4. 所以条件2，就是说，同一行（或者同一列）最多只能有一个
5. 然后(从左边开始的第i列，和从上到下的第i行，不能同时有)
6. 