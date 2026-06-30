# C - Flapping Takahashi

[Problem link](https://atcoder.jp/contests/abc434/tasks/abc434_c)

**Contest:** [Sky Inc, Programming Contest 2025 (AtCoder Beginner Contest 434)](https://atcoder.jp/contests/abc434)

time limit: 2 sec

memory limit: 1024 MiB

score: 300 points

Takahashi has decided to fly in the sky with balloons. Takahashi is at altitude `H` at time `0`
(seconds), and will fly for `10^9` seconds from now.

Takahashi can change his flying altitude by at most `1` per second. However, he cannot make his
altitude `0` or less.

In other words, if `F(t)` denotes Takahashi's altitude at time `t`, then `F(t)` satisfies all of
the following:

- `F(0) = H`
- `-1 <= (F(u) - F(t)) / (u - t) <= 1` for `0 <= t < u <= 10^9`
- `F(t) > 0` for `0 <= t <= 10^9`

There are `N` goals regarding altitude. The `i`-th goal is to make the altitude at time `t_i` at
least `l_i` and at most `u_i`.

Is it possible for Takahashi to fly in a way that achieves all goals?

You are given `T` test cases; solve each of them.

## Constraints

- `1 <= T <= 10^5`
- `1 <= N <= 10^5`
- `1 <= H <= 10^9`
- `1 <= t_1 < t_2 < ... < t_N <= 10^9`
- `1 <= l_i <= u_i <= 10^9`
- The sum of `N` over all test cases is at most `10^5`.
- All input values are integers.

## Input

```text
T
case_1
case_2
...
case_T
```

Each test case is given in the following format:

```text
N H
t_1 l_1 u_1
t_2 l_2 u_2
...
t_N l_N u_N
```

## Output

Print `T` lines. The `i`-th line should contain `Yes` if it is possible to achieve all goals in the
`i`-th test case, and `No` otherwise.

## Sample Input 1

```text
3
2 5
3 1 4
8 9 11
2 6
1 1 4
3 5 8
10 36
27 37 38
30 34 54
38 20 77
45 1 36
49 38 51
52 31 58
65 43 60
71 14 42
73 36 38
85 14 29
```

## Sample Output 1

```text
Yes
No
Yes
```

### Note

- In the first test case, one valid flight is: stay at altitude 5 until time 2, remain at 4 from
  time 2 to 3, then ascend at rate 1 to reach altitude 9 at time 8.
- In the second test case, no flight satisfies both goals.

### ideas
1. 
