# Problem Description

Rumors say that one of Kamal-ol-molk's paintings has been altered. A rectangular brush has been moved right and down on the painting.

Consider the painting as a n × m rectangular grid. At the beginning an x × y rectangular brush is placed somewhere in the frame, with edges parallel to the frame, (1 ≤ x ≤ n, 1 ≤ y ≤ m). Then the brush is moved several times. Each time the brush is moved one unit right or down. The brush has been strictly inside the frame during the painting. The brush alters every cell it has covered at some moment.

You have found one of the old Kamal-ol-molk's paintings. You want to know if it's possible that it has been altered in described manner. If yes, you also want to know minimum possible area of the brush.

## Input

The first line of input contains two integers n and m, (1 ≤ n, m ≤ 1000), denoting the height and width of the painting.

The next n lines contain the painting. Each line has m characters. Character 'X' denotes an altered cell, otherwise it's showed by '.'. There will be at least one altered cell in the painting.

## Output

Print the minimum area of the brush in a line, if the painting is possibly altered, otherwise print -1.

## Examples

### Example 1
**Input:**
```
4 4
XX..
XX..
XXXX
XXXX
```

**Output:**
```
4
```

### Example 2
**Input:**
```
4 4
....
.XXX
.XXX
....
```

**Output:**
```
2
```

### Example 3
**Input:**
```
4 5
XXXX.
XXXX.
.XX..
.XX..
```

**Output:**
```
-1
```

## ideas
1. 是否能用一个 x * y 的刷子，刷出图中X组成的区域？
2. 有点懵～ 已知x * y 的情况下，应该是可以搞的，如果同一行的右边存在，就只能往右，然后往下
3. 如果同时存在往下、往右的情况，就无解～
4. 但是关键是 x 和 y的组合就是 n * n ， 这个判断，又是max(n, m)
5. 考虑第一个区域（完全的长方形), w * h, 要么 x = h, 要么 y = w ?
6. 找到所有这样的长方形区域，都有对应的要求，那么x或者y必须同时满足这些区域的条件
7. 