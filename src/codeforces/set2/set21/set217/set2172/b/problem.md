There is a long straight road of length $\ell$ meters, where position $p$ denotes the point on the road that is $p$ meters away from the starting point. Along this road, there are $n$ buses moving in the positive direction, each traveling at the same constant speed of $x$ meters per minute. The $i$-th bus is currently at position $s_i$ and continues moving until it reaches its designated destination at position $t_i$. Once a bus reaches its destination, it ceases operation and all passengers must disembark.

There are also $m$ people who wish to reach the end of the road (position $\ell$). The current position of the $i$-th person is $p_i$, and each person can walk at a speed of at most $y$ meters per minute. If a person is at the same position as a bus, they may hop on the bus instantly. While riding a bus, they may hop off at any moment. The time required to board or leave a bus is considered negligible. Buses always move at a constant speed $x$ and never wait for passengers.

Your task is to determine the minimum possible time for each person to reach the end of the road.

*Figure 1: An illustration for sample input 1.*

## Input

The first line contains five integers $n$, $m$, $\ell$, $x$ and $y$, representing the number of buses, the number of people, the length of the road, the speed of the buses, and the walking speed of the people, respectively.

The $i$-th of the following $n$ lines contains two integers $s_i$ and $t_i$, representing the starting position and the destination position of the $i$-th bus.

The $i$-th of the following $m$ lines contains one integer $p_i$, representing the current position of the $i$-th person.

**Constraints:**

- $1 \le n \le 2 \times 10^5$
- $1 \le m \le 2 \times 10^5$
- $1 \le \ell \le 10^9$
- $1 \le y < x \le 10^6$
- $0 \le s_i < t_i \le \ell$
- $0 \le p_i \le \ell$

## Output

Print $m$ lines. The $i$-th line contains a number which is the minimum time (in minutes) for the $i$-th person to reach the end of the road.

Your answer will be accepted if the absolute or relative error does not exceed $10^{-6}$. Formally, let your answer be $a$, and the jury's answer be $b$. Your answer is considered correct if $\frac{|a - b|}{\max(1, |b|)} \le 10^{-6}$.

## Examples

### Input

```
3 3 10 4 1
0 5
2 4
7 9
3
8
5
```

### Output

```
6.25
1.5
5
```

### Input

```
1 3 100 100 1
1 2
0
1
2
```

### Output

```
100
98.01
98
```

## Note

**Explanation of Sample 1:** A person initially at position $p = 3$ can reach the end of the road in $6.25$ minutes as follows:

1. Wait for Bus 1 to arrive.
2. Hop on the bus and ride it until it reaches its destination at position $t_1 = 5$.
3. Get off the bus and walk the remaining distance to position $\ell = 10$.

As shown in Figure 1, the total time spent is $6.25$ minutes, which is the minimum possible.


### ideas
1. 假设当前用户在为止p[i], 他正在以速度y往前运动，后面有一辆车在以速度x追他，
2. 那么z = (x - y) 是速度差， 也就是说 t = (p[i] - s[?]) / z 是他们处在同一个位置的时间（不考虑结束位置)
3. 那么这个t可以计算出来，此时q[i] = p[i] + t * y
4. 如果 q[i] > e[j], 那么i乘坐不到这辆车
5. 要找到i能乘坐到的最远的车，然后从t[j]开始步行计算
6. s[j] <= p[i], 且 (p[i] - s[j]) / (x - y) * y + p[i] <= t[j]
7. 如果 (p[i] - s[j]) / z * y + p[i] > t[j] 
8. (p[i] - s[j]) * y > (t[j] - p[i]) * z
9. p[i] * y + p[i] * z > s[j] * y + t[j] * z
10. p[i] * x > s[j] * y + t[j] * z