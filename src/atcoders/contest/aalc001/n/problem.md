# N - Times Tables

[Problem link](https://atcoder.jp/contests/aalc001/tasks/aalc001_n)

**Contest:** [AAL Contest 001: Let's use segtree!](https://atcoder.jp/contests/aalc001)

time limit: 2 sec

memory limit: 1024 MiB

score: 4 points

## Problem

There is an `N` by `N` grid. Cell `(i, j)` starts with the integer `i * j`.

Perform `Q` operations. Operation `k` overwrites every cell `(i, j)` with `a_k ≤ i ≤ c_k` and `b_k ≤ j ≤ d_k` by `0`.

Print the sum of all cells after every operation, modulo `998244353`.

## Constraints

- `1 ≤ N ≤ 2 · 10^5`
- `1 ≤ Q ≤ 2 · 10^5`
- `1 ≤ a_k ≤ c_k ≤ N`
- `1 ≤ b_k ≤ d_k ≤ N`
- All input values are integers

## Input

```text
N Q
a_1 b_1 c_1 d_1
a_2 b_2 c_2 d_2
...
a_Q b_Q c_Q d_Q
```

## Output

Print the sum modulo `998244353`.

## Sample Input 1

```text
3 1
2 2 3 3
```

## Sample Output 1

```text
11
```

The grid starts as

```text
1 2 3
2 4 6
3 6 9
```

and ends as

```text
1 2 3
2 0 0
3 0 0
```

## Sample Input 2

```text
3 2
1 2 2 3
2 1 3 2
```

## Sample Output 2

```text
10
```

## Sample Input 3

```text
4 3
1 1 2 4
3 1 4 2
3 3 4 4
```

## Sample Output 3

```text
0
```

## Solution

*(TODO)*
