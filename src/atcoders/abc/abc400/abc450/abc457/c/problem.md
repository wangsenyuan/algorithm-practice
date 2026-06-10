# C - Long Sequence

[Problem link](https://atcoder.jp/contests/abc457/tasks/abc457_c)

time limit: 2 seconds

memory limit: 1024 MiB

## Problem Statement

You are given integers `N` and `K`, together with `N` integer sequences
`A_1, A_2, ..., A_N` and an integer sequence
`C = (C_1, C_2, ..., C_N)`.

The length of `A_i` is `L_i`:

```text
A_i = (A_{i,1}, A_{i,2}, ..., A_{i,L_i})
```

Construct a sequence `B` as follows:

1. Initially, `B` is empty.
2. For `i = 1, 2, ..., N` in this order, append the entire sequence `A_i` to
   the end of `B` exactly `C_i` times.

Find the value of the `K`-th element of `B`.

It is guaranteed that:

```text
1 <= K <= sum(C_i * L_i)
```

## Constraints

- `1 <= N`
- `1 <= L_i`
- `sum(L_i) <= 2 * 10^5`
- `1 <= A_{i,j} <= 10^9`
- `1 <= C_i <= 10^9`
- `1 <= K <= sum(C_i * L_i)`
- All input values are integers.

## Input

The input is given in the following format:

```text
N K
L_1 A_{1,1} A_{1,2} ... A_{1,L_1}
L_2 A_{2,1} A_{2,2} ... A_{2,L_2}
...
L_N A_{N,1} A_{N,2} ... A_{N,L_N}
C_1 C_2 ... C_N
```

## Output

Print the value of `B_K`.

## Sample 1

### Input

```text
3 9
3 1 3 2
1 3
2 4 3
1 3 2
```

### Output

```text
4
```

The resulting sequence is:

```text
B = (1, 3, 2, 3, 3, 3, 4, 3, 4, 3)
```

Therefore, `B_9 = 4`.

## Sample 2

### Input

```text
3 1
1 7
1 111
1 5
1 100 10000
```

### Output

```text
7
```

## Sample 3

### Input

```text
3 3163812
5 1 2 3 4 5
4 9 8 7 6
2 10 11
87043 908415 9814
```

### Output

```text
9
```
