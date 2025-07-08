# Robot Movement Problem

## Problem Description

Fox Ciel has a robot on a 2D plane. Initially, it is located at position (0, 0). Fox Ciel codes a command for the robot, represented by string `s`. Each character of `s` is one move operation. There are four move operations in total:

- **'U'**: go up, (x, y) → (x, y+1)
- **'D'**: go down, (x, y) → (x, y-1)  
- **'L'**: go left, (x, y) → (x-1, y)
- **'R'**: go right, (x, y) → (x+1, y)

The robot will execute the operations in `s` from left to right, and repeat them infinitely. Help Fox Ciel determine if after some steps the robot will be located at position (a, b).

## Input

- The first line contains two integers `a` and `b` (-10^9 ≤ a, b ≤ 10^9)
- The second line contains a string `s` (1 ≤ |s| ≤ 100, s only contains characters 'U', 'D', 'L', 'R') — the command

## Output

Print "Yes" if the robot will be located at (a, b), and "No" otherwise.