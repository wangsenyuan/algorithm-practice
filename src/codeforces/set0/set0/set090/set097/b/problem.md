# Problem 97B - Good Points

## Problem Statement

A set of points on a plane is called **good**, if for any two points at least one of the three conditions is true:

1. Those two points lie on same horizontal line;
2. Those two points lie on same vertical line;
3. The rectangle, with corners in these two points, contains inside or on its borders at least one point of the set, 
   other than these two. We mean here a rectangle with sides parallel to coordinates' axes, the so-called bounding 
   box of the two points.

You are given a set consisting of n points on a plane. Find any good superset of the given set whose size would not 
exceed 2·10⁵ points.

## Input

The first line contains an integer `n` (1 ≤ n ≤ 10⁴) — the number of points in the initial set. 

Next `n` lines describe the set's points. Each line contains two integers `xi` and `yi` (-10⁹ ≤ xi, yi ≤ 10⁹) — a 
corresponding point's coordinates. 

It is guaranteed that all the points are different.

## Output

Print on the first line the number of points `m` (n ≤ m ≤ 2·10⁵) in a good superset, print on next `m` lines the 
points. 

The absolute value of the points' coordinates should not exceed 10⁹. Note that you should not minimize `m`, it is 
enough to find any good superset of the given set, whose size does not exceed 2·10⁵.

All points in the superset should have integer coordinates.

## Example

### Input
```
2
1 1
2 2
```

### Output
```
3
1 1
2 2
1 2
```


### ideas
1. 如果n个点在一条直线上 => 这n个点
2. 那么必然存在一个box, 包含了所有的点（有可能corner不存在，那么需要增加最多4个corner)
3. 然后所有不在这个box上面的点，那些内部的点(x, y)
4. 对于点(x1, y1) 假设在右上角离他最近的点(x2, y2)
5. 它们组成了一个矩形，且它们中间没有其他的点，那么就需要添加一个点(x2, y1) 或者 (x1, y2) 
6. 但是怎么保证新加的点符合条件呢？
7. 考虑 y = x 这条线上的点，
8. 一开始，任意相邻的两个点，不符合条件（只要是不相邻的都是good的）
9. 大概有个想法了，就是按层处理，每层只需要和上一层进行匹配。（从下到上，从左到右）
10. 最后再增加corner就可以了
11. 