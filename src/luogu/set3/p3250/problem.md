# P3250 [HNOI2016] 网络

[Problem link](https://www.luogu.com.cn/problem/P3250)

time limit per test: 2.00s

memory limit per test: 512.00MB

A network is an unrooted tree. Each node is a server. When two servers exchange data, the data
passes through every server on the unique path between them, including the two endpoints.

Each request has an importance. If a server on that path fails, the request is affected. A server
may send a request to itself.

There are three kinds of events, given in time order. Event `i` happens at time `i`:

- `0 a b v` — a request of importance `v` appears between servers `a` and `b`
- `1 t` — the request that appeared at time `t` ends
- `2 x` — server `x` fails; among requests that are still active and do not pass through `x`,
  what is the maximum importance? If there is no such request, the answer is `-1`

## Input

The first line contains two positive integers `n` and `m`, the number of servers and the number of
events. Servers are numbered `1` through `n`.

Each of the next `n - 1` lines contains two positive integers `u` and `v`, a tree edge.

Each of the next `m` lines describes one event, as above.

## Output

For each type-`2` event, print one integer on its own line.

## Sample

### Input

```text
13 23
1 2
1 3
2 4
2 5
3 6
3 7
4 8
4 9
6 10
6 11
7 12
7 13
2 1
0 8 13 3
0 9 12 5
2 9
2 8
2 2
0 10 12 1
2 2
1 3
2 7
2 1
0 9 5 6
2 4
2 5
1 7
0 9 12 4
0 10 5 7
2 1
2 4
2 12
1 2
2 5
2 3
```

### Output

```text
-1
3
5
-1
1
-1
1
1
3
6
7
7
4
6
```

### Note

Write a request as `(a, b; t, v)` for a request of importance `v` between `a` and `b` that starts
at time `t`.

- Time 1: no requests, answer `-1`
- Time 6: requests `(8, 13; 2, 3)` and `(9, 12; 3, 5)` both go through server 2, answer `-1`
- Time 8: only `(10, 12; 7, 1)` avoids server 2, answer `1`
- Time 23: requests `(9, 5; 12, 6)`, `(9, 12; 16, 4)`, and `(10, 5; 17, 7)`; only `(9, 5; 12, 6)`
  avoids server 3, answer `6`

## Constraints

- `2 <= n <= 10^5`
- `1 <= m <= 2 * 10^5`
- All other input values are at most `10^9`

## Status

I/O and the official sample are in place. `solve` is left as a TODO.
