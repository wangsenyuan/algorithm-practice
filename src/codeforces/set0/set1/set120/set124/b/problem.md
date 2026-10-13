# B. Permutations

[Problem link](https://codeforces.com/problemset/problem/124/B)

**Contest:** [Codeforces contest 124](https://codeforces.com/contest/124)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given `n` integers, each with exactly `k` digits. Rearrange the digits of every integer by the same rule, then minimize the difference between the largest and the smallest resulting number. Leading zeros are allowed both before and after the rearrangement.

## Input

The first line contains two integers `n` and `k` (`1 <= n, k <= 8`).

Each of the next `n` lines contains a `k`-digit integer. Leading zeros are allowed.

## Output

Print one integer: the minimum possible difference between the largest and the smallest number after one common rearrangement of the digits.

## Example

### Input 1

```text
6 4
5237
2753
7523
5723
5327
2537
```

### Output 1

```text
2700
```

### Input 2

```text
3 3
010
909
012
```

### Output 2

```text
3
```

### Input 3

```text
7 5
50808
36603
37198
44911
29994
42543
50156
```

### Output 3

```text
20522
```

In the first sample, the digit order `(3, 1, 4, 2)` turns the second and fourth numbers into `5237` and `2537`.

In the second sample, swapping the first and second digits produces `100`, `99`, and `102`.

## Status

I/O and official sample cases are in place. `solve` is left as a TODO.
