# F. Minimize Fixed Points

[Problem link](https://codeforces.com/problemset/problem/2123/F)

time limit per test: 3 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

A permutation `p` of length `n` is **good** if `gcd(p_i, i) > 1` for all `2 <= i <= n`.

Find a good permutation with the minimum number of fixed points among all good
permutations of length `n`. A fixed point is an index `j` (`1 <= j <= n`) with
`p_j = j`. If multiple answers exist, print any of them.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases.

Each test case contains a single integer `n` (`2 <= n <= 10^5`).

The sum of `n` over all test cases does not exceed `10^5`.

## Output

For each test case, print one line: a good permutation of length `n` with the
minimum possible number of fixed points.

## Example

### Input

```text
4
2
3
6
13
```

### Output

```text
1 2
1 2 3
1 4 6 2 5 3
1 12 9 6 10 8 7 4 3 5 11 2 13
```

## Note

For `n = 6`, the permutation `[1, 4, 6, 2, 5, 3]` is good:

| i | p_i | gcd(p_i, i) |
|---|-----|-------------|
| 1 | 1   | 1           |
| 2 | 4   | 2           |
| 3 | 6   | 3           |
| 4 | 2   | 2           |
| 5 | 5   | 5           |
| 6 | 3   | 3           |

`gcd(p_i, i) > 1` holds for all `i >= 2`. Fixed points are at indices `1` and `5`
only; fewer fixed points are impossible.
