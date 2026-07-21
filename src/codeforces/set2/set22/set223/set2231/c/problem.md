# C. Chipmunk Theo and Equality

[Problem link](https://codeforces.com/problemset/problem/2231/C)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

Theo has a sequence of positive integers. In one operation, he chooses one
element and:

- if it is even, divides it by `2`;
- if it is odd, increases it by `1`.

Find the minimum number of operations to make all numbers in the sequence
equal.

## Constraints

- `1 ≤ t ≤ 10^4`
- `1 ≤ n ≤ 10^5`
- `1 ≤ a_i ≤ 10^9`
- Sum of `n` over all test cases ≤ `10^5`

## Input

```
t
n
a_1 a_2 ... a_n
...
```

## Output

For each test case, print one integer — the minimum number of operations.

## Samples

### Sample 1

Input:

```
1
3
3 2 4
```

Output:

```
3
```

One way: `[3,2,4] → [4,2,4] → [2,2,4] → [2,2,2]`.

### Sample 2

Input:

```
1
7
3 6 7 16 8 8 7
```

Output:

```
11
```

### Sample 3

Input:

```
1
3
1 4 2
```

Output:

```
2
```

### Sample 4

Input:

```
1
5
10 10 10 10 10
```

Output:

```
0
```

### Sample 5

Input:

```
1
6
1 1 3 1 1 1
```

Output:

```
3
```
