# C. Loyalty

[Problem link](https://codeforces.com/problemset/problem/2161/C)

**Contest:** [Pinely Round 5 (Div. 1 + Div. 2)](https://codeforces.com/contest/2161)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

## Problem

You want to buy `n` items. Item `i` costs `a_i`, and `1 ≤ a_i ≤ X`.

Your loyalty level is `floor(S / X)`, where `S` is the total cost of items bought so far.
`S` starts at `0`.

If buying an item of price `p` makes the loyalty level increase, you earn `p` bonus points.
Choose an order of purchase that maximizes the bonus points.

## Constraints

- `1 ≤ t ≤ 2 · 10^4`
- `1 ≤ n ≤ 10^5`
- `1 ≤ X ≤ 10^9`
- `1 ≤ a_i ≤ X`
- The sum of `n` over all test cases does not exceed `10^5`

## Input

The first line contains `t` — the number of test cases.

Each test case:

```text
n X
a1 a2 ... an
```

## Output

For each test case, print two lines: the maximum bonus points, then the prices in a
purchase order that achieves it. Any optimal order is accepted.

## Example

### Input

```text
7
10 2
1 2 1 2 1 2 1 2 1 2
5 10
2 2 2 2 5
11 23
5 5 22 1 21 2 10 3 1 1 2
1 1
1
1 17
11
3 100
44 32 1
16 100500
42801 73112 95296 68791 42217 21871 29316 84405 24273 42894 63370 53473 57156 61369 80 27290
```

### Output

```text
12
1 2 2 2 2 2 1 1 1 1
5
2 2 2 2 5
53
1 1 5 2 1 2 5 3 10 21 22
1
1
0
11
0
44 32 1
503499
53473 42894 80 57156 42801 61369 42217 63370 29316 68791 27290 73112 24273 84405 21871 95296
```

## Note

In the first test case, `X = 2`. Buying `1 2 2 2 2 2 1 1 1 1` crosses a multiple of `2`
on the purchases of prices `2, 2, 2, 2, 2, 1, 1`, for `12` points.

In the second test case, the first four items leave `S = 8`. The last item raises it to
`13` and earns `5` points.

## Solution

*(TODO)*
