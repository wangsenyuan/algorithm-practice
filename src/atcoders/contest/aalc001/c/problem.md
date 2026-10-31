# C - Binary Search on a Segment Tree

[Problem link](https://atcoder.jp/contests/aalc001/tasks/aalc001_c)

**Contest:** [AAL Contest 001: Let's use segtree!](https://atcoder.jp/contests/aalc001)

time limit: 2 sec

memory limit: 1024 MiB

score: 3 points

## Problem

You are given a sequence `A` of `N` nonnegative integers.

Process `Q` queries. Each query is one of:

- `1 i x`: set `A_i` to `x`.
- `2 l t`: print the smallest `r` with `l ≤ r ≤ N` such that `A_l + ... + A_r ≥ t`. If none exists, print `N + 1`.
- `3 r t`: print the largest `l` with `1 ≤ l ≤ r` such that `A_l + ... + A_r ≥ t`. If none exists, print `0`.

## Constraints

- `1 ≤ N ≤ 2 · 10^5`
- `1 ≤ Q ≤ 2 · 10^5`
- `0 ≤ A_i ≤ 10^9`
- For a type `1` query, `1 ≤ i ≤ N` and `0 ≤ x ≤ 10^9`
- For a type `2` query, `1 ≤ l ≤ N` and `1 ≤ t ≤ 10^15`
- For a type `3` query, `1 ≤ r ≤ N` and `1 ≤ t ≤ 10^15`
- All input values are integers

## Input

```text
N Q
A_1 A_2 ... A_N
query_1
query_2
...
query_Q
```

A query is `1 i x`, `2 l t`, or `3 r t`.

## Output

Let `q` be the number of type `2` and type `3` queries. Print `q` lines, in the order those queries appear.

## Sample Input 1

```text
5 8
2 1 3 0 4
2 2 4
3 5 5
1 3 1
2 2 4
3 5 5
2 4 5
3 2 4
2 1 2
```

## Sample Output 1

```text
3
3
5
3
6
0
1
```

`A_2 = 1 < 4` and `A_2 + A_3 = 4`, so the first query prints `3`. From the right, `A_4 + A_5 = 4 < 5` and `A_3 + A_4 + A_5 = 7`, so the second query prints `3`.

## Sample Input 2

```text
1 6
0
2 1 1
3 1 1
1 1 1000000000
2 1 1000000000
3 1 1000000000
2 1 1000000000000000
```

## Sample Output 2

```text
2
0
1
1
2
```

## Solution

*(TODO)*
