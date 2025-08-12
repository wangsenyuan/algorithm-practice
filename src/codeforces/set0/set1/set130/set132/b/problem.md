# Problem B - Piet Programming Language

## Problem Description

Piet is one of the most known visual esoteric programming languages. The programs in Piet are constructed from colorful blocks of pixels and interpreted using pretty complicated rules. In this problem we will use a subset of Piet language with simplified rules.

## Program Structure

The program will be a rectangular image consisting of colored and black pixels. The color of each pixel will be given by an integer number between 0 and 9, inclusive, with 0 denoting black. A block of pixels is defined as a rectangle of pixels of the same color (not black). It is guaranteed that all connected groups of colored pixels of the same color will form rectangular blocks. Groups of black pixels can form arbitrary shapes.

## Instruction Pointer (IP) Components

The program is interpreted using movement of instruction pointer (IP) which consists of three parts:

- **Current Block Pointer (BP)**: Note that there is no concept of current pixel within the block
- **Direction Pointer (DP)**: Can point left, right, up or down
- **Block Chooser (CP)**: Can point to the left or to the right from the direction given by DP; in absolute values CP can differ from DP by 90 degrees counterclockwise or clockwise, respectively

## Initial State

Initially BP points to the block which contains the top-left corner of the program, DP points to the right, and CP points to the left (see the orange square on the image below).

## Program Execution Rules

One step of program interpretation changes the state of IP in the following way:

1. The interpreter finds the furthest edge of the current color block in the direction of the DP
2. From all pixels that form this edge, the interpreter selects the furthest one in the direction of CP
3. After this, BP attempts to move from this pixel into the next one in the direction of DP

**If the next pixel belongs to a colored block:**
- This block becomes the current one
- Two other parts of IP stay the same

**If the next pixel is black or outside of the program:**
- BP stays the same but two other parts of IP change
- If CP was pointing to the left, now it points to the right, and DP stays the same
- If CP was pointing to the right, now it points to the left, and DP is rotated 90 degrees clockwise

**Important Note:** BP will never point to a black block (it is guaranteed that top-left pixel of the program will not be black).

## Task

You are given a Piet program. You have to figure out which block of the program will be current after n steps.

## Input

The first line of the input contains two integer numbers m (1 ≤ m ≤ 50) and n (1 ≤ n ≤ 5·10^7). Next m lines contain the rows of the program. All the lines have the same length between 1 and 50 pixels, and consist of characters 0-9. The first character of the first line will not be equal to 0.

## Output

Output the color of the block which will be current after n steps of program interpretation.

## Examples

### Example 1
**Input:**
```
2 10
12
43
```
**Output:**
```
1
```

### Example 2
**Input:**
```
3 12
1423
6624
6625
```
**Output:**
```
6
```

### Example 3
**Input:**
```
5 9
10345
23456
34567
45678
56789
```
**Output:**
```
5
```

## Note

In the first example IP changes in the following way:
- After step 1: block 2 becomes current one and stays it after two more steps
- After step 4: BP moves to block 3
- After step 7: BP moves to block 4
- Finally after step 10: BP returns to block 1

The sequence of states of IP is shown on the image: the arrows are traversed clockwise, the main arrow shows direction of DP, the side one — the direction of CP.


### ideas
1. 所有非黑色的块，都是矩形（规则，只需要关注边即可）
2. 假设当前BP块，那么下一个块的规则是: 设定DP/CP的方向起始为(R/T)
3.  如果对应位置（和最右边上边界接触的位置）没有新的block，那么变成(R/D)
4.  如果对应位置（和最右边下边界接触的位置) 没有新的block，那么变成(D/R)
5.  如果对应位置 (和最下边右边界接触的位置) 没有新的block， 那么变成(D/L)
6.  如果对应位置 (和最下边左边界接触的位置) 没有新的block, 那么变成(L/D)
7.  如果对应位置 (和最左边下边界接触的位置) 没有新的block， 那么变成(L/T)
8.  如果对应位置 （和最左边上边界接触的位置）没有新的block，那么变成(T/L)
9.  如果对应位置 (和最上边左边界接触的位置) 没有新的block, 那么变成(T/R)
10. 如果对应位置（和最上边右边界接触的位置) 没有新的block，那么变成（R/T)
11. 所以有可能回到初始位置 dp[i][x][y] 表示到达位置block i时， DP = X， CP = y时的操作数
12. 然后如果再遇到这个状态，说明进入了循环