# D. Yaroslav and Productivity

[Problem link](https://codeforces.com/problemset/problem/2244/D)

**Contest:** [Codeforces Round 1109 (Div. 3)](https://codeforces.com/contest/2244)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Yaroslav's productivity during the day is described by an array `a` of length `n`. His productivity can be negative if he watches short videos, or positive if he works. The total productivity is defined as the sum of all values in the array.

Sometimes Yaroslav reads motivational posts. He has `m` posts, where the `j`-th post has an impact value `b_j`. If Yaroslav reads a post with value `b_j`, then all productivity values from the beginning of the day up to position `b_j` change sign, that is, all integers `a_1, a_2, a_3, …, a_{b_j}` are multiplied by `-1`.

For example, let `a = [1, -4, 3, -4]`, and Yaroslav reads a post with value `3`. Then the signs of the first three elements are flipped, and the array becomes `[-1, 4, -3, -4]`. If he then reads a post with value `1`, the first element changes sign again, and the array becomes `[1, 4, -3, -4]`.

What is the maximum possible total productivity Yaroslav can achieve by reading any (possibly none) of the posts?

## Input

The first line contains a single integer `t` (`1 ≤ t ≤ 10^4`) — the number of test cases.

The first line of each test case contains two integers `n` and `m` (`1 ≤ m ≤ n ≤ 2 · 10^5`) — the number of measurements and the number of motivational posts.

The second line of each test case contains `n` integers `a_i` (`-10^9 ≤ a_i ≤ 10^9`) — the productivity values.

The third line of each test case contains `m` integers `b_i` (`1 ≤ b_i ≤ n`) — the impact values of the posts. It is guaranteed that all `b_i` are distinct.

It is guaranteed that the sum of `n` over all test cases does not exceed `2 · 10^5`.

## Output

For each test case, output a single integer — the maximum possible total productivity.

## Example

### Input

```text
4
5 3
-1 2 -3 4 -5
1 5 3
4 2
3 -1 3 -1
4 2
3 1
-5 -5 -5
2
4 3
3 -1 1 -3
4 3 2
```

### Output

```text
3
4
5
6
```
