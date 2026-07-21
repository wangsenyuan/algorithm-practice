# H. La Vaca Saturno Saturnita

[Problem link](https://codeforces.com/problemset/problem/2094/H)

time limit per test: 4 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Saturnita's mood depends on an array `a` of length `n` and a function
`f(k, a, l, r)`. The pseudocode for `f` is:

```text
function f(k, a, l, r):
  ans := 0
  for i from l to r (inclusive):
    while k is divisible by a[i]:
      k := k / a[i]
    ans := ans + k
  return ans
```

You are given `q` queries, each containing integers `k`, `l`, and `r`. For each
query, output `f(k, a, l, r)`.

## Constraints

- `1 <= t <= 10^4`
- `1 <= n <= 10^5`
- `1 <= q <= 5 * 10^4`
- `2 <= a_i <= 10^5`
- `1 <= k <= 10^5`
- `1 <= l <= r <= n`
- Sum of `n` over all test cases does not exceed `10^5`
- Sum of `q` over all test cases does not exceed `5 * 10^4`

## Input

The first line contains an integer `t` — the number of test cases.

For each test case:

- The first line contains two integers `n` and `q`.
- The second line contains `n` integers `a_1, a_2, ..., a_n`.
- Each of the next `q` lines contains three integers `k`, `l`, and `r`.

## Output

For each query, output the answer on a new line.

## Example

```text
Input
2
5 3
2 3 5 7 11
2 1 5
2 2 4
2310 1 5
4 3
18 12 8 9
216 1 2
48 2 4
82944 1 4

Output
5
6
1629
13
12
520
```


### ideas
1. k <= 1e5
2. 在区间[l...r]中, 找到第一个可以整除k的数
3. 如果能快速找到的话, 那么这样的操作不会超过20次
4. 把这个数组分块, 每个分块内dp[block][x] 表示第一个能整除x的最小的下标
5. 那么可以在sqrt(n) * k的时间内处理完

## Solution

The block table above is too expensive in the worst case. If every `a[i]` is
`2`, filling all multiples of every array element performs about five billion
updates. The table itself also needs roughly `sqrt(n) * 100000` integers.

Instead, store the sorted occurrence positions of every value in `a`.

For a query with initial value `k`, the value can change at position `i` only
when `a[i]` divides the current `k`. The current value of `k` always divides its
initial value, so every relevant `a[i]` must be a divisor of the initial `k`.

For each divisor `d` of the initial `k`, binary-search its first occurrence in
`[l, r]`. No later occurrence of the same `d` can matter:

- if `d` divides the current `k` at its first occurrence, the loop removes all
  possible factors of `d`, so `d` cannot divide `k` afterward;
- if `d` does not divide the current `k`, later values of `k` are divisors of
  this one, so `d` can never become a divisor again.

Sort these first occurrences by position. Between two consecutive candidate
positions, `k` cannot change, so the whole segment contributes its length times
the current `k`. At each candidate, perform the divisions exactly as in the
statement.

Let `tau(k)` be the number of divisors of `k`. Preprocessing takes `O(n)` time
and memory. Each query takes
`O(sqrt(k) + tau(k) log n + tau(k) log tau(k))` time and uses `O(tau(k))`
temporary memory. Under `k <= 100000`, `tau(k) <= 128`.

## Solution Summary

Build `positions[v]`, the sorted list of indices where `a[i] = v`. For each
query, enumerate every divisor `d` of the initial `k` and binary-search the
first occurrence of `d` inside `[l, r]`. Sort the resulting `(position, d)`
events by position.

Only these events can change `k`. Moreover, only the first occurrence of each
divisor matters: after processing it, either all possible factors of `d` have
been removed, or `d` already does not divide the current `k`; subsequent
divisions can never make `d` divide `k` again. Therefore, all positions between
two events have the same `k` and their contribution is computed at once as
`segmentLength * k`.

The occurrence lists require `O(n)` preprocessing time and memory. A query uses
`O(sqrt(k) + tau(k) log n + tau(k) log tau(k))` time and `O(tau(k))` temporary
memory, where `tau(k) <= 128` for `k <= 100000`.
