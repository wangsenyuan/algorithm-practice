# Domino Arrangement Problem

## Problem Description

You have a set of dominoes. Each domino is a rectangular tile with a line dividing its face into two square ends. Can you put all dominoes in a line one by one from left to right so that any two dominoes touch with the sides that had the same number of points? You can rotate the dominoes, changing the left and the right side (domino "1-4" turns into "4-1").

## Input

- The first line contains the number **n** (1 ≤ n ≤ 100)
- Next **n** lines contain the dominoes
- Each of these lines contains two numbers — the number of points (spots) on the left and the right half, correspondingly
- The numbers of points (spots) are non-negative integers from 0 to 6
- Duplicates (identical tiles) may be present in the given set

## Output

Print "No solution", if it is impossible to arrange the dominoes in the required manner.

If the solution exists, then describe any way to arrange the dominoes. You put the dominoes from left to right. In each of the **n** lines print:
- The index of the domino to put in the corresponding position
- Then, after a space, the character:
  - **"+"** (if you don't need to turn the domino)
  - **"–"** (if you need to turn it)