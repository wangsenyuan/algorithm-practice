# Problem E: Horse Land Cards

## Problem Description

There are three horses living in a horse land: one gray, one white and one gray-and-white. The horses are really amusing animals, which is why they adore special cards. Each of those cards must contain two integers, the first one on top, the second one in the bottom of the card. Let's denote a card with `a` on the top and `b` in the bottom as `(a, b)`.

Each of the three horses can paint the special cards:

- **Gray horse**: If you show an `(a, b)` card to the gray horse, then the horse can paint a new `(a + 1, b + 1)` card.
- **White horse**: If you show an `(a, b)` card, such that `a` and `b` are even integers, to the white horse, then the horse can paint a new `(a/2, b/2)` card.
- **Gray-and-white horse**: If you show two cards `(a, b)` and `(b, c)` to the gray-and-white horse, then he can paint a new `(a, c)` card.

Polycarpus really wants to get `n` special cards `(1, a₁)`, `(1, a₂)`, ..., `(1, aₙ)`. For that he is going to the horse land. He can take exactly one `(x, y)` card to the horse land, such that `1 ≤ x < y ≤ m`. 

**Question**: How many ways are there to choose the card so that he can perform some actions in the horse land and get the required cards?

**Note**: Polycarpus can get cards from the horses only as a result of the actions that are described above. Polycarpus is allowed to get additional cards besides the cards that he requires.

## Input

- The first line contains two integers `n, m` (`1 ≤ n ≤ 10⁵`, `2 ≤ m ≤ 10⁹`).
- The second line contains the sequence of integers `a₁, a₂, ..., aₙ` (`2 ≤ aᵢ ≤ 10⁹`). Note that the numbers in the sequence can coincide.

The numbers in the lines are separated by single spaces.

## Output

Print a single integer — the answer to the problem.

**Note**: Please, do not use the `%lld` specifier to read or write 64-bit integers in C++. It is preferred to use the `cin`, `cout` streams or the `%I64d` specifier.

## Examples

### Example 1
**Input:**
```
1 6
2
```

**Output:**
```
11
```

### Example 2
**Input:**
```
1 6
7
```

**Output:**
```
14
```

### Example 3
**Input:**
```
2 10
13 7
```

**Output:**
```
36
```


## ideas
1. 完全没有想法～
2. (x, y) => 通过操作1、和操作2，都不能得到 x >= y 的情况出现
3. 也就是说，无论什么时候，都有 x < y的条件成立
4. 操作3， (x, y), (y, z) => (x, z) 同样满足 x < z的情况
5. 考虑如何将x变成1， 如果x是偶数，那么直接操作2
6. 如果x是奇数，那么(x + 1) / 2
7. 考虑y的影响