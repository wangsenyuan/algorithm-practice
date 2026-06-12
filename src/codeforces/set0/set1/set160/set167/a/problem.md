# A. Wizards and Trolleybuses

[Problem link](https://codeforces.com/problemset/problem/167/A)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: stdin

output: stdout

In some country live wizards. They love to ride trolleybuses.

A city in this country has a trolleybus depot with `n` trolleybuses. Every day the
trolleybuses leave the depot one by one and go to the final station. The final
station is at a distance of `d` meters from the depot.

For the `i`-th trolleybus:

- it leaves at time `t_i` seconds;
- its speed never exceeds `v_i` meters per second;
- its acceleration never exceeds `a` meters per second squared.

A trolleybus can decelerate arbitrarily fast (magic). It can also change its
acceleration arbitrarily fast. The maximum acceleration `a` is the same for all
trolleybuses.

Despite the magic, trolleybuses are powered by an electric circuit and **cannot
overtake each other**. If a trolleybus catches up with another one, they travel
together one right after the other until they reach the final station. Drivers
always act to arrive at the final station as quickly as possible.

You, as head of the trolleybuses' fans' club, must determine for each trolleybus
the minimum time by which it can reach the final station. At arrival, speed need
not be zero. When a trolley leaves the depot, its speed is zero. Trolleybuses may
be treated as material points; ignore everything except engine acceleration and
deceleration.

## Input

The first line contains three space-separated integers `n`, `a`, and `d`
(`1 <= n <= 10^5`, `1 <= a, d <= 10^6`) — the number of trolleybuses, their
maximum acceleration, and the distance from the depot to the final station.

Each of the next `n` lines contains two integers `t_i` and `v_i`
(`0 <= t_1 < t_2 < ... < t_{n-1} < t_n <= 10^6`, `1 <= v_i <= 10^6`) — the
departure time and maximum speed of the `i`-th trolleybus.

## Output

For each trolleybus, print one line with the time it arrives at the final
station, in input order.

The answer is accepted if the absolute or relative error does not exceed
`10^-4`.

## Example

### Input

```text
3 10 10000
0 10
5 11
1000 1
1 2 2628
2 9
```

### Output

```text
1000.5000000000
1000.5000000000
11000.0500000000
33.0990195136
```

## Note

In the first sample, the second trolleybus catches up with the first at distance
510.5 meters from the depot. They travel the remaining 9489.5 meters together at
speed 10 meters per second. Both arrive at time 1000.5 seconds. The third
trolleybus does not catch up with them and arrives at time 11000.05 seconds.
