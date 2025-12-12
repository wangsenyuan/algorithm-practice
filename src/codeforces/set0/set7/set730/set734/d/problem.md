# Problem Description

Anton likes to play chess. Also, he likes to do programming. That is why he decided to write the program that plays chess. However, he finds the game on 8x8 board too simple, he uses an infinite one instead.

The first task he faced is to check whether the king is in check. Anton doesn't know how to implement this so he asks you to help.

Consider that an infinite chess board contains one white king and the number of black pieces. There are only rooks, bishops and queens, as the other pieces are not supported yet. The white king is said to be in check if at least one black piece can reach the cell with the king in one move.

Help Anton and write the program that for the given position determines whether the white king is in check.

## Reminder on how chess pieces move

- **Bishop** moves any number of cells diagonally, but it can't "leap" over the occupied cells.
- **Rook** moves any number of cells horizontally or vertically, but it also can't "leap" over the occupied cells.
- **Queen** is able to move any number of cells horizontally, vertically or diagonally, but it also can't "leap".

## Input

The first line of the input contains a single integer n (1 ≤ n ≤ 500 000) — the number of black pieces.

The second line contains two integers x0 and y0 (-10^9 ≤ x0, y0 ≤ 10^9) — coordinates of the white king.

Then follow n lines, each of them contains a character and two integers xi and yi (-10^9 ≤ xi, yi ≤ 10^9) — type of the i-th piece and its position. Character 'B' stands for the bishop, 'R' for the rook and 'Q' for the queen. It's guaranteed that no two pieces occupy the same position.

## Output

The only line of the output should contain "YES" (without quotes) if the white king is in check and "NO" (without quotes) otherwise.

## Examples

### Example 1

**Input:**

```text
2
4 2
R 1 1
B 1 5
```

**Output:**

```text
YES
```

### Example 2

**Input:**

```text
2
4 2
R 3 3
B 1 5
```

**Output:**

```text
NO
```
