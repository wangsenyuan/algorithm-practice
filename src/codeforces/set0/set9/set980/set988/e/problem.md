# E. Divisibility by 25

[Problem link](https://codeforces.com/problemset/problem/988/E)

**Contest:** [Codeforces Round #488 (Div. 2)](https://codeforces.com/contest/988)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given an integer `n` from `1` to `10^18` without leading zeroes.

In one move you can swap any two **adjacent** digits, but the resulting number must not have leading
zeroes.

Find the minimum number of moves needed to obtain a number divisible by `25`. Print `-1` if it is
impossible.

## Input

One line contains `n` (`1 <= n <= 10^18`). The first digit of `n` is not zero.

## Output

If it is impossible, print `-1`. Otherwise print the minimum number of adjacent swaps.

## Example

### Input

```text
5071
```

### Output

```text
4
```

### Input

```text
705
```

### Output

```text
1
```

### Input

```text
1241367
```

### Output

```text
-1
```

### Note

One possible sequence for the first example is `5071 → 5701 → 7501 → 7510 → 7150`.

## Solution

A decimal number is divisible by `25` iff its last two digits are one of `00`, `25`, `50`, or `75`.
So the task is to try making each of these four suffixes and take the minimum number of adjacent
swaps.

For a target suffix `ab`, simulate the cheapest way to place it:

1. Scan from right to left to find digit `b`, then bubble it to the last position. If no such digit
   exists, this suffix is impossible.
2. In the updated string, scan from right to left before the last position to find digit `a`, then
   bubble it to the second-last position. If no such digit exists, this suffix is impossible.
3. If the first digit is not `0`, the current swap count is valid.
4. Otherwise the moves created a leading zero. Since the last two digits are already fixed, find the
   first non-zero digit among the remaining prefix and bubble it to the front. If no such digit
   exists, this suffix is impossible.

Adjacent swaps are exactly what bubbling counts: moving a digit from index `i` to the last position
costs the number of places it crosses, and after moving the rightmost required digit first, the
relative positions for the second digit match the simulated string.

There are only four possible suffixes. Each simulation scans and bubbles through at most all digits,
so the total complexity is `O(n)` with `O(n)` extra space for the temporary digit buffer.
