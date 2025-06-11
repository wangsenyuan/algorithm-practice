# Domino Puzzle

## Problem Description

Xenia likes puzzles very much. She is especially fond of the puzzles that consist of domino pieces. Look at the picture that shows one of such puzzles.

A puzzle is a 3 × n table with forbidden cells (black squares) containing dominoes (colored rectangles on the picture). A puzzle is called correct if it meets the following conditions:

- each domino occupies exactly two non-forbidden cells of the table;
- no two dominoes occupy the same table cell;
- exactly one non-forbidden cell of the table is unoccupied by any domino (it is marked by a circle in the picture).

To solve the puzzle, you need multiple steps to transport an empty cell from the starting position to some specified position. A move is transporting a domino to the empty cell, provided that the puzzle stays correct. The horizontal dominoes can be moved only horizontally, and vertical dominoes can be moved only vertically. You can't rotate dominoes. The picture shows a probable move.

## Task

Xenia has a 3 × n table with forbidden cells and a cell marked with a circle. Also, Xenia has very many identical dominoes. Now Xenia is wondering, how many distinct correct puzzles she can make if she puts dominoes on the existing table. Also, Xenia wants the circle-marked cell to be empty in the resulting puzzle. The puzzle must contain at least one move.

Help Xenia, count the described number of puzzles. As the described number can be rather large, print the remainder after dividing it by 1000000007 (10^9 + 7).

## Input

The first line contains integer n (3 ≤ n ≤ 10^4) — the puzzle's size. 

Each of the following three lines contains n characters — the description of the table. The j-th character of the i-th line equals:
- "X" if the corresponding cell is forbidden
- "." if the corresponding cell is non-forbidden 
- "O" if the corresponding cell is marked with a circle

**Constraints:**
- It is guaranteed that exactly one cell in the table is marked with a circle
- It is guaranteed that all cells of a given table having at least one common point with the marked cell is non-forbidden

## Output

Print a single number — the answer to the problem modulo 1000000007 (10^9 + 7).

## ideas
1. 3 * n, 感觉要从两头开始，到O所在的列
2. dp[i][state] 表示从左到右，到第i列时，它的状态为state时的计数
3. state表示从上到下3个格子的状态。（1表示被占用，0表示还需要填入）
4. 且在处理第i列的时候，前面的所有的个字都被占满了
5. 如果在第一行放置一个横向的（如果可以），得到一个新的状态
6. 如果dp[i] state的第一位是空的，那么这里必须放置一个横向的（把它给填满）
7. 这个地方感觉还是的枚举，放置/不放置，横向/竖向，的牌的状态
8. 但必须保证i的被填满