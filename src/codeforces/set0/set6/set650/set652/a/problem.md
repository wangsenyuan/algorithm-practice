# A. Gabriel and Caterpillar

[Problem link](https://codeforces.com/problemset/problem/652/A)

**Contest:** [Educational Codeforces Round 10](https://codeforces.com/contest/652)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

A caterpillar is at height `h1` cm. An apple hangs at height `h2` cm on the same tree, with `h2 > h1`.

The day runs from 10 am to 10 pm. During the day the caterpillar climbs `a` cm per hour. During the night it slips `b` cm per hour. Gabriel notices it at 2 pm, just after class, so the first day has 8 hours of climbing. Each later day has 12 hours of climbing, and each night has 12 hours of slipping.

The forest is magic: the caterpillar may slip below the ground and still climb back up to the apple.

Print the number of days Gabriel must wait before he can see the caterpillar reach the apple. The day he notices it is day 0. If the caterpillar never reaches the apple, print `-1`.

## Constraints

- `1 ≤ h1 < h2 ≤ 10^5`
- `1 ≤ a, b ≤ 10^5`

## Input

```text
h1 h2
a b
```

## Output

Print one integer, the number of days, or `-1`.

## Sample Input 1

```text
10 30
2 1
```

## Sample Output 1

```text
1
```

At 10 pm on day 0 the caterpillar is at height 26. At 10 am the next day it has slipped to 14. At 6 pm that day it reaches the apple.

## Sample Input 2

```text
10 13
1 1
```

## Sample Output 2

```text
0
```

## Sample Input 3

```text
10 19
1 2
```

## Sample Output 3

```text
-1
```

## Sample Input 4

```text
1 50
5 4
```

## Sample Output 4

```text
1
```

The caterpillar slips below the ground, then reaches the apple on the next day.

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
