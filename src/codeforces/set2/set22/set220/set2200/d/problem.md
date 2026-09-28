# D. Portal

[Problem link](https://codeforces.com/problemset/problem/2200/D)

**Contest:** [Codeforces Round 1084 (Div. 3)](https://codeforces.com/contest/2200)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given a permutation `p` of length `n`. There are also two portals located at positions
`x` and `y` (`x < y`).

A portal at position `i` is initially located between the `i`-th and `(i+1)`-th elements of the
array. Specifically, if `i = 0`, then the portal is located before the first element of the array,
and if `i = n`, then the portal is located after the last element.

You may perform either of the following two operations as many times as you like:

1. Remove the element to the immediate left of one portal and insert it to the immediate right of
   the other portal.
2. Remove the element to the immediate right of one portal and insert it to the immediate left of
   the other portal.

Let `O` denote a portal. For example, if `p` is `[3, O, 2, 4, O, 1]`:

- Using operation 1 on the left and right portals respectively results in the arrays
  `[O, 2, 4, O, 3, 1]` and `[3, O, 4, 2, O, 1]`.
- Using operation 2 on the left and right portals respectively results in the arrays
  `[3, O, 4, 2, O, 1]` and `[3, 1, O, 2, 4, O]`.

Find the lexicographically smallest permutation you can obtain using these operations. Note that
portals do not affect the lexicographical comparison of permutations.

A permutation of length `n` is an array of length `n` containing each integer from `1` to `n`
exactly once.

A permutation `a` is lexicographically smaller than permutation `b` if there exists an index `i`
such that `a_j = b_j` for all indices `1 ≤ j < i` and `a_i < b_i`.

## Input

The first line contains an integer `t` (`1 ≤ t ≤ 2 · 10^4`) — the number of test cases.

For each test case, the first line contains three integers `n`, `x`, and `y`
(`1 ≤ n ≤ 2 · 10^5`, `0 ≤ x < y ≤ n`).

The second line of each test case contains `n` integers `p_1, p_2, …, p_n` — a permutation of
length `n`.

The sum of `n` over all test cases does not exceed `2 · 10^5`.

## Output

For each test case, output a line with `n` integers — the lexicographically smallest permutation
you can obtain.

## Example

### Input

```text
4
4 0 4
3 1 4 2
3 1 2
3 2 1
5 1 3
1 3 5 2 4
2 0 1
1 2
```

### Output

```text
1 4 2 3
2 3 1
1 2 3 5 4
1 2
```

### Note

In the first test case, the array is `[O, 3, 1, 4, 2, O]`. Using operation 2 on the left portal
results in `[O, 1, 4, 2, 3, O]`, which is the lexicographically smallest possible permutation that
can be obtained.

In the second test case, the array is `[3, O, 2, O, 1]`. Using operation 1 on the left portal
results in `[O, 2, O, 3, 1]`, which is the lexicographically smallest possible permutation that can
be obtained.

In the fourth test case, it is optimal not to do any operations.
