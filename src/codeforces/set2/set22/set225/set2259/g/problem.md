# G. Index Removal

[Problem link](https://codeforces.com/problemset/problem/2259/G)

**Contest:** [Codeforces Round 1119 (Div. 3)](https://codeforces.com/contest/2259)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

An array `b_1, b_2, …, b_m` is good if, for each `i` (`1 <= i < m`), `b_{i+1} >= b_i` and `b_{i+1} - b_i <= k`. An array of length `1` is always good.

You are given an initially good array `a_1, a_2, …, a_n`. Solve the following problem for each `i` (`1 <= i <= n`) independently:

- The `i`-th element of `a` is removed, making `a = [a_1, a_2, …, a_{i-1}, a_{i+1}, …, a_n]`. In one operation, you are able to subtract `1` from any element of `a`. What is the minimum number of operations required to make `a` good again?

## Input

The first line of each input contains `t` (`1 <= t <= 10^4`) — the number of test cases.

The first line of each test case contains two integers `n` and `k` (`2 <= n <= 2 * 10^5`, `1 <= k <= 10^9`).

The second line of each test case contains `n` integers `a_1, a_2, …, a_n` (`1 <= a_i <= 10^9`). It is guaranteed that `a` is initially good.

It is guaranteed that the sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output `n` space-separated integers on a single line: the `i`-th integer denoting the solution to the problem for the `i`-th index.

## Example

### Input

```text
7
4 2
1 2 4 5
4 1
1 2 3 4
5 7
1 8 9 16 20
5 1000000000
1 6 7 67 6767
6 1
1 1 2 2 3 4
4 1
1 2 3 3
4 2
1 2 4 6
```

### Output

```text
0 1 1 0
0 2 1 0
0 2 1 4 0
0 0 0 0 0
0 0 0 0 1 0
0 1 0 0
0 2 2 0
```

In the first test case:

- `i = 1`, `a = [2, 4, 5]`. Since `a` is still good, we do not need to perform any operations.
- `i = 2`, `a = [1, 4, 5]`. If we perform the operation once on the 2nd element, `a = [1, 3, 5]`, which is a good array.
- `i = 3`, `a = [1, 2, 5]`. If we perform the operation once on the 3rd element, `a = [1, 2, 4]`, which is a good array.
- `i = 4`, `a = [1, 2, 4]`. Since `a` is still good, we do not need to perform any operations.

In the fourth test case, `a` will remain good no matter what element we remove, so the answer for each index is `0`.

## Status

I/O and official sample cases are in place. `solve` is left as a TODO.
