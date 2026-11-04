# E. Little Elephant and Tree

[Problem link](https://codeforces.com/problemset/problem/258/E)

**Contest:** [Codeforces Round 157 (Div. 1)](https://codeforces.com/contest/258)

time limit per test: 4 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

A rooted tree has `n` nodes. Node `1` is the root. Every node starts with an empty list.

There are `m` operations. Operation `i` appends the number `i` to every node in the subtree of `a_i`, then appends `i` to every node in the subtree of `b_i`.

For each node `i`, let `c_i` be the number of nodes `j ≠ i` whose list shares at least one number with the list of `i`. Print `c_1, c_2, ..., c_n`.

## Constraints

- `1 ≤ n, m ≤ 10^5`
- The `n - 1` edges form an undirected tree
- `1 ≤ a_i, b_i ≤ n` and `a_i ≠ b_i`

## Input

```text
n m
u_1 v_1
u_2 v_2
...
u_{n-1} v_{n-1}
a_1 b_1
a_2 b_2
...
a_m b_m
```

## Output

Print `n` integers, `c_1` through `c_n`.

## Sample Input 1

```text
5 1
1 2
1 3
3 5
3 4
2 3
```

## Sample Output 1

```text
0 3 3 3 3
```

## Sample Input 2

```text
11 3
1 2
2 3
2 4
1 5
5 6
5 7
5 8
6 9
8 10
8 11
2 9
3 6
2 8
```

## Sample Output 2

```text
0 6 7 6 0 2 0 5 4 5 5
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
