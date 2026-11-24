# B. Towers

[Problem link](https://codeforces.com/problemset/problem/479/B)

**Contest:** [Codeforces Round 274 (Div. 2)](https://codeforces.com/contest/479)

time limit per test: 1 second

memory limit per test: 256 megabytes

## Problem

Petya has `n` towers of equal-sized cubes; tower `i` has height `a_i`. The instability
of the set is the difference between the highest and the lowest tower.

In one operation he takes the top cube of some tower and puts it on top of a different
tower. He can perform at most `k` operations. Minimize the instability.

## Input

The first line contains two integers `n` and `k` (`1 <= n <= 100`, `1 <= k <= 1000`).

The second line contains `n` integers `a_i` (`1 <= a_i <= 10^4`) — the initial heights.

## Output

Print two integers `s` and `m` (`m <= k`) — the minimum possible instability and the
number of operations used.

Then print `m` lines, each with two integers `i` and `j` (`i != j`) — move the top cube
from tower `i` to tower `j`. Heights may become zero. Any optimal sequence is accepted.

## Examples

```text
Input
3 2
5 8 5

Output
0 2
2 1
2 3
```

```text
Input
3 4
2 2 4

Output
1 1
3 2
```

```text
Input
5 3
8 3 2 6 3

Output
3 3
1 3
1 2
1 3
```

## Solution

TODO
