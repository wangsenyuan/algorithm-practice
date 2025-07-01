# Problem D: Stone Movement Game

## Problem Description

You are playing the following game. There are n points on a plane. They are the vertices of a regular n-polygon. Points are labeled with integer numbers from 1 to n. Each pair of distinct points is connected by a diagonal, which is colored in one of 26 colors. Points are denoted by lowercase English letters. There are three stones positioned on three distinct vertices. All stones are the same. With one move you can move the stone to another free vertex along some diagonal. The color of this diagonal must be the same as the color of the diagonal, connecting another two stones.

Your goal is to move stones in such way that the only vertices occupied by stones are 1, 2 and 3. You must achieve such position using minimal number of moves. Write a program which plays this game in an optimal way.

## Input

- **First line**: One integer n (3 ≤ n ≤ 70) — the number of points
- **Second line**: Three space-separated integers from 1 to n — numbers of vertices, where stones are initially located
- **Next n lines**: Each line contains n symbols — the matrix denoting the colors of the diagonals
  - Colors are denoted by lowercase English letters
  - The symbol j of line i denotes the color of diagonal between points i and j
  - Matrix is symmetric, so j-th symbol of i-th line is equal to i-th symbol of j-th line
  - Main diagonal is filled with '*' symbols because there is no diagonal, connecting point to itself

## Output

- If there is no way to put stones on vertices 1, 2 and 3, print `-1` on a single line
- Otherwise:
  - **First line**: Minimal required number of moves
  - **Next lines**: Description of each move, one move per line
  - To describe a move print two integers: the point from which to remove the stone, and the point to which move the stone
  - If there are several optimal solutions, print any of them

## ideas
1. 假设当前3个石头所在的位置是（a, b, c), 现在要移动到新的位置(d, b, c)
2. 必须color[a][d] = color[b][c] 
3. n * n * n 似乎也ok