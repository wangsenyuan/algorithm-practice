# A. Accounting

[Problem link](https://codeforces.com/problemset/problem/30/A)

**Contest:** [Codeforces Beta Round #30](https://codeforces.com/contest/30)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

A long time ago in some far country lived king Copa. After the recent king's reform, he got so
large powers that he started to keep the books by himself.

The total income `A` of his kingdom during year 0 is known, as well as the total income `B` during
year `n` (these numbers can be negative — that means there was a loss in the corresponding year).

The king wants to show financial stability. To do this, he needs a common coefficient `X` — the
coefficient of income growth during one year. This coefficient should satisfy the equation:

```text
A · X^n = B
```

Fractional numbers are not used in the kingdom's economy. All input numbers as well as the
coefficient `X` must be integers. `X` may be zero or negative.

## Input

The input contains three integers `A`, `B`, `n` (`|A|, |B| ≤ 1000`, `1 ≤ n ≤ 10`).

## Output

Output the required integer coefficient `X`, or `No solution` if such a coefficient does not exist
or it is fractional. If there are several possible solutions, output any of them.

## Examples

### Input

```text
2 18 2
```

### Output

```text
3
```

### Input

```text
-1 8 3
```

### Output

```text
-2
```

### Input

```text
0 0 10
```

### Output

```text
5
```

### Input

```text
1 16 5
```

### Output

```text
No solution
```
