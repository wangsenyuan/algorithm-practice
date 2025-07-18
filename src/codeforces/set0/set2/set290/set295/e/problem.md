# Problem E: Points on Line

## Problem Description

Yaroslav has n points that lie on the Ox axis. The coordinate of the first point is x1, the coordinate of the second point is x2, ..., the coordinate of the n-th point is xn. Now Yaroslav wants to execute m queries, each of them is of one of the two following types:

1. **Move query**: Move the pj-th point from position xpj to position xpj + dj. At that, it is guaranteed that after executing such query all coordinates of the points will be distinct.

2. **Count query**: Count the sum of distances between all pairs of points that lie on the segment [lj, rj] (lj ≤ rj). In other words, you should count the sum of: |xi - xj| for all pairs (i, j) where lj ≤ xi, xj ≤ rj.

Help Yaroslav.

## Input

- The first line contains integer n — the number of points (1 ≤ n ≤ 10^5)
- The second line contains distinct integers x1, x2, ..., xn — the coordinates of points (|xi| ≤ 10^9)
- The third line contains integer m — the number of queries (1 ≤ m ≤ 10^5)
- The next m lines contain the queries:
  - The j-th line first contains integer tj (1 ≤ tj ≤ 2) — the query type
  - If tj = 1, then it is followed by two integers pj and dj (1 ≤ pj ≤ n, |dj| ≤ 1000)
  - If tj = 2, then it is followed by two integers lj and rj (-10^9 ≤ lj ≤ rj ≤ 10^9)

**Note**: It is guaranteed that at any moment all the points have distinct coordinates.

## Output

For each type 2 query print the answer on a single line. Print the answers in the order, in which the queries follow in the input.

**Important**: Please, do not use the %lld specifier to read or write 64-bit integers in C++. It is preferred to use the cin, cout streams or the %I64d specifier.

## Examples

### Input
```
8
36 50 28 -75 40 -60 -95 -48
20
2 -61 29
1 5 -53
1 1 429
1 5 130
2 -101 -71
2 -69 53
1 1 404
1 5 518
2 -101 53
2 50 872
1 1 -207
2 -99 -40
1 7 -389
1 6 -171
1 2 464
1 7 -707
1 1 -730
1 1 560
2 635 644
1 7 -677
```

### Output
```
176
20
406
1046
1638
156
0
```

## ideas
1. 所有的点，加上一开始的，和变化后的，共有（最多） n + m 个点
2. 考虑在segment tree上的点x, y (x < y>)
3. 可以计算了