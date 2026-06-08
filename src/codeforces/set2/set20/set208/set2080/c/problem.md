# C. Card Flip

[Problem link](https://codeforces.com/problemset/problem/2080/C)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

There are `n` double-sided cards and `m` single-sided cards.

- The `i`-th double-sided card has number `a_i` on the front side and number
  `b_i` on the back side.
- The `i`-th single-sided card has number `c_i` on its front side.

All numbers written on all card sides are distinct. Initially, every card lies
on the table with its front side up.

On each turn, the current player performs exactly one action:

1. Remove from the table the card with the smallest visible number.
2. If the card with the smallest visible number is double-sided and is currently
   front side up, flip it.

The player who removes the last card from the table wins. Petya moves first.
Determine the winner if both players play optimally.

## Input

The first line contains two integers `n` and `m`
(`1 <= n, m <= 500000`) — the number of double-sided and single-sided cards.

The second line contains `n` integers `a_1, a_2, ..., a_n`
(`1 <= a_i <= 2n + m`) — the front numbers of the double-sided cards.

The third line contains `n` integers `b_1, b_2, ..., b_n`
(`1 <= b_i <= 2n + m`) — the back numbers of the double-sided cards.

The fourth line contains `m` integers `c_1, c_2, ..., c_m`
(`1 <= c_i <= 2n + m`) — the numbers of the single-sided cards.

It is guaranteed that every number from `1` to `2n + m` appears exactly once in
one of the arrays `a`, `b`, or `c`.

## Output

Print `First` if Petya wins, and `Second` if Vasya wins.

## Examples

### Input 1

```text
2 1
5 3
1 2
4
```

### Output 1

```text
First
```

### Input 2

```text
1 2
2
3
4 1
```

### Output 2

```text
Second
```

## Note

In the first example, the visible cards are initially `3`, `4`, and `5`. Petya
removes `3`, Vasya removes `4`, and Petya removes `5`, so Petya wins.

In the second example, visible cards are initially `1`, `2`, and `4`. Petya must
remove the single-sided card `1`. Then Vasya can flip the card showing `2`,
revealing `3`; after Petya removes `3`, Vasya removes `4` and wins.

## Solution explanation

The game is only interesting when the smallest visible card is a double-sided
card that is still front side up. At that moment, the current player can choose:

- remove it, decreasing the number of remaining cards by `1`;
- flip it, keeping the same physical card on the table but changing the visible
  number.

So a double-sided card can let the current player change the parity of the
number of future removals. Since the winner is determined by who removes the
last card, controlling this parity is the key.

### Deciding card

Consider double-sided cards as intervals from their front value to their back
value:

```text
a -> b
```

Initially only `a` is visible. If the card is flipped, `b` becomes visible.

Sort all double-sided cards by front value `a` in decreasing order. We look for
the important, or "deciding", card.

Start from the double-sided card with the largest front value. Call it `last`.
Then scan the remaining double-sided cards in decreasing `a`.

For a card `cur = (a, b)` before `last`, there are two cases:

1. `b > last.a`

   If `cur` is flipped, its back value still appears after `last.a`. Whether
   `cur` is removed or flipped does not change who gets control before the
   deciding card `last`. This card is not important.

2. `b < last.a`

   If `cur` is flipped, its back value appears before `last.a`. Now this earlier
   card can change the parity before reaching `last`, so `cur` itself becomes
   the new deciding card.

That is why the code does:

```go
last := 0
for i := 1; i < len(cards); i++ {
    if cards[i].b < cards[last].a {
        last = i
    }
}
```

After this scan, `cards[last]` is the first double-sided card that can actually
decide the game under optimal play.

### Counting cards before the deciding card

Let the deciding card have front value:

```text
x = cards[last].a
```

Before the deciding card is reached, some cards must be removed no matter what.
They are:

1. Every single-sided card with value `< x`.
2. Every double-sided card that appears before `last` in the sorted-by-decreasing
   order after the deciding chain is fixed. In the code this count is:

```go
len(cards) - 1 - last
```

So:

```go
before := number of single cards with c < x
before += len(cards) - 1 - last
```

This `before` is the number of forced removals before the deciding move.

### Parity

Petya moves first.

If `before` is even, then Petya is the player who reaches the deciding card.
The player who reaches the deciding card can choose remove/flip in the way that
makes them win, so Petya wins.

If `before` is odd, Vasya reaches the deciding card and can force the win.

Therefore:

```go
return before%2 == 0
```

### Complexity

Sorting the `n` double-sided cards costs:

```text
O(n log n)
```

The scan over double-sided cards and the count over single-sided cards are
linear.

Overall complexity:

```text
O(n log n + m)
```

Space complexity:

```text
O(n + m)
```
