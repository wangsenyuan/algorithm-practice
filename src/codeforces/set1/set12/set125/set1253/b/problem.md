# B. Silly Mistake

[Problem link](https://codeforces.com/problemset/problem/1253/B)

**Contest:** [Codeforces Round 600 (Div. 2)](https://codeforces.com/contest/1253)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

The log `a` has `n` events. A positive value `i` means employee `i` enters the office. A negative value `-i` means employee `i` leaves.

Split the log into one or more days. A day is valid when all of these hold:

- each employee enters at most once that day
- an employee can leave only after entering, and only once
- the office is empty at the start and at the end of the day

Print the lengths of the days. Any valid split is accepted. If no split exists, print `-1`.

## Constraints

- `1 ≤ n ≤ 10^5`
- `a_i ≠ 0`
- `1 ≤ |a_i| ≤ 10^6`

## Input

```text
n
a_1 a_2 ... a_n
```

## Output

If there is no valid split, print `-1`.

Otherwise print the number of days `d`, then `d` positive integers, the lengths of the days. Their sum must be `n`.

## Sample Input 1

```text
6
1 7 -7 3 -1 -3
```

## Sample Output 1

```text
1
6
```

The whole log is one valid day.

## Sample Input 2

```text
8
1 -1 1 2 -1 -2 3 -3
```

## Sample Output 2

```text
2
2 6
```

One valid split is a day of length `2` followed by a day of length `6`. Other valid splits are also accepted.

## Sample Input 3

```text
6
2 5 -5 5 -5 -2
```

## Sample Output 3

```text
-1
```

## Sample Input 4

```text
3
-8 1 1
```

## Sample Output 4

```text
-1
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
