# L. Berland University

[Problem link](https://codeforces.com/problemset/problem/1090/L)

**Contest:** [Codeforces contest 1090](https://codeforces.com/contest/1090)

time limit: 1 second

memory limit: 512 megabytes

## Problem

There are `t` students and a course of `n` lectures. A student passes if they
attend at least `k` lectures.

There are two auditoriums of capacities `a` and `b`. Odd-numbered lectures use
the first auditorium, even-numbered lectures use the second (lecture 1 → first,
lecture 2 → second, and so on).

Find the maximum number of students that can attend at least `k` lectures.

## Constraints

- `1 <= t, n, a, b, k <= 10^9`

## Input

```text
t n a b k
```

## Output

Print a single integer — the maximum number of students that can pass.

## Sample 1

```text
Input
10 3 4 4 3

Output
4
```

## Sample 2

```text
Input
10 3 4 4 5

Output
0
```

## Sample 3

```text
Input
100000 100000 100000 100000 1

Output
100000
```

## Sample 4

```text
Input
5 4 5 3 3

Output
5
```

## Sample 5

```text
Input
100 9 6 3 6

Output
7
```

## Solution

Only students who can pass matter; anyone else can skip every lecture. So the
answer is some `w` with `0 <= w <= t`: maximize `w` such that `w` students can
each attend at least `k` lectures.

If `k > n`, nobody can pass → answer `0`.

### Feasibility of a fixed `w`

There are `n1 = (n + 1) / 2` odd lectures (capacity `a` each) and
`n2 = n / 2` even lectures (capacity `b` each).

Among these `w` candidates, each odd lecture can seat at most `min(a, w)` of
them, and each even lecture at most `min(b, w)`. Total usable seats:

```text
seats = n1 * min(a, w) + n2 * min(b, w)
```

They need `w * k` attendances, so `w` is feasible iff

```text
seats >= w * k
```

(equivalently `seats / k >= w` with integer division, which avoids multiplying
`w * k` when values are large).

Odd/even lectures of the same type are interchangeable. If the total seat count
is enough and `k <= n`, the attendances can always be distributed so each of the
`w` students gets at least `k`.

### Binary search

`check(w)` is monotonic: if `w` works, every smaller value works. Search the
largest feasible `w` in `[0, t]` (upper bound is the student count `t`, not the
lecture count `n` — many students can share a short course when `k` is small).

```text
answer = max { w | 0 <= w <= t, seats(w) >= w * k }
```

### Complexity

- Time: `O(log t)` binary searches, each check `O(1)`
- Space: `O(1)`
