# Problem: New Year Snow Blower

Peter got a new snow blower as a New Year present. Of course, Peter decided to try it immediately. After reading the instructions he realized that it does not work like regular snow blowing machines. In order to make it work, you need to tie it to some point that it does not cover, and then switch it on. As a result it will go along a circle around this point and will remove all the snow from its path.

Formally, we assume that Peter's machine is a polygon on a plane. Then, after the machine is switched on, it will make a circle around the point to which Peter tied it (this point lies strictly outside the polygon). That is, each of the points lying within or on the border of the polygon will move along the circular trajectory, with the center of the circle at the point to which Peter tied his machine.

Peter decided to tie his car to point P and now he is wondering what is the area of the region that will be cleared from snow. Help him.

## Input

The first line of the input contains three integers — the number of vertices of the polygon $n$ $(3 \le n \le 100000)$, and coordinates of point $P$ $(x, y)$.

Each of the next $n$ lines contains two integers — coordinates of the vertices of the polygon in the clockwise or counterclockwise order. It is guaranteed that no three consecutive vertices lie on a common straight line.

All the numbers in the input are integers that do not exceed $1000000$ in their absolute value.

## Output

Print a single real value number — the area of the region that will be cleared. Your answer will be considered correct if its absolute or relative error does not exceed $10^{-6}$.

Namely: let's assume that your answer is $a$, and the answer of the jury is $b$. The checker program will consider your answer correct, if $\frac{|a - b|}{\max(1, |b|)} \le 10^{-6}$.

## Examples

### Example 1

**Input:**
```
3 0 0
0 1
-1 2
1 2
```

**Output:**
```
12.566370614359172464
```

### Example 2

**Input:**
```
4 1 -1
0 0
1 2
2 0
1 1
```

**Output:**
```
21.991148575128551812
```

## Note

In the first sample, snow will be removed from that area.


### ideas
1. 很有意思的题目， 虽然我感觉不会这个题目
2. 要找到两个半径r1, r2， r1是离P最近的点？r2是离它最远的点？
3. 如果只是这两个点，就太简单了～
4. r2没有问题，问题出在r1, r1应该是一条连线到P的最短距离
5. 所以要找到convex，然后找邻近点之间连线到P的距离