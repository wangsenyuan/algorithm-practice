# B. Colorful Field

[Problem link](https://codeforces.com/problemset/problem/79/B)

**Contest:** [Codeforces Beta Round 70 (Div. 2)](https://codeforces.com/contest/79)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Fox Ciel saw a large field while she was on a bus. The field was an `n × m`
rectangle divided into `1 × 1` cells. Some cells were wasteland, and every other
cell contained crop plants: either carrots or kiwis or grapes.

After seeing the field carefully, Ciel found that the crop plants of each cell
were planted in the following procedure:

Assume that the rows are numbered `1` to `n` from top to bottom and the columns
are numbered `1` to `m` from left to right, and a cell in row `i` and column `j`
is represented as `(i, j)`.

- First, each field is either cultivated or waste. Crop plants will be planted
  in the cultivated cells in the order of
  `(1, 1) → … → (1, m) → (2, 1) → … → (2, m) → … → (n, 1) → … → (n, m)`. Waste
  cells will be ignored.
- Crop plants (either carrots or kiwis or grapes) will be planted in each cell
  one after another cyclically. Carrots will be planted in the first cell, then
  kiwis in the second one, grapes in the third one, carrots in the fourth one,
  kiwis in the fifth one, and so on.

Now she is wondering how to determine the crop plants in some certain cells.

## Constraints

- `1 ≤ n ≤ 4 · 10^4`
- `1 ≤ m ≤ 4 · 10^4`
- `1 ≤ k ≤ 10^3`
- `1 ≤ t ≤ 10^3`
- `1 ≤ a, i ≤ n`, `1 ≤ b, j ≤ m`
- Waste cells are distinct

## Input

The first line contains four positive integers `n`, `m`, `k`, `t` — the height
of the field, the width of the field, the number of waste cells, and the number
of queries.

Each of the next `k` lines contains two integers `a`, `b`, denoting that cell
`(a, b)` is waste.

Each of the next `t` lines contains two integers `i`, `j` — a query asking the
kind of crop plants in cell `(i, j)`.

## Output

For each query, if the cell is waste, print `Waste`. Otherwise, print the name
of the crop plants in the cell: either `Carrots` or `Kiwis` or `Grapes`.

## Example

### Input

```text
4 5 5 6
4 3
1 3
3 3
2 5
3 2
1 3
1 4
2 3
2 4
1 1
1 1
```

### Output

```text
Waste
Grapes
Carrots
Kiwis
Carrots
Carrots
```

## Note

The sample corresponds to the figure in the statement.

## Solution

*(TODO)*
