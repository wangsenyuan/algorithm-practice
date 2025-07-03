# String Diversity

## Problem Description

String diversity is the number of symbols that occur in the string at least once. Diversity of `s` will be denoted by `d(s)`. 

**Examples:**
- `d("aaa") = 1`
- `d("abacaba") = 3`

Given a string `s`, consisting of lowercase Latin letters. Consider all its substrings. Obviously, any substring diversity is a number from 1 to `d(s)`. Find statistics about substrings diversity: for each `k` from 1 to `d(s)`, find how many substrings of `s` has a diversity of exactly `k`.

## Input

The input consists of a single line containing `s`. It contains only lowercase Latin letters, the length of `s` is from 1 to 3·10^5.

## Output

Print to the first line the value `d(s)`. Print sequence `t1, t2, ..., td(s)` to the following lines, where `ti` is the number of substrings of `s` having diversity of exactly `i`.

## ideas
1. 考虑一个字符x, 它的贡献 = (i - l[x]) * (n - i)