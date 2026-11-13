# P3069 [USACO13JAN] Cow Lineup G

[Problem link](https://www.luogu.com.cn/problem/P3069)

time limit per test: 1.00s

memory limit per test: 125.00MB

## Problem

Farmer John has `N` cows in a line. Cow `i` has breed ID `B(i)`, an integer from `0` to `10^9`. Several cows may share a breed ID.

FJ may choose at most `K` breed IDs and remove every cow of those breeds. After the removals, the remaining cows stay in their original order.

Print the longest contiguous block of remaining cows that all have the same breed ID.

## Input

```text
N K
B(1)
B(2)
...
B(N)
```

## Output

Print one integer, the maximum length.

## Constraints

- `1 ≤ N ≤ 10^5`
- `0 ≤ B(i) ≤ 10^9`

## Sample Input

```text
9 1
2
7
3
7
7
3
7
5
7
```

## Sample Output

```text
4
```

The breeds are `2 7 3 7 7 3 7 5 7`. Removing every cow of breed `3` leaves `2 7 7 7 7 5 7`. The longest run of one breed is the four `7`s.

## Status

I/O and the official sample are in place. `solve` is left as a TODO.
