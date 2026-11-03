# C. Fixing Typos

[Problem link](https://codeforces.com/problemset/problem/363/C)

**Contest:** [Codeforces Round 211 (Div. 2)](https://codeforces.com/contest/363)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Delete the fewest letters from a word `s` so the result has neither of these typos:

- three identical letters in a row, such as `helllo`
- two identical letters immediately followed by two identical letters, such as `helloo` or `wwaatt`

Letters may be deleted from either end or from the middle. Print any shortest word that can be obtained this way.

## Constraints

- `1 ≤ |s| ≤ 2 · 10^5`
- `s` consists of lowercase English letters

## Input

```text
s
```

## Output

Print one shortest typo-free word obtained by deleting letters from `s`.

## Sample Input 1

```text
helloo
```

## Sample Output 1

```text
hello
```

`heloo` is also valid.

## Sample Input 2

```text
woooooow
```

## Sample Output 2

```text
woow
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO. Any shortest typo-free word is accepted.
