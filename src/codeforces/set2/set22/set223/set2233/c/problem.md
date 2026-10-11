# C. Cost of a Bracket Sequence

[Problem link](https://codeforces.com/problemset/problem/2233/C)

**Contest:** [Educational Codeforces Round 191 (Rated for Div. 2)](https://codeforces.com/contest/2233)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

The cost of a bracket string is the length of its longest subsequence that is a regular bracket sequence.

A sequence `a` is a subsequence of `b` if `a` can be obtained by deleting zero or more characters from `b`.

A bracket sequence is regular if inserting `+` and `1` can turn it into a correct arithmetic expression. For example, `(())()`, `()`, and `(()(()))` are regular, while `)(`, `(()`, and `(()))(` are not.

You are given a bracket string `s` and an integer `k`. Delete at most `k` characters so that the cost of the remaining string is minimized.

## Input

The first line contains `t` (`1 <= t <= 10^3`), the number of test cases.

Each test case is two lines:

- `n` and `k` (`1 <= n <= 5000`, `0 <= k <= n`)
- a string `s` of length `n` consisting of `(` and `)`

The sum of `n` over all test cases does not exceed `5000`.

## Output

For each test case, print a binary string of length `n`. The `i`-th character is `1` if `s[i]` is deleted and `0` otherwise.

The number of ones must be at most `k`, and the remaining string must have the minimum possible cost. If several answers exist, print any of them.

## Example

### Input

```text
10
2 1
)(
2 0
()
4 1
(())
4 1
())(
5 1
((())
6 2
()()()
6 2
(()())
6 2
())(()
7 3
(()((()
10 3
(()())())(
```

### Output

```text
00
00
1000
1000
00010
101000
001001
100001
1100001
0101001000
```

In the first test case the cost is already `0`, so deleting nothing is optimal.

In the third test case one deletion cannot reach cost `0`, but deleting any one character leaves cost `2`.

## ideas
1. 感觉上, 存在, 只删除一种, 左括号或者有括号, 从而得到最小cost的方案?
2. 首先在左端, 肯定删除左括号更优, 假设不是这样; 在右端删除右括号; 
3. 所以, 确实存在同时删除左右括号的情况; 
4. dp[i][j] 表示在前i个位置删除j个左括号时的最优解(这个时候右边删除了k-j个右括号)
5. dp[i][j] = 如果当前是右括号 dp[i-1][j] + 1 如果左边还有扩括号 else 0
6.    如果s[i]是左括号 dp[i-1][j-1] min dp[i-1][j] + ? 右边是否有 > k - j 个右括号
7.    

## Solution

The cost of a bracket string is twice the number of pairs in a left-to-right matching: every `)` claims one unmatched `(`. Deleting a character never creates a pair, so an optimal answer may spend the whole budget of `k` deletions.

It is enough to spend that budget on a prefix of the `(` positions and a suffix of the `)` positions. Suppose a maximum matching leaves out the leftmost `(`. That bracket can replace the opening bracket of any pair to its right, so the matching size does not drop. Deleting any other `(` is therefore no better than deleting this one. The same exchange applies to the rightmost `)`.

Let `k' = min(k, number of '(', number of ')')`. `solve` tries each split `x = 0, 1, …, k'`: delete the `x` leftmost `(` and the `k' - x` rightmost `)`. The mask with the smallest matching cost is the answer. The cap at `k'` is safe: once every `(` or every `)` is gone, the cost is already `0`.

Each of the `k' + 1` candidates is scored in `O(n)`. One test can have `n = 5000`, so the cost is `O(n^2)` and fits the limit. The sum of `n` over tests is `5000`.