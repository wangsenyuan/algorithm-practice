# D. Horse Races

https://codeforces.com/problemset/problem/95/D

**Time Limit:** 2 seconds

**Memory Limit:** 256 megabytes

## Problem

Petya likes horse racing. Horses numbered from `l` to `r` take part in a race.
To evaluate the probability of victory, Petya needs to know how many horse
numbers are nearly lucky.

A number is nearly lucky if it contains at least two lucky digits such that the
distance between their positions is at most `k`. Lucky digits are `4` and `7`.

The distance between two digits is the absolute difference between their
positions in the number.

For example, when `k = 2`, the numbers `412395497`, `404`, and
`4070400000070004007` are nearly lucky, while `4`, `4123954997`, and
`4007000040070004007` are not.

Petya has `t` intervals `[l_i, r_i]`, and the same value of `k` applies to all
of them. For each interval, count how many nearly lucky numbers it contains.
Output each answer modulo `1000000007`.

## Input

The first line contains two integers `t` and `k`
(`1 <= t, k <= 1000`) -- the number of intervals and the maximum allowed
distance between lucky digits.

Each of the next `t` lines contains two integers `l_i` and `r_i`
(`1 <= l_i <= r_i <= 10^1000`).

All numbers are written without leading zeroes. Numbers in each line are
separated by exactly one space.

## Output

Output `t` lines. For each interval, print the number of nearly lucky numbers in
that interval modulo `1000000007`.

## Example 1

```text
Input
1 2
1 100

Output
4
```

In this interval, the nearly lucky numbers are `44`, `47`, `74`, and `77`.

## Example 2

```text
Input
1 2
70 77

Output
2
```

Only `74` and `77` are nearly lucky.

## Example 3

```text
Input
2 1
1 20
80 100

Output
0
0
```


## Solution

Count the complement first.

A number is not nearly lucky iff every pair of lucky digits has distance greater
than `k`. So we count numbers whose lucky digits are sparse, then subtract that
from the number of positive integers not exceeding the bound.

For a fixed `k`, precompute:

```text
free[rem][dist]
```

where `rem` is the number of remaining positions, and `dist` is the distance
from the next position to the most recent lucky digit. The value is the number
of ways to fill the remaining positions without creating two lucky digits within
distance `k`.

`dist` is capped at `k + 1`, which also represents "there is no lucky digit
nearby". For each next digit:

- Choose a non-lucky digit: `8` choices, and the distance increases by one.
- Choose a lucky digit (`4` or `7`): `2` choices, allowed only when
  `dist > k`, then the next distance becomes `1`.

This gives:

```text
free[0][dist] = 1
free[rem][dist] =
    8 * free[rem-1][min(k+1, dist+1)]
  + 2 * free[rem-1][1]   if dist > k
```

To count sparse-lucky numbers `<= s`, first add all shorter lengths. For a
fixed length, the first digit has `7` non-lucky choices because zero is not
allowed, and `2` lucky choices.

Then scan `s` from left to right. At each position, try digits smaller than the
current bound digit, add the corresponding `free[remaining][newDist]`, and then
continue only if the actual bound digit itself keeps the lucky digits sparse.
If the scan finishes, include `s` itself.

Let:

```text
nearly(x) = x - sparseLucky(x)
```

where `x` is taken modulo `1000000007`. The answer for `[l, r]` is:

```text
nearly(r) - nearly(l - 1)
```

The precomputation costs `O(maxLen * k)`, and each bound is processed in
`O(maxLen * 10)`. Here `maxLen <= 1000`, so this fits easily.
