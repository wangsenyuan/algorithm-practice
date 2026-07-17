# F - Gather Coins (ABC369)

**Contest:** [ABC369](https://atcoder.jp/contests/abc369) — AtCoder Beginner Contest 369  
**Task:** [https://atcoder.jp/contests/abc369/tasks/abc369_f](https://atcoder.jp/contests/abc369/tasks/abc369_f)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 500 points

## Problem Statement

There is a grid with `H` rows and `W` columns. Cell `(i, j)` is the cell at the `i`-th row from the top and `j`-th column from the left.

There are `N` coins on this grid. The `i`-th coin can be picked up by passing through cell `(R_i, C_i)`.

Start from cell `(1, 1)`, repeatedly move either **down** or **right** by one cell, and reach cell `(H, W)` while picking up as many coins as possible.

Find the **maximum** number of coins you can pick up and **one** path that achieves this maximum.

## Constraints

- `2 ≤ H, W ≤ 2 × 10^5`
- `1 ≤ N ≤ min(HW − 2, 2 × 10^5)`
- `1 ≤ R_i ≤ H`
- `1 ≤ C_i ≤ W`
- `(R_i, C_i) ≠ (1, 1)`
- `(R_i, C_i) ≠ (H, W)`
- `(R_i, C_i)` are pairwise distinct.
- All input values are integers.

## Input

From standard input:

```text
H W N
R_1 C_1
R_2 C_2
...
R_N C_N
```

## Output

Print two lines.

1. The maximum number of coins you can pick up.
2. One optimal path as a string of length `H + W − 2`. The `i`-th character is:
   - `D` if the `i`-th move is downward.
   - `R` if the `i`-th move is rightward.

If multiple paths maximize the number of coins, print any of them.

## Sample Input 1

```text
3 4 4
3 3
2 1
2 3
1 4
```

## Sample Output 1

```text
3
DRRDR
```

By moving `(1,1) → (2,1) → (2,2) → (2,3) → (3,3) → (3,4)`, you pick up three coins at `(2,1)`, `(2,3)`, and `(3,3)`.

## Sample Input 2

```text
2 2 2
2 1
1 2
```

## Sample Output 2

```text
1
DR
```

The path `RD` is also acceptable.

## Sample Input 3

```text
10 15 8
2 7
2 9
7 9
10 3
7 11
8 12
9 6
8 1
```

## Sample Output 3

```text
5
DRRRRRRRRDDDDDRRDRDDRRR
```
