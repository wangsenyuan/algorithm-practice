# C. Rabbits

[Problem link](https://codeforces.com/problemset/problem/2147/C)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: stdin

output: stdout

You have `n` flower pots arranged in a line numbered from `1` to `n` left to right.
Some of the pots contain flowers, while others are empty. You are given a binary
string `s` describing which pots contain flowers (`s_i = 1`) and which are empty
(`s_i = 0`). You also have some rabbits, and you want to take a nice picture of
rabbits and flowers. You want to put rabbits in every empty pot (`s_i = 0`), and
for each rabbit, you can put it looking either to the left or to the right.
Unfortunately, the rabbits are quite naughty, and they will try to jump, which
will ruin the picture.

Each rabbit will prepare to jump into the next pot in the direction they are
looking, but they won't jump if there is a rabbit in that pot already or if there
is another rabbit that prepares to jump into the same pot from the opposite side.
Rabbits won't jump out of the borders (a rabbit at pot `1` looking to the left
won't jump, same for a rabbit looking to the right at pot `n`).

Your goal is to choose the directions of the rabbits so that they never jump,
allowing you to take your time to take the picture. You need to determine if
there is a valid arrangement of rabbits such that no rabbit ever jumps.

## Solution

The useful way to look at the problem is to ignore flower pots unless some
rabbit points into them. A rabbit is safe if it points to the border, to another
rabbit, or to a flower that is also targeted by a rabbit from the other side.
So a flower can only be used as a target when it is a single `1` between rabbits
on both sides; two rabbits then prepare to jump into the same pot and both stay.

Scan the string by equal-character runs. Long zero-runs are easy: if a run of
`0`s has length at least `2`, the rabbits inside can be oriented so that the run
is self-contained, and one endpoint can still be chosen later if it needs to
help protect a neighboring single flower. This is the flexible state.

The only hard case is a single zero. A single rabbit may be forced by its left
side:

- If it is after a run of at least two `1`s, it cannot look left, because the
  adjacent flower has no rabbit on the other side to collide with. Therefore it
  must look right.
- If it is after a single `1`, then it may need to pair with the previous rabbit
  across that single flower. When the previous single rabbit was forced to look
  right, this rabbit must look left. Otherwise, this rabbit may become the one
  that looks right for the next boundary.
- If the previous zero-run was long, there is still enough freedom, so the
  current single zero does not create an immediate contradiction.

The implementation keeps this information in `state`:

- `state = 0`: the latest relevant single rabbit is fixed to the left, so there
  is no pending right-looking rabbit.
- `state = 1`: the latest single rabbit is forced to look right. The next
  flower-run must be only one flower and must have a rabbit on the other side.
- `state = 2`: the latest zero-run is flexible.

When the scan reaches a run of `1`s, only `state = 1` can be dangerous. That
state means the previous rabbit is forced to point into the current flower-run.
This is valid only if the current run is exactly one `1` and is not at the end of
the string, because then a rabbit on the right side can point left and collide
with it. If the run has length at least `2`, or if it reaches the right border,
there is no possible opposite rabbit, so the answer is `NO`.

All transitions depend only on the previous run and this small state, so the
algorithm is linear in the string length. The total complexity over all test
cases is `O(sum n)`, and the extra memory is `O(number of runs)` in this
implementation.
