# B. Alice, Bob, Two Teams

[Problem link](https://codeforces.com/problemset/problem/632/B)

**Contest:** [Educational Codeforces Round 9](https://codeforces.com/contest/632)

time limit per test: 1.5 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

There are `n` pieces. Piece `i` has strength `p_i`. Alice assigns each piece to team `A` or team `B`.

Bob may flip one prefix or one suffix of that assignment, changing `A` to `B` and `B` to `A`. He may also flip nothing. Alice keeps the pieces marked `A`. Bob keeps the pieces marked `B`.

Print the maximum total strength Bob can obtain.

## Constraints

- `1 ≤ n ≤ 5·10^5`
- `1 ≤ p_i ≤ 10^9`
- The assignment is a string of length `n` over `A` and `B`

## Input

```text
n
p_1 p_2 ... p_n
s
```

## Output

Print one integer, Bob's maximum strength.

## Sample Input 1

```text
5
1 2 3 4 5
ABABA
```

## Sample Output 1

```text
11
```

Bob flips the suffix of length 1. The assignment becomes `ABABB`, and his pieces are `2`, `4`, and `5`.

## Sample Input 2

```text
5
1 2 3 4 5
AAAAA
```

## Sample Output 2

```text
15
```

Bob flips the whole string.

## Sample Input 3

```text
1
1
B
```

## Sample Output 3

```text
1
```

Bob flips nothing.

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
