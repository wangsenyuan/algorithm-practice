# C - 1122 Substring 2 (ABC433)

**Contest:** [ABC433](https://atcoder.jp/contests/abc433) — AtCoder Beginner Contest 433  
**Task:** [https://atcoder.jp/contests/abc433/tasks/abc433_c](https://atcoder.jp/contests/abc433/tasks/abc433_c)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 300 points

## Problem Statement

You are given a string `S` consisting of digits.

A string `T` is called a **1122-string** if it satisfies all of the following
(the same definition as in Problem F):

- `T` is a non-empty string consisting of digits
- `|T|` is even
- The first `|T|/2` characters of `T` are all the same digit
- The last `|T|/2` characters of `T` are all the same digit
- Adding `1` to the first digit of `T` gives the last digit of `T`

For example, `1122`, `01`, and `444555` are 1122-strings, but `1222` and `90`
are not.

Find the number of substrings of `S` that are 1122-strings.

Count substrings by position: two substrings are different if they start or end
at different indices, even if they are equal as strings.

## Constraints

- `S` is a string of digits with length between `1` and `10^6`, inclusive

## Input

The input is given from Standard Input in the following format:

```text
S
```

## Output

Output the number of non-empty substrings of `S` that are 1122-strings.

## Sample Input 1

```text
1122
```

## Sample Output 1

```text
2
```

The valid substrings are:

- `12` from positions 2–3
- `1122` from positions 1–4

## Sample Input 2

```text
7788788
```

## Sample Output 2

```text
3
```

Substrings with the same text but different positions are counted separately.

## Sample Input 3

```text
2025
```

## Sample Output 3

```text
0
```

## Sample Input 4

```text
1112222334445556555
```

## Sample Output 4

```text
11
```
