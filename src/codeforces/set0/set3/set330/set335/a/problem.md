# A. Banana

[Problem link](https://codeforces.com/problemset/problem/335/A)

**Contest:** [MemSQL start[c]up Round 2 - online version](https://codeforces.com/contest/335)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Piegirl wants to build a string `s` from stickers. Every sheet contains exactly `n` stickers, each printed with one lowercase letter, so a sheet is a string of length `n`. She may buy any number of sheets, but every sheet must be the same string. After buying them, she may use any subset of the stickers and arrange them in any order.

Print the minimum number of sheets needed, and one sheet that achieves it. If `s` cannot be formed, print `-1`.

## Constraints

- `1 ≤ |s| ≤ 1000`
- `s` consists of lowercase English letters
- `1 ≤ n ≤ 1000`

## Input

```text
s
n
```

## Output

Print the minimum number of sheets, then a string of `n` lowercase letters describing one optimal sheet. If no sheet works, print `-1`.

## Sample Input 1

```text
banana
4
```

## Sample Output 1

```text
2
baan
```

## Sample Input 2

```text
banana
3
```

## Sample Output 2

```text
3
nab
```

Three sheets of `nab` supply `banana`.

## Sample Input 3

```text
banana
2
```

## Sample Output 3

```text
-1
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO. Any sheet that achieves the minimum number of sheets is accepted.
