# P4390 [BalkanOI 2007] Mokia 摩基亚

[Problem link](https://www.luogu.com.cn/problem/P4390)

time limit per test: 2.00s

memory limit per test: 125.00MB

Mokia, a Moldova mobile-phone company, built a user-location system. It can answer
where a given user is, and also how many users lie in a given region.

The world is a `W x W` square of `1 x 1` cells. Cell `(x, y)` has `1 <= x, y <= W`.
The matrix starts at all zeros. Count the users in axis-aligned rectangles.

## Input

The input is a sequence of commands, one per line:

| Command | Arguments | Meaning |
| --- | --- | --- |
| `0` | `w` | Initialize an all-zero `w x w` matrix. Appears once, at the start. |
| `1` | `x y a` | Add `a` users to cell `(x, y)`. `a` is a positive integer. |
| `2` | `x1 y1 x2 y2` | Query the number of users in `x1 <= x <= x2`, `y1 <= y <= y2`. |
| `3` | (none) | End the program. Appears once, at the end. |

## Output

For every command `2`, print one line with the rectangle sum.

## Sample

### Input

```text
0 4
1 2 3 3
2 1 1 3 3
1 2 2 2
2 2 2 3 4
3
```

### Output

```text
3
5
```

### Note

- `0 4` creates a `4 x 4` zero matrix.
- `1 2 3 3` adds 3 users at `(2, 3)`.
- `2 1 1 3 3` sums the square `[1, 3] x [1, 3]` and reports `3`.
- `1 2 2 2` adds 2 users at `(2, 2)`.
- `2 2 2 3 4` sums `[2, 3] x [2, 4]` and reports `5`.

## Constraints

For all tests:

- `1 <= w <= 2 * 10^6`
- `1 <= x1 <= x2 <= w`, `1 <= y1 <= y2 <= w`, `1 <= x, y <= w`
- `0 < a <= 10000`
- at most `160000` commands of type `1`
- at most `10000` commands of type `2`

## Status

I/O and the official sample are in place. `solve` is left as a TODO.
