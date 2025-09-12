# Problem Description

Vasily has a deck of cards consisting of n cards. There is an integer on each of the cards, this integer is between 1 and 100,000, inclusive. It is possible that some cards have the same integers on them.

Vasily decided to sort the cards. To do this, he repeatedly takes the top card from the deck, and if the number on it equals the minimum number written on the cards in the deck, then he places the card away. Otherwise, he puts it under the deck and takes the next card from the top, and so on. The process ends as soon as there are no cards in the deck. You can assume that Vasily always knows the minimum number written on some card in the remaining deck, but doesn't know where this card (or these cards) is.

You are to determine the total number of times Vasily takes the top card from the deck.

## Input

The first line contains single integer n (1 ≤ n ≤ 100,000) — the number of cards in the deck.

The second line contains a sequence of n integers a1, a2, ..., an (1 ≤ ai ≤ 100,000), where ai is the number written on the i-th from top card in the deck.

## Output

Print the total number of times Vasily takes the top card from the deck.

## Examples

### Example 1

**Input:**
```
4
6 3 1 2
```

**Output:**
```
7
```

### Example 2

**Input:**
```
1
1000
```

**Output:**
```
1
```

### Example 3

**Input:**
```
7
3 3 3 3 3 3 3
```

**Output:**
```
7
```

## Note

In the first example Vasily at first looks at the card with number 6 on it, puts it under the deck, then on the card with number 3, puts it under the deck, and then on the card with number 1. He places away the card with 1, because the number written on it is the minimum among the remaining cards. After that the cards from top to bottom are [2, 6, 3]. Then Vasily looks at the top card with number 2 and puts it away. After that the cards from top to bottom are [6, 3]. Then Vasily looks at card 6, puts it under the deck, then at card 3 and puts it away. Then there is only one card with number 6 on it, and Vasily looks at it and puts it away. Thus, in total Vasily looks at 7 cards.


### ideas
1. 反过来考虑，在去到最小的牌前，需要经过多少次？这个第一次很好计算，然后再计算下一个，再下一个。可以搞