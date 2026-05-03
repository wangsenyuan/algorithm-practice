# D2. Inversion Graph Coloring (Hard Version)

This is the **hard** version; here `n ≤ 2000` (see also D1 for the easier constraints).

A sequence `b_1, b_2, …, b_k` is **good** if you can color each index `i` **red** or **blue** so that: for every pair `i < j` with `b_i > b_j` (an inversion), indices `i` and `j` receive **different** colors.

You are given a sequence `a_1, a_2, …, a_n`. Compute the number of **good subsequences** of `a` (including the empty subsequence), modulo `10^9 + 7`.

> **\*** A sequence `b` is a subsequence of `a` if `b` can be obtained from `a` by deleting zero or more elements from arbitrary positions.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 ≤ t ≤ 100`). The description of the test cases follows.

For each test case:

- The first line contains an integer `n` (`1 ≤ n ≤ 2000`) — the length of the sequence.
- The second line contains `n` integers `a_1, a_2, …, a_n` (`1 ≤ a_i ≤ n`).

It is guaranteed that the sum of `n` over all test cases does not exceed **2000**.

## Output

For each test case, output a single integer: the number of good subsequences modulo `10^9 + 7`.

## Example

**Input**

```text
4
4
4 2 3 1
7
7 6 1 2 3 3 2
5
1 1 1 1 1
11
7 2 1 9 7 3 4 1 3 5 3
```

**Output**

```text
13
73
32
619
```

## Note

In the first test case, the subsequences that are **not** good are `[4, 3, 1]`, `[4, 2, 1]`, and `[4, 2, 3, 1]`. There are `16` subsequences in all, so `16 − 3 = 13` are good.

In the third test case, every subsequence is good.
