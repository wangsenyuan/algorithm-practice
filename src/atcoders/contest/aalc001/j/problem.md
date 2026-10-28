# J - Long Binary String

[Problem link](https://atcoder.jp/contests/aalc001/tasks/aalc001_j)

**Contest:** [AAL Contest 001: Let's use segtree!](https://atcoder.jp/contests/aalc001)

time limit: 2 sec

memory limit: 1024 MiB

score: 4 points

## Problem

`S` is a binary string of length `10^18`. Every character starts as `0`.

Process `Q` queries. Each query is one of:

- `1 l r`: flip `S_l, S_{l+1}, ..., S_r`. Change each `0` to `1` and each `1` to `0`.
- `2 l r`: print the number of `1`s among `S_l, S_{l+1}, ..., S_r`.

## Constraints

- `1 ≤ Q ≤ 2 · 10^5`
- `1 ≤ l ≤ r ≤ 10^18`
- All input values are integers

## Input

```text
Q
query_1
query_2
...
query_Q
```

A query is either `1 l r` or `2 l r`.

## Output

Let `q` be the number of type `2` queries. Print `q` lines. The `j`-th line is the answer to the `j`-th type `2` query.

## Sample Input 1

```text
8
2 1 10
1 3 7
2 1 10
1 5 12
2 4 9
2 3 12
1 1 4
2 1 12
```

## Sample Output 1

```text
0
5
3
7
7
```

Every character starts as `0`. After the second query, the `1`s are `S_3` through `S_7`. After the fourth query, the `1`s are `S_3`, `S_4`, and `S_8` through `S_12`.

## Sample Input 2

```text
6
1 999999999999999998 1000000000000000000
2 999999999999999997 1000000000000000000
1 1 1000000000000000000
2 1 1000000000000000000
2 999999999999999997 1000000000000000000
2 1 1
```

## Sample Output 2

```text
3
999999999999999997
1
1
```

## Solution

*(TODO)*
