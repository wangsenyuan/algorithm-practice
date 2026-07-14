# C - Inverse of Permutation

[Problem link](https://atcoder.jp/contests/abc217/tasks/abc217_c)

Time Limit: 2 sec / Memory Limit: 1024 MiB

Score: 300 points

## Problem Statement

We will call a sequence of length `N` where each of `1, 2, ..., N` occurs once
a permutation of length `N`.

Given a permutation of length `N`, `P = (p_1, p_2, ..., p_N)`, print a
permutation of length `N`, `Q = (q_1, ..., q_N)`, that satisfies:

- For every `i` (`1 <= i <= N`), the `p_i`-th element of `Q` is `i`.

It can be proved that there exists a unique `Q` that satisfies the condition.

## Constraints

- `1 <= N <= 2 * 10^5`
- `(p_1, p_2, ..., p_N)` is a permutation of length `N`
- All values in input are integers

## Input

```
N
p_1 p_2 ... p_N
```

## Output

Print the sequence `Q` in one line, with spaces in between.

```
q_1 q_2 ... q_N
```

## Samples

### Sample 1

Input:

```
3
2 3 1
```

Output:

```
3 1 2
```

### Sample 2

Input:

```
3
1 2 3
```

Output:

```
1 2 3
```

### Sample 3

Input:

```
5
5 3 2 4 1
```

Output:

```
5 3 2 4 1
```
