# Problem A: Monster Battle

A monster is attacking the Cyberland!

Master Yang, a brave warrior, is going to beat the monster. Yang and the monster each have 3 attributes: hitpoints (HP),
offensive power (ATK) and defensive power (DEF).

## Battle Mechanics

During the battle, every second:

- The monster's HP decreases by `max(0, ATK_Y - DEF_M)`
- Yang's HP decreases by `max(0, ATK_M - DEF_Y)`

Where:

- Index Y denotes Master Yang
- Index M denotes monster

Both decreases happen simultaneously. Master Yang wins when the monster's HP ≤ 0 and at the same time Master Yang's HP >

0.

## Shop System

Master Yang can buy attributes from the magic shop of Cyberland:

- **h** bitcoins per HP
- **a** bitcoins per ATK
- **d** bitcoins per DEF

Now Master Yang wants to know the minimum number of bitcoins he can spend in order to win.

## Input

The first line contains three integers `HP_Y`, `ATK_Y`, `DEF_Y`, separated by a space, denoting the initial HP, ATK and
DEF of Master Yang.

The second line contains three integers `HP_M`, `ATK_M`, `DEF_M`, separated by a space, denoting the HP, ATK and DEF of
the monster.

The third line contains three integers `h`, `a`, `d`, separated by a space, denoting the price of 1 HP, 1 ATK and 1 DEF.

**Constraints:** All numbers in input are integers and lie between 1 and 100 inclusively.

## Output

The only output line should contain an integer, denoting the minimum bitcoins Master Yang should spend in order to win.

## Examples

### Example 1

**Input:**

```
1 2 1
1 100 1
1 100 100
```

**Output:**

```
99
```

### Example 2

**Input:**

```
100 100 100
1 1 1
1 1 1
```

**Output:**

```
0
```

## Notes

- For the first sample, prices for ATK and DEF are extremely high. Master Yang can buy 99 HP, then he can beat the
  monster with 1 HP left.

- For the second sample, Master Yang is strong enough to beat the monster, so he doesn't need to buy anything.

## ideas

1. 首先要保证不败， 假设一共经过了n秒钟, 那么Y收到的伤害 $S_1$= n * ($ATK_m$ - $DEF_y$)
2. 那么必须有 $HP_y - S_1 $ > 0
3. $S_2$ = n * ($ATK_y - DEF_m$)
4. $HP_m - S_2$ <= 0
5. 所以攻击越高，n越小; 在给定n的情况下，增加HP或者DEF;
6. 迭代$ATK_y$， 这样子n就确定了，在给定n的情况下，计算最优的cost(要增加的HP和DEF)
7. 假设分别增加x和y，那么要满足
8. $ h \times x + d \times y $ <= cost
9. $ HP_y + x > n \times (ATK_m - DEF_y - y)$
10. $x + n \times y > n \times (ATK_m - DEF_y) - HP_y$
11. $(d - h \times n) \times y < cost + h \times n \times (ATK_m - DEF_y) - h \times HP_y $ 
12. 如果上面存在解，那么就可以得到cost (y >= 0)