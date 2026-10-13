# A. Common Prefixes

[Problem link](https://codeforces.com/problemset/problem/1384/A)

**Contest:** [Codeforces Round 659 (Div. 2)](https://codeforces.com/contest/1384)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

The longest common prefix of `s` and `t` is the largest `k` such that the first `k` characters of both strings are equal. If the first characters differ, `k` is `0`.

Koa had `n + 1` strings `s_1, s_2, …, s_{n+1}` and wrote down `a_i`, the longest common prefix length of `s_i` and `s_{i+1}`. Reconstruct any strings that produce this array. A solution always exists under the constraints.

## Input

The first line contains `t` (`1 <= t <= 100`), the number of test cases.

Each test case is two lines:

- `n` (`1 <= n <= 100`)
- `n` integers `a_1, a_2, …, a_n` (`0 <= a_i <= 50`)

The sum of `n` over all test cases does not exceed `100`.

## Output

For each test case, print `n + 1` lines. The `i`-th line is a string `s_i` of lowercase letters with length from `1` to `200`. The longest common prefix of `s_i` and `s_{i+1}` must have length `a_i`. Any valid list of strings is accepted.

## Example

### Input

```text
4
4
1 2 4 2
2
5 3
3
1 3 1
3
0 0 0
```

### Output

```text
aeren
ari
arousal
around
ari
monogon
monogamy
monthly
kevinvu
kuroni
kurioni
korone
anton
loves
adhoc
problems
```

In the first test case, one possible answer is `aeren`, `ari`, `arousal`, `around`, `ari`. The consecutive prefix lengths are `1`, `2`, `4`, `2`.

## Status

I/O and official sample cases are in place. `solve` is left as a TODO.
