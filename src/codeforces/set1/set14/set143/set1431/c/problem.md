# C. Black Friday

[Problem link](https://codeforces.com/problemset/problem/1431/C)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

A local shop is preparing for Black Friday. There are `n` items with prices
`p_1, p_2, ..., p_n` on display, sorted so that `p_1 <= p_2 <= ... <= p_n`.

The shop uses a fixed discount parameter `k`. For a single purchase of `x` items
(you may buy each item at most once), you get the cheapest `floor(x / k)` purchased
items for free.

You do not care how much money you spend. You want the total price of the items you
receive for free to be as large as possible.

Find the maximum possible total price of free items from one purchase.

## Input

The first line contains an integer `t` (`1 <= t <= 500`) — the number of test cases.

Each test case contains:

- one line with two integers `n` and `k` (`1 <= n <= 200`; `1 <= k <= n`)
- one line with `n` integers `p_1, p_2, ..., p_n` (`1 <= p_i <= 10^6`), sorted in
  non-decreasing order

## Output

For each test case, print one integer — the maximum total price of items you can
get for free in a single purchase.

## Example

### Input

```text
5
5 2
1 3 4 6 7
6 3
1 2 3 4 5 6
6 3
3 3 4 4 5 5
1 1
7
4 4
1 3 3 7
```

### Output

```text
7
4
6
7
1
```

## Note

In the first test case, one optimal purchase is to buy items with prices
`[3, 4, 6, 7]`. Since `x = 4` and `k = 2`, you get `floor(4 / 2) = 2` cheapest
items for free: `3` and `4`, for a total of `7`.

In the problem statement example, if the shop has prices
`[1, 1, 2, 2, 2, 3, 4, 5, 6]` and you buy `[1, 2, 2, 4, 5]` with `k = 2`, then
`floor(5 / 2) = 2` free items are the cheapest purchased ones: `1` and `2`.
