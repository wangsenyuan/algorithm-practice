# A - Segment Tree

[Problem link](https://atcoder.jp/contests/aalc001/tasks/aalc001_a)

**Contest:** [AAL Contest 001: Let's use segtree!](https://atcoder.jp/contests/aalc001)

time limit: 2 sec

memory limit: 1024 MiB

score: 3 points

## Problem

You are given a sequence `A` of `N` nonnegative integers.

Process `Q` queries. Each query is one of:

- `1 i x`: set `A_i` to `x`.
- `2 l r`: print the minimum of `A_l, A_{l+1}, ..., A_r`.

## Constraints

- `1 ≤ N ≤ 2 · 10^5`
- `1 ≤ Q ≤ 2 · 10^5`
- `0 ≤ A_i ≤ 10^9`
- For a type `1` query, `1 ≤ i ≤ N` and `0 ≤ x ≤ 10^9`
- For a type `2` query, `1 ≤ l ≤ r ≤ N`
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

A query is either `1 i x` or `2 l r`.

## Output

Let `q` be the number of type `2` queries. Print `q` lines. The `j`-th line is the answer to the `j`-th type `2` query.

## Sample Input 1

```text
5 5
5 3 7 9 1
2 2 4
1 3 2
2 2 4
1 5 8
2 4 5
```

## Sample Output 1

```text
3
2
8
```

The first query prints `3`, the minimum of `A_2, A_3, A_4`. The update changes `A` to `(5, 3, 2, 9, 1)`, so the next query prints `2`.

## Sample Input 2

```text
3 4
0 1000000000 7
2 1 1
2 2 2
1 1 1000000000
2 1 3
```

## Sample Output 2

```text
0
1000000000
7
```

## Solution

*(TODO)*
