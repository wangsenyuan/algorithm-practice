# F. New Year and Cleaning

[Problem link](https://codeforces.com/problemset/problem/611/F)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Limak is a little polar bear. His parents told him to clean a house before the New
Year's Eve. Their house is a rectangular grid with `h` rows and `w` columns. Each
cell is an empty square.

He is a little bear and thus he can't clean a house by himself. Instead, he is
going to use a cleaning robot.

A cleaning robot has a built-in pattern of `n` moves, defined by a string of the
length `n`. A single move (character) moves a robot to one of four adjacent
cells. Each character is one of the following four: `'U'` (up), `'D'` (down),
`'L'` (left), `'R'` (right). One move takes one minute.

A cleaning robot must be placed and started in some cell. Then it repeats its
pattern of moves till it hits a wall (one of four borders of a house). After
hitting a wall it can be placed and used again.

Limak isn't sure if placing a cleaning robot in one cell will be enough. Thus,
he is going to start it `w * h` times, one time in each cell. Maybe some cells
will be cleaned more than once but who cares?

Limak asks you one question. How much time will it take to clean a house? Find
and print the number of minutes modulo `10^9 + 7`. It's also possible that a
cleaning robot will never stop — then print `"-1"` (without the quotes) instead.

Placing and starting a robot takes no time, however, you must count a move when
robot hits a wall. Take a look into samples for further clarification.

## Input

The first line contains three integers `n`, `h` and `w`
(`1 <= n, h, w <= 500000`) — the length of the pattern, the number of rows and
the number of columns, respectively.

The second line contains a string of length `n` — the pattern of `n` moves. Each
character is one of uppercase letters `'U'`, `'D'`, `'L'` or `'R'`.

## Output

Print one line with the answer.

If a cleaning robot will never stop, print `"-1"` (without the quotes).
Otherwise, print the number of minutes it will take to clean a house modulo
`10^9 + 7`.

## Examples

### Input 1

```text
1 10 2
R
```

### Output 1

```text
30
```

### Input 2

```text
3 4 6
RUL
```

### Output 2

```text
134
```

### Input 3

```text
4 1 500000
RLRL
```

### Output 3

```text
-1
```

## Note

In the first sample house is a grid with 10 rows and 2 columns. Starting a robot
anywhere in the second column will result in only one move (thus, one minute of
cleaning) in which robot will hit a wall — he tried to go right but there is no
third column. Starting a robot anywhere in the first column will result in two
moves. The total number of minutes is `10 * 1 + 10 * 2 = 30`.

In the second sample a started robot will try to move `"RULRULRULR..."`. For
example, for the leftmost cell in the second row robot will make 5 moves before
it stops because of hitting an upper wall.

## Solution

For every starting cell, the robot follows the same infinite sequence of moves.
The only difference is where the relative path is placed inside the grid.

So instead of simulating from every cell, simulate the path once as relative
coordinates from the starting cell.

Let the relative position after some moves be `(x, y)`. The code tracks the
bounding box of all relative positions visited so far:

```text
loX <= x <= hiX
loY <= y <= hiY
```

A starting cell is still valid exactly when this whole bounding box can fit
inside the `w * h` grid.

The number of valid horizontal placements is:

```text
w - (hiX - loX)
```

and the number of valid vertical placements is:

```text
h - (hiY - loY)
```

The implementation stores these two values as:

```go
dim[0] = number of still-valid horizontal placements
dim[1] = number of still-valid vertical placements
```

Initially the bounding box contains only `(0, 0)`, so:

```go
dim = []int{w, h}
```

### Counting Cells That Stop at One Move

When a move expands the horizontal bounding box by one, some starting positions
become invalid for the first time. For each vertical placement, exactly one
horizontal placement is lost, so the number of starts that stop at this move is:

```text
dim[1]
```

Similarly, when a move expands the vertical bounding box by one, the number of
starts that stop is:

```text
dim[0]
```

That is why `play` does:

```go
ans = (ans + (moves+1)%mod*dim[i^1]) % mod
dim[i]--
```

Here:

- `i = 0` means the horizontal range expanded, so multiply by `dim[1]`;
- `i = 1` means the vertical range expanded, so multiply by `dim[0]`;
- `moves + 1` is the time when those starts hit a wall.

Once either dimension becomes zero, all starting cells have already stopped and
the answer is complete.

### First Two Pattern Runs

The code first simulates the first full pattern normally:

```go
for moves := 0; inGame() && moves < n; moves++ {
    play(moves, false)
}
```

Then it simulates the second full pattern normally as well. During this second
run, it records which pattern indices actually expand the bounding box:

```go
if play(moves, false) {
    todo = append(todo, moves%n)
}
```

After two full periods, the useful expansion positions have stabilized. Later
periods can only expand the same sides at those recorded pattern indices.

### Fast Forwarding Later Periods

For later repetitions, the code does not need to track the exact relative
position. It only needs to apply the same kind of boundary expansion.

That is the purpose of `cheating = true` in `play`.

For example:

```go
case 'L':
    if cheating {
        cur[0] = hi[0] + 1
    }
```

This directly forces the current point to be just outside the known horizontal
maximum, causing one more horizontal expansion. The other directions are handled
symmetrically:

```text
L -> expand right side
R -> expand left side
U -> expand top side
D -> expand bottom side
```

The loop:

```go
for k := 2; inGame(); k++ {
    for _, moves := range todo {
        play(moves+k*n, true)
    }
}
```

keeps applying only those useful expansions until every starting cell has
stopped.

### Infinite Case

If after a whole pattern the robot returns to relative position `(0, 0)` while
some starting cells are still alive, then the same relative path will repeat
forever.

Those still-alive starts will never hit a wall, so the answer is `-1`.

The code checks this at period boundaries:

```go
if moves%n == 0 && cur[0] == 0 && cur[1] == 0 {
    return -1
}
```

### Complexity

The first two pattern runs cost `O(n)`.

After that, every useful fast-forward step decreases either `dim[0]` or
`dim[1]`, so there are at most `h + w` such steps.

Therefore the total complexity is:

```text
O(n + h + w)
```
