# E - A > B substring (ABC441)

**Contest:** [ABC441](https://atcoder.jp/contests/abc441) — AtCoder Beginner Contest 441  
**Task:** [https://atcoder.jp/contests/abc441/tasks/abc441_e](https://atcoder.jp/contests/abc441/tasks/abc441_e)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 450 points

## Problem Statement

You are given a string `S` of length `N` consisting of three kinds of characters:
`A`, `B`, and `C`.

There are `N(N+1)/2` non-empty contiguous substrings of `S`. Find how many of
them contain more `A`s than `B`s.

Note that two substrings are counted separately if they are taken from different
positions in `S`, even if they are equal as strings.

## What Is A Substring?

A substring of `S` is a string obtained by deleting zero or more characters from
the beginning and zero or more characters from the end of `S`.

For example, `AB` is a substring of `ABC`, but `AC` is not a substring of `ABC`.

## Constraints

- `1 <= N <= 5 * 10^5`
- `S` is a string of length `N` consisting of `A`, `B`, and `C`.
- `N` is an integer.

## Input

The input is given from Standard Input in the following format:

```text
N
S
```

## Output

Print the number of contiguous substrings of `S` that contain more `A`s than
`B`s.

## Sample Input 1

```text
10
ACBBCABCAB
```

## Sample Output 1

```text
8
```

The following eight substrings satisfy the condition:

- `A`: 1st to 1st character of `S`
- `AC`: 1st to 2nd character of `S`
- `CA`: 5th to 6th character of `S`
- `CABCA`: 5th to 9th character of `S`
- `A`: 6th to 6th character of `S`
- `ABCA`: 6th to 9th character of `S`
- `CA`: 8th to 9th character of `S`
- `A`: 9th to 9th character of `S`

All other substrings do not satisfy the condition, so print `8`.

Note that `A` and `CA` can be taken from multiple positions, but they are counted
separately if taken from different positions.

## Sample Input 2

```text
4
CCBC
```

## Sample Output 2

```text
0
```

There may be no substrings that satisfy the condition.

## Sample Input 3

```text
36
CABACBBBBBAABABACCBCABCCABAABABBCBAC
```

## Sample Output 3

```text
136
```


### ideas
1. 在区间[l..r]中间，要有更多cnt[a] > cnt[b]
2. cnt[a][r] - cnt[a][l-1] > cnt[b][r] - cnt[b][l-1]
3. cnt[a][r] - cnt[b][r] > cnt[a][l-1] - cnt[b][l-1]
4. 比较简单，区间查询就可以了（但是有负数，所以需要offset一下）