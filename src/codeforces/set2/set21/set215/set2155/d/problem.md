# D. Batteries

https://codeforces.com/problemset/problem/2155/D

**Time Limit:** 2 seconds

**Memory Limit:** 256 megabytes

## Problem

This is an interactive problem, and hacks are disabled.

There are `n` batteries numbered from `1` to `n`
(`2 <= n <= 40`). Some batteries work and the rest do not. Let `a` be the
number of batteries that work; it is guaranteed that `a >= 2`.

You are given `n`, but not `a`.

A flashlight can hold two batteries. It turns on only when both inserted
batteries work. You may choose two batteries and test them in the flashlight.

Your goal is to find a pair of working batteries while using at most:

```text
floor(n^2 / a)
```

trials.

The interactor is adaptive. This means the set of working batteries is not fixed
in advance and may change during the interaction, but after every answer there
must still exist some configuration of `a` working batteries consistent with all
information received so far.

## Solution

Put the batteries on a circle in order `1, 2, ..., n`, with battery `1` after
battery `n`.

Suppose there are `a` working batteries. Look only at these `a` batteries on the
circle, and sort them in circular order:

```text
b_1, b_2, ..., b_a
```

Consider the circular gaps between consecutive working batteries. These `a`
gaps cover the whole circle, so their sum is exactly `n`. Therefore, by the
pigeonhole principle, at least one gap is at most:

```text
floor(n / a)
```

So there exists a pair of working batteries whose clockwise distance is at most
`floor(n / a)`.

The value `a` is unknown, so query by increasing circular distance:

```text
for d = 1, 2, ..., n - 1:
    for i = 1, 2, ..., n:
        query i and i + d on the circle
```

When `i + d > n`, wrap around by subtracting `n`.

If the answer is `1`, both queried batteries work, and we are done.

Why the query limit is satisfied:

For the real value of `a`, a working pair exists at some distance
`d <= floor(n / a)`. Before finishing distance `d`, the algorithm has made at
most:

```text
n * floor(n / a)
```

queries. Since:

```text
n * floor(n / a) <= floor(n^2 / a)
```

the required limit is not exceeded.

This also handles the adaptive interactor. After all queries with distances
`1..floor(n/a)` return `0`, no set of `a` working batteries can remain
consistent with the answers, because every possible configuration must contain
a close consecutive working pair. Thus the interactor must answer `1` before
that point.
