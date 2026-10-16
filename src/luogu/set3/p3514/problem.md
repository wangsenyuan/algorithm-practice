# P3514 [POI2011] LIZ-Lollipop

[Problem link](https://www.luogu.com.cn/problem/P3514)

time limit per test: 2.00s

memory limit per test: 128.00MB

A lollipop is a row of `n` segments. `W` is vanilla and costs 1, `T` is strawberry and costs 2. For each query `x`, find a contiguous fragment whose costs sum to `x`, or report that none exists.

## Input

The first line contains two integers `n` and `m` (`1 <= n, m <= 10^6`), the number of segments and the number of queries.

The second line is a string of length `n` consisting of `T` and `W`.

Each of the next `m` lines contains an integer `x` (`1 <= x <= 2 * 10^6`).

## Output

Print `m` lines. For each query, print `l r` (`1 <= l <= r <= n`) if segments `l` through `r` sum to `x`. Any such pair is accepted. If none exists, print `NIE`.

Each answer is followed by a newline, and the output ends after the last newline.

## Sample

### Input

```text
5 3
TWTWT
5
1
7
```

### Output

```text
1 3
2 2
NIE
```

Segments 1 through 3 are `TWT` and sum to 5. Segment 2 is `W` and costs 1. No fragment sums to 7.

## Solution

Every cost is positive, so the largest subarray sum is the whole string. The largest sum of the other parity is a prefix or a suffix: any subarray is the whole string minus a prefix and a suffix, and the smallest removal of the required parity is a prefix alone or a suffix alone. Scanning both directions records one segment for every prefix sum and suffix sum, and keeps the best even sum and the best odd sum.

A segment of sum `k >= 2` contains a segment of sum `k - 2`. Drop the left end when it is `T`, otherwise the right end when it is `T`, otherwise both ends are `W` and dropping them also removes 2. Starting from each recorded best segment and repeating this step fills in every smaller positive sum of the same parity. A sum that is never filled in, including every query larger than `2n`, is `NIE`.

The scans and the two downward walks are `O(n)`, and each query is `O(1)`.
