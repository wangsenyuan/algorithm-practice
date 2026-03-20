# Problem

There are two decks of cards on the table. Some cards are face up, some face down. You want to merge them into one deck where every card is face down. This happens in two stages.

## Stage 1 — merge

Merge the two decks so that within each original deck, relative order is preserved: if card `i` is above card `j` in the same deck before the merge, then after the merge card `i` is still above card `j`.

## Stage 2 — turns

On the merged deck, you may repeat the **turn** operation: take some number of cards from the top, flip all of them (face up ↔ face down), and put them back on top **in reverse order** (the card that was bottom among the taken cards becomes the new top).

Your task is to choose a merge order for stage 1 and a sequence of turn operations for stage 2 so that all cards end face down, using the **minimum** possible number of turns.

## Input

- First line: integer `n` (`1 <= n <= 10^5`) — number of cards in the first deck.
- Second line: `n` integers `a_1, a_2, ..., a_n` (`0 <= a_i <= 1`). `a_i = 0` means face down, `a_i = 1` means face up. Order is from top to bottom.
- Third line: integer `m` (`1 <= m <= 10^5`) — number of cards in the second deck.
- Fourth line: `m` integers `b_1, b_2, ..., b_m` (`0 <= b_i <= 1`), same meaning, top to bottom.

## Output

- **Line 1:** `n + m` space-separated integers — the card indices from top to bottom after stage 1. Cards from the first deck are numbered `1` to `n` (top to bottom in that deck). Cards from the second deck are numbered `n + 1` to `n + m` (top to bottom in that deck).
- **Line 2:** integer `x` — the minimum number of turn operations.
- **Line 3:** `x` integers `c_1, c_2, ..., c_x` (`1 <= c_i <= n + m`) — each `c_i` is how many cards to take from the top in that turn.

If several solutions are optimal, print any. It is guaranteed that the minimum number of operations does not exceed `6 * 10^5`.

## Examples

### Example 1

Input

```text
3
1 0 1
4
1 1 1 1
```

Output

```text
1 4 5 6 7 2 3
3
5 6 7
```

### Example 2

Input

```text
5
1 1 1 1 1
5
0 1 0 1 0
```

Output

```text
6 1 2 3 4 5 7 8 9 10
4
1 7 8 9
```

## ideas
1. 只考虑a的话，目前都是face down的情况，遇到faceup时，必须先操作1次，编程face up
2. 然后继续。换句话说，操作的次数 = a[i] != a[i+1]的位置，
3. 如果操作完以后时faceup得，需要一次额外的操作，进行face down
4. 两遍合并，如果前面是face down, 那么就一直合并a/b中face down的，直到遇到两遍都是face up 的情况
5. 然后操作反转
6. 是个贪心的过程