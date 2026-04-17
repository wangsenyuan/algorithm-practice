# Problem

The game consists of multiple rounds.  
In each round, a natural number $k$ is chosen. The player who says (or barks) it faster wins the round.

After that:

- the winner's score is multiplied by $k^2$,
- the loser's score is multiplied by $k$.

Initially, both Slastyona and Pushok have score $1$.

Slastyona lost the notebook containing the history of all games, but she remembers the final scores.  
For each given pair of final scores, determine whether such a result is possible.

## Input

The first line contains an integer $n$ ($1 \le n \le 350000$) — the number of games.

Each of the next $n$ lines contains two integers $a, b$ ($1 \le a, b \le 10^9$) — final scores of Slastyona and Pushok, respectively.

## Output

For each pair, print:

- `Yes` if such a final result is possible,
- `No` otherwise.

You may print letters in any case.

## Example

### Input

```text
6
2 4
75 45
8 8
16 16
247 994
1000000000 1000000
```

### Output

```text
Yes
Yes
Yes
No
No
Yes
```

## Note

In the first game, there could be one round with $k = 2$, won by Pushok.

The second game needs exactly two rounds:

- first round: Slastyona says $k = 5$ faster,
- second round: Pushok barks $k = 3$ faster.

## ideas
1. 感觉还挺难的～
2. 没有说，一个数k，只能被选一次。
3. 考虑一个质数，它在两边可以出现相同的次数（3的倍数）
4. 也可以左边比右边多, 假设左边出现了x次，右边出现了y次, x > y
5. 那么必须存在一个数 w, w % 3 = 0, 且 (x - 3 * w) / (y - 3 * w) = 2
6. x - 3 * w = 2 * y - 6 * w
7. => 3 * w = 2 * y - x
8. 2 * y - x 要能整除3
9. 但是a、b很大，通过sqrt(a)找出所有的质因数，太慢了（考虑到n有350000）
10. 

## Thoughts

这题最自然的办法是盯住每个质因子的指数。

因为每一轮操作本质上只是：

- 赢家乘上 `k^2`
- 输家乘上 `k`

所以对任意一个固定质数 `p`，如果这一轮里 `v_p(k) = t`，那么它对两个分数的指数贡献只有两种可能：

- `(2t, t)`
- `(t, 2t)`

也就是说，每轮对某个质数的贡献，都是：

- 一边加两份
- 另一边加一份

### 1. 最终指数必须满足什么形式

设最后：

- `v_p(a) = α`
- `v_p(b) = β`

把所有轮次对这个质数的贡献累加起来，一定可以写成：

- `α = 2x + y`
- `β = x + 2y`

其中 `x, y >= 0` 是整数。

这里：

- `x` 表示那些“第一人赢”的轮次在这个质数上的总贡献
- `y` 表示那些“第二人赢”的轮次在这个质数上的总贡献

所以问题就变成：

- 给定 `(α, β)`，什么时候能表示成 `(2x+y, x+2y)`？

### 2. 先看和

把两式相加：

- `α + β = 3(x + y)`

所以必须有：

- `α + β` 是 `3` 的倍数

这意味着：

- 对每个质数 `p`，`v_p(a) + v_p(b)` 都是 `3` 的倍数

而这恰好等价于：

- `a * b` 是完全立方数

设：

- `a * b = c^3`

那么对每个质数 `p`，

- `v_p(c) = (α + β) / 3`

### 3. 再看非负性

由

- `α = 2x + y`
- `β = x + 2y`

解出：

- `x = (2α - β) / 3`
- `y = (2β - α) / 3`

为了让 `x, y` 都是非负整数，必须有：

- `2α >= β`
- `2β >= α`

把 `v_p(c) = (α + β) / 3` 代进去看，其实这等价于：

- `α >= v_p(c)`
- `β >= v_p(c)`

也就是：

- `c | a`
- `c | b`

### 4. 最终判定条件

所以答案为 `Yes` 当且仅当：

1. `a * b` 是完全立方数
2. 设 `c = cbrt(a * b)`，则：
   - `c` 整除 `a`
   - `c` 整除 `b`

这两个条件既是必要的，也是充分的。

### 5. 为什么这样就够了

原题看起来像是：

- 要猜很多轮
- 每轮选哪个 `k`
- 每轮谁赢

但其实完全不用恢复这些过程。

因为一旦每个质因子的指数条件成立，就一定存在某种轮次安排把它们拼出来。

所以实现只需要：

1. 计算 `prod = a * b`
2. 求它的整数立方根 `c`
3. 检查：
   - `c^3 == prod`
   - `a % c == 0`
   - `b % c == 0`

满足则输出 `Yes`，否则输出 `No`。
