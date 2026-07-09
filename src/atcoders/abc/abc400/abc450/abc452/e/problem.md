# E - You WILL Like Sigma Problem

[Problem link](https://atcoder.jp/contests/abc452/tasks/abc452_e)

**Contest:** [AtCoder Beginner Contest 452](https://atcoder.jp/contests/abc452)

time limit: 2 sec

memory limit: 1024 MiB

score: 450 points

You are given a sequence of positive integers `A = (A_1, ..., A_N)` of length `N`
and a sequence of positive integers `B = (B_1, ..., B_M)` of length `M`.

Find the value, modulo `998244353`, of

```text
sum_{i=1}^{N} sum_{j=1}^{M} A_i * B_j * (i mod j)
```

## Constraints

- `1 <= N, M <= 5 * 10^5`
- `1 <= A_i, B_j <= 5 * 10^5`
- All input values are integers

## Input

```text
N M
A_1 A_2 ... A_N
B_1 B_2 ... B_M
```

## Output

Print the answer on a single line.

## Sample Input 1

```text
6 4
1 6 9 2 3 1
1 10 3 7
```

## Sample Output 1

```text
508
```

The sum of the 24 pairwise terms is `508`. For example:

- `A_1 * B_2 * (1 mod 2) = 1 * 10 * 1 = 10`
- `A_2 * B_3 * (2 mod 3) = 6 * 3 * 2 = 36`
- `A_3 * B_4 * (3 mod 4) = 9 * 7 * 3 = 189`

## Sample Input 2

```text
20 20
36625 195265 98908 111868 111868 47382 147644 472464 472464 416653 111868 195265 327972 327972 262769 75439 381156 451275 36625 195265
327972 111868 416653 177330 340019 262769 47382 262769 47382 340019 47382 262769 327972 327972 359676 381156 327972 36625 451275 381156
```

## Sample Output 2

```text
58141644
```


## Solution

Rewrite the double sum by fixing `j` first:

```text
answer = sum_j B_j * (sum_i A_i * (i mod j))
```

So for each `j`, we only need:

```text
cur(j) = sum_i A_i * (i mod j)
```

Use the identity:

```text
i mod j = i - floor(i / j) * j
```

For fixed `j`, the value `floor(i / j)` is constant on blocks:

```text
[0*j, 1*j), [1*j, 2*j), [2*j, 3*j), ...
```

If the current block is `[q*j, (q+1)*j)`, then every index `i` in this block has:

```text
i mod j = i - q*j
```

The contribution of this whole block is:

```text
sum A_i * (i - q*j)
= sum A_i * i - q*j * sum A_i
```

Therefore we precompute two prefix sums:

- `sum1[x] = A_0 + A_1 + ... + A_{x-1}`
- `sum2[x] = A_0*0 + A_1*1 + ... + A_{x-1}*(x-1)`

The implementation first inserts a dummy `0` at the front of `A` and `B`, so the original
problem indices become 1-based. Then for each fixed `j`, it iterates quotient blocks:

```go
l := q * j
r := min(n, l+j)
cur += sum2[r] - sum2[l]
cur -= (sum1[r] - sum1[l]) * (q*j)
```

This adds:

```text
sum_{i=l}^{r-1} A_i * (i mod j)
```

After all blocks for `j` are processed, multiply by `B_j` and add it to the answer:

```go
ans += cur * B_j
```

All operations are done modulo `998244353`.

The total number of processed blocks is:

```text
sum_{j=1}^{M} ceil(N / j) = O(N log M)
```

The memory complexity is `O(N)` for the two prefix-sum arrays.
