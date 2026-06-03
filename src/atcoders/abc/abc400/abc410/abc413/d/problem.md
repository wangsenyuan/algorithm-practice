# D - Make Geometric Sequence (ABC413)

**Contest:** [ABC413](https://atcoder.jp/contests/abc413) — AtCoder Beginner Contest 413  
**Task:** [https://atcoder.jp/contests/abc413/tasks/abc413_d](https://atcoder.jp/contests/abc413/tasks/abc413_d)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 425 points

## Problem Statement

You are given an integer sequence `A = (A_1, A_2, ..., A_N)` of length `N`. It
is guaranteed that `A_i != 0` for all `1 <= i <= N`.

Determine whether there exists a permutation `B = (B_1, B_2, ..., B_N)` of `A`
such that `B` forms a geometric sequence.

A sequence `S = (S_1, S_2, ..., S_N)` is a geometric sequence if there exists a
real number `r` such that `S_{i+1} = r S_i` for all integers `1 <= i < N`.

Solve `T` test cases.

## Constraints

- `1 <= T <= 10^5`
- `2 <= N <= 2 * 10^5`
- `-10^9 <= A_i <= 10^9`
- `A_i != 0`
- The sum of `N` over all test cases is at most `2 * 10^5`.
- All input values are integers.

## Input

The input is given from standard input in the following format:

```text
T
testcase_1
testcase_2
...
testcase_T
```

Each test case is given in the following format:

```text
N
A_1 A_2 ... A_N
```

## Output

Output `T` lines. For each test case, print `Yes` if `A` can be rearranged to
form a geometric sequence, and `No` otherwise.

## Sample Input 1

```text
3
5
1 8 2 4 16
5
-16 24 54 81 -36
7
90000 8100 -27000 729 -300000 -2430 1000000
```

## Sample Output 1

```text
Yes
No
Yes
```

In the first test case, the rearrangement `(16, 8, 4, 2, 1)` forms a geometric
sequence with common ratio `1/2`.

In the second test case, no rearrangement satisfies the condition.
