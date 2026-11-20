# C. Ice Cream

[Problem link](https://codeforces.com/problemset/problem/1211/C)

**Contest:** [Kotlin Heroes: Episode 2](https://codeforces.com/contest/1211)

time limit per test: 3 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Summer lasts `n` days. On day `i`, one portion of ice cream costs `c_i`. Tanya wants to eat exactly `k` portions in total. On day `i` she eats `d_i` portions, and `a_i ≤ d_i ≤ b_i`.

A portion must be eaten on the day it is bought. Find the minimum amount of money she can spend. If no schedule eats exactly `k` portions, print `-1`.

## Constraints

- `1 ≤ n ≤ 2 · 10^5`
- `0 ≤ k ≤ 10^9`
- `0 ≤ a_i ≤ b_i ≤ 10^9`
- `1 ≤ c_i ≤ 10^6`

## Input

```text
n k
a_1 b_1 c_1
a_2 b_2 c_2
...
a_n b_n c_n
```

## Output

Print one integer, the minimum cost, or `-1`.

## Sample Input 1

```text
3 7
3 5 6
0 3 4
3 3 3
```

## Sample Output 1

```text
31
```

Eat `3`, `1`, and `3` portions. The cost is `3·6 + 1·4 + 3·3 = 31`.

## Sample Input 2

```text
1 45000
40000 50000 100000
```

## Sample Output 2

```text
4500000000
```

## Sample Input 3

```text
3 100
2 10 50
50 60 16
20 21 25
```

## Sample Output 3

```text
-1
```

## Sample Input 4

```text
4 12
2 5 1
1 2 2
2 3 7
3 10 4
```

## Sample Output 4

```text
35
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
