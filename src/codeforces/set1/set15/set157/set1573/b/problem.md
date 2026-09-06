# B. Swaps

[Problem link](https://codeforces.com/problemset/problem/1573/B)

**Contest:** [Codeforces Round 743 (Div. 2)](https://codeforces.com/contest/1573)

time limit per test: 1 second

memory limit per test: 256 megabytes

input: standard input

output: standard output

You are given two arrays `a` and `b` of length `n`. Array `a` contains each odd integer from `1` to `2n` in an arbitrary order, and array `b` contains each even integer from `1` to `2n` in an arbitrary order.

You can perform the following operation on those arrays:

- choose one of the two arrays
- pick an index `i` from `1` to `n-1`
- swap the `i`-th and the `(i+1)`-th elements of the chosen array

Compute the minimum number of operations needed to make array `a` lexicographically smaller than array `b`.

For two different arrays `x` and `y` of the same length `n`, we say that `x` is lexicographically smaller than `y` if in the first position where `x` and `y` differ, the array `x` has a smaller element than the corresponding element in `y`.

An answer always exists.

## Input

The first line contains the number of test cases `t` (`1 <= t <= 10^4`).

The first line of each test case contains a single integer `n` (`1 <= n <= 10^5`) — the length of the arrays.

The second line contains `n` integers `a_1, a_2, …, a_n` (`1 <= a_i <= 2n`, all `a_i` are odd and pairwise distinct).

The third line contains `n` integers `b_1, b_2, …, b_n` (`1 <= b_i <= 2n`, all `b_i` are even and pairwise distinct).

The sum of `n` over all test cases does not exceed `10^5`.

## Output

For each test case, print one integer: the minimum number of operations needed to make array `a` lexicographically smaller than array `b`.

## Example

### Input

```text
3
2
3 1
4 2
3
5 3 1
2 4 6
5
7 5 9 1 3
2 4 6 10 8
```

### Output

```text
0
2
3
```

### Note

In the first example, `a` is already lexicographically smaller than `b`, so no operations are required.

In the second example, one way is to swap `5` and `3` and then swap `2` and `4`, which results in `[3, 5, 1]` and `[4, 2, 6]`.

## Solution

Every `a_i` is odd and every `b_j` is even, so the arrays never tie at
index `0`. Lexicographic order is therefore just `a[0] < b[0]`. Adjacent
swaps bubble any chosen pair to the front: bringing `a[j]` to position
`0` costs `j` swaps and bringing `b[i]` to position `0` costs `i`. The
rest of each array does not matter.

The answer is `min(i + j)` over all pairs with `a[j] < b[i]`. A segment
tree stores each odd value's index in `a`. For even `v = b[i]`, a
range-min query on values `[1, v)` returns the leftmost (cheapest) odd
strictly less than `v`, and the scan keeps the best `i + j`.

### Correctness sketch

After any sequence of adjacent swaps the first position still compares an
odd to an even, so the only way `a` becomes lex-smaller is `a[0] < b[0]`.
Any such first pair `(a[j], b[i])` can be moved to the front independently
in exactly `i + j` swaps, and no cheaper way exists because each element
must travel its own distance. The range query enumerates, for every even,
the cheapest compatible odd, so the recorded minimum is global.

### Complexity

Build and `n` updates plus `n` range-min queries on a tree of size
`O(n)`: `O(n log n)` time and `O(n)` memory. The sum of `n` over tests
is at most `10^5`.
