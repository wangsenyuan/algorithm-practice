# D. World of Darkraft

[Problem link](https://codeforces.com/problemset/problem/138/D)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

Recently Roma has become the happy owner of a new game World of Darkraft. This
game combines elements of virtually all known genres, and on one of the later
stages of the game Roma faced difficulties solving a puzzle.

In this part Roma fights with a cunning enemy magician. The battle takes place
on a rectangular field plaid `n x m`. Each cell contains one magical character:
`L`, `R` or `X`. Initially all the squares of the field are active.

The players, Roma and enemy magician, take turns. Roma makes the first move.
During a move a player selects one of the active cells. Then depending on the
image in the character in the cell one of the following actions takes place:

- `L` -- magical waves radiate from the cell to the left downwards and to the
  right upwards along diagonal paths. All cells on the path of the waves,
  including the selected cell, become inactive. The waves continue until the
  next inactive cell or to the edge of the field if there are no inactive cells
  on the way.
- `R` -- the magical waves radiate to the left upwards and to the right
  downwards.
- `X` -- the magical waves radiate in all four diagonal directions.

If the next player cannot make a move, that is, all cells are inactive, he
loses.

Roma has been trying to defeat the computer opponent for three days but he just
keeps losing. He asks you to help him and determine whether it is guaranteed
that he can beat the opponent, or he will have to hack the game.

## Input

The first line contains two space-separated integers `n` and `m`
(`1 <= n, m <= 20`).

The next `n` lines contain `m` characters describing the playing field: the
`j`-th character of the `i`-th line equals the magical character of the
corresponding field square.

## Output

On the first line print `WIN` if Roma can win, or `LOSE` if it is impossible to
win considering that the opponent plays optimally.

## Examples

### Input 1

```text
2 2
RL
LR
```

### Output 1

```text
LOSE
```

### Input 2

```text
2 2
RR
RR
```

### Output 2

```text
WIN
```

## Note

In the first test, each move makes one diagonal line of the square inactive, so
it is guaranteed that Roma loses after two moves.

In the second test, there are three variants of making a move: to finish off the
main diagonal line or any of the squares that are left. That means that after
three moves the game stops and Roma wins.


## Solution

This is an impartial game, so we use Sprague-Grundy numbers.

### Split by Parity

Every move removes cells only along diagonals. Moving along a diagonal preserves
the parity of `row + column`, so cells with even `row + column` and cells with
odd `row + column` never interact.

Therefore, the original game is the xor-sum of two independent games:

- the game on even cells;
- the game on odd cells.

If their Grundy numbers are `g_even` and `g_odd`, then the whole position is
winning iff:

```text
g_even xor g_odd != 0
```

### Shape of a Remaining Component

Consider only one parity class, for example the even cells.

During the game, every connected remaining piece can be represented as the
intersection of the board rectangle with four diagonal half-planes. There is one
boundary for each diagonal direction.

So a piece can be described by four diagonal bounds:

```text
lower/upper bound on one diagonal family
lower/upper bound on the other diagonal family
```

Equivalently, each cell belongs to two diagonal lines:

- one going right-up;
- one going right-down.

The four bounds say which range of those two diagonal indices is still active.
Enumerating all possible quadruples of bounds enumerates all possible pieces
that can appear during the game.

### Grundy DP for One Piece

Let `grundy(piece)` be the Grundy number of a remaining piece.

To compute it, try every active cell inside this piece as the next move.
Depending on the character in that cell:

- `L` cuts/removes along the left-down and right-up diagonal directions;
- `R` cuts/removes along the left-up and right-down diagonal directions;
- `X` cuts/removes in all four diagonal directions.

After making the move, the current piece is split into smaller independent
pieces. The Grundy number of that move is the xor of the Grundy numbers of all
pieces that remain.

Collect all such xor values, then:

```text
grundy(piece) = mex(all reachable xor values)
```

All resulting pieces are strictly smaller than the current one, so the values
can be computed with dynamic programming over the four diagonal bounds.

To iterate over the cells of a piece, iterate over the two diagonal indices that
identify a cell, and check whether the corresponding board cell is valid and
belongs to the current parity class.

### Final Decision

Compute the Grundy number for the full board restricted to even cells and the
full board restricted to odd cells.

Then:

```text
if g_even xor g_odd == 0:
    LOSE
else:
    WIN
```

The number of possible pieces is polynomial in the number of diagonals, and for
each piece we scan possible cells and compute a mex over reachable states.
