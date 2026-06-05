# D. Shoe Store

[Problem link](https://codeforces.com/problemset/problem/166/D)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

The warehouse has `n` shoe pairs. Pair `j` is described by price `c_j` and size `s_j`.
All sizes `s_j` are distinct (at most one pair per size).

There are `m` customers. Customer `i` has `d_i` money and foot size `l_i`.

Customer `i` can buy pair `j` if:

- `c_j <= d_i` (enough money), and
- `l_i = s_j` or `l_i = s_j - 1` (foot size matches, or is one size smaller)

Each customer buys at most one pair, and each pair is sold to at most one customer.

Maximize the total profit — the sum of prices of sold pairs.

## Input

The first line contains an integer `n` (`1 <= n <= 10^5`) — the number of shoe
pairs.

Each of the next `n` lines contains two integers `c_i` and `s_i`
(`1 <= c_i, s_i <= 10^9`) — the price and size of the `i`-th pair. All `s_i` are
different.

The next line contains an integer `m` (`1 <= m <= 10^5`) — the number of
customers.

Each of the next `m` lines contains two integers `d_i` and `l_i`
(`1 <= d_i, l_i <= 10^9`) — the money and foot size of the `i`-th customer.

## Output

In the first line, print the maximum profit.

In the second line, print an integer `k` — the number of sold pairs.

In the next `k` lines, print the sales. Each line contains two integers: the
customer number and the shoe pair number. Customers and pairs are numbered from
`1` in input order.

You may print the `k` lines in any order. If several optimal answers exist, print
any of them.

## Example

### Input 1

```text
3
10 1
30 2
20 3
2
20 1
20 2
```

### Output 1

```text
30
2
2 3
1 1
```

Customer `2` (money `20`, size `2`) buys pair `3` (price `20`, size `3`).
Customer `1` (money `20`, size `1`) buys pair `1` (price `10`, size `1`).
Total profit is `30`.

### Input 2

```text
3
10 4
20 5
30 6
2
70 4
50 5
```

### Output 2

```text
50
2
2 3
1 2
```

Customer `2` buys pair `3` (price `30`). Customer `1` buys pair `2` (price `20`).
Total profit is `50`.


### ideas
1. 把鞋子按照s升序处理, s[0](最小的鞋)， 只能sell给l[?] = s[0] 或者l[?] = s[0] - 1 的人
2. 那么就从中找出出价最高的人d[?] >= c[0]
3. 但是这里有个问题，就是 l[?] = s[0]的人其实还有可能购买s[1] = s[0] + 1的鞋
4. 不大好处理
5. 假设 s[0], s[1], ... s[i] 满足 s[0] + 1 = s[1], s[1] + 1 = s[2], ... 一直到s[i]
6. 那么这个时候，似乎先处理s[i]是最优的? 
7. 但是貌似还有一个决策，就是选择l[?] = s[i], 还是l[?] = s[i] - 1 的也需要考虑
8. 这里对于s[i]去连接l[j], 如果s[i] = l[j], 或者 s[i] = l[j] + 1
9. 且这里最多选择2个最大的l[j] = s[i], 2个最大的l[j] = s[i] - 1 
10. dp[i][0] = s[i]匹配0个（不管是l[?] = s[i], 或者是l[?] = s[i] - 1)
11. dp[i][1] = s[i]匹配 l[?] = s[i] - 1 的
12. dp[i][2] = s[i]匹配 l[?] = s[i](最大的那个)
13. 这样子，就可以迁移状态了

## Detailed solution explanation

Process all relevant sizes in increasing order. A relevant size is either a shoe
size `s` or a customer foot size `l`. For each foot size `l`, keep only the two
customers with the largest money among customers whose foot size is `l`.

Why two customers are enough:

- A customer with foot size `l` can only buy shoe size `l` or shoe size `l + 1`.
- Since shoe sizes are distinct, foot size `l` can be involved with at most two
  shoes.
- Therefore, at most two customers of the same foot size can be used in an
  optimal answer.
- If we need one or two customers from the same foot size, choosing among the
  two richest customers is always enough, because poorer customers cannot buy
  any shoe that a richer customer of the same foot size cannot buy.

The subtle part is that we cannot always greedily use the richest customer for
the current same-size shoe. For example:

```text
shoes:
1 1
2 2

customers:
1 1
2 1
```

The optimal answer is `3`: sell the size `1` shoe to the poorer customer, and
sell the size `2` shoe to the richer customer. If we greedily sell the size `1`
shoe to the richer customer, the size `2` shoe cannot be sold. So the DP must
remember which of the top two customers of the current foot size has already
been used.

### DP state

After processing coordinate `arr[i]`, `dp[mask]` is the maximum profit where
`mask` describes which of the top two customers with foot size `arr[i]` were
used by the shoe of the same size `arr[i]`.

There are only four masks:

- `0`: no top customer of this foot size was used as a same-size buyer.
- `1`: the richest customer was used.
- `2`: the second richest customer was used.
- `3`: impossible in this problem's transition, but keeping four states makes
  the mask logic simple.

Only same-size sales need to be remembered in the mask. A customer with foot
size `x` can still affect the next coordinate only when the next shoe has size
`x + 1`. Once we move beyond `x + 1`, customers of foot size `x` are no longer
relevant.

### Transitions

For each coordinate `arr[i]`, start with all `ndp` values as impossible.

From every previous state `mask`:

1. Skip the shoe at this size, or if there is no shoe at this coordinate:
   update `ndp[0]` with the same profit. The new mask becomes `0` because we
   are now describing usage of customers with foot size `arr[i]`, not the
   previous foot size.

2. If there is a shoe of size `arr[i]`, sell it to one of the top two customers
   with foot size `arr[i]`, if that customer has enough money. This creates a
   new mask `1 << j`, because customer `j` of the current foot size has now
   been used and must not be reused by the next size.

3. If `arr[i-1] + 1 = arr[i]`, sell the current shoe to one of the top two
   customers with foot size `arr[i-1]`. This is allowed only if that customer's
   bit is not already set in the previous `mask`. The new mask is `0`, because
   this sale uses a previous-foot-size customer, not a current-foot-size
   customer.

This handles the local conflict exactly: a foot-size group can only conflict
between two adjacent shoe sizes.

### Reconstruction

The profit DP alone is not enough to output assignments. If reconstruction only
walks backward by matching total profit values, it can choose a transition that
has the same value but was not the transition actually used by the optimum.
That can lead to missing sales or reusing a customer.

So every successful relaxation stores a parent:

- previous mask,
- customer id used by this transition, or `-1` for skip,
- shoe id used by this transition, or `-1` for skip.

After all coordinates are processed, choose the mask with maximum profit and
walk backward through the stored parents. Whenever the parent contains a
customer and shoe, append that assignment. This reconstructs the exact DP path,
not just any path with the same total.

The complexity is `O((n + m) log(n + m))` for sorting/compression plus constant
work per coordinate, because each coordinate has only four masks and each foot
size keeps at most two customers.
