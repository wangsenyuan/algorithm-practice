# Grid Connected Components

## Problem Description

You have a grid with $n$ rows and $n$ columns. Each cell is either empty (denoted by '.') or blocked (denoted by 'X').

### Connected Cells Definition

Two empty cells are **directly connected** if they share a side. Two cells $(r_1, c_1)$ (located in row $r_1$ and column $c_1$) and $(r_2, c_2)$ are **connected** if there exists a sequence of empty cells that:
- Starts with $(r_1, c_1)$
- Finishes with $(r_2, c_2)$
- Any two consecutive cells in this sequence are directly connected

A **connected component** is a set of empty cells such that:
- Any two cells in the component are connected
- There is no cell in this set that is connected to some cell not in this set

### Limak's Special Ability

Your friend Limak is a big grizzly bear. He can destroy obstacles in a specific range:
- You can choose a square of size $k \times k$ in the grid
- Limak will transform all blocked cells in that square to empty ones
- You can ask Limak to help only **once**
- The chosen square must be completely inside the grid
- It's possible that Limak won't change anything if all cells are already empty

### Goal

You like big connected components. After Limak helps you, what is the **maximum possible size** of the biggest connected component in the grid?

## Input

- The first line contains two integers $n$ and $k$ $(1 \leq k \leq n \leq 500)$ — the size of the grid and Limak's range, respectively
- Each of the next $n$ lines contains a string with $n$ characters, denoting the $i$-th row of the grid
- Each character is either '.' (empty cell) or 'X' (blocked cell)

## Output

Print the maximum possible size (number of cells) of the biggest connected component after using Limak's help.

