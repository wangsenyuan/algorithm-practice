# P3188 [HNOI2007] 梦幻岛宝珠

[Problem link](https://www.luogu.com.cn/problem/P3188)

time limit per test: 1.00s

memory limit per test: 125.00MB

There are `n` gems. Each gem has a weight and a value. Choose a subset whose total weight is at most `W` and whose total value is as large as possible.

## Input

The input contains several test cases. Each case is:

- one line with two positive integers `n` and `W`, the number of gems and the weight limit
- `n` lines, each with two integers `weight_i` and `value_i`

After the last case there is a line `-1 -1`. That line is a terminator, not a case, and produces no output.

## Output

For each case, print one integer: the maximum total value.

## Sample

### Input

```text
4 10
8 9
5 8
4 6
2 5
4 13
8 9
5 8
4 6
2 5
16 75594681
393216 5533
2 77
32768 467
29360128 407840
112 68
24576 372
768 60
33554432 466099
16384 318
33554432 466090
2048 111
24576 350
9216 216
12582912 174768
16384 295
1024 76
-1 -1
```

### Output

```text
14
19
1050650
```

## Constraints

- at most `20` test cases
- `1 <= n <= 100`
- `1 <= W <= 2^30`
- `1 <= weight_i <= 2^30`, `0 <= value_i <= 2^30`
- each `weight_i` can be written as `a * 2^b` with `1 <= a <= 10` and `0 <= b <= 30`
- the answer does not exceed `2^30`

## Solution

`W` is too large for a plain 0-1 knapsack, but every weight has the form `a * 2^b` with `a <= 10`. Split each gem by `b = bits.TrailingZeros(weight)` and `a = weight >> b`, and solve the groups from low `b` to high `b`.

For a single exponent `b`, the items are an ordinary 0-1 knapsack on the coefficient `a`. Let `g[b][j]` be the best value whose coefficients sum to at most `j`. Values are nonnegative and the capacity loop runs downward, so each gem is used once and `g[b][j]` is the at-most-`j` optimum. No group can spend more than `10 * n <= 1000` coefficient, so the second dimension stops at `1000`. A gem with `b >= bitlen(W)` weighs at least `2^b > W` and is discarded.

Merge these groups into

```text
f[i][j] = best value from gems with exponent <= i
          and weight <= j * 2^i + (W mod 2^i)
```

`f[0]` is `g[0]`. For `i >= 1`, write `bit` for bit `i-1` of `W` and split the `j` high units into `j-k` units taken from group `i` and `k` units left for the lower groups:

```text
k * 2^i + (W mod 2^i) = (2k + bit) * 2^{i-1} + (W mod 2^{i-1})
```

so

```text
f[i][j] = max over k = 0..j of
    g[i][j-k] + f[i-1][min(2k + bit, 1000)]
```

The `1000` cap is safe: every lower gem contributes at most `a` units of `2^{i-1}`, and the sum of those `a` is at most `1000`. The code stores `g[i]` and `f[i]` in the same row. `j` descends and each transition reads `f[i][j-k]` with `j-k < j`, which has not been merged yet, so that entry is still `g[i][j-k]`.

`bits.Len(W) - 1` is the top bit of `W`, and `W` shifted by that much is `1`. Therefore `f[bits.Len(W)-1][1]` is the best value for weight limit `W`.

Each of the at most `31` merges tries `O(1000^2)` splits, and there are at most `20` cases.
