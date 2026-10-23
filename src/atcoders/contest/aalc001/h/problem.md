# H - Add and Mod

[Problem link](https://atcoder.jp/contests/aalc001/tasks/aalc001_h)

**Contest:** [AAL Contest 001: Let's use segtree!](https://atcoder.jp/contests/aalc001)

time limit: 2 sec

memory limit: 1024 MiB

score: 4 points

## Problem

You are given a sequence `A` of `N` nonnegative integers.

Process `Q` queries. Each query is one of:

- `1 i x`: set `A_i` to `x`.
- `2 l r x`: start from the given `x`. For `i = l, l+1, ..., r` in that order,
  if `A_i = 0` replace `x` with `x mod 998244353`, otherwise replace `x` with
  `x + A_i`. Print the final `x`.

## Constraints

- `1 ≤ N ≤ 2 · 10^5`
- `1 ≤ Q ≤ 2 · 10^5`
- `0 ≤ A_i ≤ 10^9`
- For a type `1` query, `1 ≤ i ≤ N` and `0 ≤ x ≤ 10^9`
- For a type `2` query, `1 ≤ l ≤ r ≤ N` and `0 ≤ x ≤ 10^9`
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

A query is either `1 i x` or `2 l r x`.

## Output

Let `q` be the number of type `2` queries. Print `q` lines. The `j`-th line is
the answer to the `j`-th type `2` query.

## Sample Input 1

```text
4 5
8 0 11 5
2 1 4 998244350
1 2 6
2 2 4 10
1 4 0
2 3 4 1000000000
```

## Sample Output 1

```text
21
32
1755658
```

For the first query, `x` changes as `998244350, 998244358, 5, 16, 21`.

After the second query, `A = (8, 6, 11, 5)`. The third query then changes `x`
as `10, 16, 27, 32`.

## Sample Input 2

```text
6 7
650000000 0 910000007 123456789 0 870000013
2 3 4 750000021
2 2 6 987654321
1 5 420000009
2 3 6 314159265
1 1 0
2 1 3 271828182
2 6 6 1000000000
```

## Sample Output 2

```text
1783456817
894622424
2637616083
1181828189
1870000013
```

## Solution

*(TODO)*
