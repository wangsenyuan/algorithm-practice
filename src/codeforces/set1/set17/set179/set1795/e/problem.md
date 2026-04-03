# E. Explosion Finisher

You are playing a game where you kill monsters with magic spells. There are \(n\) cells in a row, numbered \(1\) to \(n\). Initially, the \(i\)-th cell contains the \(i\)-th monster with health \(h_i\).

You have a **basic spell** that costs **1 MP** and deals **1 damage** to one monster you choose; you may cast it any number of times. You also have a **special scroll** with an **“Explosion”** spell that you can use **only once**. You must finish killing all monsters with Explosion, so you first cast the basic spell some number of times (possibly zero), then cast Explosion **once**.

### How Explosion works

1. Choose a **power** \(x\) (MP you pour in): Explosion deals **\(x\)** damage to the targeted monster.
2. Choose a monster \(i\) as the target.

Then:

- If its **current** health \(h_i > x\), it **survives** with health reduced by \(x\).
- If \(h_i \le x\), monster \(i\) **dies**. Its death triggers an explosion that deals **\(h_i - 1\)** damage to monsters in cells \(i-1\) and \(i+1\) (if those cells exist and those monsters are still alive).

**Chain reactions:** If the damage is enough to kill a neighbor — i.e. the neighbor’s **current** health is at most the damage it receives (for \(i-1\): damage is \(h_i - 1\) from this step; similarly for \(i+1\)) — that monster also dies and creates a **secondary** explosion of power **(that monster’s health minus 1)** before it dies, which may hit **its** neighbors, and so on, until no more explosions occur.

Your goal is to eliminate all remaining monsters using these chained explosions. You may use basic spells earlier to lower some \(h_i\) or kill monsters outright (a monster dies when its current health becomes \(\le 0\)). Monsters **do not move**, so monsters in cells \(i\) and \(i+2\) are never neighbors.

**Task:** Minimize **total MP** = (number of basic spell casts) + (Explosion power \(x\)).

## Input

The first line contains one integer \(t\) (\(1 \le t \le 10^4\)) — the number of test cases.

The first line of each test case contains a single integer \(n\) (\(1 \le n \le 3 \cdot 10^5\)) — the number of cells / monsters.

The second line of each test case contains \(n\) integers \(h_1, h_2, \ldots, h_n\) (\(1 \le h_i \le 10^6\)) — initial healths.

It is guaranteed that the sum of \(n\) over all test cases does not exceed \(3 \cdot 10^5\).

## Output

For each test case, print one integer — the minimum total MP needed to kill all monsters while finishing with Explosion.

## Example

**Input**

```
5
3
1 1 1
4
4 1 2 1
4
5 10 15 10
1
42
9
1 2 3 2 2 2 3 2 1
```

**Output**

```
3
6
15
42
12
```

## Note

**First test case:** For example, use the basic spell on monsters \(1\) and \(2\) (once each) to kill them. Then cast Explosion with power \(x = 1\) on monster \(3\) to kill it. Total MP: \(2 + 1 = 3\).

**Second test case:** Optimal to cast the basic spell **4** times on monster \(1\) to kill it. Then cast Explosion with power \(x = 2\) on monster \(3\). It dies, releasing an explosion of power \(1\) that kills monsters \(2\) and \(4\). Total MP: \(4 + 2 = 6\).

**Third test case:** Cast Explosion with power \(15\) on monster \(3\). The explosion from monster \(3\) (power \(14\)) kills monsters \(2\) and \(4\). The secondary explosion from monster \(2\) (power \(9\)) kills monster \(1\).


### ideas
1. 假设从i开始释放爆炸魔法，且假设释放了x的能量
2. 那么可以计算出，范围（以及这个范围内相邻位置应该剩余的能量），从而可以计算出basic的能力(sum - sum of current range)
3. 这里有两个难点，一个是范围的计算，一个是x的确定。
4. x似乎就是h[i]? 假设不是，那么必然是通过操作1，将h[i]降低到了h[i] - y, 但是这时候，它对两边的伤害也小了
5. 不是最优的方案，所以x = h[i]是最优的选择，
6. 现在的问题，就是确定伤害范围，可以两边分开计算
7. 假设L[i], 表示如果i收到伤害w >= h[i]的时候，能够覆盖的范围
8. 那么如果h[i-1] <= h[i] - 1, 那么 L[i] = L[i-1], i受到的伤害可以传导
9. 如果h[i-1] > h[i] - 1, 那么这时候，应该找h[j] <= h[i] - (i - j)的位置?
10. 中间部分，是要通过操作1，削掉的。h[j] - j <= h[i] - i 的位置（最靠近的位置）
11. 这样子，形成了一棵树。为了计算，需要多少次操作1，还必须在节点上计算能删除的量
12. 假设 j....i, 那么从j+1到 i-1,依次只能保留 .. h[i] - (i - pos)    h[i] - 2, h[i] - 1
13. 保留的部分= h[i] * (i - j - 1) - (1 + .. i - (j+1)) * (i - j - 1) / 2
14. 需要操作1的部分 = sum(j+1...i-1) - 保留的部分
15. 如果j不存在呢？比如序列[10, 9, 8, 3] ?
16. j = i - h[i]?