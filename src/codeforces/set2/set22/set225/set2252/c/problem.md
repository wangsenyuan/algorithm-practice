# C. Risky Tower

[Problem link](https://codeforces.com/problemset/problem/2252/C)

**Contest:** [Codeforces Round 1115 (Div. 2)](https://codeforces.com/contest/2252)

time limit per test: 3 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are playing a 2D Jenga game represented by a grid with `n` rows and `m` columns. The `1`-st row is the top level of the tower, and the `n`-th row is the bottom level.

Each piece at row `i` and column `j` has a destabilization factor `a_{i,j}`. Additionally, each row `i` has an initial stability index of `v_i`.

When you remove a piece, it is completely discarded from the game. Removing a piece from row `i` damages all levels at and above it. Specifically, for every row `k` such that `1 ≤ k ≤ i`, its stability index is decreased by `a_{i,j}`.

The tower collapses if either of the following conditions is met:

- The stability index of any level drops to `0` or less.
- Any level is left with exactly `0` pieces (even if it is the topmost level).

Find the minimum number of pieces you must remove such that the tower collapses.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 ≤ t ≤ 10^4`). The description of the test cases follows.

The first line of each test case contains two integers `n` and `m` (`1 ≤ n, m ≤ 10^6`) — the number of rows and columns of the tower.

The second line contains `n` integers `v_1, v_2, …, v_n` (`1 ≤ v_i ≤ 10^9`) — the initial stability indices of each level from top to bottom.

Each of the next `n` lines contains `m` integers. The `j`-th integer on the `i`-th line is `a_{i,j}` (`1 ≤ a_{i,j} ≤ 10^9`) — the destabilization factor of the piece at row `i` and column `j`.

It is guaranteed that the sum of `n · m` over all test cases does not exceed `10^6`.

## Output

For each test case, output a single integer — the minimum number of pieces that must be removed to collapse the tower.

## Example

### Input

```text
2
2 3
10 20
2 2 2
5 5 5
3 1
100 100 100
1
2
3
```

### Output

```text
2
1
```

### Note

In the first test case, we have a `2 × 3` tower. The initial stabilities are `v_1 = 10` and `v_2 = 20`. The top level (`i = 1`) has pieces with destabilization factors `2, 2, 2`. The bottom level (`i = 2`) has pieces `5, 5, 5`.

If we remove two pieces from the bottom level, it inflicts `5 + 5 = 10` damage to level `2`, and `5 + 5 = 10` damage to level `1`. The remaining stability of level `1` becomes `10 − 10 = 0`. Since its stability dropped to `≤ 0`, the tower collapses. Thus, the minimum number of pieces we must remove is `2`.

In the second test case, `m = 1`. The rule states that the tower collapses if any level is left with exactly `0` pieces. Since every level initially has `1` piece, removing any `1` piece from the tower will instantly empty a level and cause a collapse. Therefore, the minimum number of pieces to remove is `1`.
