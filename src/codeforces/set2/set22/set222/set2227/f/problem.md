# F. It Just Keeps Going Sideways

[Problem link](https://codeforces.com/problemset/problem/2227/F)

**Contest:** [Codeforces contest 2227](https://codeforces.com/contest/2227)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem Statement

Yousef has `n` columns of cubes standing side by side. The `i`-th column
contains `a_i` identical unit cubes stacked vertically. Initially gravity pulls
cubes downwards, so every column `i` contains exactly `a_i` cubes at heights
`1, 2, ..., a_i`.

Suddenly, gravity shifts to the right. Each cube slides horizontally as far to
the right as possible. A cube cannot pass through or overlap other cubes, and
it must remain at its original height. The final configuration is uniquely
determined by the initial heights.

Before the gravity shift, Yousef may choose at most one column and decrease its
height by `1` (remove one cube from that column; a height may become `0`).

A cube's movement distance is `|j - i|`, where `i` is its original column index
and `j` is its column index after the gravity shift.

Find the maximum possible total movement distance (the sum of the movement
distances of all remaining cubes).

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test
cases.

Each test case:

- the first line contains a single integer `n` (`1 <= n <= 2 * 10^5`);
- the second line contains `n` integers `a_1, ..., a_n` (`1 <= a_i <= n`).

It is guaranteed that the sum of `n` over all test cases does not exceed
`2 * 10^5`.

## Output

For each test case, output a single integer — the maximum total movement
distance.

## Sample Input 1

```text
5
5
1 2 3 2 1
7
5 4 1 1 1 1 3
6
1 2 3 4 5 6
6
4 1 6 3 2 6
7
1 3 2 7 2 3 1
```

## Sample Output 1

```text
9
37
0
17
29
```

## Note

In the first test case, the initial total movement distance is `5`. Removing
the cube at index `5` makes the array `[1, 2, 3, 2, 0]`, after which cubes can
slide farther right, for a total distance of `9`.

In the third test case, the initial total is `0`, and removing any cube still
leaves total distance `0`.


## ideas
1. 对于第i列, 假设a[i] = 3, 那么删除下面两层, 会造成上层*塌缩*, 所以相当于删除了顶部的元素
2. 所以只需要考虑每列的顶部元素;
3. 那么就需要知道, 这一行, 在删除前有多少个(➡️的堆叠数量, 已经它右边最近的位置)
4. 最后的形状是确定的,
