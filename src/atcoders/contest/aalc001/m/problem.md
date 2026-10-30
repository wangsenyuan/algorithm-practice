# M - NAND

[Problem link](https://atcoder.jp/contests/aalc001/tasks/aalc001_m)

**Contest:** [AAL Contest 001: Let's use segtree!](https://atcoder.jp/contests/aalc001)

time limit: 2 sec

memory limit: 1024 MiB

score: 4 points

## Problem

For integers `a` and `b` in `[0, 2^30)`, their NAND is the 30-bit value whose bit `d` is `0` when both numbers have bit `d` set, and `1` otherwise.

You are given a sequence `A` of `N` integers in that range.

Process `Q` queries. Each query is one of:

- `1 p x`: set `A_p` to `x`.
- `2 l r`: set `B_l = A_l`, then `B_i = NAND(B_{i-1}, A_i)` for `i = l+1, ..., r`. Print `B_r`.

## Constraints

- `1 ≤ N ≤ 2 · 10^5`
- `1 ≤ Q ≤ 2 · 10^5`
- `0 ≤ A_i < 2^30`
- For a type `1` query, `1 ≤ p ≤ N` and `0 ≤ x < 2^30`
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

A query is either `1 p x` or `2 l r`.

## Output

Let `q` be the number of type `2` queries. Print `q` lines. The `j`-th line is the answer to the `j`-th type `2` query.

## Sample Input 1

```text
4 3
5 3 12 10
2 1 4
2 2 3
2 3 3
```

## Sample Output 1

```text
1073741821
1073741823
12
```

`B` starts at `5`. NAND with `3`, `12`, and `10` produces `1073741822`, `1073741811`, and `1073741821`.

## Sample Input 2

```text
5 6
7 10 12 3 5
2 2 4
1 3 6
2 1 5
2 3 3
1 1 1073741823
2 1 2
```

## Sample Output 2

```text
1073741820
1073741819
6
1073741813
```

## Sample Input 3

```text
3 5
1073741823 0 536870912
2 1 2
2 1 3
1 2 1073741823
2 1 3
2 2 2
```

## Sample Output 3

```text
1073741823
536870911
1073741823
1073741823
```

## Solution

NAND 按位独立，30 个位分别维护。对一个位，进入的位 `s` 经过 `A_i` 之后的结果只取决于 `s`：`s = 0` 时结果一定是 `1`，`s = 1` 时结果是 `A_i` 该位的相反值。线段树的每个节点存这两个结果 `(f0, f1)`。

左段再接右段时，用左段的两个结果去查右段。点修改只重写对应叶子。查询 `l = r` 时答案就是 `A_l`。否则取出 `A_{l+1} … A_r` 的复合函数，作用在 `A_l` 的这一位上，再左移回原位。

### 复杂度

每个位一棵线段树。建树 `O(30 n)`，每次修改或查询 `O(30 log n)`。`n, Q ≤ 2 · 10^5`。
