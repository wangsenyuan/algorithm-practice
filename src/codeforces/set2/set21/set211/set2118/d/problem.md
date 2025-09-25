# Problem D: Traffic Light Strip

## Problem Statement

You are given a strip of length $10^{15}$ and a constant $k$. There are exactly $n$ cells that contain a traffic light; each has a position $p_i$ and an initial delay $d_i$ for which $d_i < k$. The $i$-th traffic light works the following way:

- it shows red at the $l \cdot k + d_i$-th second, where $l$ is an integer
- it shows green otherwise

At second 0, you are initially positioned at some cell on the strip, facing the positive direction. At each second, you perform the following actions in order:

1. If the current cell contains a red traffic light, you turn around
2. Move one cell in the direction you are currently facing

You are given $q$ different starting positions. For each one, determine whether you will eventually leave the strip within $10^{100}$ seconds.

## Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \leq t \leq 2 \cdot 10^5$). The description of the test cases follows.

The first line of each test case contains two integers $n$, $k$ ($1 \leq n \leq 2 \cdot 10^5$ and $1 \leq k \leq 10^{15}$) — the number of traffic lights and the length of the period.

The second line of each test case contains $n$ integers $p_1, p_2, \ldots, p_n$ ($1 \leq p_1 < p_2 < \ldots < p_n \leq 10^{15}$) — the positions of the traffic lights.

The third line of each test case contains $n$ integers $d_1, d_2, \ldots, d_n$ ($0 \leq d_i < k$) — the delays of the traffic lights.

The fourth line of each test case contains one integer $q$ ($1 \leq q \leq 2 \cdot 10^5$) — the number of queries.

The fifth line of each test case contains $q$ integers $a_1, a_2, \ldots, a_q$ ($1 \leq a_i \leq 10^{15}$) — the starting positions.

It is guaranteed that the sum of $n$ and $q$ over all test cases does not exceed $2 \cdot 10^5$.

## Output

For each test case, output $q$ lines. Each line should contain "YES" if you will eventually leave the strip and "NO" otherwise. You can output the answer in any case (upper or lower). For example, the strings "yEs", "yes", "Yes", and "YES" will be recognized as positive responses.

## Example

### Input
```
4
2 2
1 4
1 0
3
1 2 3
9 4
1 2 3 4 5 6 7 8 9
3 2 1 0 1 3 3 1 1
5
2 5 6 7 8
4 2
1 2 3 4
0 0 0 0
4
1 2 3 4
3 4
1 2 3
3 1 1
3
1 2 3
```

### Output
```
YES
NO
YES
YES
YES
YES
NO
NO
YES
YES
NO
NO
YES
NO
YES
```