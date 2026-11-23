# B. Game with Strings

Source: <https://codeforces.com/problemset/problem/354/B>

Players alternately choose a letter while preserving the prefix of a path from
the top-left to the bottom-right of the grid. Determine whether optimal play
ends with more `a` letters, more `b` letters, or a draw.

## Discussion

A move appends a letter, not a direction. Therefore, a current string can
correspond to several paths and have several possible endpoints. Those
endpoints always lie on one diagonal, so represent a game state as
`(d, mask)`, where `d = row + column` is the diagonal and bit `k` in `mask`
marks a reachable cell on that diagonal.

For each state, build one successor mask per letter. For every reachable cell
in `mask`, inspect its right and down neighbors. If a neighbor contains letter
`ch`, add its position on diagonal `d + 1` to `next[ch]`. The bitwise OR is
important: choosing `ch` preserves every path that can append `ch`; a player
does not choose one particular path.

Let `f(d, mask)` be the optimal future value of `count(a) - count(b)` after
the current prefix. The initial top-left letter is forced. At diagonal `d`, the
next player chooses a non-empty `next[ch]`:

```text
f(d, mask) = min(score(ch) + f(d + 1, next[ch]))  when d is even
f(d, mask) = max(score(ch) + f(d + 1, next[ch]))  when d is odd
```

Here, `score(a) = 1`, `score(b) = -1`, and every other letter scores zero.
Even `d` means the next move belongs to the second player, because the first
player has already made the forced move at `(0, 0)`. The final diagonal has no
future move, so its value is zero.

The implementation stores each diagonal mask in local coordinates. In the
expanding half of the grid, a right move maps local bit `k` to `k` and a down
move maps it to `k + 1`. In the contracting half, the next diagonal starts one
row lower, so right maps `k` to `k - 1` and down maps `k` to `k`.

There are `O(2^n)` masks across all diagonals. Constructing all 26 successors
for one state takes `O(n + 26)`, giving total time `O(2^n * (n + 26))` and
peak DP memory `O(2^n)`.
