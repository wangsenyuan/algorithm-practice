# J. The Ultimate Wine Tasting Event

[Problem link](https://codeforces.com/problemset/problem/2068/J)

time limit per test: 2 seconds

memory limit per test: 1024 megabytes

input: standard input

output: standard output

Rumors of the excellence of Gabriella's wine tasting events have toured the world and made
it to the headlines of prestigious wine magazines. Now, she has been asked to organize an
event at the EUC 2025!

This time she selected `2n` bottles of wine, of which exactly `n` are white wine, and
exactly `n` are red wine. She arranged them in a line as usual, in a predetermined order
described by a string `s` of length `2n`: for `1 <= i <= 2n`, the `i`-th bottle from the
left is white wine if `s_i = W` and red wine if `s_i = R`.

Consider a way of dividing the `2n` bottles into two disjoint subsets, each containing
`n` bottles. Then, for every `1 <= i <= n`, swap the `i`-th bottle in the first subset
(from the left) and the `i`-th bottle of the second subset (also from the left). Is it
possible to choose the subsets so that, after this operation is done exactly once, the
white wines occupy the first `n` positions?

## Input

The first line contains an integer `t` (`1 <= t <= 500`) — the number of test cases.

The first line of each test case contains an integer `n` (`1 <= n <= 100`) — where `2n`
is the total number of bottles.

The second line of each test case contains a string `s` of length `2n`, describing the
bottle arrangement — the `i`-th character of `s` (`1 <= i <= 2n`) is `W` for white wine
and `R` for red wine.

It is guaranteed that `s` contains exactly `n` `W`'s and `n` `R`'s.

## Output

For each test case, print `YES` if it is possible to divide the bottles as explained in
the statement. Otherwise, print `NO`.

## Example

### Input

```text
3
4
WRRWWWRR
1
WR
20
WWWWRRWRRRRRWRRWRWRRWRRWWWWWWWRWWRWWRRRR
```

### Output

```text
YES
NO
YES
```

## Note

In the first test case, one subset can be bottles at positions `1, 2, 3, 7` (`W, R, R, R`)
and the other at positions `4, 5, 6, 8` (`W, W, W, R`). After swapping pairs `(1,4)`,
`(2,5)`, `(3,6)`, `(7,8)`, whites occupy the first 4 positions.

In the second test case, the only partition swaps the two bottles and yields `RW`, so there
is no solution.

### ideas
1. 每一个位置,都必须参加一次swap, 结果是W在前, R在后
2. dp[i][j] 表示处理了前(i + j)个后, 其他第一个集合有i个元素, 第二个集合有j个元素, 是否可行