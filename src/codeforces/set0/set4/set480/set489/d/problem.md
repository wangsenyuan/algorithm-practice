# Problem: Damn Rhombus

## Problem Statement

Tomash keeps wandering off and getting lost while he is walking along the streets of Berland. It's no surprise! In his home town, for any pair of intersections there is exactly one way to walk from one intersection to the other one. The capital of Berland is very different!

Tomash has noticed that even simple cases of ambiguity confuse him. So, when he sees a group of four distinct intersections a, b, c and d, such that there are two paths from a to c — one through b and the other one through d, he calls the group a "damn rhombus". Note that pairs (a, b), (b, c), (a, d), (d, c) should be directly connected by the roads. Schematically, a damn rhombus is shown on the figure below:

```
    a
   / \
  b   d
   \ /
    c
```

Other roads between any of the intersections don't make the rhombus any more appealing to Tomash, so the four intersections remain a "damn rhombus" for him.

Given that the capital of Berland has n intersections and m roads and all roads are unidirectional and are known in advance, find the number of "damn rhombi" in the city.

When rhombi are compared, the order of intersections b and d doesn't matter.

## Input

The first line of the input contains a pair of integers n, m (1 ≤ n ≤ 3000, 0 ≤ m ≤ 30000) — the number of intersections and roads, respectively. Next m lines list the roads, one per line. Each of the roads is given by a pair of integers ai, bi (1 ≤ ai, bi ≤ n; ai ≠ bi) — the number of the intersection it goes out from and the number of the intersection it leads to. Between a pair of intersections there is at most one road in each of the two directions.

It is not guaranteed that you can get from any intersection to any other one.

## Output

Print the required number of "damn rhombi".

## Examples

### Example 1

**Input:**
```
5 4
1 2
2 3
1 4
4 3
```

**Output:**
```
1
```

### Example 2

**Input:**
```
4 12
1 2
1 3
1 4
2 1
2 3
2 4
3 1
3 2
3 4
4 1
4 2
4 3
```

**Output:**
```
12
```


### ideas 
1. 如果a, c中间存在x个点，那么是否增加了 C(x, 2)？
2. 从a出发，找到所有距离第二层的点，
3. 这样子可以在最多m的时间内完成？
4. 似乎不是的，特别是中间层的(b -> c) 或者(d -> c) 会被重复访问（如果有很多节点到bd)
5. 假设可以计算出，有多少对(b, d) 有多少个相同的c
6. 如果也能计算出, (b, d) => a 的个数，两个相乘就可以了
7. 