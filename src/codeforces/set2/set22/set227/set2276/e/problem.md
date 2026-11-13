# E. Good Substrings

[Problem link](https://codeforces.com/problemset/problem/2276/E)

**Contest:** [Educational Codeforces Round 195 (Rated for Div. 2)](https://codeforces.com/contest/2276)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

## Problem

You are given a string `s` of length `n` consisting of digits.

A substring of `s` is a contiguous segment of its characters. A substring `s_{l...r}` is bad when it contains `2025` as a subsequence and does not contain `2026` as a subsequence. Otherwise it is good.

For example, `2025`, `72707275`, and `20625` are bad, while `20256`, `2026`, `777`, and `1290723576` are good.

For every index `i` from `1` to `n`, count the good substrings that contain `i`: the number of pairs `(l, r)` such that `1 ≤ l ≤ i ≤ r ≤ n` and `s_{l...r}` is good.

## Constraints

- `1 ≤ t ≤ 10^4`
- `1 ≤ n ≤ 3 · 10^5`
- `s` has length `n` and contains only digits
- The sum of `n` over all test cases does not exceed `3 · 10^5`

## Input

The first line contains `t`, the number of test cases.

Each test case is two lines:

```text
n
s
```

## Output

For each test case, print `n` integers. The `i`-th integer is the number of good substrings that contain index `i`.

## Sample Input 1

```text
6
202516
```

## Sample Output 1

```text
4 8 10 10 9 6
```

## Sample Input 2

```text
5
20265
```

## Sample Output 2

```text
5 8 9 8 5
```

## Status

1. f(i)指包含i的子串, 在其中, 不包含2025, 或者同时包含2026
2. F(i) = i子串, 且包含2025
3. G(i) = i子串, 且同时包含 2025 & 2026
4. f(i) = (i + 1) * (n - i) - (F(i) - G(i))
5. F(i) = 对于l, 如果 s[l..i] 包含空, 那么r = 第一个2025的位置, (n - r)
6.       + 如果 s[l..i] 包含2, 那么就是 025的位置
7.       + ...            20,  25
8. 需要找到i前面最右边的 2025, 202, 20, 2 的位置, 和后面 5, 25, 025, 2025的位置
9. 那么这样子, 左右最多有5个位置需要考虑
10. 让后处理G, 类似的处理

## 题解

`ans[i]` 是下标 `i`（从 0 计）被多少个好子串盖住。打印出来的第 `i+1` 个数就是它，对应题目里从 1 计的位置 `i+1`。

字符串 `s` 只含数字。子串是连续的一段。一段是坏的，当且仅当它把 `2025` 当作子序列，同时不把 `2026` 当作子序列；其余子串都是好的。对每个位置，要数有多少好子串包含它。`solve` 用「盖住这个位置的全部子串」减去「盖住它的坏子串」。上面的式子是同一件事：好子串数 = `(i+1)*(n-i) - (F(i)-G(i))`，其中 `F-G` 就是「含 `2025` 且不含 `2026`」。代码没有分开算 `F` 和 `G`，而是直接累加这个坏覆盖。

固定左端点 `l` 时，`endOf(l, "2025")` 得到的 `end25` 是从 `l` 出发、凑齐子序列 `2025` 的最左右端点；凑不齐就是 `n`。`end26` 对 `2026` 同样定义。子串 `s[l..r]` 含有这条子序列，正好是 `r` 不小于这个最左右端点。因此坏的右端点是那些已经凑齐 `2025`、还没凑齐 `2026` 的 `r`，也就是半开区间 `[end25, end26)`。`end25 >= end26` 时这个区间是空的：要么根本没有 `2025`，要么一旦 `2025` 进了子串，`2026` 也已经在里面。

样本 `s = 202516`，`n = 6`：

| 0-based `i` | 0 | 1 | 2 | 3 | 4 | 5 |
| --- | --- | --- | --- | --- | --- | --- |
| 字符 | 2 | 0 | 2 | 5 | 1 | 6 |
| 题目位置 | 1 | 2 | 3 | 4 | 5 | 6 |

只有 `l = 0` 产生坏区间。`2025` 依次落在下标 0、1、2、3，所以 `end25 = 3`。`2026` 的前三个字符仍是 0、1、2，最后的 `6` 在下标 5，所以 `end26 = 5`。坏右端点是 `r = 3, 4`，两条坏子串是 `s[0..3] = "2025"` 和 `s[0..4] = "20251"`。整串 `"202516"` 含有 `2026`，是好的，不计入坏覆盖。

`l = 1` 时，第一个 `2` 在下标 2，它后面没有 `0`，`end25` 和 `end26` 都是 6。`l = 2, 3, 4, 5` 同样凑不出长度为 4 的子序列。这些左端点没有坏右端点。

两条坏子串的覆盖：

- `[0,3]` 盖住 0、1、2、3
- `[0,4]` 盖住 0、1、2、3、4

坏覆盖因此是 `2 2 2 2 1 0`。

盖住下标 `i` 的子串，左端点有 `i+1` 种（`0..i`），右端点有 `n-i` 种（`i..n-1`），总数是 `(i+1)*(n-i)`：`6 10 12 12 10 6`。相减得到样本：

| `i` | 0 | 1 | 2 | 3 | 4 | 5 |
| --- | --- | --- | --- | --- | --- | --- |
| 总数 | 6 | 10 | 12 | 12 | 10 | 6 |
| 坏覆盖 | 2 | 2 | 2 | 2 | 1 | 0 |
| `ans` | 4 | 8 | 10 | 10 | 9 | 6 |

位置 1 的 6 个子串都从下标 0 开始：`"2" "20" "202" "2025" "20251" "202516"`，其中两条是坏的，剩下 4。位置 6 的 6 个子串都以下标 5 结尾；整串含 `2026`，较短的几条不含 `2025`，坏覆盖是 0，答案是 6。

`next` 从右往左填。`at[d]` 一开始全是 `n`。扫到 `i` 时把 `at[s[i]-'0']` 改成 `i`，再整份抄进 `next[i]`。于是 `next[i][d]` 是 `≥ i` 的第一个数字 `d`，没有就是 `n`。

```go
// next[i][d] is the first index >= i where digit d occurs, or n.
next := make([][10]int, n)
var at [10]int
for d := range at {
    at[d] = n
}
for i := n - 1; i >= 0; i-- {
    at[s[i]-'0'] = i
    next[i] = at
}
```

样本里和这题有关的几列（`6` 表示不存在）：

| `i` | `next` 0 | `next` 2 | `next` 5 | `next` 6 |
| --- | --- | --- | --- | --- |
| 0 | 1 | 0 | 3 | 5 |
| 1 | 1 | 2 | 3 | 5 |
| 2 | 6 | 2 | 3 | 5 |
| 3 | 6 | 6 | 3 | 5 |
| 4 | 6 | 6 | 6 | 5 |
| 5 | 6 | 6 | 6 | 5 |

`endOf` 从 `r = l` 开始，对模式的每个字符取 `next[r][digit]`，匹配成功后 `r++`，下一个字符必须严格在这一位的右边。四个位置因此互不相同，模式里的两个 `2` 不会用同一个下标。某一步 `r` 已经越出字符串，或者 `next` 返回 `n`，函数返回 `n`。

```go
endOf := func(l int, pat string) int {
    r := l
    for j := range len(pat) {
        if r >= n {
            return n
        }
        r = next[r][pat[j]-'0']
        if r >= n {
            return n
        }
        r++
    }
    return r - 1
}
```

每个左端点只做两次长度为 4 的查找，所以这一步是 O(n)。

差分数组记的是坏覆盖。扫完下标 `i` 之后，`bad` 是盖住 `i` 的坏子串个数，`step` 是再往右一格时 `bad` 的增量。

- `diff[p] += v`：从 `p` 起，`bad` 永久加上 `v`。
- `slope[q] += s`：从 `q` 起，每一格都比前一格多加 `s`。`step` 加上 `s` 后会一直留着，而每一格都执行 `bad += step`。

设 `L = end25`，`R = end26`，`C = R-L`。这一组坏子串的覆盖是：`l ≤ i ≤ L` 时为 `C`（每个坏右端点都 `≥ L ≥ i`）；`L < i < R` 时为 `R-i`（右端点要满足 `i ≤ r < R`）；`i ≥ R` 时为 0。三次 O(1) 更新做出这个形状：`diff[l] += C` 让 `≥ l` 先加上 `C`；`slope[L+1]--` 让 `≥ L+1` 再减去 `i-L`，所以 `L` 仍是 `C`，`L+1` 变成 `C-1`，到下标 `R` 正好减成 0；`R < n` 时 `slope[R+1]++` 把每格多出来的 `-1` 补回去，`R` 之后保持 0。多个左端点加进同一组数组，最后一次从左到右的扫描把它们叠加起来。

```go
for l := range n {
    end25 := endOf(l, "2025")
    end26 := endOf(l, "2026")
    if end25 >= end26 {
        continue
    }
    diff[l] += int64(end26 - end25)
    slope[end25+1]--
    if end26 < n {
        slope[end26+1]++
    }
}

ans := make([]int64, n)
var bad, step int64
for i := range n {
    bad += diff[i]
    step += slope[i]
    bad += step
    // (i+1)*(n-i) substrings contain index i (0-based).
    ans[i] = int64(i+1)*int64(n-i) - bad
}
```

样本这一组是 `l = 0, L = 3, R = 5, C = 2`，所以 `diff[0] = 2`，`slope[4] = -1`，`slope[6] = 1`。扫描是：`i = 0..3` 时 `bad` 停在 2；`i = 4` 时 `step` 变成 -1，`bad` 变成 1；`i = 5` 时 `step` 仍是 -1，`bad` 变成 0。`slope[6]` 写在数组最后一格，循环只读到 `i = 5`，这一格不会被用到。下标 5 已经是 0，后面也没有位置了。

这些行都在 `solution.go`：`next` 是第 33–42 行，`endOf` 是第 49–62 行，每个左端点的差分更新是第 63–74 行，把差分还原成 `ans` 是第 76–84 行。`solve` 整体是第 31–86 行。

每个测试是 O(n) 时间和 O(n) 额外内存。`t ≤ 10^4`，所有测试的 `n` 之和不超过 `3·10^5`，总工作量就是这个量级；模式长度 4 和 10 个数字都是常数。`n = 3·10^5` 时 `next` 大约 24MB，加上两条 `int64` 差分，远小于 512MB。单个答案最大大约是 `(n/2)^2 ≈ 2.25·10^10`，放得进 `int64`，所以 `diff`、`slope`、`bad`、`step` 和返回值都是 `int64`。

字符串里没有 `6` 时，`end26` 一直是 `n`。只要 `end25 < n`，坏右端点就是 `[end25, n)`，每条已经凑齐 `2025` 的子串都是坏的，代码也不会去写 `slope[n+1]`。`"2025"` 的唯一坏子串是整串，四个位置的坏覆盖都是 1，答案是 `3 5 5 3`。`2026` 不比 `2025` 更晚完成时，第 66–68 行直接跳过；两个都凑不齐时两端都是 `n`，同样跳过。`"20265"` 就是这种重叠：`end26 = 3`，`end25 = 4`，`6` 在 `5` 左边，任何含 `2025` 的子串都已经含 `2026`。没有任何坏子串，答案就是总数 `5 8 9 8 5`，和第二组样例一致。`n < 4` 时长度为 4 的子序列无法完成，答案同样只剩 `(i+1)*(n-i)`。若 `6` 后面还有字符，例如 `"202560"`，`end25 = 3`，`end26 = 4`，`slope[5]++` 会在扫描里被读到，把最后一个 `0` 的坏覆盖钉在 0。`end26 = n-1` 时写的是 `slope[n]`，扫描用不到它，因为坏覆盖在下标 `end26` 已经降到 0。
