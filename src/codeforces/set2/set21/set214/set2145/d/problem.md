# D. Inversion Value of a Permutation

[Problem link](https://codeforces.com/problemset/problem/2145/D)

time limit per test: 3 seconds

memory limit per test: 512 megabytes

input: stdin

output: stdout

A permutation of length `n` is an array of `n` integers, where each number from
`1` to `n` appears exactly once.

An inversion in a permutation `p` is a pair of indices `(i, j)` such that
`i < j` and `p[i] > p[j]`.

For a permutation `p`, we define its **inversion value** as the number of its
subsegments that contain at least one inversion. Formally, this is the number of
pairs of integers `(l, r)` (`1 <= l <= r <= n`) such that the subarray
`p[l], p[l+1], ..., p[r]` contains at least one inversion.

For example, for the permutation `[3, 1, 4, 2]`, the inversion value is `5`.

You are given two integers `n` and `k`. Your task is to construct a permutation of
length `n` with an inversion value equal to exactly `k`.

## Input

The first line contains one integer `t` (`1 <= t <= 500`) — the number of test
cases.

Each test case consists of a single line containing two integers `n` and `k`
(`2 <= n <= 30`; `0 <= k <= n * (n - 1) / 2`).

## Output

For each test case, output the answer as follows:

- if the desired permutation does not exist, output a single integer `0`;
- otherwise, output `n` distinct integers from `1` to `n` — the desired
  permutation. If there are multiple such permutations, you may output any of
  them.

## Example

### Input

```text
5
4 5
5 10
5 0
6 8
3 1
```

### Output

```text
3 1 4 2
5 4 3 2 1
1 2 3 4 5
2 3 5 6 1 4
0
```

## Note

For the first test case, the permutation `[3, 1, 4, 2]` has inversion value `5`.

For the third test case, the sorted permutation `[1, 2, 3, 4, 5]` contains no
inversions, so its inversion value is `0`.

For the last test case, no permutation of length `3` has inversion value `1`, so
the answer is `0`.


### ideas
1. 考虑给定一个P，怎么计算inversion value
2. P[i] 添加进来以后, dp[i] = dp[i-1] (i-1构成的inversion pair仍然是有效的)
3. 然后+那些i-1后缀不满足的部分；所以，假设i-1，后缀是一个递增序列，那么P[l] > P[i]
4. 那么增加的部分 = i - l
5. 考虑前面是一个很长的递增序列，现在添加1，那么贡献就是这个序列的长度
6. 54321 => 1 + 2 + 3 + 4 = 10
7. inverse(132) = 2 (1, 3) and (2, 3)
8. 