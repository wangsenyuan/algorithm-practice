# Dima and Inna's Chessboard Game

## Problem Description

Dima and Inna are doing so great! At the moment, Inna is sitting on the magic lawn playing with a pink pony. Dima wanted to play too. He brought an n × m chessboard, a very tasty candy and two numbers a and b.

Dima put the chessboard in front of Inna and placed the candy in position (i, j) on the board. The boy said he would give the candy if it reaches one of the corner cells of the board. He's got one more condition. There can only be actions of the following types:

- Move the candy from position (x, y) on the board to position (x - a, y - b)
- Move the candy from position (x, y) on the board to position (x + a, y - b)
- Move the candy from position (x, y) on the board to position (x - a, y + b)
- Move the candy from position (x, y) on the board to position (x + a, y + b)

Naturally, Dima doesn't allow to move the candy beyond the chessboard borders.

Inna and the pony started shifting the candy around the board. They wonder what is the minimum number of allowed actions that they need to perform to move the candy from the initial position (i, j) to one of the chessboard corners. Help them cope with the task!

## Input

The first line of the input contains six integers n, m, i, j, a, b (1 ≤ n, m ≤ 10^6; 1 ≤ i ≤ n; 1 ≤ j ≤ m; 1 ≤ a, b ≤ 10^6).

You can assume that the chessboard rows are numbered from 1 to n from top to bottom and the columns are numbered from 1 to m from left to right. Position (i, j) in the statement is a chessboard cell on the intersection of the i-th row and the j-th column. You can consider that the corners are: (1, m), (n, 1), (n, m), (1, 1).

## Output

In a single line print a single integer — the minimum number of moves needed to get the candy.

If Inna and the pony cannot get the candy playing by Dima's rules, print on a single line "Poor Inna and pony!" without the quotes.

## ideas
1. 考虑移动到位置(n, m)
2. 那么移动的距离是a的倍数，同时也是b的倍数
3. (n - i) / a = (m - j) / b
4. 可以左右移动，然后向下移动，消耗掉一些步骤
5. 