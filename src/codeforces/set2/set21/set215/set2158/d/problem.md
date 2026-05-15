# D. Palindrome Flipping (Codeforces 2158D)

**Limits:** 2 seconds per test, 256 MB  
**I/O:** standard input / standard output

Source: [https://codeforces.com/problemset/problem/2158/D](https://codeforces.com/problemset/problem/2158/D)

## Problem Statement

You are given two binary strings `s` and `t`, each of the same length `n`.

You may perform the following operation:

- Choose indices `l` and `r` (`1 <= l < r <= n`) such that the substring `s[l:r]` is a palindrome.
- Flip all bits in the substring `s[l:r]`.

The goal is to make `s` equal to `t` using at most `2n` operations, possibly zero.

Here:

- `s[l:r]` is the contiguous sequence of characters from index `l` to index `r`, inclusive.
- A string is a palindrome if it reads the same forwards and backwards.
- Flipping all bits means changing every `0` to `1` and every `1` to `0`.

For example, flipping `101` gives `010`.

## Input

Each test contains multiple test cases.

The first line contains an integer `T` (`1 <= T <= 5 * 10^3`) — the number of test cases.

For each test case:

- The first line contains an integer `n` (`4 <= n <= 100`) — the length of `s` and `t`.
- The next two lines contain the binary strings `s` and `t`.

It is guaranteed that the sum of `n^2` over all test cases does not exceed `5 * 10^5`.

## Output

For each test case:

- If it is impossible to make `s` equal to `t`, print `-1`.
- Otherwise, first print an integer `k` (`0 <= k <= 2n`) — the number of operations.
- Then print `k` lines, each containing two integers `l` and `r` (`1 <= l < r <= n`) — the operation indices.

For every printed operation, the current substring `s[l:r]` must be a palindrome at the moment the operation is performed.

## Example

### Input

```text
3
5
01011
10000
7
1010101
0101010
4
0010
0010
```

### Output

```text
2
1 3
3 5
1
1 7
0
```

## Note

For the first test case:

- Initially, `s = 01011` and `t = 10000`.
- Choose `l = 1` and `r = 3`. This is valid because `s[1:3] = 010` is a palindrome. After flipping it, `s = 10111`.
- Choose `l = 3` and `r = 5`. This is valid because `s[3:5] = 111` is a palindrome. After flipping it, `s = 10000`.
- Now `s = t`.

For the second test case:

- Initially, `s = 1010101` and `t = 0101010`.
- Choose `l = 1` and `r = 7`. This is valid because `s[1:7] = 1010101` is a palindrome. After flipping it, `s = 0101010`.

For the third test case, `s` is already equal to `t`, so no operations are needed.


### ideas
1. 如果选择的长度为2， 考虑10, 01, 11, 那么a[i] = (s[i] != t[i]) 的 xor的结果是不会变的
2. 如果选择的长度为3， 考虑101 => 010, 111 => 000, 奇偶性发生了变化
3. 如果长度等于2，那么直接检查就可以了
4. 如果长度>2, 如果s[i] != t[i], 但是s[i] = s[i+1] 也可以直接翻转
5. 但是如果s[i] != s[i+1], 那么s[i+2] = s[i], 那么也可以翻转，
6. 但是如果s[i+2] = s[i+1], 那么必须先翻转s[i+1:i+2], 再一起翻转。
7. 这样做似乎可以保证次数（i作为起点，最多参与了两次翻转）
8. 假设到了最后阶段， s[n] != t[n], 它是最后的位置
9. 那么选择(n-2, n-1), 一起翻转，然后再翻转一次
10. 
