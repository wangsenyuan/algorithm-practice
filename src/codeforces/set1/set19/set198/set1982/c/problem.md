On another boring day, Egor got bored and decided to do something. But since he has no friends, he came up with a game to play.

Egor has a deck of 𝑛 cards, the 𝑖-th card from the top has a number 𝑎ᵢ written on it. Egor wants to play a certain number of rounds until the cards run out. In each round, he takes a non-zero number of cards from the top of the deck and finishes the round. If the sum of the numbers on the cards collected during the round is between 𝑙 and 𝑟, inclusive, the round is won; otherwise, it is lost.

Egor knows by heart the order of the cards. Help Egor determine the maximum number of rounds he can win in such a game. Note that Egor is not required to win rounds consecutively.

## Input

Each test consists of several test cases. The first line contains an integer 𝑡 (1 ≤ 𝑡 ≤ 10⁴) — the number of test cases. This is followed by a description of the test cases.

The first line of each test case contains three integers 𝑛, 𝑙, and 𝑟 (1 ≤ 𝑛 ≤ 10⁵, 1 ≤ 𝑙 ≤ 𝑟 ≤ 10⁹).

The second line of each test case contains 𝑛 integers 𝑎₁, 𝑎₂, …, 𝑎ₙ (1 ≤ 𝑎ᵢ ≤ 10⁹) — the numbers on the cards from top to bottom.

It is guaranteed that the sum of 𝑛 for all test cases does not exceed 2⋅10⁵.

## Output

For each test case, output a single number — the maximum number of rounds Egor can win.

## Example

**Input:**
```
8
5 3 10
2 1 11 3 7
10 1 5
17 8 12 11 7 11 21 13 10 8
3 4 5
3 4 2
8 12 25
10 7 5 13 8 9 12 7
2 3 3
5 2
9 7 9
2 10 5 1 3 7 6 2 3
1 8 10
9
5 5 6
1 4 2 6 4
```

**Output:**
```
3
0
1
4
0
3
1
2
```

## Note

In the first test case, Egor can win 3 rounds:

- In the first round, take the top 2 cards with values 2 and 1 and win, as their sum is 3. After this, the deck will look like this: [11, 3, 7].
- In the second round, take the top card and lose, as its value 11 is greater than 𝑟 = 10. After this, the deck will look like this: [3, 7].
- In the third round, take the top card with value 3 and win. After this, the deck will look like this: [7].
- After this, in the fourth round, Egor only has to take the last card in the deck with value 7 and win again.

In the second test case, Egor cannot win any rounds, no matter how hard he tries.

In the third test case, you can take one card in each round, then the first and third rounds will be losing, and the second round will be winning.

In the fourth test case, you can take two cards in each round and always win.
