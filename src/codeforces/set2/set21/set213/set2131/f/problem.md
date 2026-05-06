# F. Unjust Binary Life

For a fixed target `(x, y)`, every monotone path from `(1, 1)` to `(x, y)` visits every row `1..x` and every column `1..y` at least once. The zero cells on that path form equality constraints

`a_i = b_j`

for all visited cells. Since the path is connected, all visited row and column variables must become one common bit.

Therefore the exact shape of the path does not matter. For prefixes `a[1..x]` and `b[1..y]`, the minimum number of flips is

`min(onesA[x] + onesB[y], zeroesA[x] + zeroesB[y])`.

The task is to sum this over all prefix pairs.

Let

`diffA[x] = onesA[x] - zeroesA[x]`

and similarly for `diffB[y]`. For one fixed `x`, choosing common bit `0` is no worse exactly when

`diffA[x] + diffB[y] <= 0`.

Sort all `b` prefixes by `diffB`. For each `a` prefix, binary-search the last `b` prefix satisfying the inequality. The left part contributes `onesA + onesB`; the right part contributes `zeroesA + zeroesB`. Prefix sums of `onesB` and `zeroesB` make each query `O(log n)`.

Overall complexity is `O(n log n)` per test case.
