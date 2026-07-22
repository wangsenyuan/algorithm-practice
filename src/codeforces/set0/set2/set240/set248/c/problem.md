# C. Robo-Footballer

[Problem link](https://codeforces.com/problemset/problem/248/C)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

## Problem

The football field is rectangular. The opponent's goal lies on the line
`x = 0`, between the posts `(0, y1)` and `(0, y2)`. The upper side wall is the
line `y = yw`.

The ball initially has center `(xb, yb)` and radius `r`. Robo-Wallace must kick
it toward the upper wall. After exactly one perfectly elastic bounce, the ball
must enter the goal without touching another wall or a goal post.

Find the x-coordinate `xw` of a valid aiming point on the upper wall, or output
`-1` if scoring is impossible. Any valid answer is accepted with absolute
error at most `1e-8`.

## Constraints

- `1 <= y1, y2, yw, xb, yb, r <= 10^6`
- `y1 < y2 < yw`
- `yb + r < yw`
- `2r < y2 - y1`
- The initial ball position is valid.

## Input

```text
y1 y2 yw xb yb r
```

## Output

Print `-1` if no valid shot exists. Otherwise, print the x-coordinate of a
valid aiming point.

## Examples

### Sample 1

```text
Input
4 10 13 10 3 1

Output
4.3750000000
```

### Sample 2

```text
Input
1 4 6 2 2 1

Output
-1
```

### Sample 3

```text
Input
3 10 15 17 9 2

Output
11.3333333333
```

## Solution

### Remove the ball radius

The center of the ball cannot approach the upper wall more closely than `r`.
Move that wall down by `r` and treat the ball as a point. Its effective
y-coordinate is

```text
H = yw - r.
```

The ball center must cross `x = 0` at least `r` above the lower post. Choose
the lowest candidate goal point

```text
G = (0, Y), where Y = y1 + r.
```

This choice is optimal: moving `G` upward makes the trajectory closer to the
upper post `(0, y2)`. Therefore, if this lowest trajectory hits the upper post,
every other trajectory does too.

### Unfold the reflection

Reflect the initial ball center `(xb, yb)` across the effective wall `y = H`:

```text
R = (xb, 2H - yb).
```

In the unfolded picture, the two segments before and after the bounce become
the straight line from `R` to `G`. Its intersection with `y = H` has the same
x-coordinate as the real bounce point.

Let

```text
RY = 2H - yb
D  = RY - Y.
```

By similar triangles, the required coordinate is

```text
xw = xb * (H - Y) / D.
```

### Check the upper goal post

It remains to check whether the ball touches the upper post `P = (0, y2)`.
The distance from `P` to the line through `G` and `R` is

```text
distance = xb * (y2 - Y) / sqrt(xb^2 + D^2).
```

A valid trajectory requires `distance >= r`. To avoid floating-point error in
this yes/no decision, compare the squares exactly:

```text
[xb * (y2 - Y)]^2 >= r^2 * (xb^2 + D^2).
```

The products can reach about `10^24`, so they do not fit in `int64`. The
implementation uses `math/big.Int` for this single comparison. Floating-point
arithmetic is used only for the final coordinate that the problem explicitly
allows to have a small error.

### Correctness

After shifting the upper wall, the ball center behaves like a point reflecting
from `y = H`. Reflecting the initial point across that line turns the reflected
path into the straight segment `RG`, so their intersection with `y = H`
produces equal incidence and reflection angles.

Among all permissible goal points, `G = (0, y1+r)` gives the trajectory with
maximum distance from the upper post. Thus the shot is impossible exactly when
this trajectory has distance less than `r` from that post. When it is valid,
the computed intersection reaches the goal after exactly one bounce without
touching the upper post, so the returned coordinate is correct.

### Complexity

The solution uses `O(1)` arithmetic operations and `O(1)` memory.
