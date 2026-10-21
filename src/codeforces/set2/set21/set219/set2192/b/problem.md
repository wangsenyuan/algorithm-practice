# B. Flipping Binary String

[Problem link](https://codeforces.com/problemset/problem/2192/B)

**Contest:** [Codeforces Round 1081 (Div. 2)](https://codeforces.com/contest/2192)

time limit per test: 1.5 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

You are given a binary string `s` of length `n`. In one operation, choose an index `i`
(`1 ≤ i ≤ n`) and flip every bit except `s[i]`. Each index may be chosen at most once.
You may perform any number of operations.

Make every bit of `s` equal to `0`, or report that it is impossible. The number of
operations does not need to be minimized.

## Constraints

- `1 ≤ t ≤ 10^4`
- `1 ≤ n ≤ 2 · 10^5`
- `s` is a binary string of length `n`
- The sum of `n` over all test cases does not exceed `2 · 10^5`

## Input

The first line contains `t` — the number of test cases.

Each test case:

```text
n
s
```

## Output

For each test case, print `-1` if it is impossible.

Otherwise print two lines: the number of operations `x` (`0 ≤ x ≤ n`), then the `x`
chosen indices in order. Each index appears at most once. Any valid sequence is
accepted. If `x = 0`, the second line is empty.

## Example

### Input

```text
4
3
101
3
100
4
0000
4
1010
```

### Output

```text
1
2
-1
0

2
1 3
```

## Note

In the first test case, choosing index `2` flips indices `1` and `3`, so `101`
becomes `000`.

In the second test case, `100` cannot be turned into `000`.

## Solution

选下标 `i` 会把除 `i` 以外的每一位都翻转。异或与顺序无关，所以只需要决定操作集合 `S`，且每个下标最多进 `S` 一次。

位置 `j` 被翻转的次数：`j ∉ S` 时是 `|S|`，`j ∈ S` 时是 `|S| - 1`。设 `p = |S| mod 2`，最后要变成 `0`，于是

- `j ∉ S` 时，原来的 `s[j]` 必须等于 `p`
- `j ∈ S` 时，原来的 `s[j]` 必须等于 `1 - p`

因此 `S` 里的位全都相同，设为 `b`；`S` 外面的位全都是 `1 - b`。也就是说 `S` 必须恰好是「所有等于 `b` 的位置」，并且 `|S|` 的奇偶性等于 `1 - b`。

只有两个候选：

| 操作谁 | 条件 |
|---|---|
| 每一个 `1` | `1` 的个数是偶数 |
| 每一个 `0` | `0` 的个数是奇数 |

全是 `0` 时什么都不用做。这覆盖了「`0` 的个数是偶数」时第二条规则不成立的情况。

两条都不成立时无解：`1` 的个数是奇数，并且 `0` 的个数是偶数。这正好是 `n` 为奇数且 `1` 的个数也是奇数。两条都成立时任意一条都可以；实现里取个数更少的那一组。

### 样例

`101` 有两个 `1`、一个 `0`。两个 `1` 是偶数，一个 `0` 是奇数，两条都合法，更短的是操作那个 `0`（下标 `2`）：`101 → 000`。

`100` 有一个 `1`、两个 `0`，两条都不成立，答案是 `-1`。

`1010` 有两个 `1`、两个 `0`。只有「操作全部 `1`」合法，下标 `1 3`。先操作 `3`：`1010 → 0111`；再操作 `1`：`0111 → 0000`。

### 复杂度

每个测试扫一遍字符串，总长度不超过 `2 · 10^5`。
