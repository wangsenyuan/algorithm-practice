# Problem Statement

You are given `n` strings `s1, s2, ..., sn` consisting of characters `0` and `1`. `m` operations are performed, on each of them you concatenate two existing strings into a new one. On the `i-th` operation the concatenation `sa[i] + sb[i]` is saved into a new string `s[n + i]` (the operations are numbered starting from 1). 

After each operation you need to find the maximum positive integer `k` such that all possible strings consisting of `0` and `1` of length `k` (there are `2^k` such strings) are substrings of the new string. If there is no such `k`, print `0`.

## Input

The first line contains single integer `n` (1 ≤ n ≤ 100) — the number of strings. The next `n` lines contain strings `s1, s2, ..., sn` (1 ≤ |si| ≤ 100), one per line. The total length of strings is not greater than 100.

The next line contains single integer `m` (1 ≤ m ≤ 100) — the number of operations. `m` lines follow, each of them contains two integers `ai` and `bi` (1 ≤ ai, bi ≤ n + i - 1) — the number of strings that are concatenated to form `s[n + i]`.

## Output

Print `m` lines, each should contain one integer — the answer to the question after the corresponding operation.

## Example

### Input
```
5
01
10
101
11111
0
3
1 2
6 5
4 4
```

### Output
```
1
2
0
```

## Note

- On the first operation, a new string `"0110"` is created. For `k = 1` the two possible binary strings of length `k` are `"0"` and `"1"`, they are substrings of the new string. For `k = 2` and greater there exist strings of length `k` that do not appear in this string (for `k = 2` such string is `"00"`). So the answer is `1`.

- On the second operation the string `"01100"` is created. Now all strings of length `k = 2` are present.

- On the third operation the string `"1111111111"` is created. There is no zero, so the answer is `0`.

### ideas
1. 也就是要找到最大的k，对于这个k，所有$pow(2, k)$ 中 01...01 (长度为k)都存在
2. 如果 $s_c$ = $s_a$ + $s_b$
3. 如果 f(a) = ka, f(b) = kb, 那么 f(c) >= max(ka, kb)
4. 大于的部分，要在连接处找
5. 考虑 0001 + 0111
6. 00, 01, 10, 11 所以k等于2
7. dp[k]表示要等到k的全部子串，需要的最短长度
8. 要在所有k-1长度的子串前面增加0/1, dp[1] = 01, dp[2] = 0110
9. dp[3] = 000 001 011 010 110 111 101 100
10. 最多只能到9，那就简单了