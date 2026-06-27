# D. Fullmetal Bitchemist

[Problem link](https://codeforces.com/problemset/problem/2237/D)

**Contest:** [Order Capital Round 2 (Codeforces Round 1104, Div. 1 + Div. 2)](https://codeforces.com/contest/2237)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

A binary string consists only of characters `0` and `1`.

A binary string `t` is **beautiful** if it can be reduced to length exactly `1` by repeatedly choosing two equal adjacent characters, removing both, and inserting the opposite character in their place.

You are given a binary string `s`. Count the number of non-empty beautiful substrings of `s`.

A string `a` is a substring of `b` if `a` can be obtained by deleting zero or more characters from the beginning and end of `b`.

## Input

The first line contains an integer `t` (`1 <= t <= 10^4`) — the number of test cases.

Each test case contains:

- one line with an integer `n` (`1 <= n <= 10^6`) — the length of `s`;
- one line with a binary string `s` of length `n`.

It is guaranteed that the sum of `n` over all test cases does not exceed `10^6`.

## Output

For each test case, print one integer — the number of beautiful substrings of `s`.

## Example

### Input

```text
10
1
0
2
01
5
01001
3
001
6
011110
9
010110110
12
010000101001
16
1010011010010110
20
11110101101101001110
30
000101100011111001111100000010
```

### Output

```text
1
2
10
5
15
30
47
81
139
316
```

### Note

- In the first test case, the only substring is `0`, which has length `1` and is beautiful.
- In the second test case, the beautiful substrings are `0` and `1`; `01` is not beautiful.
- In the third test case, there are `10` beautiful substrings; one optimal example is the full string `01001`.

### ideas
1. 看不出来重点~
2. 假设知道以i-1的贡献, dp[i-1][0/1]
3. 表示i-1 最后变成0/1的区间数 
4. 如果 s[i] = '0' => dp[i][0] = 1 
5.                   dp[i][1] = dp[i-1][0]
6. 如果 s[i] = '1' => dp[i][0] = dp[i-1][0]
7.                    dp[i][1] = 1
8. 