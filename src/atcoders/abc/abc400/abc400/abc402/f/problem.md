# F - Path to Integer (ABC402)

**Contest:** [ABC402](https://atcoder.jp/contests/abc402) — Tokio Marine & Nichido Fire Insurance Programming Contest 2025  
**Task:** [https://atcoder.jp/contests/abc402/tasks/abc402_f](https://atcoder.jp/contests/abc402/tasks/abc402_f)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 525 points

## Problem Statement

There is an `N × N` grid. Cell `(i, j)` is the cell in the `i`-th row from the top and `j`-th column from the left. Each cell contains a digit from `1` to `9`; cell `(i, j)` contains `A_{i,j}`.

Initially, a token is on cell `(1, 1)`. Let `S` be an empty string. Repeat the following operation **`2N − 1` times**:

1. Append to the end of `S` the digit in the current cell.
2. Move the token one cell **down** or one cell **right**. Do **not** move on the **(2N − 1)-th** operation.

After all operations, the token is on cell `(N, N)` and `|S| = 2N − 1`.

Interpret `S` as an integer. The **score** is the remainder when this integer is divided by `M`.

Find the **maximum achievable** score over all valid paths.

## Constraints

- `1 ≤ N ≤ 20`
- `2 ≤ M ≤ 10^9`
- `1 ≤ A_{i,j} ≤ 9`
- All input values are integers.

## Input

From standard input:

```text
N M
A_{1,1} A_{1,2} ... A_{1,N}
A_{2,1} A_{2,2} ... A_{2,N}
...
A_{N,1} A_{N,2} ... A_{N,N}
```

## Output

Print one integer — the maximum score.

## Sample Input 1

```text
2 7
1 2
3 1
```

## Sample Output 1

```text
5
```

## Sample Input 2

```text
3 100000
1 2 3
3 5 8
7 1 2
```

## Sample Output 2

```text
13712
```

## Sample Input 3

```text
5 402
8 1 3 8 9
8 2 4 1 8
4 1 8 5 9
6 2 1 6 7
6 6 7 7 6
```

## Sample Output 3

```text
384
```

## Note

**Sample 1** (`N = 2`, `M = 7`):

| Path                    | `S`   | Score (`S mod M`) |
| ----------------------- | ----- | ----------------- |
| `(1,1) → (1,2) → (2,2)` | `121` | `121 mod 7 = 2`   |
| `(1,1) → (2,1) → (2,2)` | `131` | `131 mod 7 = 5`   |

The maximum is **5**.

Each path from `(1, 1)` to `(N, N)` using only right/down moves has exactly `2N − 2` moves and visits `2N − 1` cells (including start and end).


### ideas
1. 可以bruteforce吗？
2. dp[i][j] = set 所有可以得到的num % M?
3. 如果m比较小， m /64 <= 1e6 (那么用bitset)？
4. m <= 64 * 1e6
5. m比较大的时候，把它分成两部分？