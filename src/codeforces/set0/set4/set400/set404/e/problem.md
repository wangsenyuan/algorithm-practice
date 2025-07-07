# Problem E: Robot Instructions

## Problem Description

Valera has a strip infinite in both directions and consisting of cells. The cells are numbered by integers. The cell number 0 has a robot.

The robot has instructions — the sequence of moves that he must perform. In one move, the robot moves one cell to the left or one cell to the right, according to instructions. Before the robot starts moving, Valera puts obstacles in some cells of the strip, excluding cell number 0. If the robot should go into the cell with an obstacle according the instructions, it will skip this move.

Also Valera indicates the finish cell in which the robot has to be after completing the entire instructions. The finishing cell should be different from the starting one. It is believed that the robot completed the instructions successfully, if during the process of moving he visited the finish cell exactly once — at its last move. Moreover, the latter move cannot be skipped.

Let's assume that k is the minimum number of obstacles that Valera must put to make the robot able to complete the entire sequence of instructions successfully and end up in some finishing cell. You need to calculate in how many ways Valera can choose k obstacles and the finishing cell so that the robot is able to complete the instructions successfully.

## Input

The first line contains a sequence of characters without spaces s₁s₂...sₙ (1 ≤ n ≤ 10⁶), consisting only of letters "L" and "R". If character sᵢ equals "L", then the robot on the i-th move must try to move one cell to the left. If the sᵢ-th character equals "R", then the robot on the i-th move must try to move one cell to the right.

## Output

Print a single integer — the required number of ways. It's guaranteed that this number fits into 64-bit signed integer type.

## Examples

### Example 1
**Input:**
```
RR
```

**Output:**
```
1
```

**Explanation:** Valera mustn't add any obstacles and his finishing cell must be cell 2.

### Example 2
**Input:**
```
RRL
```

**Output:**
```
1
```

**Explanation:** Valera must add an obstacle in cell number 1, and his finishing cell must be cell number -1. In this case robot skips the first two moves and on the third move he goes straight from the starting cell to the finishing one. But if Valera doesn't add any obstacles, or adds an obstacle to another cell, then the robot visits the finishing cell more than once.


## ideas
1. k貌似要么等于0（不增加任何限制），等于1，只在一边增加限制，或者2，两边增加限制（这个似乎不可能）
2. 不增加限制的话，就是检查是否能够最后停在一个（只有最后才能访问到的位置）
3. 增加1个的时候，有不同的位置可以选
4. 假设在左边位置l0增加限制后，最后停在某个位置r0，那么似乎在l0+1处放置也是可行的
5. 假设在放置任何限制前，能够到达的位置为x
6. 那么新的位置 x[l] = x + ？
7. 如果在l的左边，有w次先左移动, 是不是增加的距离 = w?
8. 