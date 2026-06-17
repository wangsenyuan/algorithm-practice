# E. New Year's Gifts

[Problem link](https://codeforces.com/problemset/problem/2182/E)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

Monocarp has `n` friends and `m` gift boxes. Box `i` has beauty `a_i`. Each box
holds at most one gift.

He must buy every friend a gift costing at least `y_i` coins. A gift may be placed
in a box or given without a box.

Friend `i` is **happy** if at least one of these holds:

- the gift is **in** a box with beauty at least `x_i`;
- the gift costs at least `z_i` coins (`z_i > y_i`).

Given `k` coins, maximize the number of friends Monocarp can make happy.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases.

Each test case:

- One line: `n`, `m`, `k` (`1 <= n, m <= 2 * 10^5`, `1 <= k <= 10^15`).
- One line: `m` integers `a_1, …, a_m` (`1 <= a_i <= m`) — box beauties.
- Then `n` lines: `x_i`, `y_i`, `z_i` (`1 <= x_i <= m`, `1 <= y_i < z_i <= 10^9`).

Additional constraints:

- `sum(y_i) <= k` for each test case.
- Sum of `n` over all test cases `<= 2 * 10^5`.
- Sum of `m` over all test cases `<= 2 * 10^5`.

## Output

For each test case, print one integer — the maximum number of friends that can be
made happy.

## Example

### Input

```text
3
2 1 6
1
1 2 3
1 2 7
2 2 3
1 1
2 1 5
2 1 5
3 4 11
1 1 1 2
1 2 13
3 2 5
3 1 3
```

### Output

```text
2
0
2
```

## Note

**Test 1:** Give friend 1 a gift costing `3` coins (happy via `z_1`). Give friend 2 a
gift costing `2` coins **in** the box with beauty `1` (happy via `x_2`). Both are
happy.

**Test 2:** No friend can be happy: `k = 3` is not enough for any `z_i`, and every
`x_i` exceeds all box beauties (`1`).

**Test 3:** One optimal plan: gifts costing `2`, `6`, and `3` coins for friends 1–3.
Friends 2 and 3 become happy; friend 1’s gift is not in a box and is too cheap for
`z_1`.


### ideas
1. sort boxes, 尽量使用高价值的盒子; 假设有w个盒子可以用(装了w个礼物) 那么这些礼物可以使用y[i]来支付
2. 其他的没有装入盒子的,必须使用z[i]来支付
3. 所以没有装入盒子的, 应该选择z小的那部分
4. 但是装入盒子的部分, 尽量选择y小的那部分
5. 按照 z - y升序排列, 这样子, 左端应该是不使用盒子的, 右端是使用盒子的部分
6. d[i] = z[i] - y[i], d[i] < d[j] 如果j不使用盒子, i使用盒子, 如果某个盒子, w >= x[i], w >= x[j]
7. 那么显然交换它们, 收益更大. 原来的付出 = y[i] + z[j], 交换后, z[i] + y[j]
8. y[i] + z[j] - (z[i] + y[j]) = (z[j] - y[j]) - (z[i] - y[i]) > 0 (收益为正)
9. 但是还有两个因素, 应该是总支出 <= k, 还有x
10. 所有的礼物都必须买, 且总支出 <= k 的前提下, 让最多的人开心;
11. 不需要最小化总支出, 是要最大化开心的人数(似乎可以二分)
12. 如果不使用盒子, 那么支出 = sum(z), 装入一个盒子, 相当于节省 z[i] - y[i]的收益
13. 现在考虑盒子w, 要转入它的, 应该是 x[?] <= w 中最大节省的那个(这样能)
14. 