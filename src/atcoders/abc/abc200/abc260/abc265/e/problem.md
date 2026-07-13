# E - Warp

[Problem link](https://atcoder.jp/contests/abc265/tasks/abc265_e)

**Contest:** [AtCoder Beginner Contest 265](https://atcoder.jp/contests/abc265)

time limit: 3 sec

memory limit: 1024 MiB

score: 500 points

Takahashi is at the origin of a two-dimensional plane. He will teleport `N` times.
In each teleportation, he makes one of the following moves:

- from `(x, y)` to `(x + A, y + B)`
- from `(x, y)` to `(x + C, y + D)`
- from `(x, y)` to `(x + E, y + F)`

There are obstacles on `M` points `(X_1, Y_1), ..., (X_M, Y_M)`. He cannot
teleport onto these coordinates.

How many paths of `N` teleportations are there? Print the count modulo
`998244353`.

## Constraints

- `1 <= N <= 300`
- `0 <= M <= 10^5`
- `-10^9 <= A, B, C, D, E, F <= 10^9`
- `(A, B)`, `(C, D)`, and `(E, F)` are pairwise distinct
- `-10^9 <= X_i, Y_i <= 10^9`
- `(X_i, Y_i) != (0, 0)`
- All `(X_i, Y_i)` are distinct
- All input values are integers

## Input

```text
N M
A B C D E F
X_1 Y_1
...
X_M Y_M
```

## Output

Print the answer.

## Sample Input 1

```text
2 2
1 1 1 2 1 3
1 2
2 2
```

## Sample Output 1

```text
5
```

Possible paths:

- `(0,0) -> (1,1) -> (2,3)`
- `(0,0) -> (1,1) -> (2,4)`
- `(0,0) -> (1,3) -> (2,4)`
- `(0,0) -> (1,3) -> (2,5)`
- `(0,0) -> (1,3) -> (2,6)`

## Sample Input 2

```text
10 3
-1000000000 -1000000000 1000000000 1000000000 -1000000000 1000000000
-1000000000 -1000000000
1000000000 1000000000
-1000000000 1000000000
```

## Sample Output 2

```text
0
```

## Sample Input 3

```text
300 0
0 0 1 0 0 1
```

## Sample Output 3

```text
292172978
```

## ideas
1. dp[i][x][y] 表示经过i次操作后, 目前在位置(x, y)处时的ways
2. 这里有个问题时, (x, y)可能会被重复的到达, 这样子有没有关系? 因为有次数, 是没有关系的.
3. 空间复杂性多少呢? 似乎是 pow(3, n)个的.
4. 如果真这么大, 肯定是不大行的. 可能这里分析不对.

## Solution

Let the three move vectors be `move[0]`, `move[1]`, and `move[2]`.

Instead of using the current coordinate as the DP state, record how many times each kind of move has been selected. After exactly `w` teleports, suppose:

- the first move has been used `x` times;
- the second move has been used `y` times;
- the third move has been used `z = w - x - y` times.

Therefore, once `w`, `x`, and `y` are known, `z` is also determined. We only need a two-dimensional DP table:

```text
dp[x][y] = number of valid sequences after w teleports
           that use the first move x times and the second move y times
```

The current coordinate of this state is:

```text
(row, column)
    = x * move[0] + y * move[1] + z * move[2]
```

This formula works because vector addition is commutative: the final coordinate depends only on how many times each move was used, not on their order. However, `dp[x][y]` still counts all different orders that produce those move counts.

Store all obstacle coordinates in a hash set so that each destination can be checked in expected `O(1)` time.

Initially, no move has been used, so:

```text
dp[0][0] = 1
```

For every step `w`, consider each state `(x, y)` with `x + y <= w`. Compute its current coordinate and try all three next moves:

- First move: add `dp[x][y]` to `ndp[x+1][y]`.
- Second move: add `dp[x][y]` to `ndp[x][y+1]`.
- Third move: add `dp[x][y]` to `ndp[x][y]`.

A transition is performed only if its destination is not an obstacle. All additions are taken modulo `998244353`.

After processing one teleport, swap `dp` and `ndp`. Thus, only two `O(N^2)` tables are needed. After all `N` teleports, sum every value in `dp` to obtain the answer.

### Why move counts are convenient states

One could also group paths by their current coordinate, but that would require a map for every DP layer. Move counts give a fixed `O(N^2)` array: `x` and `y` are both between `0` and `N`, and the current coordinate can be calculated directly from them.

Different move-count combinations may arrive at the same coordinate when the move vectors are linearly dependent. The code may keep those combinations as separate states; this does not double-count any path because every sequence has exactly one move-count triple. Different orders with the same counts are intentionally combined in `dp[x][y]`.

### Correctness Proof

We prove by induction on `w` that `dp[x][y]` equals the number of valid sequences of `w` teleports that use the first move `x` times, the second move `y` times, and the third move `w-x-y` times.

For `w = 0`, the empty sequence is the only sequence. It uses every move zero times, and the initialization `dp[0][0] = 1` is correct.

Assume the statement is true after `w` teleports. Consider a state `(x, y)`. Every sequence counted by this state can choose one of the three moves next:

- choosing the first move produces state `(x+1, y)`;
- choosing the second move produces state `(x, y+1)`;
- choosing the third move produces state `(x, y)` in the next layer.

The algorithm computes the exact destination coordinate and adds the sequence count only when that destination is not blocked. Thus, it adds exactly every valid sequence of `w+1` moves and rejects exactly every sequence that lands on an obstacle. Every non-empty sequence has a unique last move, so no sequence is omitted or counted twice by these transitions.

Therefore, the invariant also holds after `w+1` teleports. By induction, it holds after `N` teleports. Summing all states then counts exactly all valid sequences of `N` teleports, so the algorithm is correct.

### Complexity

There are `O(w^2)` states at step `w`, and each state has three constant-time transitions. Summed over all `N` steps:

- Time: `O(N^3 + M)`
- Space: `O(N^2 + M)`
