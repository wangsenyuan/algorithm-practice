# B - All Minus (AtCoder ARC218)

**Limits:** 2 sec / 1024 MiB  
**Score:** 400 points

Source: [https://atcoder.jp/contests/arc218/tasks/arc218_b](https://atcoder.jp/contests/arc218/tasks/arc218_b)

## Problem Statement

There are `N` non-negative integers `A_1, A_2, ..., A_N` written on a blackboard.

Alice and Bob play a game. Starting with Alice, they alternately perform the following operation. The player who reduces the number of integers written on the blackboard to `0` wins.

- Let `m` be the minimum non-negative integer currently written on the blackboard.
  - If `m > 0`, choose a positive integer `x` between `1` and `m`, inclusive. Replace every integer written on the blackboard with its current value minus `x`.
  - If `m = 0`, erase one or more of the `0`s written on the blackboard.

Determine who wins when both players play optimally.

You are given `T` test cases; solve each of them.

## Constraints

- `1 <= T <= 2 * 10^5`
- `1 <= N <= 2 * 10^5`
- `0 <= A_i <= 10^9`
- The sum of `N` over all test cases is at most `2 * 10^5`.
- All input values are integers.

## Solution

Sort the numbers:

```text
A_1 <= A_2 <= ... <= A_N
```

The key point is that the operation `m > 0` subtracts the same `x` from every remaining number.  Therefore, while no zero is erased, all differences between numbers stay unchanged.

So the game can be viewed by distinct value groups.

Let the sorted distinct values be:

```text
v_1 < v_2 < ... < v_k
```

and let `cnt_i` be the number of occurrences of `v_i`.

After all groups with values less than `v_i` have disappeared, the next group becomes the minimum.  Its current value is not `v_i` itself, but the gap:

```text
d_i = v_i - v_{i-1}
```

where `v_0 = 0`.

This is because all previous subtraction moves have already subtracted exactly `v_{i-1}` from every still remaining number.

Thus the game is a sequence of phases:

```text
gap d_i, then erase cnt_i zeros
```

### When the Gap Is `1`

If the current minimum is positive and equal to `1`, the player has no choice:

```text
x = 1
```

After this move, the current group becomes zero, and the turn passes to the opponent.

So a gap of `1` only flips the turn.

Then the opponent must deal with `cnt_i` zeros.

- If `cnt_i = 1` and there are more positive groups after it, the opponent is forced to erase that one zero, and the turn flips again.
- Otherwise, the opponent can win immediately:
  - if this is the last group, erasing all zeros empties the board;
  - if `cnt_i > 1`, the opponent has enough freedom to choose how many zeros to erase, so they can force the favorable parity.

This is why the implementation does:

```go
id ^= 1
if i-j > 1 || i == n {
    return player[id]
}
id ^= 1
```

for `diff == 1`.

### When the Gap Is Greater Than `1`

If the current gap is:

```text
d_i > 1
```

then the player to move has a choice.

They may choose:

- `x = d_i`, making the current group zero immediately;
- or `x = d_i - 1`, leaving the current minimum equal to `1`.

These two choices differ by exactly one forced turn before the zero-erasing phase.  Therefore, the current player can choose which player faces the next critical zero phase.

So the current player can force a win immediately once a gap greater than `1` appears.

This is why the code returns the current player when:

```go
diff > 1
```

### Initial Zeros

If the smallest value is already `0`, then the first phase is a zero-erasing phase.

If there is more than one zero, or if all numbers are zero, the current player can win immediately by erasing a suitable positive number of zeros.

If there is exactly one zero and some positive number remains, the current player must erase that zero, and the turn flips.

This is the `diff == 0` branch in the implementation.

### Algorithm

1. Sort `A`.
2. Scan equal-value groups.
3. Maintain `id`, the player whose turn it is:

```text
id = 0 -> Alice
id = 1 -> Bob
```

4. For each group `[j, i)`:
   - compute `diff = current_value - previous_distinct_value`;
   - apply the rules above.

The first branch that determines the winner returns immediately.

### Complexity

Sorting dominates:

```text
O(N log N)
```

per test case, with total `N <= 2 * 10^5`, so this is fast enough.
