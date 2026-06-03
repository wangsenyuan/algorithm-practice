# F - All Included (ABC419)

**Contest:** [ABC419](https://atcoder.jp/contests/abc419) — AtCoder Beginner Contest 419  
**Task:** [https://atcoder.jp/contests/abc419/tasks/abc419_f](https://atcoder.jp/contests/abc419/tasks/abc419_f)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 550 points

## Problem Statement

You are given `N` lowercase English strings `S_1, S_2, ..., S_N` and an integer
`L`.

Find the number, modulo `998244353`, of length-`L` lowercase English strings that
contain all of `S_1, S_2, ..., S_N` as substrings.

A **substring** of `S` is a string obtained by deleting zero or more characters from
the beginning and zero or more characters from the end of `S`. For example, `ab`,
`bc`, and `bcd` are substrings of `abcd`, while `ac`, `dc`, and `e` are not
substrings of `abcd`.

## Constraints

- `1 <= N <= 8`
- `1 <= L <= 100`
- `N` and `L` are integers
- Each `S_i` is a string of length between `1` and `10`, inclusive, consisting of
  lowercase English letters
- `S_i != S_j` for `i != j`

## Input

The input is given from standard input in the following format:

```text
N L
S_1
S_2
:
S_N
```

## Output

Print the answer.

## Sample Input 1

```text
2 4
ab
c
```

## Sample Output 1

```text
153
```

Some strings that satisfy the condition are `abcz` and `cabc`. The string `acbd`
does not contain `ab` as a substring, so it does not satisfy the condition.

## Sample Input 2

```text
2 6
abc
cde
```

## Sample Output 2

```text
54
```

## Sample Input 3

```text
5 50
bbfogggj
zkbach
eedirhyc
ffgd
oemmswj
```

## Sample Output 3

```text
689020583
```


### ideas
1. 如果 s[i]被s[j]完全包含，那么可以把s[i]丢弃掉（因为结果中出现了s[j], 那么肯定也出现了s[i])
2. dp[i][mask][k]表示到i为止，已经完全包含了mask表示的集合，且最后一个是s[k]时的方案数
3. 现在加入字符s[w], 如果s[k]和s[w]有一段*公共*的部分，那么起点加入的长度，会不同
4. 但是有一个的问题，就是中间可以掺入一些随意的字符，会造成重复计算
5. dp[i][?] 这个状态，应该是和整个匹配有关系的
6. 用trie来定位状态？不对。
7. 比如对于abc, 再加入一个d，其实就切换到了cde, 如果是加入c, 就回到了c
8. 