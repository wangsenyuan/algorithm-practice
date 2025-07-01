# Grid Maze Problem

## Problem Description

Pavel loves grid mazes. A grid maze is an n × m rectangle maze where each cell is either empty, or is a wall. You can go from one cell to another only if both cells are empty and have a common side.

Pavel drew a grid maze with all empty cells forming a connected area. That is, you can go from any empty cell to any other one. Pavel doesn't like it when his maze has too little walls. He wants to turn exactly k empty cells into walls so that all the remaining cells still formed a connected area. Help him.

## Input

The first line contains three integers n, m, k (1 ≤ n, m ≤ 500, 0 ≤ k < s), where:
- n and m are the maze's height and width, correspondingly
- k is the number of walls Pavel wants to add
- s represents the number of empty cells in the original maze

Each of the next n lines contains m characters. They describe the original maze:
- If a character on a line equals `.`, then the corresponding cell is empty
- If a character equals `#`, then the cell is a wall

## Output

Print n lines containing m characters each: the new maze that fits Pavel's requirements. Mark the empty cells that you transformed into walls as `X`, the other cells must be left without changes (that is, `.` and `#`).

## Constraints

- 1 ≤ n, m ≤ 500
- 0 ≤ k < s (where s is the number of empty cells)

## Notes

- It is guaranteed that a solution exists
- If there are multiple solutions you can output any of them

## ideas
1. 感觉还有点麻烦的
2. 把k个转换掉，那么剩余s - k 个就是联通的，那么先标记这部分，剩余的部分再标记成X
3. 