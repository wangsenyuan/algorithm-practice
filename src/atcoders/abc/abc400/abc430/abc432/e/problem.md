# E - Clamp

[Problem link](https://atcoder.jp/contests/abc432/tasks/abc432_e)

**Contest:** [OMRON Corporation Programming Contest 2025 #2 (AtCoder Beginner Contest 432)](https://atcoder.jp/contests/abc432)

time limit: 2 sec

memory limit: 1024 MiB

score: 450 points

You are given an integer sequence `A = (A_1, A_2, ..., A_N)` of length `N`.

Process `Q` queries in order. Each query is one of:

- Type 1: `1 x y` — set `A_x` to `y`
- Type 2: `2 l r` — compute `sum_{1 <= i <= N} max(l, min(r, A_i))`

## Constraints

- `1 <= N <= 5 * 10^5`
- `1 <= Q <= 2 * 10^5`
- `0 <= A_i <= 5 * 10^5`
- For type 1 queries: `1 <= x <= N`, `0 <= y <= 5 * 10^5`
- For type 2 queries: `0 <= l, r <= 5 * 10^5`
- All input values are integers

## Input

```text
N Q
A_1 A_2 ... A_N
query_1
query_2
...
query_Q
```

Each query is either `1 x y` or `2 l r`.

## Output

Let `K` be the number of type 2 queries. Print `K` lines. The i-th line is the
answer to the i-th type 2 query.

## Sample Input 1

```text
3 4
4 3 2
1 1 7
2 3 5
1 2 0
2 4 2
```

## Sample Output 1

```text
11
12
```

Initially `A = (4, 3, 2)`.

- Query 1: set `A_1` to `7`, so `A = (7, 3, 2)`
- Query 2: `max(3, min(5, 7)) + max(3, min(5, 3)) + max(3, min(5, 2)) = 5 + 3 + 3 = 11`
- Query 3: set `A_2` to `0`, so `A = (7, 0, 2)`
- Query 4: `max(4, min(2, 7)) + max(4, min(2, 0)) + max(4, min(2, 2)) = 4 + 4 + 4 = 12`

## Sample Input 2

```text
8 10
320 578 244 604 145 839 156 857
2 400 556
1 5 168
2 254 62
2 145 301
1 1 23
1 3 0
2 413 758
2 297 613
1 8 451
2 598 692
```

## Sample Output 2

```text
3824
2032
2073
4350
3596
4884
```
