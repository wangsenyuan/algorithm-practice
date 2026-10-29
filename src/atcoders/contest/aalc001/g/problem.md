# G - Sum of k-th Powers

[Problem link](https://atcoder.jp/contests/aalc001/tasks/aalc001_g)

**Contest:** [AAL Contest 001: Let's use segtree!](https://atcoder.jp/contests/aalc001)

time limit: 4 sec

memory limit: 1024 MiB

score: 3 points

## Problem

You are given a sequence `A` of `N` nonnegative integers.

Process `Q` queries. Each query is one of:

- `1 l r x`: add `x` to every `A_i` with `l ≤ i ≤ r`.
- `2 l r k`: print the sum of `A_i^k` for `i` from `l` to `r`, modulo `998244353`.

## Constraints

- `1 ≤ N ≤ 2 · 10^5`
- `1 ≤ Q ≤ 2 · 10^5`
- `0 ≤ A_i ≤ 10^9`
- For a type `1` query, `1 ≤ l ≤ r ≤ N` and `0 ≤ x ≤ 10^9`
- For a type `2` query, `1 ≤ l ≤ r ≤ N` and `1 ≤ k ≤ 4`
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

A query is either `1 l r x` or `2 l r k`.

## Output

Let `q` be the number of type `2` queries. Print `q` lines. The `j`-th line is the answer to the `j`-th type `2` query.

## Sample Input 1

```text
4 5
1 2 3 4
2 1 4 1
2 2 4 2
1 2 3 5
2 1 3 3
2 2 4 4
```

## Sample Output 1

```text
10
29
856
6753
```

The first query prints `1 + 2 + 3 + 4 = 10`. After the third query, `A = (1, 7, 8, 4)`, so the fourth query prints `1^3 + 7^3 + 8^3 = 856`.

## Sample Input 2

```text
3 5
998244352 0 2
2 1 3 2
1 1 2 1
2 1 3 4
1 2 3 1000000000
2 1 3 1
```

## Sample Output 2

```text
5
17
3511297
```

## Solution

*(TODO)*
