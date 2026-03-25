It is so boring in the summer holiday, isn't it? So Alice and Bob have invented a new game to play. The rules are as follows. First, they get a set of `n` distinct integers. And then they take turns to make the following moves. During each move, either Alice or Bob (the player whose turn is the current) can choose two distinct integers `x` and `y` from the set, such that the set doesn't contain their absolute difference `|x - y|`. Then this player adds integer `|x - y|` to the set (so, the size of the set increases by one).

If the current player has no valid move, he (or she) loses the game. The question is who will finally win the game if both players play optimally. Remember that Alice always moves first.

## Input

The first line contains an integer `n` (`2 <= n <= 100`) - the initial number of elements in the set.

The second line contains `n` distinct space-separated integers `a1, a2, ..., an` (`1 <= ai <= 10^9`) - the elements of the set.

## Output

Print a single line with the winner's name. If Alice wins print `Alice`, otherwise print `Bob` (without quotes).

## Examples

### Example 1

**Input**

```
2
2 3
```

**Output**

```
Alice
```

### Example 2

**Input**

```
2
5 3
```

**Output**

```
Alice
```

### Example 3

**Input**

```
3
5 6 7
```

**Output**

```
Bob
```

## Note

Consider the first test sample. Alice moves first, and the only move she can do is to choose `2` and `3`, then to add `1` to the set. Next Bob moves, there is no valid move anymore, so the winner is Alice.

## Ideas (scratch)

- Adding `|x - y|` to the set: what invariant stays the same?
- Example: for `(x, y) = (87, 17)`, you can add `70, 53, 36, 19, 2, 15, 13, 11, 9, 7, 5, 3, 1` (each step adds a value not yet in the set, and new values tend to be smaller than `x`).
- Each added number is new to the set and is typically smaller than the larger of the pair.
- If `1` is already in the set, you can often generate more numbers; if not, you may still be able to reach `1`.
- The outcome may depend only on the maximum element (and related structure).
- Consider `gcd` of the set: the smallest value you can reach is related to the gcd.
