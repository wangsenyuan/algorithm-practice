# Problem B

The city of Fishtopia can be imagined as a grid of 4 rows and an odd number of columns. It has two main villages; the first is located at the top-left cell $(1,1)$, people who stay there love fishing at the Tuna pond at the bottom-right cell $(4,n)$. The second village is located at $(4,1)$ and its people love the Salmon pond at $(1,n)$.

The mayor of Fishtopia wants to place $k$ hotels in the city, each one occupying one cell. To allow people to enter the city from anywhere, hotels should not be placed on the border cells.

A person can move from one cell to another if those cells are not occupied by hotels and share a side.

Can you help the mayor place the hotels in a way such that there are equal number of shortest paths from each village to its preferred pond?

## Input

The first line of input contain two integers, $n$ and $k$ ($3 \leq n \leq 99$, $0 \leq k \leq 2 \times (n-2)$), $n$ is odd, the width of the city, and the number of hotels to be placed, respectively.

## Output

Print "YES", if it is possible to place all the hotels in a way that satisfies the problem statement, otherwise print "NO".

If it is possible, print an extra 4 lines that describe the city, each line should have $n$ characters, each of which is "#" if that cell has a hotel on it, or "." if not.

## Examples

### Example 1
**Input:**
```
7 2
```

**Output:**
```
YES
.......
.#.....
.#.....
.......
```

### Example 2
**Input:**
```
5 3
```

**Output:**
```
YES
.....
.###.
.....
.....
```

### ideas
1. 还有点难么～
2. 考虑只有一行，只在第二行，如果它右边有w个空位置，那么，对于第二个村子，
3. 计数 = 到达第一个空列的方案数 = n - w(在任何一个列处选择走到第3行，一旦到达，只能一直往前)
4. n - w + n - w + 1, ..+ n
5. 但是对于第一个存在，如果它一直往右边（选择了第一行）C(4 + w - 1, w - 1)种选择
6. 到达(1, n - w)时，往下或者往右移动
7. 如果 w != 1, 似乎就没法保证是相等的
8. 如果 k 是偶数，就比较简单， 第一行和第二行相同就可以了
9. 如果k是奇数, 如果 k >= 5, 似乎也有办法
10. 下面留一个中间的空位就好了
11. 只有当 k = 1,3的时候，
12. k = 1 的时候， 假设就放在位置(2, 2)处
13. 对于第一个存在来说来说 = C(n - 2 - 1 + 3, 3) 如果先到达位置(1, 3)
14.  + n (如果先到达位置 (3, 2))
15.  对于第二个村子来说, C(n - 2 - 1 + 3, 3) (如果先到达(4, 3))
16.   + 1 (如果到(1, 1))
17.   + 2 * n
18.  好像是无解的