# Hanabi Game Problem

Have you ever played Hanabi? If not, then you've got to try it out! This problem deals with a simplified version of the game.

## Problem Description

Overall, the game has 25 types of cards (5 distinct colors and 5 distinct values). Borya is holding n cards. The game is somewhat complicated by the fact that everybody sees Borya's cards except for Borya himself. Borya knows which cards he has but he knows nothing about the order they lie in. Note that Borya can have multiple identical cards (and for each of the 25 types of cards he knows exactly how many cards of this type he has).

The aim of the other players is to achieve the state when Borya knows the color and number value of each of his cards. For that, other players can give him hints. The hints can be of two types: color hints and value hints.

- **Color hint**: A player names some color and points at all the cards of this color.
- **Value hint**: A player names some value and points at all the cards that contain the value.

Determine what minimum number of hints the other players should make for Borya to be certain about each card's color and value.

## Input

The first line contains integer n (1 ≤ n ≤ 100) — the number of Borya's cards. The next line contains the descriptions of n cards. The description of each card consists of exactly two characters:
- The first character shows the color (overall this position can contain five distinct letters — R, G, B, Y, W)
- The second character shows the card's value (a digit from 1 to 5)

Borya doesn't know the exact order of the cards they lie in.

## Output

Print a single integer — the minimum number of hints that the other players should make.

## Examples

### Example 1
**Input:**
```
2
G3 G3
```

**Output:**
```
0
```

**Explanation:** In the first sample Borya already knows for each card that it is a green three.

### Example 2
**Input:**
```
4
G4 R4 R3 B3
```

**Output:**
```
2
```

**Explanation:** In the second sample we can show all fours and all red cards.

### Example 3
**Input:**
```
5
B1 Y1 W1 G1 R1
```

**Output:**
```
4
```

**Explanation:** In the third sample you need to make hints about any four colors.


### ideas
1. 没有头绪～
2. 共有n张牌，每种牌属于某种类型，由5种颜色和5个数字组成, 共25种类型
3. 然后要确定每张牌的类型，需要多少次操作
4. 每次操作，要么指定某个颜色，并挑出所有这个颜色的牌
5. 或者指定某个数字，并挑出所有这个颜色的牌
6. 假设选出了所有红色的牌，然后如果这些红色的牌的数字是一样的，那么没有必要继续处理
7. 如果红色牌的数字不一样，那么就需要继续处理
8. 假设现在有两种颜色R、B，然后有3种数字，那么似乎选择颜色是更优的选择？
9. 假设数字1都是R，数字2、3由R、B组成
10. 那么选择数字1似乎更优？
11. 如果数字1被选择后，那么后续还需要再选择数字1吗？似乎是不用的
12. 所有dp[color][num] = 目前为止选择了color和num状态后，还需要进行多少操作？