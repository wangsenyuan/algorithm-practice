# C. Network Mask

[Problem link](https://codeforces.com/problemset/problem/291/C)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Polycarpus has `n` distinct IPv4 addresses. Each address is four 8-bit numbers
(without leading zeroes) separated by dots.

A **subnet mask** is an IP address whose 32-bit form is `11...1100...00`: one or
more leading ones followed by one or more trailing zeros.

The network address of an IP is the bitwise AND of the IP and the subnet mask.

Find a subnet mask such that the given IPs fall into exactly `k` distinct
networks. If several masks work, choose the one with the fewest ones. If none
exists, print `-1`.

## Constraints

- `1 <= k <= n <= 10^5`
- All IP addresses are distinct and valid in the simplified model of the
  statement

## Input

The first line contains two integers `n` and `k`.

Each of the next `n` lines contains one IP address.

## Output

Print the subnet mask in dotted form, or `-1` if it does not exist.

## Examples

### Sample 1

```text
Input
5 3
0.0.0.1
0.1.1.2
0.0.2.1
0.1.1.0
0.0.2.3

Output
255.255.254.0
```

### Sample 2

```text
Input
5 2
0.0.0.1
0.1.1.2
0.0.2.1
0.1.1.0
0.0.2.3

Output
255.255.0.0
```

### Sample 3

```text
Input
2 1
255.0.0.1
0.0.0.2

Output
-1
```
