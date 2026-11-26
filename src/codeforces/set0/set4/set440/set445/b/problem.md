# B. DZY Loves Chemistry

[Problem link](https://codeforces.com/problemset/problem/445/B)

**Contest:** [Codeforces Round 254 (Div. 2)](https://codeforces.com/contest/445)

time limit per test: 1 second

memory limit per test: 256 megabytes

## Problem

DZY has `n` chemicals, and `m` pairs of them react.

He pours the chemicals into a test tube one by one in any order.
The danger of an empty test tube is `1`. When a chemical is poured, if at least
one chemical already in the test tube reacts with it, the danger is multiplied
by `2`; otherwise it remains unchanged.

Find the maximum possible danger after pouring all chemicals.

## Input

The first line contains two space-separated integers `n` and `m`.

Each of the next `m` lines contains two integers `x_i` and `y_i`
(`1 <= x_i < y_i <= n`). These chemicals react with each other.
Each pair appears at most once.

## Output

Print one integer: the maximum possible danger.

## Sample

```text
Input
1 0

Output
1
```

```text
Input
2 1
1 2

Output
2
```

```text
Input
3 2
1 2
2 3

Output
4
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
