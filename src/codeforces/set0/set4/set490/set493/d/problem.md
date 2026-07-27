# D. Vasya and Chess

[Problem link](https://codeforces.com/problemset/problem/493/D)

**Contest:** [Codeforces Round #281 (Div. 2)](https://codeforces.com/contest/493)

time limit: 2 seconds

memory limit: 256 megabytes

## Problem

There is an `n × n` chessboard. Cell `(1, 1)` has the white queen and cell
`(1, n)` has the black queen. Every other cell has a green pawn that belongs
to nobody.

Players alternate turns; white moves first. On each turn a player must capture
some piece with their queen (a green pawn or the enemy queen), moving like a
chess queen and unable to jump over pieces. A player loses if they cannot
capture anything, or if their queen was taken on the previous turn.

Determine who wins with optimal play. If white wins, also output the
lexicographically smallest first move `(r, c)` (minimize `r`, then `c`).

## Constraints

- `2 <= n <= 10^9`

## Input

```text
n
```

## Output

Print `white` or `black` on the first line.

If the answer is `white`, print a second line with two integers `r` and `c` —
white's first winning move.

## Sample 1

```text
Input
2

Output
white
1 2
```

White can capture the black queen on the first move.

## Sample 2

```text
Input
3

Output
black
```


## ideas
1. 如果n是奇数, 当white向black移动的时候, black可以镜像white的操作,
2. 这样子, black和white始终处于同一列, 且它们之间的距离始终是奇数.
3. 当white移动到最中间一行时, black就能攻击到它了
4. 现在考虑n是偶数. n = 2, white胜
5. n = 4, white胜利, 且第一步移动到 (1, 2) 这时候, black只能往(4, 2)移动
6. 下一步white移动到(2, 2), 就进入了n = 3 的状态, 但是由black启动
7. n = 6
