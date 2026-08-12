# B. Zhily and Barknights

[Problem link](https://codeforces.com/problemset/problem/2223/B)

**Contest:** [Codeforces Round 1097 (Div. 1, Based on Zhili Cup 2026)](https://codeforces.com/contest/2223)

time limit per test: 4 seconds

memory limit per test: 512 megabytes

input: standard input

output: standard output

## Problem

Zhily developed a game called Barknights, and she is preparing to release a
major update on March 25. Specifically, she plans to add a Module to each
Operator in the game, multiplying their power level by the Module's power.

After the update is released, the famous game streamer Jily will evaluate the
Operators' power levels and rank them in a tier list. Whenever an
earlier-released Operator is ranked higher than a later-released Operator, it
will cause a wave of drama.

Unfortunately, Zhily accidentally knocked over a hot water kettle and broke her
computer, which caused all the Modules to be randomly rearranged. Now Zhily
wants to know the expected number of waves of drama that will be generated, but
since she has to move on to the next problem, she has entrusted this task to
you.

You are given two arrays `a` and `b` of `n` positive integers. Let `b'` be a
permutation of array `b` chosen uniformly at random among all `n!` possible
permutations. Define `c_i = a_i · b'_i` for `1 ≤ i ≤ n`.

Find the expected number of inversions\* of array `c`.

\*An inversion in array `c` is a pair of indices `(i, j)` such that
`1 ≤ i < j ≤ n` and `c_i > c_j`.

## Constraints

- `1 ≤ t ≤ 100`
- `1 ≤ n ≤ 2000`
- `1 ≤ a_i, b_i ≤ 10^9`
- Sum of `n` over all test cases does not exceed `2000`

## Input

Each test contains multiple test cases. The first line contains the number of
test cases `t`.

The first line of each test case contains a single integer `n` — the length of
the arrays `a` and `b`.

The second line contains `n` integers `a_1, a_2, …, a_n`.

The third line contains `n` integers `b_1, b_2, …, b_n`.

## Output

For each test case, output the expected number of inversions in `c` modulo
`998244353`.

Formally, let `M = 998244353`. It can be shown that the exact answer can be
expressed as an irreducible fraction `p/q`, where `p` and `q` are integers and
`q ≢ 0 (mod M)`. Output the integer equal to `p · q^{-1} mod M`.

## Example

### Input

```text
3
5
1 14 5 1 4
1 1 1 1 1
3
3 2 5
3 2 5
10
10 72 65 43 73 23 78 13 49 99
31 90 45 19 44 18 59 31 48 29
```

### Output

```text
5
665496236
820778710
```

## Note

In the first test case, since all elements of `b` are `1`, any of the `5!`
permutations results in `b' = (1,1,1,1,1)`. Thus `c = (1,14,5,1,4)` is always
constant. The inversions are `(2,3), (2,4), (2,5), (3,4),` and `(3,5)`. The
expected number of inversions is `5`.

In the second test case, there are `3! = 6` equally likely permutations `b'`.
The expected number of inversions is `(1+0+0+0+2+1)/6 = 2/3`.

Modulo `998244353`, the answer is `2 · 3^{-1} ≡ 665496236`.

## Solution

For every fixed pair of indices `i < j`, only the two modules assigned to
these two positions matter.

Suppose the module at `i` is `x` and the module at `j` is `y`. This pair
contributes one inversion exactly when

```text
a_i * x > a_j * y
```

Since all values are positive, this can be rearranged without changing the
inequality direction:

```text
y / x < a_i / a_j
```

So each index pair `(i, j)` asks: among all ordered choices of two distinct
modules `(x, y)`, how many ratios `y / x` are strictly smaller than
`a_i / a_j`?

There are `n * (n - 1)` ordered choices for `(x, y)`. In a uniformly random
permutation, every ordered pair of distinct original modules has exactly this
same probability of occupying positions `(i, j)`, so the contribution of this
index pair is:

```text
count(y / x < a_i / a_j) / (n * (n - 1))
```

By linearity of expectation, the final answer is the sum of this value over all
`i < j`.

The implementation builds two sorted ratio lists:

- `arr1`: all `a_i / a_j` for `i < j`
- `arr2`: all `b_q / b_p` for ordered pairs `p != q`

Ratios are stored as reduced numerator/denominator pairs and compared by cross
multiplication:

```text
u / v < x / y  <=>  u * y < x * v
```

After sorting both lists, a two-pointer scan counts, for every ratio in
`arr1`, how many ratios in `arr2` are strictly smaller. This total count is the
sum of all numerators above. Multiplying it by the modular inverse of
`n * (n - 1)` gives the expected value modulo `998244353`.

### Correctness

For a fixed pair `i < j`, an inversion happens exactly when
`a_i * x > a_j * y`, where `x` and `y` are the modules assigned to positions
`i` and `j`. Because all numbers are positive, this is equivalent to
`y / x < a_i / a_j`.

In a random permutation, the ordered module pair `(x, y)` at positions `(i, j)`
is uniformly distributed over all `n * (n - 1)` ordered pairs of distinct
modules. Therefore, the expected contribution of `(i, j)` is precisely the
number of ordered module ratios `y / x` smaller than `a_i / a_j`, divided by
`n * (n - 1)`.

The list `arr1` contains exactly one threshold ratio for every possible index
pair `i < j`. The list `arr2` contains exactly one ratio `y / x` for every
possible ordered choice of distinct modules. The sorted two-pointer scan counts
exactly the number of `arr2` ratios strictly smaller than each `arr1` ratio.
Thus the accumulated count is the total numerator of the expected inversion
sum over all pairs `i < j`.

Finally, multiplying by `(n * (n - 1))^{-1}` modulo `998244353` performs the
required division in modular arithmetic. Hence the algorithm returns the
expected number of inversions.

### Complexity

There are `O(n^2)` ratios in both lists. Sorting dominates the runtime, so the
time complexity is `O(n^2 log n)`, and the memory complexity is `O(n^2)`.
