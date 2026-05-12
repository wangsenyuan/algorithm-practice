# Codeforces 2203C - Test Generator

https://codeforces.com/problemset/problem/2203/C

## Statement

You are developing a test generator. It takes two integers `s` and `m` as input.
You need to construct an array of non-negative integers:

```text
a = [a1, a2, ..., an]
```

such that:

```text
sum(a[i]) = s
```

and for each `i`:

```text
a[i] & m = a[i]
```

where `&` denotes the bitwise AND operator.

In other words, in each number `a[i]`, bits set to `1` may appear only in
positions where the corresponding bit of `m` is also `1`.

Determine whether at least one such array exists. If it exists, find the minimum
possible length `n`.

## Input

Each test contains multiple test cases.

The first line contains the number of test cases `t` (`1 <= t <= 10^4`).

Each test case consists of one line containing two integers `s` and `m`
(`1 <= s, m <= 10^18`) — the parameters of the generator.

## Output

For each test case, print one integer:

- if such an array does not exist, print `-1`;
- otherwise, print the minimum possible value of `n`, the array length.

## Sample

```text
6
13 5
13 3
13 6
1000000007 2776648
99999999999 1
998244353 1557287
```

```text
3
5
-1
-1
99999999999
642
```

## Note

For `s = 13`, `m = 5`, the answer is `3`, because one suitable array is:

```text
[5, 4, 4]
```

For `s = 13`, `m = 3`, the answer is `5`, because one suitable array is:

```text
[3, 3, 3, 3, 1]
```

For `s = 13`, `m = 6`, the answer is `-1`, because no suitable array exists.


### ideas
1. a & m = a => a 是m的一个子集
2. 考虑每一位的情况
3. 