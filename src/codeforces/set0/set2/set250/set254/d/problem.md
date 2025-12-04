Rats have bred to hundreds and hundreds in the basement of the store, owned by Vasily Petrovich. Vasily Petrovich may have not noticed their presence, but they got into the habit of sneaking into the warehouse and stealing food from there. Vasily Petrovich cannot put up with it anymore, he has to destroy the rats in the basement. Since mousetraps are outdated and do not help, and rat poison can poison inattentive people as well as rats, he chose a radical way: to blow up two grenades in the basement (he does not have more).

In this problem, we will present the shop basement as a rectangular table of $n \times m$ cells. Some of the cells are occupied by walls, and the rest of them are empty. Vasily has been watching the rats and he found out that at a certain time they go to sleep, and all the time they sleep in the same places. He wants to blow up a grenade when this convenient time comes. On the plan of his basement, he marked cells with sleeping rats in them. Naturally, these cells are not occupied by walls.

Grenades can only blow up in a cell that is not occupied by a wall. The blast wave from a grenade distributes as follows. We assume that the grenade blast occurs at time 0. During this initial time only the cell where the grenade blew up gets 'clear'. If at time $t$ some cell is clear, then at time $t + 1$ those side-neighbouring cells which are not occupied by the walls get clear too (some of them could have been cleared before). The blast wave distributes for exactly $d$ seconds, then it dies immediately.

An example of a distributing blast wave: Picture 1 shows the situation before the blast, and the following pictures show "clear" cells by time 0, 1, 2, 3 and 4. Thus, the blast wave on the picture distributes for $d = 4$ seconds.

Vasily Petrovich wonders, whether he can choose two cells to blast the grenades so as to clear all cells with sleeping rats. Write the program that finds it out.

## Input

The first line contains three integers $n$, $m$ and $d$, separated by single spaces ($4 \le n, m \le 1000$, $1 \le d \le 8$). Next $n$ lines contain the table that represents the basement plan. Each row of the table consists of $m$ characters. Character "X" means that the corresponding cell is occupied by the wall, character "." represents an empty cell, character "R" represents an empty cell with sleeping rats.

It is guaranteed that the first and the last row, as well as the first and the last column consist of characters "X". The plan has at least two empty cells. There is at least one cell with sleeping rats.

## Output

If it is impossible to blow up all cells with sleeping rats, print a single integer -1. Otherwise, print four space-separated integers $r_1$, $c_1$, $r_2$, $c_2$, that mean that one grenade should go off in cell $(r_1, c_1)$, and the other one — in cell $(r_2, c_2)$.

Consider the table rows numbered from top to bottom from 1 to $n$ and the table columns — from left to right from 1 to $m$. As $r_1$ and $r_2$ represent the row numbers, and $c_1$ and $c_2$ represent the column numbers in the table, they should fit the limits: $1 \le r_1, r_2 \le n$, $1 \le c_1, c_2 \le m$. It is forbidden to blow a grenade twice in the same cell. The blast waves of the grenades can intersect. It is possible that one grenade blast destroys no rats, and the other one destroys all of them.

## Examples

### Example 1

**Input**
```
4 4 1
XXXX
XR.X
X.RX
XXXX
```

**Output**
```
2 2 2 3
```

### Example 2

**Input**
```
9 14 5
XXXXXXXXXXXXXX
X....R...R...X
X..R.........X
X....RXR..R..X
X..R...X.....X
XR.R...X.....X
X....XXR.....X
X....R..R.R..X
XXXXXXXXXXXXXX
```

**Output**
```
2 3 6 9
```

### Example 3

**Input**
```
7 7 1
XXXXXXX
X.R.R.X
X.....X
X..X..X
X..R..X
X....RX
XXXXXXX
```

**Output**
```
-1
```


## ideas
1. 找出两个位置，能够在d时间内，覆盖到所有的老鼠
2. dp[r][c][d] 表示在时间d内，能够覆盖的老鼠的数量？好像不大行，因为可能有重叠的部分
3. 考虑最边上的rat，肯定需要一个炸弹去覆盖到它们
4. 考虑两个rat, 它们的炸弹埋点不重叠（没有一个炸弹埋点能同时砸到它们）
5. 那么两个炸弹能够放置的范围基本确定了
6. 这个范围，大概是 2 * d * 2 * d = 4 * d * d ~ 130
7. 所以最后的结果  = 130 * 130 * 130 * 130 = 1e8?
8. 这两个rat也比较好找，rat的数量，不能超过 4 * d * d * 2(否则肯定clear不了)
9. 那么找到距离最远的两个rat