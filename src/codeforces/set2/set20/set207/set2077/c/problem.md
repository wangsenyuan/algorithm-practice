# C. Binary Subsequence Value Sum

[Problem link](https://codeforces.com/problemset/problem/2077/C)

time limit per test: 3 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

For a binary string `v`, define

```text
F(v, l, r) = length(v[l..r]) - 2 * number_of_zeroes(v[l..r])
```

and `F(v, l, r) = 0` when `l > r`.

The score of `v` is the maximum value of

```text
F(v, 1, i) * F(v, i + 1, |v|)
```

over all split positions `0 <= i <= |v|`.

You are given a binary string `s`. Each query flips one position of `s`, and the
modifications are persistent. After every flip, output the sum of scores over all
non-empty subsequences of the current `s`, modulo `998244353`.

## Input

The first line contains `t`, the number of test cases.

For each test case:

- The first line contains `n` and `q`.
- The second line contains the binary string `s`.
- The next `q` lines each contain an index to flip.

The total `n` and total `q` over all test cases are both at most `2 * 10^5`.

## Output

For each query, output the required sum after applying that flip.

## Example

### Input

```text
3
3 2
010
1
3
10 3
0101000110
3
5
10
24 1
011001100110000101111000
24
```

### Output

```text
1
5
512
768
1536
23068672
```

## Solution Summary

Only the number of zeroes matters.

Treat `1` as `+1` and `0` as `-1`. For a subsequence `v`, let

```text
d = count_1(v) - count_0(v)
```

For any split, the two parts have sums `x` and `d - x`, so their product is
`x(d - x)`. The maximum over integer splits is determined by `d`, and equals

```text
floor(d / 2) * ceil(d / 2)
```

So for the whole string `s`, we need to sum this value over all non-empty
subsequences. If `cnt0` is the number of zeroes in `s`, then among all
subsequences the possible signed sums can be counted by binomial coefficients.
After simplifying the binomial sums, the answer is:

```text
2^(n - 4) * (n * (n + 1) - 4 * cnt0 * n + 4 * cnt0^2 - 2)
```

Modulo arithmetic handles `2^(n - 4)` as `2^n / 16`, using the modular inverse
of `16`.

Each query only flips one bit, so update `cnt0` by `+1` or `-1` and recompute the
closed form in `O(1)`.

## Complexity

For each test case, counting the initial zeroes costs `O(n)`, and all queries
cost `O(q)`. Memory usage is `O(n + q)` for the mutable string and answers.


### ideas
1. F(v, l, r) = 在v中, 处于l...r段1的个数 - 0的个数
2. S(v) = max over i, F(v, 1, i) * F(v, i+1, k), k = len(v)
3. 