The Robot is in a rectangular maze of size `n x m`. Each cell of the maze is either empty or occupied by an obstacle. The Robot can move between neighboring cells to the left (symbol `L`), right (`R`), up (`U`) or down (`D`). The Robot can move to a cell only if it is empty. Initially, the Robot is in an empty cell.

Your task is to find the lexicographically minimal Robot cycle of length exactly `k`, which begins and ends in the cell where the Robot was initially. The Robot is allowed to visit any cell many times (including the starting cell).

The Robot's path is given as a string consisting of symbols `L`, `R`, `U` and `D`. For example, if the Robot goes down, then left, then right, and up, its path is written as `DLRU`.

In this task you don't need to minimize the length of the path. Find the lexicographically smallest (dictionary order) string that satisfies the requirements above.

## Input
The first line contains three integers `n`, `m` and `k` (`1 <= n, m <= 1000`, `1 <= k <= 10^6`) — the size of the maze and the length of the cycle.

Each of the following `n` lines contains `m` symbols — the description of the maze:

- `.` means the cell is empty
- `*` means the cell is occupied by an obstacle
- `X` means the Robot initially is in this cell (and it is empty)

It is guaranteed that the symbol `X` appears in the maze exactly once.

## Output
Print the lexicographically minimum Robot path of length exactly `k`, which starts and ends in the cell where the Robot initially is. If there is no such path, print `IMPOSSIBLE` (without quotes).

## Examples

### Example 1
**Input**
```
2 3 2
.**
X..
```
**Output**
```
RL
```

### Example 2
**Input**
```
5 6 14
..***.
*...X.
..*...
..*.**
....*.
```
**Output**
```
DLDDLLLRRRUURU
```

### Example 3
**Input**
```
3 3 4
***
*X*
***
```
**Output**
```
IMPOSSIBLE
```

## Note
In the first sample, two cyclic paths of length `2` exist — `UD` and `RL`. The second cycle is lexicographically smaller.

In the second sample the Robot should move in the following way: down, left, down, down, left, left, left, right, right, right, up, up, right, up.

In the third sample the Robot can't move to the neighboring cells, because they are occupied by obstacles.


### ideas
1. k长度必须是偶数，且出现一个L的情况下，肯定要出现一个R, 出现一个U，必须有一个D
2. ABCDEFGHIJKLMNOPQRSTUVWXYZ
3. D < L < R < U
4. 如果能够移动到D就移动，
5. 不能移动到D得时候，就移动L
6. 如果不能就R？
7. 但是如何保证能回去呢？记录目前离起点的最短的距离？
8. 剩余的移动次数 >= 最短距离