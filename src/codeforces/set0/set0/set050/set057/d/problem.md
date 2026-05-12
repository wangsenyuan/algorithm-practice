# D. Average lifespan (Codeforces 57D)

Stewie the Rabbit explores a new parallel universe shaped as a rectangular grid with `n` rows and `m` columns. The universe is small: each cell holds **at most one** particle.

Each particle is either **static** or **dynamic**.

- **Static** particles never move. No two static particles share a **row** or a **column**, and none occupy **diagonally adjacent** cells.
- A **dynamic** particle appears in a **uniformly random empty** cell, then chooses a **uniformly random empty** cell as its destination (the destination may equal the start). It moves along a **shortest path** through cells not occupied by static particles. When it reaches the destination, it disappears. Only one dynamic particle exists at a time.

The particle may move between **edge-adjacent** cells; each such step takes **one galactic second**.

Find the **expected** (average) lifespan of one such particle.

## Input

The first line contains two integers `n` and `m` (`2 <= n, m <= 1000`) — grid dimensions.

The next `n` lines each contain `m` characters:

- `'X'` — static particle in that cell  
- `'.'` — empty cell  

The grid is guaranteed to satisfy the static-particle rules above.

## Output

Print a single real number — the average lifespan, with **at least 6** decimal digits shown.

The answer is accepted if the absolute or relative error is at most `10^-6`.

## Examples

### Example 1

**Input**

```text
2 2
..
.X
```

**Output**

```text
0.888888888889
```

### Example 2

**Input**

```text
3 3
...
.X.
...
```

**Output**

```text
2.000000000000
```

Original: [Codeforces 57D](https://codeforces.com/problemset/problem/57/D)


### ideas
1. 固定起点(r1, c1), 计算所有其他节点到它的最少距离的和，除以(free - 1)
2. 加起来，最后再除以 free?
3. 但是这样的复杂性 = n * m * n * m too much
4. 所以要利用没有同一行，用一列有X的性质
5. 如果在某一行里面没有X，那么这一行的贡献很容易计算（同列同理）
6. 如果有一个X，那么+2？因为，没有相邻的斜向的X