# C. Levko and Strings

[Problem link](https://codeforces.com/problemset/problem/360/C)

**Contest:** [Codeforces Round 210 (Div. 1)](https://codeforces.com/contest/360)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

`s` is a string of length `n` made of lowercase English letters. For every string `t` of the same length, its beauty relative to `s` is the number of pairs `(i, j)` with `1 ≤ i ≤ j ≤ n` such that the substring `t[i..j]` is lexicographically larger than `s[i..j]`.

Two equal-length strings are compared at the first position where they differ, and the one with the larger character there is larger. Identical substrings do not count.

Count the strings `t` whose beauty is exactly `k`. Print that count modulo `1000000007` (`10^9 + 7`).

## Constraints

- `1 ≤ n ≤ 2000`
- `0 ≤ k ≤ 2000`
- `s` has length `n` and contains only lowercase English letters

## Input

```text
n k
s
```

## Output

Print the number of strings modulo `1000000007`.

## Sample Input 1

```text
2 2
yz
```

## Sample Output 1

```text
26
```

## Sample Input 2

```text
2 3
yx
```

## Sample Output 2

```text
2
```

## Sample Input 3

```text
4 7
abcd
```

## Sample Output 3

```text
21962
```

## Solution

子串 `t[l..r]` 比 `s[l..r]` 大，当且仅当第一处不同的位置 `i` 满足 `t[i] > s[i]`。beauty 只由这些严格更大的位置贡献。

按与 `s` 不同的位置把 `t` 切开。上一次不同在 `p`（`p = 0` 表示还没有），这一次在 `i`，中间 `t[p+1..i-1]` 必须等于 `s`。

- `t[i] < s[i]`：从 `p+1` 到 `i` 开头的子串都更小，beauty 不变。
- `t[i] > s[i]`：这些子串都更大。起点有 `i-p` 个，右端点可以是 `i..n` 的任意一个，共 `n-i+1` 个，所以增加 `(i-p)·(n-i+1)`。

`t[i]` 与 `s[i]` 不同之后，已经开头的子串都比较完了，后面的字符不再影响它们。

`dp[i][j]` 是填好前 `i` 位、第 `i` 位与 `s` 不同、已确定的 beauty 为 `j` 的方案数。`dp[0][0] = 1`。记 `R = n-i+1`。

- 更小的字符有 `s[i]-'a'` 种。上一次不同可以是任意 `p < i`，用前缀和一次转移。
- 更大的字符有 `'z'-s[i]` 种。长度 `L = i-p` 的代价是 `L·R`，只枚举 `L ≤ j/R`。

最后一次不同之后的后缀全部抄 `s`，beauty 不再增加，所以答案是所有 `dp[i][k]` 之和。整串等于 `s` 时 beauty 为 `0`，因此 `k = 0` 时再加 `1`。

### 样例 1

`s = yz`，`k = 2`。第 1 位只有 `z` 比 `y` 大，代价 `1·2 = 2`。后面再抄 `z` 得到 `zz`。第 1 位已经是 `z` 时，第 2 位改成 `a..y` 不会再增加 beauty，得到另外 25 个串。一共 `26`。

### 复杂度

更小的转移是 `O(nk)`。位置 `i` 上更大的转移枚举 `O(k^2 / (n-i+1))` 次，对所有位置求和是 `O(k^2 log n)`。`n, k ≤ 2000`。
