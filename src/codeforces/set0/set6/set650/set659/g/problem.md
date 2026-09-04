# G. Fence Divercity

[Problem link](https://codeforces.com/problemset/problem/659/G)

**Contest:** [Codeforces Round #346 (Div. 2)](https://codeforces.com/contest/659)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Long ago, Vasily built a good fence at his country house. Vasily calls a fence good if it is a series of `n` consecutively fastened vertical boards of centimeter width, the height of each in centimeters being a positive integer. The height of the `i`-th board from the left is `h_i`.

Today Vasily decided to change the design by cutting out exactly one connected top part so that the fence remains good. The cut part must consist only of the upper parts of the boards, and adjacent cut pieces must be interconnected (they share a non-zero length before the cut).

Two ways to cut a part are distinct if the remaining fences differ in the height of at least one board.

Count the number of ways to cut exactly one such connected part, modulo `1 000 000 007` (`10^9 + 7`).

## Input

The first line contains integer `n` (`1 <= n <= 10^6`) — the number of boards.

The second line contains `n` space-separated integers `h_1, h_2, ..., h_n` (`1 <= h_i <= 10^9`) — the heights of the boards from left to right.

## Output

Print the remainder after dividing the number of valid cuts by `1 000 000 007`.

## Examples

### Input

```text
2
1 1
```

### Output

```text
0
```

### Input

```text
3
3 4 2
```

### Output

```text
13
```

### Note

From the first fence it is impossible to cut exactly one piece and leave a good fence (every board has height 1, so any cut would drop a height to 0).

The second sample has 13 distinct remaining fences after a valid connected top cut.

## Status

I/O and official samples are in place. `solve` is left as a TODO.
