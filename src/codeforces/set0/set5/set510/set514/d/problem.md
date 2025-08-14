# Problem D: Droid Army

## Problem Description

An army of $n$ droids is lined up in one row. Each droid is described by $m$ integers $a_1, a_2, \ldots, a_m$, where $a_i$ is the number of details of the $i$-th type in this droid's mechanism. R2-D2 wants to destroy the sequence of consecutive droids of maximum length. He has $m$ weapons, the $i$-th weapon can affect all the droids in the army by destroying one detail of the $i$-th type (if the droid doesn't have details of this type, nothing happens to it).

A droid is considered to be destroyed when all of its details are destroyed. R2-D2 can make at most $k$ shots. How many shots from the weapon of what type should R2-D2 make to destroy the sequence of consecutive droids of maximum length?

## Input

The first line contains three integers $n, m, k$ ($1 \leq n \leq 10^5$, $1 \leq m \leq 5$, $0 \leq k \leq 10^9$) — the number of droids, the number of detail types and the number of available shots, respectively.

Next $n$ lines follow describing the droids. Each line contains $m$ integers $a_1, a_2, \ldots, a_m$ ($0 \leq a_i \leq 10^8$), where $a_i$ is the number of details of the $i$-th type for the respective robot.

## Output

Print $m$ space-separated integers, where the $i$-th number is the number of shots from the weapon of the $i$-th type that the robot should make to destroy the subsequence of consecutive droids of the maximum length.

If there are multiple optimal solutions, print any of them.

**Note:** It is not necessary to make exactly $k$ shots, the number of shots can be less.

## Examples

### Example 1

**Input:**
```
5 2 4
4 0
1 2
2 1
0 2
1 3
```

**Output:**
```
2 2
```

**Explanation:** In the first test the second, third and fourth droids will be destroyed.

### Example 2

**Input:**
```
3 2 4
1 2
1 3
2 2
```

**Output:**
```
1 3
```

**Explanation:** In the second test the first and second droids will be destroyed.


## ideas
1. 假设可以消灭l...r区间，且x[1], x[2]... x[5]表示每种类型使用了x[i]次shot
2. 那么要满足 a[l][1],,a[j][1]... a[r][1] <= x[1]
3. ... 
4. 且 x[1] + x[2] + .. x[5] <= k
5. 对于给定的l，可以2分r
6. 似乎，用双指针也可以