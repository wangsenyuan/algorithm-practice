# B. Composite Coloring

[Problem link](https://codeforces.com/problemset/problem/1332/B)

**Contest:** [Codeforces Round 630 (Div. 2)](https://codeforces.com/contest/1332)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

## Problem

A composite number is a product of two integers that are both greater than `1`.

Alice has `n` composite numbers. Choose `m` with `1 ≤ m ≤ 11` and color each number with one color from `1` to `m` so that:

- every color from `1` to `m` is used at least once
- any two numbers of the same color have gcd greater than `1`

Equal numbers may receive different colors. Any valid coloring is accepted.

## Constraints

- `1 ≤ t ≤ 1000`
- `1 ≤ n ≤ 1000`
- `4 ≤ a_i ≤ 1000`
- every `a_i` is composite
- the sum of `n` over all test cases does not exceed `10^4`

## Input

The first line contains `t`, the number of test cases.

Each test case is two lines:

```text
n
a_1 a_2 ... a_n
```

## Output

For each test case print `m`, then `n` integers `c_1 ... c_n`, the color of each element.

## Sample Input 1

```text
3
6 10 15
```

## Sample Output 1

```text
1
1 1 1
```

`gcd(6, 10) = 2`, `gcd(6, 15) = 3`, and `gcd(10, 15) = 5`, so one color is valid. Other colorings are also accepted.

## Sample Input 2

```text
2
4 9
```

## Sample Output 2

```text
2
2 1
```

Each color is used once, so the gcd condition holds.

## Sample Input 3

```text
23
437 519 865 808 909 391 194 291 237 395 323 365 511 497 781 737 871 559 731 697 779 841 961
```

## Sample Output 3

```text
11
4 7 8 10 7 3 10 7 7 8 3 1 1 5 5 9 2 2 3 3 4 11 6
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
