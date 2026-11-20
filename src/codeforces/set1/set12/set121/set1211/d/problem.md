# D. Teams

[Problem link](https://codeforces.com/problemset/problem/1211/D)

**Contest:** [Kotlin Heroes: Episode 2](https://codeforces.com/contest/1211)

time limit per test: 3 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

There are `n` players. Player `i` has rating `r_i`. Form as many teams as possible. Each player is in at most one team, and each team has exactly `a + b` players:

- `a` players share one rating `x`
- `b` players share another rating `y`
- `y = k · x`

## Constraints

- `1 ≤ n, a, b ≤ 3 · 10^5`
- `2 ≤ k ≤ 1000`
- `1 ≤ r_i ≤ 10^6`

## Input

```text
n a b k
r_1 r_2 ... r_n
```

## Output

Print one integer, the maximum number of teams.

## Sample Input 1

```text
12 1 2 2
1 1 2 2 2 2 2 3 3 4 6 6
```

## Sample Output 1

```text
3
```

Two teams use ratings `1` and `2`: each is one `1` and two `2`s. One team uses ratings `3` and `6`: one `3` and two `6`s.

## Sample Input 2

```text
14 1 1 3
3 3 1 1 9 9 2 3 6 6 3 18 3 18
```

## Sample Output 2

```text
6
```

## Sample Input 3

```text
1 2 3 10
1000000
```

## Sample Output 3

```text
0
```

## Status
1. x * k, x * k * k, x * k * k * k 
2. 固定x的情况下, 这个链条最长是10 (1024)
3. 
