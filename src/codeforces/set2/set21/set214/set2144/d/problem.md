# D. Price Tags

[Problem link](https://codeforces.com/problemset/problem/2144/D)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

You are the owner of a store with `n` items. The `i`-th item costs `c_i` coins and
has a price tag showing `c_i`.

You hold a sale in the format: "we divided all the prices `x` times." Formally,
choose an integer coefficient `x` (the same for every item). During the sale, the
`i`-th item costs `ceil(c_i / x)` coins, where `ceil(y)` denotes rounding up.

You may reuse existing price tags by pinning them on other items, and only print
new tags for prices that are not already available. Printing one price tag costs
`y` coins.

Choose an integer `x` strictly greater than `1` to maximize total income:

```text
(total value of items at sale prices) - (cost of printed price tags)
```

For each test case, output the maximum possible total income.

## Input

The first line contains a single integer `t` (`1 <= t <= 10`) — the number of test
cases.

The first line of each test case contains two integers `n` and `y`
(`1 <= n <= 2 * 10^5`; `1 <= y <= 10^9`) — the number of items and the cost of
printing one price tag.

The second line contains `n` integers `c_1, c_2, ..., c_n`
(`1 <= c_i <= 2 * 10^5`) — the initial prices.

## Output

For each test case, output a single integer — the maximum total income.

## Example

### Input

```text
4
5 51
150 150 50 148 150
3 1000000000
42 42 42
10 5
43211 8088 45 1 73 1 9198 4991 1 833 1001
11 5
8088 45 1 73 1 9198 4991 1 833 1001 1 1
```

### Output

```text
31
-2999999937
-1627553
21242
```

## Note

In the first test case, it is optimal to choose `x = 3`. The new prices become
`[50, 50, 17, 50, 50]` (in some order). Two old tags with price `50` can be reused,
and three new tags must be printed. Income is:

```text
50 + 50 + 17 + 50 + 50 - 51 * 3 = 31
```

In the second test case, it is optimal to choose `x = 2`. The new prices are
`[21, 21, 21]`, and all three price tags must be printed. With `y = 10^9`, the
income is negative.

In the third test case, it is optimal to choose `x = 111`.

In the fourth test case, it is optimal to choose `x = 2`. The new prices match the
old prices, so no new tags need to be printed and income equals the sum of all
initial prices.


### ideas
1. 假设 x, 那么c[1]/x, c[2]/x, ... c[i]/x, 假设其中有k个相同，那么需要支付 (n - k) * y
2. x越小（但是大于1）, income越大
3. 给定x的时候，如果能快速计算k的话，那么就好办了
4. 比如 8 / 3 = 2, 8 / 4 = 2
5. 对于一个数w, 当x处在一个区间[lo, hi)时，w/lo = w/(hi - 1)
6. 排序后， 对于c[1], 它能够被重复使用的情况是否是一个区间？
7. 假设c[1] = 3, 那么x = 2时，如果存在6,7, 那么c[1]可以被覆盖
8. 如果x = 3, 那么存在8，9，10时可以被覆盖，
9. c[1]可以贡献2,3这个可以检查出来（但是这样的x最多时多少个呢？）
10. 好像可以搞

## Solution

Fix the sale coefficient `x`. An item with original price `c` becomes:

```text
ceil(c / x)
```

For a target sale price `i`, the items that become `i` are exactly:

```text
(i - 1) * x < c <= i * x
```

So if we know the frequency prefix sum of original prices, we can count these
items quickly.

Suppose:

- `v` items become sale price `i`,
- `w` old price tags with value `i` already exist.

Then at most `min(v, w)` of those `v` items can reuse existing tags. The remaining
`v - min(v, w)` tags must be printed. Therefore the contribution of sale price
`i` for this fixed `x` is:

```text
v * i - (v - min(v, w)) * y
```

The code accumulates this contribution into:

```go
dp[x] += v*i - (v-z)*y
```

where `z = min(v, w)`.

### Enumerating `x`

The maximum original price is at most `2 * 10^5`, so we can enumerate all possible
sale prices `i`.

For a fixed `i`, iterate over multiples:

```go
j = i*x
```

where `x >= 2`. Then the original prices that become `i` are:

```text
((i - 1) * x, i * x]
```

In the code:

```go
v := freq[min(mx, j)] - freq[(i-1)*x]
```

Here `freq` is already a prefix count array, so this gives the number of original
prices in that interval.

The loop stops once `(i - 1) * x > mx`, because then no original price can still
fall into the interval.

### Handling Price `1`

Items with original price `1` are special. For every `x > 1`:

```text
ceil(1 / x) = 1
```

They always keep price `1`, and there are exactly enough old `1` tags for them, so
they always contribute `1` each and never require printing. The code removes them
from the main counting:

```go
special := freq[1]
n := len(c) - freq[1]
freq[1] = 0
```

Then it solves only the non-`1` items and adds `special` back at the end.

If all items are `1`, the answer is simply `n`.

### Validating Each `x`

`fp[x]` counts how many non-`1` items have been assigned a sale price while
computing contributions for coefficient `x`.

Only coefficients with:

```go
fp[x] == n
```

are valid complete evaluations. For those, `dp[x]` is the total income for that
coefficient, and the answer is the maximum `dp[x] + special`.

### Complexity

For each sale price `i`, the inner loop iterates over multiples of `i`, so the
total work is:

```text
sum(mx / i) = O(mx log mx)
```

where `mx = max(c_i) <= 2 * 10^5`. This is fast enough for all test cases.
