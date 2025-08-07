# Water Vessel Transfusion Problem

## Problem Description

A system of n vessels with water is given. Several pairs of vessels are connected by tubes with transfusion mechanisms. One may transfer an integer amount of liters of water between two vessels connected by such tube (tube works in both directions). There might be multiple tubes between two vessels. Total number of tubes equals e. Volume of each vessel equals v liters. Of course, the amount of the water in any vessel cannot exceed v liters in the process of transfusions.

Given the initial amounts ai of water in the vessels and the desired amounts bi find a sequence of transfusions that deals with the task. Total number of transfusions must not exceed 2·n².

## Input

First line of the input contains integers n, v, e (1 ≤ n ≤ 300, 1 ≤ v ≤ 10⁹, 0 ≤ e ≤ 50000).

Next two lines contain n integers each: initial ai and the desired amounts bi of water in corresponding vessels (0 ≤ ai, bi ≤ v).

Next e lines describe one tube each in the format x y (1 ≤ x, y ≤ n, x ≠ y) for a tube between vessels number x and y. There might be multiple tubes between two vessels. You may assume that vessels are numbered from 1 to n in some way.

## Output

Print "NO" (without quotes), if such sequence of transfusions does not exist.

Otherwise print any suitable sequence in the following format. On the first line print the total number of transfusions k (k should not exceed 2·n²). In the following k lines print transfusions in the format x y d (transfusion of d liters from the vessel number x to the vessel number y, x and y must be distinct). For all transfusions d must be a non-negative integer.

## Examples

### Example 1

**Input:**
```
2 10 1
1 9
5 5
1 2
```

**Output:**
```
1
2 1 4
```

### Example 2

**Input:**
```
2 10 0
5 2
4 2
```

**Output:**
```
NO
```

### Example 3

**Input:**
```
2 10 0
4 2
4 2
```

**Output:**
```
0
```

### ideas
1. 假设是一棵树，这棵树sum(a[u]) = sum(b[u]) 否则没有答案
2. 从叶子节点开始处理，假设节点u，它的子节点都搞定了
3. 现在处理u。且假设a[u] < b[u], 也就是是它需要新的流量
4. 如果能够从parent获得足够的流量, 就补充给它
5. 但是如果不够，那么就记录下来？这时候，a[p] = 0
6. 但是如果a[u] > b[u], 也就是多出来了，如果能够把增加的部分，给消耗给p, 那么就消耗掉
7. 否则也需要记录下来; 一个节点u上，只可能记录，要么是需要补充的（a[u] = 0)
8. 要么是堵在u的， a[u] = v
9. 