# B. Physics Practical

[Problem link](https://codeforces.com/problemset/problem/253/B)

**Contest:** [Codeforces Round 154 (Div. 2)](https://codeforces.com/contest/253)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: input.txt

output: output.txt

## Problem

Vasya recorded `n` capacitance measurements. The teacher rejects the work when the largest remaining result is more than twice the smallest one.

Erase as few results as possible so that the remaining values `x` (minimum) and `y` (maximum) satisfy `y ≤ 2x`. Print that minimum number of erasures.

## Constraints

- `2 ≤ n ≤ 10^5`
- `1 ≤ c_i ≤ 5000`

## Input

```text
n
c_1 c_2 ... c_n
```

## Output

Print one integer, the minimum number of results to erase.

## Sample Input 1

```text
6
4 5 3 8 3 7
```

## Sample Output 1

```text
2
```

Erase `8` and `7`. The remaining values are `4 5 3 3`, so the maximum is `5` and the minimum is `3`. Erasing both `3`s also works: the remaining maximum is `8` and the minimum is `4`.

## Sample Input 2

```text
4
4 3 2 4
```

## Sample Output 2

```text
0
```

The maximum is `4` and the minimum is `2`, and `4 ≤ 2·2`.

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
