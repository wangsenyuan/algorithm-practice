There is an $N \times N$ grid. The cell at the $i$-th row from the top and $j$-th column from the left is called $(i, j)$.

The grid is described by $N$ strings $S_1, S_2, \ldots, S_N$. If the $j$-th character of $S_i$ is `.`, $(i, j)$ is an empty cell; if it is `#`, $(i, j)$ is a wall cell.

Initially, Takahashi-kun is at empty cell $(N, C)$, and repeats the following movement $N-1$ times:

- If he is currently at $(r, c)$, he specifies one of $(r-1, c-1)$, $(r-1, c)$, $(r-1, c+1)$ as the destination. Here, he cannot specify a cell that does not exist in the grid as the destination.
- If the destination $(a, b)$ is a wall cell, the following occurs:
  - If $(i, b)$ is currently an empty cell for all integers satisfying $a < i \le N$, he destroys the wall at $(a, b)$ and moves there. That is, $(a, b)$ becomes an empty cell and he moves to $(a, b)$.
  - Otherwise, he fails to move. In this case, he immediately ends the repetition of movements, even if he has not made $N-1$ movements.
- If the destination $(a, b)$ is an empty cell, he moves to $(a, b)$.

Output a string $R$ of length $N$ satisfying the following conditions:

- If he can reach $(1, i)$ without failing during the movements, the $i$-th character of $R$ is `1`.
- Otherwise, the $i$-th character of $R$ is `0`.

You are given $T$ test cases; solve each of them.

## Constraints

All input values are integers.

- $1 \le T \le 50000$
- $2 \le N \le 3000$
- $1 \le C \le N$
- $S_i$ is a string of length $N$ consisting of `.` and `#`.
- The $C$-th character of $S_N$ is `.`.
- For each input, the sum of $N^2$ does not exceed $9 \times 10^6$.

## Input

The input is given from Standard Input in the following format:

```
T
case_1
case_2
⋮
case_T
```

Each test case is given in the following format:

```
N C
S_1
S_2
⋮
S_N
```

## Output

Print $T$ lines.

The $i$-th line should contain the answer for the $i$-th test case.

## Sample Input 1

```
5
5 3
.###.
..#..
#.#.#
#...#
##..#
2 2
##
..
4 1
####
####
####
.###
3 3
...
...
...
10 3
##.##.##.#
.####..#..
...#.#..#.
.#.#.#.#..
...####...
#.#.##....
.##...#...
#.##.....#
#....###.#
.#..#.#...
```

## Sample Output 1

```
10111
11
1000
111
0011010010
```

This input contains five test cases.

For the first test case, for example, he can reach $(1, 3)$ without failing during the movements as follows:

- Initially, he is at $(5, 3)$.
- He moves to empty cell $(4, 2)$.
- $(3, 3)$ is a wall cell, but since $(4, 3)$, $(5, 3)$ are currently both empty cells, he destroys the wall at $(3, 3)$ and moves to $(3, 3)$.
- $(2, 3)$ is a wall cell, but since $(3, 3)$, $(4, 3)$, $(5, 3)$ are currently all empty cells, he destroys the wall at $(2, 3)$ and moves to $(2, 3)$.
- $(1, 3)$ is a wall cell, but since $(2, 3)$, $(3, 3)$, $(4, 3)$, $(5, 3)$ are currently all empty cells, he destroys the wall at $(1, 3)$ and moves to $(1, 3)$.

He can reach $(1, 1)$, $(1, 3)$, $(1, 4)$, $(1, 5)$ without failing during the movements, so we print `10111`.

There is an 
N×N grid. The cell at the 
i-th row from the top and 
j-th column from the left is called 
(i,j).
The grid is described by 
N strings 
S 
1
​
 ,S 
2
​
 ,…,S 
N
​
 . If the 
j-th character of 
S 
i
​
  is ., 
(i,j) is an empty cell; if it is #, 
(i,j) is a wall cell.

Initially, Takahashi-kun is at empty cell 
(N,C), and repeats the following movement 
N−1 times:

If he is currently at 
(r,c), he specifies one of 
(r−1,c−1),(r−1,c),(r−1,c+1) as the destination. Here, he cannot specify a cell that does not exist in the grid as the destination.
If the destination 
(a,b) is a wall cell, the following occurs:
If 
(i,b) is currently an empty cell for all integers satisfying 
a<i≤N, he destroys the wall at 
(a,b) and moves there. That is, 
(a,b) becomes an empty cell and he moves to 
(a,b).
Otherwise, he fails to move. In this case, he immediately ends the repetition of movements, even if he has not made 
N−1 movements.
If the destination 
(a,b) is an empty cell, he moves to 
(a,b).
Output a string 
R of length 
N satisfying the following conditions:

If he can reach 
(1,i) without failing during the movements, the 
i-th character of 
R is 1.
Otherwise, the 
i-th character of 
R is 0.
You are given 
T test cases; solve each of them.

Constraints
T,N,C are integers.
1≤T≤50000
2≤N≤3000
1≤C≤N
S 
i
​
  is a string of length 
N consisting of . and #.
The 
C-th character of 
S 
N
​
  is ..
For each input, the sum of 
N 
2
  does not exceed 
9×10 
6
 .
Input
The input is given from Standard Input in the following format:

T
case 
1
​
 
case 
2
​
 
⋮
case 
T
​
 
Each test case is given in the following format:

N 
C
S 
1
​
 
S 
2
​
 
⋮
S 
N
​
 
Output
Print 
T lines.

The 
i-th line should contain the answer for the 
i-th test case.

Sample Input 1
Copy
5
5 3
.###.
..#..
#.#.#
#...#
##..#
2 2
##
..
4 1
####
####
####
.###
3 3
...
...
...
10 3
##.##.##.#
.####..#..
...#.#..#.
.#.#.#.#..
...####...
#.#.##....
.##...#...
#.##.....#
#....###.#
.#..#.#...
Sample Output 1
Copy
10111
11
1000
111
0011010010
This input contains five test cases.

For the first test case, for example, he can reach 
(1,3) without failing during the movements as follows:

Initially, he is at 
(5,3).
He moves to empty cell 
(4,2).
(3,3) is a wall cell, but since 
(4,3),(5,3) are currently both empty cells, he destroys the wall at 
(3,3) and moves to 
(3,3).
(2,3) is a wall cell, but since 
(3,3),(4,3),(5,3) are currently all empty cells, he destroys the wall at 
(2,3) and moves to 
(2,3).
(1,3) is a wall cell, but since 
(2,3),(3,3),(4,3),(5,3) are currently all empty cells, he destroys the wall at 
(1,3) and moves to 
(1,3).
He can reach 
(1,1),(1,3),(1,4),(1,5) without failing during the movements, so print 10111.


### ideas
1. 在移动的过程中，如果目标是一个#，那么如果它下面全部是., 那么就可以摧毁它，并到达那里
2. 如果是一个empty, 那么就可以直接到达那里
3. 如果要到达(1, i), 那么有两种策略，一种就是顺着空的格子移动
4. 还有一种是，到达它的(r, i) (必须尽快到达) 然后一路往上
5. 如果这两个策略都不行，那么就肯定不行～
6. 