# D. Good Schedule

[Problem link](https://codeforces.com/problemset/problem/2230/D)

**Contest:** [Educational Codeforces Round 190 (Rated for Div. 2)](https://codeforces.com/contest/2230)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem Statement

Alice and Bob decided to watch a TV series consisting of `n` episodes, numbered
from `1` to `n`. The series will be shown on television over the next `n` days.
On the `i`-th day, episode `a_i` is shown in Alice's city, and episode `b_i` in
Bob's city.

They plan to select a segment of days `[L, R]` (`1 <= L <= R <= n`) to watch the
series. Initially, neither of them has seen any episodes. Each day `i` in this
segment:

- if Alice has already watched episodes `1, 2, ..., a_i - 1`, but not `a_i`,
  then she watches `a_i` on day `i`; otherwise, she watches nothing;
- if Bob has already watched episodes `1, 2, ..., b_i - 1`, but not `b_i`, then
  he watches `b_i` on day `i`; otherwise, he watches nothing.

To avoid spoilers, they want a segment `[L, R]` such that on every day in the
segment, one of the following holds:

- both watch the same episode on that day; or
- neither watches anything on that day.

Count the number of suitable segments `[L, R]`.

## Input

The first line contains a single integer `t` (`1 <= t <= 10^4`) — the number of
test cases.

Each test case consists of three lines:

- the first line contains a single integer `n` (`1 <= n <= 5 * 10^5`);
- the second line contains `n` integers `a_1, a_2, ..., a_n` (`1 <= a_i <= n`);
- the third line contains `n` integers `b_1, b_2, ..., b_n` (`1 <= b_i <= n`).

The sum of `n` over all test cases does not exceed `5 * 10^5`.

## Output

For each test case, print a single integer — the number of suitable segments.

## Sample Input 1

```text
4
3
1 2 1
1 2 2
2
1 1
1 2
5
1 3 2 1 4
1 4 2 3 2
9
1 1 3 1 1 3 2 3 1
1 3 1 1 3 1 2 1 3
```

## Sample Output 1

```text
4
0
7
12
```

In the first test case, the suitable segments are `[1,1]`, `[1,2]`, `[1,3]`, and
`[2,2]`.


### ideas
1. 还很难么~
2. 越来越乱了