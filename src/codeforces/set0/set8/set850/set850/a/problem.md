# Problem Description

You are given a set of $n$ points in 5-dimensional space. The points are labeled from $1$ to $n$. No two points coincide.

We will call a point $a$ **bad** if there are different points $b$ and $c$, not equal to $a$, from the given set such that the angle between vectors $\vec{ab}$ and $\vec{ac}$ is acute (i.e. strictly less than $\frac{\pi}{2}$). Otherwise, the point is called **good**.

The angle between vectors $\vec{ab}$ and $\vec{ac}$ in 5-dimensional space is defined as $\arccos\left(\frac{\vec{ab} \cdot \vec{ac}}{|\vec{ab}| \cdot |\vec{ac}|}\right)$, where $\vec{ab} \cdot \vec{ac}$ is the scalar product and $|\vec{ab}|$ is the length of $\vec{ab}$.

Given the list of points, print the indices of the good points in ascending order.

## Input

The first line of input contains a single integer $n$ ($1 \leq n \leq 10^3$) — the number of points.

The next $n$ lines of input contain five integers $a_i, b_i, c_i, d_i, e_i$ ($|a_i|, |b_i|, |c_i|, |d_i|, |e_i| \leq 10^3$) — the coordinates of the $i$-th point. All points are distinct.

## Output

First, print a single integer $k$ — the number of good points.

Then, print $k$ integers, each on their own line — the indices of the good points in ascending order.

## Examples

### Example 1

**Input:**

```text
6
0 0 0 0 0
1 0 0 0 0
0 1 0 0 0
0 0 1 0 0
0 0 0 1 0
0 0 0 0 1
```

**Output:**

```text
1
1
```

### Example 2

**Input:**

```text
3
0 0 1 2 0
0 0 9 2 0
0 0 5 9 0
```

**Output:**

```text
0
```

## Note

In the first sample, the first point forms exactly a $\frac{\pi}{2}$ angle with all other pairs of points, so it is good.

In the second sample, along the $cd$ plane, we can see the points look as follows:

### ideas
1. 给定a，遍历其他的节点b，形成向量 x = (a -> b)
2. cos(0) = 1, cos(90) = 0, cos(w) < 0, when w > 90
3. when (x * y) / (sqrt(x) * sqrt(y)) < 0, 
4. 那就是看b, c 相对于a 的向量x, y 的符号