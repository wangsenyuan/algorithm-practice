# C. MEXOR

[Problem link](https://codeforces.com/problemset/problem/2245/C)

**Contest:** [Spectral::Cup 2026 Round 3 (Codeforces Round 1110, Div. 1 + Div. 2)](https://codeforces.com/contest/2245)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

## Problem

You are given a positive integer `n` and a non-negative integer `k`. Construct a
permutation `p` of length `n` such that the following condition is satisfied:

- Let `f(i) = mex([p_0, p_1, …, p_i])` for every `0 ≤ i < n`. Then
  `f(0) ⊕ f(1) ⊕ … ⊕ f(n-1)` should equal `k`, where `⊕` denotes bitwise XOR.

A permutation of length `n` is an array of `n` distinct integers from `0` to
`n-1` in arbitrary order. For example, `[1, 2, 0, 4, 3]` is a permutation, but
`[0, 1, 1]` is not (duplicate), and `[0, 2, 3]` is not (`n = 3` but `3` appears).

The minimum excluded (MEX) of a collection of integers is the smallest
non-negative integer that does not occur in the collection.

## Constraints

- `1 ≤ t ≤ 10^4`
- `1 ≤ n ≤ 2 · 10^5`
- `0 ≤ k ≤ 10^9`
- Sum of `n` over all test cases does not exceed `2 · 10^5`

## Input

The first line contains the number of test cases `t`.

The only line of each test case contains two integers `n` and `k`.

## Output

For each test case, if no such permutation exists, output `NO`.

Otherwise, first output `YES` on a single line. Then output `n` distinct
integers `p_0, p_1, …, p_{n-1}` (`0 ≤ p_i < n`).

If there exist multiple permutations satisfying the requirement, you may output
any of them.

You can output the answer in any case (upper or lower).

## Example

### Input

```text
6
1 0
1 1
3 0
4 8
5 1
9 12
```

### Output

```text
NO
YES
0
YES
0 1 2
NO
YES
3 0 2 1 4
YES
1 4 0 8 2 3 7 5 6
```

## Note

In the first and second test cases, the only permutation of length `1` is
`[0]`, with `f(0) = mex([0]) = 1`.

In the fourth test case, it can be proven that no permutation of length `4`
exists such that `f(0) ⊕ f(1) ⊕ f(2) ⊕ f(3) = 8`.

In the fifth test case, the values of `f(i)` for `0 ≤ i < n` are:

- `f(0) = mex([3]) = 0`
- `f(1) = mex([3, 0]) = 1`
- `f(2) = mex([3, 0, 2]) = 1`
- `f(3) = mex([3, 0, 2, 1]) = 4`
- `f(4) = mex([3, 0, 2, 1, 4]) = 5`

Since `0 ⊕ 1 ⊕ 1 ⊕ 4 ⊕ 5 = 1`, `p = [3, 0, 2, 1, 4]` is valid.

## ideas
1. f(n-1) = n always
2. bit by bit

## Solution

The full permutation has MEX `n`, so `f(n-1) = n` always. The target on
the earlier prefix MEXes is therefore `k ⊕ n`. A prefix that does not
contain `0` has MEX `0`, and XOR with `0` does nothing, so those positions
are free.

A MEX value that stays for an even number of steps cancels. To avoid that,
put `0` immediately before a short increasing suffix of “holes.” After `0`
appears, MEX jumps through those holes one by one (each used once), then
finishes at `n`.

Let `x = k ⊕ n`. The earlier XOR is then exactly the XOR of the holes:

- If `x < n`, one hole: `x`.
- If `x ≥ n`, peel the highest bit `w = 2^{⌊log₂ x⌋}`. Then `x = r ⊕ w`
  with `r < w`. If `w ≥ n`, that bit cannot appear as a prefix MEX → `NO`.
  Otherwise the holes are `r` then `w`.

Fill the prefix before `0` with the unused values in increasing order so
that `1, 2, …, r-1` are already present when `0` is placed. MEX then jumps
`r → (w) → n` as required.

`x` is possible iff it uses only bits that appear in `0..n-1` (equivalently
`(k ⊕ n)` has no bit `≥ 2^{⌊log₂(n-1)⌋+1}`). The construction realizes
every such `x`: one hole when `x < n`, two holes `r, w` otherwise. A final
scan of prefix-MEX XOR is only a check; it is not part of the argument.
