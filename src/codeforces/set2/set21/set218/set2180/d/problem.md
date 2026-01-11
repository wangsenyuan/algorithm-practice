# Problem D

You are given $n$ different integer points $x_1 < x_2 < \ldots < x_n$ on the X-axis of the plane. For each $1 \leq i \leq n$, you have to pick a real value $r_i > 0$ and draw a disk with radius $r_i$ and center $x_i$ such that no two disks overlap.

You want to choose values $r_i$ in such a way that the number of tangent pairs of disks is maximized, and you have to find this maximum value.

We say that two disks overlap if the area of their intersection is positive. In particular, two outer tangent disks do not overlap.

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \leq t \leq 10^4$). The description of the test cases follows.

The first line of each test case contains a single integer $n$ ($1 \leq n \leq 2 \cdot 10^6$) — the number of points in that test case.

The second line of each test case contains $n$ integers $x_1, x_2, \ldots x_n$ ($0 \leq x_1 < x_2 < \ldots < x_n \leq 10^9$) — the coordinates of the points in that test case.

It's also guaranteed that the sum of $n$ over all test cases is at most $2 \cdot 10^6$.

## Output

For each test case, print a single integer denoting the maximum number of tangent pairs of disks in that test case.

## Example

**Input:**

```text
4
3
1 2 3
4
1 2 4 5
6
1 2 11 12 21 22
7
0 1 2 3 5 8 13
```

**Output:**

```text
2
2
3
6
```

## Note

In the first test case, you can let $r_1 = r_3 = 0.25$ and $r_2 = 0.75$ so the 1-st and 2-nd disks and the 2-nd and 3-rd disks will be pairwise tangent.


### ideas
1. 假设有一段m个点，已经连接在一起了，现在要添加一个点进去
2. 新的点，会被下一个点限制r的上限；如果还是够不到前面点的圆，那么需要增加前面点的半径
3. 但是增加前面的，又会挤压更前面的半径，类似这样的过程
4. 假设上一个点增加了 dr, 那么前面的点，需要减少dr， 再前面的点需要增加dr, ...
5. 所以，需要知道这段的距离的上限，是相隔一个的每个点的上限
6. 奇数位和偶数位要分开考虑