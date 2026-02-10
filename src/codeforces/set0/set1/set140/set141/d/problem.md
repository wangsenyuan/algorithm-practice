Vasya participates in a ski race along the X axis. The start is at point $0$, and the finish is at $L$, that is, at a distance $L$ meters from the start in the positive direction of the axis. Vasya has been training so hard that he can run one meter in exactly one second.

Besides, there are $n$ take-off ramps on the track. Each ramp is characterized by four numbers:

- $x_i$ — the ramp's coordinate
- $d_i$ — from how many meters Vasya will land if he goes down this ramp
- $t_i$ — the flight time in seconds
- $p_i$ — for how many meters Vasya should gather speed to get ready and fly off the ramp. As Vasya gathers speed, he should ski on the snow (that is, he should not be flying), but his speed still equals one meter per second.

Vasya is allowed to move in any direction on the X axis, but he is prohibited to cross the start line (that is, go to the negative semiaxis). Vasya himself chooses which take-off ramps he will use and in what order; he is not obliged to take off from all the ramps he encounters. Specifically, Vasya can skip a ramp. It is guaranteed that $x_i + d_i \le L$, that is, Vasya cannot cross the finish line in flight.

Vasya can jump from the ramp only in the positive direction of the X axis. More formally, when using the $i$-th ramp, Vasya starts gathering speed at point $x_i - p_i$, jumps at point $x_i$, and lands at point $x_i + d_i$. He cannot use the ramp in the opposite direction.

Your task is to find the minimum time that Vasya will spend to cover the distance.

## Input

The first line contains two integers $n$ and $L$ ($0 \le n \le 10^5$, $1 \le L \le 10^9$). Then $n$ lines contain the descriptions of the ramps, each description on a single line. Each description is a group of four non-negative integers $x_i$, $d_i$, $t_i$, $p_i$ ($0 \le x_i \le L$, $1 \le d_i, t_i, p_i \le 10^9$, $x_i + d_i \le L$).

## Output

Print in the first line the minimum time in seconds Vasya needs to complete the track. Print in the second line $k$ — the number of take-off ramps that Vasya needs to use. Print on the third line $k$ numbers: the indices of the take-off ramps Vasya used, in the order in which he used them. Print each number exactly once; separate the numbers with a space. The ramps are numbered starting from $1$ in the order in which they are given in the input.

## Examples

### Input

```
2 20
5 10 5 5
4 16 1 7
```

### Output

```
15
1
1
```

### Input

```
2 20
9 8 12 6
15 5 1 1
```

### Output

```
16
1
2
```

## Note

In the first sample, Vasya cannot use ramp 2, because then he will need to gather speed starting from point $-3$, which is not permitted by the statement. The optimal option is using ramp 1. The resulting time is: moving to the point of gathering speed + gathering speed until reaching the takeoff ramp + flight time + moving to the finish line $= 0 + 5 + 5 + 5 = 15$.

In the second sample, using ramp 1 is not optimal for Vasya as $t_1 > d_1$. The optimal option is using ramp 2. The resulting time is: moving to the point of gathering speed + gathering speed until reaching the takeoff ramp + flight time + moving to the finish line $= 14 + 1 + 1 + 0 = 16$.


### ideas
1. L是最差值，
2. 假设在ramp[i]上起跳（前面不管）它落地的位置 y[i] = x[i] + d[i], 花费的时间是t[i]
3. 假设后面有个ramp[j], 如果y[i] <= x[j] - p[j], 那么i可以直接贡献j， t[i] + (x[j] - y[i])
4. 如果y[i] > x[j] - p[j], 那么必须留出p[j]的距离，往回运动 y[i] - (x[j] - p[j])的距离，然后起跳
5. 贡献 = t[i] + y[i] - (x[j] - p[j]) + p[j] = t[i] + y[i] + 2 * p[j] - x[j]
6. 那么反过来，对于j来说，就是找到那些y[i] <= x[j] - p[j]的最小值(t[i] - y[i]) 的最小值
7. 或者是 y[i] > x[j] - p[j]部分, t[i] + y[i] 的最小值 + 2 * p[j] - x[j]
8. 对于两个i, j, 有没有可能存在 x[i] < x[j], 但是先跳x[j], 再跳x[i]更快的情况呢？
9. 不可能。所以按照x升序处理。需要两个segment tree
10.  