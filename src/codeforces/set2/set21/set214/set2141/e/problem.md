# Problem

You are given a string `s` consisting of characters `0`, `1`, and/or `?`.

Call a binary string perfect if it can be split into two non-empty parts:

- a prefix `a`,
- a suffix `b`,

such that `a[i] >= b[i]` for every `i` from `1` to `min(|a|, |b|)`.

Your task is to count the number of ways to replace all question marks independently with `0` or `1` so that the resulting string is perfect.

Two ways are different if they differ in at least one position.

Print the answer modulo `998244353`.

## Input

The first line contains a single integer `t` (`1 <= t <= 10^4`) — the number of test cases.

Each test case contains one line with a string `s` (`2 <= |s| <= 2 * 10^5`) consisting of `0`, `1`, and/or `?`.

It is guaranteed that the sum of `|s|` over all test cases does not exceed `2 * 10^5`.

## Output

For each test case, print one integer — the number of valid replacements modulo `998244353`.

## Example

**Input**

```text
5
0?1
011
????
110?0
0?1?
```

**Output**

```text
1
0
15
2
3
```

## Note

In the first test case, string `001` is valid: it can be split into a prefix of length `1` and a suffix of length `2`.

In the fourth test case, two valid resulting strings are:

- `11010`, which can be split into prefix length `2` and suffix length `3`;
- `11000`, which can be split into prefix length `4` and suffix length `1`.

## ideas
1. 假设处理到位置i，把s分成了两部分
2. 这时候看，能够有多少种方案
3. 不考虑复杂性，就从头开始, 如果 a[0] > b[0]， 那么方案数 = pow(2, ? 的数量)
4. 如果a[0] = '?', 且 b[0] = '0' = pow(2, ? - 1) 继续
5. 分类讨论，处理到尾巴
6. 但是问题在于，这个复杂性是n*n的；
7. 那么也可以是 pow(2, ?) - (a < b)
8. a < b, a[0] < b[0], 或者 a[0...i] = b[0...i], 且 a[i+1] < b[i+1]
9. 那么这个时候，a[0...i] = b[0...i] 对于b来说就是前缀匹配了
10. 如果是 s[.] = 0/1, 那么就是直接匹配，如果是?, 看情况，（两个？贡献2）
11. dp[i] 表示到i为止, s[:i/2] = s[i/2:i] (i只能是偶数)时的计数
12. 如果s[i+1] > s[i/2+1] => dp[i] * (那些还没有确定的部分的贡献)
13. 都不需要kmp
14. 上面理解错了。不是按照i分割后计算，而是先分配？，如果这个分配可以保证 a >= b ， 不管在哪里分割，都是算一次
15. 