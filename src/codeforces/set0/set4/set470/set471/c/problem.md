# Problem: House of Cards

Polar bears Menshykov and Uslada from the zoo of St. Petersburg and elephant Horace from the zoo of Kiev decided to build a house of cards. For that they've already found a hefty deck of n playing cards. Let's describe the house they want to make:

## House Structure

The house consists of some non-zero number of floors.

Each floor consists of a non-zero number of rooms and the ceiling. A room is two cards that are leaned towards each other. The rooms are made in a row, each two adjoining rooms share a ceiling made by another card.

Each floor besides for the lowest one should contain less rooms than the floor below.

Please note that the house may end by the floor with more than one room, and in this case they also must be covered by the ceiling. Also, the number of rooms on the adjoining floors doesn't have to differ by one, the difference may be more.

## Problem Statement

While bears are practicing to put cards, Horace tries to figure out how many floors their house should consist of. The height of the house is the number of floors in it. It is possible that you can make a lot of different houses of different heights out of n cards. It seems that the elephant cannot solve this problem and he asks you to count the number of the distinct heights of the houses that they can make using exactly n cards.

## Input

The single line contains integer n (1 ≤ n ≤ 10^12) — the number of cards.

## Output

Print the number of distinct heights that the houses made of exactly n cards can have.

## Examples

### Example 1
**Input:**
```
13
```

**Output:**
```
1
```

### Example 2
**Input:**
```
6
```

**Output:**
```
0
```

## Note

In the first sample you can build only these two houses (remember, you must use all the cards):

[House diagrams would go here]

Thus, 13 cards are enough only for two floor houses, so the answer is 1.

The six cards in the second sample are not enough to build any house.


### ideas
1. 每一层使用 x张，x % 3 = 2的条件必须满足
2. 假设每层依次为m, m - 1, m - 2,...
3. f(m) = 2 * m + m - 1 使用了这么多张牌
4. 3 * m - 1 + 3 * (m - 1) - 1 ... >= n
5. 这样子，可以计算出高度
6. 最上面是1（可能会多出几张牌）
7. 多出来的，必须是3的倍数（否则的话，就没法用完）