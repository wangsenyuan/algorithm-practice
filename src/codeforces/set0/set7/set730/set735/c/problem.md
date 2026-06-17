# C. Tennis Championship

[Problem link](https://codeforces.com/problemset/problem/735/C)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

A tennis tournament in Rio has `n` players. It is single elimination from the
first match: a player who loses leaves immediately.

Organizers fix one pairing rule before the bracket is built:

> Two players may meet only if the number of matches one has already played differs
> from the other's by **at most 1**. (Both must have won all their previous matches
> to still be in the tournament.)

The bracket order is not fixed yet, but this rule is. Find the **maximum** number
of matches the **champion** can play, over all valid arrangements.

## Input

One integer `n` (`2 <= n <= 10^18`) — the number of players.

## Output

Print one integer — the maximum possible number of games played by the tournament
winner.

## Examples

### Input

```text
2
```

### Output

```text
1
```

### Input

```text
3
```

### Output

```text
2
```

### Input

```text
4
```

### Output

```text
2
```

### Input

```text
10
```

### Output

```text
4
```

## Note

Assume player 1 is the champion in each sample.

- `n = 2`: one final → **1** game.
- `n = 3`: player 1 can beat 2, then 3 → **2** games.
- `n = 4`: player 1 cannot play everyone; after beating 2 and 3 (2 games), player
  4 has played 0 games, so the difference is 2. Pair `(1,2)` and `(3,4)`, then the
  winners meet → champion plays **2** games.
- `n = 10`: answer is **4**.


## ideas
1. let F[x] = n 表示经过x次获胜, 至少需要n个人参赛
2. F[x+1] = F[x] + F[x-1]
3. F[1] = 2 
4. F[2] = 3 (fib序列了)