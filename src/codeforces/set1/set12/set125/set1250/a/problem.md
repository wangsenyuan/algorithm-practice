# A. Berstagram

[Problem link](https://codeforces.com/problemset/problem/1250/A)

**Contest:** [2019-2020 ICPC, NERC, Northern Eurasia Onsite](https://codeforces.com/contest/1250)

time limit per test: 3 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

## Problem

Polycarp published `n` posts. The feed starts as `1, 2, ..., n`, where position `1` is the highest and position `n` is the lowest.

Then `m` likes arrive. The `j`-th like is for post `a_j`. If that post is not already highest, it swaps with the post directly above it. If it is already highest, the feed does not change.

For each post, print the highest and lowest positions it ever occupies. Consider the feed before any likes, after each like, and after all likes. The highest position is the smaller index.

## Constraints

- `1 ≤ n ≤ 10^5`
- `1 ≤ m ≤ 4 · 10^5`
- `1 ≤ a_j ≤ n`

## Input

```text
n m
a_1 a_2 ... a_m
```

## Output

Print `n` lines. Line `i` contains the minimum and maximum positions of post `i`.

## Sample Input 1

```text
3 5
3 2 1 3 3
```

## Sample Output 1

```text
1 2
2 3
1 3
```

The feed is `[1, 2, 3]`, then `[1, 3, 2]`, `[1, 2, 3]`, `[1, 2, 3]`, `[1, 3, 2]`, and finally `[3, 1, 2]`.

## Sample Input 2

```text
10 6
7 3 5 7 3 6
```

## Sample Output 2

```text
1 2
2 3
1 3
4 7
4 5
6 7
5 7
8 8
9 9
10 10
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
