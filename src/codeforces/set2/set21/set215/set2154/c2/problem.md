# C2. No Cost Too Great (Hard Version)

[Problem link](https://codeforces.com/problemset/problem/2154/C2)

time limit per test: 3 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

This is the hard version. The difference is that `1 <= b_i <= 10^9`.

You are given two arrays of positive integers `a` and `b`, both of length `n`.
You may perform the following operation any number of times:

- choose an index `i` (`1 <= i <= n`) and increase `a_i` by `1`; this operation
  costs `b_i`.

Find the minimum total cost needed so that there exist two indices `i < j` with
`gcd(a_i, a_j) > 1`.

`gcd(x, y)` denotes the greatest common divisor of integers `x` and `y`.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test
cases.

For each test case:

- The first line contains an integer `n` (`2 <= n <= 2 * 10^5`) — the length of
  the arrays.
- The second line contains `n` integers `a_1, a_2, ..., a_n`
  (`1 <= a_i <= 2 * 10^5`).
- The third line contains `n` integers `b_1, b_2, ..., b_n`
  (`1 <= b_i <= 10^9`).

The sum of `n` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, output the minimum cost.

## Example

### Input

```text
6
2
1 1
1 2
2
4 8
41 67
5
1 1 727 1 1
1 1 1000 1 1
2
3 11
1 1
3
2 7 11
1 6 6
3
2 7 11
100 1 2
```

### Output

```text
3
0
2
1
5
1
```

## Notes

In the first test case, one optimal sequence is:

```text
[1, 1] -> [2, 1] -> [2, 2]
```

The cost is `1 + 2 = 3`, and then `gcd(2, 2) = 2`.

In the second test case, `gcd(4, 8) = 4`, so the condition is already satisfied
and the answer is `0`.
