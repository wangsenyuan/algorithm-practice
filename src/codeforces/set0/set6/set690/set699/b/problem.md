# B. One Bomb

[Problem link](https://codeforces.com/problemset/problem/699/B)

**Contest:** [Codeforces Round 363 (Div. 2)](https://codeforces.com/contest/699)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

A depot is an `n` by `m` grid. Each cell is empty (`.`) or a wall (`*`).

You have one bomb. Placed at cell `(x, y)`, it wipes out every wall in row `x` and every wall in column `y`. The cell may be empty or a wall.

Decide whether one bomb can wipe out every wall. If several cells work, any of them is acceptable.

## Input

The first line contains two integers `n` and `m` (`1 <= n, m <= 1000`).

Each of the next `n` lines contains `m` characters `.` or `*`. The `j`-th character of the `i`-th line is cell `(i, j)`.

## Output

If no placement wipes out every wall, print `NO`.

Otherwise print `YES`, and on the next line the 1-based coordinates `x y` of a valid cell.

## Example

### Input 1

```text
3 4
.*..
....
.*..
```

### Output 1

```text
YES
1 2
```

### Input 2

```text
3 3
..*
.*.
*..
```

### Output 2

```text
NO
```

### Input 3

```text
6 5
..*..
..*..
*****
..*..
..*..
..*..
```

### Output 3

```text
YES
3 3
```

## Status

I/O and official sample cases are in place. `solve` is left as a TODO.
