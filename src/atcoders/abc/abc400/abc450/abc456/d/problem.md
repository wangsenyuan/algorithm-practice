# D - Not Adjacent 2 (ABC456)

**Contest:** [ABC456](https://atcoder.jp/contests/abc456) — AtCoder Beginner Contest 456  
**Task:** [https://atcoder.jp/contests/abc456/tasks/abc456_d](https://atcoder.jp/contests/abc456/tasks/abc456_d)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 400 points

## Problem Statement

You are given a string `S` consisting of `a`, `b`, and `c`.

Find the number of non-empty subsequences of `S` in which no two adjacent
characters are the same, modulo `998244353`.

Two subsequences are considered distinct if they are taken from different
positions, even if they are identical as strings.

A subsequence of `S` is a string obtained by removing zero or more characters
from `S` and concatenating the remaining characters in their original order. For
example, `ab` and `ac` are subsequences of `abc`, but `ca` and `bb` are not
subsequences of `abc`.

## Constraints

- `S` is a string of length between `1` and `3 * 10^5`, inclusive.
- `S` consists only of `a`, `b`, and `c`.

## Input

The input is given from Standard Input in the following format:

```text
S
```

## Output

Output the answer.

## Sample Input 1

```text
abbc
```

## Sample Output 1

```text
11
```

The subsequences in which no two adjacent characters are the same are the
following `11`:

- `a` (the 1st character of `S`)
- `b` (the 2nd character of `S`)
- `b` (the 3rd character of `S`)
- `c` (the 4th character of `S`)
- `ab` (the 1st and 2nd characters of `S`)
- `ab` (the 1st and 3rd characters of `S`)
- `ac` (the 1st and 4th characters of `S`)
- `bc` (the 2nd and 4th characters of `S`)
- `bc` (the 3rd and 4th characters of `S`)
- `abc` (the 1st, 2nd, and 4th characters of `S`)
- `abc` (the 1st, 3rd, and 4th characters of `S`)

Note that two subsequences are considered distinct if they are taken from
different positions, even if they are identical as strings.

## Sample Input 2

```text
cabcabcbcaccacbcbcaabacbacaabccacbccbcacbacbacabcacabcaccaaaaabababcbabacaccabbcacbcbcbcababcbcbabca
```

## Sample Output 2

```text
378217423
```

Output the count modulo `998244353`.
