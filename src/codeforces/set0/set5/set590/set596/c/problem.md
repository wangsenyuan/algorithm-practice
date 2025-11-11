## Problem Statement

Wilbur is playing with a set of $n$ points on the coordinate plane. All points have non-negative integer coordinates. Moreover, if some point $(x, y)$ belongs to the set, then all points $(x', y')$, such that $0 \le x' \le x$ and $0 \le y' \le y$ also belong to this set.

Now Wilbur wants to number the points in the set he has, that is assign them distinct integer numbers from $1$ to $n$. In order to make the numbering aesthetically pleasing, Wilbur imposes the condition that if some point $(x, y)$ gets number $i$, then all $(x',y')$ from the set, such that $x' \ge x$ and $y' \ge y$ must be assigned a number not less than $i$. 

For example, for a set of four points $(0, 0), (0, 1), (1, 0)$ and $(1, 1)$, there are two aesthetically pleasing numberings. One is $1, 2, 3, 4$ and another one is $1, 3, 2, 4$.

Wilbur's friend comes along and challenges Wilbur. For any point he defines its special value as $s(x, y) = y - x$. Now he gives Wilbur some $w_1, w_2, \ldots, w_n$, and asks him to find an aesthetically pleasing numbering of the points in the set, such that the point that gets number $i$ has its special value equal to $w_i$, that is $s(x_i, y_i) = y_i - x_i = w_i$.

Now Wilbur asks you to help him with this challenge.

## Input

The first line of the input consists of a single integer $n$ $(1 \le n \le 100\,000)$ — the number of points in the set Wilbur is playing with.

Next follow $n$ lines with points descriptions. Each line contains two integers $x$ and $y$ $(0 \le x, y \le 100\,000)$, that give one point in Wilbur's set. It's guaranteed that all points are distinct. Also, it is guaranteed that if some point $(x, y)$ is present in the input, then all points $(x', y')$, such that $0 \le x' \le x$ and $0 \le y' \le y$, are also present in the input.

The last line of the input contains $n$ integers. The $i$-th of them is $w_i$ $(-100\,000 \le w_i \le 100\,000)$ — the required special value of the point that gets number $i$ in any aesthetically pleasing numbering.

## Output

If there exists an aesthetically pleasant numbering of points in the set, such that $s(x_i, y_i) = y_i - x_i = w_i$, then print "YES" on the first line of the output. Otherwise, print "NO".

If a solution exists, proceed output with $n$ lines. On the $i$-th of these lines print the point of the set that gets number $i$. If there are multiple solutions, print any of them.

## Examples

### Example 1
**Input:**
```
5
2 0
0 0
1 0
1 1
0 1
0 -1 -2 1 0
```
**Output:**
```
YES
0 0
1 0
2 0
0 1
1 1
```

### Example 2
**Input:**
```
3
1 0
0 0
2 0
0 1 2
```
**Output:**
```
NO
```

## Note

In the first sample, point $(2, 0)$ gets number $3$, point $(0, 0)$ gets number one, point $(1, 0)$ gets number $2$, point $(1, 1)$ gets number $5$ and point $(0, 1)$ gets number $4$. One can easily check that this numbering is aesthetically pleasing and $y_i - x_i = w_i$.

In the second sample, the special values of the points in the set are $0, -1$, and $-2$ while the sequence that the friend gives to Wilbur is $0, 1, 2$. Therefore, the answer does not exist.


### ideas
1. 就是要给所有的点(x, y)进行编号，满足两个条件
2.   a) 如果这个点的编号是i, 那么所有右上的点的编号 > i
3.   b) y - x = wi
4. 这些点可以组成一个图，每个点，都指向它垂直上方的第一个点，水平右方的第一个点
5. 这样子就可以得到一个图，然后进行bfs，对于当前层的，看能否匹配