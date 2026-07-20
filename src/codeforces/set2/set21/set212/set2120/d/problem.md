# D. Matrix game

[Problem link](https://codeforces.com/problemset/problem/2120/D)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Aryan and Harshith play a game with integers `a`, `b`, and `k`.

1. Aryan chooses two integers `n` and `m`.
2. Harshith chooses an `n × m` matrix with entries in `{1, ..., k}`.
3. Aryan wins if the matrix contains an `a × b` submatrix (contiguous block of
   `a` rows and `b` columns) whose entries are all equal.

Given `a`, `b`, and `k`, find the lexicographically minimum pair `(n, m)` such
that Aryan **always** wins no matter how Harshith fills the matrix (Harshith
plays optimally). Output `n` and `m` modulo `10^9 + 7`.

`(n1, m1)` is lexicographically smaller than `(n2, m2)` if `n1 < n2`, or
`n1 = n2` and `m1 < m2`.

## Constraints

- Multi-test
- Sample values: `a, b, k` can be as large as `9 · 10^4` (see samples)
- Output `n` and `m` modulo `10^9 + 7`

## Input

```
t
a b k
...
```

## Output

For each test case, print two integers — `n` and `m` modulo `10^9 + 7`.

## Samples

### Sample 1

Input:

```
1
1 1 5
```

Output:

```
1 1
```

### Sample 2

Input:

```
1
2 2 2
```

Output:

```
3 7
```

### Sample 3

Input:

```
1
90000 80000 70000
```

Output:

```
299929959 603196135
```

### ideas
1. 没有想法~
2. 