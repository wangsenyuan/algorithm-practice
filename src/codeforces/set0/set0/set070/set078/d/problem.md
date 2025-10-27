A breakthrough among computer games, "Civilization XIII", is striking in its scale and elaborate details. Let's take a closer look at one of them.

The playing area in the game is split into congruent cells that are regular hexagons. The side of each cell is equal to 1. Each unit occupies exactly one cell of the playing field. The field can be considered infinite.

Let's take a look at the battle unit called an "Archer". Each archer has a parameter "shot range". It's a positive integer that determines the radius of the circle in which the archer can hit a target. The center of the circle coincides with the center of the cell in which the archer stays. A cell is considered to be under the archer's fire if and only if all points of this cell, including border points are located inside the circle or on its border.

The picture below shows the borders for shot ranges equal to 3, 4 and 5. The archer is depicted as A.

Find the number of cells that are under fire for some archer.

## Input

The first and only line of input contains a single positive integer k — the archer's shot range (1 ≤ k ≤ 10^6).

## Output

Print the single number, the number of cells that are under fire.

Note: Please do not use the %lld specificator to read or write 64-bit integers in C++. It is preferred to use the cout stream (also you may use the %I64d specificator).

## Examples

### Example 1
```
Input
3

Output
7
```

### Example 2
```
Input
4

Output
13
```

### Example 3
```
Input
5

Output
19
```


### ideas
1. 符合二分的性质，但是怎么check呢？