# E - Max Matrix 2 (ABC433)

**Contest:** [ABC433](https://atcoder.jp/contests/abc433) — AtCoder Beginner Contest 433  
**Task:** [https://atcoder.jp/contests/abc433/tasks/abc433_e](https://atcoder.jp/contests/abc433/tasks/abc433_e)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 450 points

## Problem Statement

You are given integers `N`, `M`, a sequence of `N` integers
`X = (X_1, X_2, ..., X_N)`, and a sequence of `M` integers
`Y = (Y_1, Y_2, ..., Y_M)`.

Determine whether there exists an `N` by `M` integer matrix
`A = (A_{i,j})` (`1 <= i <= N`, `1 <= j <= M`) satisfying all of the following.
If one exists, output any such matrix.

- `1 <= A_{i,j} <= N * M`
- All `N * M` entries of `A` are distinct
- For each `i = 1, 2, ..., N`:
  `max_{1 <= j <= M} A_{i,j} = X_i`
- For each `j = 1, 2, ..., M`:
  `max_{1 <= i <= N} A_{i,j} = Y_j`

You are given `T` test cases. Solve each one.

## Constraints

- `1 <= T <= 10^5`
- `1 <= N, M`
- The sum of `N * M` over all test cases is at most `2 * 10^5`
- `1 <= X_i, Y_j <= N * M`
- All input values are integers

## Input

The input is given from Standard Input in the following format:

```text
T
case_1
case_2
⋮
case_T
```

Each test case is given in the following format:

```text
N M
X_1 X_2 ... X_N
Y_1 Y_2 ... Y_M
```

## Output

Output the answer for each test case in order, one block per test case.

If no matrix satisfies all conditions, output:

```text
No
```

Otherwise, output:

```text
Yes
A_{1,1} A_{1,2} ... A_{1,M}
A_{2,1} A_{2,2} ... A_{2,M}
⋮
A_{N,1} A_{N,2} ... A_{N,M}
```

Any valid matrix is accepted if multiple solutions exist.

## Sample Input 1

```text
3
2 3
5 6
5 3 6
3 3
5 4 6
6 2 4
5 4
18 20 19 14 17
18 20 14 15
```

## Sample Output 1

```text
Yes
5 1 4
2 3 6
No
Yes
18 12 4 9
13 20 1 10
16 19 6 8
2 5 14 3
11 17 7 15
```

For the first test case, all entries are between `1` and `6`, they are distinct,
and:

- `max(A_{1,1}, A_{1,2}, A_{1,3}) = max(5, 1, 4) = 5 = X_1`
- `max(A_{2,1}, A_{2,2}, A_{2,3}) = max(2, 3, 6) = 6 = X_2`
- `max(A_{1,1}, A_{2,1}) = max(5, 2) = 5 = Y_1`
- `max(A_{1,2}, A_{2,2}) = max(1, 3) = 3 = Y_2`
- `max(A_{1,3}, A_{2,3}) = max(4, 6) = 6 = Y_3`

Another accepted output for the first test case is:

```text
Yes
5 3 1
4 2 6
```

## Solution

Every cell `(i, j)` must satisfy `A[i,j] <= X[i]` and `A[i,j] <= Y[j]`, so only
values `<= min(X[i], Y[j])` can appear there. Build buckets `ok[t]`: all cells
with `min(X[i], Y[j]) = t + 1` (using 1-based values).

### Necessary checks

- All `X[i]` must be distinct, and all `Y[j]` must be distinct. Two rows (or two
  columns) cannot share the same maximum while using distinct entries.
- Map `p1[v]` = row index with `X[i] = v`, and `p2[v]` = column index with
  `Y[j] = v` (if a value appears twice, answer is `No`).

### Greedy assignment (largest value first)

Process values `v = N*M, N*M-1, ..., 1` in descending order. Maintain a queue of
cells waiting to receive “filler” values.

| Case | Condition | Action |
| ---- | --------- | ------ |
| Both | `v` is some `X[i]` and some `Y[j]` | Place `v` at `(p1[v], p2[v])`; enqueue other cells from `ok[v-1]` |
| One  | `v` is only a row max or only a col max | Place at the first cell in `ok[v-1]`; enqueue the rest (if `ok[v-1]` is empty → `No`) |
| None | `v` is not a row or column maximum | Pop a cell from the queue (if empty → `No`); place `v`; enqueue `ok[v-1]` |

Intuition: row/column maxima are forced onto specific rows/columns (and onto one
cell when a value is both). Those placements reserve the tight cells in `ok[v]`;
smaller values that are not maxima fill queued cells left over from larger steps.

### Example (first sample)

`N = 2`, `M = 3`, `X = [5, 6]`, `Y = [5, 3, 6]`.

- `v = 6`: row 2 and column 3 → place at `(2, 3)`; queue cells with
  `min = 6` except `(2, 3)`.
- `v = 5`: row 1 and column 1 → place at `(1, 1)`; queue other `min = 5` cells.
- `v = 4, 3, 2, 1`: not row/col maxima → pop from queue and fill; each step
  extends the queue with cells whose `min` equals the current value.

This yields a valid matrix such as `[[5,1,4],[2,3,6]]`.

### Complexity

- Time: `O(N * M)` per test case
- Space: `O(N * M)`