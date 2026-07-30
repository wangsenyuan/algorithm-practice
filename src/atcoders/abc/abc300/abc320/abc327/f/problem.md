# F - Apples

[Problem link](https://atcoder.jp/contests/abc327/tasks/abc327_f)

**Contest:** [AtCoder Beginner Contest 327](https://atcoder.jp/contests/abc327)

time limit: 2 sec

memory limit: 1024 MiB

score: 550 points

## Problem Statement

There are apple trees lined up on a number line, and `N` apples fall from the trees.
Specifically, for each `1 <= i <= N`, an apple falls at coordinate `X_i` at time `T_i`.

Takahashi has a basket with durability `D` and length `W`, and he can take the following
action **exactly once**.

> Choose positive integers `S` and `L`. He sets up the basket to cover the range
> `L - 0.5 <= x <= L + W - 0.5` at time `S - 0.5`, and retrieves it at time `S + D - 0.5`.
> He gets all the apples that fell into the range covered by the basket between the time it
> was set up and the time it was retrieved.

He cannot move the basket once it has been set up, nor can he set it up again once it has
been retrieved.

Find the maximum number of apples that he can get.

## Constraints

- `1 <= N <= 2 * 10^5`
- `1 <= D <= 2 * 10^5`
- `1 <= W <= 2 * 10^5`
- `1 <= T_i <= 2 * 10^5`
- `1 <= X_i <= 2 * 10^5`
- All pairs `(T_i, X_i)` are different.
- All input values are integers.

## Input

```text
N D W
T_1 X_1
T_2 X_2
...
T_N X_N
```

## Output

Print the maximum number of apples that Takahashi can get.

## Sample Input 1

```text
8 4 3
1 1
3 4
6 4
5 2
4 2
4 3
5 5
7 3
```

## Sample Output 1

```text
5
```

If Takahashi chooses `S = 3` and `L = 2`, he sets up the basket to cover
`1.5 <= x <= 4.5` from time `2.5` to `6.5`, collecting these five apples:

- `(T, X) = (3, 4)`
- `(T, X) = (6, 4)`
- `(T, X) = (5, 2)`
- `(T, X) = (4, 2)`
- `(T, X) = (4, 3)`

No choice collects six or more apples.

## Solution Summary

For a basket placed at integer position `L` and used from integer start time
`S`, an apple `(T, X)` is collected exactly when

\[
S \le T < S+D,
\qquad
L \le X < L+W.
\]

The task is therefore to find an axis-aligned `D × W` rectangle containing the
largest number of apple points in the `(time, coordinate)` plane.

### Sweep the time dimension

Instead of directly choosing the start time, sweep an equivalent end-time
coordinate. Apple `(T, X)` is active while that sweep coordinate is in

\[
[T,T+D).
\]

So it is added at time `T` and removed at time `T+D`. The active apples at one
sweep position are exactly the apples that can be collected by one valid
duration-`D` basket interval. If a sweep position would correspond to a
non-positive start time, shifting the basket start to `1` can only retain or
add apples because all fall times are positive.

Only event times `T` and `T+D` matter. The code coordinate-compresses them in
`times`, then places apple indices in `active` and `expire` buckets.

### Convert an apple into a range of basket positions

For a fixed active apple at coordinate `X`, the basket's left endpoint must
satisfy

\[
X-W+1 \le L \le X.
\]

Thus the apple contributes `+1` to every valid integer `L` in that interval.
When the apple expires, it contributes `-1` to the same interval.

The value as a function of `L` changes only at endpoints `X-W+1` and `X`, so
the code compresses both kinds of endpoints into `pos`. For each active apple,
`tr.update(l, r, 1)` range-adds its contribution; expiration performs the same
update with `-1`.

### Segment tree invariant

At every processed time event, the leaf for compressed position `L` stores the
number of currently active apples whose valid-left-endpoint interval contains
`L`. Therefore it is exactly the number of apples collected by using that
horizontal basket placement at the current sweep time.

`tree.val[0]` is the maximum over all leaves, so it gives the best basket
placement for the current time window. Taking the maximum of this value during
the sweep gives the answer.

### Correctness sketch

Every valid basket has a duration-`D` time interval and a length-`W` coordinate
interval. At the corresponding sweep event, precisely its collectable apples
are active. Each such apple adds one to the leaf representing the basket's
left endpoint, while every non-collectable apple does not contribute. Hence the
segment-tree maximum equals the best basket at that time. Sweeping all event
times covers every point where the active set can change, so the global maximum
is the optimum answer.

### Complexity

There are `O(N)` time and position endpoints. Sorting costs `O(N log N)`, and
each apple performs one insertion and one deletion, each in `O(log N)` time.
The total time complexity is `O(N log N)` and the memory complexity is `O(N)`.
