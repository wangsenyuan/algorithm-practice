# D. Caveman

[Problem link](https://codeforces.com/problemset/problem/2276/D)

**Contest:** [Educational Codeforces Round 195 (Rated for Div. 2)](https://codeforces.com/contest/2276)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

## Problem

Monocarp drew an undirected graph with `n` vertices and `m` edges. He wants to color each vertex red or blue so that every vertex has at least one neighbor of the opposite color.

Count the valid colorings. Two colorings are different when at least one vertex has a different color.

Monocarp only knows the numbers `1`, `2`, and "many" (any number greater than `2`). He does not know zero, so no valid coloring makes his brain melt.

## Constraints

- `1 ≤ t ≤ 10^4`
- `3 ≤ n ≤ 3 · 10^5`
- `0 ≤ m ≤ min(n(n-1)/2, 3 · 10^5)`
- `1 ≤ u_i, v_i ≤ n`
- The graph has no self-loops and no multiple edges
- The sum of `n` over all test cases does not exceed `3 · 10^5`
- The sum of `m` over all test cases does not exceed `3 · 10^5`

## Input

The first line contains `t`, the number of test cases.

Each test case starts with:

```text
n m
```

Then `m` lines follow, each with the endpoints of one edge:

```text
u_i v_i
```

## Output

For each test case, print one line:

- `1` or `2`, when that is the number of valid colorings
- `Many`, when there are more than two
- `Brain Melt`, when there are none

## Sample Input 1

```text
3 1
1 2
```

## Sample Output 1

```text
Brain Melt
```

Vertex `3` is isolated, so it has no neighbor of the opposite color.

## Sample Input 2

```text
3 2
1 2
1 3
```

## Sample Output 2

```text
2
```

Vertices `2` and `3` must both differ from `1`. The two colorings are `1` red with the others blue, and `1` blue with the others red.

## Sample Input 3

```text
3 3
1 2
2 3
1 3
```

## Sample Output 3

```text
Many
```

## Sample Input 4

```text
4 3
1 2
2 3
3 4
```

## Sample Output 4

```text
Many
```

## Status

I/O and the official samples are in place. `solve` is left as a TODO.
