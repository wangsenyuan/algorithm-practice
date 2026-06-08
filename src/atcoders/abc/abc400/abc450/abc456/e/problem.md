# E - Endless Holidays (ABC456)

**Contest:** [ABC456](https://atcoder.jp/contests/abc456) — AtCoder Beginner Contest 456  
**Task:** [https://atcoder.jp/contests/abc456/tasks/abc456_e](https://atcoder.jp/contests/abc456/tasks/abc456_e)

**Time limit:** 3 sec / **Memory limit:** 1024 MiB  
**Score:** 450 points

## Problem Statement

The Kingdom of AtCoder has `N` cities, called city `1, 2, ..., N`. There are
`M` bidirectional roads connecting pairs of cities, where the `i`-th road
connects cities `U_i` and `V_i`. Any pair of cities can be reached from each
other by traversing some roads.

In the Kingdom of AtCoder, a week consists of `W` days. A week proceeds through
days `1, 2, ..., W`, and the day after day `W` is day `1`.

Each city has certain days of the week that are holidays. The holiday
information for city `i` is given as a string `S_i` of length `W`: if the
`j`-th character of `S_i` is `o`, day `j` is a holiday; if it is `x`, day `j`
is a weekday.

Takahashi chooses a city he likes and visits it at noon on day `1`. Each night
thereafter, he repeatedly chooses to either stay in his current city or move to
a city directly connected by a road.

Output `Yes` if it is possible for him to keep moving so that the city he is in
at noon every day is a holiday, and `No` otherwise.

`T` test cases are given; solve each of them.

## Constraints

- `1 <= T <= 10^5`
- `1 <= N <= 10^5`
- `N - 1 <= M <= 10^5`
- `1 <= U_i < V_i <= N`
- Any pair of cities can be reached from each other by traversing some roads.
- `2 <= W <= 10`
- `T, N, M, U_i, V_i, W` are integers.
- `S_i` is a string of length `W` consisting of `o`, `x`.
- The sum of `N` over all test cases is at most `10^5`.
- The sum of `M` over all test cases is at most `10^5`.

## Input

The input is given from Standard Input in the following format:

```text
T
case_1
case_2
...
case_T
```

Here, `case_i` denotes the input for the `i`-th test case. Each test case is
given in the following format:

```text
N M
U_1 V_1
U_2 V_2
...
U_M V_M
W
S_1
S_2
...
S_N
```

## Output

Output `T` lines. The `i`-th line should contain the answer for the `i`-th test
case.

## Sample Input 1

```text
3
4 4
1 2
1 4
2 4
2 3
3
xxo
xox
oxo
oxx
1 0
4
oooo
5 5
1 4
2 3
4 5
3 4
2 5
7
oxxxxxx
xxoxxxo
xxxoxox
xoxxoxx
oxxxoxx
```

## Sample Output 1

```text
Yes
Yes
No
```

For the first test case, for example, the condition can be satisfied by
repeatedly moving along cities `4 -> 2 -> 1 -> 4 -> 2 -> 1 -> ...`.
Alternatively, the condition can also be satisfied by repeatedly moving along
cities `3 -> 2 -> 3 -> 3 -> 2 -> 3 -> ...`.

For the second test case, the condition can be satisfied by staying in city `1`
indefinitely.

For the third test case, it is impossible to keep moving while satisfying the
condition.


### ideas
1. 感觉要形成一个环