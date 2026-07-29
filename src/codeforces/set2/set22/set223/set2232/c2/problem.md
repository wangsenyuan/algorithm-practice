# C2. Seating Arrangement (Hard Version)

[Problem link](https://codeforces.com/problemset/problem/2232/C2)

**Contest:** [Codeforces contest 2232](https://codeforces.com/contest/2232)

## Problem

This is the hard version of the problem. Constraints on `n`, `x`, `s`, and `t`
are larger than in the easy version.

Alice has `x` tables with `s` seats each. Friends line up in a fixed order; for
each person Alice either seats them at a valid table or kicks them out. People
are processed in order.

Personalities:

- `I` (introvert) — must sit at an **empty** table
- `E` (extrovert) — must sit at a **non-empty** table
- `A` (ambivert) — may sit at any table

Maximize the number of seated friends.

## Constraints

- `1 <= t <= 10^4`
- `1 <= n, x, s <= 2 · 10^5`
- Sum of `n` over all test cases `<= 2 · 10^5`
- String `u` consists of `A`, `E`, `I` only

## Input

The first line contains `t`.

For each test case:

- The first line contains three integers `n`, `x`, and `s`
- The second line contains a string `u` of length `n`

## Output

For each test case, print one integer — the maximum number of people seated.

## Sample 1

```text
Input
5 2 2
EIAIE

Output
4
```

## Sample 2

```text
Input
20 5 5
AEIEEEEIEAAEIEEEEIEA

Output
20
```

## Sample 3

```text
Input
8 2 4
AAAAAIEE

Output
7
```

## Sample 4

```text
Input
8 4 2
AIEAEAAI

Output
7
```

## Sample 5

```text
Input
8 3 3
AIEAEAAI

Output
7
```

## Sample 6

```text
Input
4 2 2
IAEE

Output
4
```

## ideas
1. 每个桌子的第一个人, 必须是I或者是A, E不能是第一个
2. 那么对于I来说, 当他进入的时候,只要有空桌子, 就坐下来, 否则只能离开
3. 然后后面的人, 尽量的去占有人的桌子
4. 当出现一个E的时候, 这个时候, 只能调配A的人过来; 如果存在这样free的人, 就安排
