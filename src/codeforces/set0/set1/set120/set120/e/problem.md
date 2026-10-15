# E. Put Knight!

[Problem link](https://codeforces.com/problemset/problem/120/E)

**Contest:** [School Regional Team Contest, Saratov, 2011](https://codeforces.com/contest/120)

**Tags:** games, math, *1400

time limit per test: 1 second

memory limit per test: 256 megabytes

input: input.txt

output: output.txt

Petya and Gena play "Put a Knight!" on an `n` by `n` chessboard. They take turns placing knights so that no two knights threaten each other. A knight on square `(r, c)` threatens `(r-1, c+2)`, `(r-1, c-2)`, `(r+1, c+2)`, `(r+1, c-2)`, `(r-2, c+1)`, `(r-2, c-1)`, `(r+2, c+1)`, and `(r+2, c-1)`. Squares outside the board are ignored. The player who cannot place a knight loses. Both play optimally, and Petya starts. Decide who wins.

## Input

The first line contains an integer `T` (`1 <= T <= 100`), the number of boards.

Each of the next `T` lines contains one integer `n_i` (`1 <= n_i <= 10000`), the side length of that board.

## Output

For each board, print `0` if Petya wins and `1` otherwise.

## Example

### Input

```text
2
2
1
```

### Output

```text
1
0
```
